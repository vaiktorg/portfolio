package main

import (
	"github.com/maxence-charriere/go-app/pkg/app"
	"github.com/vaiktorg/portfolio/src/ui/pages"
	"github.com/vaiktorg/portfolio/src/ui/pages/admin"
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

	// Route Registration
	app.Route("/", dash)
	app.Route("/error", &pages.ErrorPage{})

	app.Run()
}
