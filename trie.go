package anagrambler

const slabSize int = 100000

type Word struct {
	s *string
	next *Word
}

type Node struct {
	Words    *Word
	a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w, x, y, z *Node
}

func (node *Node) getChild (letter rune) *Node {
	switch letter {
	case 'a':
		return node.a
	case 'b':
		return node.b
	case 'c':
		return node.c
	case 'd':
		return node.d
	case 'e':
		return node.e
	case 'f':
		return node.f
	case 'g':
		return node.g
	case 'h':
		return node.h
	case 'i':
		return node.i
	case 'j':
		return node.j
	case 'k':
		return node.k
	case 'l':
		return node.l
	case 'm':
		return node.m
	case 'n':
		return node.n
	case 'o':
		return node.o
	case 'p':
		return node.p
	case 'q':
		return node.q
	case 'r':
		return node.r
	case 's':
		return node.s
	case 't':
		return node.t
	case 'u':
		return node.u
	case 'v':
		return node.v
	case 'w':
		return node.w
	case 'x':
		return node.x
	case 'y':
		return node.y
	case 'z':
		return node.z
	default:
		panic(letter)
	}
}

func (node *Node) setChild (letter rune, child *Node) {
	switch letter {
	case 'a':
		node.a = child
	case 'b':
		node.b = child
	case 'c':
		node.c = child
	case 'd':
		node.d = child
	case 'e':
		node.e = child
	case 'f':
		node.f = child
	case 'g':
		node.g = child
	case 'h':
		node.h = child
	case 'i':
		node.i = child
	case 'j':
		node.j = child
	case 'k':
		node.k = child
	case 'l':
		node.l = child
	case 'm':
		node.m = child
	case 'n':
		node.n = child
	case 'o':
		node.o = child
	case 'p':
		node.p = child
	case 'q':
		node.q = child
	case 'r':
		node.r = child
	case 's':
		node.s = child
	case 't':
		node.t = child
	case 'u':
		node.u = child
	case 'v':
		node.v = child
	case 'w':
		node.w = child
	case 'x':
		node.x = child
	case 'y':
		node.y = child
	case 'z':
		node.z = child
	}
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


func NextNode(t *Trie) *Node {
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
