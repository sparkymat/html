package html

import "fmt"

type headNode struct {
	name             string
	attributes       map[string]string
	attributeOrder   []string
	unsafeHTMLString string
}

func (hn headNode) String() string {
	return nodeAsString(hn)
}

// Title returns a <title> tag with the specified text
func Title(text string) headNode {
	return headNode{
		unsafeHTMLString: fmt.Sprintf("<title>%s</title>\n", text),
	}
}

func Base(href string, target string) headNode {
	return headNode{
		name: "base",
		attributes: map[string]string{
			"href":   href,
			"target": target,
		},
		attributeOrder: []string{"href", "target"},
	}
}

func Style(rules string) headNode {
	return headNode{
		unsafeHTMLString: fmt.Sprintf("<style>%s</style>\n", rules),
	}
}

func Link(href string) headNode {
	return headNode{
		name: "link",
		attributes: map[string]string{
			"rel":  "stylesheet",
			"href": href,
		},
		attributeOrder: []string{"rel", "href"},
	}
}

func ScriptInline(body string) headNode {
	return headNode{
		unsafeHTMLString: fmt.Sprintf("<script>%s</script>\n", body),
	}
}

func ScriptExternal(href string) headNode {
	return headNode{
		name: "script",
		attributes: map[string]string{
			"src": href,
		},
	}
}

func (hn headNode) getUnsafeHTMLString() string      { return hn.unsafeHTMLString }
func (hn headNode) getName() string                  { return hn.name }
func (hn headNode) getChildrenCount() int            { return 0 }
func (hn headNode) getChildAt(index int) node        { return nil }
func (hn headNode) getAttributes() map[string]string { return hn.attributes }
func (hn headNode) getAttributeOrder() []string      { return hn.attributeOrder }
