package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
	// For reading touchscreen events
	"github.com/gvalkov/golang-evdev"
	// For executing keyboard actions
	"github.com/micmonay/keybd_event"
)

type screenData struct {
	rotation string
	resolX   int32
	resolY   int32
}

func isEventMatch(event *evdev.InputEvent, matchEvent tevMatch) bool {
	switch matchEvent.Ineq {
	case "eq":
		if event.Code == matchEvent.Code && event.Type == matchEvent.Type && event.Value == matchEvent.Value {
			return true
		}
	case "gt":
		if event.Code == matchEvent.Code && event.Type == matchEvent.Type && event.Value > matchEvent.Value {
			return true
		}
	case "ge":
		if event.Code == matchEvent.Code && event.Type == matchEvent.Type && event.Value >= matchEvent.Value {
			return true
		}
	case "lt":
		if event.Code == matchEvent.Code && event.Type == matchEvent.Type && event.Value < matchEvent.Value {
			return true
		}
	case "le":
		if event.Code == matchEvent.Code && event.Type == matchEvent.Type && event.Value <= matchEvent.Value {
			return true
		}
	}
	return false
}

func execKeys(keyboard keybd_event.KeyBonding) {
	keyboard.Press()
	time.Sleep(10 * time.Millisecond)
	keyboard.Release()
}

func getScreenSpecs() screenData {
	var sd screenData
	xrcmd := "xrandr --current --verbose|grep primary"
	output, _ := exec.Command("bash", "-c", xrcmd).Output()
	if len(output) == 0 {
		return sd
	}
	strAry := strings.Split(string(output), " ")
	resolAry := strings.Split(strAry[3], "x")
	intX, _ := strconv.Atoi(resolAry[0])
	sd.resolX = int32(intX)
	yAry := strings.Split(resolAry[1], "+")
	intY, _ := strconv.Atoi(yAry[0])
	sd.resolY = int32(intY)
	sd.rotation = strAry[5]
	return sd
}

func screenDataTicker() {
	for {
		screenInfo := getScreenSpecs()
		if screenInfo.rotation != "" {
			resolX = screenInfo.resolX
			resolY = screenInfo.resolY
			halfResolX = resolX / 2
			halfResolY = resolY / 2
			touchState.rotation = screenInfo.rotation
		}
		time.Sleep(3 * time.Second)
	}
}

func resetTouchState() {
	rotation := touchState.rotation
	touchState = blankState
	touchState.rotation = rotation
}

func parseFatal(err error, msg string) {
	if err != nil {
		if msg != "" {
			fmt.Println(msg)
		}
		fmt.Println(err)
		os.Exit(0)
	} else {
		return
	}
}
