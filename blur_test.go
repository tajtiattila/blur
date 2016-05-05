package blur_test

import (
	"image"
	"image/png"
	"os"
	"testing"

	"github.com/tajtiattila/blur"
)

func TestBlur(t *testing.T) {
	f, err := os.Open("sample/cballs.png")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	im, _, err := image.Decode(f)
	if err != nil {
		t.Fatal(err)
	}

	dst := blur.Gaussian(im, 5, blur.ReuseSrc)

	f, err = os.Create("sample/cballs_blur.png")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	if err := png.Encode(f, dst); err != nil {
		t.Fatal(err)
	}
}
