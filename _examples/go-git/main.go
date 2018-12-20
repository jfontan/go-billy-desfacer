package main

import (
	"os"
	"path/filepath"

	"github.com/spf13/afero"
	"gopkg.in/jfontan/go-billy-desfacer.v0"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/cache"
	"gopkg.in/src-d/go-git.v4/storage/filesystem"
)

func main() {
	path := filepath.Join(os.Getenv("GOPATH"), "src", "gopkg.in",
		"jfontan", "go-billy-desfacer.v0")

	workdir := afero.NewBasePathFs(afero.NewOsFs(), path)
	gitdir := afero.NewBasePathFs(afero.NewOsFs(), filepath.Join(path, ".git"))

	billywork := desfacer.New(workdir)
	billygit := desfacer.New(gitdir)

	storage := filesystem.NewStorage(billygit, cache.NewObjectLRUDefault())

	repo, err := git.Open(storage, billywork)
	if err != nil {
		panic(err)
	}
}
