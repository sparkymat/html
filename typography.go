package html

func H1(text string) bodyNode     { return nodeWithText("h1", text) }
func H2(text string) bodyNode     { return nodeWithText("h2", text) }
func H3(text string) bodyNode     { return nodeWithText("h3", text) }
func H4(text string) bodyNode     { return nodeWithText("h4", text) }
func H5(text string) bodyNode     { return nodeWithText("h5", text) }
func H6(text string) bodyNode     { return nodeWithText("h6", text) }
func P(text string) bodyNode      { return nodeWithText("p", text) }
func Strong(text string) bodyNode { return nodeWithText("strong", text) }
func Em(text string) bodyNode     { return nodeWithText("em", text) }
func Code(text string) bodyNode   { return nodeWithText("code", text) }
