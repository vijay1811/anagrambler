package anagrambler

const (
	slabSize int = 100000
	poolSize int = 10
)

type Node struct {
	Words    []string
	Children map[rune]*Node
}

type slab struct {
	head  int
	nodes [slabSize]Node
}

type slabPool []*slab

type Trie struct {
	Root *Node
	pool slabPool
}

func NewNode() *Node {
	return &Node{
		Words:    make([]string, 0, 1),
		Children: make(map[rune]*Node),
	}
}

func newSlab() *slab {
	// The zero values of slab.head and slab.nodes are fine
	return &slab{}
}

func NewTrie() *Trie {
	trie := &Trie{
		Root: NewNode(),
		pool: make(slabPool, poolSize),
	}

	trie.pool = append(trie.pool, newSlab())

	return trie
}
