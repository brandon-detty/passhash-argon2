package main

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/brandon-detty/passhash-argon2/hasher"
)

func newPasswordCLI() {
	password, err := getNewPassword()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	pepper, err := getPepper()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	salt, hash, err := hasher.CreateHash(password, pepper)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	printWithLabel("Salt", hex.EncodeToString(salt))
	printWithLabel("Hash", hex.EncodeToString(hash))
}

func getNewPassword() ([]byte, error) {
	p, err := quietPrompt("Password: ")
	if err != nil {
		return []byte{}, err
	}
	p2, err := quietPrompt("Re-enter Password: ")
	if err != nil {
		return []byte{}, err
	}
	if !bytes.Equal(p, p2) {
		return []byte{}, errors.New("passwords do not match")
	}
	return p, nil
}

func getPepper() ([]byte, error) {
	pepper, err := quietPrompt("Pepper (32-char hex; blank to generate randomly): ")
	if err != nil {
		return []byte{}, err
	}
	switch len(pepper) {
	case 0:
		pepper, err = hasher.GenerateCondiment()
		if err != nil {
			return []byte{}, err
		}
		printWithLabel("Random Pepper", hex.EncodeToString(pepper))
	case 32:
	default:
		return []byte{}, errors.New("pepper must be 32 characters (hex)")
	}
	return pepper, nil
}
