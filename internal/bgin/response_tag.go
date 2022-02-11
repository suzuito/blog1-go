package bgin

import "github.com/suzuito/blog1-go/pkg/entity"

// ResponseTag ...
type ResponseTag struct {
	Name string `json:"name"`
}

// NewResponseTag ...
func NewResponseTag(a *entity.Tag) *ResponseTag {
	return &ResponseTag{
		Name: a.Name,
	}
}

// NewResponseTags ...
func NewResponseTags(a *[]entity.Tag) *[]ResponseTag {
	b := []ResponseTag{}
	for _, v := range *a {
		b = append(b, *NewResponseTag(&v))
	}
	return &b
}
