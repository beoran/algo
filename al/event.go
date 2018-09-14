package al

/*
#cgo pkg-config: allegro-5
#cgo CFLAGS: -I/usr/local/include
#cgo linux LDFLAGS: -lc_nonshared
#include <stdlib.h>
#include <allegro5/allegro.h>
#include <allegro5/events.h>
#include "helpers.h"
*/
import "C"

import "unsafe"
import "runtime"


// EventUnion Type, not to avoid complications.
// type EVENT_TYPE C.ALLEGRO_EVENT_TYPE

// EventUnion Type constants
const (
    EVENT_JOYSTICK_AXIS             = C.ALLEGRO_EVENT_JOYSTICK_AXIS
    EVENT_JOYSTICK_BUTTON_DOWN      = C.ALLEGRO_EVENT_JOYSTICK_BUTTON_DOWN
    EVENT_JOYSTICK_BUTTON_UP        = C.ALLEGRO_EVENT_JOYSTICK_BUTTON_UP
    EVENT_JOYSTICK_CONFIGURATION    = C.ALLEGRO_EVENT_JOYSTICK_CONFIGURATION
    EVENT_KEY_DOWN                  = C.ALLEGRO_EVENT_KEY_DOWN
    EVENT_KEY_CHAR                  = C.ALLEGRO_EVENT_KEY_CHAR
    EVENT_KEY_UP                    = C.ALLEGRO_EVENT_KEY_UP
    EVENT_MOUSE_AXES                = C.ALLEGRO_EVENT_MOUSE_AXES
    EVENT_MOUSE_BUTTON_DOWN         = C.ALLEGRO_EVENT_MOUSE_BUTTON_DOWN
    EVENT_MOUSE_BUTTON_UP           = C.ALLEGRO_EVENT_MOUSE_BUTTON_UP
    EVENT_MOUSE_ENTER_DISPLAY       = C.ALLEGRO_EVENT_MOUSE_ENTER_DISPLAY
    EVENT_MOUSE_LEAVE_DISPLAY       = C.ALLEGRO_EVENT_MOUSE_LEAVE_DISPLAY
    EVENT_MOUSE_WARPED              = C.ALLEGRO_EVENT_MOUSE_WARPED
    EVENT_TIMER                     = C.ALLEGRO_EVENT_TIMER
    EVENT_DISPLAY_EXPOSE            = C.ALLEGRO_EVENT_DISPLAY_EXPOSE
    EVENT_DISPLAY_RESIZE            = C.ALLEGRO_EVENT_DISPLAY_RESIZE
    EVENT_DISPLAY_CLOSE             = C.ALLEGRO_EVENT_DISPLAY_CLOSE
    EVENT_DISPLAY_LOST              = C.ALLEGRO_EVENT_DISPLAY_LOST
    EVENT_DISPLAY_FOUND             = C.ALLEGRO_EVENT_DISPLAY_FOUND
    EVENT_DISPLAY_SWITCH_IN         = C.ALLEGRO_EVENT_DISPLAY_SWITCH_IN
    EVENT_DISPLAY_SWITCH_OUT        = C.ALLEGRO_EVENT_DISPLAY_SWITCH_OUT
    EVENT_DISPLAY_ORIENTATION       = C.ALLEGRO_EVENT_DISPLAY_ORIENTATION
    EVENT_TOUCH_BEGIN               = C.ALLEGRO_EVENT_TOUCH_BEGIN
    EVENT_TOUCH_END                 = C.ALLEGRO_EVENT_TOUCH_END
    EVENT_TOUCH_MOVE                = C.ALLEGRO_EVENT_TOUCH_MOVE
    EVENT_TOUCH_CANCEL              = C.ALLEGRO_EVENT_TOUCH_CANCEL
    EVENT_DISPLAY_CONNECTED         = C.ALLEGRO_EVENT_DISPLAY_CONNECTED
    EVENT_DISPLAY_DISCONNECTED      = C.ALLEGRO_EVENT_DISPLAY_DISCONNECTED
)

func EVENT_TYPE_IS_USER(t int) bool {
    return ((t) >= 512)
}

func GET_EVENT_TYPE(a, b, c, d int) int {
    return AL_ID(a, b, c, d)
}

