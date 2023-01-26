package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		usage()
		os.Exit(1)
	}
	switch os.Args[1] {
	case "new":
		newPasswordCLI()
	case "check":
		checkPasswordCLI()
	default:
		usage()
	}
}

func usage() {
	fmt.Println("Possible Arguments:")
	fmt.Println("    new - hash a password")
	fmt.Println("    check - verify a password")
}
