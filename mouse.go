package main

import (
	"fmt"
	"time"
	// For moving mouse cursor
	"github.com/go-vgo/robotgo"
)

// Handles all basic mouse controls.
// movement, drag, left/right click
func mouseCtrl(StateChan chan mpts) {
	//var TapCnt int32
	var isStart bool
	var isMouseDown bool
	var TouchX, TouchY int = -1, -1
	var TouchTime time.Time
	for {
		tempState := <-StateChan
		if !isLocked {
			if tempState.isDown[1] {
				isStart, isMouseDown = false, false
				TouchX, TouchY = -1, -1
				continue
			}
			if tempState.isDown[0] && !isStart {
				isStart = true
				TouchTime = time.Now()
				//TapCnt = tempState.tapCnt
			}
			if isStart {
				if TouchX != int(touchState.X[0]) {
					TouchX = int(touchState.X[0])
				}
				if TouchY != int(touchState.Y[0]) {
					TouchY = int(touchState.Y[0])
				}
				if TouchX > 0 && TouchY > 0 { //works but wrong, cursor cant be at 0,0
					switch touchState.rotation {
					case inv:
						xval := float32(TouchX) / (float32((maxX + 1)) / float32(resolX))
						xinv := halfResolX + (halfResolX - int32(xval))
						yval := float32(TouchY) / (float32((maxY + 1)) / float32(resolY))
						yinv := halfResolY + (halfResolY - int32(yval))
						robotgo.MoveMouse(int(xinv), int(yinv))
					case norm:
						xval := float32(TouchX) / (float32((maxX + 1)) / float32(resolX))
						yval := float32(TouchY) / (float32((maxY + 1)) / float32(resolY))
						robotgo.MoveMouse(int(xval), int(yval))
					}
					if !isMouseDown {
						mouse.LeftPress()
						fmt.Println("LEFT CLICK DOWN!")
						isMouseDown = true
					}
				}
				if !touchState.isDown[0] {
					TimeDiff := time.Now().Sub(TouchTime)
					if TimeDiff > 600*time.Millisecond && TimeDiff < 2250*time.Millisecond {
						mouse.LeftRelease()
						fmt.Println("LEFT CLICK UP!")
						mouse.RightClick()
						fmt.Println("RIGHT CLICK!")
					} else {
						mouse.LeftRelease()
						fmt.Println("LEFT CLICK UP!")
					}
					isStart, isMouseDown = false, false
					TouchX, TouchY = -1, -1
				}
			}
		}
	}
}
