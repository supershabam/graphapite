package graphapite

import (
	"fmt"
	"time"
)

type StoreResolver struct {
	Functions map[string]SeriesFn
	Store     Store
}

// Resolve turns a target string into a list of series
func (r StoreResolver) Resolve(target string, from, until time.Time) ([]Series, error) {
	return []Series{}, fmt.Errorf("NOT IMPLEMENTED")
}
