package models

// RequestedData represents the main Github API response
type RequestedData struct {
	Count      int        `json:"total_count"`
	Incomplete bool       `json:"incomplete_results"`
	Items      []RepoInfo `json:"items"`
}

// RepoInfo represents a repository it's the item of interest in the JSON response
// of the Github api.
type RepoInfo struct {
	// Name of the repository
	Name string `json:"name"`
	// Link to the Github repository
	URL string `json:"html_url"`
	// Number of stargazers
	Stars int `json:"stargazers_count"`
	// Repo defined language
	Language string `json:"language,omitempty"`
}
