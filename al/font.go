// Font extension
package al

/*
#cgo pkg-config: allegro_font-5
#cgo CFLAGS: -I/usr/local/include
#cgo linux LDFLAGS: -lc_nonshared
#include <stdlib.h>
#include <allegro5/allegro.h>
#include <allegro5/allegro_font.h>
#include "helpers.h"
*/
import "C"
import "runtime"
import "unsafe"
import "fmt"

type Font struct {
    handle *C.ALLEGRO_FONT
}

const (
    ALIGN_LEFT    = C.ALLEGRO_ALIGN_LEFT
    ALIGN_CENTRE  = C.ALLEGRO_ALIGN_CENTRE
    ALIGN_CENTER  = C.ALLEGRO_ALIGN_CENTER
    ALIGN_RIGHT   = C.ALLEGRO_ALIGN_RIGHT
    ALIGN_INTEGER = C.ALLEGRO_ALIGN_INTEGER
)

// Converts a font to it's underlying C pointer
func (self *Font) toC() *C.ALLEGRO_FONT {
    return (*C.ALLEGRO_FONT)(self.handle)
}

// Destroys the font.
func (self *Font) Destroy() {
    if self.handle != nil {
        C.al_destroy_font(self.toC())
    }
    self.handle = nil
}

// Wraps a C font into a go font
func wrapFontRaw(data *C.ALLEGRO_FONT) *Font {
    if data == nil {
        return nil
    }
    return &Font{data}
}

// Sets up a finalizer for this Font that calls Destroy()
func (self *Font) SetDestroyFinalizer() *Font {
    if self != nil {
        runtime.SetFinalizer(self, func(me *Font) { me.Destroy() })
    }
    return self
}

// Wraps a C voice into a go mixer and sets up a finalizer that calls Destroy()
func wrapFont(data *C.ALLEGRO_FONT) *Font {
    self := wrapFontRaw(data)
    return self.SetDestroyFinalizer()
}

/* 
TODO:
ALLEGRO_FONT_FUNC(bool, al_register_font_loader, (const char *ext, ALLEGRO_FONT *(*load)(const char *filename, int size, int flags)));
*/

func loadBitmapFont(filename string) *C.ALLEGRO_FONT {
    cfilename := cstr(filename)
    defer cstrFree(cfilename)
    return C.al_load_bitmap_font(cfilename)
}

func loadBitmapFontFlags(filename string, flags int) *C.ALLEGRO_FONT {
    cfilename := cstr(filename)
    defer cstrFree(cfilename)
    return C.al_load_bitmap_font_flags(cfilename, C.int(flags))
}

func loadFont(filename string, size, flags int) *C.ALLEGRO_FONT {
    cfilename := cstr(filename)
    defer cstrFree(cfilename)
    return C.al_load_font(cfilename, C.int(size), C.int(flags))
}

func (self *Bitmap) grabFont(ranges []int) *C.ALLEGRO_FONT {
    cn := C.int(len(ranges) / 2)
    cranges := (*C.int)(unsafe.Pointer(&ranges[0]))
    return C.al_grab_font_from_bitmap(self.handle, cn, cranges)
}

func createBuiltinFont() *C.ALLEGRO_FONT {
    return C.al_create_builtin_font()
}

// Loads a font from the give bitmap filename
func LoadBitmapFontRaw(filename string) *Font {
    return wrapFontRaw(loadBitmapFont(filename))
}

// Loads a font from the give bitmap filename
func LoadBitmapFont(filename string) *Font {
    return LoadBitmapFontRaw(filename).SetDestroyFinalizer()
}

// Loads a font from the give bitmap filename with the given flags
func LoadBitmapFontFlagsRaw(filename string, flags int) *Font {
    return wrapFontRaw(loadBitmapFontFlags(filename, flags))
}

// Loads a font from the give bitmap filename with the given flags
func LoadBitmapFontFlags(filename string, flags int) *Font {
    return LoadBitmapFontFlagsRaw(filename, flags).SetDestroyFinalizer()
}

