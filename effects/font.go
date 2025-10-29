package effects

import "github.com/SirMoM/go-wasm/shared"

// Char is bit representation or a font char
type Char8 [8]uint8

var A = Char8{
	0b00011000,
	0b00100100,
	0b01000010,
	0b01111110,
	0b01000010,
	0b01000010,
	0b01000010,
	0b00000000,
}

func (c Char8) getPixel(x int, y int) uint8 {
	if (x >= 8 || x < 0) || (y >= 8 || y < 0) {
		shared.ERR("Pixel outside of Char")
		return 255
	}
	if c[y]&(1<<(7-x)) != 0 {
		return 0
	}

	return 255

}

func drawText(imageIn shared.ImgData) (manipulatedImageOut shared.ImgData) {
	// Validate that data length matches width*height*4 (RGBA)
	if len(imageIn.Data)/4 != imageIn.Width*imageIn.Height {
		shared.ERR("Image dimensions are inconsistent did not change image!")
		return imageIn
	}

	clusterSize := 8
	rgbaImage := RgbaFromBytes(imageIn.Data)
	width := imageIn.Width
	height := imageIn.Height

	for y := 0; y < height; y += clusterSize {
		for x := 0; x < width; x += clusterSize {

			for yy := y; yy < clusterSize; yy++ {
				rowBase := yy * width
				for xx := x; xx < clusterSize; xx++ {
					i := rowBase + xx
					p := A.getPixel(xx, yy)
					rgbaImage[i].R = p
					rgbaImage[i].G = p
					rgbaImage[i].B = p
					if p == 0 {
						rgbaImage[i].A = 0
					} else {
						rgbaImage[i].A = 255
					}
				}
			}
		}
	}

	imageIn.Data = rgbaImage.toBytes()
	return imageIn
}
