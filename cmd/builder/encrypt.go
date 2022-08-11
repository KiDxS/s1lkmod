package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"log"
)

func AesEncrypt0(text string) (encryptedText string) {
	key := []byte("c5abedfe64975b437a9d92f76705cd43")
	plainText := []byte(text)

	aesCipher, _ := aes.NewCipher(key)

	gcm, _ := cipher.NewGCM(aesCipher)

	// creates a new byte array the size of the nonce
	// which must be passed to Seal
	nonce := make([]byte, gcm.NonceSize())
	// populates our nonce with a cryptographically secure
	// random sequence
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatalln(err)
	}

	// here we encrypt our text using the Seal function
	// Seal encrypts and authenticates plaintext, authenticates the
	// additional data and appends the result to dst, returning the updated
	// slice. The nonce must be NonceSize() bytes long and unique for all
	// time, for a given key.
	encrypted := gcm.Seal(nonce, nonce, plainText, nil)
	encryptedText = base64.URLEncoding.EncodeToString(encrypted)
	return
}

func AesEncrypt1(text string) (encryptedText string) {
	key := []byte("fa2fd1d8beadf10fb32ffff8a1be0c4a")
	plainText := []byte(text)

	aesCipher, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	gcm, _ := cipher.NewGCM(aesCipher)
	// if any error generating new GCM
	// handle them

	// creates a new byte array the size of the nonce
	// which must be passed to Seal
	nonce := make([]byte, gcm.NonceSize())
	// populates our nonce with a cryptographically secure
	// random sequence
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		log.Println(err)
	}

	// here we encrypt our text using the Seal function
	// Seal encrypts and authenticates plaintext, authenticates the
	// additional data and appends the result to dst, returning the updated
	// slice. The nonce must be NonceSize() bytes long and unique for all
	// time, for a given key.
	encrypted := gcm.Seal(nonce, nonce, plainText, nil)
	encryptedText = base64.URLEncoding.EncodeToString(encrypted)
	return
}
