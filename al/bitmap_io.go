// albitmap
package al

/*
#include <stdlib.h>
#include <allegro5/allegro.h>
#include "helpers.h"
#include "callbacks.h"
*/
import "C"

const KEEP_BITMAP_FORMAT        = C.ALLEGRO_KEEP_BITMAP_FORMAT 
const NO_PREMULTIPLIED_ALPHA    = C.ALLEGRO_NO_PREMULTIPLIED_ALPHA
const KEEP_INDEX                = C.ALLEGRO_KEEP_INDEX

func LoadBitmap(name string) * Bitmap {
    cname := cstr(name)
    defer cstrFree(cname)
    return wrapBitmap(C.al_load_bitmap(cname))
}

func (bmp * Bitmap) Save(name string) bool {
    cname := cstr(name)
    defer cstrFree(cname)
    return bool(C.al_save_bitmap(cname, bmp.toC()))
}


