package subcompo

import (
	. "github.com/maxence-charriere/go-app/pkg/app"
)

type AccordionCompo struct {
	Compo
	isHidden     bool
	ButtonClass  string
	contentClass string
	ButtonText   string
	ContentBody  []UI
}

func (a *AccordionCompo) Render() UI {
	return Div().Body(
		Button().Class(a.ButtonClass).Text(a.ButtonText).OnClick(a.ToggleNotifs),
		Div().Class(a.contentClass+" w3-padding").Body(
			a.ContentBody...,
		),
	)
}

func (a *AccordionCompo) OnMount(Context) {
	a.contentClass = "w3-hide"
	a.isHidden = true
	a.Update()
}

func (a *AccordionCompo) ToggleNotifs(Context, Event) {
	if a.isHidden {
		a.contentClass = "w3-show"
	}
	if !a.isHidden {
		a.contentClass = "w3-hide"
	}
	a.isHidden = !a.isHidden
	a.Update()
}
