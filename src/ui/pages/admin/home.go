package admin

import (
	. "github.com/maxence-charriere/go-app/pkg/app"
)

type HomePage struct {
	Compo
	email string
	msg   string
	links map[string]string
}

func (h *HomePage) Render() UI {
	return Div().Body(
		Div().Body(
			H1().Text("WELCOME TO HEPTACODE"),
			P().Text("Personal project showcasing tools, tutorials and more."),
		).Class("w3-center"),
		Hr(),
		Div().Body(
			H3().Text("About").Class("w3-center"),
			P().Text("go-app is a package to build progressive web apps (PWA) with Go programming language and WebAssembly."),
			A().Text("Check Out Go-app").Href("https://github.com/maxence-charriere/go-app"),
			Ul().Body(
				Li().Text("Language is Golang"),
				Li().Text("Exported for WebAssembly (GOARCH = wasm | GOOS = js)"),
			),
		).Class("w3-center").
			Style("margin-top", "64px").
			Style("margin-bottom", "64px"),
		Hr(),
		Div().Class("w3-jumbo w3-center").Body(
			Range(h.links).Map(
				func(s string) UI {
					return A().Href(h.links[s]).Body(I().Class(s)).Class("w3-mobile w3-margin")
				},
			),
		),
		H5().Text("Made By Vaiktorg").Class("w3-center"),
	).Class("w3-animate-opacity").ID("content")
}

func (h *HomePage) OnMount(Context) {
	h.links = map[string]string{
		"mdi mdi-github":   "https://github.com/vaiktorg",
		"mdi mdi-facebook": "https://www.facebook.com/TheOnlyVaiktorg/",
		"mdi mdi-linkedin": "https://www.linkedin.com/in/victorhernandezcandelaria/",
	}
	h.Update()
}

func (h *HomePage) OnInputChange(ctx Context, _ Event) {
	h.email = ctx.JSSrc.Get("value").String()
}

func (h *HomePage) OnClick(Context, Event) {

}
