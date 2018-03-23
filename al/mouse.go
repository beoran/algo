package al

/*
#include <stdlib.h>
#include <allegro5/allegro.h>
#include "helpers.h"
*/
import "C"

// import "unsafe"
import "runtime"


// Type that wraps a mouse 
type Mouse struct {
    handle *C.ALLEGRO_MOUSE
}

// Wraps a C Allegro mouse in a Mouse. Sets no finalizer.
func wrapMouseRaw(handle *C.ALLEGRO_MOUSE) *Mouse {
    if handle == nil {
        return nil
    }
    return &Mouse{handle}
}

// Destroys a mouse. Use this only when really needed!
func (self *Mouse) Destroy() {
    if self.handle != nil {
        // do nothing
    }
    self.handle = nil
}


// Wraps a C Allegro mouse cursor in a MouseCursor. Sets a finalizer that calls Destroy.
func wrapMouse(handle *C.ALLEGRO_MOUSE) *MouseCursor {
    self := wrapMouse(handle)
    if self != nil {
        runtime.SetFinalizer(self, func(me *Mouse) { me.Destroy() })
    }
    return self
}

// Type that wraps a mouse cursor
type MouseCursor struct {
    handle *C.ALLEGRO_MOUSE_CURSOR
}

// Returns low level hande for cursor
func (cursor *MouseCursor) toC() *C.ALLEGRO_MOUSE_CURSOR {
    return cursor.handle
}


// Destroys a mouse cursor. Use this only when really needed!
func (self *MouseCursor) Destroy() {
    if self.handle != nil {
        C.al_destroy_mouse_cursor(self.handle)
    }
    self.handle = nil
}

// Wraps a C Allegro mouse cursor in a MouseCursor. Sets no finalizer.
func wrapMouseCursorRaw(handle *C.ALLEGRO_MOUSE_CURSOR) *MouseCursor {
    if handle == nil {
        return nil
    }
    return &MouseCursor{handle}
}

// Wraps a C Allegro mouse cursor in a MouseCursor. Sets a finalizer that calls Destroy.
func wrapMouseCursor(handle *C.ALLEGRO_MOUSE_CURSOR) *MouseCursor {
    self := wrapMouseCursor(handle)
    if self != nil {
        runtime.SetFinalizer(self, func(me *MouseCursor) { me.Destroy() })
    }
    return self
}

// Mouse state type
type MouseState C.ALLEGRO_MOUSE_STATE

// Convert from C
func wrapMouseState(state C.ALLEGRO_MOUSE_STATE) MouseState {
    return MouseState(state)
}

// Convert to C
func (state MouseState) toC() C.ALLEGRO_MOUSE_STATE {
    return C.ALLEGRO_MOUSE_STATE(state)
}

// Convert to C
func (state * MouseState) toCPointer() * C.ALLEGRO_MOUSE_STATE {
    return (* C.ALLEGRO_MOUSE_STATE)(state)
}

func (state MouseState) X() int {
    return int(state.x)
}

func (state MouseState) Y() int {
    return int(state.y)
}

func (state MouseState) Z() int {
    return int(state.z)
}

func (state MouseState) W() int {
    return int(state.w)
}

func (state MouseState) Buttons() int {
    return int(state.buttons)
}


func (state MouseState) Pressure() float32 {
    return float32(state.pressure)
}

func (state MouseState) Display() * Display {
    return wrapDisplayRaw(state.display)
}

type SystemMouseCursor int

