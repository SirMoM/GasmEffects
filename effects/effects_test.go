package effects

import (
	"bytes"
	"testing"
)

func TestGreyScale(t *testing.T) {
	// Create a sample 2x1 pixel image (Red, Green)
	inputBytes := []byte{
		255, 0, 0, 255, // Red pixel
		0, 255, 0, 255, // Green pixel
	}
	img := RgbaFromBytes(inputBytes)

	// Expected output after grayscale conversion
	// R: 0.3 * 255 = 76.5 -> 76
	// G: 0.59 * 255 = 150.45 -> 150
	expectedBytes := []byte{
		76, 76, 76, 255,
		150, 150, 150, 255,
	}

	// Apply the grayscale effect
	img.GreyScale()
	outputBytes := img.ToBytes()

	// Compare the result
	if !bytes.Equal(outputBytes, expectedBytes) {
		t.Errorf("GreyScale was incorrect, got: %v, want: %v.", outputBytes, expectedBytes)
	}
}
