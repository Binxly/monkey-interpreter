package main

import (
	"fmt"
	"os"
	"os/user"
	"monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("How's it going, %s? Welcome to my Monkey programming language interpreter.\n",
		user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
