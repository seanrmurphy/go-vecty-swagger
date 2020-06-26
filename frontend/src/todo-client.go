package main

import (
	"log"

	"github.com/gopherjs/vecty"
	"github.com/seanrmurphy/go-vecty-swagger/frontend/src/components"
	"github.com/seanrmurphy/go-vecty-swagger/frontend/src/store"
)

var (
	RestEndpoint = "https://3t9ljow8x2.execute-api.eu-west-1.amazonaws.com/prod/"
)

func main() {

	vecty.SetTitle("GopherJS • TodoMVC")
	vecty.AddStylesheet("https://rawgit.com/tastejs/todomvc-common/master/base.css")
	vecty.AddStylesheet("https://rawgit.com/tastejs/todomvc-app-css/master/index.css")

	store.Initialize(RestEndpoint)
	log.Printf("store.items length = %v\n", len(store.Items))

	p := &components.PageView{}
	store.Listeners.Add(p, func() {
		p.Items = store.Items
		log.Printf("in func p.items length = %v\n", len(p.Items))
		vecty.Rerender(p)
	})
	log.Printf("p.items length = %v\n", len(p.Items))
	vecty.RenderBody(p)
}
