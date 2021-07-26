package main

import (
	"fmt"
	// For reading touchscreen events
	"github.com/gvalkov/golang-evdev"
)

func calibrateEdges() bool {
	var x, y int32
	var QuadTouch = false
	var TouchCnt int8
	fmt.Printf("Max X %d Max Y %d", x, y)
	device, _ := evdev.Open(devPath)
	for {
		event, _ := device.ReadOne()
		if isEventMatch(event, touchXEv) {
			if event.Value > x {
				x = event.Value
			}
		} else if isEventMatch(event, touchYEv) {
			if event.Value > y {
				y = event.Value
			}
		} else if isEventMatch(event, touchIs4) {
			QuadTouch = true
		} else if isEventMatch(event, touchUpEv) && QuadTouch == true {
			TouchCnt++
		}
		if TouchCnt == 4 {
			break
		}
		fmt.Printf("\rMax X %d Max Y %d", x, y)
	}
	fmt.Printf("\n4+ Finger Exit Tap Detected!\n")
	fmt.Println("Max X set to", x)
	fmt.Println("Max Y set to", y)
	maxX = x
	maxY = y
	return true
}
