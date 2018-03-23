#include <stdio.h>
#include <stdlib.h>
#include <allegro5/allegro.h>
#include "helpers.h"
#include "callbacks.h"
#include "_cgo_export.h"

/* This wraps the C api function that takes a callback, to take go_take_callback_callback
 as the callback */
int go_take_callback(void * context) {
    return take_callback((function_pointer)go_take_callback_callback, context);
}

 


