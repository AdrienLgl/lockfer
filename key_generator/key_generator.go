package key_generator

import "crypto/rand"

func createKey(value int) []byte {
	key := make([]byte, value)
	_, err := rand.Read(key)
	if err != nil {
		panic("Erreur lors de la génération de la clé")
	}
	return key
}

func GetKey() []byte {
	key := createKey(64)
	return key
}

func CreateIdentifier() []byte {
	identifier := createKey(12)
	return identifier
}
