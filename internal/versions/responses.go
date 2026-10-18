package versions

type VersionResponse struct {
	Environment string `json:"environment"`
	Title       string `json:"title"`
	Version     string `json:"version"`
}
