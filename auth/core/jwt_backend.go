package Authentication

import (
	"bufio"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/Vijayaraagavan/backengine/models"
	"github.com/Vijayaraagavan/backengine/auth/settings"
)

type JWTAuthenticationBackend struct {
	privateKey *rsa.PrivateKey
	PublicKey *rsa.PublicKey
}

const (
	tokenDuration = 72
	expireOffset = 3600
)

var authBackendInstance *JWTAuthenticationBackend

func InitJWTAuthenticationBackend() *JWTAuthenticationBackend {
	if authBackendInstance == nil {
		authBackendInstance = &JWTAuthenticationBackend{
			privateKey: getPrivateKey(),
			PublicKey: getPublicKey(),
		}
	}
	return authBackendInstance
}

func (backend *JWTAuthenticationBackend) GenerateToken(person models.Person) (string, error) {
	var signingKey []byte
	signingKey, err := json.Marshal(backend.PublicKey)

	claims := jwt.MapClaims{
		"exp":          time.Now().Add(time.Hour * time.Duration(settings.Get().JWTExpirationDelta)).Unix(),
		"iat":          time.Now().Unix(),
		"sub":          person.Uuid,
		"userName": 	person.Name,
		"userPhoneNo": 	person.PhoneNo,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(signingKey)

	return ss, err
}

func getPrivateKey() *rsa.PrivateKey {
	privateKeyFile, err := os.Open(settings.Get().PrivateKeyPath)
	if err != nil {
		panic(err)
	}
	pemFileInfo, _ := privateKeyFile.Stat()
	var size int64 = pemFileInfo.Size()
	pembytes := make([]byte, size)

	buffer := bufio.NewReader(privateKeyFile)
	_, err = buffer.Read(pembytes)

	data, _ := pem.Decode([]byte(pembytes))

	privateKeyFile.Close()
	privateKeyImported, err := x509.ParsePKCS1PrivateKey(data.Bytes)

	if err != nil {
		panic(err)
	}

	return privateKeyImported
}

func getPublicKey() *rsa.PublicKey {
	publicKeyFile, err := os.Open(settings.Get().PublicKeyPath)
	if err != nil {
		panic(err)
	}

	pemfileinfo, _ := publicKeyFile.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)

	buffer := bufio.NewReader(publicKeyFile)
	_, err = buffer.Read(pembytes)

	data, _ := pem.Decode([]byte(pembytes))

	publicKeyFile.Close()

	publicKeyImported, err := x509.ParsePKIXPublicKey(data.Bytes)

	if err != nil {
		panic(err)
	}

	
	rsaPub, ok := publicKeyImported.(*rsa.PublicKey)

	if !ok {
		panic(err)
	}

	return rsaPub
}

// func GetAuthMethod() jwt.Keyfunc {
// 	backend := InitJWTAuthenticationBackend() 
// 	signingKey, _ := json.Marshal(backend.PublicKey)
// 	return func (token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(jwt.SigningMethodHS256); !ok {
// 			fmt.Println("Invalid token")
// 			return nil, fmt.Error("unexpected signing method")
// 		} else {
// 			return signingKey, nil
// 		}
// 	}
// }