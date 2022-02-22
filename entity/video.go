package entity

type Video struct {
	Title string `json:"title" validate:"is-cool"`
	Description string `json:"description"`
	URL string `json:"url"`
}