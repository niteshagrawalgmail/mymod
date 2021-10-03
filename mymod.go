package mymod

import "fmt"

// Print prints String "Hello, World!"

func Print(salutation, name string) {
	fmt.Printf("\n%s, %s!\n", salutation, name)
}
