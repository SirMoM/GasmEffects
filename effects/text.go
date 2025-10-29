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
