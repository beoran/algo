package al

/*
#include <stdlib.h>
#include <allegro5/allegro.h>
#include "helpers.h"
*/
import "C"


/* Enum: ALLEGRO_TOUCH_INPUT_MAX_TOUCH_COUNT
 */
const TOUCH_INPUT_MAX_TOUCH_COUNT = C.ALLEGRO_TOUCH_INPUT_MAX_TOUCH_COUNT

type TouchInput struct {
    handle * C.ALLEGRO_TOUCH_INPUT
}

func wrapTouchInputRaw(cti * C.ALLEGRO_TOUCH_INPUT) * TouchInput {
    res := &TouchInput{cti}
    return res
}

type TouchInputState  C.ALLEGRO_TOUCH_INPUT_STATE

func wrapTouchInputStateRaw(cti C.ALLEGRO_TOUCH_INPUT_STATE) TouchInputState {
    return TouchInputState(cti)
}

type TouchState C.ALLEGRO_TOUCH_STATE

func wrapTouchStateRaw(ts C.ALLEGRO_TOUCH_STATE) TouchState {
    return TouchState(ts)
}

func (ts TouchState) Id() int {
        return int(ts.id)
}

func (ts TouchState) X() float32 { return float32(ts.x); }
func (ts TouchState) Y() float32 { return float32(ts.y); }
func (ts TouchState) DX() float32 { return float32(ts.dx); }
func (ts TouchState) DY() float32 { return float32(ts.dy); }
func (ts TouchState) Primary() bool { return bool(ts.primary); }
func (ts TouchState) Display() * Display { return wrapDisplayRaw(ts.display); }

func (tsi TouchInputState) Touch(index int) TouchState {
    return wrapTouchStateRaw(tsi.touches[index])
} 
 
func IsTouchInputInstalled() bool {
    return bool(C.al_is_touch_input_installed())
}

func InstallTouchInput() bool {
    return bool(C.al_install_touch_input())
}

func UninstallTouchInput() { 
    C.al_uninstall_touch_input()
}

func GetTouchInputState() TouchInputState {
    cstate := C.struct_ALLEGRO_TOUCH_INPUT_STATE{}
    C.al_get_touch_input_state(&cstate)
    return wrapTouchInputStateRaw(cstate)
}

func GetTouchInputEventSource() * EventSource {
    return wrapEventSourceRaw(C.al_get_touch_input_event_source())
}

