package server

import (
	"net/http"

	"github.com/corvuscrypto/birdnest/requests"
	"github.com/julienschmidt/httprouter"
)

//Router is an adapter for the httprouter.Router
type Router struct {
	router *httprouter.Router
}

func transformRequest(w http.ResponseWriter, r *http.Request, p httprouter.Params) *requests.Request {
	req := new(requests.Request)
	req.Response = w
	req.Params = p
	return req
}

func wrapHandler(h RequestHandler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		req := transformRequest(w, r, p)
		applyMiddleware(req)
		h(req)
	}
}

//Handle is the adapter for birdnest routing to utilize httprouter's Handle method
func (r *Router) Handle(method, path string, handle RequestHandler) {
	r.router.Handle(method, path, wrapHandler(handle))
}

//OPTIONS is the adapter for the httprouter OPTIONS shortcut
func (r *Router) OPTIONS(path string, handle RequestHandler) {
	r.router.Handle("OPTIONS", path, wrapHandler(handle))
}

//PATCH is the adapter for the httprouter PATCH shortcut
func (r *Router) PATCH(path string, handle RequestHandler) {
	r.router.Handle("PATCH", path, wrapHandler(handle))
}

//POST is the adapter for the httprouter POST shortcut
func (r *Router) POST(path string, handle RequestHandler) {
	r.router.Handle("POST", path, wrapHandler(handle))
}

//PUT is the adapter for the httprouter PUT shortcut
func (r *Router) PUT(path string, handle RequestHandler) {
	r.router.Handle("PUT", path, wrapHandler(handle))
}

//GET is the adapter for the httprouter GET shortcut
func (r *Router) GET(path string, handle RequestHandler) {
	r.router.Handle("GET", path, wrapHandler(handle))
}

//DELETE is the adapter for the httprouter DELETE shortcut
func (r *Router) DELETE(path string, handle RequestHandler) {
	r.router.Handle("DELETE", path, wrapHandler(handle))
}

//HEAD is the adapter for the httprouter HEAD shortcut
func (r *Router) HEAD(path string, handle RequestHandler) {
	r.router.Handle("HEAD", path, wrapHandler(handle))
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(w, req)
}

//NewRouter returns a Router instance. If an *httprouter.Router instance is passed into NewRouter, the Router uses
//it, otherwise if it is nil the default is used
func NewRouter(r *httprouter.Router) *Router {
	if r == nil {
		r = httprouter.New()
	}
	ret := new(Router)
	ret.router = r
	return ret
}
