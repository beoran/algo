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

const PI = 3.14159265358979323846

// Allegro library ID calculation.
func AL_ID(a, b, c, d int) int {
	return (((a) << 24) | ((b) << 16) | ((c) << 8) | (d))
}

const VERSION = 5
const SUB_VERSION = 1
const WIP_VERSION = 5
const RELEASE_NUMBER = 1
const VERSION_STR = "5.0.7"
const DATE_STR = "2012"
const DATE = 20120624 /* yyyymmdd */
const VERSION_INT = ((VERSION << 24) | (SUB_VERSION << 16) |
	(WIP_VERSION << 8) | RELEASE_NUMBER)

// Converts bool to Allegro's C.bool
func b2cb(res bool) C.bool {
	if res {
		return C.bool(true)
	}
	return C.bool(false)
}

// Memory allocation, use this in stead of malloc for allegro stuff
func alMalloc(size uint) unsafe.Pointer {
	return C.al_malloc_with_context(C.size_t(size), 0, nil, nil)
}

// Memory allocation, use this in stead of calloc for allegro stuff
func alCalloc(size, n uint) unsafe.Pointer {
	return C.al_calloc_with_context(C.size_t(size), C.size_t(n), 0, nil, nil)
}

// Free memory, use this in stead of free for allegro stuff
func alFree(ptr unsafe.Pointer) {
	C.al_free_with_context(ptr, 0, nil, nil)
}

// Converts C.bool to Allegro's C.bool
func cb2b(res C.bool) bool {
	if res {
		return true
	}
	return false
}

// Checks if the basic Allegro system is installed or not.
func IsSystemInstalled() bool {
	return bool(C.al_is_system_installed())
}

// Gets the raw version of Allegro linked to as an integer.
func GetAllegroVersion() uint32 {
	return uint32(C.al_get_allegro_version())
}

// Initializes the Allegro system.
func Initialize() bool {
	return bool(C.al_install_system(VERSION_INT, nil))
	//	return bool(C.algo_initialize())
}

// Cleans up the Allegro system. Needed after calling Initialize.
func Cleanup() {
	C.al_uninstall_system()
}

// Installs the Allegro system. 
func InstallSystem() bool {
	return bool(C.al_install_system(VERSION_INT, nil))
}

// Uninstalls the Allegro system. Must be called after using InstallSystem.
func UninstallSystem() {
	C.al_uninstall_system()
}

// allegro5/path.h

// Wrapper for an Allegro path.
type Path struct {
	handle *C.ALLEGRO_PATH
}

// Wraps an Allegro path into the go struct above, but does not set a finalizer
func wrapPathRaw(handle *C.ALLEGRO_PATH) *Path {
	if handle == nil {
		return nil
	}
	return &Path{handle}
}

// Wraps an Allegro path into the go struct above, and sets the finalizer
// to be the struct's Destroy method
func wrapPath(handle *C.ALLEGRO_PATH) *Path {
	result := wrapPathRaw(handle)
	if result != nil {
		runtime.SetFinalizer(result, func(me *Path) { me.Destroy() })
	}
	return result
}

// Creates an Allegro path.
func CreatePath(str string) *Path {
	cstr := cstr(str)
	defer cstrFree(cstr)
	return wrapPath(C.al_create_path(cstr))
}

// Creates an allegro path for a directory.
func CreatePathForDirectory(str string) *Path {
	cstr := cstr(str)
	defer cstrFree(cstr)
	return wrapPath(C.al_create_path_for_directory(cstr))
}

// Clones an allegro path.
func (self *Path) ClonePath() *Path {
	return wrapPath(C.al_clone_path(self.handle))
}

// Destroys an Allegro path. It may not be used after this.
// Destroy may be called many times.
func (self *Path) Destroy() {
	if self.handle != nil {
		C.al_destroy_path(self.handle)
	}
	self.handle = nil
}

// Gets amount of components of the path name
func (self *Path) GetPathNumComponents() int {
	return int(C.al_get_path_num_components(self.handle))
}

// converst the allegro path to a string 
func (self *Path) String() string {
	return gostr(C.al_path_cstr(self.handle, C.char(NATIVE_PATH_SEP)))
}

