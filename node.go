package graphapite

import (
	"encoding/json"
	"strings"
)

type Node struct {
	Path []string
	Name string
	Leaf bool
}

// MarshalJSON allows us to override how we actually write this object
// to JSON. We do this because the json graphite wants makes no domain sense.
// https://github.com/graphite-project/graphite-web/blob/master/webapp/graphite/metrics/views.py
func (n Node) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"id":            strings.Join(n.Path, "."),
		"text":          n.Name,
		"allowChildren": !n.Leaf,
		"expandable":    !n.Leaf,
		"leaf":          n.Leaf,
	}
	return json.Marshal(m)
}
