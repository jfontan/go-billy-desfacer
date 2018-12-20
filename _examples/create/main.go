package main

import (
	"fmt"

	"github.com/spf13/afero"
	"gopkg.in/jfontan/go-billy-desfacer.v0"
)

func main() {
	// wrap an afero filesystem to billy interface
	aferofs := afero.NewMemMapFs()
	billyfs := desfacer.New(aferofs)

	// create a file with billy interface
	billyfile, err := billyfs.Create("file")
	if err != nil {
		panic(err)
	}

	_, err = billyfile.Write([]byte("some data"))
	if err != nil {
		panic(err)
	}
	_ = billyfile.Close()

	// read file directly in afero filesystem
	aferofile, err := aferofs.Open("file")
	if err != nil {
		panic(err)
	}

	buf := make([]byte, 32)
	n, err := aferofile.Read(buf)
	if err != nil {
		panic(err)
	}
	_ = aferofile.Close()

	fmt.Println(string(buf[:n]))
}
