//Package birdnest is just a fun little framework that aims at making things work (for me). This was built with my own
//mind in mind (heheh) and if you are like-minded, maybe you'll find this framework useful. If not, there's tons of
//others. Seriously. I shouldn't have made this one, which reminds me of Robert Frost's poem about two roads. Think of
//the two roads like two frameworks. Well in this case the choice I made was to just plow right through the sodding
//middle of the fork like Nigel Farage's rhetoric toward a busfull of European liberals waving the EU flag around while
//driving in England on the right side of the road.
package birdnest

//These imports are to make getting all of the subpackages easier
import (
	_ "github.com/corvuscrypto/birdnest/config"
	_ "github.com/corvuscrypto/birdnest/logging"
	_ "github.com/corvuscrypto/birdnest/rendering"
	_ "github.com/corvuscrypto/birdnest/requests"
	_ "github.com/corvuscrypto/birdnest/security"
	_ "github.com/corvuscrypto/birdnest/server"
	_ "github.com/corvuscrypto/birdnest/sessions"
)
