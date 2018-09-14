package al

/*
#cgo pkg-config: allegro-5
#cgo CFLAGS: -I/usr/local/include
#cgo linux LDFLAGS: -lc_nonshared
#include <stdlib.h>
#include <allegro5/allegro.h>
#include <allegro5/events.h>
#include "helpers.h"
#include "callbacks.h"
*/
import "C"

import "unsafe"
import "runtime"

const PI = 3.14159265358979323846

// Allegro library ID calculation.
func AL_ID(a, b, c, d int) int {
    return (((a) << 24) | ((b) << 16) | ((c) << 8) | (d))
}


const (
    VERSION         =  C.ALLEGRO_VERSION
    SUB_VERSION     =  C.ALLEGRO_SUB_VERSION
    WIP_VERSION     =  C.ALLEGRO_WIP_VERSION
    UNSTABLE_BIT    =  0
    RELEASE_NUMBER  =  C.ALLEGRO_RELEASE_NUMBER
    VERSION_STR     =  C.ALLEGRO_VERSION_STR
    DATE_STR        =  C.ALLEGRO_DATE_STR
    VERSION_INT     =  ((VERSION << 24) | (SUB_VERSION << 16) |
                        (WIP_VERSION << 8) | RELEASE_NUMBER | UNSTABLE_BIT)
 )

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
func AllegroVersion() uint32 {
    return uint32(C.al_get_allegro_version())
}

// Initializes the Allegro system.
func Initialize() bool {
    return bool(C.al_install_system(VERSION_INT, nil))
    //  return bool(C.algo_initialize())
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
func (self *Path) NumComponents() int {
    return int(C.al_get_path_num_components(self.handle))
}

// Gets the index'th component of the path name
func (path *Path) Component(index int) string {
    return C.GoString(C.al_get_path_component(path.handle, C.int(index)))
}


// Converts the allegro path to a string 
func (self *Path) String() string {
    return gostr(C.al_path_cstr(self.handle, C.char(NATIVE_PATH_SEP)))
}


func (path * Path) MakeCanonical() * Path {
    C.al_make_path_canonical(path.handle)
    return path
}


/*
func (self * Path) 


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

const (
    RESOURCES_PATH          = C.ALLEGRO_RESOURCES_PATH
    TEMP_PATH               = C.ALLEGRO_TEMP_PATH
    USER_DATA_PATH          = C.ALLEGRO_USER_DATA_PATH
    USER_HOME_PATH          = C.ALLEGRO_USER_HOME_PATH
    USER_SETTINGS_PATH      = C.ALLEGRO_USER_SETTINGS_PATH
    USER_DOCUMENTS_PATH     = C.ALLEGRO_USER_DOCUMENTS_PATH
    EXENAME_PATH            = C.ALLEGRO_EXENAME_PATH
    LAST_PATH               = C.ALLEGRO_LAST_PATH
)

// Gets a standard path location.
func StandardPath(id int) *Path {
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

// Gets the name of the organisation for an app
func OrgName() string {
    return gostr(C.al_get_org_name())
}

// Gets the name of the app
func AppName() string {
    return gostr(C.al_get_app_name())
}

// Inibits the screensaver, or not debending on inhibit.
func InhibitScreensaver(inhibit bool) bool {
    return bool(C.al_inhibit_screensaver(C.bool(inhibit)))
}

// Might be needed on OSX.
func RunMain(args []string, callback RunMainCallback, data interface{}) int {
    runMain.fn   = callback
    runMain.data = data
    argc, argv := CStrings(args) ; defer CStringsFree(argc, argv)
    return int(C.al_run_main(argc, argv, (*[0]byte)(C.go_run_main_callback)))
}

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

// Gets the time since allegro was initialized in seconds 
func Time() float64 {
    return float64(C.al_get_time())
}

// Sleeps the given amount of seconds
func Rest(seconds float64) {
    C.al_rest(C.double(seconds))
}


type Timeout = C.ALLEGRO_TIMEOUT

func (tout * Timeout) Init(seconds float64) {
    C.al_init_timeout(tout, C.double(seconds))
}

func (tout * Timeout) toC() (* C.ALLEGRO_TIMEOUT) {
    return (* C.ALLEGRO_TIMEOUT)(tout)
}

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

// Creates a timer with the given tick speed.
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
func (self *Timer) Speed() float64 {
    return float64(C.al_get_timer_speed(self.handle))
}

// Gets the count (in ticks) of the timer
func (self *Timer) Count() int {
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
func TimerEventSource(self * Timer) *EventSource {
    return (*EventSource)(C.al_get_timer_event_source(self.handle))
}

// Do nothing function for benchmarking only
func DoNothing() {
}

// Returns Allehro's error number
func Errno() int {
    return int(C.al_get_errno())
}
