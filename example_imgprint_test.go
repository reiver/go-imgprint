package imgprint_test

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
//	"image/png"
//	"os"

	"github.com/reiver/go-imgprint"
)

func ExampleImgPrint_simple() {

	// Here we define some colors.
	//
	// We use these in 2 places in this example:
	//
	// Place №1:
	//
	// We use this in the palette we use when creating the image.
	//
	// (Note that we don't have to use a palette when creating an image.
	// But, for this example we do.)
	//
	// Place №2:
	//
	// We use it to give the image an initial background color.
	//
	// ...
	//
	// Note that we have a number of colors here that we do NOT use in this example.
	// We included them to make this a more realistic example.
	transparent := color.NRGBA{0x00,0x00,0x00, 0x00}
	black       := color.NRGBA{0x1B,0x2A,0x34, 0xFF}
	blue        := color.NRGBA{0x1E,0x5A,0xA8, 0xFF}
	gray        := color.NRGBA{0x8A,0x92,0x8D, 0xFF}
	magenta     := color.NRGBA{0xFF,0x1D,0xCE, 0xFF}
	red         := color.NRGBA{0xB4,0x00,0x00, 0xFF}
	yellow      := color.NRGBA{0xFA,0xC8,0x0A, 0xFF}
	white       := color.NRGBA{0xFF,0xFF,0xFF, 0xFF}

	// Here we create a palette (of colors).
	//
	// We use this palette when creating the image.
	//
	// Note that, in general, a lot of images do NOT have to have a palette.
	// But there are times when you might want your image to use a palette.
	//
	// For example, when an image only has a limited number of colors,
	// we can sometimes make the image file smaller (by using a palette).
	//
	// Also, some people use palettes to get a certain style.
	var palette []color.Color = []color.Color{
		transparent,
		black,
		blue,
		gray,
		magenta,
		red,
		yellow,
		white,
	}

	// In this example the image we created will have the following dimensions:
	//
	//	1080x566
	//
	// This gets used later when we create our image.
	var bounds = image.Rectangle{
		Min: image.Point{0,0},
		Max: image.Point{1080,566},
	}

	// Initially create the image.
	//
	// Note that image is just in the computer's memory at this point.
	// It is NOT stored in a .GIF, .PNG, .JPEG file, or any other type of image file.
	//
	// Turning an image into a file is an extra step.
	palimg := image.NewPaletted(bounds, palette)

	// Give the image a background color.
	var background image.Image = &image.Uniform{black}
	drawer := draw.Drawer(draw.Src)
	drawer.Draw(palimg, palimg.Bounds(), background, image.ZP)

	// Print the text on it.
	//
	// THIS IS THE IMPORTANT PART OF THIS EXAMPLE.
	x,y := 10, 50
	str := "Hello world!"
	err := imgprint.ImgPrint(palimg, x, y, str) // <---------
	if nil != err {
		fmt.Println("ERROR:", err)
		return
	}

	saveImage("example-hello-world.png", palimg)

	// Output:
	//
	// FILENAME: example-hello-world.png
}

func ExampleImgPrint_coloredString() {

	// Here we define some colors.
	//
	// We use these in 2 places in this example:
	//
	// Place №1:
	//
	// We use this in the palette we use when creating the image.
	//
	// (Note that we don't have to use a palette when creating an image.
	// But, for this example we do.)
	//
	// Place №2:
	//
	// We use it to give the image an initial background color.
	//
	// ...
	//
	// Note that we have a number of colors here that we do NOT use in this example.
	// We included them to make this a more realistic example.
	transparent := color.NRGBA{0x00,0x00,0x00, 0x00}
	black       := color.NRGBA{0x1B,0x2A,0x34, 0xFF}
	blue        := color.NRGBA{0x1E,0x5A,0xA8, 0xFF}
	gray        := color.NRGBA{0x8A,0x92,0x8D, 0xFF}
	magenta     := color.NRGBA{0xFF,0x1D,0xCE, 0xFF}
	red         := color.NRGBA{0xB4,0x00,0x00, 0xFF}
	yellow      := color.NRGBA{0xFA,0xC8,0x0A, 0xFF}
	white       := color.NRGBA{0xFF,0xFF,0xFF, 0xFF}

	// Here we create a palette (of colors).
	//
	// We use this palette when creating the image.
	//
	// Note that, in general, a lot of images do NOT have to have a palette.
	// But there are times when you might want your image to use a palette.
	//
	// For example, when an image only has a limited number of colors,
	// we can sometimes make the image file smaller (by using a palette).
	//
	// Also, some people use palettes to get a certain style.
	var palette []color.Color = []color.Color{
		transparent,
		black,
		blue,
		gray,
		magenta,
		red,
		yellow,
		white,
	}

	// In this example the image we created will have the following dimensions:
	//
	//	1080x566
	//
	// This gets used later when we create our image.
	var bounds = image.Rectangle{
		Min: image.Point{0,0},
		Max: image.Point{1080,566},
	}

	// Initially create the image.
	//
	// Note that image is just in the computer's memory at this point.
	// It is NOT stored in a .GIF, .PNG, .JPEG file, or any other type of image file.
	//
	// Turning an image into a file is an extra step.
	palimg := image.NewPaletted(bounds, palette)

	// Give the image a background color.
	var background image.Image = &image.Uniform{black}
	drawer := draw.Drawer(draw.Src)
	drawer.Draw(palimg, palimg.Bounds(), background, image.ZP)

	// Print the text on it.
	//
	// THIS IS THE IMPORTANT PART OF THIS EXAMPLE.
	x,y := 10, 50
	str := "Hello world!"
	err := imgprint.ImgPrint(palimg, x, y, imgprint.ColoredString(blue, str)) // <---------
	if nil != err {
		fmt.Println("ERROR:", err)
		return
	}

	saveImage("example-hello-world_colored.png", palimg)

	// Output:
	//
	// FILENAME: example-hello-world_colored.png
}

func saveImage(filename string, img image.Image) {
	fmt.Println("FILENAME:", filename)

/*
	file, err := os.Create(filename)
	if nil != err {
		fmt.Fprintf(os.Stderr, "ERROR: problem creating (empty) file %q: %s\n", filename, err)
		return
	}

	defer file.Close()

	{
		var pngEncoder png.Encoder = png.Encoder{
			CompressionLevel:png.BestCompression,
		}

		err := pngEncoder.Encode(file, img)
		if nil != err {
			fmt.Fprintf(os.Stderr, "ERROR: problem writing .PNG file %q: %s\n", filename, err)
			return
		}
	}
*/
}
