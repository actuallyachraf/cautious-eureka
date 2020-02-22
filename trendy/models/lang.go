package models

// LanguageInfo represents the top repositories in a given language
type LanguageInfo struct {
	// Language we're looking for information about
	Language string `json:"language"`
	// Count of repositories in a given language
	Count int `json:"count"`
	// Links to the Githu repos for said language
	Links []string `json:"links"`
}
