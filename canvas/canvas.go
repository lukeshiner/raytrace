package canvas

import (
	"fmt"
	"math"
	"strings"

	"github.com/lukeshiner/raytrace/colour"
)

// Canvas holds pixel grid data
type Canvas struct {
	Width, Height int
	Pixels        [][]colour.Colour
}

// WritePixel writes a pixel to the canvas.
func (c *Canvas) WritePixel(x, y int, colour colour.Colour) {
	c.Pixels[x][y] = colour
}

// Pixel returns the colour at pixel co-ordinates (x,y).
func (c *Canvas) Pixel(x, y int) colour.Colour {
	return c.Pixels[x][y]
}

// ToPPM returns the canvas as a PPM string.
func (c *Canvas) ToPPM() string {
	return c.ppmHeader() + c.ppmColours()
}

func (c *Canvas) ppmHeader() string {
	return fmt.Sprintf("P3\n%d %d\n255\n", c.Width, c.Height)
}

func (c *Canvas) ppmColours() string {
	line := ""
	colours := ""
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			line += ppmFormatPixel(c.Pixel(x, y)) + " "
		}
		line = formatPPMLine(line)
		colours += line
		line = ""
	}
	return colours
}

func formatPPMLine(line string) string {
	lines := []string{line}
	lines = splitLine(lines, 70)
	return strings.TrimSpace(strings.Join(lines, "\n")) + "\n"
}

func splitLine(lines []string, length int) []string {
	line := lines[len(lines)-1]
	if len(line) > length {
		index := length - 1
		for line[index] != ' ' {
			index--
		}
		lines[len(lines)-1] = line[:index]
		lines = append(lines, line[index+1:])
	} else {
		return lines
	}
	return splitLine(lines, length)
}

func clampColour(c float64) int {
	newC := int(math.Ceil(255 * c))
	if newC >= 255 {
		return 255
	}
	if newC <= 0 {
		return 0
	}
	return newC
}

func ppmFormatPixel(c colour.Colour) string {
	return fmt.Sprintf(
		"%d %d %d", clampColour(c.Red), clampColour(c.Green), clampColour(c.Blue))
}

// New creates a new Canvas.
func New(width, height int) Canvas {
	var pixels [][]colour.Colour
	for x := 0; x < width; x++ {
		row := []colour.Colour{}
		for y := 0; y < height; y++ {
			row = append(row, colour.New(0, 0, 0))
		}
		pixels = append(pixels, row)
	}
	return Canvas{width, height, pixels}
}
