package encyption

***REMOVED***
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
***REMOVED***
	"io"
***REMOVED***

type Encyption struct {
	Key []byte
***REMOVED***

func (e *Encyption***REMOVED*** Encrypt(text string***REMOVED*** (string, error***REMOVED*** {
	plaintent := []byte(text***REMOVED***

	block, err := aes.NewCipher(e.Key***REMOVED***
***REMOVED***
		return "", err
***REMOVED***

	cipherText := make([]byte, aes.BlockSize+len(plaintent***REMOVED******REMOVED***
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv***REMOVED***; err != nil {
		return "", err
***REMOVED***

	stream := cipher.NewCFBEncrypter(block, iv***REMOVED***
	stream.XORKeyStream(cipherText[aes.BlockSize:], plaintent***REMOVED***

	return base64.URLEncoding.EncodeToString(cipherText***REMOVED***, nil

***REMOVED***

func (e *Encyption***REMOVED*** Decrypt(crytoText string***REMOVED*** (string, error***REMOVED*** {
	cipherText, _ := base64.URLEncoding.DecodeString(crytoText***REMOVED***

	block, err := aes.NewCipher(e.Key***REMOVED***
***REMOVED***
		return "", err
***REMOVED***

	if len(cipherText***REMOVED*** < aes.BlockSize {
		return "", nil
***REMOVED***

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv***REMOVED***
	stream.XORKeyStream(cipherText, cipherText***REMOVED***

	return fmt.Sprintf("%s", cipherText***REMOVED***, nil
***REMOVED***
