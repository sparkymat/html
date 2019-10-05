package html

import "testing"

func TestHtml(t *testing.T) {
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
