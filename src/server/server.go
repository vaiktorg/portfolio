//+build windows linux

package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/vaiktorg/grimoire/errs"

	"github.com/gin-gonic/gin"

	"github.com/maxence-charriere/go-app/pkg/app"
)

var (
	Title     string
	Desc      string
	Name      string
	ShortName string

	CSSPath string
	WEBPath string
)

func main() {
	fmt.Println(os.Getwd())
	mux := http.NewServeMux()
	uiHandler := &app.Handler{
		Title:        Title,
		Author:       "Vaiktorg",
		Name:         Name,
		ShortName:    ShortName,
		Description:  Desc,
		LoadingLabel: "Reading Grimoire...",
		Styles: []string{
			"https://www.w3schools.com/w3css/4/w3pro.css",
			"http://cdn.materialdesignicons.com/5.4.55/css/materialdesignicons.min.css",
			CSSPath,
		},
		Icon: struct {
			Default    string
			Large      string
			AppleTouch string
		}{
			Default:    "https://i.imgur.com/vNxAhoY.png",
			Large:      "https://i.imgur.com/vNxAhoY.png",
			AppleTouch: "",
		},
	}

	// Register the handler for ui
	mux.Handle("/", uiHandler)

	// Server functionality
	serv := gin.Default()
	mux.Handle("/s/", serv)

	serv.GET("/s/", func(c *gin.Context) {
		c.String(200, "WOWOW")
	})

	serv.GET("/s/:user/*action", func(c *gin.Context) {
		user := c.Param("user")
		action := c.Param("action")
		c.String(200, "%s\t%s", user, action)
	})

	// Run this shit!
	addr := ":8080"
	fmt.Println(Title + ": " + Desc)
	fmt.Printf("Listening on: %v\n", addr)

	err := http.ListenAndServe(addr, mux)
	errs.CheckError(err)
}
