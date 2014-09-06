package graphapite

import "github.com/supershabam/graphapite/structs"

type Transformer interface {
	Transform([]structs.Series) ([]structs.Series, error)
}
