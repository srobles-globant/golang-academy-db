// Package model defines the structs of the app's data model and the interface to persist them
package model

// Cart struct
type Cart struct {
	ID    int    `json:"id"`
	Owner string `json:"owner"`
	Items []Item `json:"items"`
}

// Item struct
type Item struct {
	ID        int      `json:"id"`
	ArticleID int      `json:"articleId"`
	Article   *Article `json:"-"`
	Quantity  int      `json:"quantity"`
}

// Article struct
type Article struct {
	ID    int     `json:"id"`
	Title string  `json:"title"`
	Price float32 `json:"price"`
}
