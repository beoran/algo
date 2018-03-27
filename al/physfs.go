package al

// Native dialogs extension

/*
#cgo pkg-config: allegro_physfs-5
#cgo CFLAGS: -I/usr/local/include
#cgo linux LDFLAGS: -lc_nonshared
#include <stdlib.h>
#include <allegro5/allegro.h>
#include <allegro5/allegro_physfs.h>
#include "helpers.h"
*/
import "C"

// Sets up to use the PHYSFS helper to easily read data in zip file .
func SetPhysfsFileInterface() {
    C.al_set_physfs_file_interface()
}

func PhysfsVersion() uint32 {
    return uint32(C.al_get_allegro_physfs_version())
}
