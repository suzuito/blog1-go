package model

// ArticleID ...
type ArticleID string

// Article ...
type Article struct {
	ID          ArticleID
	Title       string
	Description string
	CreatedAt   int64
	UpdatedAt   int64
	PublishedAt int64
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
