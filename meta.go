package html

import (
	"strings"

	"github.com/sparkymat/html/meta"
	"github.com/sparkymat/html/viewport"
)

// MetaCharset returns a <meta> tag with the specified text
func MetaCharset(charset string) headNode {
	return headNode{
		name: "meta",
		attributes: map[string]string{
			"charset": charset,
		},
		attributeOrder: []string{"charset"},
	}
}

func MetaHTTPEquiv(header meta.HTTPEquiv, value string) headNode {
	return headNode{
		name: "meta",
		attributes: map[string]string{
			"http-equiv": header.String(),
			"content":    value,
		},
		attributeOrder: []string{"http-equiv", "content"},
	}
}

func MetaInfo(name meta.Info, value string) headNode {
	return headNode{
		name: "meta",
		attributes: map[string]string{
			"name":    name.String(),
			"content": value,
		},
		attributeOrder: []string{"name", "content"},
	}
}

func MetaViewport(attributes ...viewport.Attribute) headNode {
	attributesStrings := []string{}
	for _, attribute := range attributes {
		attributesStrings = append(attributesStrings, attribute.String())
	}

	return MetaInfo(meta.Viewport, strings.Join(attributesStrings, ", "))
}
