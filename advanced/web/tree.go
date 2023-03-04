package web

type node struct {
	path       string
	children   map[string]*node
	handleFunc HandleFunc
}

func (n *node) getOrCreateChild(seg string) *node {
	if n.children == nil {
		n.children = make(map[string]*node)
	}
	child, ok := n.children[seg]
	if !ok {
		child = &node{path: seg}
		n.children[seg] = child
	}
	return child
}