func getAnyEvenTimestamp(any *C.ALLEGRO_ANY_EVENT) float64 {
    return float64(any.timestamp)
}

// Wrapper interface for different event sources
type EventSourcer interface {
    EventSource() * EventSource
}

func (disp * Display) EventSource() * EventSource {
    return (*EventSource)(unsafe.Pointer(disp))
}

func (joy * Joystick) EventSource() * EventSource {
    return (*EventSource)(unsafe.Pointer(joy))
}

func (key * Keyboard) EventSource() * EventSource {
    return (*EventSource)(unsafe.Pointer(key))
}

func (mouse * Mouse) EventSource() * EventSource {
    return (*EventSource)(unsafe.Pointer(mouse))
}

func (touch * TouchInput) EventSource() * EventSource {
    return (*EventSource)(unsafe.Pointer(touch))
}

func (timer * Timer) EventSource() * EventSource {
    return (*EventSource)(unsafe.Pointer(timer))
}

func (es * EventSource) EventSource() * EventSource {
    return es
}

 
// Event sources that emit events.
type EventSource C.ALLEGRO_EVENT_SOURCE

// Wraps an event source pointer but sets no finalizer (not needed anyway) 
func wrapEventSourceRaw(ptr *C.ALLEGRO_EVENT_SOURCE) *EventSource {
    return (*EventSource)(ptr)
}

// Converts wrapper EventUnion source pointer to C Allegro event source pointer
func (self *EventSource) toC() *C.ALLEGRO_EVENT_SOURCE {
    return (*C.ALLEGRO_EVENT_SOURCE)(self)
}

// Events that the event system emits.
type EventUnion     C.ALLEGRO_EVENT

// use an EventUnion interface to wrap the 
// C union's members in

type Event interface {
    Type() int
    Source() EventSourcer
    Timestamp() float64
}

type AnyEvent       C.ALLEGRO_ANY_EVENT
type DisplayEvent   C.ALLEGRO_DISPLAY_EVENT
type JoystickEvent  C.ALLEGRO_JOYSTICK_EVENT
type KeyboardEvent  C.ALLEGRO_KEYBOARD_EVENT
type MouseEvent     C.ALLEGRO_MOUSE_EVENT
type TimerEvent     C.ALLEGRO_TIMER_EVENT
type TouchEvent     C.ALLEGRO_TOUCH_EVENT
type UserEvent      C.ALLEGRO_USER_EVENT


func (ev JoystickEvent        ) MarkJoystick()  { ; }
func (ev KeyboardEvent        ) MarkKey()       { ; }
func (ev MouseEvent           ) MarkMouse()     { ; }
func (ev DisplayEvent         ) MarkDisplay() { ; }
func (ev TouchEvent           ) MarkTouch() { ; }             

// Converts wrapper EventUnion pointer to C Allegro event pointer
func (self *EventUnion) toC() *C.ALLEGRO_EVENT {
    return (*C.ALLEGRO_EVENT)(self)
}

// Returns an unsafe pointer to the event 
func (self *EventUnion) toPointer() unsafe.Pointer {
    return unsafe.Pointer(self.toC())
}

func (self *EventUnion) toUP() unsafe.Pointer {
    return unsafe.Pointer(self.toC())
}

