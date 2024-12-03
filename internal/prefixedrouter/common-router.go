package prefixedrouter

import "github.com/julienschmidt/httprouter"

type CommonRouter interface {
	GET(path string, handle httprouter.Handle)
	POST(path string, handle httprouter.Handle)
	DELETE(path string, handle httprouter.Handle)
	PATCH(path string, handle httprouter.Handle)
	PUT(path string, handle httprouter.Handle)
	OPTIONS(path string, handle httprouter.Handle)
}
