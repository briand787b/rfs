package models

// MediaType is a type of media
type MediaType struct {
	// DB-mapped
	ID   int    `json:"id"`
	Name string `json:"name"`
}
