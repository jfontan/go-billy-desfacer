package desfacer

import (
	"github.com/spf13/afero"
	billy "gopkg.in/src-d/go-billy.v4"
)

var _ billy.File = new(File)

// File wraps an afero.File.
type File struct {
	path string
	a    afero.File
}

// NewFile creates a new File.
func NewFile(path string, a afero.File) *File {
	return &File{
		path: path,
		a:    a,
	}
}

// Name implements billy.File interface.
func (f *File) Name() string {
	return f.path
}

// Write implements billy.File interface.
func (f *File) Write(p []byte) (n int, err error) {
	return f.a.Write(p)
}

// Read implements billy.File interface.
func (f *File) Read(p []byte) (n int, err error) {
	return f.a.Read(p)
}

// ReadAt implements billy.File interface.
func (f *File) ReadAt(p []byte, off int64) (n int, err error) {
	return f.a.ReadAt(p, off)
}

// Seek implements billy.File interface.
func (f *File) Seek(offset int64, whence int) (int64, error) {
	return f.a.Seek(offset, whence)
}

// Close implements billy.File interface.
func (f *File) Close() error {
	return f.a.Close()
}

// Lock implements billy.File interface.
func (f *File) Lock() error {
	return nil
}

// Unlock implements billy.File interface.
func (f *File) Unlock() error {
	return nil
}

// Truncate implements billy.File interface.
func (f *File) Truncate(size int64) error {
	return f.a.Truncate(size)
}
