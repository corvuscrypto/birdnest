package sessions

import (
	"crypto/rand"
	"encoding/base64"
	"time"

	"github.com/corvuscrypto/birdnest/config"
	"github.com/corvuscrypto/birdnest/security"
)

//SessionValidator is the interface through which most sessions can be validated
type SessionValidator interface {
	IsGuest() bool
	IsValid() bool
}

//Session is the struct representation of a web session
type Session struct {
	Owner          interface{}
	Expiration     time.Time
	rawToken       []byte
	encryptedToken string
}

//Serialize satisfies the Serializer interface
func (s *Session) Serialize() []byte {
	return s.rawToken
}

//Deserialize satisfies the Serializer interface
func (s *Session) Deserialize(data []byte) {
	s.rawToken = data
}

//NewSession generates a new session
func NewSession(owner interface{}) *Session {
	sess := new(Session)
	sess.Owner = owner
	sessBytes := make([]byte, 32)
	rand.Read(sessBytes)
	base64.StdEncoding.Encode(sessBytes, sessBytes)
	sess.rawToken = sessBytes
	sess.encryptedToken = string(security.EncryptData(sess))
	//set the Expiration
	sess.Expiration = time.Unix(int64(config.Config.GetInt("SessionExpiration", 0)), 0)

	return sess
}
