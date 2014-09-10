package graphapite

// SeriesFn can be executed to return an array of series.
type SeriesFn func() ([]Series, error)

// A Resolver turns a target string into a SeriesFn which can later be called to
// get an array of Series.
type Resolver interface {
	Resolve(rawtarget string) (SeriesFn, error)
}
