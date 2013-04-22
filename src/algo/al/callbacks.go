package al

/*
#include <stdlib.h>
#include <allegro5/allegro.h>
#include "helpers.h"
*/
import "C"

import "unsafe"

//export go_upload_bitmap
func go_upload_bitmap(bitmap unsafe.Pointer, data unsafe.Pointer) C.bool {
	return false
}
