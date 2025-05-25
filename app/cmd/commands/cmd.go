package commands

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/codecrafters-io/git-starter-go/app/internal/objects"
)

func Init() {
	for _, dir := range []string{objects.ROOT, objects.OBJECT_STORE, objects.REFS_STORE} {
		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating directory: %s\n", err)
		}
	}

	headFileContents := []byte("ref: refs/heads/main\n")
	if err := os.WriteFile(objects.HEAD, headFileContents, 0644); err != nil {
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

func HashObject(args []string) {
	fs := flag.NewFlagSet("hash-object", flag.ExitOnError)
	var w string
	fs.StringVar(&w, "w", "", "Text to be hashed and inserted into the object storage")
	fs.Parse(args)

	blob := objects.Blob{
		Content: w,
		Size:    uint(len(w)),
	}

	err := blob.Write()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing blob")
		os.Exit(1)
	}

	fmt.Print(blob.Hash)
}
