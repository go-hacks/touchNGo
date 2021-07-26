# touchNGo
Complete Linux Touchscreen Control Replacement

Current features

Full mouse controls(left/right clicks, drag-select, etc)

(to use this software effectively, you must know how the clicks work)

(left click and drag works as you would think, however when you let up
after holding for about a second, you are instead prompted with a right
click on release instead of just a released left click as you would think)

(this is especially useful for drag selecting a bunch of files and immediately
having a right click when you release for a whole saved click)

(it is quite slick when you get used to it)

1 Finger swipe (up/down)(left/right in upgrades, unusable)

(obviously some minor conflicts arise while using 1 finger swiping. it is
pretty much proof of concept and does work fine when you're used to where
you can swipe and not cause issues. naturally, 2+ finger swiping will come
in the future)

2-10 Finger tap gestures (3 & 5 configured to F12 for yakuake and onboard keyboard)

Auto-updates screen rotation & resolution

Supports inverted touchscreen (left/right rotation yet to come)

Supports touchscreen locking via vol buttons (Dn/Up 2x)

^this can be finicky depending on your buttons

All active functions now use MPTS (multi-point touchstate)

To use this, you will currently have to modify some sections of
code for your setup but I will point you in the right directions.

If you want a more complete description of what this is, see touchngo.go

git clone https://github.com/go-hacks/touchNGo.git
to your home directory. (This is important!)

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
