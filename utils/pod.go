package utils

import (
	"fmt"
	"path/filepath"

	"github.com/hashicorp/go-version"
)

type Poddir struct {
	Path    string
	Name    string
	Version *version.Version
}

func NewPoddir(path string) *Poddir {
	poddir := Poddir{
		Path: path,
		Name: filepath.Base(filepath.Dir(path)),
	}

	if v, err := version.NewVersion(filepath.Base(path)); err != nil {
		poddir.Version, _ = version.NewVersion("0.0.0")
	} else {
		poddir.Version = v
	}

	return &poddir
}

func (p *Poddir) String() string {
	size, _ := DirSize(p.Path)
	return fmt.Sprintf("%v %v, %.2f MB", p.Name, p.Version, SizeToMB(size))
}

func (p *Poddir) Size() float64 {
	size, _ := DirSize(p.Path)
	return SizeToMB(size)
}
