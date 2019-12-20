package button

type Type string

const (
	Button Type = "button"
	Submit Type = "submit"
)

func (t Type) String() string {
	return string(t)
}
