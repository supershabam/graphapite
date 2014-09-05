package graphapite

// A Pattern is a graphite key that may or may not contain * characters
type Pattern string

// Match returns true iff the key matches the pattern
func (p Pattern) Match(key Key) bool {
	return false
}
