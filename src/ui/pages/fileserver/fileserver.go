package fileserver

import (
	"fmt"
	"time"

	. "github.com/maxence-charriere/go-app/pkg/app"
)

type FileServerPage struct {
	Compo
}

var (
	fileserverAddr = "10.0.0.94:2000"
)

func (d *FileServerPage) Render() UI {
	// Title
	return Div().Class("w3-padding-64").Body(
		Div().Class("w3-bar w3-black").Body(
			Button().Class(" w3-bar-item w3-button w3-right").Body(I().Class("mdi mdi-login-variant")),
			Label().Class("w3-bar-item").Text("Welcome! [Username]"),
		).Style("top", "0").
			Style("position", "fixed"),

		Div().Class("w3-container w3-bar w3-center").Body(
			H1().Class("w3-left").Text("Vaiktorg's FileServer"),
		),
		// Directory options
		Div().Class("w3-container w3-xlarge w3-bar").Body(

			Button().Class(" w3-bar-item w3-button w3-right").Body(I().Class("mdi mdi-folder-upload")).OnClick(d.OpenFileDialog),
			Button().Class(" w3-bar-item w3-button w3-right").Body(I().Class("mdi mdi-folder-plus")),

			Button().Class(" w3-bar-item w3-button w3-left").Body(I().Class("mdi mdi-home")),
			Button().Class(" w3-bar-item w3-button w3-left").Body(I().Class("mdi mdi-keyboard-backspace")),
		),

		// Upload Input
		Input().Class("w3-hide").
			Name("Upload").
			Type("file").
			ID("file").
			Multiple(true).
			Accept("image/*").
			OnChange(d.OnChange),

		Div().Class("w3-container w3-bar").Body(
			Button().Class(" w3-button").Body(Small().Text("home")),
			Button().Class(" w3-button").Body(Small().Text("NotPr0n")),
			Button().Class(" w3-button").Body(Small().Text("seriously_dont")),
		),
		// Directory files
		Div().Class("w3-container ").Body(
			// Crumbs path

			Div().Class("w3-hide-medium w3-hide-large").Body(
				card("Pr0n.mp4", "/root/NotMyPr0n/pr0n", "63MB"),
				card("Pr0n.mp4", "/root/NotMyPr0n/pr0n", "63MB"),
				card("Pr0n.mp4", "/root/NotMyPr0n/pr0n", "63MB"),
			),

			// Table of files
			Table().Class("w3-hide-small w3-table-all w3-black w3-mobile w3-centered").Body(
				THead().Body(
					Tr().Body(
						Th().Text("name"),
						Th().Text("timestamp").Class(""),
						Th().Text("CreatedBy"),
						Th().Text("Size"),
						Th().Text(""),
					).Class("w3-black"),
				),
				// This should be a "Range(elems...)"
				Tr().Body(
					Td().Text("clownPr0n.mp4"),
					Td().Text(time.Now().Format("02 Jan 06 15:04 MST")),
					Td().Text("Bebo"),
					Td().Text("30MB"),
					Td().Class("w3-button w3-large").Body(I().Class("mdi mdi-delete")),
				),
			).Class("w3-table-all w3-hoverable"),

			Div().Class("w3-margin-top w3-bar w3-center").Body(
				A().Class("w3-button").Text("«"),
				A().Class("w3-button").Text("1"),
				A().Class("w3-button").Text("2"),
				A().Class("w3-button").Text("3"),
				A().Class("w3-button").Text("4"),
				A().Class("w3-button").Text("»"),
			),
		),
	)
}

func (d *FileServerPage) OnChange(ctx Context, _ Event) {
	fmt.Println("BITCHES")

	/*	let photo = document.getElementById("image-file").files[0];
		let formData = new FormData();

		formData.append("photo", photo);
		fetch('/upload/image', {method: "POST", body: formData});*/

	//files := Window().GetElementByID("file").Get("files")
	//fmt.Printf("%T\n", files)
	//var buff []byte
	//CopyBytesToGo(buff, files)

	//rdr := bytes.NewReader(buff)

}

func (d *FileServerPage) OpenFileDialog(_ Context, e Event) {
	e.PreventDefault()
	Window().GetElementByID("file").Call("click")
}

func card(filename, path, size string) UI {
	return Div().Class("w3-card-2 w3-margin-bottom").Body(
		Header().Class("w3-bar w3-black").Body(
			Button().Class("w3-bar-item w3-btn w3-right").Body(I().Class("mdi mdi-delete")),
			Button().Class("w3-bar-item w3-btn w3-black").Text(filename).Style("margin", "auto"),
		),
		Div().Class("w3-bar w3-padding").Body(
			Small().Class("w3-left").Text(path),
			Small().Class("w3-right").Text("Size: "+size),
		),
	)
}
