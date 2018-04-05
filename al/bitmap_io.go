// albitmap
package al

/*
#include <stdlib.h>
#include <allegro5/allegro.h>
#include "helpers.h"
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

func (file * File) LoadBitmap(ident string) * Bitmap {
    var cident  *C.char = nil
    if ident != "" {
        cident = cstr(ident)
        defer cstrFree(cident)
    }
    
    return wrapBitmap(C.al_load_bitmap_f(file.toC(), cident))
}

func (file * File) SaveBitmap(ident string, bmp * Bitmap) bool {
    var cident  *C.char = nil
    if ident != "" { 
        cident = cstr(ident)
        defer cstrFree(cident)
    }
    
    return bool(C.al_save_bitmap_f(file.toC(), cident, bmp.toC()))
}


