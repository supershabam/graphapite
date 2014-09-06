package graphapite

import "github.com/supershabam/graphapite/structs"

type AliasTransformer struct {
	NewName string
}

func (t AliasTransformer) Transform(in []structs.Series) ([]structs.Series, error) {
	out := make([]structs.Series, len(in))
	for _, s := range in {
		out = append(out, structs.Series{
			Name:                 t.NewName,
			TimesortedDatapoints: s.TimesortedDatapoints,
		})
	}
	return out, nil
}
