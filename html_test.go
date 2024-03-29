package html

import (
	"testing"

	"github.com/andreyvit/diff"
	"github.com/sparkymat/html/meta"
	"github.com/sparkymat/html/size"
	"github.com/sparkymat/html/viewport"
)

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
	hString := h.String()

	if hString != expectedString {
		t.Errorf("Expected: %s\nHTML(): %s\n", expectedString, hString)
	}
}

func TestHtmlHead(t *testing.T) {
	h := HTML(
		Head(
			Title("Hello World"),
			MetaCharset("utf-8"),
			MetaHTTPEquiv(meta.Refresh, "30"),
			MetaInfo(meta.Author, "Charlie"),
			MetaInfo(meta.Generator, "go"),
			MetaViewport(
				viewport.HeightDeviceHeight(),
				viewport.Height(size.Px(200)),
				viewport.WidthDeviceWidth(),
				viewport.Width(size.Px(100)),
				viewport.InitialScale(1.0),
				viewport.MinimumScale(1.0),
				viewport.MaximumScale(2.0),
				viewport.UserScalable(true),
			),
			Base("http://www.example.com/app/", "_blank"),
			Link("/style.css"),
			Style("h1 {color: red;}"),
			ScriptInline("console.log('Hello, world');"),
			ScriptExternal("/index.js"),
		),
		Body(),
	)

	hString := h.String()

	expectedString := `<!DOCTYPE html>
<html>
<head>
<title>Hello World</title>
<meta charset="utf-8" />
<meta http-equiv="refresh" content="30" />
<meta name="author" content="Charlie" />
<meta name="generator" content="go" />
<meta name="viewport" content="height=device-height, height=200px, width=device-width, width=100px, initial-scale=1, minimum-scale=1, maximum-scale=2, user-scalable=yes" />
<base href="http://www.example.com/app/" target="_blank" />
<link rel="stylesheet" href="/style.css" />
<style>h1 {color: red;}</style>
<script>console.log('Hello, world');</script>
<script src="/index.js" />
</head>
<body>
</body>
</html>
`
	if hString != expectedString {
		t.Errorf("%v", diff.LineDiff(expectedString, hString))
	}
}
