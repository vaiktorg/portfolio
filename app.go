package main

import (
	"github.com/maxence-charriere/go-app/pkg/app"
	"github.com/vaiktorg/portfolio/src/input"
	"github.com/vaiktorg/portfolio/src/pages"
)

func main() {

	dash := &pages.DashboardPage{}
	config := &input.YamlInput{}

	dash.SetContent(&pages.HomePage{})
	config.Init()

	navbar := dash.GetNavBar()
	navbar.RegisterPage("home", &pages.HomePage{})
	navbar.RegisterPage("tools", &pages.ToolsPage{})
	navbar.RegisterPage("shuffler", &pages.GUIDCardShuffler{})

	app.Route("/", dash)
	app.Route("/dashboard", dash)

	app.Run()
}
