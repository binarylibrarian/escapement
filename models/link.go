package models

// Link represents the links users are keeping
type Link struct {
	ID     string   `json:"id"`
	URL    string   `json:"url"`
	Read   bool     `json:"read"`
	Shared []string `json:"shared"`
	Notes  string   `json:"notes"`
}