// Loads a font from the give font filename with the given size and flags.
func LoadFontRaw(filename string, size, flags int) *Font {
    return wrapFontRaw(loadFont(filename, size, flags))
}

// Loads a font from the give font filename with the given size and flags.
func LoadFont(filename string, size, flags int) *Font {
    return LoadFontRaw(filename, size, flags).SetDestroyFinalizer()
}

// Converts this bitmap into a font
func (self *Bitmap) GrabFontRaw(ranges []int) *Font {
    return wrapFontRaw(self.grabFont(ranges))
}

// Converts this bitmap into a font
func (self *Bitmap) GrabFont(ranges []int) *Font {
    return self.GrabFontRaw(ranges).SetDestroyFinalizer()
}

// Creates a builtin font. It must be Destroy() when done using it just like any other font. 
func CreateBuiltinFontRaw() *Font {
    return wrapFontRaw(createBuiltinFont())
}

// Creates a builtin font. Has a finalizer set that will call Destroy().
func CreateBuiltinFont() *Font {
    return wrapFont(createBuiltinFont())
}

// Ustr basics 
type Ustr struct {
    handle *C.ALLEGRO_USTR
}

// Converts a USTR to it's underlying C pointer
func (self *Ustr) toC() *C.ALLEGRO_USTR {
    return (*C.ALLEGRO_USTR)(self.handle)
}

// Destroys the USTR.
func (self *Ustr) Destroy() {
    if self.handle != nil {
        C.al_ustr_free(self.toC())
    }
    self.handle = nil
}

// Wraps a C USTR into a go font
func wrapUstrRaw(data *C.ALLEGRO_USTR) *Ustr {
    if data == nil {
        return nil
    }
    return &Ustr{data}
}

// Sets up a finalizer for this Ustr that calls Destroy()
func (self *Ustr) SetDestroyFinalizer() *Ustr {
    if self != nil {
        runtime.SetFinalizer(self, func(me *Ustr) { me.Destroy() })
    }
    return self
}

// Wraps a C Ustr into go Ustr  and sets up a finalizer that calls Destroy()
func wrapUstr(data *C.ALLEGRO_USTR) *Ustr {
    self := wrapUstrRaw(data)
    return self.SetDestroyFinalizer()
}

// Draws an allegro UTF8 string
func (font *Font) DrawUstr(color Color, x, y float32, flags int, ustr *Ustr) {
    C.al_draw_ustr(font.toC(), color.toC(), C.float(x), C.float(y), C.int(flags), ustr.toC())
}

// Draws a C string
func (font *Font) DrawText(color Color, x, y float32, flags int, text string) {
    ctext := cstr(text)
    defer cstrFree(ctext)
    C.al_draw_text(font.toC(), color.toC(), C.float(x), C.float(y), C.int(flags), ctext)
}

// Draws an allegro UTF8 string, justified
func (font *Font) DrawJustifiedUstr(color Color, x1, x2, y, diff float32, flags int, ustr *Ustr) {
    C.al_draw_justified_ustr(font.toC(), color.toC(), cf(x1), cf(x2), cf(y), cf(diff), ci(flags), ustr.toC())
}

// Draws a C string, justified
func (font *Font) DrawJustifiedText(color Color, x1, x2, y, diff float32, flags int, text string) {
    ctext := cstr(text)
    defer cstrFree(ctext)
    C.al_draw_justified_text(font.toC(), color.toC(), cf(x1), cf(x2), cf(y), cf(diff), ci(flags), ctext)
}

// Formats a single line of text
func (font * Font) DrawTextf(color Color, x, y float32, flags int, 
                             format string, args ...interface{}) {
    text := fmt.Sprintf(format, args...)
    font.DrawText(color, x, y, flags, text)
}

// Formats justified text.
func (font * Font) DrawJustifiedTextf(color Color, x1, x2, y, diff float32, flags int, 
                             format string, args ...interface{}) {
    text := fmt.Sprintf(format, args...)
    font.DrawJustifiedText(color, x1, x2, y, diff, flags, text)
}

