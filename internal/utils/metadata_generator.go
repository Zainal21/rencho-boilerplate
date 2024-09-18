package utils

import (
	"strings"
	"time"

	"github.com/lucsky/cuid"
)

type Metadata struct {
	UID       func() string
	Slug      func(str string) string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func GenerateMetadata() Metadata {
	now := time.Now().UTC()

	return Metadata{
		UID: func() string {
			return cuid.New()
		},
		Slug: func(str string) string {
			return strings.ToLower(strings.Join(strings.Split(str, " "), "-"))
		},
		CreatedAt: now,
		UpdatedAt: now,
	}
}
