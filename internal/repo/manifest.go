package repo

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type Manifest struct {
	ManifestVersion int                  `toml:"manifest_version"`
	Package         ManifestPackage      `toml:"package"`
	Source          ManifestSource       `toml:"source"`
	Dependencies    ManifestDependencies `toml:"dependencies"`
	Build           ManifestBuild        `toml:"build"`
	Install         ManifestInstall      `toml:"install"`
	Metadata        ManifestMetadata     `toml:"metadata"`
}

type ManifestPackage struct {
	Name        string `toml:"name"`
	Version     string `toml:"version"`
	Description string `toml:"description"`
	License     string `toml:"license"`
	Homepage    string `toml:"homepage"`
	Maintainer  string `toml:"maintainer"`
}

type ManifestSource struct {
	URL    string `toml:"url"`
	SHA256 string `toml:"sha256"`
	Size   int64  `toml:"size"`
}

type ManifestDependencies struct {
	Runtime   []string `toml:"runtime"`
	Build     []string `toml:"build"`
	Optional  []string `toml:"optional"`
	Conflicts []string `toml:"conflicts"`
}

type ManifestBuild struct {
	System string `toml:"system"`
	Recipe string `toml:"recipe"`
}

type ManifestInstall struct {
	Prefix string `toml:"prefix"`
}

type ManifestMetadata struct {
	Architecture []string `toml:"architecture"`
}

func FetchManifest(name, version string) (Manifest, error) {
	url := fmt.Sprintf(
		"https://repo.yosi-repo.ru/v1/packages/%s/%s/manifest.toml",
		name,
		version,
	)

	resp, err := httpClient.Get(url)
	if err != nil {
		return Manifest{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return Manifest{}, fmt.Errorf(
			"repository returned HTTP status %d",
			resp.StatusCode,
		)
	}

	var manifest Manifest

	_, err = toml.NewDecoder(resp.Body).Decode(&manifest)
	if err != nil {
		return Manifest{}, err
	}

	return manifest, nil
}
