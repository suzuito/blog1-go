package usecase

// Usecase ...
type Usecase interface {
	SyncArticles(
		source ArticleReader,
	) error
}

// Impl ...
type Impl struct {
	db DB
}

// NewImpl ...
func NewImpl(db DB) *Impl {
	return &Impl{
		db: db,
	}
}
