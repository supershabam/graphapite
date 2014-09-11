package graphapite

import "time"

type MockResolver struct {
	ResolveFn func(target string, from, until time.Time) ([]Series, error)
}

func (r MockResolver) Resolve(target string, from, until time.Time) ([]Series, error) {
	return r.ResolveFn(target, from, until)
}
