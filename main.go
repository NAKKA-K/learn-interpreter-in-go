package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/NAKKA-K/learn-interpreter-in-go/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the interpreter-in-go programming ranguage!\n", user.Username)
	fmt.Println("Feel free to type in commands")
	repl.Start(os.Stdin, os.Stdout)
}
