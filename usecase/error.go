package usecase

import "fmt"

var (
	// ErrNotFound ...
	ErrNotFound = fmt.Errorf("Not found")
	// ErrAlreadyExists ...
	ErrAlreadyExists = fmt.Errorf("Already exists")
)
