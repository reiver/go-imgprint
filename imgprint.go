package imgprint

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"sourcecode.social/reiver/go-imgrow"
)

// ImgPrint draw text on an image.
//
// The simplest example of using ImgPrint is:
//
//	imgprint.ImgPrint(dstimg, x, y, "Hello world!")
//
// ... which would draw the text "Hello world!" onto the image 'dstimg' starting a (x,y).
//
// 'dstimg' is the image where the text will be draw.
// Whomever calls ImagePrint needs to create this image before hand.
//
// 'x', 'y' is where the text starts to be drawn.
//
// `value` is what will be drawn.
// It could just a Go string.
// Or imgprint.ColoredString() could be used to give a string a color.
// Or imgprint.StyledString() could be used to give a string a style.
func ImgPrint(dstimg draw.Image, x int, y int, value any) error {
	if nil == dstimg {
		return errNilDestinationImage
	}

	// The default style is set here.
	var foregroundColor color.Color = color.NRGBA{0xFF,0x1D,0xCE, 0xFF} // hot magenta
	var fontFace font.Face = basicfont.Face7x13

	// See if the caller provided any style.
	{
		styledstringer, casted := value.(StyledStringer)
		if casted {
			{
				c := styledstringer.ForegroundColor()
				if nil != c  {
					foregroundColor = c
				}
			}
			{
				f := styledstringer.FontFace()
				if nil != f {
					fontFace = f
				}
			}
		}
	}

	// Figure out how much to scale the font-face.
	var drawScalar int
	{
		var dstimgBounds image.Rectangle = dstimg.Bounds()

		var imgWidth = dstimgBounds.Max.X - dstimgBounds.Min.X

		const approximateNumberOfCharacterAcross = 48

		var widthOfCharacterM int = 7
		{
			fontFaceBounds, _, ok := fontFace.GlyphBounds('M')
			if ok {
				widthOfCharacterM = fontFaceBounds.Max.X.Ceil()
			}
		}

		drawScalar = imgWidth / (approximateNumberOfCharacterAcross * widthOfCharacterM)
	}

	var drawer *font.Drawer
	{
		point := fixed.Point26_6{fixed.I(x), fixed.I(y)}

		drawer = &font.Drawer{
			Dst:  imgrow.DrawImage(drawScalar, dstimg),
			Src:  image.NewUniform(foregroundColor),
			Face: fontFace,
			Dot:  point,
		}
	}

	var str string
	{
		switch casted := value.(type) {
		case interface{String()string}:
			str = casted.String()
		case string:
			str = casted
		default:
			str = fmt.Sprintf("%v", value)
		}
	}

	// Because of how we are scaling the font, we need to move the starting point of where we start drawing back to the original x,y.
	xAdjusted, yAdjusted := x / drawScalar, y / drawScalar

	return imgPrint(drawer, xAdjusted, yAdjusted, str)
}


func imgPrint(drawer *font.Drawer, x int, y int, str string) error {
	if nil == drawer {
		return errNilFontDrawer
	}

	drawer.DrawString(str)

	return nil
}
