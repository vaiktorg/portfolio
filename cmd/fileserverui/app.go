package main

import (
	"github.com/maxence-charriere/go-app/pkg/app"
	"github.com/vaiktorg/portfolio/src/pages"
	"github.com/vaiktorg/portfolio/src/pages/fileserver"
)

// TODO: Make everything mobile/responsive.

func main() {

	// FileServerUI
	fs := &fileserver.Index{}

	// MiscPages
	e := &pages.ErrorPage{}

	// Route Registration
	app.Route("/", fs)
	app.Route("/error", e)

	app.Run()
}
