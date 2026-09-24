package repo

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type BuildRecipe struct {
	RecipeVersion string            `toml:"recipe_version"`
	Requirements  BuildRequirements `toml:"requirements"`
	Policy        BuildPolicy       `toml:"policy"`
	Configure     BuildStage        `toml:"configure"`
	Build         BuildStage        `toml:"build"`
	Test          BuildStage        `toml:"test"`
	Install       BuildStage        `toml:"install"`
}

type BuildRequirements struct {
	Commands []string `toml:"commands"`
}

type BuildPolicy struct {
	Network bool `toml:"network"`
}

type BuildStage struct {
	Steps []BuildStep `toml:"steps"`
}

type BuildStep struct {
	Command          string            `toml:"command"`
	Args             []string          `toml:"args"`
	WorkingDirectory string            `toml:"working_directory"`
	Environment      map[string]string `toml:"environment"`
	DestDir          bool              `toml:"destdir"`
}

func FetchBuildRecipe(name, version, file string) (BuildRecipe, error) {
	url := fmt.Sprintf(
		"https://repo.yosi-repo.ru/v1/packages/%s/%s/%s",
		name,
		version,
		file,
	)

	resp, err := httpClient.Get(url)
	if err != nil {
		return BuildRecipe{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return BuildRecipe{}, fmt.Errorf(
			"repository returned HTTP status %d",
			resp.StatusCode,
		)
	}

	var recipe BuildRecipe

	_, err = toml.NewDecoder(resp.Body).Decode(&recipe)
	if err != nil {
		return BuildRecipe{}, err
	}

	return recipe, nil
}
