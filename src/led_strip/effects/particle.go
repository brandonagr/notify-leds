package effects

import (
	. "led_strip"
	"math"
)

// Particle that will be drawn on the screen
type Particle struct {

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

var _ Drawable = &Particle{}

// Construct a Particle
func NewParticle(position, velocity, size float64, color RGBA) *Particle {
	return &Particle{
		color:    color,
		position: position,
		velocity: velocity,
		size:     size,
		zindex:   100,
		lifeMax:  2,
	}
}

// Returns the color at position blended on top of baseColor
func (this *Particle) ColorAt(position float64, baseColor RGBA) (color RGBA) {

	distance := math.Abs(position - this.position)

	if distance < this.size {
		fadedColor := this.color
		fadedColor.A = uint8((this.size - distance) * 255.0)
		color = fadedColor.BlendWith(baseColor)
	} else {
		color = baseColor
	}

	return color
}

// ZIndex of the ball
func (this *Particle) ZIndex() ZIndex {
	return this.zindex
}

// Animate ball
func (this *Particle) Animate(dt float64) bool {
	this.position += this.velocity * dt
	this.lifeSoFar += dt

	return this.lifeSoFar < this.lifeMax
}
