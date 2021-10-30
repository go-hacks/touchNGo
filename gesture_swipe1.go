package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"time"
)

func downCheckOne (checkChan chan bool, StateChan chan mpts, startX int32, startY int32) {
	startTime := time.Now()
	var swipeLength = halfY
	var subSwipeLength = swipeLength / 3
	for {
		tempState := <-StateChan
		currX := int32(tempState.X[0])
		currY := int32(tempState.Y[0])
		// fail if angled or over scrollbar(both screen edges)
		if currX > startX+(maxX/10) || currX < startX-(maxX/10) || currX > maxX-(maxX/20) || currX < (maxX/20) {
			checkChan <- false
			break
		}
		// fail if multi-touch, swipe up, or over 2 sec(timeout)
		if tempState.isDown[1] || currY < startY - 15 || time.Now().Sub(startTime) > 2*time.Second {
			checkChan <- false
			break
		}
		// pass if reaches down far enough without previous fails
		if currY >= startY + subSwipeLength {
			checkChan <- true
			break
		}
	}
	return
}

func downCheckTwo (checkChan chan bool, StateChan chan mpts, startX int32, startY int32) {
	startTime := time.Now()
	var swipeLength = halfY
	var subSwipeLength = swipeLength / 3
	for {
		tempState := <-StateChan
		currX := int32(tempState.X[0])
		currY := int32(tempState.Y[0])
		// fail if angled or over scrollbar(both screen edges)
		if currX > startX+(maxX/10) || currX < startX-(maxX/10) || currX > maxX-(maxX/20) || currX < (maxX/20) {
			checkChan <- false
			break
		}
		// fail if multi-touch, swipe up, or over 2 sec(timeout)
		if tempState.isDown[1] || currY < startY - 15 || time.Now().Sub(startTime) > 2*time.Second {
			checkChan <- false
			break
		}
		// pass if reaches down far enough without previous fails
		if currY >= startY + (subSwipeLength * 2) {
			checkChan <- true
			break
		}
	}
	return
}

func downCheckThree (checkChan chan bool, StateChan chan mpts, startX int32, startY int32) {
	startTime := time.Now()
	var swipeLength = halfY
	//var subSwipeLength = swipeLength / 3
	for {
		tempState := <-StateChan
		currX := int32(tempState.X[0])
		currY := int32(tempState.Y[0])
		// fail if angled or over scrollbar(both screen edges)
		if currX > startX+(maxX/10) || currX < startX-(maxX/10) || currX > maxX-(maxX/20) || currX < (maxX/20) {
			checkChan <- false
			break
		}
		// fail if multi-touch, swipe up, or over 2 sec(timeout)
		if tempState.isDown[1] || currY < startY - 15 || time.Now().Sub(startTime) > 2*time.Second {
			checkChan <- false
			break
		}
		// pass if reaches down far enough without previous fails
		if currY >= startY + swipeLength {
			checkChan <- true
			break
		}
	}
	return
}

func gestureSwipeDown1 (StateChan chan mpts) {
	var startX, startY int32
	var lastTap int32
	var isStart bool
	checkChannel := make(chan bool)
	for {
		tempState := <-StateChan
		if !isLocked && !isStart && tempState.tapCnt > lastTap {
			if tempState.isGesture {
				isStart = true
				startX = int32(tempState.X[0])
				startY = int32(tempState.Y[0])
				lastTap = tempState.tapCnt
				go downCheckOne(checkChannel, StateChan, startX, startY)
				go downCheckTwo(checkChannel, StateChan, startX, startY)
				go downCheckThree(checkChannel, StateChan, startX, startY)
				passOne := <- checkChannel
				passTwo := <- checkChannel
				passThree := <- checkChannel
				if passOne && passTwo && passThree {
					switch touchState.rotation {
					case inv:
						fmt.Println("UP SWIPE!")
						fnKeys := []string{"lctrl", "lalt"}
						robotgo.KeyTap("up", fnKeys)
					case norm:
						fmt.Println("DOWN SWIPE!")
						fnKeys := []string{"lctrl", "lalt"}
						robotgo.KeyTap("down", fnKeys)
					}
				}
				isStart = false
			}
		}
	}
}

