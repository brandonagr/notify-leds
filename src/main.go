package main

import (
	"encoding/xml"
	"flag"
	"github.com/gmallard/stompngo"
	. "led_strip"
	. "led_strip/effects"
	"log"
	_ "log"
	_ "math"
	"net"
	"runtime"
	"strings"
	"time"
)

// Command line flags
var webDisplay = flag.Bool("webdisplay", false, "use webhost on localhost:8080 for the display")
var brokerHost = flag.String("brokerHost", "localhost", "Broker host to connect to")
var brokerPort = flag.String("brokerPort", "61613", "Broker port to connect to")

// Application entry point
func main() {

	Settings.Read()

	flag.Parse()

	log.Print("MinFrameTime is ", Settings.MinFrameTime)

	// channel that will allow background goroutine to send new Drawable objects to the render thread
	newDrawables := make(chan Drawable, 100)

	var display Display

	if *webDisplay || runtime.GOOS == "windows" {
		display = NewWebDisplay(Settings)
	} else {
		display = NewLedDisplay(Settings)
	}

	if runtime.GOOS == "windows" {
		//go GenerateDrawablesTimer(newDrawables)
		go GenerateDrawablesLogs(*brokerHost, *brokerPort, newDrawables)
	} else {
		go GenerateDrawablesLogs(*brokerHost, *brokerPort, newDrawables)
	}

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
func GenerateDrawablesTimer(newDrawables chan<- Drawable) {

	var count int = 0
	for {
		time.Sleep(500 * time.Millisecond)
		newDrawables <- NewParticle(0, 32, 1, RGBA{128, 128, 0, 255})

		count++
		if count%20 == 0 {
			newDrawables <- NewFlash(5, 0.5, 0.5, [2]RGBA{RGBA{255, 0, 0, 128}, RGBA{255, 255, 255, 255}})
		}
	}
}

// Type that is received from a message
type LogMessage struct {
	ApplicationName string
	LogType         string
	EntryDate       string
	Description     string
}

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

// Generate drawables from the broker
// would it be better to have this return a channel of LogMessage, and chain channels together to form pipeline to generate Drawable? Only if want to add another stage to the filter in the future
func GenerateDrawablesLogs(host, port string, newDrawables chan<- Drawable) {
	log.Println("Connecting to ", host, ":", port, " ...")
	tcpConnection, err := net.Dial("tcp", net.JoinHostPort(host, port))
	if err != nil {
		log.Fatalln(err)
	}

	headers := stompngo.Headers{"accept-version", "1.1", "host", host}
	connection, err := stompngo.Connect(tcpConnection, headers)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Stomp connect complete ...", connection.Protocol())

	u := stompngo.Uuid()
	s := stompngo.Headers{"destination", "/topic/*", "id", u}
	r, err := connection.Subscribe(s)
	if err != nil {
		log.Fatalln(err)
	}
	for message := range r {
		if message.Error != nil {
			log.Fatalln(message.Error)
		}

		log.Println("Received message", strings.Join(message.Message.Headers, ","))

		logMessage := LogMessage{}
		if err = xml.Unmarshal(message.Message.Body, &logMessage); err != nil {
			log.Printf("Failed to unmarshal", err, message.Message.Body)
		} else {
			newDrawables <- CreateDrawableFromLog(logMessage)
		}
	}
}

// Convert LogMessage to Drawable
func CreateDrawableFromLog(message LogMessage) Drawable {
	logType := strings.ToLower(message.LogType)
	particleColor := RGBA{}

	switch logType {
	case "trace":
		particleColor = RGBA{128, 128, 128, 255}
		break
	case "debug":
		particleColor = RGBA{0, 0, 255, 255}
		break
	case "info", "informational":
		particleColor = RGBA{0, 255, 0, 255}
		break
	case "warn":
		particleColor = RGBA{255, 128, 0, 255}
		break
	case "error":
		return NewFlash(10, 0.25, 0.25, [2]RGBA{RGBA{255, 0, 0, 255}, RGBA{0, 0, 255, 128}})
		break
	case "fatal":
		return NewFlash(20, 0.5, 0.5, [2]RGBA{RGBA{255, 255, 255, 255}, RGBA{255, 0, 255, 255}})
		break
	default:
		log.Fatalln("Unexpected logType of ", logType)
		break
	}

	// need to generate a random position
	// random position in first 2/3 of strip
	// velocity based on size of message
	// size based on size of message
	// color based on LogType
	return NewFadingParticle(0, 32, 2, particleColor) //position, velocity, size float64)
}
