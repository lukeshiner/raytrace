package canvas

import (
	"strings"
	"testing"

	"github.com/lukeshiner/raytrace/colour"
)

func TestCanvas(t *testing.T) {
	c := New(10, 20)
	if c.Width != 10 {
		t.Error("Canvas width not set correctly.")
	}
	if c.Height != 20 {
		t.Error("Canvas height not set correctly.")
	}
	for _, row := range c.Pixels {
		for _, pixel := range row {
			if pixel.Red != 0 || pixel.Green != 0 || pixel.Blue != 0 {
				t.Error("Canvas initialized with non black pixel.")
			}
		}
	}
}

func TestWritePixel(t *testing.T) {
	canvas := New(10, 20)
	red := colour.Colour{Red: 1, Green: 0, Blue: 0}
	canvas.WritePixel(2, 3, red)
	if canvas.Pixel(2, 3).Equal(red) != true {
		t.Error("Error retriening pixel from canvas.")
	}
}

func TestPPMHeader(t *testing.T) {
	canvas := New(5, 3)
	expected := "P3\n5 3\n255"
	output := canvas.ToPPM()
	lines := strings.Split(output, "\n")
	header := strings.Join(lines[:3], "\n")
	if header != expected {
		t.Errorf("PPM header incorrect: expected \"%s\", got \"%s\".", expected, header)
	}
}

func TestPPMFormatPixel(t *testing.T) {
	var tests = []struct {
		colour   colour.Colour
		expected string
	}{
		{colour.Colour{Red: 0, Green: 0, Blue: 0}, "0 0 0"},
		{colour.Colour{Red: 1, Green: 0, Blue: 0}, "255 0 0"},
		{colour.Colour{Red: 0, Green: 0.5, Blue: 0}, "0 128 0"},
		{colour.Colour{Red: 0, Green: 0, Blue: -5}, "0 0 0"},
		{colour.Colour{Red: 1.5, Green: 0, Blue: -5}, "255 0 0"},
	}

	for _, test := range tests {
		output := ppmFormatPixel(test.colour)
		if output != test.expected {
			t.Errorf(
				"Incorrect PPM pixel output: Colour: %+v, expected \"%v\", recieved \"%v\".",
				test.colour, test.expected, output)
		}
	}
}

func TestPPMPixelData(t *testing.T) {
	canvas := New(5, 3)
	c1 := colour.Colour{Red: 1.5, Green: 0, Blue: 0}
	c2 := colour.Colour{Red: 0, Green: 0.5, Blue: 0}
	c3 := colour.Colour{Red: -0.5, Green: 0, Blue: 1}
	canvas.WritePixel(0, 0, c1)
	canvas.WritePixel(2, 1, c2)
	canvas.WritePixel(4, 2, c3)
	output := canvas.ToPPM()
	expected := "255 0 0 0 0 0 0 0 0 0 0 0 0 0 0\n0 0 0 0 0 0 0 128 0 0 0 0 0 0 0\n0 0 0 0 0 0 0 0 0 0 0 0 0 0 255"
	lines := strings.Split(output, "\n")
	testLines := strings.Join(lines[3:6], "\n")
	if testLines != expected {
		t.Errorf("Expected: \"%v\", got \"%v\".", expected, testLines)
	}
}

func TestPPMLineLength(t *testing.T) {
	canvas := New(10, 2)
	colour := colour.Colour{Red: 1, Green: 0.8, Blue: 0.6}
	for x := 0; x < canvas.Width; x++ {
		for y := 0; y < canvas.Height; y++ {
			canvas.WritePixel(x, y, colour)
		}
	}
	output := canvas.ToPPM()
	expected := "255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204\n153 255 204 153 255 204 153 255 204 153 255 204 153\n255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204\n153 255 204 153 255 204 153 255 204 153 255 204 153\n"
	lines := strings.Split(output, "\n")
	testLines := strings.Join(lines[3:], "\n")
	if testLines != expected {
		t.Errorf("PPM line length incorrect:\n Expected:\t%q\nGot:\t\t%q.", expected, testLines)
	}
}

func TestPPMEndsWithNewline(t *testing.T) {
	canvas := New(5, 3)
	output := canvas.ToPPM()
	if output[len(output)-1] != '\n' {
		t.Error("PPM output did not end with a newline.")
	}
}
