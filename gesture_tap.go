package main

import (
	"fmt"
	"os/exec"
	"time"
	"github.com/go-vgo/robotgo"
)

func gestureTap(StateChan chan mpts) {
	var isStart bool
	var touchTime time.Time
	var touchCnt int
	var touchMax int
	for {
		tempState := <-StateChan
		switch isStart {
		case false:
			if tempState.isGesture {
				isStart = true
				touchTime = time.Now()
				continue
			}
		case true:
			for i := 0; i < 10; i++ {
				if tempState.isDown[i] {
					touchCnt++
				}
			}
			if touchCnt > touchMax {
				touchMax = touchCnt
			}
			touchCnt = 0
			timeDiff := time.Now().Sub(touchTime)
			if !tempState.isGesture {
				isStart = false
				touchCnt = 0
				if timeDiff < 600*time.Millisecond {
					switch touchMax {
					// case 2:
						// 	fmt.Println("TWO FINGER TAP!")
						// 	cmdErr := exec.Command("onboard").Run()
						// 	parseFatal(cmdErr, "2 finger tap action failure")
						case 3:
							fmt.Println("THREE FINGER TAP!")
							if !isLocked {
								robotgo.KeyTap("f12")
							}
						case 4:
							fmt.Println("FOUR FINGER TAP!")
							if !isLocked {
								cmdErr := exec.Command("onboard").Run()
								parseFatal(cmdErr, "5 finger tap action failure")
							}
						case 5:
							fmt.Println("FIVE FINGER TAP!")
							if !isLocked {
								prevX, prevY := robotgo.GetMousePos()
								switch touchState.rotation {
									case inv:
										xval := float32(7615) / (float32((maxX + 1)) / float32(resolX))
										xinv := halfResolX + (halfResolX - int32(xval))
										yval := float32(60) / (float32((maxY + 1)) / float32(resolY))
										yinv := halfResolY + (halfResolY - int32(yval))
										robotgo.MoveMouse(int(xinv), int(yinv))
									case norm:
										xval := float32(7615) / (float32((maxX + 1)) / float32(resolX))
										yval := float32(60) / (float32((maxY + 1)) / float32(resolY))
										robotgo.MoveMouse(int(xval), int(yval))
								}
								time.Sleep(15 * time.Millisecond)
								mouse.LeftPress()
								time.Sleep(15 * time.Millisecond)
								mouse.LeftRelease()
								time.Sleep(50 * time.Millisecond)
								robotgo.MoveMouse(prevX, prevY)
							}
						case 6:
						// fmt.Println("SIX FINGER TAP!")
						// 	cmdErr := exec.Command("onboard").Run()
						// 	parseFatal(cmdErr, "6 finger tap action failure")
						case 7:
						//	fmt.Println("SEVEN FINGER TAP!")
						// 	cmdErr := exec.Command("onboard").Run()
						// 	parseFatal(cmdErr, "7 finger tap action failure")
						case 8:
						// 	fmt.Println("EIGHT FINGER TAP!")
						// 	cmdErr := exec.Command("onboard").Run()
						// 	parseFatal(cmdErr, "8 finger tap action failure")
						case 9:
							fmt.Println("NINE FINGER TAP!")
							screenLocker()
						case 10:
					 		fmt.Println("10 FINGER TAP!")
							screenLocker()
					}
				}
				touchMax = 0
			}
		}
	}
}
