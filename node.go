package graphapite

type Node struct {
	Path []string
	Name string
	Leaf bool
}

func (n Node) Key() Key {
}
