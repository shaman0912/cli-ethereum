package internal

import (
	"fmt"
	"os"
	"strings"
)

func HashWriteToFile(name string, pass []byte) error {
	filename := strings.TrimSpace(name) + ".txt"
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error when creating a file:", err)
		return err
	}
	defer file.Close()

	_, err = file.Write(pass)
	if err != nil {
		fmt.Println("Error when writing to a file:", err)
		return err
	}

	fmt.Printf("You have successfully registered\n")
	return nil
}
