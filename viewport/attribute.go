package viewport

import (
	"fmt"
	"strconv"

	"github.com/sparkymat/html/size"
)

type Attribute string

func (a Attribute) String() string {
	return string(a)
}

func HeightDeviceHeight() Attribute {
	return "height=device-height"
}

func Height(height size.Size) Attribute {
	return Attribute(fmt.Sprintf("height=%s", height.String()))
}

func WidthDeviceWidth() Attribute {
	return "width=device-width"
}

func Width(width size.Size) Attribute {
	return Attribute(fmt.Sprintf("width=%s", width.String()))
}

func InitialScale(scale float64) Attribute {
	return Attribute(fmt.Sprintf("initial-scale=%s", formatFloat(scale)))
}

func MinimumScale(scale float64) Attribute {
	return Attribute(fmt.Sprintf("minimum-scale=%s", formatFloat(scale)))
}

func MaximumScale(scale float64) Attribute {
	return Attribute(fmt.Sprintf("maximum-scale=%s", formatFloat(scale)))
}

func UserScalable(isScalable bool) Attribute {
	boolString := map[bool]string{
		false: "no",
		true:  "yes",
	}[isScalable]
	return Attribute(fmt.Sprintf("user-scalable=%s", boolString))
}

func formatFloat(v float64) string {
	return strconv.FormatFloat(v, 'f', -1, 64)
}
