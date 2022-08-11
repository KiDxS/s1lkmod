package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func AesDecrypt0(text string) (plaintext string) {
	base64Decoded, _ := base64.URLEncoding.DecodeString(text)
	key := []byte("c5abedfe64975b437a9d92f76705cd43")
	encryptedText := base64Decoded

	aesCipher, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(aesCipher)
	if err != nil {
		panic(err)
	}

	nonceSize := gcm.NonceSize()
	if len(encryptedText) < nonceSize {
		panic(err)
	}

	nonce, encryptedText := encryptedText[:nonceSize], encryptedText[nonceSize:]
	decrypt, err := gcm.Open(nil, nonce, encryptedText, nil)
	if err != nil {
		panic(err)
	}
	plaintext = string(decrypt)
	return
}
func AesDecrypt1(text string) (plaintext string) {
	base64Decoded, _ := base64.URLEncoding.DecodeString(text)
	key := []byte("fa2fd1d8beadf10fb32ffff8a1be0c4a")
	encryptedText := base64Decoded

	aesCipher, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(aesCipher)
	if err != nil {
		panic(err)
	}

	nonceSize := gcm.NonceSize()
	if len(encryptedText) < nonceSize {
		panic(err)
	}

	nonce, encryptedText := encryptedText[:nonceSize], encryptedText[nonceSize:]
	decrypt, err := gcm.Open(nil, nonce, encryptedText, nil)
	if err != nil {
		panic(err)
	}
	plaintext = string(decrypt)
	return
}
