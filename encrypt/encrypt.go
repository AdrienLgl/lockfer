package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func EncryptFile(file []byte, key []byte, filename string, id string) (bool, error) {
	// The key should be 16 bytes (AES-128), 24 bytes (AES-192) or
	// 32 bytes (AES-256)
	fmt.Printf("Encryption...\n")
	block, err := aes.NewCipher(key)
	if err != nil {
		return false, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return false, err
	}
	// Never use more than 2^32 random nonces with a given key
	// because of the risk of repeat.
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return false, err
	}
	// TODO create directory
	errM := os.MkdirAll("files/encrypted/"+id, os.ModePerm)
	if errM != nil {
		fmt.Println("Error during directory creation")
		fmt.Println(errM.Error())
		return false, errM
	}
	filePath := "files/encrypted/" + id + "/" + filename + ".bin"
	ciphertext := gcm.Seal(nonce, nonce, file, nil)
	// Save back to file
	err = ioutil.WriteFile(filePath, ciphertext, 0777)
	if err != nil {
		fmt.Println("Error during file writing: ")
		fmt.Printf("%s\n", filePath)
		return false, err
	}
	fmt.Printf("File %s encrypted successfully \n", filename)
	return true, nil
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
