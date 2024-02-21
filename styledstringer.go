package imgprint

import (
	"fmt"

	"image/color"

	"golang.org/x/image/font"
)

// StyledStringer represents a text string that has a (single) style.
type StyledStringer interface {
	fmt.Stringer
	ForegroundColor() color.Color
	FontFace() font.Face
}
