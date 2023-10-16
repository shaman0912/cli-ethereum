package internal

import (
	"bufio"
	variables "cli-ethereum"
	"fmt"
	"os"
	"strings"
)

func CreateAccount() {
	fmt.Println("Please enter your name and the name must not be longer than 16 characters:")
	input := bufio.NewReader(os.Stdin)
	name, err := input.ReadString('\n')
	if err != nil {
		fmt.Printf("Wrong your enter: %v\n", err)
		return
	}
	name = ValidateAndPad(strings.TrimSpace(name))
	block := InitBlock(variables.EncryptionKey)
	passAes := EncryptAes([]byte(name), block)
	HashWriteToFile(name, passAes)

}
