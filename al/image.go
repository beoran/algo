// Image extension
package al

/*
#cgo pkg-config: allegro_image-5
#cgo CFLAGS: -I/usr/local/include
#cgo linux LDFLAGS: -lc_nonshared
#include <stdlib.h>
#include <allegro5/allegro.h>
#include <allegro5/allegro_image.h>
#include "helpers.h"
*/
import "C"

// Initializes the image format addon
func InitImageAddon() bool {
    return cb2b(C.al_init_image_addon())
}

// Shuts down the image format addon
func ShutdownImageAddon() {
    C.al_shutdown_image_addon()
}

// Returns the version of the image loading addon.
// Gets the allegro font addon version
func AllegroImageVersion() uint32 {
    return (uint32)(C.al_get_allegro_image_version())
}
