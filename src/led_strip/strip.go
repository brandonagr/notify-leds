package led_strip

import (
	"container/list"
)

// Defines all of the information
type LedStrip struct {

	// Size of the strip, from 0 to width exclusive
	width int

	// All of the drawable items, stored in increasing ZIndex order
	drawables *list.List

	// Buffer used to render the strip
	renderBuffer []RGBA
}

// Initialized a new strip
func NewLedStrip(width int) *LedStrip {

	return &LedStrip{
		width:        width,
		drawables:    list.New(),
		renderBuffer: make([]RGBA, width),
	}
}

// Adds a drawable to the strip
func (strip *LedStrip) Add(addDrawable Drawable) {

	curElement := strip.drawables.Front()

	if curElement == nil {
		strip.drawables.PushFront(addDrawable)
		return
	}

	for ; curElement != nil; curElement = curElement.Next() {

		curDrawable := curElement.Value.(Drawable)

		if addDrawable.ZIndex() <= curDrawable.ZIndex() {
			strip.drawables.InsertBefore(addDrawable, curElement)
			return
		}
	}

	// got here so it wasn't less than any
	strip.drawables.PushBack(addDrawable)
}

// Determines the color at the given position
func (strip *LedStrip) ColorAt(position float64) RGBA {

	color := RGBA{0, 0, 0, 255}

	for curElement := strip.drawables.Front(); curElement != nil; curElement = curElement.Next() {

		drawable := curElement.Value.(Drawable)

		color = drawable.ColorAt(position, color)
	}

	return color
}

// Animate all Drawables
func (strip *LedStrip) Animate(dt float64) {

	for curElement := strip.drawables.Front(); curElement != nil; {

		drawable := curElement.Value.(Drawable)

		if !drawable.Animate(dt) {
			nextElement := curElement.Next()
			strip.drawables.Remove(curElement)
			curElement = nextElement
		} else {
			curElement = curElement.Next()
		}
	}
}

// Render each integer position and pass that to the Display
func (strip *LedStrip) RenderTo(display Display) {

	for ledIndex := 0; ledIndex < strip.width; ledIndex++ {
		strip.renderBuffer[ledIndex] = strip.ColorAt(float64(ledIndex))
	}
	display.Render(strip.renderBuffer)
}

// Returns true if the strip of drawables is valid
func (strip *LedStrip) IsValid() bool {

	var prevDrawable Drawable = nil

	for curElement := strip.drawables.Front(); curElement != nil; curElement = curElement.Next() {

		curDrawable := curElement.Value.(Drawable)

		if prevDrawable != nil {
			if prevDrawable.ZIndex() > curDrawable.ZIndex() {
				return false
			}
		}

		prevDrawable = curDrawable
	}

	return true
}

// Return number of drawables in the strip
func (strip *LedStrip) DrawableLen() int {
	return strip.drawables.Len()
}

// Width of the strip
func (strip *LedStrip) Width() int {
	return strip.width
}
