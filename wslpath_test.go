// +build linux

package wslpath

import (
	"testing"

	"github.com/matryer/is"
)

func TestFromWindows(t *testing.T) {
	is := is.New(t)
	path := "C:\\Users\\user\\file.txt"
	lpath, err := FromWindows(path)
	is.NoErr(err)
	is.Equal(lpath, "/mnt/c/Users/user/file.txt")
}

func TestToWindows(t *testing.T) {
	is := is.New(t)
	path := "/mnt/c/Users/user/file.txt"
	wpath, err := ToWindows(path)
	is.NoErr(err)
	is.Equal(wpath, "C:\\Users\\user\\file.txt")
}
