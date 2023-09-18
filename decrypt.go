package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"

	"golang.org/x/crypto/pbkdf2"
)

func aesDecrypt(cipherAndSalts string, password string, keySize AESKeySize, iterations int) (string, error) {
	byteKey := []byte(password)

	saltKey, saltEnc, ciphertext := decodeCipherAndSalts(cipherAndSalts)

	rawKeySize := int(keySize)
	key := pbkdf2.Key(byteKey, saltKey, iterations, rawKeySize, sha256.New)

	// Init Key
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Init AES GCM (AEAD)
	aesGcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	message, err := aesGcm.Open(nil, saltEnc, ciphertext, nil)
	if err != nil {
		return "", err
	}
	
	return string(message), nil
}