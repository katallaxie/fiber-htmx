package unpkg

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/katallaxie/fiber-htmx/imports"
)

const DefaultUrl = "https://unpkg.com/%s@%s/?meta"
const DefaultCdnUrl = "https://unpkg.com/%s@%s/"

var _ imports.Resolver = (*client)(nil)

type client struct{}

type Response struct {
	Type  string `json:"type"`
	Path  string `json:"path"`
	Files []struct {
		Path string `json:"path"`
		Type string `json:"type"`
	} `json:"files"`
}

// New returns a new unpkg provider.
func New() *client {
	return &client{}
}

// Resolve resolves the package to a URL.
func (c *client) Resolve(ctx context.Context, pkg *imports.Package) error {
	metaUrl := fmt.Sprintf(DefaultUrl, pkg.Name, pkg.Version)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, metaUrl, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unpkg API responded with status %d", resp.StatusCode)
	}

	var meta Response
	if err := json.NewDecoder(resp.Body).Decode(&meta); err != nil {
		return err
	}

	fmt.Println(meta)

	// basePath := fmt.Sprintf(DefaultCdnUrl, pkg.Name, pkg.Version)

	// // Recursively collect all files
	// var files library.Files
	// c.walkFiles(meta.Files, basePath, &files)

	// return files, version, nil

	return nil
}