// Converts wrapper EventUnion pointer to an Event interface struct
func (evun *EventUnion) Event() Event {
    ae := (*C.ALLEGRO_ANY_EVENT)(evun.toPointer())
    switch (C.int(ae._type)) {
        case EVENT_JOYSTICK_AXIS            : fallthrough 
        case EVENT_JOYSTICK_BUTTON_DOWN     : fallthrough 
        case EVENT_JOYSTICK_BUTTON_UP       : fallthrough
        case EVENT_JOYSTICK_CONFIGURATION   : 
            return (*JoystickEvent)(evun.toPointer())
        case EVENT_KEY_CHAR                 : fallthrough
        case EVENT_KEY_DOWN                 : fallthrough
        case EVENT_KEY_UP                   : 
            return (*KeyboardEvent)(evun.toPointer())
 
        case EVENT_MOUSE_AXES               : fallthrough    
        case EVENT_MOUSE_BUTTON_DOWN        : fallthrough    
        case EVENT_MOUSE_BUTTON_UP          : fallthrough    
        case EVENT_MOUSE_ENTER_DISPLAY      : fallthrough    
        case EVENT_MOUSE_LEAVE_DISPLAY      : fallthrough    
        case EVENT_MOUSE_WARPED             :
            return (*MouseEvent)(evun.toPointer())
            
        case EVENT_TIMER                    :
            return (*TimerEvent)(evun.toPointer())
            
        case EVENT_DISPLAY_EXPOSE           : fallthrough    
        case EVENT_DISPLAY_RESIZE           : fallthrough    
        case EVENT_DISPLAY_CLOSE            : fallthrough    
        case EVENT_DISPLAY_LOST             : fallthrough    
        case EVENT_DISPLAY_FOUND            : fallthrough    
        case EVENT_DISPLAY_SWITCH_IN        : fallthrough    
        case EVENT_DISPLAY_SWITCH_OUT       : fallthrough    
        case EVENT_DISPLAY_ORIENTATION      :
            return (*DisplayEvent)(evun.toPointer())

        case EVENT_TOUCH_BEGIN              : fallthrough
        case EVENT_TOUCH_MOVE               : fallthrough
        case EVENT_TOUCH_CANCEL             : fallthrough
        case EVENT_TOUCH_END                :
            return (*TouchEvent)(evun.toPointer())        
        default: break
    }
    
    if EVENT_TYPE_IS_USER(int(ae._type)) {
            return (*UserEvent)(evun.toPointer())
    }
    
    return (*AnyEvent)(evun.toPointer())
}

/* These wrappers implement the interface Event for all event types */

func (ev AnyEvent     ) Type() int { return int(ev._type); }
func (ev DisplayEvent ) Type() int { return int(ev._type); }
func (ev JoystickEvent) Type() int { return int(ev._type); }
func (ev KeyboardEvent) Type() int { return int(ev._type); }
func (ev MouseEvent   ) Type() int { return int(ev._type); }
func (ev TimerEvent   ) Type() int { return int(ev._type); }
func (ev TouchEvent   ) Type() int { return int(ev._type); }
func (ev UserEvent    ) Type() int { return int(ev._type); }

func (ev AnyEvent     ) Source() EventSourcer { return wrapEventSourceRaw   (ev.source); }
func (ev DisplayEvent ) Source() EventSourcer { return wrapDisplayRaw       (ev.source); }
func (ev JoystickEvent) Source() EventSourcer { return wrapJoystickRaw      (ev.source); }
func (ev KeyboardEvent) Source() EventSourcer { return wrapKeyboardRaw      (ev.source); }
func (ev MouseEvent   ) Source() EventSourcer { return wrapMouseRaw         (ev.source); }
func (ev TimerEvent   ) Source() EventSourcer { return wrapTimerRaw         (ev.source); }
func (ev TouchEvent   ) Source() EventSourcer { return wrapTouchInputRaw    (ev.source); }
func (ev UserEvent    ) Source() EventSourcer { return wrapEventSourceRaw   (ev.source); }

func (ev AnyEvent     ) Timestamp() float64 { return float64(ev.timestamp); }
func (ev DisplayEvent ) Timestamp() float64 { return float64(ev.timestamp); }
func (ev JoystickEvent) Timestamp() float64 { return float64(ev.timestamp); }
func (ev KeyboardEvent) Timestamp() float64 { return float64(ev.timestamp); }
func (ev MouseEvent   ) Timestamp() float64 { return float64(ev.timestamp); }
func (ev TimerEvent   ) Timestamp() float64 { return float64(ev.timestamp); }
func (ev TouchEvent   ) Timestamp() float64 { return float64(ev.timestamp); }
func (ev UserEvent    ) Timestamp() float64 { return float64(ev.timestamp); }


