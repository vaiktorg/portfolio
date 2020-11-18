package components

import "github.com/maxence-charriere/go-app/pkg/app"

type NotifBar struct {
	app.Compo
	err   []string
	stats []string
}

func (n *NotifBar) Render() app.UI {
	return app.Div().Body(
		app.Div().Body(app.Text("Message: ")).Class("w3-third"),
		app.Div().Body(app.Text("Errors:")).Class("w3-third"),
	).Class("w3-container w3-row w3-bar w3-black w3-bottom w3-animate-opacity")

}

func (n *NotifBar) NotifError(msg string, err error) {

}
