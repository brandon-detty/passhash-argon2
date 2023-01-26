package main

import (
	"bytes"
	"fmt"
	"syscall"

	"golang.org/x/term"
)

const (
	maxLabelWidth = "14"
)

func quietPrompt(prompt string) ([]byte, error) {
	fmt.Print(prompt)
	input, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Print("\n")
	if err != nil {
		return []byte{}, err
	}
	return bytes.TrimSpace(input), nil
}

func printWithLabel(label string, value string) {
	fmt.Printf("%-"+maxLabelWidth+"s %v\n", label, value)
}
