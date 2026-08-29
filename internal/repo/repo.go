package repo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Package struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type Index struct {
	Repository string    `json:"repository"`
	Version    int       `json:"version"`
	Packages   []Package `json:"packages"`
}

const indexURL = "https://repo.yosi-repo.ru/v1/index.json"

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

func FetchIndex() (Index, error) {
	resp, err := httpClient.Get(indexURL)
	if err != nil {
		return Index{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Index{}, fmt.Errorf("repository returned HTTP status %d", resp.StatusCode)
	}
	var index Index
	err = json.NewDecoder(resp.Body).Decode(&index)
	if err != nil {
		return Index{}, err
	}

	return index, nil
}
