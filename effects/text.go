package effects

import "github.com/SirMoM/go-wasm/shared"

func drawText(imageIn shared.ImgData) (manipulatedImageOut shared.ImgData) {
	// Validate that data length matches width*height*4 (RGBA)
	if len(imageIn.Data)/4 != imageIn.Width*imageIn.Height {
		shared.ERR("Image dimensions are inconsistent did not change image!")
		return imageIn
	}

	rgbaImage := RgbaFromBytes(imageIn.Data)
	width := imageIn.Width
	text := "Hello World"
	textPos := len(rgbaImage)/2 + (width / 2) - (len(text) / 2 * 8)
	image, err := drawString(textPos, width, text, rgbaImage)
	if err != nil {
		shared.ERR(err.Error())
		return shared.ImgData{}
	}

	imageIn.Data = image.toBytes()
	return imageIn
}

const asciiRamp = "@#%8&WM*oahkbdpqwm0=-:. "

func turnToAscii(imageIn shared.ImgData) (manipulatedImageOut shared.ImgData) {
	if len(imageIn.Data)/4 != imageIn.Width*imageIn.Height {
		shared.ERR("Image dimensions are inconsistent did not change image!")
		return imageIn
	}

	scaledGreyscale := Greyscale(bilinear(imageIn))
	rgbaImage := RgbaFromBytes(scaledGreyscale.Data)

	for x := 0; x < scaledGreyscale.Width/Char8Width; x++ {
		for y := 0; y < scaledGreyscale.Width/Char8Width; y++ {
			// THis is not right
			rgba := rgbaImage[x*scaledGreyscale.Width+y]
			charIndex := int(rgba.G) * (len(asciiRamp) - 1) / 255
			var err error
			// THis is not right
			rgbaImage, err = drawString(x*scaledGreyscale.Width, y*scaledGreyscale.Height, string(asciiRamp[charIndex]), rgbaImage)
			if err != nil {
				shared.ERR("Error drawing ascii ramp")
				shared.ERR(err.Error())
				return shared.ImgData{}
			}

		}

	}
	imageIn.Data = rgbaImage.toBytes()

	return imageIn
}
