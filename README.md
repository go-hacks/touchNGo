# touchNGo
Complete Linux Touchscreen Control Replacement

To use this, you will currently have to modify some sections of
code for your setup but I will point you in the right directions.

If you want a more complete description of what this is, see touchngo.go

First you will want to build it by running the included build script ./build
You may require these go get dependencies
  For reading touchscreen events
	"github.com/gvalkov/golang-evdev"
	For executing mouse clicks
	"github.com/bendahl/uinput"
  For executing keyboard actions
  "github.com/micmonay/keybd_event"

Next, run ./touchNGo -l to get the list of input devices.
Then, put your touchscreen devices name in the included launch script touchStart
Replace mine, GXTP7386:00 27C6:0113, with whatever yours is.
Run ./touchStart and you should be prompted with calibration.
Take your calibration maxX and maxY values and put them in touchngo.go on lines 43 & 44.
Rebuild with ./build and re-run with ./touchStart.
If you have panel launchers, you may set them to the included flip-screen.pl which will
allow for a button to invert the touchscreen and back again and so on.

Note: Make sure all executables are marked as such. touchNGo, touchStart,
flip-screen.pl, and rot.sh may all require chmod +x

To set the specific actions for gestures, you will find them in gesture_tap.go line 40
and init.go line 149 for the up/down swipe keyboards.

You can see what was done and can probably mix n match keyboard calls and system cmds
into any of the gestures. Eventually this will all be controlled via a config file
but for now I have it hard coded for my setup with the other options commented out.
Swipe R/L function does work but I need to finish rewriting it into the new MPTS
format that I just finished before uploading all of this.
