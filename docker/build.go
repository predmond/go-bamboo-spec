package docker

import (
	"strings"
)

type Build struct {
	path string
	tag  string
}

func NewBuild(path string) *Build {
	return &Build{
		path: path,
	}
}

func (b *Build) String() string {
	args := make([]string, 0, 2)
	args = append(args, "docker", "build")
	if len(b.tag) > 0 {
		args = append(args, "--tag", b.tag)
	}
	args = append(args, b.path)
	return strings.Join(args, " ")
}

func (b *Build) WithTag(tag string) *Build {
	b.tag = tag
	return b
}
