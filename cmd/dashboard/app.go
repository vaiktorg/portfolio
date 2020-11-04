package main

import (
	"github.com/maxence-charriere/go-app/pkg/app"
	"github.com/vaiktorg/portfolio/src/pages"
	"github.com/vaiktorg/portfolio/src/pages/admin"
	"github.com/vaiktorg/portfolio/src/pages/fileserver"
)

// TODO: Make everything mobile/responsive.

func main() {

	// Dashboard Portfolio
	dash := &admin.DashboardPage{}
	dash.SetContent(&admin.HomePage{})

	navbar := dash.GetNavBar()
	navbar.RegisterPage("home", &admin.HomePage{})
	navbar.RegisterPage("tools", &admin.ToolsPage{})
	navbar.RegisterPage("test", &admin.TestPage{})
	navbar.RegisterPage("shuffler", &admin.GUIDCardShuffler{})

	// FileServerUI
	fs := &fileserver.Index{}

	// MiscPages

	// Route Registration
	app.Route("/admin", dash)
	app.Route("/fileserverui", fs)
	app.Route("/error", &pages.ErrorPage{})

	app.Run()
}
