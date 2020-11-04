package services

import (
	"fmt"
	"github.com/vaiktorg/portfolio/src/services/notifs"
	"io"
	"io/ioutil"
	"net/http"
)

type HTTPReqs struct {
	client http.Client
}

func (h *HTTPReqs) GET(addr string, handler func(data []byte)) {
	resp, err := h.client.Get(addr)
	if err != nil {
		notifs.NotifError("", err)
		return
	}

	go h.processHandler(resp, handler)
}

func (h *HTTPReqs) POST(addr string, contentType string, body io.Reader, handler func(data []byte)) {
	resp, err := h.client.Post(addr, contentType, body)
	if err != nil {
		notifs.NotifError("shit wont work", err)
		return
	}

	go h.processHandler(resp, handler)
}

func (h *HTTPReqs) processHandler(resp *http.Response, handler func(data []byte)) {
	if resp.Header.Get("Content-Type") == "text/html" {
		notifs.NotifError("html not allowed on httpReqs", nil)
		return
	}
	if resp.StatusCode != http.StatusOK {
		notifs.NotifError(fmt.Sprintf("GetReq came back with status code: %v\n", resp.StatusCode), nil)
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		notifs.NotifError("Could not read response body", err)
		return
	}
	defer resp.Body.Close()

	handler(data)
}
