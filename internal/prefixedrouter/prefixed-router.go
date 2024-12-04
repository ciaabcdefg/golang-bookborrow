package prefixedrouter

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Middleware func(httprouter.Handle) httprouter.Handle

type PrefixedRouter struct {
	Prefix     string
	Router     *httprouter.Router
	Middleware Middleware
}

var EmptyMiddleware Middleware = func(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		next(w, r, p)
	}
}

func New(prefix string, router *httprouter.Router, middleware Middleware) *PrefixedRouter {
	return &PrefixedRouter{
		Prefix:     prefix,
		Router:     router,
		Middleware: middleware,
	}
}

func (r PrefixedRouter) GET(path string, handle httprouter.Handle) {
	log.Println("Mapped GET", r.Prefix+path)
	r.Router.GET(r.Prefix+path, r.Middleware(handle))
}

func (r PrefixedRouter) POST(path string, handle httprouter.Handle) {
	log.Println("Mapped POST", r.Prefix+path)
	r.Router.POST(r.Prefix+path, r.Middleware(handle))
}

func (r PrefixedRouter) DELETE(path string, handle httprouter.Handle) {
	log.Println("Mapped DELETE", r.Prefix+path)
	r.Router.DELETE(r.Prefix+path, r.Middleware(handle))
}

func (r PrefixedRouter) PATCH(path string, handle httprouter.Handle) {
	log.Println("Mapped PATCH", r.Prefix+path)
	r.Router.PATCH(r.Prefix+path, r.Middleware(handle))
}

func (r PrefixedRouter) PUT(path string, handle httprouter.Handle) {
	log.Println("Mapped PUT", r.Prefix+path)
	r.Router.PATCH(r.Prefix+path, r.Middleware(handle))
}

func (r PrefixedRouter) OPTIONS(path string, handle httprouter.Handle) {
	log.Println("Mapped OPTIONS", r.Prefix+path)
	r.Router.OPTIONS(r.Prefix+path, r.Middleware(handle))
}