func upCheckOne (checkChan chan bool, StateChan chan mpts, startX int32, startY int32) {
	startTime := time.Now()
	var swipeLength = halfY
	var subSwipeLength = swipeLength / 3
	for {
		tempState := <-StateChan
		currX := int32(tempState.X[0])
		currY := int32(tempState.Y[0])
		// fail if angled or over scrollbar(both screen edges)
		if currX > startX+(maxX/10) || currX < startX-(maxX/10) || currX > maxX-(maxX/20) || currX < (maxX/20) {
			checkChan <- false
			break
		}
		// fail if multi-touch, swipe down, or over 2 sec(timeout)
		if tempState.isDown[1] || currY > startY + 15 || time.Now().Sub(startTime) > 2*time.Second {
			checkChan <- false
			break
		}
		// pass if reaches down far enough without previous fails
		if currY <= startY - subSwipeLength {
			checkChan <- true
			break
		}
	}
	return
}

func upCheckTwo (checkChan chan bool, StateChan chan mpts, startX int32, startY int32) {
	startTime := time.Now()
	var swipeLength = halfY
	var subSwipeLength = swipeLength / 3
	for {
		tempState := <-StateChan
		currX := int32(tempState.X[0])
		currY := int32(tempState.Y[0])
		// fail if angled or over scrollbar(both screen edges)
		if currX > startX+(maxX/10) || currX < startX-(maxX/10) || currX > maxX-(maxX/20) || currX < (maxX/20) {
			checkChan <- false
			break
		}
		// fail if multi-touch, swipe down, or over 2 sec(timeout)
		if tempState.isDown[1] || currY > startY + 15 || time.Now().Sub(startTime) > 2*time.Second {
			checkChan <- false
			break
		}
		// pass if reaches down far enough without previous fails
		if currY <= startY - (subSwipeLength * 2) {
			checkChan <- true
			break
		}
	}
	return
}

func upCheckThree (checkChan chan bool, StateChan chan mpts, startX int32, startY int32) {
	startTime := time.Now()
	var swipeLength = halfY
	//var subSwipeLength = swipeLength / 3
	for {
		tempState := <-StateChan
		currX := int32(tempState.X[0])
		currY := int32(tempState.Y[0])
		// fail if angled or over scrollbar(both screen edges)
		if currX > startX+(maxX/10) || currX < startX-(maxX/10) || currX > maxX-(maxX/20) || currX < (maxX/20) {
			checkChan <- false
			break
		}
		// fail if multi-touch, swipe down, or over 2 sec(timeout)
		if tempState.isDown[1] || currY > startY + 15 || time.Now().Sub(startTime) > 2*time.Second {
			checkChan <- false
			break
		}
		// pass if reaches down far enough without previous fails
		if currY <= startY - swipeLength {
			checkChan <- true
			break
		}
	}
	return
}

func gestureSwipeUp1 (StateChan chan mpts) {
	var startX, startY int32
	var lastTap int32
	var isStart bool
	checkChannel := make(chan bool)
	for {
		tempState := <-StateChan
		if !isLocked && !isStart && tempState.tapCnt > lastTap {
			//if tempState.isGesture {
				isStart = true
				startX = int32(tempState.X[0])
				startY = int32(tempState.Y[0])
				lastTap = tempState.tapCnt
				go upCheckOne(checkChannel, StateChan, startX, startY)
				go upCheckTwo(checkChannel, StateChan, startX, startY)
				go upCheckThree(checkChannel, StateChan, startX, startY)
				passOne := <- checkChannel
				passTwo := <- checkChannel
				passThree := <- checkChannel
				if passOne && passTwo && passThree {
					switch touchState.rotation {
					case inv:
						fmt.Println("DOWN SWIPE!")
						fnKeys := []string{"lctrl", "lalt"}
						robotgo.KeyTap("down", fnKeys)
					case norm:
						fmt.Println("UP SWIPE!")
						fnKeys := []string{"lctrl", "lalt"}
						robotgo.KeyTap("up", fnKeys)
					}
				}
				isStart = false
			//}
		}
	}
}

