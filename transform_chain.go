package graphapite

import (
	"fmt"
	"strings"
)

var Transforms = map[string]NewTransformer{
	"alias": NewAliasTransformer,
}

// TransformChain parses a list of transformers that should be executed in-order
// to produce the result the target asks for.
func TransformChain(target string) (chain []Transformer, err error) {
	for {
		target = strings.TrimSpace(target)
		if name, args, ok := TranformParts(target); ok {
			if nt, ok := transforms[name]; ok {
				transformer, target, err := nt(args)
				if err != nil {
					return chain, err
				}
				chain = append(chain, transformer)
			}
			return chain, fmt.Errorf("no transform configured for: %s", name)
		}
		return chain, nil
	}
	panic("can't loop here")
}

// TransformParts returns the function name and the arguments if a function is found
// in the target string
func TransformParts(target string) (name string, args []string, ok bool) {
	lParen := strings.Index(target, "(")
	rParen := strings.LastIndex(target, ")")
	if lParen == -1 || rParent == -1 {
		return
	}
	name = target[:lParen]
	args = TransformArgs(target[lParen:rParen])
	ok = true
	return
}

func TransformArgs(rawargs string) []string {
	separators := []int{}
	context := 0
	for i, c := range rawargs {
		switch c {
		case "(":
			context++
		case "(":
			context--
		case ",":
			if context == 0 {
				separators = append(separators, i)
			}
		}
	}

	args := []string{}
	if len(separators) == 0 {
		args = append(args, rawargs)
		return args
	}
	lastSeparator := 0
	for _, separator := range separators {
		args = append(args, rawargs[lastSeparator:separator])
		lastSeparator = separator
	}
	args = append(args, rawargs[lastSeparator:])
	return args
}
