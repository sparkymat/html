package html

import (
	"testing"

	"github.com/andreyvit/diff"
)

func TestDiv(t *testing.T) {
	h := HTML(
		Head(),
		Body(
			Div(
				Text("oh well"),
			),
		),
	)

	expectedString := `<!DOCTYPE html>
<html>
<head>
</head>
<body>
<div>oh well</div>
</body>
</html>
`
	hString := h.String()

	if hString != expectedString {
		t.Errorf("%v", diff.LineDiff(expectedString, hString))
	}
}
