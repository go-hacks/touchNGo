package main

/*
              +---Version 0.52---+
<---Complete Linux Touchscreen Control Replacement--->
	!No OS or application specific touch
	!controls will work while in use.
	!Disables touchscreen via xinput while
        !in use and re-enables upon exiting.
	*In theory, the main controls could
	*be disabled along with leaving xinput
	*running to achieve a "supplemental" mode
	*for only adding gestures to stock Linux.
	*You could also have this on a dock to
	*disable while using a specific application
	*with special touch controls in which
	*case TouchNGo reverts back to stock
	*controls upon exiting or being aborted.
    <<<--------------- Features --------------->>>
        Full mouse controls(left/right clicks, drag-select, etc)
        1 Finger swipe (up/down/left/right)
        2-10 Finger tap gestures
	Auto-updates screen rotation & resolution
	Supports inverted touchscreen (left/right rotation yet to come)
	Supports touchscreen locking via vol buttons (Dn/Up 2x)
	All active functions now use MPTS (multi-point touchstate)
*/

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
	// For reading touchscreen events
	"github.com/gvalkov/golang-evdev"
	// For executing mouse clicks
	"github.com/bendahl/uinput"
)

// Set global values
var maxX int32 = 0
var maxY int32 = 0
var isLocked bool = false
var resolX int32
var resolY int32
var halfResolX int32
var halfResolY int32
var halfX, halfY int32
var calibrated bool
var devPath string
var isDebugMode bool
var mouse uinput.Mouse

func main() {
	// Create virtual mouse
	mouse, _ = uinput.CreateMouse("/dev/uinput", []byte("touchNGoMouse"))
	defer mouse.Close()
	// Check Input Sanity
	if len(os.Args) < 2 {
		fmt.Println("\rList Devices -l")
		fmt.Println("Usage:", os.Args[0], "touchDevice (DEBUG)")
		os.Exit(0)
	}
	// Check for debugging mode
	if len(os.Args) == 3 {
		if os.Args[2] == "DEBUG" {
			isDebugMode = true
		}
	}
	// Check for calibration
	if maxX > 0 && maxY > 0 {
		calibrated = true
	}
	// Get touch event paths or list devices
	devices, err := evdev.ListInputDevices()
	parseFatal(err, "Fail to get input device list")
	for _, dev := range devices {
		if os.Args[1] == "-l" {
			fmt.Println(dev.Name)
		} else if os.Args[1] == dev.Name {
			devPath = dev.Fn
			break
		}
	}
	if os.Args[1] == "-l" {
		os.Exit(0)
	}
	// Start Ctrl+C & sigterm hook to re-enable
	// xinput before disabling the touchscreen
	SigChan := make(chan os.Signal, 1)
	signal.Notify(SigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-SigChan
		cmd := "xinput"
		arg := "--enable"
		cmdErr := exec.Command(cmd, arg, "pointer:"+os.Args[1]).Run()
		parseFatal(cmdErr, "Touchscreen re-enable failure")
		os.Exit(1)
	}()
	// Disable touchscreen via xinput as we
	// will now be hijacking the events :)
	cmd := "xinput"
	arg := "--disable"
	cmdErr := exec.Command(cmd, arg, "pointer:"+os.Args[1]).Run()
	parseFatal(cmdErr, "Touchscreen disable failure")
	// Get screen information
	go screenDataTicker()
	// Give screenDataTicker() a chance to
	// update once before sending frames.
	time.Sleep(250 * time.Millisecond)
	fmt.Printf("READY!\n")
	// Do calibration if needed and set HalfX/Y
	if calibrated == false {
		fmt.Println("SCREEN BOUNDARIES ARE UNSET!")
		fmt.Println("Slowly swipe finger from center to edges until values max out.")
		fmt.Println("To finish calibration, tap the screen with at least 4 fingers.")
		fmt.Println("Calibrating Now...")
		calibrated = calibrateEdges()
		fmt.Println("You may set these values above main() to avoid calibration in the future.")
	}
	halfX = (maxX + 1) / 2
	halfY = (maxY + 1) / 2
	// Initiate event channel (buffered for less frame loss)
	// for the analyzer and start main Go routine
	frameAnalyzerEVCH := make(chan []evdev.InputEvent, 5)
	go frameAnalyzer(frameAnalyzerEVCH)
	// Run screen lock routine to handle touch locking w/vol buttons
	go screenLock()
	// Setup touchscreen event device
	// and begin sending frames to analyzer
	device, _ := evdev.Open(devPath)
	for {
		eventFrame, _ := device.Read()
		if isDebugMode {
			fmt.Println(eventFrame)
		}
		frameAnalyzerEVCH <- eventFrame
	}
}
