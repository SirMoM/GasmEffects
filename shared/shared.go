package shared

type ImgData struct {
	Data        []byte `gasm:"data,clamped"`
	ColorSpace  string `gasm:"colorSpace"`
	Height      int    `gasm:"height"`
	PixelFormat string `gasm:"pixelFormat"`
	Width       int    `gasm:"width"`
}
