package html

import "fmt"

type node interface {
	getUnsafeHTMLString() string
	getName() string
	getChildrenCount() int
	getChildAt(int) node
	getAttributes() map[string]string
	getAttributeOrder() []string
}

func nodeAsString(n node) string {
	if n.getUnsafeHTMLString() != "" {
		return n.getUnsafeHTMLString()
	}

	var closeTag string

	openTag := fmt.Sprintf("<%s", n.getName())
	hasCloseTag := true
	multiline := true

	if n.getChildrenCount() == 0 {
		hasCloseTag = false
		multiline = false
	}
	if n.getChildrenCount() == 1 {
		if n.getChildAt(0).getUnsafeHTMLString() != "" {
			multiline = false
		}
	}

	if len(n.getAttributes()) > 0 {
		if len(n.getAttributeOrder()) > 0 {
			for _, key := range n.getAttributeOrder() {
				value := n.getAttributes()[key]
				openTag = fmt.Sprintf("%s %s=\"%s\"", openTag, key, value)
			}
		} else {
			for key, value := range n.getAttributes() {
				openTag = fmt.Sprintf("%s %s=\"%s\"", openTag, key, value)
			}
		}
	}

	if multiline {
		closeTag = fmt.Sprintf("\n")
	}

	if hasCloseTag {
		openTag = fmt.Sprintf("%s>", openTag)
		if multiline {
			openTag = fmt.Sprintf("%s\n", openTag)
		}
		closeTag = fmt.Sprintf("%s</%s>", closeTag, n.getName())
	} else {
		closeTag = fmt.Sprintf("%s />", closeTag)
	}
	var body string
	for i := 0; i < n.getChildrenCount(); i++ {
		child := n.getChildAt(i)
		body = fmt.Sprintf("%s%s", body, nodeAsString(child))
	}
	return fmt.Sprintf("%s%s%s\n", openTag, body, closeTag)
}
