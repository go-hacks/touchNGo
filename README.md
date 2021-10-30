# touchNGo
Complete Linux Touchscreen Control Replacement

To use this, you will currently have to modify some sections of
code for your setup but I will point you in the right directions.

If you want a more complete description of what this is, see touchngo.go

git clone https://github.com/go-hacks/touchNGo.git
to your home directory. (This is important!)

First you will want to build it by running the included build script ./build

Next, run ./touchNGo -l to get the list of input devices.

Then, put your touchscreen device's name in the included launch script touchStart

Replace mine, GXTP7386:00 27C6:0113, with whatever yours is.

Run ./touchStart and you should be prompted with calibration.
Take your calibration maxX and maxY values and put them in init.go on lines 8 & 9.

Rebuild with ./build and re-run with ./touchStart.

If you have panel launchers, you may set them to the included flip-screen.pl which will
allow for a button to invert the touchscreen and back again and so on.

Note: Make sure all executables are marked as such. touchNGo, touchStart,
flip-screen.pl, rot.sh, and the build script may all require chmod +x

To set the specific actions for gestures, you will find them in
gesture_tap.go at line 40 where the case begins and gesture_swipe1.go
at lines 111 & 229 where the cases begin there as well.

You can see what was done and can probably mix n match keyboard calls and system cmds
into any of the gestures. Eventually this will all be controlled via a JSON config file
but for now I have it hard coded for my setup with the other options commented out.
Swipe R/L function is disabled for now as I need to finish rewriting it into the new
MPTS (Multi-Point Touch State) format like the rest of the code uses.

I have changed the screen locking function to 9 & 10 finger tap instead of using the
volume buttons as it works better and doesn't require the use of a 2nd event device.
It's set to 9 & 10 because it can be tricky to get all 10 taps to register every time
and it makes the whole process feel smoother and more responsive.

If you prefer to use the volume buttons, I left the module intact and you can remove
the .dpcd from its name and uncomment the subroutine in touchngo.go and the init
stuff for it in init.go. It is pretty straightforward.
