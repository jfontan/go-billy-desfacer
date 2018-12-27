# go-billy-desfacer [![GoDoc](https://godoc.org/gopkg.in/jfontan/go-billy-desfacer.v0?status.svg)](https://godoc.org/github.com/jfontan/go-billy-desfacer)[![Build Status](https://travis-ci.com/jfontan/go-billy-desfacer.svg?branch=master)](https://travis-ci.com/jfontan/go-billy-desfacer)[![codecov](https://codecov.io/gh/jfontan/go-billy-desfacer/branch/master/graph/badge.svg)](https://codecov.io/gh/jfontan/go-billy-desfacer)

[go-billy](https://github.com/src-d/go-billy) filesystem that wraps [afero](https://github.com/spf13/afero). It lets use afero filesystems with software that expects go-billy, for example with [go-git](https://github.com/src-d/go-git).


# Installation

```
go get gopkg.in/jfontan/go-billy-desfacer.v0
```

# Example of use

```go
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
```

# Notes

* The functions `Symlink` and `Readlink` are not implemented as afero does not have that functionality.


# The name

"desfacer" means "to undo" or "to unmake" in Galician and old Spanish.

# License

Apache License Version 2.0, see [LICENSE](LICENSE).

