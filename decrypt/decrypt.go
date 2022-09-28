package decrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

func DecryptAES(key []byte, ct string) {
	ciphertext, _ := hex.DecodeString(ct)

	c, err := aes.NewCipher(key)
	CheckError(err)

	pt := make([]byte, len(ciphertext))
	c.Decrypt(pt, ciphertext)

	s := string(pt[:])
	fmt.Println("DECRYPTED:", s)
}

func DecryptAESImage(cipherData, secret []byte) (plainData []byte) {
	block, err := aes.NewCipher(secret)
	if err != nil {
		return
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := cipherData[:nonceSize], cipherData[nonceSize:]
	plainData, err = gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return
	}
	return
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
