package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: bcrypter <password>")
		os.Exit(1)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(os.Args[1]), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Println("Hash:", string(hashedPassword))
}
