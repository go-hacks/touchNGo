package main

import (
	"fmt"
	"syscall"
	// For reading touchscreen events
	"github.com/gvalkov/golang-evdev"
)

// Handles screen locking with vol dn/up side buttons
func screenLock() {
	var buttons [8]tevMatch
	var times [8]syscall.Timeval
	device, _ := evdev.Open("/dev/input/event3")
	for {
		eventFrame, _ := device.Read()
		for i := 0; i < len(eventFrame); i++ {
			event := eventFrame[i]
			if event.Code == 114 || event.Code == 115 {
				buttons = btnShift(buttons)
				times = timeShift(times)
				times[0] = event.Time
				buttons[0].Code = event.Code
				buttons[0].Type = event.Type
				buttons[0].Value = event.Value
				buttons[0].Ineq = "eq"
				if buttons == lockMatch && times[0].Sec-times[7].Sec < 4 {
					if isLocked == false {
						isLocked = true
						fmt.Println("SCREEN LOCKED!")
					} else {
						isLocked = false
						fmt.Println("SCREEN UNLOCKED!")
					}
				}
			}
		}
	}
}

func btnShift(arr [8]tevMatch) [8]tevMatch {
	var arrNew [8]tevMatch
	for i := 0; i < 7; i++ {
		arrNew[i+1] = arr[i]
	}
	return arrNew
}

func timeShift(arr [8]syscall.Timeval) [8]syscall.Timeval {
	var arrNew [8]syscall.Timeval
	for i := 0; i < 7; i++ {
		arrNew[i+1] = arr[i]
	}
	return arrNew
}
