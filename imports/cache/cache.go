package cache

import (
	"github.com/hashicorp/golang-lru/v2/simplelru"
	"github.com/katallaxie/fiber-htmx/imports"
	"github.com/zeiss/pkg/errorx"
)

var _ imports.Resolver = (*cache)(nil)

var lru *simplelru.LRU[string, *imports.ExactPackage]

func init() {
	lru = errorx.Must(simplelru.NewLRU[string, *imports.ExactPackage](1000, nil))
}

type cache struct {
	resolver imports.Resolver
}

// New returns a new unpkg provider.
func New(resolver imports.Resolver) imports.Resolver {
	return &cache{resolver}
}

// Resolve resolves the package to a URL.
func (c *cache) Resolve(pkg *imports.ExactPackage) error {
	if v, ok := lru.Get(pkg.Hash()); ok {
		pkg.Files = v.Files
		return nil
	}

	if err := c.resolver.Resolve(pkg); err != nil {
		return err
	}

	lru.Add(pkg.Hash(), pkg)

	return nil
}
