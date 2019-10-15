package html

import "fmt"

type htmlHeadNode struct {
	children []headNode
}

// Head returns a <head> node
func Head(children ...headNode) htmlHeadNode {
	return htmlHeadNode{children: children}
}

func (h htmlHeadNode) String() string {
	nodeString := ""
	for _, childNode := range h.children {
		nodeString = fmt.Sprintf("%s%s", nodeString, childNode.String())
	}
	return fmt.Sprintf(`<head>
%s</head>
`, nodeString)
}
