package graphapite

import "time"

// A SeriesFn is a type that can be called with args to return an array of Series
type SeriesFn func(r Resolver, args []string, from, until time.Time) ([]Series, error)

// A Resolver turns a target string into a SeriesFn which can later be called to
// get an array of Series.
type Resolver interface {
	Resolve(target Target, from, until time.Time) ([]Series, error)
}
