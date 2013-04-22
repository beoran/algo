package al

/*
#include <stdlib.h>
#include <allegro5/allegro.h>
#include "helpers.h"
*/
import "C"

import "runtime"

// Usful regexp for KATE:  ALLEGRO_([A-Z0-9_]+)(.*) -> \1 = C.ALLEGRO_\1

// Joystick functionality.

type Joystick struct {
    handle *C.ALLEGRO_JOYSTICK
}

// Destroyst he joystick. According to the Allegro documentation, this 
// does nothing.
func (self *Joystick) Destroy() {
    // Some problems is this is enabled so make sure this does nothing...
    // C.al_release_joystick(self.handle)
}

// Wraps a C joystick handler into the Go Joystick wrapper.
func wrapJoystickRaw(handle *C.ALLEGRO_JOYSTICK) *Joystick {
    if handle == nil {
        return nil
    }
    return &Joystick{handle}
}

// Wraps a C joystick handler into the Go Joystick wrapper.
// Also sets a finalizer that calls joystick.Destroy().
func wrapJoystick(handle *C.ALLEGRO_JOYSTICK) *Joystick {
    self := wrapJoystickRaw(handle)
    if self != nil {
        runtime.SetFinalizer(self, func(me *Joystick) { me.Destroy() })
    }
    return self
}

// Struct that holds the state of the joystick.
type JoystickState C.ALLEGRO_JOYSTICK_STATE

// Converts a wrapped joystick state to a C joystick state.
func (self *JoystickState) toC() *C.ALLEGRO_JOYSTICK_STATE {
    return (*C.ALLEGRO_JOYSTICK_STATE)(self)
}

// Gets the state of the axis for the given stick on the joystick.
// returns 0.0 if the stick or axis are out of range. May return
// garbage for nonexisting sticks and axes.
func (self *JoystickState) GetStickAxis(stick, axis int) float32 {
    if stick >= int(C._AL_MAX_JOYSTICK_STICKS) {
        return 0.0
    }
    if axis >= int(C._AL_MAX_JOYSTICK_AXES) {
        return 0.0
    }
    if axis < 0 {
        return 0.0
    }
    if stick < 0 {
        return 0.0
    }
    return float32(self.stick[C.int(stick)].axis[C.int(axis)])
}

// Gerts the state of the button with the given index on the joystick.
// Will return -1 if the button is out of range.
func (self *JoystickState) GetButton(button int) int {
    if button >= int(C._AL_MAX_JOYSTICK_BUTTONS) {
        return -1
    }
    if button < 0 {
        return -1
    }
    return int(self.button[C.int(button)])
}

// Joystick flags that determine the type of the joystick.
const (
    JOYFLAG_DIGITAL  = C.ALLEGRO_JOYFLAG_DIGITAL
    JOYFLAG_ANALOGUE = C.ALLEGRO_JOYFLAG_ANALOGUE
)

// Installs the Allegro Joystick module.
func InstallJoystick() bool {
    return bool(C.al_install_joystick())
}

// Uninstalls the Allegro Joystick module.
func UninstallJoystick() {
    C.al_uninstall_joystick()
}

// Returns true if the Allegro joystick module ws instaled, false if not.
func IsJoystickInstalled() bool {
    return bool(C.al_is_joystick_installed())
}

// Returns the amount of joysticks detected.
func GetNumJoysticks() int {
    return int(C.al_get_num_joysticks())
}

// Returns the joyn'th joystick, or nil if no such stick exists. 
func GetJoystick(joyn int) *Joystick {
    return wrapJoystick(C.al_get_joystick(C.int(joyn)))
}

// Joystick properties.

// Returns true if the joystick self is active, false if not.
func (self *Joystick) IsActive() bool {
    return bool(C.al_get_joystick_active(self.handle))
}

// Returns the name of the joystick self.
func (self *Joystick) GetName() string {
    return gostr(C.al_get_joystick_name(self.handle))
}

// Returns the amount of sticks the joystick self has.
func (self *Joystick) GetNumSticks() int {
    return int(C.al_get_joystick_num_sticks(self.handle))
}

// Returns the joystick flags for the numbered stick on the joystick self.
func (self *Joystick) GetStickFlags(stick int) int {
    return int(C.al_get_joystick_stick_flags(self.handle, C.int(stick)))
}

