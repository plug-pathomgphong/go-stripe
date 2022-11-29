package encyption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

type Encyption struct {
	Key []byte
}

func (e *Encyption) Encrypt(text string) (string, error) {
	plaintent := []byte(text)

	block, err := aes.NewCipher(e.Key)
	if err != nil {
		return "", err
	}

	cipherText := make([]byte, aes.BlockSize+len(plaintent))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plaintent)

	return base64.URLEncoding.EncodeToString(cipherText), nil

}

func (e *Encyption) Decrypt(crytoText string) (string, error) {
	cipherText, _ := base64.URLEncoding.DecodeString(crytoText)

	block, err := aes.NewCipher(e.Key)
	if err != nil {
		return "", err
	}

	if len(cipherText) < aes.BlockSize {
		return "", nil
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return fmt.Sprintf("%s", cipherText), nil
}
