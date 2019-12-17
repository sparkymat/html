package html

import "fmt"

type HeadNode struct {
	name             string
	attributes       map[string]string
	attributeOrder   []string
	unsafeHTMLString string
	alwaysCloseTag   bool
}

func (hn HeadNode) String() string {
	return nodeAsString(hn)
}

// Title returns a <title> tag with the specified text
func Title(text string) HeadNode {
	return HeadNode{
		unsafeHTMLString: fmt.Sprintf("<title>%s</title>\n", text),
	}
}

func Base(href string, target string) HeadNode {
	return HeadNode{
		name: "base",
		attributes: map[string]string{
			"href":   href,
			"target": target,
		},
		attributeOrder: []string{"href", "target"},
	}
}

func Style(rules string) HeadNode {
	return HeadNode{
		unsafeHTMLString: fmt.Sprintf("<style>%s</style>\n", rules),
	}
}

func Link(href string) HeadNode {
	return HeadNode{
		name: "link",
		attributes: map[string]string{
			"rel":  "stylesheet",
			"href": href,
		},
		attributeOrder: []string{"rel", "href"},
	}
}

func ScriptInline(body string) HeadNode {
	return HeadNode{
		unsafeHTMLString: fmt.Sprintf("<script>%s</script>\n", body),
	}
}

func ScriptExternal(href string) HeadNode {
	return HeadNode{
		name: "script",
		attributes: map[string]string{
			"src": href,
		},
		alwaysCloseTag: true,
	}
}

func (hn HeadNode) getUnsafeHTMLString() string      { return hn.unsafeHTMLString }
func (hn HeadNode) getName() string                  { return hn.name }
func (hn HeadNode) getChildrenCount() int            { return 0 }
func (hn HeadNode) getChildAt(index int) node        { return nil }
func (hn HeadNode) getAttributes() map[string]string { return hn.attributes }
func (hn HeadNode) getAttributeOrder() []string      { return hn.attributeOrder }
func (hn HeadNode) getAlwaysCloseTag() bool          { return hn.alwaysCloseTag }
