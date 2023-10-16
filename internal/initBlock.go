package internal

import (
	"crypto/aes"
	"crypto/cipher"
	"log"
)

var EncryptionKey = []byte("16byteAESKey1234")

func InitBlock(key []byte) cipher.Block {
	block, err := aes.NewCipher(EncryptionKey)
	if err != nil {
		log.Fatal(err)
	}
	return block
}
