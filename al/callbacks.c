#include <stdio.h>
#include <stdlib.h>
#include <allegro5/allegro.h>
#include "helpers.h"
#include "callbacks.h"
#include "_cgo_export.h"

typedef bool (*al_create_custom_bitmap_upload)(ALLEGRO_BITMAP *bitmap, void *data);

/* Wrapper for al_create_custom_bitmap. */
ALLEGRO_BITMAP * go_create_custom_bitmap(int w, int h, void * context) {	
    return al_create_custom_bitmap(w, h, 
        (al_create_custom_bitmap_upload) go_create_custom_bitmap_callback, context);
}


/* This wraps the C api function that takes a callback, to take go_take_callback_callback
 as the callback */
int go_take_callback(void * context) {
	return take_callback((function_pointer)go_take_callback_callback, context);
}

 