// Gets the width of a UTF8 encoded string for this font.
func (font *Font) UstrWidth(ustr *Ustr) int {
    return int(C.al_get_ustr_width(font.toC(), ustr.toC()))
}

// Gets the width of a string for this font.
func (font *Font) TextWidth(text string) int {
    ctext := cstr(text)
    defer cstrFree(ctext)
    return int(C.al_get_text_width(font.toC(), ctext))
}

// Gets the line height of this font.
func (font *Font) LineHeight() int {
    return int(C.al_get_font_line_height(font.toC()))
}

// Gets the ascent of this font.
func (font *Font) Ascent() int {
    return int(C.al_get_font_ascent(font.toC()))
}

// Gets the descent of this font.
func (font *Font) Descent() int {
    return int(C.al_get_font_descent(font.toC()))
}

// Gets the dimension of a UTF-8 text in this font. 
func (font *Font) UstrDimension(ustr *Ustr) (bbx, bby, bbw, bbh int) {
    var cbbx, cbby, cbbw, cbbh C.int
    C.al_get_ustr_dimensions(font.toC(), ustr.toC(), &cbbx, &cbby, &cbbw, &cbbh)
    return int(cbbx), int(cbby), int(cbbw), int(cbbh)
}

// Gets the dimension of a text in this font.
func (font *Font) TextDimension(text string) (bbx, bby, bbw, bbh int) {
    ctext := cstr(text)
    defer cstrFree(ctext)
    var cbbx, cbby, cbbw, cbbh C.int
    C.al_get_text_dimensions(font.toC(), ctext, &cbbx, &cbby, &cbbw, &cbbh)
    return int(cbbx), int(cbby), int(cbbw), int(cbbh)
}

// Initializes the font addon 
func InitFontAddon() {
    C.al_init_font_addon()
}

// Close the font addon
func ShutdownFontAddon() {
    C.al_init_font_addon()
}

// Gets the allegro font addon version
func GetAllegroFontVersion() uint32 {
    return (uint32)(C.al_get_allegro_font_version())
}

// Gets the range of characters supported by the font
func (font *Font) Ranges() (ranges []int, count int) {
    count = int(C.al_get_font_ranges(font.toC(), 0, nil))
    ranges = make([]int, count * 2)
    isize := C.sizeof_int
    cranges := C.malloc(C.size_t(isize * count * 2))
    defer C.free(cranges)
    C.al_get_font_ranges(font.toC(), ci(count), (*C.int)(cranges))
    for i := 0 ; i < count * 2; i++ {
        ranges[i] = int(*(*C.int)(unsafe.Pointer(uintptr(cranges) + uintptr(i * isize))))
    }
    return ranges, count
}


// Draws a C string text over multiple lines
func (font *Font) DrawMultilineText(color Color, x, y, max_width, line_height float32, 
                                    flags int, text string) {
    ctext := cstr(text)
    defer cstrFree(ctext)
    C.al_draw_multiline_text(font.toC(), color.toC(), C.float(x), C.float(y), 
                             C.float(max_width), C.float(line_height), C.int(flags), ctext)
}

// Formats text over multiple lines
func (font * Font) DrawMultilineTextf(color Color, x, y, max_width, line_height float32, flags int, 
                             format string, args ...interface{}) {
    text := fmt.Sprintf(format, args...)
    font.DrawMultilineText(color, x, y, max_width, line_height, flags, text)
}

func (font * Font) SetFallbackFont(fallback * Font) {
    C.al_set_fallback_font(font.toC(), fallback.toC())
}

func (font * Font) FallbackFont() (* Font) {
    return wrapFontRaw(C.al_get_fallback_font(font.toC()))
}

/*
The fallback is a hassle, might be better to do this in Go.
ALLEGRO_FONT_FUNC(void, al_do_multiline_text, (const ALLEGRO_FONT *font,
   float max_width, const char *text,
   bool (*cb)(int line_num, const char *line, int size, void *extra),
   void *extra));

*/


