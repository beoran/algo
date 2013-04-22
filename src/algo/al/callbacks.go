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

var CallbackInt func() int = nil

// generic function pointer caller

//export go_generic_callback_int
func go_generic_callback_int() int {
	if CallbackInt != nil {
		return CallbackInt()
	}
	return 0
}
