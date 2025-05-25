package main

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/git-starter-go/app/cmd/commands"
)

// Usage local_build.sh <command> <arg1> <arg2> ...
func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: mygit <command> [<args>...]\n")
		os.Exit(1)
	}

	switch command := os.Args[1]; command {
	case "init":
		commands.Init()
	case "cat-file":
		commands.CatFile(os.Args[2:])
	case "hash-object":
		commands.HashObject(os.Args[2:])
	default:
		fmt.Fprintf(os.Stderr, "Unknown command %s\n", command)
		os.Exit(1)
	}
}
