package model

import "time"

type Category struct {
	ID               string
	Version          uint64
	Name             string
	Enabled          bool
	CreatedDate      time.Time
	LastModifiedDate time.Time
}
