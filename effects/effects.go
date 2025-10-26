package effects

import "math"

// RGBA represents a single pixel.
type RGBA struct {
	R, G, B, A uint8
}

// RgbaImage is a slice of RGBA pixels.
type RgbaImage []RGBA

// ToBytes converts the RgbaImage back to a flat byte slice.
func (rgbaImage *RgbaImage) ToBytes() []byte {
	pixelCount := len(*rgbaImage)
	data := make([]byte, pixelCount*4)

	for i, rgba := range *rgbaImage {
		base := i * 4
		data[base] = rgba.R
		data[base+1] = rgba.G
		data[base+2] = rgba.B
		data[base+3] = rgba.A
	}
	return data
}

// RgbaFromBytes converts a flat byte slice into an RgbaImage.
func RgbaFromBytes(data []byte) RgbaImage {
	pixelCount := len(data) / 4
	imgAsRGBA := make(RgbaImage, pixelCount)
	for p := 0; p < pixelCount; p++ {
		base := p * 4
		rgba := RGBA{
			R: data[base],
			G: data[base+1],
			B: data[base+2],
			A: data[base+3],
		}
		imgAsRGBA[p] = rgba
	}
	return imgAsRGBA
}

// GreyScale uses the Luminosity Method to turn an image to a gray scale.
// It modifies the RgbaImage in place.
// grayscale = 0.3 * R + 0.59 * G + 0.11 * B
func (rgbaImage *RgbaImage) GreyScale() {
	for i := range *rgbaImage {
		r := float64((*rgbaImage)[i].R)
		g := float64((*rgbaImage)[i].G)
		b := float64((*rgbaImage)[i].B)
		gray := uint8(math.Floor(0.3*r + 0.59*g + 0.11*b))
		(*rgbaImage)[i].R = gray
		(*rgbaImage)[i].G = gray
		(*rgbaImage)[i].B = gray
	}
}