// Old format code, will be converted to mpts code like the above f(x)
// func isSwipeRL1(EventChan chan *evdev.InputEvent, RightKeyboard keybd_event.KeyBonding, LeftKeyboard keybd_event.KeyBonding) {
// 	var StartX, StartY int32 = -1, -1
// 	var CurrX, CurrY int32
// 	var SwipeLength = (maxX + 1) / 3
// 	var SubSwipeLength = SwipeLength / 3
// 	var TouchStart = false
// 	//var TapCnt int32 //currently unused but update value with event.
// 	var SwipeRtCheck1, SwipeRtCheck2, SwipeRtCheck3 = false, false, false
// 	var SwipeLtCheck1, SwipeLtCheck2, SwipeLtCheck3 = false, false, false
// 	var StartTime time.Time
// 	for {
// 		event := <-EventChan
// 		if isEventMatch(event, touchDnEv) {
// 			StartTime = time.Now()
// 			TouchStart = true
// 			//TapCnt = event.Value
// 		} else if TouchStart == true {
// 			if isEventMatch(event, touchXEv) {
// 				//TapCnt = event.Value
// 				if StartX != -1 {
// 					CurrX = event.Value
// 					if CurrX >= StartX+SubSwipeLength {
// 						SwipeRtCheck1 = true
// 					} else if CurrX <= StartX-SubSwipeLength {
// 						SwipeLtCheck1 = true
// 					}
// 					if CurrX >= StartX+(SubSwipeLength*2) {
// 						SwipeRtCheck2 = true
// 					} else if CurrX <= StartX-(SubSwipeLength*2) {
// 						SwipeLtCheck2 = true
// 					}
// 					if CurrX >= StartX+SwipeLength {
// 						SwipeRtCheck3 = true
// 					} else if CurrX <= StartX-SwipeLength {
// 						SwipeLtCheck3 = true
// 					}
// 				} else {
// 					StartX = event.Value
// 				}
// 			} else if isEventMatch(event, touchYEv) {
// 				if StartY != -1 {
// 					CurrY = event.Value
// 					if CurrY > StartY+(maxY/10) || CurrY < StartY-(maxY/10) {
// 						//Out of tolerance relook for new touch
// 						TouchStart = false
// 						SwipeRtCheck1, SwipeRtCheck2, SwipeRtCheck3 = false, false, false
// 						SwipeLtCheck1, SwipeLtCheck2, SwipeLtCheck3 = false, false, false
// 						StartX, StartY = -1, -1
// 					}
// 				} else {
// 					StartY = event.Value
// 				}
// 			} else if isEventMatch(event, touchUpEv) {
// 				if time.Now().Sub(StartTime) < 2*time.Second {
// 					if SwipeRtCheck1 == true && SwipeRtCheck2 == true && SwipeRtCheck3 == true {
// 						switch touchState.rotation {
// 						case inv:
// 							fmt.Println("LEFT SWIPE!")
// 							execKeys(LeftKeyboard)
// 						case norm:
// 							fmt.Println("RIGHT SWIPE!")
// 							execKeys(RightKeyboard)
// 						}
// 					} else if SwipeLtCheck1 == true && SwipeLtCheck2 == true && SwipeLtCheck3 == true {
// 						switch touchState.rotation {
// 						case inv:
// 							fmt.Println("RIGHT SWIPE!")
// 							execKeys(RightKeyboard)
// 						case norm:
// 							fmt.Println("LEFT SWIPE!")
// 							execKeys(LeftKeyboard)
// 						}
// 					}
// 				}
// 				TouchStart = false
// 				SwipeRtCheck1, SwipeRtCheck2, SwipeRtCheck3 = false, false, false
// 				SwipeLtCheck1, SwipeLtCheck2, SwipeLtCheck3 = false, false, false
// 				StartX, StartY = -1, -1
// 			}
// 		}
// 	}
// }