const (
    SYSTEM_MOUSE_CURSOR_NONE       = SystemMouseCursor(C.ALLEGRO_SYSTEM_MOUSE_CURSOR_NONE)       
    SYSTEM_MOUSE_CURSOR_DEFAULT    = SystemMouseCursor(C.ALLEGRO_SYSTEM_MOUSE_CURSOR_DEFAULT)    
    SYSTEM_MOUSE_CURSOR_ARROW      = SystemMouseCursor(C.ALLEGRO_SYSTEM_MOUSE_CURSOR_ARROW)      
    SYSTEM_MOUSE_CURSOR_BUSY       = SystemMouseCursor(C.ALLEGRO_SYSTEM_MOUSE_CURSOR_BUSY)       
    SYSTEM_MOUSE_CURSOR_QUESTION   = SystemMouseCursor(C.ALLEGRO_SYSTEM_MOUSE_CURSOR_QUESTION)   
    SYSTEM_MOUSE_CURSOR_EDIT       = SystemMouseCursor(C.ALLEGRO_SYSTEM_MOUSE_CURSOR_EDIT)       
    SYSTEM_MOUSE_CURSOR_MOVE       = SystemMouseCursor(C.ALLEGRO_SYSTEM_MOUSE_CURSOR_MOVE)       
    SYSTEM_MOUSE_CURSOR_RESIZE_N   = SystemMouseCursor(C.ALLEGRO_SYSTEM_MOUSE_CURSOR_RESIZE_N)   
    SYSTEM_MOUSE_CURSOR_RESIZE_W   = SystemMouseCursor(C.ALLEGRO_SYSTEM_MOUSE_CURSOR_RESIZE_W)   
    SYSTEM_MOUSE_CURSOR_RESIZE_S   = SystemMouseCursor(C.ALLEGRO_SYSTEM_MOUSE_CURSOR_RESIZE_S)   
    SYSTEM_MOUSE_CURSOR_RESIZE_E   = SystemMouseCursor(C.ALLEGRO_SYSTEM_MOUSE_CURSOR_RESIZE_E)   
    SYSTEM_MOUSE_CURSOR_RESIZE_NW  = SystemMouseCursor(C.ALLEGRO_SYSTEM_MOUSE_CURSOR_RESIZE_NW)  
    SYSTEM_MOUSE_CURSOR_RESIZE_SW  = SystemMouseCursor(C.ALLEGRO_SYSTEM_MOUSE_CURSOR_RESIZE_SW)  
    SYSTEM_MOUSE_CURSOR_RESIZE_SE  = SystemMouseCursor(C.ALLEGRO_SYSTEM_MOUSE_CURSOR_RESIZE_SE)  
    SYSTEM_MOUSE_CURSOR_RESIZE_NE  = SystemMouseCursor(C.ALLEGRO_SYSTEM_MOUSE_CURSOR_RESIZE_NE)  
    SYSTEM_MOUSE_CURSOR_PROGRESS   = SystemMouseCursor(C.ALLEGRO_SYSTEM_MOUSE_CURSOR_PROGRESS)  
    SYSTEM_MOUSE_CURSOR_PRECISION  = SystemMouseCursor(C.ALLEGRO_SYSTEM_MOUSE_CURSOR_PRECISION)  
    SYSTEM_MOUSE_CURSOR_LINK       = SystemMouseCursor(C.ALLEGRO_SYSTEM_MOUSE_CURSOR_LINK)       
    SYSTEM_MOUSE_CURSOR_ALT_SELECT = SystemMouseCursor(C.ALLEGRO_SYSTEM_MOUSE_CURSOR_ALT_SELECT) 
    SYSTEM_MOUSE_CURSOR_UNAVAILABLE= SystemMouseCursor(C.ALLEGRO_SYSTEM_MOUSE_CURSOR_UNAVAILABLE)
    NUM_SYSTEM_MOUSE_CURSORS       = SystemMouseCursor(C.ALLEGRO_NUM_SYSTEM_MOUSE_CURSORS)       
)

func IsMouseInstalled() bool {
    return cb2b(C.al_is_mouse_installed())
}

func InstallMouse() bool {
    return cb2b(C.al_install_mouse())
}

func UninstallMouse() {
    C.al_uninstall_mouse()
}

func GetMouseNumButtons() uint {
    return uint(C.al_get_mouse_num_buttons())
}

func GetMouseNumAxes() uint {
    return uint(C.al_get_mouse_num_axes())
}

func (display * Display) SetMouseXY(x , y int) bool {
    return cb2b(C.al_set_mouse_xy(display.toC(), ci(x) , ci(y)))
}

func SetMouseZ(z int) bool {
    return cb2b(C.al_set_mouse_z(ci(z)))
}

func SetMouseW(w int) bool {
    return cb2b(C.al_set_mouse_w(ci(w)))
}

func SetMouseAxis(axis, value int) bool {
    return cb2b(C.al_set_mouse_axis(ci(axis), ci(value)))
}

func AlGetMouseState() MouseState {
    var state C.ALLEGRO_MOUSE_STATE
    C.al_get_mouse_state(&state)
    return wrapMouseState(state)
}

func (state * MouseState) ButtonDown(button int) bool {
    return cb2b(C.al_mouse_button_down(state.toCPointer(), ci(button)))
}

func (state * MouseState) Axis(axis int) int {
    return int(C.al_get_mouse_state_axis(state.toCPointer(), ci(axis)))
}

func GetMouseEventSource() * EventSource {
    return wrapEventSourceRaw(C.al_get_mouse_event_source())
}

func CreateMouseCursor(sprite * Bitmap, xfocus, yfocus int) * MouseCursor {
    return wrapMouseCursor(C.al_create_mouse_cursor(sprite.toC(), ci(xfocus), ci(yfocus)))
}

func (display * Display) SetMouseCursor(cursor * MouseCursor) bool {
    return cb2b(C.al_set_mouse_cursor(display.toC(), cursor.toC()))
}

func (cursor SystemMouseCursor) toC() C.ALLEGRO_SYSTEM_MOUSE_CURSOR {
    return C.ALLEGRO_SYSTEM_MOUSE_CURSOR(cursor)
}

func (display * Display) SetSystemMouseCursor(cursor SystemMouseCursor) bool {
    return cb2b(C.al_set_system_mouse_cursor(display.toC(), cursor.toC()))
}

func (display * Display) ShowMouseCursor() bool {
    return cb2b(C.al_show_mouse_cursor(display.toC()))
}

func (display * Display) HideMouseCursor() bool {
    return cb2b(C.al_hide_mouse_cursor(display.toC()))
}

func GetMouseCursorPosition() (ok bool, x, y int) {
    var cx, cy C.int
    ok = cb2b(C.al_get_mouse_cursor_position(&cx, &cy))
    x = int(cx)
    y = int(cy)
    return ok, x, y
}

func (display * Display) GrabMouse() bool {
    return cb2b(C.al_grab_mouse(display.toC()))
}

func UngrabMouse() bool {
    return cb2b(C.al_ungrab_mouse())
}


