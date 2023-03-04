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

func (n *node) getChild(seg string) (*node, bool) {
	if n.children == nil {
		return nil, false
	}
	child, ok := n.children[seg]
	return child, ok
}
