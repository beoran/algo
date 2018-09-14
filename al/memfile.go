// Memfile extension
package al

/*
#cgo pkg-config: allegro_memfile-5
#cgo CFLAGS: -I/usr/local/include
#cgo linux LDFLAGS: -lc_nonshared
#include <stdlib.h>
#include <allegro5/allegro.h>
#include <allegro5/allegro_memfile.h>
#include "helpers.h"
*/
import "C"
import "unsafe"

// Returns the version of the image loading addon.
// Gets the allegro font addon version
func AllegroMemfileVersion() uint32 {
    return (uint32)(C.al_get_allegro_memfile_version())
}

// Opens a memfile. Data can be put in a buffer and given to Allegro.
func OpenMemfile(buffer []byte, mode string) *File {
    cmode := cstr(mode)
    defer cstrFree(cmode)
    csize := C.int64_t(len(buffer))
    cmem := unsafe.Pointer(&buffer[0])
    return wrapFile(C.al_open_memfile(cmem, csize, cmode))
}
