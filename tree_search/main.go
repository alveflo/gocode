package main

import (
	"fmt"
	"./finder"
)

func createTree() finder.NodeTree {
	var tree finder.NodeTree
	one := finder.Node {
		1,
		&finder.Node {},
		&finder.Node {},
	}
	three := finder.Node {
		3,
		&one,
		&finder.Node {},		
	}
	ten := finder.Node {
		10,
		&finder.Node {},
		&finder.Node {},
	};
	twelve := finder.Node {
		12,
		&finder.Node {},
		&finder.Node {},
	};
	eleven := finder.Node {
		11,
		&twelve,
		&finder.Node {},
	};
	seven := finder.Node {
		7,
		&ten,
		&eleven,
	}
	five := finder.Node {
		5,
		&three,
		&seven,
	}
	tree.Nodes = []finder.Node {
		five,
	}
	return tree;
}

func main() {
	var f finder.Finder
	tree := createTree();
	fmt.Println("Looking for eleven (11):")
	result := f.Find(11, tree);
	fmt.Println(result);
}