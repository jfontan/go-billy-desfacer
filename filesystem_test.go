package desfacer

import (
	"testing"

	"github.com/go-git/go-billy/v5/test"
	"github.com/go-git/go-billy/v5/util"
	"github.com/spf13/afero"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

var _ = Suite(&FilesystemSuite{})

type FilesystemSuite struct {
	test.BasicSuite
	test.TempFileSuite
	test.DirSuite
	test.ChrootSuite

	baseFS *FS
	tmp    string
}

func (s *FilesystemSuite) SetUpTest(c *C) {
	a := afero.NewOsFs()
	fs := New(a)
	s.baseFS = fs

	var err error
	s.tmp, err = util.TempDir(fs, "", "billy")
	c.Assert(err, IsNil)

	tmp, err := fs.Chroot(s.tmp)
	c.Assert(err, IsNil)

	s.BasicSuite.FS = tmp
	s.TempFileSuite.FS = tmp
	s.DirSuite.FS = tmp
	s.ChrootSuite.FS = tmp
}

func (s *FilesystemSuite) TearDownTest(c *C) {
	if s.baseFS != nil {
		err := util.RemoveAll(s.baseFS, s.tmp)
		c.Assert(err, IsNil)
	}
}
