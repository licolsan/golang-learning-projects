package models

// Article - this is model of Article
type Article struct {
	ID      string `json:"id"`
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"Content"`
}
