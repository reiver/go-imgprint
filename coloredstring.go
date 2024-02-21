package imgprint

import (
	"image/color"

	"golang.org/x/image/font"
)

// ColoredString is used, with imgprint.ImgPrint(), to give a text string a color.
//
// For example:
//
//	var red color.NRGBA{0xB4,0x00,0x00, 0xFF}
//	
//	imgprint.ImgPrint(
//		dstimg,
//		x, y,
//		imgprint.ColoredImage(red, "Hello world!"),
//	)
func ColoredString(c color.Color, str string) StyledStringer {
	return internalColoredString{
		str:str,
		foregroundColor: c,
	}
}

type internalColoredString struct {
	str string
	foregroundColor color.Color
}

func (internalColoredString) FontFace() font.Face {
	return nil
}

func (receiver internalColoredString) ForegroundColor() color.Color {
	return receiver.foregroundColor
}

func (receiver internalColoredString) String() string {
	return receiver.str
}

