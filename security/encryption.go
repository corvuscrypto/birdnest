package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/binary"
	"sync/atomic"

	"github.com/corvuscrypto/birdnest/config"
)

//Serializer is an interface that allows one to encrypt/decrypt data directly from/to a struct
type Serializer interface {
	Serialize() []byte
	Deserialize([]byte)
}

var encrypter cipher.AEAD
var nonceSeed []byte
var nonceCounter uint32

func init() {
	if config.Config.Get("SecretKey") == nil {
		secretKey := make([]byte, 32)
		rand.Read(secretKey)
		config.Config.Set("SecretKey", string(secretKey))
	}

	//setup the default cipher
	block, err := aes.NewCipher([]byte(config.Config.GetString("SecretKey")))
	if err != nil {
		panic(err)
	}

	setEncryptionBlock(block)

	//generate a random nonce seed
	nonceSeed = make([]byte, 8)
	rand.Read(nonceSeed)
}

func generateNonce() []byte {
	//Generate a nonce per guidelines specified in RFC 5116 §3.2
	//nonceSeed is the fixed field of the nonce
	//the counter field is an atomically incremented value initialized to a random value
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, atomic.AddUint32(&nonceCounter, 1))
	return append(nonceSeed, buf...)
}

//EncryptData takes a value and encrypts it using a GCM cipher block
func EncryptData(value interface{}) []byte {
	nonce := generateNonce()
	s, isSerializer := value.(Serializer)
	if !isSerializer {
		data, isByteSlice := value.([]byte)
		if !isByteSlice {
			panic("EncryptData expects either a Serializer or byte slice!")
		}
		return append(encrypter.Seal(data[:0], nonce, data, nil), nonce...)
	}
	data := s.Serialize()
	return append(encrypter.Seal(data[:0], nonce, data, nil), nonce...)
}

//DecryptData takes a destination interface{} and a byte slice and decrypts the bytes into the destination
//using a GCM cipher block. If dst is not nil, the decrypted data will be passed into the Serializer's
//Deserialize method. This method returns the decrypted bytes regardless.
func DecryptData(src []byte, dst Serializer) []byte {
	nonce := src[len(src)-12:]
	data := src[:len(src)-12]
	data, _ = encrypter.Open(data[:0], nonce, data, nil)
	if dst != nil {
		dst.Deserialize(data)
	}
	return data
}

func setEncryptionBlock(cphr cipher.Block) {
	var err error
	encrypter, err = cipher.NewGCM(cphr)
	if err != nil {
		panic(err)
	}
}
