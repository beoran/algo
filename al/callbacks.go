package al

/*
#include <stdlib.h>
#include <allegro5/allegro.h>
#include "helpers.h"
#include "callbacks.h"
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


// generic function pointer caller

var CallbackInt func() int = nil

//export go_generic_callback_int
func go_generic_callback_int() int {
    if CallbackInt != nil {
        return CallbackInt()
    }
    return 0
}

type CallbackFunction func() int

type callbackData struct {
    fn   CallbackFunction
    data unsafe.Pointer
}

// the function below is the C callback that will be given to the 
// C api that needs a callback. It uses a callback context with a 
// Go function pointer in it to call that go function.

//export go_take_callback_callback
func go_take_callback_callback(context unsafe.Pointer) int {
    cbd := (*callbackData)(context)
    fn := cbd.fn
    return fn()
}

// Finally wrap the C callback caller function.
func TakeCallback(fn CallbackFunction) int {
    ctx := unsafe.Pointer(&callbackData{fn, nil})
    return int(C.go_take_callback(ctx))
}
