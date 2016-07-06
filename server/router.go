package server

import (
	"net/http"

	"github.com/corvuscrypto/birdnest/rendering"
	"github.com/corvuscrypto/birdnest/requests"
	"github.com/julienschmidt/httprouter"
)

//RequestHandler type is the function signature required in
//order to register a new middleware request handler
type RequestHandler func(*requests.Request)

//Router is an adapter for the httprouter.Router. DO NOT INSTANTIATE THIS MANUALLY!
//This struct is exported only for documentation purposes and you must use the factory method provided (NewRouter)
type Router struct {
	router       *httprouter.Router
	PanicHandler RequestHandler
	NotFHandler  RequestHandler
}

func transformRequest(w http.ResponseWriter, r *http.Request, p httprouter.Params) *requests.Request {
	req := new(requests.Request)
	req.Request = r
	req.Response = w
	req.Params = p
	req.Ctx = make(requests.Context)
	return req
}

func (r *Router) panicHandler() {
	if p := recover(); p != nil {

	}
}

func (r *Router) wrapHandler(h RequestHandler, renderer rendering.Renderer) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		req := transformRequest(w, r, p)
		applyMiddleware(req)
		h(req)
		if renderer != nil {
			renderer.Render(req)
		}
	}
}

//Handle is the adapter for birdnest routing to utilize httprouter's Handle method
func (r *Router) Handle(method, path string, handle RequestHandler, renderer rendering.Renderer) {
	r.router.Handle(method, path, r.wrapHandler(handle, renderer))
}

//OPTIONS is the adapter for the httprouter OPTIONS shortcut.
//If renderer is not nil then the request will be rendered using that rendering agent
func (r *Router) OPTIONS(path string, handle RequestHandler, renderer rendering.Renderer) {
	r.router.Handle("OPTIONS", path, r.wrapHandler(handle, renderer))
}

//PATCH is the adapter for the httprouter PATCH shortcut.
//If renderer is not nil then the request will be rendered using that rendering agent
func (r *Router) PATCH(path string, handle RequestHandler, renderer rendering.Renderer) {
	r.router.Handle("PATCH", path, r.wrapHandler(handle, renderer))
}

//POST is the adapter for the httprouter POST shortcut.
//If renderer is not nil then the request will be rendered using that rendering agent
func (r *Router) POST(path string, handle RequestHandler, renderer rendering.Renderer) {
	r.router.Handle("POST", path, r.wrapHandler(handle, renderer))
}

//PUT is the adapter for the httprouter PUT shortcut.
//If renderer is not nil then the request will be rendered using that rendering agent
func (r *Router) PUT(path string, handle RequestHandler, renderer rendering.Renderer) {
	r.router.Handle("PUT", path, r.wrapHandler(handle, renderer))
}

//GET is the adapter for the httprouter GET shortcut.
//If renderer is not nil then the request will be rendered using that rendering agent
func (r *Router) GET(path string, handle RequestHandler, renderer rendering.Renderer) {
	r.router.Handle("GET", path, r.wrapHandler(handle, renderer))
}

//DELETE is the adapter for the httprouter DELETE shortcut.
//If renderer is not nil then the request will be rendered using that rendering agent
func (r *Router) DELETE(path string, handle RequestHandler, renderer rendering.Renderer) {
	r.router.Handle("DELETE", path, r.wrapHandler(handle, renderer))
}

//HEAD is the adapter for the httprouter HEAD shortcut.
//If renderer is not nil then the request will be rendered using that rendering agent
func (r *Router) HEAD(path string, handle RequestHandler, renderer rendering.Renderer) {
	r.router.Handle("HEAD", path, r.wrapHandler(handle, renderer))
}

//ServeFiles is wrapper to set a static fileserver route onto the underlying httprouter.Router instance
func (r *Router) ServeFiles(path string, root http.FileSystem) {
	r.router.ServeFiles(path, root)
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
