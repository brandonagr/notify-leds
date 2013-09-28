package effects

import (
	//"fmt"
	. "led_strip"
)

// One color with alpha transparency that
type Gradient struct {
	maxPosition float64
	color       RGBA
}

// Construct a Flash
func NewGradient(max float64, color RGBA) Drawable {
	return &Gradient{
		maxPosition: max,
		color:       color,
	}
}

// Returns the color at position blended on top of baseColor
func (this *Gradient) ColorAt(position float64, baseColor RGBA) (color RGBA) {

	color = this.color
	color.A = uint8(255.0 * (this.maxPosition - position) / this.maxPosition)

	result := color.BlendWith(baseColor)
	//fmt.Println(result, color, baseColor)
	return result
}

// ZIndex of the drawable
func (this *Gradient) ZIndex() ZIndex {
	return 1
}

// Animate forward in time
func (this *Gradient) Animate(dt float64) bool {
	return true
}
