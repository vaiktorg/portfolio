package pages

import (
	"github.com/maxence-charriere/go-app/pkg/app"
	"github.com/vaiktorg/portfolio/src/components"
	"github.com/vaiktorg/portfolio/src/models"
)

type DashboardPage struct {
	app.Compo
	content  app.Composer
	sidebar  components.NavBar
	LoggedIn models.User
}

func (d *DashboardPage) Render() app.UI {
	return app.Div().Body(
		&d.sidebar,
		d.content,
	).Class("w3-row").Style("margin-bottom", "5%")
}

func (d *DashboardPage) OnRightClick(_ app.Context, ev app.Event) {
	app.NewContextMenu(
		app.MenuItem().Label("TEST"),
	)
	ev.PreventDefault()
}

func (d *DashboardPage) OnMount(app.Context) {
	d.sidebar.SetContentChanger(func(comp app.Composer) {
		d.content = comp
		d.Update()
	})
}

// Functionality
//======================================================================================================================
func (d *DashboardPage) GetContent() app.UI {
	return d.content
}

func (d *DashboardPage) SetContent(comp app.Composer) {
	d.content = comp
}

func (d *DashboardPage) GetNavBar() *components.NavBar {
	return &d.sidebar
}

// Events
//======================================================================================================================

func (d *DashboardPage) OnContentChange(ev *components.ContentChangeEvent) {
	d.content = ev.Content
	d.Update()
}
