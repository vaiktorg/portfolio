package notifs

import (
	"fmt"
	"sync"

	"github.com/vaiktorg/grimoire/uid"
)

type NotificationCenter struct {
	notif   chan Notif  //recieves message and sends them to be notified.
	handler func(Notif) //calls the functionality that handles the notification
}

var (
	once     sync.Once
	center   *NotificationCenter
	idlenght = 10
)

type Notif struct {
	ID       uid.UID
	Typ      Type
	Msg      string
	Error    string
	Data     interface{}
	isHidden bool
}

type Type uint8

const (
	Info Type = iota
	Success
	Error
	Closer //Stops NotifCenter, Init starts it.
)

func (n *NotificationCenter) Init() {
	n.notif = make(chan Notif)
	n.handler = defaultHandler
	go func() {
		for notif := range n.notif {
			switch notif.Typ {
			case Closer:
				n.handler(notif)
				return
			default:
				n.handler(notif)
			}

		}
	}()
}

// NotifyInfo
func (n *NotificationCenter) NotifInfo(msg string, data interface{}) {
	n.notif <- Notif{
		ID:   uid.NewUID(idlenght),
		Typ:  Info,
		Msg:  msg,
		Data: data,
	}
}

// NotifSuccess
func (n *NotificationCenter) NotifSuccess(msg string) {
	n.notif <- Notif{
		ID:  uid.NewUID(idlenght),
		Typ: Success,
		Msg: msg,
	}
}

// NotifError
func (n *NotificationCenter) NotifError(msg string, err error) {
	notif := Notif{
		ID:  uid.NewUID(idlenght),
		Typ: Error,
		Msg: msg,
	}
	if err != nil {
		notif.Error = err.Error()
	}
	n.notif <- notif
}

func (n *NotificationCenter) Close() {
	n.notif <- Notif{
		Typ: Closer,
	}
}
func (n *NotificationCenter) OnNotif(notifHandler func(notif Notif)) {
	n.handler = notifHandler
}

// =======================
func defaultHandler(notif Notif) {
	switch notif.Typ {
	case Error:
		fmt.Println(notif.Error, notif.Msg)
	case Info:
		fmt.Println(notif.Msg, notif.Data)
	case Closer:
		fmt.Println("Closed")
	}
}

// Globals
//======================
func GetNotificationCenter() *NotificationCenter {
	once.Do(func() {
		center = new(NotificationCenter)
		center.Init()
	})
	return center
}

func NotifError(msg string, err error) {
	center.NotifError(msg, err)
}
func NotifSuccess(msg string) {
	center.NotifSuccess(msg)
}
func NotifInfo(msg string, data interface{}) {
	center.NotifInfo(msg, data)
}
