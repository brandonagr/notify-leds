package effects

import (
	//"fmt"
	. "led_strip"
)

// Flash that changes frequency and fades out
type Flash struct {

	// time in seconds this has been happening
	lifeSoFar float64

	// max time flash will take
	lifeMax float64

	// hz, beginning at a higher rate
	frequencyBegin float64

	// hz, ending at a lower rate
	frequencyEnd float64

	// Colors that are alternatied between
	colors [2]RGBA
}

// Construct a Flash
func NewFlash(lifeMax, frequencyBegin, frequencyEnd float64, colors [2]RGBA) Drawable {
	return &Flash{
		lifeMax:        lifeMax,
		frequencyBegin: frequencyBegin,
		frequencyEnd:   frequencyEnd,
		colors:         colors,
	}
}

// Returns the color at position blended on top of baseColor
func (this *Flash) ColorAt(position float64, baseColor RGBA) (color RGBA) {

	lifePercentage := (this.lifeMax - this.lifeSoFar) / this.lifeMax
	frequency := (this.frequencyBegin-this.frequencyEnd)*lifePercentage + this.frequencyBegin

	colorIndex := int(frequency/this.lifeSoFar) % 2
	color = this.colors[colorIndex]
	if color.A > 0 {
		color.A = uint8(lifePercentage * 255.0)
	}

	//fmt.Println(lifePercentage, frequency, colorIndex, color.A)

	return color.BlendWith(baseColor)
}

// ZIndex of the ball
func (this *Flash) ZIndex() ZIndex {
	return 1000
}

// Animate forward in time
func (this *Flash) Animate(dt float64) bool {
	this.lifeSoFar += dt

	return this.lifeSoFar < this.lifeMax
}
