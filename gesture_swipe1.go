package main

import (
	"fmt"
	// For reading touchscreen events
	"github.com/gvalkov/golang-evdev"
	// For executing keyboard actions
	"github.com/micmonay/keybd_event"
	"time"
)

func isGestureSwipeUD1 (StateChan chan mpts, UpKeyboard keybd_event.KeyBonding, DownKeyboard keybd_event.KeyBonding) {
	var startX, startY int32
	var currX, currY int32
	var swipeLength = halfY
	var subSwipeLength = swipeLength / 3
	var lastTap int32
	var isStart bool
	var touchTime time.Time
	var swipeUpCheck, swipeDnCheck [3]bool
	for {
		tempState := <-StateChan
		if !isLocked && !isStart && tempState.tapCnt > lastTap {
			if tempState.isGesture {
				isStart = true
				touchTime = time.Now()
				startX = int32(tempState.X[0])
				startY = int32(tempState.Y[0])
				lastTap = tempState.tapCnt
				continue
			}
		}
		if !isLocked && isStart {
			if tempState.isDown[0] {
				currX = int32(tempState.X[0])
				currY = int32(tempState.Y[0])
			}
			if currY >= startY + subSwipeLength {
				swipeDnCheck[0] = true
			} else if currY <= startY - subSwipeLength {
				swipeUpCheck[0] = true
			}
			if currY >= startY + (subSwipeLength * 2) {
				swipeDnCheck[1] = true
			} else if currY <= startY - (subSwipeLength * 2) {
				swipeUpCheck[1] = true
			}
			if currY >= startY + swipeLength {
				swipeDnCheck[2] = true
			} else if currY <= startY - swipeLength {
				swipeUpCheck[2] = true
			}
			if !tempState.isGesture && time.Now().Sub(touchTime) < 2*time.Second {
				if swipeUpCheck[0] == true && swipeUpCheck[1] == true && swipeUpCheck[2] == true {
					switch touchState.rotation {
					case inv:
						fmt.Println("DOWN SWIPE!")
						execKeys(DownKeyboard)
					case norm:
						fmt.Println("UP SWIPE!")
						execKeys(UpKeyboard)
					}
				} else if swipeDnCheck[0] == true && swipeDnCheck[1] == true && swipeDnCheck[2] == true {
					switch touchState.rotation {
					case inv:
						fmt.Println("UP SWIPE!")
						execKeys(UpKeyboard)
					case norm:
						fmt.Println("DOWN SWIPE!")
						execKeys(DownKeyboard)
					}
				}
				isStart = false
				swipeUpCheck[0], swipeUpCheck[1], swipeUpCheck[2] = false, false, false
				swipeDnCheck[0], swipeDnCheck[1], swipeDnCheck[2] = false, false, false
				continue
			}
			if tempState.isDown[1] || currX > startX+(maxX/10) || currX < startX-(maxX/10) || currX > maxX-(maxX/20) || currX < (maxX/20) || time.Now().Sub(touchTime) > 2*time.Second {
				fmt.Println("BAIL FROM SWIPE!")
				// Bail-out check includes multi-touch
				// and out of tolerance for being too angled
				// or by screen edges to avoid scrollbar
				// and lastly if over 2 seconds
				isStart = false
				swipeUpCheck[0], swipeUpCheck[1], swipeUpCheck[2] = false, false, false
				swipeDnCheck[0], swipeDnCheck[1], swipeDnCheck[2] = false, false, false
			}
		}
	}
}

// Old format code, will be converted to mpts code like the above f(x)
func isSwipeRL1(EventChan chan *evdev.InputEvent, RightKeyboard keybd_event.KeyBonding, LeftKeyboard keybd_event.KeyBonding) {
	var StartX, StartY int32 = -1, -1
	var CurrX, CurrY int32
	var SwipeLength = (maxX + 1) / 3
	var SubSwipeLength = SwipeLength / 3
	var TouchStart = false
	//var TapCnt int32 //currently unused but update value with event.
	var SwipeRtCheck1, SwipeRtCheck2, SwipeRtCheck3 = false, false, false
	var SwipeLtCheck1, SwipeLtCheck2, SwipeLtCheck3 = false, false, false
	var StartTime time.Time
	for {
		event := <-EventChan
		if isEventMatch(event, touchDnEv) {
			StartTime = time.Now()
			TouchStart = true
			//TapCnt = event.Value
		} else if TouchStart == true {
			if isEventMatch(event, touchXEv) {
				//TapCnt = event.Value
				if StartX != -1 {
					CurrX = event.Value
					if CurrX >= StartX+SubSwipeLength {
						SwipeRtCheck1 = true
					} else if CurrX <= StartX-SubSwipeLength {
						SwipeLtCheck1 = true
					}
					if CurrX >= StartX+(SubSwipeLength*2) {
						SwipeRtCheck2 = true
					} else if CurrX <= StartX-(SubSwipeLength*2) {
						SwipeLtCheck2 = true
					}
					if CurrX >= StartX+SwipeLength {
						SwipeRtCheck3 = true
					} else if CurrX <= StartX-SwipeLength {
						SwipeLtCheck3 = true
					}
				} else {
					StartX = event.Value
				}
			} else if isEventMatch(event, touchYEv) {
				if StartY != -1 {
					CurrY = event.Value
					if CurrY > StartY+(maxY/10) || CurrY < StartY-(maxY/10) {
						//Out of tolerance relook for new touch
						TouchStart = false
						SwipeRtCheck1, SwipeRtCheck2, SwipeRtCheck3 = false, false, false
						SwipeLtCheck1, SwipeLtCheck2, SwipeLtCheck3 = false, false, false
						StartX, StartY = -1, -1
					}
				} else {
					StartY = event.Value
				}
			} else if isEventMatch(event, touchUpEv) {
				if time.Now().Sub(StartTime) < 2*time.Second {
					if SwipeRtCheck1 == true && SwipeRtCheck2 == true && SwipeRtCheck3 == true {
						switch touchState.rotation {
						case inv:
							fmt.Println("LEFT SWIPE!")
							execKeys(LeftKeyboard)
						case norm:
							fmt.Println("RIGHT SWIPE!")
							execKeys(RightKeyboard)
						}
					} else if SwipeLtCheck1 == true && SwipeLtCheck2 == true && SwipeLtCheck3 == true {
						switch touchState.rotation {
						case inv:
							fmt.Println("RIGHT SWIPE!")
							execKeys(RightKeyboard)
						case norm:
							fmt.Println("LEFT SWIPE!")
							execKeys(LeftKeyboard)
						}
					}
				}
				TouchStart = false
				SwipeRtCheck1, SwipeRtCheck2, SwipeRtCheck3 = false, false, false
				SwipeLtCheck1, SwipeLtCheck2, SwipeLtCheck3 = false, false, false
				StartX, StartY = -1, -1
			}
		}
	}
}
