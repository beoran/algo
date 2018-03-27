// color and color pixel format support
package al

/*
#include <stdlib.h>
#include <allegro5/allegro.h>
#include "helpers.h"
#include "callbacks.h"
*/
import "C"

// Color type
type Color C.ALLEGRO_COLOR

// Convert from
func wrapColor(color C.ALLEGRO_COLOR) Color {
    return Color(color)
}

// Convert to C
func (self Color) toC() C.ALLEGRO_COLOR {
    return C.ALLEGRO_COLOR(self)
}

// Creates a new color 
func CreateColor(r, g, b, a float32) Color {
    return Color{C.float(r), C.float(g), C.float(b), C.float(a)}
}

// Returns the R component of the color self.
func (self Color) R() float32 {
    return float32(self.r)
}

// Returns the G component of the color self.
func (self Color) G() float32 {
    return float32(self.g)
}

// Returns the B component of the color self.
func (self Color) B() float32 {
    return float32(self.b)
}

// Returns the A component of the color self.
func (self Color) A() float32 {
    return float32(self.a)
}

type PixelFormat C.enum_ALLEGRO_PIXEL_FORMAT

const(
   PIXEL_FORMAT_ANY                 = PixelFormat(C.ALLEGRO_PIXEL_FORMAT_ANY                 )
   PIXEL_FORMAT_ANY_NO_ALPHA        = PixelFormat(C.ALLEGRO_PIXEL_FORMAT_ANY_NO_ALPHA        )
   PIXEL_FORMAT_ANY_WITH_ALPHA      = PixelFormat(C.ALLEGRO_PIXEL_FORMAT_ANY_WITH_ALPHA      )
   PIXEL_FORMAT_ANY_15_NO_ALPHA     = PixelFormat(C.ALLEGRO_PIXEL_FORMAT_ANY_15_NO_ALPHA     )
   PIXEL_FORMAT_ANY_16_NO_ALPHA     = PixelFormat(C.ALLEGRO_PIXEL_FORMAT_ANY_16_NO_ALPHA     )
   PIXEL_FORMAT_ANY_16_WITH_ALPHA   = PixelFormat(C.ALLEGRO_PIXEL_FORMAT_ANY_16_WITH_ALPHA   )
   PIXEL_FORMAT_ANY_24_NO_ALPHA     = PixelFormat(C.ALLEGRO_PIXEL_FORMAT_ANY_24_NO_ALPHA     )
   PIXEL_FORMAT_ANY_32_NO_ALPHA     = PixelFormat(C.ALLEGRO_PIXEL_FORMAT_ANY_32_NO_ALPHA     )
   PIXEL_FORMAT_ANY_32_WITH_ALPHA   = PixelFormat(C.ALLEGRO_PIXEL_FORMAT_ANY_32_WITH_ALPHA   )
   PIXEL_FORMAT_ARGB_8888           = PixelFormat(C.ALLEGRO_PIXEL_FORMAT_ARGB_8888           )
   PIXEL_FORMAT_RGBA_8888           = PixelFormat(C.ALLEGRO_PIXEL_FORMAT_RGBA_8888           )
   PIXEL_FORMAT_ARGB_4444           = PixelFormat(C.ALLEGRO_PIXEL_FORMAT_ARGB_4444           )
   PIXEL_FORMAT_RGB_888             = PixelFormat(C.ALLEGRO_PIXEL_FORMAT_RGB_888             )
   PIXEL_FORMAT_RGB_565             = PixelFormat(C.ALLEGRO_PIXEL_FORMAT_RGB_565             )
   PIXEL_FORMAT_RGB_555             = PixelFormat(C.ALLEGRO_PIXEL_FORMAT_RGB_555             )
   PIXEL_FORMAT_RGBA_5551           = PixelFormat(C.ALLEGRO_PIXEL_FORMAT_RGBA_5551           )
   PIXEL_FORMAT_ARGB_1555           = PixelFormat(C.ALLEGRO_PIXEL_FORMAT_ARGB_1555           )
   PIXEL_FORMAT_ABGR_8888           = PixelFormat(C.ALLEGRO_PIXEL_FORMAT_ABGR_8888           )
   PIXEL_FORMAT_XBGR_8888           = PixelFormat(C.ALLEGRO_PIXEL_FORMAT_XBGR_8888           )
   PIXEL_FORMAT_BGR_888             = PixelFormat(C.ALLEGRO_PIXEL_FORMAT_BGR_888             )
   PIXEL_FORMAT_BGR_565             = PixelFormat(C.ALLEGRO_PIXEL_FORMAT_BGR_565             )
   PIXEL_FORMAT_BGR_555             = PixelFormat(C.ALLEGRO_PIXEL_FORMAT_BGR_555             )
   PIXEL_FORMAT_RGBX_8888           = PixelFormat(C.ALLEGRO_PIXEL_FORMAT_RGBX_8888           )
   PIXEL_FORMAT_XRGB_8888           = PixelFormat(C.ALLEGRO_PIXEL_FORMAT_XRGB_8888           )
   PIXEL_FORMAT_ABGR_F32            = PixelFormat(C.ALLEGRO_PIXEL_FORMAT_ABGR_F32            )
   PIXEL_FORMAT_ABGR_8888_LE        = PixelFormat(C.ALLEGRO_PIXEL_FORMAT_ABGR_8888_LE        )
   PIXEL_FORMAT_RGBA_4444           = PixelFormat(C.ALLEGRO_PIXEL_FORMAT_RGBA_4444           )
   PIXEL_FORMAT_SINGLE_CHANNEL_8    = PixelFormat(C.ALLEGRO_PIXEL_FORMAT_SINGLE_CHANNEL_8    )
   PIXEL_FORMAT_COMPRESSED_RGBA_DXT1= PixelFormat(C.ALLEGRO_PIXEL_FORMAT_COMPRESSED_RGBA_DXT1)
   PIXEL_FORMAT_COMPRESSED_RGBA_DXT3= PixelFormat(C.ALLEGRO_PIXEL_FORMAT_COMPRESSED_RGBA_DXT3)
   PIXEL_FORMAT_COMPRESSED_RGBA_DXT5= PixelFormat(C.ALLEGRO_PIXEL_FORMAT_COMPRESSED_RGBA_DXT5)
   NUM_PIXEL_FORMATS                = PixelFormat(C.ALLEGRO_NUM_PIXEL_FORMATS                )
)


