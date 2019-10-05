package html

type headNode struct {
}

// Head returns a <head> node
func Head() *headNode {
	return &headNode{}
}

func (h *headNode) String() string {
	return `<head>
</head>
`
}
