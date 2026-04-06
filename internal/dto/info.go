package dto

type Info struct {
	ServiceName string `json:"service"`
	Version     string `json:"version"`
	Description string `json:"description"`
	ApiDocsPath string `json:"api_docs_url"`
	Status      string `json:"running"`
}