func (eu * EventUnion) AnyEvent     () *AnyEvent      {                                                 return (*AnyEvent     )(eu.toPointer());   }
func (eu * EventUnion) DisplayEvent () *DisplayEvent  { if !eu.IsDisplayEvent () { return nil; } else { return (*DisplayEvent )(eu.toPointer());}; }
func (eu * EventUnion) JoystickEvent() *JoystickEvent { if !eu.IsJoystickEvent() { return nil; } else { return (*JoystickEvent)(eu.toPointer());}; }
func (eu * EventUnion) KeyboardEvent() *KeyboardEvent { if !eu.IsKeyboardEvent() { return nil; } else { return (*KeyboardEvent)(eu.toPointer());}; }
func (eu * EventUnion) MouseEvent   () *MouseEvent    { if !eu.IsMouseEvent   () { return nil; } else { return (*MouseEvent   )(eu.toPointer());}; }
func (eu * EventUnion) TimerEvent   () *TimerEvent    { if !eu.IsTimerEvent   () { return nil; } else { return (*TimerEvent   )(eu.toPointer());}; }
func (eu * EventUnion) TouchEvent   () *TouchEvent    { if !eu.IsTouchEvent   () { return nil; } else { return (*TouchEvent   )(eu.toPointer());}; }
func (eu * EventUnion) UserEvent    () *UserEvent     { if !eu.IsUserEvent    () { return nil; } else { return (*UserEvent    )(eu.toPointer());}; }

// Returns the type of the event.
func (self *EventUnion) Type() int {
    return self.AnyEvent().Type()
}

// Returns the timestamp of the event.
func (self *EventUnion) Timestamp() float64 {
    return self.AnyEvent().Timestamp()
}

// Returns the event source of the event
func (self *EventUnion) Source() EventSourcer {
    return self.AnyEvent().Source()
}

// Returns true if this is a dispay event, false if not.
func (self *EventUnion) IsDisplayEvent() bool {
    t := self.Type()
    return ((t >= EVENT_DISPLAY_EXPOSE) && (t <= EVENT_DISPLAY_ORIENTATION) || ((t>= EVENT_DISPLAY_CONNECTED) && (t<= EVENT_DISPLAY_DISCONNECTED) ))
}

// Returns true if this is a mouse event, false if not.
func (self *EventUnion) IsMouseEvent() bool {
    t := self.Type()
    return (t >= EVENT_MOUSE_AXES) && (t <= EVENT_MOUSE_WARPED)
}

// Returns true if this is a Joystick event, false if not.
func (self *EventUnion) IsJoystickEvent() bool {
    t := self.Type()
    return (t >= EVENT_JOYSTICK_AXIS) && (t <= EVENT_JOYSTICK_CONFIGURATION)
}

// Returns true if this is a keyboard event, false if not.
func (self *EventUnion) IsKeyboardEvent() bool {
    t := self.Type()
    return (t >= EVENT_KEY_DOWN) && (t <= EVENT_KEY_UP)
}

// Returns true if this is a keyboard event, false if not.
func (self *EventUnion) IsTouchEvent() bool {
    t := self.Type()
    return (t >= EVENT_TOUCH_BEGIN) && (t <= EVENT_TOUCH_CANCEL)
}

// Returns true if this is a timer event, false if not.
func (self *EventUnion) IsTimerEvent() bool {
    t := self.Type()
    return (t == EVENT_TIMER)
}

// Returns true if this is a user event, false if not.
func (self *EventUnion) IsUserEvent() bool {
    t := self.Type()
    return EVENT_TYPE_IS_USER(t)
}

// Returns the X position of the display event.
func (de DisplayEvent) X() int {
    return int(de.x)
}

// Returns the Y position of the display event.
func (de DisplayEvent) Y() int {
    return int(de.y)
}

// Returns the width of the display event.
func (de DisplayEvent) Width() int {
    return int(de.width)
}

// Returns the height of the display event.
func (de DisplayEvent) Height() int {
    return int(de.height)
}

// Returns the orientation of the display event.
func (de DisplayEvent) Orientation() int {
    return int(de.orientation)
}

// Returns the ID of the joystick for the joystick event.
func (je JoystickEvent) ID() int {
    jsptr := je.id
    for i := 0 ; i < NumJoysticks() ; i++ {
        js := FindJoystick(i)
        if js.handle == jsptr {
            return i
        }
    }
    /* shouln't happen, but... */
    return 0xbad101
}


// Returns the stick number of the joystick event.
func (je JoystickEvent) Stick() int {
    return int(je.stick)
}

