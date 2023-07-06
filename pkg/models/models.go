package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: can't find appropriate snippet")

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
