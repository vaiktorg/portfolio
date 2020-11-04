package fileserver

import (
	"fmt"
	"time"

	"github.com/vaiktorg/portfolio/src/services"

	. "github.com/maxence-charriere/go-app/pkg/app"
)

type Index struct {
	Compo
	client services.HTTPReqs
}

func (d *Index) Render() UI {
	return Div().Body(
		H1().Text("This is a fileserverui"),
		Input().Class("w3-hide").
			Name("Upload").
			Type("file").
			ID("file").
			Multiple(true).
			Accept("image/*").
			OnChange(d.OnChange),

		Button().Class("w3-button w3-margin w3-black").Text("Upload").OnClick(d.OpenFileDialog),

		Div().Class("w3-container").Body(
			Table().Body(
				THead().Body(
					Tr().Body(
						Th().Text("Name"),
						Th().Text("Timestamp"),
						Th().Text("CreatedBy"),
						Th().Text("Size"),
					).Class("w3-black"),
				),
				Tr().Body(
					Td().Text("bhdjksd.jpg"),
					Td().Text(time.Now().String()),
					Td().Text("Bebo"),
					Td().Text("30MB"),
				),
			).Class("w3-table-all w3-hoverable"),
		),
	).
		Class("w3-container w3-padding").
		OnContextMenu(d.OnRightClick).Style("width", "100%")
}

func (d *Index) OnRightClick(_ Context, ev Event) {
	ev.PreventDefault()
	NewContextMenu(
		MenuItem().Label("Upload"),
	)
}

func (d *Index) OnChange(Context, Event) {
	fmt.Println("FILES")

	//files := Window().GetElementByID("file").Get("files")
	//reader := ctx.JSSrc.New("FileReader")
	//
	//fmt.Println(buff)
	//services.HTTPReqs{}.POST(filepath.Join(fileserverAddr, "users", "upload"),"image/*",buff, )
}

func (d *Index) OpenFileDialog(_ Context, e Event) {
	e.PreventDefault()
	Window().GetElementByID("file").Call("click")
}
