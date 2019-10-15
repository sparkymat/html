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
			Div().Class("card").Children(
				Div().Class("card-header").Children(
					H1().Class("card-title").Children(Text("Hello")),
				),
				Div().Class("card-body").Children(
					P(`
This was a triumph!
I am making a note here!
`),
				),
				Div().Class("card-footer").Children(
					A().Href("/foo").Class("btn-primary").Children(Text("Do")),
					A().Href("/cancel").Class("btn-secondary").Children(Text("Don't")),
				),
			),
		),
	)

	expectedString := `<!DOCTYPE html>
<html>
<head>
</head>
<body>
<div>oh well</div>
<div class="card">
<div class="card-header">
<h1 class="card-title">Hello</h1>
</div>
<div class="card-body">
<p>
This was a triumph!
I am making a note here!
</p>
</div>
<div class="card-footer">
<a href="/foo" class="btn-primary">Do</a>
<a href="/cancel" class="btn-secondary">Don't</a>
</div>
</div>
</body>
</html>
`
	hString := h.String()

	if hString != expectedString {
		t.Errorf("%v", diff.LineDiff(expectedString, hString))
	}
}
