package main

import (
	"fmt"
	"math"
	"syscall/js"
)

type ImgData struct {
	Data        []byte `js:"data"`
	ColorSpace  string `js:"colorSpace"`
	Height      int    `js:"height"`
	PixelFormat string `js:"pixelFormat"`
	Width       int    `js:"width"`
}

func (i *ImgData) String() string {
	return fmt.Sprintf(`{
	"data": "%v",
	"ColorSpace": "%s",
	"Height": "%d",
	"PixelFormat": "%s",
	"Width": "%d",
}`, i.Data[:12], i.ColorSpace, i.Height, i.PixelFormat, i.Width)
}

type RGBA struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

type RgbaImage []RGBA

func (rgbaImage *RgbaImage) toJsImgData(imgData *ImgData) {
	// Convert the RGBA slice back to a tightly packed []byte (length = pixels*4)
	pixelCount := len(*rgbaImage)
	data := make([]byte, pixelCount*4)

	for i, rgba := range *rgbaImage {
		base := i * 4
		data[base] = rgba.R
		data[base+1] = rgba.G
		data[base+2] = rgba.B
		data[base+3] = rgba.A
	}
	imgData.Data = data
}

func rgbaFromImageDataData(imgData ImgData) RgbaImage {
	// TODO: Check for pixelFormat and ColorSpace
	totalBytes := len(imgData.Data)
	if totalBytes != imgData.Width*imgData.Height*4 {
		warn(fmt.Sprintf("rgbaFromImageDataData: data length %d does not match width*height*4 (%d)", totalBytes, imgData.Width*imgData.Height*4))
	}
	pixelCount := totalBytes / 4
	imgAsRGBA := make(RgbaImage, pixelCount)
	for p := 0; p < pixelCount; p++ {
		base := p * 4
		rgba := RGBA{
			R: imgData.Data[base],
			G: imgData.Data[base+1],
			B: imgData.Data[base+2],
			A: imgData.Data[base+3],
		}
		imgAsRGBA[p] = rgba
	}
	return imgAsRGBA
}

// GreyScale uses the Luminosity Method to turn the image to a gray scale
// grayscale = 0.3 * R + 0.59 * G + 0.11 * B
func GreyScale(oldImgData ImgData) ImgData {

	imgAsRGBA := rgbaFromImageDataData(oldImgData)
	info("converted to RGBA array")

 for i := range imgAsRGBA {
 	r := float64(imgAsRGBA[i].R)
 	g := float64(imgAsRGBA[i].G)
 	b := float64(imgAsRGBA[i].B)
 	gray := uint8(math.Floor(0.3*r + 0.59*g + 0.11*b))
 	imgAsRGBA[i].R = gray
 	imgAsRGBA[i].G = gray
 	imgAsRGBA[i].B = gray
 	// preserve alpha
 }

 info("converted to greyscale")
	greyscaledImageData := ImgData{
		Data:        nil,
		ColorSpace:  oldImgData.ColorSpace,
		PixelFormat: oldImgData.PixelFormat,
		Height:      oldImgData.Height,
		Width:       oldImgData.Width,
	}

	imgAsRGBA.toJsImgData(&greyscaledImageData)

	info("converted to JS struct")
	return greyscaledImageData
}

func manipulateImg(this js.Value, args []js.Value) any {

	info("start manipulateImg")

	if len(args) != 1 {
		ERR(fmt.Sprintf("manipulateImg requires 1 argument %v", args))
	}
	var strct ImgData
	parseJsObject(args[0], &strct)

	info("parsed obj")
	var res = GreyScale(strct)
	info(res.String())
	return encodeJsObject(&res)
}
