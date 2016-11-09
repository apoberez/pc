package main

import (
	"net/http"
	"github.com/apoberez/pc/catalog"
	"github.com/apoberez/pc/core"
)

func main() {
	http.ListenAndServe(":9000", initRouting())
}

func initRouting() *core.Router {
	router := core.NewRouting()
	router.Register(catalog.GetRoutes())

	return router
}

