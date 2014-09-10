package graphapite

import "fmt"

type Alias struct {
	Resolver Resolver
}

func (a Alias) Execute(args []string) ([]Series, error) {
	if len(args) != 2 {
		err = fmt.Errorf("alias: expected 2 arguments but got %d", len(args))
		return
	}
	series, err = a.Resolver.Resolve(args[0])
	if err != nil {
		return
	}
	for _, series := range series {
		series.Name = args[1]
	}
	return
}
