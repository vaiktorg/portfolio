package notifs

import (
	"fmt"
	"sync"

	. "github.com/maxence-charriere/go-app/pkg/app"
	"github.com/vaiktorg/grimoire/uid"
	"github.com/vaiktorg/portfolio/src/ui/services/notifs"
)

type NotificationComponent struct {
	Compo
	center    *notifs.NotificationCenter
	m         sync.Mutex
	notifs    []notifs.Notif
	isHidden  bool
	showClass string
}

func (n *NotificationComponent) Render() UI {
	return Div().Body(
		Range(n.notifs).Slice(func(i int) UI {
			return Card(
				n.notifs[i].ID,
				n.notifs[i].Typ,
				n.notifs[i].Msg,
				n.notifs[i].Error,
				n.notifs[i].Data,
			)
		}),
	).Class("w3-row")
	// TODO: makefile notification container have better UX.

}

func Card(id uid.UID, typ notifs.Type, msg string, trace string, data interface{}) UI {
	title := ""
	color := ""
	icon := ""
	switch typ {
	case notifs.Error:
		title = " Error"
		color = "red"
		icon = "mdi mdi-alert"
	case notifs.Info:
		title = " Info"
		color = "blue"
		icon = "mdi mdi-information"
	case notifs.Success:
		title = " Success"
		color = "green"
		icon = "mdi mdi-check"
	case notifs.Closer:
		title = " Closed"
		color = "black"
		icon = "mdi mdi-message-bulleted-off"
	}
	return Div().ID(id.String()).Class("w3-container w3-half w3-border w3-"+color).Body(
		P().Body(I().Class(icon), Text(title)),
		Hr().Class("hr-right-white"),
		P().Text(msg),
		If(trace != "" && typ == notifs.Error,
			Pre().Body(Text(trace)),
		),
		If(data != nil && typ == notifs.Info,
			P().Text(fmt.Sprintf("%+v\n", data)),
		),
	)
}

func (n *NotificationComponent) OnMount(Context) {
	n.center = notifs.GetNotificationCenter()
	n.center.OnNotif(n.updateNotif)
}

func (n *NotificationComponent) updateNotif(notif notifs.Notif) {
	n.m.Lock()
	defer n.m.Unlock()
	n.notifs = append(n.notifs, notif)
	n.Update()
}
