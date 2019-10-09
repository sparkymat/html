package meta

type HTTPEquiv string

const (
	Refresh      HTTPEquiv = "refresh"
	ContentType  HTTPEquiv = "content-type"
	DefaultStyle HTTPEquiv = "default-style"
)

func (et HTTPEquiv) String() string {
	return string(et)
}

type Info string

const (
	ApplicationName Info = "application-name"
	Author          Info = "author"
	Description     Info = "description"
	Generator       Info = "generator"
	Keywords        Info = "keywords"
	Viewport        Info = "viewport"
)

func (i Info) String() string {
	return string(i)
}
