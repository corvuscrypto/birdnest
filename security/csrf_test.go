package security

import "testing"

func TestCSRF(T *testing.T) {

	//generate a CSRF token
	token := GenerateCSRFToken()

	//check to ensure that the encoded length is the expected 44 characters long
	if len(token) != 44 {
		T.Errorf("Unexpected CSRF token length!")
	}
}
