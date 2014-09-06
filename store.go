package graphapite

import (
	"time"

	"github.com/supershabam/graphapite/structs"
)

type Store interface {
	Get(key structs.Key, start, end time.Time) ([]structs.Datapoint, error)
	Nodes(pattern structs.Pattern) ([]structs.Node, error)
	Write(key structs.Key, datapoint structs.Datapoint) error
}
