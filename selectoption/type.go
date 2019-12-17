package selectoption

type Type struct {
	Value string
	Label string
}

func New(value string, label string) Type {
	return Type{
		Value: value,
		Label: label,
	}
}
