package graphapite

import "strings"

type Target struct {
	Args       []string
	IsFunction bool
	IsPattern  bool
	Name       string
	Pattern    string
}

func (t *Target) Parse(target string) error {
	target = strings.TrimSpace(target)
	t.IsFunction = false
	t.IsPattern = false
	if name, args, ok := targetFunction(target); ok {
		t.Name = name
		t.Args = args
		t.IsFunction = true
		return nil
	}
	t.IsPattern = true
	t.Pattern = target
	return nil
}

func targetFunction(target string) (name string, args []string, ok bool) {
	lParen := strings.Index(target, "(")
	// lParen must exist and must not be in first position
	if lParen == -1 || lParen == 0 {
		return
	}
	// must end in a )
	if !strings.HasSuffix(target, ")") {
		return
	}
	name = target[:lParen]
	args = targetArgs(target[lParen+1 : len(target)-1])
	ok = true
	return
}

func targetArgs(in string) (args []string) {
	context := 0
	last := 0
	for i, c := range in {
		switch c {
		case '{', '(':
			context++
		case ')', '}':
			context--
		case ',':
			if context == 0 {
				args = append(args, strings.TrimSpace(in[last:i]))
				last = i + 1
			}
		}
	}
	args = append(args, strings.TrimSpace(in[last:]))
	return
}
