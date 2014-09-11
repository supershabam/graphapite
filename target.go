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
	rParen := strings.LastIndex(target, ")")
	if lParen == -1 || rParen == -1 || lParen == 0 {
		return
	}
	name = target[:lParen]
	args = targetArgs(target[lParen+1 : rParen])
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
