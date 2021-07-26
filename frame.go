package main

import (
	"fmt"
	// For reading touchscreen events
	"github.com/gvalkov/golang-evdev"
)

// Updates the touchstate frame by frame
func frameAnalyzer(EventFrameChan chan []evdev.InputEvent) {
	isMouseCtrlChan := make(chan mpts)
	isGestureTapChan := make(chan mpts)
	isGestureSwipeUD1Chan := make(chan mpts)
	go isMouseCtrl(isMouseCtrlChan)
	go isGestureTap(isGestureTapChan, threeTapKeyboard)
	go isGestureSwipeUD1(isGestureSwipeUD1Chan, isSwipeUD1Ukb, isSwipeUD1Dkb)
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
		isMouseCtrlChan <- touchState
		isGestureTapChan <- touchState
		isGestureSwipeUD1Chan <- touchState
		if isDebugMode {
			fmt.Println(touchState)
		}
		//fmt.Println("frameSTOP")
	}
}
