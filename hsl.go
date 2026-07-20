package excelize

import (
	"image/color"
)

var HSLModel = color.ModelFunc(hslModel)

type HSL struct {
	H, S, L float64
}

func (c HSL) RGBA() (uint32, uint32, uint32, uint32) { _ = "STUB: not implemented"; return 0, 0, 0, 0 }

func hslModel(c color.Color) color.Color { _ = "STUB: not implemented"; return *new(color.Color) }

func RGBToHSL(r, g, b uint8) (h, s, l float64) { _ = "STUB: not implemented"; return 0, 0, 0 }

func HSLToRGB(h, s, l float64) (r, g, b uint8) { _ = "STUB: not implemented"; return 0, 0, 0 }

func hueToRGB(p, q, t float64) float64 { _ = "STUB: not implemented"; return 0 }
