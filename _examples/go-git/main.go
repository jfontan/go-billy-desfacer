package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/cache"
	"github.com/go-git/go-git/v5/storage/filesystem"
	"github.com/spf13/afero"
	desfacer "gopkg.in/jfontan/go-billy-desfacer.v0"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: main <path to git repo>")
		os.Exit(255)
	}
	path := os.Args[1]

	workdir := afero.NewBasePathFs(afero.NewOsFs(), path)
	gitdir := afero.NewBasePathFs(afero.NewOsFs(), filepath.Join(path, ".git"))

	billywork := desfacer.New(workdir)
	billygit := desfacer.New(gitdir)

	storage := filesystem.NewStorage(billygit, cache.NewObjectLRUDefault())

	repo, err := git.Open(storage, billywork)
	if err != nil {
		panic(err)
	}

	head, err := repo.Head()
	if err != nil {
		panic(err)
	}

	commit, err := repo.CommitObject(head.Hash())
	if err != nil {
		panic(err)
	}

	fmt.Println(commit.String())
}
