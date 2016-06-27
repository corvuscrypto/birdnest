package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/binary"
	"errors"
	"os"
	"sync/atomic"

	"github.com/corvuscrypto/birdnest/config"
	"github.com/corvuscrypto/birdnest/logging"
)

//These are the errors to be expected when things go awry during the encryption/cipher creation process
var (
	ErrInvalidType = errors.New("Input must be either a Serializer or byte slice!")
	ErrInvalidData = errors.New("Input is invalid!")
)

//Serializer is an interface that allows one to encrypt/decrypt data directly from/to a struct
type Serializer interface {
	Serialize() []byte
	Deserialize([]byte)
}

//use only this value of the key once initialized to prevent haphazard tricks re: the key
var secretKey []byte

var encrypter cipher.AEAD
var nonceSeed []byte
var nonceCounter uint32

func init() {
	//generate a random nonce seed
	nonceSeed = make([]byte, 8)
	rand.Read(nonceSeed)

	if secretKey = []byte(config.Config.GetString("SecretKey")); len(secretKey) == 0 {
		secretKey = make([]byte, 32)
		rand.Read(secretKey)
		config.Config.Set("SecretKey", string(secretKey))
	}

	//setup the default cipher
	err := setDefaultEncrypter(secretKey)
	if err != nil {
		logging.Error("Encrypter couldn't initialize", err)
		os.Exit(1)
	}
}

func generateNonce() []byte {
	//Generate a nonce per guidelines specified in RFC 5116 §3.2
	//nonceSeed is the fixed field of the nonce
	//the counter field is an atomically incremented value initialized to a random value
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, atomic.AddUint32(&nonceCounter, 1))
	return append(nonceSeed, buf...)
}

func setDefaultEncrypter(secretKey []byte) error {
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return err
	}
	err = setEncryptionBlock(block)
	return err
}

func setEncryptionBlock(cphr cipher.Block) error {
	newEnc, err := cipher.NewGCM(cphr)
	//Don't replace the encrypter if we fail to set GC mode
	if err == nil {
		encrypter = newEnc
	}
	return err
}

//SetEncryptionBlock allows the use of custom cipher blocks. All cipher blocks are wrapped in GCM
func SetEncryptionBlock(f func([]byte) cipher.Block) error {
	return setEncryptionBlock(f(secretKey))
}

//EncryptData takes a value and encrypts it using a GCM cipher block
func EncryptData(value interface{}) ([]byte, error) {
	nonce := generateNonce()
	s, isSerializer := value.(Serializer)
	if !isSerializer {
		data, isByteSlice := value.([]byte)
		if !isByteSlice {
			return nil, ErrInvalidType
		}
		return append(encrypter.Seal(data[:0], nonce, data, nil), nonce...), nil
	}
	data := s.Serialize()
	return append(encrypter.Seal(data[:0], nonce, data, nil), nonce...), nil
}

//DecryptData takes a destination interface{} and a byte slice and decrypts the bytes into the destination
//using a GCM cipher block. If dst is not nil, the decrypted data will be passed into the Serializer's
//Deserialize method. This method returns the decrypted bytes regardless.
func DecryptData(src []byte, dst Serializer) ([]byte, error) {
	if len(src) < 12 {
		return nil, ErrInvalidData
	}
	nonce := src[len(src)-12:]
	data := src[:len(src)-12]
	data, err := encrypter.Open(data[:0], nonce, data, nil)
	if err != nil {
		return nil, err
	}
	if dst != nil {
		dst.Deserialize(data)
	}
	return data, nil
}