// Returns true if the numbered stick on joystick self is digital, false if not.
// Note that theoretically, a stick could be both digital and analog...
func (self *Joystick) IsStickDigital(stick int) bool {
    return (JOYFLAG_DIGITAL & self.GetStickFlags(stick)) == JOYFLAG_DIGITAL
}

// Returns true if the numbered stick on joystick self is analog, false if not
// Note that theoretically, a stick could be both digital and analog...
func (self *Joystick) IsStickAnalog(stick int) bool {
    return (JOYFLAG_ANALOGUE & self.GetStickFlags(stick)) == JOYFLAG_ANALOGUE
}

// Returns a string that describes the joystick flags for the numbered stick 
// on the joystick self. Will return "Analog" for an analog joystick, 
// "Digital" for a digital joystick, "Hybrid" fo one that's both and 
// "None" for one that's neither
func (self *Joystick) GetStickFlagsName(stick int) string {
    if self.IsStickAnalog(stick) {
        if self.IsStickDigital(stick) {
            return "Hybrid"
        } else {
            return "Analog"
        }
    }
    if self.IsStickDigital(stick) {
        return "Digital"
    }
    return "None"
}

// Returns the name of the stick on the joystick self.
func (self *Joystick) GetStickName(stick int) string {
    return gostr(C.al_get_joystick_stick_name(self.handle, C.int(stick)))
}

// Returns the amount of axes for the stick on the joystick self.
func (self *Joystick) GetNumAxes(stick int) int {
    return int(C.al_get_joystick_num_axes(self.handle, C.int(stick)))
}

// Returns the name of the axis for the stick on the joystick self.
func (self *Joystick) GetAxisName(stick, axis int) string {
    return gostr(C.al_get_joystick_axis_name(self.handle, C.int(stick), C.int(axis)))
}

// Returns the amount of buttons on the joystick self.
func (self *Joystick) GetNumButtons() int {
    return int(C.al_get_joystick_num_buttons(self.handle))
}

// Returns the name of the button on the joystick self.
func (self *Joystick) GetButtonName(button int) string {
    return gostr(C.al_get_joystick_button_name(self.handle, C.int(button)))
}

// Gets the state of the joystick
func (self *Joystick) GetState() *JoystickState {
    state := &JoystickState{}
    C.al_get_joystick_state(self.handle, state.toC())
    return state
}

/*
AL_FUNC(bool,           al_install_joystick,    (void));
AL_FUNC(void,           al_uninstall_joystick,  (void));
AL_FUNC(bool,           al_is_joystick_installed, (void));
AL_FUNC(bool,           al_reconfigure_joysticks, (void));

AL_FUNC(int,            al_get_num_joysticks,   (void));
AL_FUNC(ALLEGRO_JOYSTICK *, al_get_joystick,    (int joyn));
AL_FUNC(void,           al_release_joystick,    (ALLEGRO_JOYSTICK *));
AL_FUNC(bool,           al_get_joystick_active, (ALLEGRO_JOYSTICK *));
AL_FUNC(const char*,    al_get_joystick_name,   (ALLEGRO_JOYSTICK *));

AL_FUNC(int,            al_get_joystick_num_sticks, (ALLEGRO_JOYSTICK *));
AL_FUNC(int, al_get_joystick_stick_flags, (ALLEGRO_JOYSTICK *, int stick)); 

AL_FUNC(const char*,    al_get_joystick_stick_name, (ALLEGRO_JOYSTICK *, int stick));

AL_FUNC(int,            al_get_joystick_num_axes,   (ALLEGRO_JOYSTICK *, int stick));
AL_FUNC(const char*,    al_get_joystick_axis_name,  (ALLEGRO_JOYSTICK *, int stick, int axis));

AL_FUNC(int,            al_get_joystick_num_buttons,  (ALLEGRO_JOYSTICK *));
AL_FUNC(const char*,    al_get_joystick_button_name,  (ALLEGRO_JOYSTICK *, int buttonn));

AL_FUNC(void,           al_get_joystick_state,  (ALLEGRO_JOYSTICK *, ALLEGRO_JOYSTICK_STATE *ret_state));

AL_FUNC(ALLEGRO_EVENT_SOURCE *, al_get_joystick_event_source, (void));
*/
