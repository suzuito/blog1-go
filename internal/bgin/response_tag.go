package bgin

import "github.com/suzuito/blog1-go/internal/entity/model"

// ResponseTag ...
type ResponseTag struct {
	Name string `json:"name"`
}

// NewResponseTag ...
func NewResponseTag(a *model.Tag) *ResponseTag {
	return &ResponseTag{
		Name: a.Name,
	}
}

// NewResponseTags ...
func NewResponseTags(a *[]model.Tag) *[]ResponseTag {
	b := []ResponseTag{}
	for _, v := range *a {
		b = append(b, *NewResponseTag(&v))
	}
	return &b
}
