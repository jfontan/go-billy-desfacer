package desfacer

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/afero"
	"gopkg.in/src-d/go-billy.v4"
	"gopkg.in/src-d/go-billy.v4/helper/chroot"
	"gopkg.in/src-d/go-billy.v4/util"
)

const defaultDirectoryMode = 0755

var (
	_ billy.Basic      = new(FS)
	_ billy.Dir        = new(FS)
	_ billy.Filesystem = new(FS)

	// ErrNotImplemented is returned when the function is not implemented.
	ErrNotImplemented = fmt.Errorf("functionality not implemented")
)

type FS struct {
	a    afero.Fs
	path string
}

func New(fs afero.Fs) *FS {
	return NewPath(fs, "/")
}

func NewPath(fs afero.Fs, path string) *FS {
	return &FS{
		a:    fs,
		path: path,
	}
}

// Create implements billy.Basic interface.
func (f *FS) Create(filename string) (billy.File, error) {
	if err := f.createDir(filename); err != nil {
		return nil, err
	}

	file, err := f.a.Create(filename)
	if err != nil {
		return nil, err
	}

	return NewFile(filename, file), nil
}

// Open implements billy.Basic interface.
func (f *FS) Open(filename string) (billy.File, error) {
	file, err := f.a.Open(filename)
	if err != nil {
		return nil, err
	}

	return NewFile(filename, file), nil
}

// OpenFile implements billy.Basic interface.
func (f *FS) OpenFile(filename string, flag int, perm os.FileMode) (billy.File, error) {
	if flag&os.O_CREATE == os.O_CREATE {
		if err := f.createDir(filename); err != nil {
			return nil, err
		}
	}

	file, err := f.a.OpenFile(filename, flag, perm)
	if err != nil {
		return nil, err
	}

	return NewFile(filename, file), nil
}

// Stat implements billy.Basic interface.
func (f *FS) Stat(filename string) (os.FileInfo, error) {
	return f.a.Stat(filename)
}

// Rename implements billy.Basic interface.
func (f *FS) Rename(oldpath string, newpath string) error {
	if err := f.createDir(newpath); err != nil {
		return err
	}

	return f.a.Rename(oldpath, newpath)
}

// Remove implements billy.Basic interface.
func (f *FS) Remove(filename string) error {
	return f.a.Remove(filename)
}

// Join implements billy.Basic interface.
func (f *FS) Join(elem ...string) string {
	return filepath.Join(elem...)
}

// ReadDir implements billy.Dir interface.
func (f *FS) ReadDir(path string) ([]os.FileInfo, error) {
	d, err := f.a.Open(path)
	if err != nil {
		return nil, err
	}

	defer d.Close()

	return d.Readdir(0)
}

// MkdirAll implements billy.Dir interface.
func (f *FS) MkdirAll(filename string, perm os.FileMode) error {
	return f.a.MkdirAll(filename, perm)
}

func (f *FS) createDir(fullpath string) error {
	dir := filepath.Dir(fullpath)
	if dir != "." {
		if err := f.MkdirAll(dir, defaultDirectoryMode); err != nil {
			return err
		}
	}

	return nil
}

// Lstat implements billy.Symlink interface.
func (f *FS) Lstat(filename string) (os.FileInfo, error) {
	if s, ok := f.a.(afero.Lstater); ok {
		stat, _, err := s.LstatIfPossible(filename)
		return stat, err
	}

	return f.a.Stat(filename)
}

// Symlink implements billy.Symlink interface.
func (f *FS) Symlink(target string, link string) error {
	return ErrNotImplemented
}

// Readlink implements billy.Symlink interface.
func (f *FS) Readlink(link string) (string, error) {
	return "", ErrNotImplemented
}

// TempFile implements billy.TempFile interface.
func (f *FS) TempFile(dir string, prefix string) (billy.File, error) {
	return util.TempFile(f, dir, prefix)
}

// Chroot implements billy.Chroot interface.
func (f *FS) Chroot(path string) (billy.Filesystem, error) {
	return chroot.New(f, path), nil
}

// Root implements billy.Chroot interface.
func (f *FS) Root() string {
	return f.path
}
