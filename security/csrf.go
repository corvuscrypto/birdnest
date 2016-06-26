package security

import (
	"crypto/rand"
	"encoding/base64"
)

//GenerateCSRFToken generates and returns a cryptographically random token for use in cross-site request forgery protection
func GenerateCSRFToken() string {
	//256 bit entropy
	token := make([]byte, 32)
	rand.Read(token)
	return base64.StdEncoding.EncodeToString(token)
}
