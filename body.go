package html

type bodyNode struct {
}

// Body returns a <body> node
func Body() *bodyNode {
	return &bodyNode{}
}

func (b *bodyNode) String() string {
	return `<body>
</body>
`
}
