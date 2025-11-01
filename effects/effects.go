package effects

import (
	"math"
	"unsafe"

	"github.com/SirMoM/go-wasm/shared"
)

// RGBA represents a single pixel.
type RGBA struct {
	R, G, B, A uint8
}

// RgbaImage is a slice of RGBA pixels.
type RgbaImage []RGBA

const (
	GREYSCALE = iota
	NEAREST_NEIGHBOUR
	BILINEAR
	TEXT
	ASCII
	END
)

type ManipulationFunction func(imageIn shared.ImgData) (manipulatedImageOut shared.ImgData)
type ManipulationFunctions []ManipulationFunction

var Functions ManipulationFunctions = ManipulationFunctions{
	Greyscale,
	nearestNeighbour,
	bilinear,
	drawText,
	turnToAscii,
}

func GetManipulationFunction(funIdx int) ManipulationFunction {
	return Functions[funIdx]
}

// ToBytes converts the RgbaImage back to a flat byte slice.
func (rgbaImage *RgbaImage) toBytes() []byte {
	if len(*rgbaImage) == 0 {
		return []byte{}
	}

	// Use unsafe to reinterpret the RGBA slice as a byte slice
	// This avoids copying and is much more efficient
	return unsafe.Slice((*byte)(unsafe.Pointer(&(*rgbaImage)[0])), len(*rgbaImage)*4)

}

// RgbaFromBytes converts a flat byte slice into an RgbaImage.
func RgbaFromBytes(data []byte) RgbaImage {
	if len(data)%4 != 0 {
		shared.ERR("data length must be a multiple of 4")
		return RgbaImage{}
	}

	pixelCount := len(data) / 4

	return RgbaImage(unsafe.Slice((*RGBA)(unsafe.Pointer(&data[0])), pixelCount))
}

// Greyscale uses the Luminosity Method to turn an image to a gray scale.
// grayscale = 0.3 * R + 0.59 * G + 0.11 * B
func Greyscale(imageIn shared.ImgData) (manipulatedImageOut shared.ImgData) {
	rgbaImage := RgbaFromBytes(imageIn.Data)
	for i := range rgbaImage {
		r := float64(rgbaImage[i].R)
		g := float64(rgbaImage[i].G)
		b := float64(rgbaImage[i].B)
		gray := uint8(math.Floor(0.3*r + 0.59*g + 0.11*b))
		rgbaImage[i].R = gray
		rgbaImage[i].G = gray
		rgbaImage[i].B = gray
	}
	imageIn.Data = rgbaImage.toBytes()
	return imageIn
}

// nearestNeighbour applies a pixelation effect by grouping pixels into clusters
// and assigning each cluster the color of a middle pixel.
func nearestNeighbour(imageIn shared.ImgData) (manipulatedImageOut shared.ImgData) {
	// Validate that data length matches width*height*4 (RGBA)
	if len(imageIn.Data)/4 != imageIn.Width*imageIn.Height {
		shared.ERR("Image dimensions are inconsistent did not change image!")
		return imageIn
	}

	clusterSize := 16
	rgbaImage := RgbaFromBytes(imageIn.Data)
	width := imageIn.Width
	height := imageIn.Height

	for y := 0; y < height; y += clusterSize {
		for x := 0; x < width; x += clusterSize {
			idx := y*width + x
			r := rgbaImage[idx].R
			g := rgbaImage[idx].G
			b := rgbaImage[idx].B

			yMax := y + clusterSize
			if yMax > height {
				yMax = height
			}
			xMax := x + clusterSize
			if xMax > width {
				xMax = width
			}

			for yy := y; yy < yMax; yy++ {
				rowBase := yy * width
				for xx := x; xx < xMax; xx++ {
					i := rowBase + xx
					rgbaImage[i].R = r
					rgbaImage[i].G = g
					rgbaImage[i].B = b
				}
			}
		}
	}

	imageIn.Data = rgbaImage.toBytes()
	return imageIn
}

// bilinear applies a pixelation effect by grouping pixels into clusters
// and assigning each cluster the median color of all pixels.
func bilinear(imageIn shared.ImgData) (manipulatedImageOut shared.ImgData) {
	// Validate that data length matches width*height*4 (RGBA)
	if len(imageIn.Data)/4 != imageIn.Width*imageIn.Height {
		shared.ERR("Image dimensions are inconsistent did not change image!")
		return imageIn
	}

	clusterSize := 8
	rgbaImage := RgbaFromBytes(imageIn.Data)
	width := imageIn.Width
	height := imageIn.Height

	// Reusable histograms for each channel (0..255)
	var histR, histG, histB [256]int

	for y := 0; y < height; y += clusterSize {
		for x := 0; x < width; x += clusterSize {
			// Determine cluster boundaries (handle edges)
			yMax := y + clusterSize
			if yMax > height {
				yMax = height
			}
			xMax := x + clusterSize
			if xMax > width {
				xMax = width
			}

			// Reset histograms
			for i := 0; i < 256; i++ {
				histR[i] = 0
				histG[i] = 0
				histB[i] = 0
			}

			// Build histograms for the block
			n := 0
			for yy := y; yy < yMax; yy++ {
				rowBase := yy * width
				for xx := x; xx < xMax; xx++ {
					p := rgbaImage[rowBase+xx]
					histR[p.R]++
					histG[p.G]++
					histB[p.B]++
					n++
				}
			}

			// Find median for each channel (lower median for even n)
			target := (n - 1) / 2
			findMedian := func(hist *[256]int, target int) uint8 {
				sum := 0
				for v := 0; v < 256; v++ {
					sum += hist[v]
					if sum > target {
						return uint8(v)
					}
				}
				return 255
			}
			mr := findMedian(&histR, target)
			mg := findMedian(&histG, target)
			mb := findMedian(&histB, target)

			// Paint the block with the median color (preserving alpha)
			for yy := y; yy < yMax; yy++ {
				rowBase := yy * width
				for xx := x; xx < xMax; xx++ {
					i := rowBase + xx
					rgbaImage[i].R = mr
					rgbaImage[i].G = mg
					rgbaImage[i].B = mb
				}
			}
		}
	}

	imageIn.Data = rgbaImage.toBytes()
	return imageIn
}
