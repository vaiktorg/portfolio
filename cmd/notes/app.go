package main

import (
	"github.com/maxence-charriere/go-app/pkg/app"
	"github.com/vaiktorg/portfolio/src/ui/pages"
	"github.com/vaiktorg/portfolio/src/ui/pages/notesui"
)

func main() {

	// FileServerUI
	ns := &notesui.NotesPage{}

	// MiscPages
	e := &pages.ErrorPage{}

	// Route Registration
	app.Route("/", ns)
	app.Route("/error", e)

	app.Run()
}
