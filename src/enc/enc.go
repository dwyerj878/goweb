package enc

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	b64 "encoding/base64"
	"io"
	"log"
)

func Encrypt(plainText string, keyText string) string {
	key := []byte(keyText)
	log.Println(len(key))
	text := []byte(plainText)
	c, err := aes.NewCipher(key)

	if err != nil {
		log.Println(err)
	}

	// gcm or Galois/Counter Mode, is a mode of operation
	// for symmetric key cryptographic block ciphers
	// - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	gcm, err := cipher.NewGCM(c)
	// if any error generating new GCM
	// handle them
	if err != nil {
		log.Println(err)
	}

	// creates a new byte array the size of the nonce
	// which must be passed to Seal
	nonce := make([]byte, gcm.NonceSize())
	// populates our nonce with a cryptographically secure
	// random sequence
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		log.Println(err)
	}

	// here we encrypt our text using the Seal function
	// Seal encrypts and authenticates plaintext, authenticates the
	// additional data and appends the result to dst, returning the updated
	// slice. The nonce must be NonceSize() bytes long and unique for all
	// time, for a given key.
	encrypted := gcm.Seal(nonce, nonce, text, nil)
	base64 := b64.StdEncoding.EncodeToString(encrypted)
	return string(base64)
}

func Decrypt(encrypted string, keyText string) string {
	key := []byte(keyText)
	ciphertext, err := b64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		log.Println(err)
	}
	c, err := aes.NewCipher(key)
	if err != nil {
		log.Println(err)
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		log.Println(err)
	}
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		log.Println(err)
	}
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		log.Println(err)
	}
	return string(plaintext)

}

func (r REQUEST) Encrypt() RESPONSE {
	encrypted := Encrypt(r.Text, r.Key)
	return RESPONSE{Result: encrypted}
}

func (r REQUEST) Decrypt() RESPONSE {
	decrypted := Decrypt(r.Text, r.Key)
	return RESPONSE{Result: decrypted}
}

type REQUEST struct {
	Text string `json:"text"`
	Key  string `json:"key"`
}

type RESPONSE struct {
	Result string `json:"result"`
}
