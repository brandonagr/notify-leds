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

	// the time at which the color will be changed next
	changeColorAtTime float64

	// Currently displayed color index, 0 or 1
	colorIndex int

	// Colors that are alternatied between
	colors [2]RGBA
}

// Construct a Flash
func NewFlash(lifeMax, frequencyBegin, frequencyEnd float64, colors [2]RGBA) Drawable {
	return &Flash{
		lifeMax:           lifeMax,
		frequencyBegin:    frequencyBegin,
		frequencyEnd:      frequencyEnd,
		colors:            colors,
		changeColorAtTime: 1.0 / frequencyBegin,
	}

}

// Returns the color at position blended on top of baseColor
func (this *Flash) ColorAt(position float64, baseColor RGBA) (color RGBA) {

	color = this.colors[this.colorIndex]
	if color.A > 0 {
		lifeRemainingPercentage := (this.lifeMax - this.lifeSoFar) / this.lifeMax
		color.A = uint8(lifeRemainingPercentage * 255.0)
	}

	return color.BlendWith(baseColor)
}

// ZIndex of the ball
func (this *Flash) ZIndex() ZIndex {
	return 1000
}

// Animate forward in time
func (this *Flash) Animate(dt float64) bool {
	this.lifeSoFar += dt

	for this.lifeSoFar > this.changeColorAtTime {

		percentage := this.changeColorAtTime / this.lifeMax
		frequency := (this.frequencyEnd-this.frequencyBegin)*percentage + this.frequencyBegin

		this.changeColorAtTime = this.changeColorAtTime + 1.0/frequency
		this.colorIndex = (this.colorIndex + 1) % 2

		//fmt.Println(percentage, frequency, this.lifeSoFar, this.colorIndex, this.changeColorAtTime)
	}

	return this.lifeSoFar < this.lifeMax
}
