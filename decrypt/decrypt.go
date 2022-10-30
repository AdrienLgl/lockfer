package decrypt

import (
	"archive/zip"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func DecryptFile(file []byte, filename string, secret_key []byte) ([]byte, error) {
	// The key should be 16 bytes (AES-128), 24 bytes (AES-192) or
	// 32 bytes (AES-256)
	block, err := aes.NewCipher(secret_key)
	if err != nil {
		fmt.Println("Erreur lors du déchiffrage du fichier")
		fmt.Println(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println("Erreur lors du déchiffrage du fichier")
		fmt.Println(err.Error())
	}
	nonce := file[:gcm.NonceSize()]
	file = file[gcm.NonceSize():]
	plaintext, err := gcm.Open(nil, nonce, file, nil)
	if err != nil {
		fmt.Println("Erreur lors du déchiffrage du fichier")
		fmt.Println(err.Error())
		return nil, err
	}
	return plaintext, nil
}

func CreateZip(path string, id string) {
	archive, err := os.Create(path + ".zip")
	if err != nil {
		panic(err)
	}
	defer archive.Close()
	zipWriter := zip.NewWriter(archive)
	files, err := ioutil.ReadDir(path)
	for _, file := range files {
		f, err := os.Open(path + "/" + file.Name())
		if err != nil {
			panic(err)
		}
		defer f.Close()
		w, err := zipWriter.Create(file.Name())
		if err != nil {
			panic(err)
		}
		if _, err := io.Copy(w, f); err != nil {
			panic(err)
		}
	}
	zipWriter.Close()
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
