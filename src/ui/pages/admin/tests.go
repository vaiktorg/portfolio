package admin

import (
	"fmt"

	. "github.com/maxence-charriere/go-app/pkg/app"
	"github.com/vaiktorg/portfolio/src/ui/services"
)

type TestPage struct {
	Compo
	client services.HTTPReqs
}

func (t *TestPage) Render() UI {
	return Div().Body(Input().Type("button").Value("HTTPReq").OnClick(t.OnClickHTTPRequest))
}

func (t *TestPage) OnClickHTTPRequest(Context, Event) {
	//t.client.GET("https://gist.githubusercontent.com/Vaiktorg/ecee0dca3da073517487a0b2049d6c2d/raw/ad835e1c49e77331c875d5fa000e6b946ac83545/Website",
	t.client.GET("",
		func(data []byte, err error) {
			fmt.Println(string(data))
			fmt.Println("FUCK")
		})
}
