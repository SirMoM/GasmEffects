package gasm

import (
	"fmt"
	"strings"
	"syscall/js"

	e "github.com/SirMoM/go-wasm/effects"
)

type ImgData struct {
	Data        []byte `gasm:"data,clamped"`
	ColorSpace  string `gasm:"colorSpace"`
	Height      int    `gasm:"height"`
	PixelFormat string `gasm:"pixelFormat"`
	Width       int    `gasm:"width"`
}

type ManipulateImgArgs struct {
	functionName string
	imageData    ImgData
}

func validateCall(args []js.Value) (mArgs ManipulateImgArgs, err error) {

	// Validate arguments: expect (funcName: string, img: object)
	if len(args) != 2 {
		ERR(fmt.Sprintf("manipulateImg expects exactly 2 arguments: (functionName: string, img: object). Got %d", len(args)))
		return mArgs, fmt.Errorf("invalid argument count: expected 2, got %d", len(args))
	}
	if args[0].Type() != js.TypeString {
		ERR(fmt.Sprintf("invalid argument[0]: expected string function name, got %s (%v)", args[0].Type().String(), args[0]))
		return mArgs, fmt.Errorf("invalid argument type: expected string, got %s", args[0].Type().String())
	}
	mArgs.functionName = strings.ToLower(strings.TrimSpace(args[0].String()))

	if args[1].Type() != js.TypeObject {
		ERR(fmt.Sprintf("invalid argument[1]: expected  ImageData object, got %s (%v)", args[1].Type().String(), args[1]))
		return mArgs, fmt.Errorf("invalid argument type: expected ImageData object, got %s", args[0].Type().String())
	}
	parseJsObject(args[1], &mArgs.imageData)

	return mArgs, nil
}

func ManipulateImg(this js.Value, args []js.Value) any {
	Info("start manipulateImg")
	manipulateImgArgs, err := validateCall(args)
	if err != nil {
		return nil
	}

	var idx int
	switch manipulateImgArgs.functionName {
	case "greyscale", "grayscale":
		idx = e.GREYSCALE
	case "nearest neighbour":
		idx = e.NEAREST_NEIGHBOUR
	default:
		ERR(fmt.Sprintf("unknown manipulation function %q. Supported: greyscale", manipulateImgArgs.functionName))
		return js.Undefined()
	}

	manipulateImgArgs.imageData.Data = e.GetManipulationFunction(idx)(manipulateImgArgs.imageData.Data)

	return encodeJsObject(&manipulateImgArgs.imageData)
}
