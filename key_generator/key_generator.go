package key_generator

import (
	"crypto/rand"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	uuid "github.com/nu7hatch/gouuid"
)

type Identifier struct {
	Id  string
	Key string
	jwt.StandardClaims
}

type TokenDecryption struct {
	Id        string
	Key       string
	Encrypted string
}

func createKey(value int) []byte {
	key := make([]byte, value)
	_, err := rand.Read(key)
	if err != nil {
		panic("Erreur lors de la génération de la clé")
	}
	return key
}

func GetKey() []byte {
	key := createKey(32)
	return key
}

func CreateIdentifier() string {
	identifier := createKey(12)
	return string(identifier)
}

func CreateToken(secret string, id string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return err.Error()
	}
	tokenSecret := os.Getenv("TOKEN")
	date := time.Now()
	date = date.AddDate(0, 0, 1)
	claims := Identifier{
		id,
		secret,
		jwt.StandardClaims{
			ExpiresAt: date.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		fmt.Println("Erreur lors de la création du token")
	}
	fmt.Println(tokenString, err)
	return tokenString
}

func CreateUUID() (string, error) {
	u, e := uuid.NewV4()
	exist, _ := exists("files/encrypted/" + u.String())
	if exist {
		CreateUUID()
	}
	return u.String(), e
}

func CreateFilename(id string) (string, error) {
	u, e := uuid.NewV4()
	exist, _ := exists("files/decrypted/" + id + "/" + u.String())
	if exist {
		CreateUUID()
	}
	return u.String(), e
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func DecryptToken(tokenString string) (TokenDecryption, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return TokenDecryption{}, err
	}
	tokenSecret := os.Getenv("TOKEN")
	token, err := jwt.ParseWithClaims(tokenString, &Identifier{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenSecret), nil
	})

	if err != nil {
		fmt.Println("Erreur sur le token")
		return TokenDecryption{}, err
	}

	if claims, ok := token.Claims.(*Identifier); ok && token.Valid {
		return TokenDecryption{
			claims.Id,
			claims.Key,
			string(claims.StandardClaims.ExpiresAt),
		}, nil
	} else {
		fmt.Println("Erreur lors du décryptge du token")
		return TokenDecryption{}, err
	}
}
