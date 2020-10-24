package pages

import (
	. "github.com/maxence-charriere/go-app/pkg/app"
	"github.com/vaiktorg/portfolio/src/models"
)

type ToolsPage struct {
	Compo
	posts []models.PostCard
}

func (t *ToolsPage) Render() UI {

	return Div().Body(
		H1().Text("Tools"),
		P().Text("Random set of casual tools"),
		Hr(),

		Range(
			Div().Body(
				Img().Src("web/static/gwarana.png").Class("w3-image"),
				Div().Body(
					H3().Text("GWARANA"),
					P().Text("This is a tool for keeping your pc awake by simulating pressing keys"),
				).Class("w3-container w3-padding"),
			).Class("w3-third w3-center w3-padding"),
		),
	).Class("w3-row w3-animate-opacity w3-center ").ID("content")
}

//func (t *ToolsPage) tools() map[string]Post {
//
//}
