package chash

type Consistent interface {
	AddNode(node string)
	RemoveNode(node string)
	GetNode(key string) string
}

type HashFunc func(data []byte) uint32
