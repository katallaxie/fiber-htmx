package imports

import "context"

type Package struct {
	Name    string
	Version string
	Files   []isFile_Type
}

type isFile_Type interface {
	isFile_Type()
}

type FileJS struct {
	Path      string
	LocalPath string
}

func (f *FileJS) isFile_Type() {}

type FileCSS struct {
	Path      string
	LocalPath string
}

func (f *FileCSS) isFile_Type() {}

type FileUnkown struct {
	Path      string
	LocalPath string
}

func (f *FileUnkown) isFile_Type() {}

type Resolver interface {
	Resolve(ctx context.Context, pkg *Package) error
}
