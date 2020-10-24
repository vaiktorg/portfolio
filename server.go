//+build windows

package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/maxence-charriere/go-app/pkg/app"
)

func main() {
	h := &app.Handler{
		Title:        "HeptaCode Dashboard",
		Author:       "Vaiktorg",
		Name:         "Heptacode Dashboard",
		ShortName:    "HptCdDshbrd",
		Description:  "Portofolio",
		LoadingLabel: "Reading Grimoire...",

		Styles: []string{
			"https://www.w3schools.com/w3css/4/w3.css",
			//"https://pro.fontawesome.com/releases/v5.10.0/css/all.css",
			//"https://maxcdn.bootstrapcdn.com/fo/nt-awesome/4.7.0/css/font-awesome.min.css",
			"http://cdn.materialdesignicons.com/5.4.55/css/materialdesignicons.min.css",
			"web/static/style.css",
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

	//repo := app.GitHubPages("vaiktorg")
	err := app.GenerateStaticWebsite("vaiktorg.github.io", h)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	port := ":8080"
	fmt.Printf("Listening on: %v\n", port)
	_ = http.ListenAndServe(port, h)
}