/*
func (self * Path) 


AL_FUNC(int, al_get_path_num_components, (const ALLEGRO_PATH *path));
AL_FUNC(const char*, al_get_path_component, (const ALLEGRO_PATH *path, int i));
AL_FUNC(void, al_replace_path_component, (ALLEGRO_PATH *path, int i, const char *s));
AL_FUNC(void, al_remove_path_component, (ALLEGRO_PATH *path, int i));
AL_FUNC(void, al_insert_path_component, (ALLEGRO_PATH *path, int i, const char *s));
AL_FUNC(const char*, al_get_path_tail, (const ALLEGRO_PATH *path));
AL_FUNC(void, al_drop_path_tail, (ALLEGRO_PATH *path));
AL_FUNC(void, al_append_path_component, (ALLEGRO_PATH *path, const char *s));
AL_FUNC(bool, al_join_paths, (ALLEGRO_PATH *path, const ALLEGRO_PATH *tail));
AL_FUNC(bool, al_rebase_path, (const ALLEGRO_PATH *head, ALLEGRO_PATH *tail));
AL_FUNC(const char*, al_path_cstr, (const ALLEGRO_PATH *path, char delim));
AL_FUNC(void, al_destroy_path, (ALLEGRO_PATH *path));

AL_FUNC(void, al_set_path_drive, (ALLEGRO_PATH *path, const char *drive));
AL_FUNC(const char*, al_get_path_drive, (const ALLEGRO_PATH *path));

AL_FUNC(void, al_set_path_filename, (ALLEGRO_PATH *path, const char *filename));
AL_FUNC(const char*, al_get_path_filename, (const ALLEGRO_PATH *path));

AL_FUNC(const char*, al_get_path_extension, (const ALLEGRO_PATH *path));
AL_FUNC(bool, al_set_path_extension, (ALLEGRO_PATH *path, char const *extension));
AL_FUNC(const char*, al_get_path_basename, (const ALLEGRO_PATH *path));

AL_FUNC(bool, al_make_path_canonical, (ALLEGRO_PATH *path));

// defer
*/

// Not wrapped yet: 
// AL_FUNC(SYSTEM *, al_get_system_driver, (void));
// AL_FUNC(CONFIG *, al_get_system_config, (void));

const (
	RESOURCES_PATH = iota
	TEMP_PATH
	USER_DATA_PATH
	USER_HOME_PATH
	USER_SETTINGS_PATH
	USER_DOCUMENTS_PATH
	EXENAME_PATH
	LAST_PATH
)

// Gets a standard path location.
func GetStandardPath(id int) *Path {
	return wrapPath(C.al_get_standard_path(C.int(id)))
}

// Sets the name of the executable.
func SetExeName(name string) {
	C.al_set_exe_name(cstr(name))
}

// Sets the name of the organisation.
func SetOrgName(name string) {
	C.al_set_org_name(cstr(name))
}

// Sets the name of the app.
func SetAppName(name string) {
	C.al_set_app_name(cstr(name))
}

// Gets the name of the organisation
func GetOrgName() string {
	return gostr(C.al_get_org_name())
}

// Sets the name of the app
func GetAppName() string {
	return gostr(C.al_get_app_name())
}

// Inibits the screensaver, or not debending on inhibit.
func InhibitScreensaver(inhibit bool) bool {
	return bool(C.al_inhibit_screensaver(C.bool(inhibit)))
}

/// XXX How to wrap this and is it needed????
// AL_FUNC(int, al_run_main, (int argc, char **argv, int (*)(int, char **)));

/** Allegro has it's own string type. While it's nice, it's 
not needed in Go, so I will just wrap the basic conversion functions. */

type USTR struct {
	handle *C.ALLEGRO_USTR
}

// Frees an Allegro unicode string.
func (self *USTR) Free() {
	if self.handle != nil {
		C.al_ustr_free(self.handle)
	}
	self.handle = nil
}

// Just for consistency and to allow SelfDestruuct to work 
func (self *USTR) Destroy() {
	self.Free()
}

// Converts an Allegro Unicode string to a Go string 
func (self *USTR) String() string {
	if self.handle == nil {
		return "<destroyed>"
	}
	return C.GoStringN(C.al_cstr(self.handle), C.int(C.al_ustr_size(self.handle)))
}

// Wraps an Allegro USTR into the go struct above, but does not set a finalizer
func wrapUSTRRaw(handle *C.ALLEGRO_USTR) *USTR {
	if handle == nil {
		return nil
	}
	return &USTR{handle}
}

// Wraps an Allegro path into the go struct above, and sets the finalizer to 
// be the Destroy method of that struct.
func wrapUSTR(handle *C.ALLEGRO_USTR) *USTR {
	result := wrapUSTRRaw(handle)
	if result != nil {
		runtime.SetFinalizer(result, func(me *USTR) { me.Destroy() })
	}
	return result
}

