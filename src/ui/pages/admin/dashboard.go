package admin

import (
	. "github.com/maxence-charriere/go-app/pkg/app"
	"github.com/vaiktorg/portfolio/src/ui/components/navigation"
	"github.com/vaiktorg/portfolio/src/ui/components/notifs"
)

type DashboardPage struct {
	Compo
	content       Composer
	sidebar       navigation.NavBar
	notifications notifs.NotificationComponent
}

func (d *DashboardPage) Render() UI {
	return Div().Body(
		&d.sidebar,
		d.content,
		&d.notifications,
	).Class("w3-row").Style("margin-bottom", "5%")
}

func (d *DashboardPage) OnRightClick(_ Context, ev Event) {
	NewContextMenu(
		MenuItem().Label("TEST"),
	)
	ev.PreventDefault()
}

func (d *DashboardPage) OnMount(Context) {
	d.sidebar.SetContentChanger(func(comp Composer) {
		d.content = comp
		d.Update()
	})
}

// Functionality
//======================================================================================================================
func (d *DashboardPage) GetContent() UI {
	return d.content
}

func (d *DashboardPage) SetContent(comp Composer) {
	d.content = comp
}

func (d *DashboardPage) GetNavBar() *navigation.NavBar {
	return &d.sidebar
}

// Events
//======================================================================================================================

func (d *DashboardPage) OnContentChange(ev *navigation.ContentChangeEvent) {
	d.content = ev.Content
	d.Update()
}
