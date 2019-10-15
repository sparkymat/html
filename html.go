package html

import "fmt"

// HTML returns an <html> document
func HTML(head htmlHeadNode, body htmlBodyNode) html {
	return html{head: head, body: body}
}

type html struct {
	head htmlHeadNode
	body htmlBodyNode
}

func (h html) String() string {
	return fmt.Sprintf(`<!DOCTYPE html>
<html>
%s%s</html>
`, h.head.String(), h.body.String())
}
