package main

import (
	"bufio"
	"cli-ethereum/internal"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Println("If you don't have an account, capitalize the word Create, if you are registered, enter your password")
		inputTemp := bufio.NewReader(os.Stdin)
		input, err := inputTemp.ReadString('\n')
		if err != nil {
			fmt.Printf("Wrong your enter: %v\n", err)
			return
		}
		input = strings.TrimSpace(input)

		internal.Authentication(input)

		fmt.Print("Do you want to perform another operation? (yes/no): ")
		anotherOp, _ := inputTemp.ReadString('\n')
		if strings.TrimSpace(anotherOp) != "yes" {
			break
		}
	}
}
