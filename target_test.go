package graphapite

import (
	"fmt"
	"strings"
	"testing"
)

func TestTarget(t *testing.T) {
	var target Target

	functionTests := []struct {
		Name string
		Args []string
	}{
		{
			Name: "alias",
			Args: []string{
				"sub(some.{multi,value}.key)",
				"name",
			},
		},
	}

	for _, ftest := range functionTests {
		raw := fmt.Sprintf("%s(%s)", ftest.Name, strings.Join(ftest.Args, ", "))
		err := target.Parse(raw)
		if err != nil {
			t.Fatal(err)
		}
		if !target.IsFunction {
			t.Fatalf("expected target to be a function target")
		}
		if target.Name != ftest.Name {
			t.Fatalf("expected target name to be %s not %s", ftest.Name, target.Name)
		}
		if len(target.Args) != len(ftest.Args) {
			t.Fatalf("expected %d arguments, not %d", len(ftest.Args), len(target.Args))
		}
		for i, arg := range target.Args {
			if ftest.Args[i] != arg {
				t.Fatalf("expected argument %d to be %s not %s", i, ftest.Args[i], arg)
			}
		}
	}
}
