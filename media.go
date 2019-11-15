package html

func Video(showControls bool) BodyNode {
	attributes := map[string]string{}
	attributesOrder := []string{}

	if showControls {
		attributes["controls"] = "controls"
		attributesOrder = append(attributesOrder, "controls")
	}

	return BodyNode{
		name:           "video",
		attributes:     attributes,
		attributeOrder: attributesOrder,
	}
}

func Source(src string, mimeType string) BodyNode {
	return BodyNode{
		name: "source",
		attributes: map[string]string{
			"src":  src,
			"type": mimeType,
		},
		attributeOrder: []string{
			"src",
			"type",
		},
	}
}
