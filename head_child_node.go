package html

import "fmt"

type headChildNode struct {
	name             string
	attributes       map[string]string
	attributeOrder   []string
	unsafeHTMLString string
}

func (hn headChildNode) String() string {
	return nodeAsString(hn)
}

// Title returns a <title> tag with the specified text
func Title(text string) headChildNode {
	return headChildNode{
		unsafeHTMLString: fmt.Sprintf("<title>%s</title>\n", text),
	}
}

func Base(href string, target string) headChildNode {
	return headChildNode{
		name: "base",
		attributes: map[string]string{
			"href":   href,
			"target": target,
		},
		attributeOrder: []string{"href", "target"},
	}
}

func Style(rules string) headChildNode {
	return headChildNode{
		unsafeHTMLString: fmt.Sprintf("<style>%s</style>\n", rules),
	}
}

func Link(href string) headChildNode {
	return headChildNode{
		name: "link",
		attributes: map[string]string{
			"rel":  "stylesheet",
			"href": href,
		},
		attributeOrder: []string{"rel", "href"},
	}
}

func ScriptInline(body string) headChildNode {
	return headChildNode{
		unsafeHTMLString: fmt.Sprintf("<script>%s</script>\n", body),
	}
}

func ScriptExternal(href string) headChildNode {
	return headChildNode{
		name: "script",
		attributes: map[string]string{
			"src": href,
		},
	}
}

func (hn headChildNode) getUnsafeHTMLString() string      { return hn.unsafeHTMLString }
func (hn headChildNode) getName() string                  { return hn.name }
func (hn headChildNode) getChildrenCount() int            { return 0 }
func (hn headChildNode) getChildAt(index int) node        { return nil }
func (hn headChildNode) getAttributes() map[string]string { return hn.attributes }
func (hn headChildNode) getAttributeOrder() []string      { return hn.attributeOrder }
