package prefixedrouter

import (
	"log"

	"github.com/julienschmidt/httprouter"
)

type PrefixedRouter struct {
	Prefix string
	Router *httprouter.Router
}

func (r PrefixedRouter) GET(path string, handle httprouter.Handle) {
	log.Println("Mapped GET", r.Prefix+path)
	r.Router.GET(r.Prefix+path, handle)
}

func (r PrefixedRouter) POST(path string, handle httprouter.Handle) {
	log.Println("Mapped POST", r.Prefix+path)
	r.Router.POST(r.Prefix+path, handle)
}

func (r PrefixedRouter) DELETE(path string, handle httprouter.Handle) {
	log.Println("Mapped DELETE", r.Prefix+path)
	r.Router.DELETE(r.Prefix+path, handle)
}

func (r PrefixedRouter) PATCH(path string, handle httprouter.Handle) {
	log.Println("Mapped PATCH", r.Prefix+path)
	r.Router.PATCH(r.Prefix+path, handle)
}

func (r PrefixedRouter) PUT(path string, handle httprouter.Handle) {
	log.Println("Mapped PUT", r.Prefix+path)
	r.Router.PATCH(r.Prefix+path, handle)
}

func (r PrefixedRouter) OPTIONS(path string, handle httprouter.Handle) {
	log.Println("Mapped OPTIONS", r.Prefix+path)
	r.Router.OPTIONS(r.Prefix+path, handle)
}
