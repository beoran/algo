// TTF extension
package al

/*
#cgo pkg-config: allegro_ttf-5
#cgo CFLAGS: -I/usr/local/include
#cgo linux LDFLAGS: -lc_nonshared
#include <stdlib.h>
#include <allegro5/allegro.h>
#include <allegro5/allegro_ttf.h>
#include "helpers.h"
*/
import "C"

// import "runtime"
// import "unsafe"
const (
    TTF_NO_KERNING = C.ALLEGRO_TTF_NO_KERNING
    TTF_MONOCHROME = C.ALLEGRO_TTF_MONOCHROME
    TTF_NO_AUTOHINT= C.ALLEGRO_TTF_NO_AUTOHINT
)

func LoadTTFFont(filename string, size, flags int) * Font {
    cfilename := cstr(filename); defer cstrFree(cfilename)
    return wrapFont(C.al_load_ttf_font(cfilename, C.int(size), C.int(flags)))
} 

func (file * File) LoadTTFFont(size, flags int) * Font {
    return wrapFont(C.al_load_ttf_font_f(file.toC(), nil, C.int(size), C.int(flags)))
} 

func LoadTTFFontStretch(filename string, w, h, flags int) * Font {
    cfilename := cstr(filename); defer cstrFree(cfilename)
    return wrapFont(C.al_load_ttf_font_stretch(cfilename,  C.int(w), C.int(h), C.int(flags)))
} 

func (file * File) LoadTTFFontStretch(w, h, flags int) * Font {
    return wrapFont(C.al_load_ttf_font_stretch_f(file.toC(), nil, C.int(w), C.int(h), C.int(flags)))
} 


func InitTTFAddon() bool {
    return bool(C.al_init_ttf_addon())
}

func ShutdownTTFAddon() {
    C.al_shutdown_ttf_addon()
}

func TTFVersion() uint32 {
    return uint32(C.al_get_allegro_ttf_version())
}

