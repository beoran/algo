package al

/*
#include <stdlib.h>
#include <allegro5/allegro.h>
#include "helpers.h"
*/
import "C"

import "unsafe"

// Callback function for CreateCustomBitmap()
type CreateCustomBitmapCallback func(bitmap *Bitmap, data interface{}) bool

// Callback context for CreateCustomBitmap()
type createCustomBitmapContext struct {
    fn   CreateCustomBitmapCallback
    data interface{}
}

//export go_create_custom_bitmap_callback
func go_create_custom_bitmap_callback(ptr unsafe.Pointer, context unsafe.Pointer) C.bool {
    ccbd := (*createCustomBitmapContext)(context)
    cbmp := (*C.ALLEGRO_BITMAP)(ptr)
    bmp := wrapBitmapRaw(cbmp)
    fptr := ccbd.fn
    return b2cb((fptr)(bmp, ccbd.data))
}

// Callback function for RunMain()
type RunMainCallback func(args []string, data interface{}) int

// Callback context for CreateCustomBitmap()
type runMainContext struct {
    fn   RunMainCallback    
    data interface{}
}

var runMain runMainContext

//export go_run_main_callback
func go_run_main_callback(argc C.int, argv ** C.char) C.int {
    args := GoStrings(argc, argv)        
    return C.int(runMain.fn(args, runMain.data))
}

type ThreadCallbackFunction func(*Thread, unsafe.Pointer) unsafe.Pointer

type threadCallbackData struct {
    fn   ThreadCallbackFunction
    data unsafe.Pointer
}

//export go_create_thread_callback
func go_create_thread_callback(thread * C.ALLEGRO_THREAD, arg unsafe.Pointer) unsafe.Pointer {
    cbd := (*threadCallbackData)(arg)
    fn  := cbd.fn
    res := fn(wrapThreadRaw(thread), cbd.data)
    return unsafe.Pointer(res)
}

