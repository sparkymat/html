package method

type Type string

const (
	Get  Type = "GET"
	Post Type = "POST"
)

func (m Type) String() string {
	return string(m)
}
