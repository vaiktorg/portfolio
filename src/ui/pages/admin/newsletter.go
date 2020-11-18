package admin

import (
	"time"

	"github.com/maxence-charriere/go-app/pkg/app"
	"github.com/vaiktorg/grimoire/helpers"
)

type NewsletterPage struct {
	app.Compo
	email string
	msg   string
}

func (h *NewsletterPage) Render() app.UI {
	return app.Div().Body(
		app.Div().Body(
			app.Img().Src("https://i.imgur.com/npQE0TM.png").Style("margin", "64px auto"),
			app.H1().Text("HEPTACODE COLLECTIVE").Style("margin-bottom", "0px"),
			app.H5().Text("The masters are those who understand the source."),
			app.B(),
			app.P().Text("Would you dare to learn the truth?"),
		).Class("w3-center"),

		app.Div().Body(
			app.If(h.email != "", app.P().Text(h.msg)),

			app.Input().
				Value(h.email).
				Placeholder("Email").
				AutoFocus(true).
				OnChange(h.OnInputChange),

			app.Input().Type("button").Value("Submit").OnClick(h.OnClick),
		).Class("w3-center"),
	).Class("w3-animate-opacity w3-animate-opacity")

}

func (h *NewsletterPage) OnInputChange(ctx app.Context, _ app.Event) {
	h.email = ctx.JSSrc.Get("value").String()
}

func (h *NewsletterPage) OnClick(ctx app.Context, _ app.Event) {
	if helpers.ValidateEmail(h.email) {
		h.msg = "We have sent you an email to: " + h.email
		h.Update()
		time.AfterFunc(time.Second, func() {
			app.Navigate("http://google.com/")
		})
		return
	}

	h.msg = "Invalid email"
	h.Update()

}
