package graphapite

import "fmt"

type TargetFn func() ([]Series, error)

type Targeter struct {
}

var NullTargetFn = func() (series []Series, err error) {
	err = fmt.Errorf("NOT IMPLEMENTED")
	return
}

func (t Targeter) Parse(rawtarget string) (TargetFn, error) {
	if name, args, ok := t.fnArgs(rawtarget); ok {
		return t.parseFunction(name, args)
	}
	if pattern, ok := t.pattern(rawtarget); ok {
		return t.fetchPattern(pattern)
	}
	return NullTargetFn, fmt.Errorf("could not parse target")
}

func (t Targeter) parseFunction(name string, args []string) (TargetFn, error) {
	return NullTargetFn, fmt.Errorf("NOT IMPLEMENTED")
}

func (t Targeter) fnArgs(rawtarget string) (name string, args []string, ok bool) {
	return "", []string{}, false
}

func (t Targeter) fetchPattern(pattern string) (TargetFn, error) {
	return NullTargetFn, fmt.Errorf("NOT IMPLEMENTED")
}
