package services

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type HTTPReqs struct {
	http.Client
}

func (h *HTTPReqs) GET(addr string, handler func([]byte, error)) {
	resp, err := h.Get(addr)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	go h.processHandler(resp, handler)
}

func (h *HTTPReqs) POST(addr string, contentType string, body io.Reader, handler func(data []byte, err error)) {
	resp, err := h.Post(addr, contentType, body)
	if err != nil {
		fmt.Println("shit wont work", err)
		return
	}

	go h.processHandler(resp, handler)
}

func (h *HTTPReqs) processHandler(resp *http.Response, handler func(data []byte, err error)) {
	if resp.Header.Get("Content-Type") == "text/views" {
		fmt.Println("views not allowed on httpReqs", nil)
		return
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println(fmt.Sprintf("GetReq came back with status code: %v\n", resp.StatusCode), nil)
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Could not read response body", err)
		return
	}
	defer resp.Body.Close()

	handler(data, nil)
}
