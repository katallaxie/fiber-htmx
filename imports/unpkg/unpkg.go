package unpkg

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/katallaxie/fiber-htmx/imports"
)

const DefaultUrl = "https://unpkg.com/%s@%s/?meta"
const DefaultCdnUrl = "https://unpkg.com/%s@%s/"

var _ imports.Resolver = (*client)(nil)

type client struct{}

// Response is the response from the unpkg API.
type Response struct {
	Package string `json:"package"`
	Version string `json:"version"`
	Prefix  string `json:"prefix"`
	Files   []struct {
		Path      string `json:"path"`
		Size      int    `json:"size"`
		Integrity string `json:"integrity"`
		Type      string `json:"type"`
	} `json:"files"`
}

// New returns a new unpkg provider.
func New() *client {
	return &client{}
}

// Resolve resolves the package to a URL.
func (c *client) Resolve(ctx context.Context, name, version string) (*imports.Package, error) {
	metaUrl := fmt.Sprintf(DefaultUrl, name, version)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, metaUrl, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unpkg API responded with status %d", resp.StatusCode)
	}

	var meta Response
	if err := json.NewDecoder(resp.Body).Decode(&meta); err != nil {
		return nil, err
	}

	pkg := &imports.Package{
		Name:    meta.Package,
		Version: meta.Version,
	}

	for _, f := range meta.Files {
		switch filepath.Ext(f.Path) {
		case ".js":
			pkg.Files = append(pkg.Files, &imports.FileJS{
				Path: f.Path,
			})
		case ".css":
			pkg.Files = append(pkg.Files, &imports.FileCSS{
				Path: f.Path,
			})
		default:
			pkg.Files = append(pkg.Files, &imports.FileUnkown{
				Path: f.Path,
			})
		}
	}

	return pkg, nil
}
