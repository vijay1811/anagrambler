package anagrambler

const slabSize int = 100000

type Node struct {
	Words    []string
	Children map[rune]*Node
}

type slab struct {
	head  int
	nodes []*Node
}

type slabPool []*slab

type Trie struct {
	Root *Node
	pool slabPool
	poolHead int
}

func NewNode() *Node {
	return &Node{
		Words:    make([]string, 0, 1),
		Children: make(map[rune]*Node),
	}
}

func newSlab() *slab {
	s := &slab{
		head: 0,
		nodes: make([]*Node, slabSize),
	}

	for i := 0; i < slabSize; i++ {
		s.nodes[i] = NewNode()
	}

	return s
}

func NewTrie() *Trie {
	t := &Trie{
		Root: NewNode(),
		pool: make(slabPool, 1),
	}

	t.pool[0] = newSlab()

	return t
}


func NextNode(t *Trie) *Node {
	source := t.pool[t.poolHead]

	// The newest slab is full. We need to make a new slab.
	if source.head == slabSize - 1 {
		t.pool = append(t.pool, newSlab())
		t.poolHead += 1
		source = t.pool[t.poolHead]
	}

	n := source.nodes[source.head]

	source.head += 1

	return n
}
