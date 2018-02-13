package finder

import "testing"

func createTree() NodeTree {
	var tree NodeTree
	one := Node {
		1,
		&Node {},
		&Node {},
	}
	three := Node {
		3,
		&one,
		&Node {},		
	}
	ten := Node {
		10,
		&Node {},
		&Node {},
	};
	twelve := Node {
		12,
		&Node {},
		&Node {},
	};
	eleven := Node {
		11,
		&twelve,
		&Node {},
	};
	seven := Node {
		7,
		&ten,
		&eleven,
	}
	five := Node {
		5,
		&three,
		&seven,
	}
	tree.Nodes = []Node {
		five,
	}
	return tree;
}

func TestFindEleven(t *testing.T) {
	var f Finder
	tree := createTree();
	result := f.Find(11, tree);

	if result != true {
		t.Error("Expected true when looking for 11!")
	}
}

func TestFindTwo(t *testing.T) {
	var f Finder
	tree := createTree();
	result := f.Find(2, tree);

	if result != false {
		t.Error("Expected false when looking for 2!")
	}
}

func BenchmarkFindTwelve(b *testing.B) {
	var f Finder
	tree := createTree();

	for i := 0; i < b.N; i++ {
		f.Find(12, tree);
	}
}