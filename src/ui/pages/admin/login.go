package admin

import (
	"time"

	"github.com/vaiktorg/grimoire/helpers"

	"github.com/maxence-charriere/go-app/pkg/app"
)

type LoginPage struct {
	app.Compo
	loggedIn bool
	password string
	uname    string
	msgs     []string
}

func (l *LoginPage) Render() app.UI {
	return app.Div().Body(
		app.H1().Text("HEPTACODE ACCESS"),
		app.Hr(),
		app.Label().Text("Username | Email").Class("w3-text-black w3-padding-top"),
		app.Input().Type("text").Placeholder("Email or Password").
			Class("w3-input w3-border").
			OnChange(l.onUsernameChange),

		app.Br(),
		app.Label().Text("Password").Class("w3-text-black w3-padding-top"),
		app.Input().Type("text").Placeholder("Password").
			Class("w3-input w3-border").
			OnChange(l.onPasswordChange),

		app.If(len(l.msgs) > 0,
			app.Range(l.msgs).Slice(func(i int) app.UI {
				return app.P().Text("* " + l.msgs[i]).Class("w3-animate-opacity")
			}),
		),

		app.Br(),
		app.Input().Type("Submit").Value("LogIn").
			Class("w3-btn w3-black").
			OnClick(l.onSubmit),
	).Class("container w3-animate-opacity").
		Style("padding", "10% 30%")

}

func (l *LoginPage) onSubmit(ctx app.Context, _ app.Event) {
	isUnameValid := helpers.ValidateUsername(l.uname)
	l.msgs = helpers.ValidatePassword(l.password)

	if !isUnameValid {
		l.msgs = append(l.msgs, "invalid username")
	}

	if len(l.msgs) > 0 {
		l.Update()
		return
	}

	l.msgs = append(l.msgs, "Welcome!")

	time.AfterFunc(time.Second, func() {
		app.Navigate("/dashboard")
	})

	l.loggedIn = true

	l.Update()
}

func (l *LoginPage) onPasswordChange(ctx app.Context, _ app.Event) {
	l.password = ctx.JSSrc.Get("value").String()
}

func (l *LoginPage) onUsernameChange(ctx app.Context, _ app.Event) {
	l.uname = ctx.JSSrc.Get("value").String()
}
