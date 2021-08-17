package notesui

import (
	"fmt"

	. "github.com/maxence-charriere/go-app/pkg/app"
	"github.com/vaiktorg/portfolio/src/ui/components"
	"github.com/vaiktorg/portfolio/src/ui/services"
)

type NotesPage struct {
	Compo
	tree components.Tree

	DirectoryVisible string
	DirActive        bool

	EditorVisible string
	EditorActive  bool

	PreviewVisible string
	PreviewActive  bool

	client services.HTTPReqs
}

func (n *NotesPage) OnMount(Context) {
	n.DirectoryVisible = "w3-show"
	n.EditorVisible = "w3-hide"
	n.PreviewVisible = "w3-show"

	n.client.GET("/s/vaiktorg/init", func(data []byte, err error) {
		fmt.Println(string(data))
	})

	n.Update()
}

func (n *NotesPage) Render() UI {
	return Div().Body(
		// NotifBar
		Div().Class("w3-bar w3-black").ID("MenuBar").Body(
			Button().Class("w3-button w3-bar-item "+n.ActiveBtn(n.DirectoryVisible)).Text("Directory").OnClick(n.ToggleDirectory),
			Button().Class("w3-button w3-bar-item "+n.ActiveBtn(n.EditorVisible)).Text("Editor").OnClick(n.ToggleEditor),
			Button().Class("w3-button w3-bar-item "+n.ActiveBtn(n.PreviewVisible)).Text("Preview").OnClick(n.TogglePreview),
		),
		Hr(),
		Div().Class("").Body(
			Div().Class("w3-row").Body(
				Div().Class("w3-half w3-container "+n.DirectoryVisible).Body(H5().Text("Directory...")),
			),

			Div().Class("w3-row").ID("NotesEditor").Body(
				Div().Class("w3-half w3-padding "+n.EditorVisible).Body(
					Div().Class("w3-bar").Body(
						Input().Class("w3-mobile w3-posts w3-border").Placeholder("Note Title..."),
						Button().Class("w3-bar-item w3-button").Body(I().Class("mdi mdi-format-bold")),
						Button().Class("w3-bar-item w3-button").Body(I().Class("mdi mdi-format-italic")),
						Button().Class("w3-bar-item w3-button").Body(I().Class("mdi mdi-link")),
						Button().Class("w3-bar-item w3-button").Body(I().Class("mdi mdi-code-braces")),
						Button().Class("w3-bar-item w3-button").Body(I().Class("mdi mdi-format-list-bulleted")),
						Button().Class("w3-bar-item w3-button").Body(I().Class("mdi mdi-file-image")),
					),
					Div().ContentEditable(true).Class("editor").ID("editor").Text("this is some shit because its shit."),
				),
				Div().Class("w3-rest w3-padding "+n.PreviewVisible).Body(H5().Text("Preview...")),
			),
		),
	)
}

func (n *NotesPage) ToggleDirectory(Context, Event) {
	n.DirectoryVisible = n.Toggle(n.DirectoryVisible)
	n.Update()
}
func (n *NotesPage) ToggleEditor(Context, Event) {
	n.EditorVisible = n.Toggle(n.EditorVisible)
	n.Update()
}
func (n *NotesPage) TogglePreview(Context, Event) {
	n.PreviewVisible = n.Toggle(n.PreviewVisible)
	n.Update()
}

func (NotesPage) Toggle(class string) string {
	hide := "w3-hide"
	show := "w3-show"

	switch class {
	case show:
		return hide
	case hide:
		return show
	}

	return show
}

func (NotesPage) ActiveBtn(class string) string {
	hide := "w3-hide"

	if class == hide {
		return ""
	}
	return "w3-active"
}
