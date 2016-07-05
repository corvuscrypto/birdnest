/*
the rendering package contains several renderers for use with birdnest.
*/
package rendering

import (
	"encoding/json"
	"errors"

	"github.com/corvuscrypto/birdnest/requests"
)

//These are the rendering errors. The UnableToRender error can be used,
//but more specific errors are always more preferable to generic ones
var (
	ErrRendererNotFound = errors.New("The requested renderer was not found!")
	ErrUnableToRender   = errors.New("The renderer was unable to render the data!")
)

//Renderer handles request contexts and marshals the data via HTTP. The Render method MUST throw an error if it is
//unable to render the data and it SHOULD set the Response code to 500.
//
//Unlike many traditional renderers which use a view configuration or REQUIRE manual rendering, birdnest renderers instead
//implement a NewView method which wraps a request handler and calls the renderers exported Render method
//after data has been handled. You may still call render manually if you choose. Multiple calls to Render should
//do nothing.
type Renderer interface {
	NewView(func(*requests.Request)) func(*requests.Request) error
	Render(*requests.Request) error
}

//DefaultRenderer is the default struct which satisfies the Renderer interface and allows manipulation of rendering
//behavior via mutation of the renderFunc struct member. NOTE: because this can be changed at runtime I'm sure there
//are a few cool tricks to be implemented, but I'd recommend against them. Changing renderer behavior on the fly can
//make for annoying debugging and also black holes that emit nad-punching unicorns will appear for like... no reason!
type DefaultRenderer struct {
	renderFunc func(*requests.Request) error
}

//NewView on the DefaultRenderer type satisfied the Renderer interface and wraps a request handler such that rendering
//occurs at the final stage of request handling. For the DefaultRenderer, if there is an error during rendering, the
//response will have a 500 code written to it.
func (d *DefaultRenderer) NewView(f func(*requests.Request)) func(*requests.Request) {
	return func(r *requests.Request) {
		f(r)
		err := d.renderFunc(r)
		if err != nil {
			r.Response.WriteHeader(500)
		}
	}
}

//Render on the DefaultRenderer type satisfies the Renderer interface and allows direct rendering of
//a web response
func (d *DefaultRenderer) Render(r *requests.Request) error {
	return d.renderFunc(r)
}

//this is for private access only
var availableRenderers map[string]Renderer

//GetRenderer retrieves a renderer that is in the registry. If the renderer cannot be found, an error will be returned
func GetRenderer(name string) (Renderer, error) {
	r := availableRenderers[name]
	if r == nil {
		return nil, ErrRendererNotFound
	}
	return r, nil
}

//JSONRenderer transforms a request context into a JSON Response. If the renderer fails to process it will
//automatically set the header to an error 500 code.
var JSONRenderer = &DefaultRenderer{
	func(req *requests.Request) error {
		data, err := json.Marshal(req.Ctx)
		if err != nil {
			req.Response.WriteHeader(500)
			return err
		}
		//ensure that the content header is set
		req.Response.Header().Set("Content-Type", "application/json")
		req.Response.Write(data)
		return nil
	},
}
