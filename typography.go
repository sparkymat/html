package html

func H1(text string) BodyNode     { return nodeWithText("h1", text) }
func H2(text string) BodyNode     { return nodeWithText("h2", text) }
func H3(text string) BodyNode     { return nodeWithText("h3", text) }
func H4(text string) BodyNode     { return nodeWithText("h4", text) }
func H5(text string) BodyNode     { return nodeWithText("h5", text) }
func H6(text string) BodyNode     { return nodeWithText("h6", text) }
func P(text string) BodyNode      { return nodeWithText("p", text) }
func Strong(text string) BodyNode { return nodeWithText("strong", text) }
func Em(text string) BodyNode     { return nodeWithText("em", text) }
func Code(text string) BodyNode   { return nodeWithText("code", text) }
