package main

import (
	"fmt"
	"net/http"
	"plugin"
)

const SOPluginServerAddr = ":8000"

type Plugin interface {
	Name(p string) string
}

var loadedPlugin Plugin

func loadSOPlugin() {
	p, err := plugin.Open("plugin.so")
	if err != nil {
		panic(err)
	}
	pp, err := p.Lookup("Plugin")
	if err != nil {
		panic(err)
	}
	var ok bool
	loadedPlugin, ok = pp.(Plugin)
	if !ok {
		panic("assertion error")
	}
}

func ServeSOPlugin() {
	loadSOPlugin()
	http.HandleFunc("/so", func(w http.ResponseWriter, r *http.Request) {
		str := loadedPlugin.Name(Name)
		fmt.Fprintln(w, str)
	})
	http.ListenAndServe(SOPluginServerAddr, nil)
}
