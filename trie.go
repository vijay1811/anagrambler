package anagrambler

const slabSize int = 100000

type Word struct {
	s *string
	next *Word
}

type Node struct {
	Words    *Word
	Children [26]*Node
}

type slab struct {
	head  int
	nodes [slabSize]Node
}

type slabPool []*slab

type Trie struct {
	Root *Node
	pool slabPool
	poolHead int
}

func NewTrie() *Trie {
	t := &Trie{
		Root: &Node{},
		pool: make(slabPool, 1),
	}

	t.pool[0] = &slab{}

	return t
}


func (t *Trie) NextNode() *Node {
	source := t.pool[t.poolHead]

	// The newest slab is full. We need to make a new slab.
	if source.head == slabSize - 1 {
		t.pool = append(t.pool, &slab{})
		t.poolHead += 1
		source = t.pool[t.poolHead]
	}

	n := &source.nodes[source.head]

	source.head += 1

	return n
}
