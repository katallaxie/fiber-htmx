package imports

// Resolver is the interface that must be implemented by all resolvers.
type Resolver interface {
	// Resolve resolves a package by its name and version.
	Resolve(pkg *ExactPackage) error
}

var _ Resolver = (*UnimplementedResolver)(nil)

// UnimplementedResolver is a resolver that is not implemented.
type UnimplementedResolver struct{}

// Resolve resolves a package by its name and version.
func (r *UnimplementedResolver) Resolve(pkg *ExactPackage) error {
	return nil
}

// Require is a struct that represents a package with its name and version.
type Require struct {
	// File is the file path of the package. This is the local path not the full url in the CDN.
	File string
	// Raw is the raw file path of the package.
	Raw string
	// As is the alias for the package.
	As string
}

// ExactPackage is a struct that represents a package with its name, version, and files.
type ExactPackage struct {
	// Name is the name of the package.
	Name string
	// Version is the version of the package.
	Version string
	// Files is a list of files associated with the package.
	Files []isFile_Type
}

type isFile_Type interface {
	isFile_Type()
}

// FileJS is a struct that represents a JavaScript file.
type FileJS struct {
	Path      string
	LocalPath string
	Integrity string
}

func (*FileJS) isFile_Type() {}

// FileCSS is a struct that represents a CSS file.
type FileCSS struct {
	Path      string
	LocalPath string
	Integrity string
}

func (*FileCSS) isFile_Type() {}

// FileUnkown is a struct that represents an unknown file type.
type FileUnkown struct {
	Path      string
	LocalPath string
	Integrity string
}

func (*FileUnkown) isFile_Type() {}
