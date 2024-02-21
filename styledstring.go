package imgprint

import (
	"image/color"

	"golang.org/x/image/font"
)

// StyledString is used, with imgprint.ImgPrint(), to give a text string a (single) style.
//
// For example:
//
//	var blue color.NRGBA{0x1E,0x5A,0xA8, 0xFF}
//	var fontface font.Face = basicfont.Face7x13
//	
//	var style imgprint.Style = imgprint.Style{
//		FontFace:
//		ForegroundColor: blue,
//	}
//	
//	imgprint.ImgPrint(
//		dstimg,
//		x, y,
//		imgprint.StyledString(style, "Hello world!"),
//	)
func StyledString(style Style, str string) StyledStringer {
	return internalStyledString{
		str:str,
		style:style,
	}
}

type internalStyledString struct {
	str string
	style Style
}

func (receiver internalStyledString) FontFace() font.Face {
	return receiver.style.FontFace
}

func (receiver internalStyledString) ForegroundColor() color.Color {
	return receiver.style.ForegroundColor
}

func (receiver internalStyledString) String() string {
	return receiver.str
}
