package boltdb

import (
	"github.com/gabesullice/phargo/storage"
)

type mockRunner func() (storage.Fetcher, error)
