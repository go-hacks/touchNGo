package main

import (
	"fmt"
	// For reading touchscreen events
	"github.com/gvalkov/golang-evdev"
)

// Updates the touchstate frame by frame
func frameAnalyzer(EventFrameChan chan []evdev.InputEvent) {
	mouseCtrlChan := make(chan mpts)
	gestureTapChan := make(chan mpts)
	gestureSwipeUp1Chan := make(chan mpts)
	gestureSwipeDown1Chan := make(chan mpts)
	go mouseCtrl(mouseCtrlChan)
	go gestureTap(gestureTapChan)
	go gestureSwipeUp1(gestureSwipeUp1Chan)
	go gestureSwipeDown1(gestureSwipeDown1Chan)
	var currPoint int32
	for {
		eventFrame := <-EventFrameChan
		for i := 0; i < len(eventFrame); i++ {
			event := eventFrame[i]
			if isEventMatch(&event, gestureStart) {
				touchState.isGesture = true
			} else if isEventMatch(&event, gestureStop) {
				touchState.isGesture = false
				resetTouchState()
			} else if isEventMatch(&event, touchDnEv) {
				currPoint = 0
				touchState.isDown[0] = true
				touchState.tapCnt = event.Value
			} else if event.Code == 47 {
				currPoint = event.Value
				touchState.isDown[currPoint] = true
			} else if isEventMatch(&event, touchXEv) {
				touchState.X[currPoint] = uint16(event.Value)
			} else if isEventMatch(&event, touchYEv) {
				touchState.Y[currPoint] = uint16(event.Value)
			} else if isEventMatch(&event, touchUpEv) {
				touchState.isDown[currPoint] = false
				touchState.X[currPoint] = 0
				touchState.Y[currPoint] = 0
			}
		}
		mouseCtrlChan <- touchState
		gestureTapChan <- touchState
		gestureSwipeUp1Chan <- touchState
		gestureSwipeDown1Chan <- touchState
		if isDebugMode {
			fmt.Println(touchState)
		}
	}
}
