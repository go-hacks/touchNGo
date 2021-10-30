package main

import (
	"fmt"
)

// Set global values
var maxX int32 = 7679
var maxY int32 = 4319
var isLocked bool = false
var resolX int32
var resolY int32
var halfResolX int32
var halfResolY int32
var halfX, halfY int32
var calibrated bool
var devPath string
var isDebugMode bool
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
//var lockMatch [8]tevMatch

func init() {
	fmt.Printf("Initializing...")

	// Build match array for screen lock/unlock
	// for i := 0; i < 2; i++ {
	// 	j := i*4
	// 	lockMatch[j] = volUpRelease
	// 	lockMatch[j+1] = volUpPress
	// 	lockMatch[j+2] = volDnRelease
	// 	lockMatch[j+3] = volDnPress
	// }
	// ^^deprecated
}
