package main

import (
	"fmt"
	"time"
	// For executing keyboard actions
	"github.com/micmonay/keybd_event"
)

// Rotational direction strings for cases
const inv string = "inverted"
const norm string = "normal"
const left string = "left"
const right string = "right"

// Define multi-point touchstate and
// touch event match structures
type mpts struct {
	tapCnt    int32
	isGesture bool
	isDown    [10]bool
	X, Y      [10]uint16
	rotation  string
}
type tevMatch struct {
	Code, Type uint16
	Ineq       string
	Value      int32
}

// Declare states and event matches globally
var touchState mpts
var blankState mpts

//Initialize match events
var touchXEv = tevMatch{
	Code: 53, Type: 3,
	Ineq: "ge", Value: 0,
}
var touchYEv = tevMatch{
	Code: 54, Type: 3,
	Ineq: "ge", Value: 0,
}
var touchDnEv = tevMatch{
	Code: 57, Type: 3,
	Ineq: "ge", Value: 0,
}
var touchUpEv = tevMatch{
	Code: 57, Type: 3,
	Ineq: "eq", Value: -1,
}
var touchIsOvr1 = tevMatch{
	Code: 47, Type: 3,
	Ineq: "ge", Value: 1,
}
var touchIsOvr2 = tevMatch{
	Code: 47, Type: 3,
	Ineq: "ge", Value: 2,
}
var touchIs4 = tevMatch{
	Code: 47, Type: 3,
	Ineq: "eq", Value: 3,
}
var touchIs3 = tevMatch{
	Code: 47, Type: 3,
	Ineq: "eq", Value: 2,
}
var touchIs5 = tevMatch{
	Code: 47, Type: 3,
	Ineq: "eq", Value: 4,
}
var touchOvr3 = tevMatch{
	Code: 47, Type: 3,
	Ineq: "gt", Value: 2,
}
var touchOvr5 = tevMatch{
	Code: 47, Type: 3,
	Ineq: "gt", Value: 4,
}
var gestureStart = tevMatch{
	Code: 330, Type: 1,
	Ineq: "eq", Value: 1,
}
var gestureStop = tevMatch{
	Code: 330, Type: 1,
	Ineq: "eq", Value: 0,
}

var volDnPress = tevMatch{
	Code: 114, Type: 1,
	Ineq: "eq", Value: 1,
}
var volDnRelease = tevMatch{
	Code: 114, Type: 1,
	Ineq: "eq", Value: 0,
}
var volUpPress = tevMatch{
	Code: 115, Type: 1,
	Ineq: "eq", Value: 1,
}
var volUpRelease = tevMatch{
	Code: 115, Type: 1,
	Ineq: "eq", Value: 0,
}
var lockMatch [8]tevMatch

/*var resolX int32 = 1920
var resolY int32 = 1280
var halfResolX int32 = resolX / 2
var halfResolY int32 = resolY / 2
*/
var isSwipeRL1Rkb keybd_event.KeyBonding
var isSwipeRL1Lkb keybd_event.KeyBonding
var isSwipeUD1Ukb keybd_event.KeyBonding
var isSwipeUD1Dkb keybd_event.KeyBonding
var threeTapKeyboard keybd_event.KeyBonding

func init() {
	fmt.Printf("Initializing...")
	// Build match array for screen lock/unlock
	for i := 0; i < 2; i++ {
		j := i*4
		lockMatch[j] = volUpRelease
		lockMatch[j+1] = volUpPress
		lockMatch[j+2] = volDnRelease
		lockMatch[j+3] = volDnPress
	}
	// Initialize Keyboards
	var err error
	isSwipeRL1Rkb, err = keybd_event.NewKeyBonding()
	parseFatal(err, "RL1Rkb failure")
	isSwipeRL1Lkb, err = keybd_event.NewKeyBonding()
	parseFatal(err, "RL1Lkb failure")
	isSwipeUD1Ukb, err = keybd_event.NewKeyBonding()
	parseFatal(err, "UD1Ukb failure")
	isSwipeUD1Dkb, err = keybd_event.NewKeyBonding()
	parseFatal(err, "UD1Dkb failure")
	threeTapKeyboard, err = keybd_event.NewKeyBonding()
	parseFatal(err, "3tapkb failure")
	// For linux, wait 2 seconds **why tho? cheat to 1 sec.
	time.Sleep(1 * time.Second)
	// Set keyboard keybindings
	isSwipeRL1Rkb.HasCTRL(true)
	isSwipeRL1Rkb.HasALT(true)
	isSwipeRL1Rkb.SetKeys(keybd_event.VK_DOWN)
	isSwipeRL1Lkb.HasCTRL(true)
	isSwipeRL1Lkb.HasALT(true)
	isSwipeRL1Lkb.SetKeys(keybd_event.VK_UP)
	isSwipeUD1Ukb.HasCTRL(true)
	isSwipeUD1Ukb.HasALT(true)
	isSwipeUD1Ukb.SetKeys(keybd_event.VK_UP)
	isSwipeUD1Dkb.HasCTRL(true)
	isSwipeUD1Dkb.HasALT(true)
	isSwipeUD1Dkb.SetKeys(keybd_event.VK_DOWN)
	threeTapKeyboard.SetKeys(keybd_event.VK_F12)
}
