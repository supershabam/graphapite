package graphapite

import "fmt"

type StoreResolver struct {
	Functions map[string]SeriesFn
	Store     Store
}

// Resolve turns a target string into a list of series
func (r StoreResolver) Resolve(rawtarget string) ([]Series, error) {
	return []Series{}, fmt.Errorf("NOT IMPLEMENTED")
}
