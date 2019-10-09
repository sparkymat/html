package html

import "fmt"

type headNode struct {
	children []headChildNode
}

// Head returns a <head> node
func Head(children ...headChildNode) headNode {
	return headNode{children: children}
}

func (h headNode) String() string {
	nodeString := ""
	for _, childNode := range h.children {
		nodeString = fmt.Sprintf("%s%s", nodeString, childNode.String())
	}
	return fmt.Sprintf(`<head>
%s</head>
`, nodeString)
}