// Returns the axis number of the joystick event.
func (je JoystickEvent) Axis() int {
    return int(je.axis)
}

// Returns the button number of the joystick event.
func (je JoystickEvent) Button() int {
    return int(je.button)
}

// Returns the position of the joystick axis during the event.
func (je JoystickEvent) Pos() float32 {
    return float32(je.pos)
}

// Returns the display that has emitted the keyboard event.
func (ke KeyboardEvent) Display() *Display {
    return wrapDisplayRaw(ke.display)
}

// Returns the key code of the keyboard event. 
func (ke KeyboardEvent) KeyCode() int {
    return int(ke.keycode)
}

// Returns the unicode character of the keyboard event. 
func (ke KeyboardEvent) Unichar() rune {
    return rune(ke.unichar)
}

// Returns the modifiers of the keyboard event. 
func (ke KeyboardEvent) Modifiers() int {
    return int(ke.modifiers)
}

// Returns is the keyboard event was automatically repeated or not.
func (ke KeyboardEvent) Repeat() bool {
    return bool(ke.repeat)
}

// Returns the X position of the mouse event.
func (me MouseEvent) X() int {
    return int(me.x)
}

// Returns the Y position of the mouse event.
func (me MouseEvent) Y() int {
    return int(me.y)
}

// Returns the Z position of the mouse event.
func (me MouseEvent) Z() int {
    return int(me.z)
}

// Returns the W position of the mouse event.
func (me MouseEvent) W() int {
    return int(me.w)
}

// Returns the delta of the X position of the mouse event.
func (me MouseEvent) DX() int {
    return int(me.dx)
}

// Returns the delta of the Y position of the mouse event.
func (me MouseEvent) DY() int {
    return int(me.dy)
}

// Returns the delta of the Z position of the mouse event.
func (me MouseEvent) DZ() int {
    return int(me.dz)
}

// Returns the delta of the W position of the mouse event.
func (me MouseEvent) DW() int {
    return int(me.dw)
}

// Returns the button of the mouse event.
func (me MouseEvent) Button() int {
    return int(me.button)
}

// Returns the pressure of the mouse event.
func (me MouseEvent) Pressure() float32 {
    return float32(me.pressure)
}

// Returns the display that has emitted the mouse event. 
func (me MouseEvent) Display() *Display {
    return wrapDisplayRaw(me.display)
}


// Returns the error of the timer event.
func (te TimerEvent) Error() float64 {
    return float64(te.error)
}

// Returns the ticks of the timer event. 
func (te TimerEvent) Count() int64 {
    return int64(te.count)
}


// Returns the X position of the touch input event.
func (te TouchEvent) X() int {
    return int(te.x)
}

// Returns the Y position of the touch input event.
func (te TouchEvent) Y() int {
    return int(te.y)
}

// Returns the ID of the touch input event.
func (te TouchEvent) ID() int {
    return int(te.id)
}

// Returns whether the touch input event is primary or not.
func (te TouchEvent) Primary() bool {
    return bool(te.primary)
}

// Returns the delta of the X position of the touch input event.
func (te TouchEvent) DX() int {
    return int(te.dx)
}

// Returns the delta of the Y position of the touch input event.
func (te TouchEvent) DY() int {
    return int(te.dy)
}

/* The safest way to use user events from Go is to use integer handles
 * as offsets into a GO-allocated map. */

func (ue UserEvent) Data1Pointer() unsafe.Pointer { return unsafe.Pointer(uintptr(ue.data1));}
func (ue UserEvent) Data1Integer() int64          { return                  int64(ue.data1); }
func (ue UserEvent) Data2Pointer() unsafe.Pointer { return unsafe.Pointer(uintptr(ue.data2));}
func (ue UserEvent) Data2Integer() int64          { return                  int64(ue.data2); }
func (ue UserEvent) Data3Pointer() unsafe.Pointer { return unsafe.Pointer(uintptr(ue.data3));}
func (ue UserEvent) Data3Integer() int64          { return                  int64(ue.data3); }
func (ue UserEvent) Data4Pointer() unsafe.Pointer { return unsafe.Pointer(uintptr(ue.data4));}
func (ue UserEvent) Data4Integer() int64          { return                  int64(ue.data4); }

