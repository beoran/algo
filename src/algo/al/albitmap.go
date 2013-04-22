// albitmap
package al

/*
#include <stdlib.h>
#include <allegro5/allegro.h>
#include "helpers.h"
*/
import "C"

import "runtime"

// Type that wraps a Bitmap
type Bitmap struct {
	handle *C.ALLEGRO_BITMAP
}

// Destroys a bitmap. Use this only when really needed!
func (self *Bitmap) Destroy() {
	if self.handle != nil {
		C.al_destroy_bitmap(self.handle)
	}
	self.handle = nil
}

// Wraps a C Allegro bitmap in a Bitmap. Sets no finalizer.
func wrapBitmapRaw(handle *C.ALLEGRO_BITMAP) *Bitmap {
	if handle == nil {
		return nil
	}
	return &Bitmap{handle}
}

// Wraps a C Allegro Bitmap in a Bitmap. Sets a finalizer that calls Destroy.
func wrapBitmap(handle *C.ALLEGRO_BITMAP) *Bitmap {
	self := wrapBitmapRaw(handle)
	if self != nil {
		runtime.SetFinalizer(self, func(me *Bitmap) { me.Destroy() })
	}
	return self
}

const (
	MEMORY_BITMAP       = C.ALLEGRO_MEMORY_BITMAP
	FORCE_LOCKING       = C.ALLEGRO_FORCE_LOCKING
	NO_PRESERVE_TEXTURE = C.ALLEGRO_NO_PRESERVE_TEXTURE
	MIN_LINEAR          = C.ALLEGRO_MIN_LINEAR
	MAG_LINEAR          = C.ALLEGRO_MAG_LINEAR
	MIPMAP              = C.ALLEGRO_MIPMAP
	VIDEO_BITMAP        = C.ALLEGRO_VIDEO_BITMAP
	CONVERT_BITMAP      = C.ALLEGRO_CONVERT_BITMAP
)

// Sets the format for new bitmaps that are created using CreateBitmap
func SetNewBitmapFormat(format int) {
	C.al_set_new_bitmap_format(C.int(format))
}

// Sets the flags for new bitmaps that are created using CreateBitmap
func SetNewBitmapFlags(flags int) {
	C.al_set_new_bitmap_flags(C.int(flags))
}

// Adds a flags to the flags that will be used for new bitmaps that are created 
// using CreateBitmap
func AddNewBitmapFlag(flags int) {
	C.al_add_new_bitmap_flag(C.int(flags))
}

// Gets the format for new bitmaps that are created using CreateBitmap
func NewBitmapFormat(format int) int {
	return int(C.al_get_new_bitmap_format())
}

// Gets the flags for new bitmaps that are created using CreateBitmap
func NewBitmapFlags(flags int) int {
	return int(C.al_get_new_bitmap_flags())
}

// Gets the width of the bitmap.
func (self *Bitmap) Width() int {
	return int(C.al_get_bitmap_width(self.handle))
}

// Gets the height of the bitmap.
func (self *Bitmap) Height() int {
	return int(C.al_get_bitmap_height(self.handle))
}

// Gets the format of the bitmap.
func (self *Bitmap) Format() int {
	return int(C.al_get_bitmap_format(self.handle))
}

// Gets the flags of the bitmap.
func (self *Bitmap) Flags() int {
	return int(C.al_get_bitmap_flags(self.handle))
}

// Creates a new RAW bitmap. It will not be automatically destroyed!
func CreateBitmapRaw(w, h int) *Bitmap {
	return wrapBitmapRaw(C.al_create_bitmap(C.int(w), C.int(h)))
}

// Creates a new bitmap. It has a finalizer in place that will let it be automatically 
// destroyed.
func CreateBitmap(w, h int) *Bitmap {
	return wrapBitmap(C.al_create_bitmap(C.int(w), C.int(h)))
}

