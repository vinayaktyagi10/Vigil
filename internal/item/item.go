package item

import (
	"time"
)

type Item struct {
	Title string
	URL	  string
	Source string
	PublishedAt time.Time
	RawContent string
}
