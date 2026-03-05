package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/KarenTsaturyan/HeroLang/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! HeroLang programming language!\n", user.Username)
	fmt.Printf("Type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
	// Result of REPL
	// Let a = 5
	// {Type:IDENT Literal:Let}
	// {Type:IDENT Literal:a}
	// {Type:= Literal:=}
	// {Type:INT Literal:5}
}
