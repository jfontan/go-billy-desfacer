package desfacer

import (
	"testing"

	"github.com/spf13/afero"
	. "gopkg.in/check.v1"
	"gopkg.in/src-d/go-billy.v4/helper/chroot"
	"gopkg.in/src-d/go-billy.v4/test"
	"gopkg.in/src-d/go-billy.v4/util"
)

func Test(t *testing.T) { TestingT(t) }

var _ = Suite(&FilesystemSuite{})

type FilesystemSuite struct {
	test.BasicSuite
	test.DirSuite

	FS  *FS
	tmp string
}

func (s *FilesystemSuite) SetUpTest(c *C) {
	a := afero.NewOsFs()
	fs := New(a)

	var err error
	s.tmp, err = util.TempDir(fs, "", "billy")
	c.Assert(err, IsNil)

	tmp := chroot.New(fs, s.tmp)
	s.BasicSuite.FS = tmp
	s.DirSuite.FS = tmp
}

func (s *FilesystemSuite) TearDownTest(c *C) {
	if s.FS != nil {
		err := util.RemoveAll(s.FS, s.tmp)
		c.Assert(err, IsNil)
	}
}
