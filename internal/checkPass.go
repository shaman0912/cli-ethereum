package internal

import (
	"fmt"
	"io/ioutil"
)

func CheckPassword(input string) {
	savedHashInFile, err := ioutil.ReadFile(input + ".txt")
	if err != nil {
		fmt.Println("You're not registered, capitalize the word Create to register.")
		CreateAccount()
		return
	}

	// Init AES with key
	block := InitBlock(EncryptionKey)
	savedHash := []byte(DecryptAes(savedHashInFile, block))

	if CompareHashes([]byte(input), savedHash) {
		fmt.Println("You have successfully logged in Wallet")
	} else {
		fmt.Println("Incorrect password. Try again.")
	}
}