func (ue UserEvent) SetData1Integer(v int64)      { ue.data1 = C.intptr_t(v); }
func (ue UserEvent) SetData2Integer(v int64)      { ue.data2 = C.intptr_t(v); }
func (ue UserEvent) SetData3Integer(v int64)      { ue.data3 = C.intptr_t(v); }
func (ue UserEvent) SetData4Integer(v int64)      { ue.data4 = C.intptr_t(v); }


/*
The compiler accept this, but it's unlikely to work correctly.
func (ue UserEvent) Data1() interface{} { return (interface{})(ue.Data1Pointer());}
func (ue UserEvent) Data2() interface{} { return (interface{})(ue.Data2Pointer());}
func (ue UserEvent) Data3() interface{} { return (interface{})(ue.Data3Pointer());}
func (ue UserEvent) Data4() interface{} { return (interface{})(ue.Data4Pointer());}
*/



// Event queues.
type EventQueue struct {
    handle *C.ALLEGRO_EVENT_QUEUE
}

// Destroys the event queue.
func (self *EventQueue) Destroy() {
    if self.handle != nil {
        C.al_destroy_event_queue(self.handle)
    }
    self.handle = nil
}

// Wraps an event queue, but does not set a finalizer.
func wrapEventQueueRaw(handle *C.ALLEGRO_EVENT_QUEUE) *EventQueue {
    if handle == nil {
        return nil
    }
    return &EventQueue{handle}
}

func (eq * EventQueue) toC() (handle *C.ALLEGRO_EVENT_QUEUE) {
    return eq.handle
}

// Wraps an event queue and sets a finalizer that calls Destroy
func wrapEventQueue(handle *C.ALLEGRO_EVENT_QUEUE) *EventQueue {
    result := wrapEventQueueRaw(handle)
    if result != nil {
        runtime.SetFinalizer(result, func(me *EventQueue) { me.Destroy() })
    }
    return result
}

// Create an event queue.
func CreateEventQueue() *EventQueue {
    return wrapEventQueue(C.al_create_event_queue())
}

// Register an event source with self. 
func (self *EventQueue) RegisterEventSource(src *EventSource) {
    C.al_register_event_source(self.handle, src.toC())
}

// Unregister an event source with self. 
func (self *EventQueue) UnregisterEventSource(src *EventSource) {
    C.al_unregister_event_source(self.handle, src.toC())
}

// Returns true if the event queue self is empty, false if not.
func (self *EventQueue) IsEmpty() bool {
    return bool(C.al_is_event_queue_empty(self.handle))
}

// Returns the next event from the event queue as well as a bool
// to signify if an event was fetched sucessfully or not.
func (self *EventQueue) NextEvent() (event *EventUnion, ok bool) {
    event = &EventUnion{}
    ok = bool(C.al_get_next_event(self.handle, event.toC()))
    return event, ok
}

// Peeks at the next event in the event queue and returns it as well as a bool
// to signify if an event was fetched sucessfully or not.
func (self *EventQueue) PeekNextEvent() (event *EventUnion, ok bool) {
    event = &EventUnion{}
    ok = bool(C.al_peek_next_event(self.handle, event.toC()))
    return event, ok
}

// Drops the next event from the event queue
func (self *EventQueue) DropNextEvent() bool {
    return bool(C.al_drop_next_event(self.handle))
}

// Flushes the event queue
func (self *EventQueue) Flush() {
    C.al_flush_event_queue(self.handle)
}

// Waits for the next event from the event queue 
func (self *EventQueue) WaitForEvent() (event *EventUnion) {
    event = &EventUnion{}
    C.al_wait_for_event(self.handle, event.toC())
    return event
}

// Waits for secs seconds the next event from the event queue 
func (self *EventQueue) WaitForEventTimed(secs float32) (event *EventUnion, ok bool) {
    event = &EventUnion{}
    ok = bool(C.al_wait_for_event_timed(self.handle, event.toC(), C.float(secs)))
    return event, ok
}

func (queue *EventQueue) WaitForEventUntil(timeout * Timeout) (event *EventUnion, ok bool) {
    event = &EventUnion{}
    ok = bool(C.al_wait_for_event_until(queue.toC(), event.toC(), timeout.toC()))
    return event, ok
}

