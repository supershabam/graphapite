package graphapite

import "github.com/supershabam/graphapite/structs"

type NewTransformer func([]string) (Transformer, string, error)

type Transformer interface {
	Transform([]structs.Series) ([]structs.Series, error)
}
