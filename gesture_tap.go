package main

import (
	"fmt"
	"os/exec"
	"time"
	"github.com/micmonay/keybd_event"
)

func isGestureTap(StateChan chan mpts, threeTapkb keybd_event.KeyBonding) {
	var isStart bool
	var touchTime time.Time
	var touchCnt int
	var touchMax int
	for {
		tempState := <-StateChan
		if !isLocked && !isStart {
			if tempState.isGesture {
				isStart = true
				touchTime = time.Now()
				continue
			}
		}
		if !isLocked && isStart {
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
						execKeys(threeTapkb)
					// case 4:
						// 	fmt.Println("FOUR FINGER TAP!")
						// 	cmdErr := exec.Command("onboard").Run()
						// 	parseFatal(cmdErr, "4 finger tap action failure")
					case 5:
						fmt.Println("FIVE FINGER TAP!")
						cmdErr := exec.Command("onboard").Run()
						parseFatal(cmdErr, "5 finger tap action failure")
					// case 6:
						// 	fmt.Println("SIX FINGER TAP!")
						// 	cmdErr := exec.Command("onboard").Run()
						// 	parseFatal(cmdErr, "6 finger tap action failure")
					// case 7:
						// 	fmt.Println("SEVEN FINGER TAP!")
						// 	cmdErr := exec.Command("onboard").Run()
						// 	parseFatal(cmdErr, "7 finger tap action failure")
					// case 8:
						// 	fmt.Println("EIGHT FINGER TAP!")
						// 	cmdErr := exec.Command("onboard").Run()
						// 	parseFatal(cmdErr, "8 finger tap action failure")
					// case 9:
						// 	fmt.Println("NINE FINGER TAP!")
						// 	cmdErr := exec.Command("onboard").Run()
						// 	parseFatal(cmdErr, "9 finger tap action failure")
					// case 10:
						// 	fmt.Println("10 FINGER TAP!")
						// 	cmdErr := exec.Command("onboard").Run()
						// 	parseFatal(cmdErr, "10 finger tap action failure")
					}
				}
				touchMax = 0
			}
		}
	}
}
