package al

/*
#cgo pkg-config: allegro-5.0
#cgo CFLAGS: -I/usr/local/include
#cgo linux LDFLAGS: -lc_nonshared
#include <stdlib.h>
#include <allegro5/allegro.h>
#include "helpers.h"
*/
import "C"

import "unsafe"
import "runtime"

const PI = 3.14159265358979323846


// Allegro library ID calculation.
func AL_ID(a,b,c,d int)  int {
  return (((a)<<24) | ((b)<<16) | ((c)<<8) | (d))
}


const VERSION = 5
const SUB_VERSION = 0
const WIP_VERSION = 7
const RELEASE_NUMBER = 1
const VERSION_STR = "5.0.7"
const DATE_STR   = "2012"
const DATE       = 20120624  /* yyyymmdd */
const VERSION_INT = 
  ((VERSION << 24) | (SUB_VERSION << 16) | 
  (WIP_VERSION << 8) | RELEASE_NUMBER)



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
  return bool(C.algo_initialize())
}

// Cleans up the Allegro system. Needed after calling Initialize.
func Cleanup() {
  C.algo_atexit_cleanup()
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
  handle * C.ALLEGRO_PATH
}

// Wraps an Allegro path into the go struct above, but does not set a finalizer
func wrapPathRaw(handle * C.ALLEGRO_PATH) (* Path) {
  if handle == nil { return nil }
  return &Path{handle}
}  

// Wraps an Allegro path into the go struct above, and sets a finalizer
func wrapPathCleanly(handle * C.ALLEGRO_PATH, 
                     clean func(path * Path) ) (* Path) {
  result := wrapPathRaw(handle);
  if (result == nil) { return result }
  runtime.SetFinalizer(result, clean)
  return result
}

// Wraps an Allegro path into the go struct above, and sets a default finalizer
func wrapPath(handle * C.ALLEGRO_PATH)  (* Path) {
  cleanup := func(path * Path) { path.Destroy() }
  return wrapPathCleanly(handle,cleanup);
}

// Creates an Allegro path.
func CreatePath(str string) *Path {
  cstr := C.CString(str)
  defer C.free(unsafe.Pointer(cstr))
  return wrapPath(C.al_create_path(cstr))
}

// Creates an allegro path for a directory.
func CreatePathForDirectory(str string) *Path {
  cstr := C.CString(str)
  defer C.free(unsafe.Pointer(cstr))
  return wrapPath(C.al_create_path_for_directory(cstr))
}

// Clones an allegro path.
func (self * Path) ClonePath() *Path { 
  return wrapPath(C.al_clone_path(self.handle))
}

// Destroys an Allegro path. It may not be used after this.
// Destroy may be called many times.
func (self * Path) Destroy() {
  if self.handle != nil { C.al_destroy_path(self.handle) }
  self.handle = nil;
}

func (self * Path) GetPathNumComponents() (int) {
  return int(C.al_get_path_num_components(self.handle))
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




/*
func GetStandardPath(int id) string {
  

AL_FUNC(PATH *, al_get_standard_path, (int id));
AL_FUNC(void, al_set_exe_name, (char const *path));

AL_FUNC(void, al_set_org_name, (const char *org_name));
AL_FUNC(void, al_set_app_name, (const char *app_name));
AL_FUNC(const char *, al_get_org_name, (void));
AL_FUNC(const char *, al_get_app_name, (void));

AL_FUNC(bool, al_inhibit_screensaver, (bool inhibit));

*/



  
// AL_FUNC(int, al_run_main, (int argc, char **argv, int (*)(int, char **)));

/** Allegro has it's own string type. While it's nice, it's 
not needed in Go, so I will just wrap the basic conversion functions */

type USTR struct {
  handle * C.ALLEGRO_USTR
}

// Frees an Allegro unicode string.
func (self * USTR) Free() {
  if self.handle != nil { C.al_ustr_free(self.handle) }
  self.handle = nil
}

// Converts an Allegro Unicode string to a Go string 
func (self * USTR) String() string {
  if (self.handle == nil) { return "<destroyed>" }
  return C.GoStringN(C.al_cstr(self.handle), C.int(C.al_ustr_size(self.handle)))
} 

// Wraps an Allegro USTR into the go struct above, but does not set a finalizer
func wrapUSTRRaw(handle * C.ALLEGRO_USTR) (* USTR) {
  if handle == nil { return nil }
  return &USTR{handle}
}  

// Wraps an Allegro path into the go struct above, and sets a finalizer
func wrapUSTRCleanly(handle * C.ALLEGRO_USTR, 
                     clean func(ustr * USTR) ) (* USTR) {
  result := wrapUSTRRaw(handle);
  if (result == nil) { return result }
  runtime.SetFinalizer(result, clean)
  return result
}


// Wraps an Allegro path into the go struct above, and sets a default finalizer
func wrapUSTR(handle * C.ALLEGRO_USTR) (* USTR) {
  cleanup := func(ustr * USTR) { ustr.Free() }
  return wrapUSTRCleanly(handle, cleanup);
}


// Converts a go string to an Allegro Unicode string
func USTRV(str string)  (* USTR) { 
  cstr := C.CString(str)
  defer C.free(unsafe.Pointer(cstr))
  return wrapUSTR(C.al_ustr_new(cstr))
}

// Converts a go string to an Allegro Unicode string
func USTRP(str * string) (* USTR) { 
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



