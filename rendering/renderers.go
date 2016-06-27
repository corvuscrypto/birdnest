package rendering

import "github.com/corvuscrypto/birdnest/requests"

type Renderer interface {
	render(*requests.Request)
}
