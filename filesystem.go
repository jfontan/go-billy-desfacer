package desfacer

import (
	"os"
	"path/filepath"

	"github.com/spf13/afero"
	"gopkg.in/src-d/go-billy.v4"
)

const defaultDirectoryMode = 0755

var _ billy.Basic = new(FS)
var _ billy.Dir = new(FS)

type FS struct {
	a afero.Fs
}

func New(fs afero.Fs) *FS {
	return &FS{
		a: fs,
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
