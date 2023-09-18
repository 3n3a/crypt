package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"strings"
)

type AESKeySize int
const (
	Size128bit AESKeySize = 16
	Size192bit AESKeySize = 24
	Size256bit AESKeySize = 32
)

func getSalt(n int) []byte {
	nonce := make([]byte, n)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	return (nonce)
}

func decodeCipherAndSalts(cipherSalt string) ([]byte, []byte, []byte) {
	// __salt_key__$__salt_enc__$__ciphertext__
	fmt.Println(cipherSalt)
	parts := strings.Split(cipherSalt, "$")
	fmt.Println(parts)
	byteParts := make([][]byte, 0)
	for _, hexPart := range parts {
		decodedHexPart, _ := hex.DecodeString(hexPart)
		byteParts = append(byteParts, decodedHexPart)
	}
	return byteParts[0], byteParts[1], byteParts[2]
}

func encodeCipherAndSalts(parts [][]byte) (string, error) {
	// __salt_key__$__salt_enc__$__ciphertext__
	hexParts := make([]string, 0)
	for _, part := range parts {
		hexParts = append(hexParts, hex.EncodeToString(part))
	}
	cipherSalt := strings.Join(hexParts, "$")
	fmt.Println(cipherSalt[:50])
	return cipherSalt, nil
}