/* TODO:
AL_FUNC(ALLEGRO_BITMAP*, al_create_custom_bitmap, (int w, int h, bool (*upload)(ALLEGRO_BITMAP *bitmap, void *data), void *data));
*/

// Puts a pixel to the currently active bitmap
func PutPixel(x, y int, color Color) {
	C.al_put_pixel(C.int(x), C.int(y), color.toC())
}

// Blends a pixel to the currently active bitmap
func PutBlendedPixel(x, y int, color Color) {
	C.al_put_blended_pixel(C.int(x), C.int(y), color.toC())
}

// Gets a pixel from the bitmap
func (self *Bitmap) GetPixel(x, y int) (color Color) {
	return wrapColor(C.al_get_pixel(self.handle, C.int(x), C.int(y)))
}

// Converts pixels of the mask color to transparent pixels (with an alpha channel) 
// for the given bitmap. Useful for, say "magic pink" backgrounds.
func (self *Bitmap) ConvertMaskToAlpha(mask_color Color) {
	C.al_convert_mask_to_alpha(self.handle, mask_color.toC())
}

// Sets the clipping rectangle for the currently active bitmap. Anything drawn outside
// this rectangle will be cut off or "clipped".
func SetClippingRectangle(x, y, w, h int) {
	C.al_set_clipping_rectangle(C.int(x), C.int(y), C.int(w), C.int(h))
}

// Resets the clipping rectangle for the currently active bitmap to the full size. 
func ResetClippingRectangle() {
	C.al_reset_clipping_rectangle()
}

// Gets the clipping rectangle for the currently active bitmap. 
func ClippingRectangle() (x, y, w, h int) {
	var cx, cy, cw, ch C.int
	C.al_get_clipping_rectangle(&cx, &cy, &cw, &ch)
	return int(cx), int(cy), int(cw), int(ch)
}

// Creates a RAW sub bitmap of the given bitmap that must be manually destoyed with 
// Destroy() 
func (self *Bitmap) CreateSubBitmapRaw(x, y, w, h int) *Bitmap {
	return wrapBitmapRaw(C.al_create_sub_bitmap(self.handle,
		C.int(x), C.int(y), C.int(w), C.int(h)))
}

// Creates a sub bitmap of the given bitmap that will automatically be destoyed
// through a finalizer. However, you must ensure that this destroy happens before the 
// parent bitmap is disposed of. You may need to call Destroy manyally anyway. 
func (self *Bitmap) CreateSubBitmap(x, y, w, h int) *Bitmap {
	return wrapBitmap(C.al_create_sub_bitmap(self.handle,
		C.int(x), C.int(y), C.int(w), C.int(h)))
}

// Returns whether or not the bitmap is a sub bitmap 
func (self *Bitmap) IsSubBitmap() bool {
	return cb2b(C.al_is_sub_bitmap(self.handle))
}

// Returns the parent bitmap of this sub bitmap, or nil if there is no parent. 
// This is a raw bitmap that has no finaliser set on it since likely 
/// this function will only be used for inspection.
func (self *Bitmap) Parent() *Bitmap {
	return wrapBitmapRaw(C.al_get_parent_bitmap(self.handle))
}

// Returns a raw clone of the bitmap, that will not be automatically 
// destroyed.
func (self *Bitmap) CloneRaw() *Bitmap {
	return wrapBitmapRaw(C.al_clone_bitmap(self.handle))
}

// Returns a clone of the bitmap, that will automatically be
// destroyed.
func (self *Bitmap) Clone() *Bitmap {
	return wrapBitmap(C.al_clone_bitmap(self.handle))
}

// Converts the bitmap to the current screen format, to ensure the blitting goes fast
func (self *Bitmap) Convert() {
	C.al_convert_bitmap(self.handle)
}

// Converts all known unconverted bitmaps to the current screen format, 
// to ensure the blitting goes fast. 
func ConvertBitmaps() {
	C.al_convert_bitmaps()
}
