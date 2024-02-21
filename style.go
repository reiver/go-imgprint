package imgprint

import (
	"image/color"

	"golang.org/x/image/font"
)

// Style represents a single style, that is applied to a text string.
//
// A style includes:
// foreground-color,
// font-face.
//
// Style is uses with imgprint.StyledString(), to give a string style.
// Which in turn is passed to imgprint.ImgPrint().
type Style struct {
	ForegroundColor color.Color
	FontFace font.Face
}
