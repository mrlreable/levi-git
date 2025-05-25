package commands

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/codecrafters-io/git-starter-go/app/internal/objects"
)

func Init() {
	for _, dir := range []string{".git", ".git/objects", ".git/refs"} {
		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating directory: %s\n", err)
		}
	}

	headFileContents := []byte("ref: refs/heads/main\n")
	if err := os.WriteFile(".git/HEAD", headFileContents, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing file: %s\n", err)
	}

	fmt.Println("Initialized git directory")
}

func CatFile(args []string) {
	fs := flag.NewFlagSet("cat-file", flag.ExitOnError)
	var p string
	fs.StringVar(&p, "p", "", "SHA-1 hash of the object to be retrieved")
	fs.Parse(args)

	p = strings.TrimSpace(p)
	if p == "" {
		fmt.Fprintf(os.Stderr, "Empty hash value\n")
		os.Exit(1)
	}

	blob, err := objects.Read(p)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading blob")
		os.Exit(1)
	}

	fmt.Print(blob.Content)
}
