package jsdeliver

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/katallaxie/fiber-htmx/imports"
)

const (
	DefaultUrl    = "https://data.jsdelivr.com/v1/package/npm/%s@%s"
	DefaultCdnUrl = "https://cdn.jsdelivr.net/npm/%s@%s/%s"
)

var _ imports.Resolver = (*client)(nil)

type client struct{}

type File struct {
	Type  string `json:"type"`
	Name  string `json:"name"`
	Hash  string `json:"hash"`
	Size  int    `json:"size"`
	Files Files  `json:"files,omitempty"`
}

type Files []File

// Response is the response from the unpkg API.
type Response struct {
	Default string `json:"default"`
	Files   Files  `json:"files"`
}

// New returns a new unpkg provider.
func New() imports.Resolver {
	return &client{}
}

// Resolve resolves the package to a URL.
func (c *client) Resolve(ctx context.Context, pkg *imports.ExactPackage) error {
	url := fmt.Sprintf(DefaultUrl, pkg.Name, pkg.Version)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
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

	walkFiles(meta.Files, "", pkg)

	return nil
}

func walkFiles(files Files, filePath string, pkg *imports.ExactPackage) {
	for _, file := range files {
		if file.Type == "directory" {
			walkFiles(file.Files, filepath.Join(filePath, file.Name), pkg)
			continue
		}

		switch filepath.Ext(file.Name) {
		case ".js":
			pkg.Files = append(pkg.Files, &imports.FileJS{
				Path:      fmt.Sprintf(DefaultCdnUrl, pkg.Name, pkg.Version, filepath.Clean(filepath.Join(filePath, file.Name))),
				Integrity: file.Hash,
				LocalPath: filepath.Join(filePath, file.Name),
			})
		case ".css":
			pkg.Files = append(pkg.Files, &imports.FileCSS{
				Path:      fmt.Sprintf(DefaultCdnUrl, pkg.Name, pkg.Version, filepath.Join(filePath, file.Name)),
				Integrity: file.Hash,
				LocalPath: filepath.Join(filePath, file.Name),
			})
		default:
			pkg.Files = append(pkg.Files, &imports.FileUnkown{
				Path:      fmt.Sprintf(DefaultCdnUrl, pkg.Name, pkg.Version, filepath.Join(filePath, file.Name)),
				Integrity: file.Hash,
				LocalPath: filepath.Join(filePath, file.Name),
			})
		}
	}
}
