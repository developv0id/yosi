package repo

type Package struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type Index struct {
	Repository string    `json:"repository"`
	Version    string    `json:"version"`
	Packages   []Package `json:"packages"`
}
