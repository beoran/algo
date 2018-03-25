// acodec
package al

/*
#cgo pkg-config: allegro_acodec-5
#cgo CFLAGS: -I/usr/local/include
#cgo linux LDFLAGS: -lc_nonshared
#include <stdlib.h>
#include <allegro5/allegro.h>
#include <allegro5/allegro_acodec.h>
#include "helpers.h"
*/
import "C"

// Initializes the audio codec addon of Allegro
func InitAcodecAddon() bool {
    return cb2b(C.al_init_acodec_addon())
}

// Gets the version of the allegro5 audio codec addon 
func AcodecVersion() uint32 {
    return uint32(C.al_get_allegro_acodec_version())
}
