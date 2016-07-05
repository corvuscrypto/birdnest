package rendering

import (
	"encoding/json"
	"errors"

	"github.com/corvuscrypto/birdnest/requests"
)

//These are the built-in renderers that are automatically supported
const (
	JSONRenderer string = "json"
)

//These are the rendering errors. The UnableToRender error can be used,
//but more specific errors are always more preferable to generic ones
var (
	ErrRendererNotFound = errors.New("The requested renderer was not found!")
	ErrUnableToRender   = errors.New("The renderer was unable to render the data!")
)

//Renderer handles request contexts and marshals the data via HTTP. The Render method MUST throw an error if it is
//unable to render the data and it SHOULD set the Response code to 500
type Renderer interface {
	Render(*requests.Request) error
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
type jsonRenderer struct{}

//Render renders the context within a request to a proper JSON response
func (r *jsonRenderer) Render(req *requests.Request) error {
	data, err := json.Marshal(req.Ctx)
	if err != nil {
		req.Response.WriteHeader(500)
		return err
	}
	//ensure that the content header is set
	req.Response.Header().Set("Content-Type", "application/json")
	req.Response.Write(data)
	return nil
}

//RegisterRenderer adds a renderer to the registry for use
func RegisterRenderer(name string, renderer Renderer) {
	availableRenderers[name] = renderer
}

func init() {
	availableRenderers = make(map[string]Renderer)
	//add the default renderers to the availableRenderers map
	availableRenderers[JSONRenderer] = new(jsonRenderer)
}
