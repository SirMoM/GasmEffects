package effects

import (
	"fmt"
	"strings"

	"github.com/SirMoM/go-wasm/shared"
)

const Char8Width = 8

// Char8 is a bitmap representation or a font char
type Char8 [Char8Width]uint8

func (c Char8) String() string {
	sB := strings.Builder{}
	sB.WriteString(fmt.Sprintf("Char8:\n"))
	for _, u := range c {
		sB.WriteString(fmt.Sprintf("%b\n", u))
	}
	return sB.String()
}

// StringToChar8 converts a string to a Char8 array to be drawn on a RgbaImage
func StringToChar8(s string) ([]*Char8, error) {
	c := make([]*Char8, len(s))
	shared.Warn(fmt.Sprintf("Converting string to []Char8 %s", strings.ToUpper(s)))
	for i, r := range strings.ToUpper(s) {

		char8, ok := FontMap[r]
		if !ok {
			return nil, fmt.Errorf("string contains invalid char can not convert to Char8 %c", r)
		} else {
			//shared.Warn(fmt.Sprintf("%c -> %s", r, char8))
			c[i] = &char8
		}
	}
	return c, nil
}

func (c Char8) getPixel(x int, y int) uint8 {
	if (x >= Char8Width || x < 0) || (y >= Char8Width || y < 0) {
		shared.ERR("Pixel outside of Char")
		return 255
	}
	if c[y]&(1<<(7-x)) != 0 {
		return 0
	}

	return 255

}

func drawChar8(startIdx int, width int, c Char8, buf RgbaImage) RgbaImage {
	// Early exit for whitespace this might be a problem if we want to draw backgrounds
	if c == WHITE_SPACE {
		return buf
	}

	for y := 0; y < Char8Width; y++ {
		for x := 0; x < Char8Width; x++ {
			p := c.getPixel(x, y)
			if p == 0 {
				// TODO This does not work well with overflows...
				idx := startIdx + (y*width + x)
				if idx >= len(buf) {
					shared.ERR("Pixel outside of image")
					continue
				}
				buf[idx].R = 0
				buf[idx].G = 0
				buf[idx].B = 0
			}
		}
	}
	return buf
}

func drawString(startIdx int, width int, text string, buf RgbaImage) (RgbaImage, error) {
	bitmaps, err := StringToChar8(text)
	if err != nil {
		return nil, err
	}

	for i, char8 := range bitmaps {
		drawChar8(startIdx+(Char8Width*i), width, *char8, buf)
	}

	return buf, nil
}

var FontMap = map[rune]Char8{
	' ': WHITE_SPACE,
	'A': A,
	'B': B,
	'C': C,
	'D': D,
	'E': E,
	'F': F,
	'G': G,
	'H': H,
	'I': I,
	'J': J,
	'K': K,
	'L': L,
	'M': M,
	'N': N,
	'O': O,
	'P': P,
	'Q': Q,
	'R': R,
	'S': S,
	'T': T,
	'U': U,
	'V': V,
	'W': W,
	'X': X,
	'Y': Y,
	'Z': Z,
	'0': ZERO,
	'1': ONE,
	'2': TWO,
	'3': THREE,
	'4': FOUR,
	'5': FIVE,
	'6': SIX,
	'7': SEVEN,
	'8': EIGHT,
	'9': NINE,
	'.': DOT,
	':': COLON,
	'-': DASH,
	'=': EQUAL,
	'+': PLUS,
	'*': STAR,
	'#': HASH,
	'%': PERCENT,
	'@': AT,
	'&': AMPERSAND,
}

var WHITE_SPACE = Char8{
	0b0,
}
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

var B = Char8{
	0b01111100,
	0b01000010,
	0b01000010,
	0b01111100,
	0b01000010,
	0b01000010,
	0b01111100,
	0b00000000,
}

var C = Char8{
	0b00111110,
	0b01000000,
	0b01000000,
	0b01000000,
	0b01000000,
	0b01000000,
	0b00111110,
	0b00000000,
}

