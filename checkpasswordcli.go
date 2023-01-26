package main

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"

	"github.com/brandon-detty/passhash-argon2/hasher"
)

func checkPasswordCLI() {
	password, err := getPassword()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	pepper, err := getHex("Pepper", hasher.CondimentLength)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	salt, err := getHex("Salt", hasher.CondimentLength)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	correctHash, err := getHex("Comparison hash", hasher.ArgonKeyLength)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Hash matches:", hasher.CheckPassword(password, salt, pepper, correctHash))
}

func getPassword() ([]byte, error) {
	p, err := quietPrompt("Password: ")
	if err != nil {
		return []byte{}, err
	}
	return p, nil
}

func getHex(which string, byteLength int) ([]byte, error) {
	input, err := quietPrompt(which + " (" + strconv.Itoa(byteLength*2) + "-char hex): ")
	if err != nil {
		return []byte{}, err
	}
	if len(input) != byteLength*2 {
		return []byte{}, errors.New("wrong length")
	}
	input, err = hex.DecodeString(string(input))
	if err != nil {
		return []byte{}, err
	}
	return input, nil
}
