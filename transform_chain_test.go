package graphapite

import (
	"fmt"
	"testing"
)

func TestTransformArgs(t *testing.T) {
	tests := []struct {
		In  string
		Out []string
	}{
		In: "something(nested,with arguments, another(next, and arguments)), arg2",
		Out: []string{
			"something(nested,with arguments, another(next, and arguments))",
			"args2",
		},
	}

	for _, test := range tests {
		output := TransformArgs(test.In)
		if len(output) != len(test.Out) {
			t.Fatal(fmt.Errorf("test %d generated %+v when expecting %+v", output, test.Out))
		}
	}
}