var D = Char8{
	0b01111100,
	0b01000010,
	0b01000010,
	0b01000010,
	0b01000010,
	0b01000010,
	0b01111100,
	0b00000000,
}

var E = Char8{
	0b01111110,
	0b01000000,
	0b01000000,
	0b01111100,
	0b01000000,
	0b01000000,
	0b01111110,
	0b00000000,
}

var F = Char8{
	0b01111110,
	0b01000000,
	0b01000000,
	0b01111100,
	0b01000000,
	0b01000000,
	0b01000000,
	0b00000000,
}

var G = Char8{
	0b00111110,
	0b01000000,
	0b01000000,
	0b01001110,
	0b01000010,
	0b01000010,
	0b00111110,
	0b00000000,
}

var H = Char8{
	0b01000010,
	0b01000010,
	0b01000010,
	0b01111110,
	0b01000010,
	0b01000010,
	0b01000010,
	0b00000000,
}

var I = Char8{
	0b01111110,
	0b00011000,
	0b00011000,
	0b00011000,
	0b00011000,
	0b00011000,
	0b01111110,
	0b00000000,
}

var J = Char8{
	0b00000110,
	0b00000100,
	0b00000100,
	0b00000100,
	0b01000100,
	0b01000100,
	0b00111000,
	0b00000000,
}

var K = Char8{
	0b01000100,
	0b01001000,
	0b01010000,
	0b01100000,
	0b01010000,
	0b01001000,
	0b01000100,
	0b00000000,
}

var L = Char8{
	0b01000000,
	0b01000000,
	0b01000000,
	0b01000000,
	0b01000000,
	0b01000000,
	0b01111110,
	0b00000000,
}

var M = Char8{
	0b01000010,
	0b01100110,
	0b01011010,
	0b01000010,
	0b01000010,
	0b01000010,
	0b01000010,
	0b00000000,
}

var N = Char8{
	0b01000010,
	0b01100010,
	0b01010010,
	0b01001010,
	0b01000110,
	0b01000010,
	0b01000010,
	0b00000000,
}

var O = Char8{
	0b00111100,
	0b01000010,
	0b01000010,
	0b01000010,
	0b01000010,
	0b01000010,
	0b00111100,
	0b00000000,
}

var P = Char8{
	0b01111100,
	0b01000010,
	0b01000010,
	0b01111100,
	0b01000000,
	0b01000000,
	0b01000000,
	0b00000000,
}

var Q = Char8{
	0b00111100,
	0b01000010,
	0b01000010,
	0b01000010,
	0b01001010,
	0b01000100,
	0b00111110,
	0b00000000,
}

var R = Char8{
	0b01111100,
	0b01000010,
	0b01000010,
	0b01111100,
	0b01010000,
	0b01001000,
	0b01000100,
	0b00000000,
}

var S = Char8{
	0b00111110,
	0b01000000,
	0b01000000,
	0b00111100,
	0b00000010,
	0b00000010,
	0b01111100,
	0b00000000,
}

var T = Char8{
	0b01111110,
	0b00011000,
	0b00011000,
	0b00011000,
	0b00011000,
	0b00011000,
	0b00011000,
	0b00000000,
}

var U = Char8{
	0b01000010,
	0b01000010,
	0b01000010,
	0b01000010,
	0b01000010,
	0b01000010,
	0b00111100,
	0b00000000,
}

var V = Char8{
	0b01000010,
	0b01000010,
	0b01000010,
	0b01000010,
	0b01000010,
	0b00100100,
	0b00011000,
	0b00000000,
}

var W = Char8{
	0b01000010,
	0b01000010,
	0b01000010,
	0b01000010,
	0b01011010,
	0b01100110,
	0b01000010,
	0b00000000,
}

var X = Char8{
	0b01000010,
	0b00100100,
	0b00011000,
	0b00011000,
	0b00100100,
	0b01000010,
	0b01000010,
	0b00000000,
}

var Y = Char8{
	0b01000010,
	0b00100100,
	0b00011000,
	0b00011000,
	0b00011000,
	0b00011000,
	0b00011000,
	0b00000000,
}

