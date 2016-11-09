package core

import (
	"fmt"
	"net/http"
	"context"

	"github.com/julienschmidt/httprouter"
)

type Route struct {
	Path   string
	Method string
	Action func(w http.ResponseWriter, r *http.Request)
}

type Router struct {
	httprouter *httprouter.Router
	container  *Container
}

func NewRouting() *Router {
	router := httprouter.New()
	// todo: move app params
	router.ServeFiles("/static/*filepath", http.Dir("/home/oroinc/work/golang/src/github.com/apoberez/pc/static"))
	return &Router{
		httprouter: router,
		container: NewContainer(),
	}
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.httprouter.ServeHTTP(w, r)
}

func (router *Router) Register(routes []*Route) {
	for _, route := range routes {
		switch route.Method {
		case "GET":
			router.httprouter.GET(route.Path, adapt(AuthMiddleware(route.Action), route, router.container))
		default:
			panic("method not supported")
		}
	}
}

func adapt(h http.HandlerFunc, route *Route, container *Container) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ctx := context.WithValue(r.Context(), "params", &p)
		ctx = context.WithValue(ctx, "route", route)
		ctx = context.WithValue(ctx, "container", container)

		// JSON API required content type
		w.Header().Set("Content-Type", "application/vnd.api+json")
		h(w, r.WithContext(ctx))
	}
}

// AuthMiddleware - takes in a http.HandlerFunc, and returns a http.HandlerFunc
var AuthMiddleware = func(f http.HandlerFunc) http.HandlerFunc {
	// one time scope setup area for middleware
	return func(w http.ResponseWriter, r *http.Request) {
		// ... pre handler functionality
		fmt.Println("start auth")
		f(w, r)
		fmt.Println("end auth")
		// ... post handler functionality
	}
}

// PrivateMiddleware - takes in a http.HandlerFunc, and returns a http.HandlerFunc
var PrivateMiddleware = func(f http.HandlerFunc) http.HandlerFunc {
	// one time scope setup area for middleware
	return func(w http.ResponseWriter, r *http.Request) {
		// ... pre handler functionality
		fmt.Println("start private")
		f(w, r)
		fmt.Println("end private")
		// ... post handler functionality
	}
}
