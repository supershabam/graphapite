package graphapite

import (
	"fmt"
	"strings"
	"time"
)

// Args => (Target, "string") string in quotes. Don't ask me, it's what graphite
// desires.
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

	name := args[1]
	name = strings.TrimSpace(name)
	if !strings.HasPrefix(name, "\"") || !strings.HasSuffix(name, "\"") {
		err = fmt.Errorf("name must be quoted")
		return
	}
	name = strings.TrimPrefix(name, "\"")
	name = strings.TrimSuffix(name, "\"")

	in, err := r.Resolve(target, from, until)
	if err != nil {
		return
	}
	for _, series := range in {
		out = append(out, Series{
			Name:       name,
			Datapoints: series.Datapoints,
		})
	}
	return
}
