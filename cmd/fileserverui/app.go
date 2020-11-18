package main

import (
	"github.com/maxence-charriere/go-app/pkg/app"
	"github.com/vaiktorg/portfolio/src/ui/pages"
	"github.com/vaiktorg/portfolio/src/ui/pages/fileserver"
)

func main() {

	// FileServerUI
	fs := &fileserver.FileServerPage{}

	// MiscPages
	e := &pages.ErrorPage{}

	// Route Registration
	app.Route("/", fs)
	app.Route("/error", e)

	app.Run()
}
