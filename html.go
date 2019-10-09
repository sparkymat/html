package html

import "fmt"

// HTML returns an <html> document
func HTML(head headNode, body bodyNode) htmlNode {
	return htmlNode{head: head, body: body}
}

type htmlNode struct {
	head headNode
	body bodyNode
}

func (h htmlNode) String() string {
	return fmt.Sprintf(`<!DOCTYPE html>
<html>
%s%s</html>
`, h.head.String(), h.body.String())
}
