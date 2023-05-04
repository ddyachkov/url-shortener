// Package cookie reads and writes encrypted cookie values.
package cookie

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

// Errors
var (
	ErrValueTooLong = errors.New("cookie value too long")
	ErrInvalidValue = errors.New("invalid cookie value")
)

// GetEncryptedValue returns decrypted value from received cookie.
func GetEncryptedValue(r *http.Request, name string, secretKey []byte) (value string, err error) {
	cookie, err := r.Cookie(name)
	if err != nil {
		return "", err
	}

	encryptedValue, err := base64.URLEncoding.DecodeString(cookie.Value)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()
	if len(encryptedValue) < nonceSize {
		return "", ErrInvalidValue
	}

	nonce := encryptedValue[:nonceSize]
	cipherText := encryptedValue[nonceSize:]

	plainText, err := aesGCM.Open(nil, []byte(nonce), []byte(cipherText), nil)
	if err != nil {
		return "", err
	}

	expectedName, value, ok := strings.Cut(string(plainText), ":")
	if !ok {
		return "", ErrInvalidValue
	}

	if expectedName != name {
		return "", ErrInvalidValue
	}

	return value, nil
}

// GetEncryptedValue writes encrypted value to a certain cookie.
func WriteEncryptedValue(w http.ResponseWriter, name string, value string, secretKey []byte) (err error) {
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	cookie := http.Cookie{
		Name:  name,
		Value: value,
		Path:  "/",
	}

	plainText := fmt.Sprintf("%s:%s", cookie.Name, cookie.Value)

	encryptedValue := aesGCM.Seal(nonce, nonce, []byte(plainText), nil)

	cookie.Value = string(encryptedValue)
	cookie.Value = base64.URLEncoding.EncodeToString([]byte(cookie.Value))

	if len(cookie.String()) > 4096 {
		return ErrValueTooLong
	}

	http.SetCookie(w, &cookie)

	return nil
}
