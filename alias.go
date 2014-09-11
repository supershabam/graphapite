package graphapite

import (
	"fmt"
	"time"
)

// Alias provides the `alias(SeriesFn, name)` method of graphite
func Alias(r Resolver, args []string, from, until time.Time) (out []Series, err error) {
	if len(args) != 2 {
		err = fmt.Errorf("alias: expected 2 arguments but got %d", len(args))
		return
	}
	var target Target
	err = target.Parse(args[0])
	if err != nil {
		return
	}
	in, err := r.Resolve(target, from, until)
	if err != nil {
		return
	}
	for _, series := range in {
		out = append(out, Series{
			Name:       args[1],
			Datapoints: series.Datapoints,
		})
	}
	return
}
