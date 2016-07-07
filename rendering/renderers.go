/*
the rendering package contains several renderers for use with birdnest.
*/
package rendering

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"text/template"

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
	Render(*requests.Request) error
}

//DefaultRenderer is the default struct which renders a request and allows manipulation of rendering
//behavior via mutation of the renderFunc struct member. NOTE: because this can be changed at runtime I'm sure there
//are a few cool tricks which can be implemented, but I'd recommend against them. Changing renderer behavior on the fly can
//make for annoying debugging and also black holes that emit nad-punching unicorns will appear for, like... no reason!
//
//Unlike many traditional renderers which use a view configuration or REQUIRE manual rendering, birdnest renderers instead
//implement a NewView method which wraps a request handler and calls the renderers exported Render method
//after data has been handled. You may still call render manually if you choose. Multiple calls to Render should
//do nothing.
type DefaultRenderer struct {
	renderFunc func(*requests.Request) error
}

//NewStaticView accepts a filepath (for consistency, absolute path is recommended) to a static file and returns a renderer that can be used
//with the internal routing system.
func NewStaticView(file string) *DefaultRenderer {
	ret := new(DefaultRenderer)
	ret.renderFunc = func(r *requests.Request) error {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}
		r.Response.Write(data)
		return nil
	}
	return ret
}

func NewTemplateView(t *template.Template) *DefaultRenderer {
	ret := new(DefaultRenderer)
	ret.renderFunc = func(r *requests.Request) error {
		err := t.Execute(r.Response, r.Ctx)
		if err != nil {
			return err
		}
		r.Rendered = true
		return nil
	}
	return ret
}

//Render on the DefaultRenderer type satisfies the Renderer interface and allows direct rendering of
//a web response
func (d *DefaultRenderer) Render(r *requests.Request) error {
	return d.renderFunc(r)
}

//JSONRenderer transforms a request context into a JSON Response. If the renderer fails to process it will
//return an error.
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

		//set the Rendered flag as true
		req.Rendered = true

		return nil
	},
}
