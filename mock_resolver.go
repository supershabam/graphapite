package graphapite

type MockResolver struct {
	ResolveFn func(string) ([]Series, error)
}

func (r MockResolver) Resolve(rawtarget string) ([]Series, error) {
	return r.ResolveFn(rawtarget)
}