var Z = Char8{
	0b01111110,
	0b00000100,
	0b00001000,
	0b00010000,
	0b00100000,
	0b01000000,
	0b01111110,
	0b00000000,
}

var ZERO = Char8{
	0b00111100,
	0b01000010,
	0b01000110,
	0b01001010,
	0b01010010,
	0b01100010,
	0b00111100,
	0b00000000,
}

var ONE = Char8{
	0b00011000,
	0b00111000,
	0b00011000,
	0b00011000,
	0b00011000,
	0b00011000,
	0b00111100,
	0b00000000,
}

var TWO = Char8{
	0b00111100,
	0b01000010,
	0b00000010,
	0b00000100,
	0b00011000,
	0b00100000,
	0b01111110,
	0b00000000,
}

var THREE = Char8{
	0b00111100,
	0b01000010,
	0b00000010,
	0b00111100,
	0b00000010,
	0b01000010,
	0b00111100,
	0b00000000,
}

var FOUR = Char8{
	0b00001100,
	0b00010100,
	0b00100100,
	0b01000100,
	0b01111110,
	0b00000100,
	0b00000100,
	0b00000000,
}

var FIVE = Char8{
	0b01111110,
	0b01000000,
	0b01111100,
	0b00000010,
	0b00000010,
	0b01000010,
	0b00111100,
	0b00000000,
}

var SIX = Char8{
	0b00111100,
	0b01000000,
	0b01111100,
	0b01000010,
	0b01000010,
	0b01000010,
	0b00111100,
	0b00000000,
}

var SEVEN = Char8{
	0b01111110,
	0b00000010,
	0b00000100,
	0b00001000,
	0b00010000,
	0b00010000,
	0b00010000,
	0b00000000,
}

var EIGHT = Char8{
	0b00111100,
	0b01000010,
	0b01000010,
	0b00111100,
	0b01000010,
	0b01000010,
	0b00111100,
	0b00000000,
}

var NINE = Char8{
	0b00111100,
	0b01000010,
	0b01000010,
	0b01000010,
	0b00111110,
	0b00000010,
	0b00111100,
	0b00000000,
}

var DOT = Char8{
	0b00000000,
	0b00000000,
	0b00000000,
	0b00000000,
	0b00000000,
	0b00011000,
	0b00011000,
	0b00000000,
}

var COLON = Char8{
	0b00000000,
	0b00011000,
	0b00011000,
	0b00000000,
	0b00011000,
	0b00011000,
	0b00000000,
	0b00000000,
}

var DASH = Char8{
	0b00000000,
	0b00000000,
	0b00000000,
	0b01111110,
	0b00000000,
	0b00000000,
	0b00000000,
	0b00000000,
}

var EQUAL = Char8{
	0b00000000,
	0b00000000,
	0b01111110,
	0b00000000,
	0b01111110,
	0b00000000,
	0b00000000,
	0b00000000,
}

var PLUS = Char8{
	0b00000000,
	0b00011000,
	0b00011000,
	0b01111110,
	0b00011000,
	0b00011000,
	0b00000000,
	0b00000000,
}

var STAR = Char8{
	0b00000000,
	0b00100100,
	0b00011000,
	0b01111110,
	0b00011000,
	0b00100100,
	0b00000000,
	0b00000000,
}

var HASH = Char8{
	0b00100100,
	0b01111110,
	0b00100100,
	0b00100100,
	0b01111110,
	0b00100100,
	0b00000000,
	0b00000000,
}

var PERCENT = Char8{
	0b01100010,
	0b01100100,
	0b00001000,
	0b00010000,
	0b00100000,
	0b01000110,
	0b01000110,
	0b00000000,
}

var AT = Char8{
	0b00111110,
	0b01000001,
	0b01011011,
	0b01011011,
	0b01011111,
	0b01000000,
	0b00111110,
	0b00000000,
}

var AMPERSAND = Char8{
	0b00000000,
	0b00110110,
	0b01001000,
	0b00100100,
	0b01001010,
	0b01010000,
	0b00110110,
	0b00000000,
}
