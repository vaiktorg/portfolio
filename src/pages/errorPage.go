package pages

import (
	. "github.com/maxence-charriere/go-app/pkg/app"
	"github.com/vaiktorg/portfolio/src/services"
)

type ErrorPage struct {
	Compo
	err        error
	ErrManager services.ErrorManager
}

func (e *ErrorPage) OnMount(Context) {
	e.ErrManager.OnError(func(err error) {
		e.err = err
		e.Update()
	})
}

func (e *ErrorPage) Render() UI {
	return Div().Class("w3-container w3-center").Body(
		H1().Text("Shit, Something Blew Up..."),
		Div().Class("w3-container").Body(
			Pre().Text(e.err.Error()),
		),
	)
}
