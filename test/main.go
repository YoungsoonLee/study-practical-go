package main

import (
	"fmt"
	"os"
	"syscall"

	"golang.org/x/term"
)

func main() {
	fmt.Print("Password: ")
	bytepw, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		os.Exit(1)
	}
	pass := string(bytepw)
	fmt.Printf("\nYou've entered: %q\n", pass)
}
