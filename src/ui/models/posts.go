package models

import (
	"time"

	"github.com/vaiktorg/grimoire/uid"
)

type PostCard struct {
	Title     string `yaml:"title"`
	Desc      string `yaml:"desc"`
	Timestamp time.Time
	Img       string `yaml:"img"`
}

type Post struct {
	id        uid.UID
	title     string
	desc      string
	timestamp time.Time
	content   string
}
