package html

import "fmt"

// HTML returns an <html> document
func HTML(head HTMLHeadNode, body HTMLBodyNode) HTMLNode {
	return HTMLNode{head: head, body: body}
}

type HTMLNode struct {
	head HTMLHeadNode
	body HTMLBodyNode
}

func (h HTMLNode) String() string {
	return fmt.Sprintf(`<!DOCTYPE html>
<html>
%s%s</html>
`, h.head.String(), h.body.String())
}
