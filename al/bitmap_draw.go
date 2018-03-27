// albitmap
package al

/*
#include <stdlib.h>
#include <allegro5/allegro.h>
#include "helpers.h"
#include "callbacks.h"
*/
import "C"

const (
    FLIP_HORIZONTAL = C.ALLEGRO_FLIP_HORIZONTAL
    FLIP_VERTICAL   = C.ALLEGRO_FLIP_VERTICAL
)

func (bmp * Bitmap) Draw(dx, dy float32, flags int) {
    C.al_draw_bitmap(bmp.handle, C.float(dx), C.float(dy), C.int(flags))
}
    

func (bmp * Bitmap) DrawRegion(sx, sy, sw, sh, dx, dy float32, flags int) {
    C.al_draw_bitmap_region(bmp.handle,  C.float(sx), C.float(sy),
        C.float(sw), C.float(sh), C.float(dx), C.float(dy), C.int(flags))
}

func (bmp * Bitmap) DrawScaled(sx, sy, sw, sh, dx, dy, dw, dh float32, flags int) {
    C.al_draw_scaled_bitmap(bmp.handle,  C.float(sx), C.float(sy),
        C.float(sw), C.float(sh), C.float(dx), C.float(dy), 
        C.float(dw), C.float(dh), C.int(flags))
}

func (bmp * Bitmap) DrawRotated(cx, cy, dx, dy, angle float32, flags int) {
    C.al_draw_rotated_bitmap(bmp.handle,  C.float(cx), C.float(cy),
        C.float(dx), C.float(dy), C.float(angle), C.int(flags))
}

func (bmp * Bitmap) DrawScaledRotated(cx, cy, dx, dy, xscale, yscale, angle float32, flags int) {
    C.al_draw_scaled_rotated_bitmap(bmp.handle,  C.float(cx), C.float(cy),
        C.float(dx), C.float(dy), 
        C.float(xscale), C.float(yscale), C.float(angle), C.int(flags))
}

func (bmp * Bitmap) DrawTinted(color Color, dx, dy float32, flags int) {
    C.al_draw_tinted_bitmap(bmp.handle, color.toC(), C.float(dx), C.float(dy), C.int(flags))
}
    

func (bmp * Bitmap) DrawTintedRegion(color Color, sx, sy, sw, sh, dx, dy float32, flags int) {
    C.al_draw_tinted_bitmap_region(bmp.handle, color.toC(),  C.float(sx), C.float(sy),
        C.float(sw), C.float(sh), C.float(dx), C.float(dy), C.int(flags))
}

func (bmp * Bitmap) DrawTintedScaled(color Color, sx, sy, sw, sh, dx, dy, dw, dh float32, flags int) {
    C.al_draw_tinted_scaled_bitmap(bmp.handle, color.toC(),  C.float(sx), C.float(sy),
        C.float(sw), C.float(sh), C.float(dx), C.float(dy), 
        C.float(dw), C.float(dh), C.int(flags))
}

func (bmp * Bitmap) DrawTintedRotated(color Color, cx, cy, dx, dy, angle float32, flags int) {
    C.al_draw_tinted_rotated_bitmap(bmp.handle, color.toC(),  C.float(cx), C.float(cy),
        C.float(dx), C.float(dy), C.float(angle), C.int(flags))
}

func (bmp * Bitmap) DrawTintedScaledRotated(color Color, cx, cy, dx, dy, xscale, yscale, angle float32, flags int) {
    C.al_draw_tinted_scaled_rotated_bitmap(bmp.handle, color.toC(),  C.float(cx), C.float(cy),
        C.float(dx), C.float(dy), 
        C.float(xscale), C.float(yscale), C.float(angle), C.int(flags))
}

func (bmp * Bitmap) DrawTintedScaledRotatedRegion(sx, sy, sw, sh float32, color Color, cx, cy, dx, dy, xscale, yscale, angle float32, flags int) {
    C.al_draw_tinted_scaled_rotated_bitmap_region(bmp.handle, 
        C.float(sx), C.float(sy), C.float(sw), C.float(sh),
        color.toC(),  C.float(cx), C.float(cy),
        C.float(dx), C.float(dy), 
        C.float(xscale), C.float(yscale), C.float(angle), C.int(flags))
}

