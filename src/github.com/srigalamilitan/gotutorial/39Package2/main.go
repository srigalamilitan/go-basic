package main

import (
	"crypto/sha1"
	"fmt"

	p "github.com/wuriyanto48/go-pbkdf2"
)

func main() {
	fmt.Println("Package External")
	pass := p.NewPassword(sha1.New, 8, 32, 15000)
	hashed := pass.HashPassword("123456")
	fmt.Println(hashed.CipherText)
	fmt.Println(hashed.Salt)

	isValid := pass.VerifyPassword("123456", hashed.CipherText, hashed.Salt)

	fmt.Println(isValid)
}
