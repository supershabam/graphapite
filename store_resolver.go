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
func (r StoreResolver) Resolve(target Target, from, until time.Time) ([]Series, error) {
	if target.IsFunction {
		if fn, ok := r.Functions[target.Name]; ok {
			return fn(r, target.Args, from, until)
		}
		return []Series{}, fmt.Errorf("function %s not found", target.Name)
	}
	datapoints, err := r.Store.Get(Key(target.Pattern), from, until)
	if err != nil {
		return []Series{}, err
	}
	return []Series{
		Series{
			Name:       target.Pattern,
			Datapoints: datapoints,
		},
	}, nil
}
