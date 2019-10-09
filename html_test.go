package html

import "testing"

func TestHtmlBasic(t *testing.T) {
	h := HTML(
		Head(),
		Body(),
	)

	expectedString := `<!DOCTYPE html>
<html>
<head>
</head>
<body>
</body>
</html>
`
	if h.String() != expectedString {
		t.Errorf("Expected: %s\nHTML(): %s\n", expectedString, h.String())
	}
}

func TestHtmlComplete(t *testing.T) {
	h := HTML(
		Head(
			Title("Hello World"),
			MetaCharset("utf-8"),
		),
		Body(),
	)

	expectedString := `<!DOCTYPE html>
<html>
<head>
<title>Hello World</title>
<meta charset="utf-8" />
</head>
<body>
</body>
</html>
`
	if h.String() != expectedString {
		t.Errorf("Expected: %s\nHTML(): %s\n", expectedString, h.String())
	}
}
