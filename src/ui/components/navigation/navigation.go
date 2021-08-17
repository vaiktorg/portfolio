package navigation

import (
	"strings"
	"sync"

	"github.com/maxence-charriere/go-app/pkg/app"
)

type NavBar struct {
	app.Compo
	once           sync.Once
	CurrentPage    string
	pages          map[string]app.Composer
	contentChanger func(compo app.Composer)
}

type ContentChangeEvent struct {
	Content app.Composer
}

func (n *NavBar) Render() app.UI {
	//	return app.Raw(`
	//<div class="w3-sidebar w3-bar-block w3-black" style="width:auto">
	//  <a href="#" class="w3-bar-item w3-button"><i class="fa fa-home"></i></a>
	//  <a href="#" class="w3-bar-item w3-button"><i class="fa fa-search"></i></a>
	//  <a href="#" class="w3-bar-item w3-button"><i class="fa fa-envelope"></i></a>
	//  <a href="#" class="w3-bar-item w3-button"><i class="fa fa-globe"></i></a>
	//</div>
	//`)
	return app.Div().Body(
		app.Img().Src("https://i.imgur.com/nRDh6ph.png").Class("w3-mobile w3-image brand-logo color-invert"),

		app.Range(n.pages).Map(func(k string) app.UI {
			return app.Input().
				Type("button").
				Value("/" + k).
				Src("/" + k).
				Class("w3-bar-item w3-btn w3-mobile w3-margin").
				OnClick(n.OnClick)
		}),
	).Class("w3-animate-opacity w3-black").
		Style("padding-left", "20%").
		Style("padding-right", "20%").
		ID("models")
}

func (n *NavBar) OnMount(app.Context) {
	n.CurrentPage = "Welcome"
	n.Update()
}

func (n *NavBar) OnClick(ctx app.Context, _ app.Event) {
	n.CurrentPage = strings.Trim(ctx.JSSrc.Get("value").String(), "/")
	n.contentChanger(n.pages[n.CurrentPage])
	n.Update()
}

func (n *NavBar) RegisterPage(name string, page app.Composer) {
	if n.pages == nil {
		println("ages init")
		n.pages = map[string]app.Composer{}
	}

	if _, ok := n.pages[name]; !ok {
		println("added page")
		n.pages[name] = page
	}
}

func (n *NavBar) SetContentChanger(contentChanger func(comp app.Composer)) {
	n.contentChanger = contentChanger
}
