package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"

	"golang.org/x/crypto/pbkdf2"
)



func aesEncrypt(message string, password string, keySize AESKeySize, iterations int, saltSize int) (string, error) {
	byteKey := []byte(password)
	byteMessage := []byte(message)

	saltKey := getSalt(saltSize)

	rawKeySize := int(keySize) 
	key := pbkdf2.Key(byteKey, saltKey, iterations, rawKeySize, sha256.New)

	// Init Key
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	saltEnc := getSalt(saltSize)

	// Init AES GCM (AEAD)
	aesGcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	ciphertext := aesGcm.Seal(nil, saltEnc, byteMessage, nil)

	cipherSalts, err := encodeCipherAndSalts([][]byte{saltKey, saltEnc, ciphertext})
	if err != nil {
		return "", err
	}
	
	return cipherSalts, nil
}