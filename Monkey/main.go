package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/Mr-xiaotian/Celestial-Language/Monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Print("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
