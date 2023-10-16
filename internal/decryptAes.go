package internal

import (
	"crypto/aes"
	"crypto/cipher"
)

func DecryptAes(data []byte, block cipher.Block) string {
	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(data, data)
	return string(data)
}
