package model

// Article ...
type Article struct {
	ID          string
	Title       string
	Description string
	CreatedAt   int64
	UpdatedAt   int64
	Versions    ArticleVersions
	Tags        []Tag
}

// ArticleVersions ...
type ArticleVersions struct {
	Current int64
}

// Tag ...
type Tag struct {
	Name      string
	CreatedAt int64
}
