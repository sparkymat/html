package html

// MetaCharset returns a <meta> tag with the specified text
func MetaCharset(charset string) headChildNode {
	return headChildNode{
		standardNode{
			name: "meta",
			attributes: map[string]string{
				"charset": charset,
			},
		},
	}
}
