package effects

import (
	. "led_strip"
	"math"
)

// Particle that will be drawn on the screen
type FadingParticle struct {

	// color of the particle
	color RGBA

	// current position of the ball
	position float64

	// direction and speed of the ball in leds / second
	velocity float64

	// radius of the particle
	size float64

	// amount of item particle has existed for
	lifeSoFar float64

	// max time particle will exist in seconds
	lifeMax float64

	// z position of ball
	zindex ZIndex
}

// Construct a Particle that fades in and then out over its lifetime
func NewFadingParticle(position, velocity, size float64, color RGBA) Drawable {
	return &FadingParticle{
		color:    color,
		position: position,
		velocity: velocity,
		size:     size,
		zindex:   100,
		lifeMax:  2,
	}
}

// Returns the color at position blended on top of baseColor
func (this *FadingParticle) ColorAt(position float64, baseColor RGBA) (color RGBA) {

	distance := math.Abs(position - this.position)

	if distance < this.size {
		fadedColor := this.color

		lifePercentage := this.lifeSoFar / this.lifeMax
		maxAlpha := 255.0
		if lifePercentage < 0.5 {
			maxAlpha = (lifePercentage * 2.0) * 255.0
		} else {
			maxAlpha = (1.0 - lifePercentage) * 2.0 * 255.0
		}

		fadedColor.A = uint8((this.size - distance) * maxAlpha)
		color = fadedColor.BlendWith(baseColor)
	} else {
		color = baseColor
	}

	return color
}

// ZIndex of the ball
func (this *FadingParticle) ZIndex() ZIndex {
	return this.zindex
}

// Animate ball
func (this *FadingParticle) Animate(dt float64) bool {
	this.position += this.velocity * dt
	this.lifeSoFar += dt

	return this.lifeSoFar < this.lifeMax
}
