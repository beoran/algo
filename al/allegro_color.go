// Color extension
package al

/*
#cgo pkg-config: allegro_color-5
#cgo CFLAGS: -I/usr/local/include
#cgo linux LDFLAGS: -lc_nonshared
#include <stdlib.h>
#include <allegro5/allegro.h>
#include <allegro5/allegro_color.h>
#include "helpers.h"
*/
import "C"

// import "runtime"
// import "unsafe"

// Gets the version of the color addon
func ColorVersion() uint32 {
	return uint32(C.al_get_allegro_color_version())
}

// Converts HSV color components to RGB 
func HsvToRgb(h, s, v float32) (r, g, b float32) {
	var cr, cg, cb C.float
	C.al_color_hsv_to_rgb(C.float(h), C.float(s), C.float(v), &cr, &cg, &cb)
	return float32(cr), float32(cg), float32(cb)
}

// Converts RGB color components to HSL
func RgbToHsl(r, g, b float32) (h, s, l float32) {
	var ch, cs, cl C.float
	C.al_color_rgb_to_hsl(C.float(r), C.float(g), C.float(b), &ch, &cs, &cl)
	return float32(ch), float32(cs), float32(cl)
}

// Converts RGB color components to HSV
func RgbToHsv(r, g, b float32) (h, s, v float32) {
	var ch, cs, cv C.float
	C.al_color_rgb_to_hsl(C.float(r), C.float(g), C.float(b), &ch, &cs, &cv)
	return float32(ch), float32(cs), float32(cv)
}

// Converts HSL color components to RGB 
func HslToRgb(h, s, l float32) (r, g, b float32) {
	var cr, cg, cb C.float
	C.al_color_hsl_to_rgb(C.float(h), C.float(s), C.float(l), &cr, &cg, &cb)
	return float32(cr), float32(cg), float32(cb)
}

// Converts a X11 color name to RGB 
func NameToRgb(name string) (r, g, b float32) {
	var cr, cg, cb C.float
	cname := cstr(name)
	defer cstrFree(cname)
	C.al_color_name_to_rgb(cname, &cr, &cg, &cb)
	return float32(cr), float32(cg), float32(cb)
}

// Converts RGB color components to an X11 color name 
func RgbToName(r, g, b float32) (name string) {
	return gostr(C.al_color_rgb_to_name(C.float(r), C.float(g), C.float(b)))
}

// Converts RGB color components to CMYK
func RgbToCmyk(r, g, b float32) (c, m, y, k float32) {
	var cc, cm, cy, ck C.float
	C.al_color_rgb_to_cmyk(C.float(r), C.float(g), C.float(b), &cc, &cm, &cy, &ck)
	return float32(cc), float32(cm), float32(cy), float32(ck)
}

// Converts CMYK color components to RGB
func CmykToRgb(c, m, y, k float32) (r, g, b float32) {
	var cr, cg, cb C.float
	C.al_color_cmyk_to_rgb(C.float(c), C.float(m), C.float(y), C.float(k), &cr, &cg, &cb)
	return float32(cr), float32(cg), float32(cb)
}

// Converts RGB color components to YUV
func RgbToYuv(r, g, b float32) (y, u, v float32) {
	var cy, cu, cv C.float
	C.al_color_rgb_to_yuv(C.float(r), C.float(g), C.float(b), &cy, &cu, &cv)
	return float32(cy), float32(cu), float32(cv)
}

// Converts HSL color components to RGB 
func YuvToRgb(y, u, v float32) (r, g, b float32) {
	var cr, cg, cb C.float
	C.al_color_yuv_to_rgb(C.float(y), C.float(u), C.float(v), &cr, &cg, &cb)
	return float32(cr), float32(cg), float32(cb)
}

// Converts RGB color components to HTML notation
func RgbToHtml(r, g, b float32) (html string) {
	chtml := alMalloc(8)
	defer alFree(chtml)
	C.al_color_rgb_to_html(C.float(r), C.float(g), C.float(b), (*C.char)(chtml))
	return gostr((*C.char)(chtml))
}

// Converts HTML notation to RGB color components
// Converts a X11 color name to RGB 
func HtmlToRgb(html string) (r, g, b float32) {
	var cr, cg, cb C.float
	chtml := cstr(html)
	defer cstrFree(chtml)
	C.al_color_name_to_rgb(chtml, &cr, &cg, &cb)
	return float32(cr), float32(cg), float32(cb)
}

// Creates a color from YUV components
func ColorYuv(y, u, v float32) Color {
	return wrapColor(C.al_color_yuv(C.float(y), C.float(u), C.float(v)))
}

// Creates a color from Cmyk components
func ColorCmyk(c, m, y, k float32) Color {
	return wrapColor(C.al_color_cmyk(C.float(c), C.float(m), C.float(y), C.float(k)))
}

// Creates a color from HSL components
func ColorHsl(h, s, l float32) Color {
	return wrapColor(C.al_color_hsl(C.float(h), C.float(s), C.float(l)))
}

// Creates a color from HSV components
func ColorHsv(h, s, v float32) Color {
	return wrapColor(C.al_color_hsv(C.float(h), C.float(s), C.float(v)))
}

// Creates a color from an X11 color name 
func ColorName(s string) Color {
	cs := cstr(s)
	defer cstrFree(cs)
	return wrapColor(C.al_color_name(cs))
}

// Creates a color from HTML notation
func ColorHtml(s string) Color {
	cs := cstr(s)
	defer cstrFree(cs)
	return wrapColor(C.al_color_html(cs))
}