func MapRGB(r, g, b uint8) Color {
    return wrapColor(C.al_map_rgb(C.uchar(r), C.uchar(g), C.uchar(b)))
}

func MapRGBA(r, g, b, a uint8) Color {
    return wrapColor(C.al_map_rgba(C.uchar(r), C.uchar(g), C.uchar(b), C.uchar(a)))
}

func PremulRGBA(r, g, b, a uint8) Color {
    return wrapColor(C.al_premul_rgba(C.uchar(r), C.uchar(g), C.uchar(b), C.uchar(a)))
}


func MapRGBF(r, g, b float32) Color {
    return wrapColor(C.al_map_rgb_f(C.float(r), C.float(g), C.float(b)))
}

func MapRGBAF(r, g, b, a float32) Color {
    return wrapColor(C.al_map_rgba_f(C.float(r), C.float(g), C.float(b), C.float(a)))
}

func PremulRGBAF(r, g, b, a float32) Color {
    return wrapColor(C.al_premul_rgba_f(C.float(r), C.float(g), C.float(b), C.float(a)))
}

func (color Color) UnmapRGB() (r, g, b uint8) {
    var cr, cg, cb C.uchar
    C.al_unmap_rgb(color.toC(), &cr, &cg, &cb)
    r = uint8(cr)
    g = uint8(cg)
    b = uint8(cb)
    return r, g, b
}

func (color Color) UnmapRGBA() (r, g, b, a uint8) {
    var cr, cg, cb, ca C.uchar
    C.al_unmap_rgba(color.toC(), &cr, &cg, &cb, &ca)
    r = uint8(cr)
    g = uint8(cg)
    b = uint8(cb)
    a = uint8(ca)
    return r, g, b, a
}

func (color Color) UnmapRGBF() (r, g, b float32) {
    var cr, cg, cb C.float
    C.al_unmap_rgb_f(color.toC(), &cr, &cg, &cb)
    r = float32(cr)
    g = float32(cg)
    b = float32(cb)
    return r, g, b
}

func (color Color) UnmapRGBAF() (r, g, b, a float32) {
    var cr, cg, cb, ca C.float
    C.al_unmap_rgba_f(color.toC(), &cr, &cg, &cb, &ca)
    r = float32(cr)
    g = float32(cg)
    b = float32(cb)
    a = float32(ca)
    return r, g, b, a
}

func (format PixelFormat) PixelSize() int {
    return int(C.al_get_pixel_size(C.int(format)))
} 

func (format PixelFormat) FormatBits() int {
    return int(C.al_get_pixel_format_bits(C.int(format)))
} 

func (format PixelFormat) BlockSize() int {
    return int(C.al_get_pixel_block_size(C.int(format)))
} 

func (format PixelFormat) BlockWidth() int {
    return int(C.al_get_pixel_block_width(C.int(format)))
} 

func (format PixelFormat) BlockHeight() int {
    return int(C.al_get_pixel_block_height(C.int(format)))
} 

