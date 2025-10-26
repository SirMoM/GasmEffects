package gasm

import (
	"fmt"
	"syscall/js"

	"github.com/SirMoM/go-wasm/effects"
)

type ImgData struct {
	Data        []byte `gasm:"data,clamped"`
	ColorSpace  string `gasm:"colorSpace"`
	Height      int    `gasm:"height"`
	PixelFormat string `gasm:"pixelFormat"`
	Width       int    `gasm:"width"`
}

func ManipulateImg(this js.Value, args []js.Value) any {
	Info("start manipulateImg")

	if len(args) != 1 {
		ERR(fmt.Sprintf("manipulateImg requires 1 argument %v", args))
	}
	if args[0].Type() != js.TypeObject {
		ERR(fmt.Sprintf("manipulateImg requires an object %v", args))
	}
	var strct ImgData
	parseJsObject(args[0], &strct)

	// Convert to RgbaImage, process, and convert back.
	imgAsRGBA := effects.RgbaFromBytes(strct.Data)
	imgAsRGBA.GreyScale()
	strct.Data = imgAsRGBA.ToBytes()

	Info("converted to greyscale")
	return encodeJsObject(&strct)
}
