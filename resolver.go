package graphapite

// ResolverFn can be executed to return an array of series.
type ResolverFn func(args []string) ([]Series, error)

type Function interface {
	Execute(args []string) ([]Series, error)
}

// A Resolver turns a target string into a SeriesFn which can later be called to
// get an array of Series.
type Resolver struct {
	Functions map[string]Function
	Store     Store
}

// Resolve turns a target string into a list of series
func (r Resolver) Resolve(rawtarget string) ([]Series, error) {
}