// Converts a go string to an Allegro Unicode string
func USTRV(str string) *USTR {
	cstr := cstr(str)
	defer cstrFree(cstr)
	return wrapUSTR(C.al_ustr_new(cstr))
}

// Converts a go string to an Allegro Unicode string
func USTRP(str *string) *USTR {
	return USTRV(*str)
}

// Allegro's timer functions 

// Gets the time the app is running in seconds 
func GetTime() float64 {
	return float64(C.al_get_time())
}

// Sleeps the given amount of seconds
func Rest(seconds float64) {
	C.al_rest(C.double(seconds))
}

// Event Type, not to avoid complications.
// type EVENT_TYPE C.ALLEGRO_EVENT_TYPE

// Event Type constants
const (
	EVENT_JOYSTICK_AXIS          = C.ALLEGRO_EVENT_JOYSTICK_AXIS
	EVENT_JOYSTICK_BUTTON_DOWN   = C.ALLEGRO_EVENT_JOYSTICK_BUTTON_DOWN
	EVENT_JOYSTICK_BUTTON_UP     = C.ALLEGRO_EVENT_JOYSTICK_BUTTON_UP
	EVENT_JOYSTICK_CONFIGURATION = C.ALLEGRO_EVENT_JOYSTICK_CONFIGURATION

	EVENT_KEY_DOWN = C.ALLEGRO_EVENT_KEY_DOWN
	EVENT_KEY_CHAR = C.ALLEGRO_EVENT_KEY_CHAR
	EVENT_KEY_UP   = C.ALLEGRO_EVENT_KEY_UP

	EVENT_MOUSE_AXES          = C.ALLEGRO_EVENT_MOUSE_AXES
	EVENT_MOUSE_BUTTON_DOWN   = C.ALLEGRO_EVENT_MOUSE_BUTTON_DOWN
	EVENT_MOUSE_BUTTON_UP     = C.ALLEGRO_EVENT_MOUSE_BUTTON_UP
	EVENT_MOUSE_ENTER_DISPLAY = C.ALLEGRO_EVENT_MOUSE_ENTER_DISPLAY
	EVENT_MOUSE_LEAVE_DISPLAY = C.ALLEGRO_EVENT_MOUSE_LEAVE_DISPLAY
	EVENT_MOUSE_WARPED        = C.ALLEGRO_EVENT_MOUSE_WARPED

	EVENT_TIMER = C.ALLEGRO_EVENT_TIMER

	EVENT_DISPLAY_EXPOSE      = C.ALLEGRO_EVENT_DISPLAY_EXPOSE
	EVENT_DISPLAY_RESIZE      = C.ALLEGRO_EVENT_DISPLAY_RESIZE
	EVENT_DISPLAY_CLOSE       = C.ALLEGRO_EVENT_DISPLAY_CLOSE
	EVENT_DISPLAY_LOST        = C.ALLEGRO_EVENT_DISPLAY_LOST
	EVENT_DISPLAY_FOUND       = C.ALLEGRO_EVENT_DISPLAY_FOUND
	EVENT_DISPLAY_SWITCH_IN   = C.ALLEGRO_EVENT_DISPLAY_SWITCH_IN
	EVENT_DISPLAY_SWITCH_OUT  = C.ALLEGRO_EVENT_DISPLAY_SWITCH_OUT
	EVENT_DISPLAY_ORIENTATION = C.ALLEGRO_EVENT_DISPLAY_ORIENTATION
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

// Event sources that emit events.
type EventSource C.ALLEGRO_EVENT_SOURCE

// Wraps an event source pointer but sets no finalizer (not needed anyway) 
func wrapEventSourceRaw(ptr *C.ALLEGRO_EVENT_SOURCE) *EventSource {
	return (*EventSource)(ptr)
}

// Converts wrapper Event source pointer to C Allegro event pointer
func (self *EventSource) toC() *C.ALLEGRO_EVENT_SOURCE {
	return (*C.ALLEGRO_EVENT_SOURCE)(self)
}

// Events that the event system emits.
type Event C.ALLEGRO_EVENT

// Converts wrapper Event pointer to C Allegro event pointer
func (self *Event) toC() *C.ALLEGRO_EVENT {
	return (*C.ALLEGRO_EVENT)(self)
}

// Returns an unsafe pointer to the event 
func (self *Event) toPointer() unsafe.Pointer {
	return unsafe.Pointer(self.toC())
}

// Converts wrapper Event pointer to C Allegro any event
func (self *Event) ANY_EVENT() *C.ALLEGRO_ANY_EVENT {
	return (*C.ALLEGRO_ANY_EVENT)(self.toPointer())
}

// Converts wrapper Event pointer to C Allegro display event
func (self *Event) DISPLAY_EVENT() *C.ALLEGRO_DISPLAY_EVENT {
	return (*C.ALLEGRO_DISPLAY_EVENT)(self.toPointer())
}

// Converts wrapper Event pointer to C Allegro joystick event
func (self *Event) JOYSTICK_EVENT() *C.ALLEGRO_JOYSTICK_EVENT {
	return (*C.ALLEGRO_JOYSTICK_EVENT)(self.toPointer())
}

// Converts wrapper Event pointer to C Allegro event
func (self *Event) KEYBOARD_EVENT() *C.ALLEGRO_KEYBOARD_EVENT {
	return (*C.ALLEGRO_KEYBOARD_EVENT)(self.toPointer())
}

// Converts wrapper Event pointer to C Allegro touch event
func (self *Event) TOUCH_EVENT() *C.ALLEGRO_TOUCH_EVENT {
	return (*C.ALLEGRO_TOUCH_EVENT)(self.toPointer())
}

// Converts wrapper Event pointer to C Allegro mouse event
func (self *Event) MOUSE_EVENT() *C.ALLEGRO_MOUSE_EVENT {
	return (*C.ALLEGRO_MOUSE_EVENT)(self.toPointer())
}

// Converts wrapper Event pointer to C Allegro timer event
func (self *Event) TIMER_EVENT() *C.ALLEGRO_TIMER_EVENT {
	return (*C.ALLEGRO_TIMER_EVENT)(self.toPointer())
}

// Converts wrapper Event pointer to C Allegro event
func (self *Event) USER_EVENT() *C.ALLEGRO_USER_EVENT {
	return (*C.ALLEGRO_USER_EVENT)(self.toPointer())
}

// Returns the type of the event.
func (self *Event) Type() int {
	return int(self.ANY_EVENT()._type)
}

// Returns the timestamp of the event.
func (self *Event) Timestamp() float64 {
	return float64(self.ANY_EVENT().timestamp)
}

// Returns the event source of the event
func (self *Event) EventSource() *EventSource {
	return (*EventSource)(self.ANY_EVENT().source)
}

// Returns true if this is a dispay event, false if not.
func (self *Event) IsDisplay() bool {
	t := self.Type()
	return (t >= EVENT_DISPLAY_EXPOSE) && (t <= EVENT_DISPLAY_ORIENTATION)
}

// Returns true if this is a mouse event, false if not.
func (self *Event) IsMouse() bool {
	t := self.Type()
	return (t >= EVENT_MOUSE_AXES) && (t <= EVENT_MOUSE_WARPED)
}

// Returns true if this is a Joystick event, false if not.
func (self *Event) IsJoystick() bool {
	t := self.Type()
	return (t >= EVENT_JOYSTICK_AXIS) && (t <= EVENT_JOYSTICK_CONFIGURATION)
}

// Returns true if this is a keyboard event, false if not.
func (self *Event) IsKeyboard() bool {
	t := self.Type()
	return (t >= EVENT_KEY_DOWN) && (t <= EVENT_KEY_UP)
}

// Returns true if this is a timer event, false if not.
func (self *Event) IsTimer() bool {
	t := self.Type()
	return (t == EVENT_TIMER)
}

// Returns true if this is a user event, false if not.
func (self *Event) IsUser() bool {
	t := self.Type()
	return EVENT_TYPE_IS_USER(t)
}

// Returns the event's source pointer
func (self *Event) EVENT_SOURCE() *C.ALLEGRO_EVENT_SOURCE {
	return self.ANY_EVENT().source
}

// Returns an unsafe pointer to the event's source pointer
func (self *Event) EVENT_SOURCE_PTR() unsafe.Pointer {
	return unsafe.Pointer(self.ANY_EVENT())
}

// Returns the display that has emitted the event. Will return nil if 
// this is not a display event.
func (self *Event) DisplayDisplay() *Display {
	if !(self.IsDisplay()) {
		return nil
	}
	return wrapDisplayRaw((*C.ALLEGRO_DISPLAY)(self.EVENT_SOURCE_PTR()))
}

// Returns the X position of the display event. Will return garbage 
// if this is not a display event.
func (self *Event) DisplayX() int {
	return int((self.DISPLAY_EVENT()).x)
}

// Returns the Y position of the display event. Will return garbage 
// if this is not a display event.
func (self *Event) DisplayY() int {
	return int(self.DISPLAY_EVENT().y)
}

// Returns the width of the display event. Will return garbage 
// if this is not a display event.
func (self *Event) DisplayWidth() int {
	return int(self.DISPLAY_EVENT().width)
}

// Returns the height of the display event. Will return garbage 
// if this is not a display event.
func (self *Event) DisplayHeight() int {
	return int(self.DISPLAY_EVENT().height)
}

// Returns the orientation of the display event. Will return garbage 
// if this is not a display event.
func (self *Event) DisplayOrientation() int {
	return int(self.DISPLAY_EVENT().orientation)
}

// XXX: maybe also wrap the source in a Joystick type? 

// Returns the stick number of the joystick event. Will return garbage 
// if this is not a joystick event.
func (self *Event) JoystickStick() int {
	return int(self.JOYSTICK_EVENT().stick)
}

// Returns the axis number of the joystick event. Will return garbage 
// if this is not a joystick event.
func (self *Event) JoystickAxis() int {
	return int(self.JOYSTICK_EVENT().axis)
}

// Returns the button number of the joystick event. Will return garbage 
// if this is not a joystick event.
func (self *Event) JoystickButton() int {
	return int(self.JOYSTICK_EVENT().button)
}

// Returns the position of the joystick event. Will return garbage 
// if this is not a joystick event.
func (self *Event) JoystickPos() float32 {
	return float32(self.JOYSTICK_EVENT().pos)
}

/// XXX also wrap Keyboard event source?

// Returns the display that has emitted the keyboard event. Will return nil if 
// this is not a keyboard event.
func (self *Event) KeyboardDisplay() *Display {
	if !(self.IsKeyboard()) {
		return nil
	}
	return wrapDisplayRaw(self.KEYBOARD_EVENT().display)
}

// Returns the keycode of the keyboard event. Returns garbage 
// if this is not a keyboard event.
func (self *Event) KeyboardKeycode() int {
	return int(self.KEYBOARD_EVENT().keycode)
}

// Returns the unichar of the keyboard event. Returns garbage 
// if this is not a keyboard event.
func (self *Event) KeyboardUnichar() rune {
	return rune(self.KEYBOARD_EVENT().unichar)
}

// Returns the modifiers of the keyboard event. Returns garbage 
// if this is not a keyboard event.
func (self *Event) KeyboardModifiers() int {
	return int(self.KEYBOARD_EVENT().modifiers)
}

// Returns is the keyboard event was autorepeated or not. Returns garbage 
// if this is not a keyboard event.
func (self *Event) KeyboardRepeat() bool {
	return bool(self.KEYBOARD_EVENT().repeat)
}

// Returns the x postion of the mouse event. Returns garbage 
// if this is not a mouse event.
func (self *Event) MouseX() int {
	return int(self.MOUSE_EVENT().x)
}

// Returns the y postion of the mouse event. Returns garbage 
// if this is not a mouse event.
func (self *Event) MouseY() int {
	return int(self.MOUSE_EVENT().y)
}

// Returns the z postion of the mouse event. Returns garbage 
// if this is not a mouse event.
func (self *Event) MouseZ() int {
	return int(self.MOUSE_EVENT().z)
}

// Returns the w postion of the mouse event. Returns garbage 
// if this is not a mouse event.
func (self *Event) MouseW() int {
	return int(self.MOUSE_EVENT().w)
}

// Returns the dx of the mouse event. Returns garbage 
// if this is not a mouse event.
func (self *Event) MouseDX() int {
	return int(self.MOUSE_EVENT().dx)
}

// Returns the dy of the mouse event. Returns garbage 
// if this is not a mouse event.
func (self *Event) MouseDY() int {
	return int(self.MOUSE_EVENT().dy)
}

// Returns the dz of the mouse event. Returns garbage 
// if this is not a mouse event.
func (self *Event) MouseDZ() int {
	return int(self.MOUSE_EVENT().dz)
}

// Returns the dw of the mouse event. Returns garbage 
// if this is not a mouse event.
func (self *Event) MouseDW() int {
	return int(self.MOUSE_EVENT().dw)
}

// Returns the button of the mouse event. Returns garbage 
// if this is not a mouse event.
func (self *Event) MouseButton() int {
	return int(self.MOUSE_EVENT().button)
}

// Returns the pressure of the mouse event. Returns garbage 
// if this is not a mouse event.
func (self *Event) MousePressure() float32 {
	return float32(self.MOUSE_EVENT().pressure)
}

// Returns the display that has emitted the mouse event. Will return nil if 
// this is not a mouse event.
func (self *Event) MouseDisplay() *Display {
	if !(self.IsMouse()) {
		return nil
	}
	return wrapDisplayRaw(self.MOUSE_EVENT().display)
}

// Returns the error of the timer event. Returns garbage 
// if this is not a timer event.
func (self *Event) TimerError() float64 {
	return float64(self.TIMER_EVENT().error)
}

// Returns the ticks of the timer event. Returns garbage 
// if this is not a timer event.
func (self *Event) TimerCount() int64 {
	return int64(self.TIMER_EVENT().count)
}

// Wrapping of user event seems not really meaningful in Go so leave that out.

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
func (self *EventQueue) GetNextEvent() (event *Event, ok bool) {
	event = &Event{}
	ok = bool(C.al_get_next_event(self.handle, event.toC()))
	return event, ok
}

// Peeks at the next event in the event queue and returns it as well as a bool
// to signify if an event was fetched sucessfully or not.
func (self *EventQueue) PeekNextEvent() (event *Event, ok bool) {
	event = &Event{}
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
func (self *EventQueue) WaitForEvent() (event *Event) {
	event = &Event{}
	C.al_wait_for_event(self.handle, event.toC())
	return event
}

// Waits for secs seconds the next event from the event queue 
func (self *EventQueue) WaitForEventTimed(secs float32) (event *Event, ok bool) {
	event = &Event{}
	ok = bool(C.al_wait_for_event_timed(self.handle, event.toC(), C.float(secs)))
	return event, ok
}

/*
// Emitting user events is omitted for now.
TODO: 
AL_FUNC(bool, al_wait_for_event_until, (ALLEGRO_EVENT_QUEUE *queue,
                                        ALLEGRO_EVENT *ret_event,
                                        ALLEGRO_TIMEOUT *timeout));
*/

// Precise (?) Timer 
type Timer struct {
	handle *C.ALLEGRO_TIMER
}

// Destroys the timer queue.
func (self *Timer) Destroy() {
	if self.handle != nil {
		C.al_destroy_timer(self.handle)
	}
	self.handle = nil
}

// Wraps a timer, but does not set a finalizer.
func wrapTimerRaw(handle *C.ALLEGRO_TIMER) *Timer {
	if handle == nil {
		return nil
	}
	return &Timer{handle}
}

// Wraps an event queue and sets a finalizer that calls Destroy
func wrapTimer(handle *C.ALLEGRO_TIMER) *Timer {
	result := wrapTimerRaw(handle)
	if result != nil {
		runtime.SetFinalizer(result, func(me *Timer) { me.Destroy() })
	}
	return result
}

// Creates a timer wih the given tick speed.
func CreateTimer(speed_secs float64) *Timer {
	return wrapTimer(C.al_create_timer(C.double(speed_secs)))
}

// Starts the timer.
func (self *Timer) Start() {
	C.al_start_timer(self.handle)
}

// Stops the timer.
func (self *Timer) Stop() {
	C.al_stop_timer(self.handle)
}

// Returns true if the timer was started, false if not 
func (self *Timer) IsStarted() bool {
	return bool(C.al_get_timer_started(self.handle))
}

// Sets the speed of the timer.
func (self *Timer) SetSpeed(speed_secs float64) {
	C.al_set_timer_speed(self.handle, C.double(speed_secs))
}

// Gets the speed of the timer.
func (self *Timer) GetSpeed() float64 {
	return float64(C.al_get_timer_speed(self.handle))
}

// Gets the count (in ticks) of the timer
func (self *Timer) GetCount() int {
	return int(C.al_get_timer_count(self.handle))
}

// Sets the count (in ticks) of the timer
func (self *Timer) SetCount(count int) {
	C.al_set_timer_count(self.handle, C.int64_t(count))
}

// Adds to the count (in ticks) of the timer
func (self *Timer) AddCount(count int) {
	C.al_add_timer_count(self.handle, C.int64_t(count))
}

// Gets the event source of this timer that can be registered 
// on an event queue with RegisterEventSource.
func (self *Timer) GetEventSource() *EventSource {
	return (*EventSource)(C.al_get_timer_event_source(self.handle))
}

// Do nothing function for benchmarking only
func DoNothing() {
}
