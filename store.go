package graphapite

import (
	"time"
)

type Store interface {
	Get(key string, start, end time.Time) ([]Datapoint, error)
	Nodes(pattern string) ([]Node, error)
	Write(key string, datapoint Datapoint) error
}
