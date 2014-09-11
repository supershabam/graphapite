package graphapite

// A SeriesFn is a type that can be called with args to return an array of Series
type SeriesFn interface {
	Execute(args []string) ([]Series, error)
}

// A Resolver turns a target string into a SeriesFn which can later be called to
// get an array of Series.
type Resolver interface {
	Resolve(rawtarget string) ([]Series, error)
}
