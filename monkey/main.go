package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/kazu69/golang-interpreter/monkey/repl"
)

func main() {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Println("Hello %s ! This is Monkey program language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
