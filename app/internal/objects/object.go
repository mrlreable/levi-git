package objects

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"os"
	"path"
)

const ROOT = ".git"
const HEAD = ".git/HEAD"
const OBJECT_STORE = ".git/objects/"
const REFS_STORE = ".git/refs/"

// Path: .git/objects/<hash first 2 characters>/<hash rest of the characters>
type Object interface {
	// Creates the object (blob, tree) and writes it to the objects directory
	Write() error
	Type() string
}

// <size>/0<content>
type Blob struct {
	Size    uint
	Content string
	Hash    string
}

func (b *Blob) Write() error {

	return nil
}

func (b *Blob) Type() string {
	return "blob"
}

func Read(p string) (Blob, error) {
	prefix := p[:2]
	file := p[2:]

	f, err := os.Open(path.Join(OBJECT_STORE, prefix, file))
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
	defer f.Close()

	reader, err := zlib.NewReader(f)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
	defer reader.Close()

	var out bytes.Buffer
	_, err = io.Copy(&out, reader)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	// eg. blob 11\0Hello world
	bytes := out.Bytes()
	var content string
	var size uint
	for i, b := range bytes {
		if b == 0 {
			content = string(bytes[i+1:])
			size = uint(len(content))
		}
	}

	b := Blob{
		Content: content,
		Size:    size,
	}

	return b, nil
}

type Tree struct {
}

func (t *Tree) Write() error {
	return nil
}

func (t *Tree) Type() string {
	return "tree"
}
