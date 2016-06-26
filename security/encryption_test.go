package security

import (
	"bytes"
	"encoding/gob"
	"runtime"
	"sync"
	"testing"
)

func TestNonceGeneration(T *testing.T) {
	//test that for 10,000 nonces generated that there are no collisions
	//really this is more of an assurance test that the atomic counter works reliably on your machine

	var wg = new(sync.WaitGroup)
	var holder = make([][]byte, 10000)
	runtime.GOMAXPROCS(runtime.NumCPU())
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(i int) {
			holder[i] = generateNonce()
			//since standard nonce length is 12 bytes lets just ensure that this is an invariant during generation
			if len(holder[i]) != 12 {
				T.Errorf("Encountered unexpected nonce length")
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	for i := 0; i < 10000; i++ {
		for j := i + 1; j < 10000; j++ {
			if string(holder[i]) == string(holder[j]) {
				T.Errorf("Encountered duplicate nonce")
			}
		}
	}
}

type TestSerializer struct {
	Secret         string
	ID             uint32
	FavoriteColour string
}

func (t *TestSerializer) Serialize() []byte {
	var buf = new(bytes.Buffer)
	encoder := gob.NewEncoder(buf)
	encoder.Encode(t)
	return buf.Bytes()
}

func (t *TestSerializer) Deserialize(b []byte) {
	var buf = bytes.NewBuffer(b)
	decoder := gob.NewDecoder(buf)
	decoder.Decode(t)
}

func TestEncryption(T *testing.T) {
	//Make a Serializer to test
	exampleSerializer := new(TestSerializer)
	exampleSerializer.FavoriteColour = "Yellow"
	exampleSerializer.ID = 45
	exampleSerializer.Secret = "Super Secret"

	first := EncryptData(exampleSerializer)
	firstNon := exampleSerializer.Serialize()

	if string(first) == string(firstNon) {
		T.Errorf("Encrypted string matched unencrypted string")
	}

	second := EncryptData(exampleSerializer)

	//the first should NOT match the second
	if string(first) == string(second) {
		T.Errorf("Encountered unexpected string match!")
	}

	//now decrypt the data into a new struct
	decrypted := new(TestSerializer)
	DecryptData(first, decrypted)

	//we wont check if the values match. That is the responsibility of the user

	//Now encrypt and decrypt a byte slice
	mockData := []byte("This is so super secret")
	bfirst := EncryptData(mockData)
	bsecond := EncryptData(mockData)

	//Again we assert that these are not the same
	if string(bfirst) == string(bsecond) {
		T.Errorf("Encountered unexpected string match!")
	}

	bdecrypted := DecryptData(bfirst, nil)
	if string(bdecrypted) != string(mockData) {
		T.Errorf("Decrypted value did not match the source")
	}

}

type FakeBlock struct{}

func (f *FakeBlock) BlockSize() int {
	return 10
}
func (f *FakeBlock) Encrypt(d, s []byte) {}
func (f *FakeBlock) Decrypt(d, s []byte) {}

func TestPanics(T *testing.T) {

	//Test the panic of the default encrypter with a bad key
	func() {
		defer func() {
			if catch := recover(); catch == nil {
				T.Errorf("Did not panic as expected!")
			}
		}()
		setDefaultEncrypter([]byte("asd"))
	}()

	//Test the panic of setting an encrypter
	func() {
		defer func() {
			if catch := recover(); catch == nil {
				T.Errorf("Did not panic as expected!")
			}
		}()
		setEncryptionBlock(new(FakeBlock))
	}()

	//Test the panic of encrypting something other than a byte slice or encrypter
	func() {
		defer func() {
			if catch := recover(); catch == nil {
				T.Errorf("Did not panic as expected!")
			}
		}()
		EncryptData("asd")
	}()
}
