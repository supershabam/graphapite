package graphapite

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Args => (Target, "string") string in quotes. Don't ask me, it's what graphite
// desires.
func AliasByNode(r Resolver, args []string, from, until time.Time) (out []Series, err error) {
	if len(args) != 2 {
		err = fmt.Errorf("alias: expected 2 arguments but got %d", len(args))
		return
	}

	var target Target
	err = target.Parse(args[0])
	if err != nil {
		return
	}

	index, err := strconv.Atoi(args[1])
	if err != nil {
		return
	}

	in, err := r.Resolve(target, from, until)
	if err != nil {
		return
	}
	for _, series := range in {
		parts := strings.Split(series.Name, ".")
		if len(parts) < index {
			err = fmt.Errorf("name does not have part %d", index)
			return
		}
		out = append(out, Series{
			Name:       parts[index],
			Datapoints: series.Datapoints,
		})
	}
	return
}
