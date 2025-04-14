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

func (*FileJS) isFile_Type() {}

type FileCSS struct {
	Path      string
	LocalPath string
}

func (*FileCSS) isFile_Type() {}

type FileUnkown struct {
	Path      string
	LocalPath string
}

func (*FileUnkown) isFile_Type() {}

type Resolver interface {
	Resolve(ctx context.Context, name, version string) (*Package, error)
}
