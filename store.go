package graphapite

import "time"

type Store interface {
	Get(key Key, start, end time.Time) ([]Datapoint, error)
	Write(key Key, datapoint Datapoint) error
}
