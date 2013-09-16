package main

import (
	"flag"
	. "led_strip"
	. "led_strip/effects"
	"log"
	_ "log"
	_ "math"
	"runtime"
	"time"
)

//point that grows and shrinks, moving in one direction with a random velocity
//color and size should depend on log level, need some random differentiator between the same log level getting received constantly
//should color depend on topic name? and size depend on log level?

//Entire screen should blink on an off color, with rate starting out at 30 hz and then growing slower over time until the timeout
//or should it alternate color? should transparency fade over time?

//trace - gray
//debug - blue
//info - green
//warn - orange
//error - red / yellow
//fatal - purple / white

var webDisplay = flag.Bool("webdisplay", false, "use webhost on localhost:8080 for the display")

// Application entry point
func main() {

	Settings.Read()

	flag.Parse()

	log.Print("MinFrameTime is ", Settings.MinFrameTime)

	var display Display
	if *webDisplay || runtime.GOOS == "windows" {
		display = NewWebDisplay(Settings)
	} else {
		display = NewLedDisplay(Settings)
	}

	newDrawables := make(chan Drawable, 100)
	go CreateDrawables(newDrawables)

	strip := NewLedStrip(Settings.LedCount)
	curTime := time.Now()
	prevTime := curTime

	renderTick := time.Tick(time.Duration(Settings.MinFrameTime*1000.0) * time.Millisecond)
	for _ = range renderTick {

		prevTime, curTime = curTime, time.Now()
		dt := curTime.Sub(prevTime).Seconds()

		// check for new things to add to strip
		for len(newDrawables) > 0 {
			strip.Add(<-newDrawables)
		}

		strip.Animate(dt)
		strip.RenderTo(display)
	}
}

// Test generate of drawables
func CreateDrawables(newDrawables chan<- Drawable) {

	var count int = 0
	for {
		time.Sleep(250 * time.Millisecond)
		newDrawables <- NewParticle(0, 32, 3)

		count++
		if count%10 == 0 {
			newDrawables <- NewFlash()
		}
	}
}
