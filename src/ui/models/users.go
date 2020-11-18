package models

import "time"

type User struct {
	name         string
	sessionStart time.Time
	connected    bool
}
