package graphapite

// A Pattern is a graphite key that may or may not contain * characters
type Pattern string

// Turns a pattern into a list of keys
func (p Pattern) Keys(store Store) ([]Key, error) {
}

// "1.2.3.*".Expand() => []string{"1.2.3.4", "1.2.3.four"}
// "1.*.3.*".Expand() => []string{"1.*".Expand()}
