package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
)

func EncryptAES(key []byte, plaintext string) string {
	// create cipher
	c, err := aes.NewCipher(key)
	CheckError(err)
	// al	locate space for ciphered data
	out := make([]byte, len(plaintext))
	// encrypt
	c.Encrypt(out, []byte(plaintext))
	// return hex string
	return hex.EncodeToString(out)
}

func EncryptAESImage(plainData, secret []byte) (cipherData []byte) {
	block, _ := aes.NewCipher(secret)
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return
	}
	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return
	}
	cipherData = gcm.Seal(
		nonce,
		nonce,
		plainData,
		nil)
	return
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
