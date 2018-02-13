package finder

type NodeTree struct {
	Nodes []Node
}

type Node struct {
	Value int
	Left *Node
	Right *Node
}

type Finder struct {}

func find(value int, current Node) bool {
	if (current == Node{}) {
		return false;
	}

	if current.Value == value {
		return true;
	}

	if find(value, *current.Left) {
		return true;
	}

	return find(value, *current.Right);
}

func (f *Finder) Find(value int, tree NodeTree) bool {
	return find(value, tree.Nodes[0]);
}