// albitmap
package al

/*
#include <stdlib.h>
#include <allegro5/allegro.h>
#include "helpers.h"
*/
import "C"

import "runtime"
import "unsafe"

// Type that wraps a Bitmap
type Bitmap struct {
    handle *C.ALLEGRO_BITMAP
}

// returns low level handle for the bitmap
func (bmp *Bitmap) toC() * C.ALLEGRO_BITMAP {
    return bmp.handle
}


// Destroys a bitmap. Use this only when really needed!
func (bmp *Bitmap) Destroy() {
    if bmp.handle != nil {
        C.al_destroy_bitmap(bmp.handle)
    }
    bmp.handle = nil
}

// Alias for destoy to make this implement a Closer interface
func (bmp *Bitmap) Close() {
    bmp.Destroy()
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
    bmp := wrapBitmapRaw(handle)
    if bmp != nil {
        runtime.SetFinalizer(bmp, func(me *Bitmap) { me.Destroy() })
    }
    return bmp
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
func NewBitmapFormat() int {
    return int(C.al_get_new_bitmap_format())
}

// Gets the flags for new bitmaps that are created using CreateBitmap
func NewBitmapFlags() int {
    return int(C.al_get_new_bitmap_flags())
}

// Gets the width of the bitmap.
func (bmp *Bitmap) Width() int {
    return int(C.al_get_bitmap_width(bmp.handle))
}

// Gets the height of the bitmap.
func (bmp *Bitmap) Height() int {
    return int(C.al_get_bitmap_height(bmp.handle))
}

// Gets the width of the bitmap as a float32.
func (bmp *Bitmap) Widthf() float32 {
    return float32(bmp.Width())
}

// Gets the height of the bitmap as a float32.
func (bmp *Bitmap) Heightf() float32 {
    return float32(bmp.Height())
}



// Gets the format of the bitmap.
func (bmp *Bitmap) Format() int {
    return int(C.al_get_bitmap_format(bmp.handle))
}

// Gets the flags of the bitmap.
func (bmp *Bitmap) Flags() int {
    return int(C.al_get_bitmap_flags(bmp.handle))
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

// Puts a pixel to the currently active bitmap
func PutPixel(x, y int, color Color) {
    C.al_put_pixel(C.int(x), C.int(y), color.toC())
}

// Blends a pixel to the currently active bitmap
func PutBlendedPixel(x, y int, color Color) {
    C.al_put_blended_pixel(C.int(x), C.int(y), color.toC())
}

// Gets a pixel from the bitmap
func (bmp *Bitmap) Pixel(x, y int) (color Color) {
    return wrapColor(C.al_get_pixel(bmp.handle, C.int(x), C.int(y)))
}

// Converts pixels of the mask color to transparent pixels (with an alpha channel) 
// for the given bitmap. Useful for, say "magic pink" backgrounds.
func (bmp *Bitmap) ConvertMaskToAlpha(mask_color Color) {
    C.al_convert_mask_to_alpha(bmp.handle, mask_color.toC())
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
func (bmp *Bitmap) CreateSubBitmapRaw(x, y, w, h int) *Bitmap {
    return wrapBitmapRaw(C.al_create_sub_bitmap(bmp.handle,
        C.int(x), C.int(y), C.int(w), C.int(h)))
}

// Creates a sub bitmap of the given bitmap that will automatically be destoyed
// through a finalizer. However, you must ensure that this destroy happens before the 
// parent bitmap is disposed of. You may need to call Destroy manyally anyway. 
func (bmp *Bitmap) CreateSubBitmap(x, y, w, h int) *Bitmap {
    return wrapBitmap(C.al_create_sub_bitmap(bmp.handle,
        C.int(x), C.int(y), C.int(w), C.int(h)))
}

// Returns whether or not the bitmap is a sub bitmap 
func (bmp *Bitmap) IsSubBitmap() bool {
    return cb2b(C.al_is_sub_bitmap(bmp.handle))
}

// Returns the parent bitmap of this sub bitmap, or nil if there is no parent. 
// This is a raw bitmap that has no finaliser set on it since likely 
/// this function will only be used for inspection.
func (bmp *Bitmap) Parent() *Bitmap {
    return wrapBitmapRaw(C.al_get_parent_bitmap(bmp.handle))
}

// Gets the Returns the X position within the parent of a sub bitmap
func (bmp *Bitmap) SubX() int {
    return int(C.al_get_bitmap_x(bmp.handle))
}
// Gets the Returns the Y position within the parent of a sub bitmap
func (bmp *Bitmap) SubY() int {
    return int(C.al_get_bitmap_y(bmp.handle))
}

// Changes the parent, size and position of a sub bitmap 
func (bmp *Bitmap) Reparent(parent * Bitmap, x, y, w, h int) {
    C.al_reparent_bitmap(bmp.handle, parent.handle, C.int(x), C.int(y), C.int(w), C.int(h))
}

// Returns a raw clone of the bitmap, that will not be automatically 
// destroyed.
func (bmp *Bitmap) CloneRaw() *Bitmap {
    return wrapBitmapRaw(C.al_clone_bitmap(bmp.handle))
}

// Returns a clone of the bitmap, that will automatically be
// destroyed.
func (bmp *Bitmap) Clone() *Bitmap {
    return wrapBitmap(C.al_clone_bitmap(bmp.handle))
}

// Converts the bitmap to the current screen format, to ensure the blitting goes fast
func (bmp *Bitmap) Convert() {
    C.al_convert_bitmap(bmp.handle)
}

// Converts all known unconverted memory bitmaps to the current screen format, 
// to ensure the blitting goes fast. 
func ConvertMemoryBitmaps() {
    C.al_convert_memory_bitmaps()
}

const (
    LOCK_READWRITE = C.ALLEGRO_LOCK_READWRITE
    LOCK_READONLY  = C.ALLEGRO_LOCK_READONLY
    LOCK_WRITEONLY = C.ALLEGRO_LOCK_WRITEONLY
)

type LockedRegion = C.ALLEGRO_LOCKED_REGION

// TODO: Provide better access to the data member if needed
func (lk * LockedRegion) dataPointer() unsafe.Pointer {
    return unsafe.Pointer(lk.data)
}

func (lk * LockedRegion) Format() int {
    return int(lk.format)
}

func (lk * LockedRegion) Pitch() int {
    return int(lk.pitch)
}

func (lk * LockedRegion) PixelSize() int {
    return int(lk.pixel_size)
}

func (bmp * Bitmap) Lock(format, flags int) *LockedRegion {
    return (*LockedRegion)(C.al_lock_bitmap(bmp.handle, C.int(format), C.int(flags)))
}

func (bmp * Bitmap) LockRegion(x, y, width, height, format, flags int) *LockedRegion {
    return (*LockedRegion)(C.al_lock_bitmap_region(bmp.handle, 
    C.int(x), C.int(y), C.int(width), C.int(height), C.int(format), C.int(flags)))
}

func (bmp * Bitmap) LockBlocked(flags int) *LockedRegion {
    return (*LockedRegion)(C.al_lock_bitmap_blocked(bmp.handle, C.int(flags)))
}

func (bmp * Bitmap) LockRegionBlocked(x, y, width, height, flags int) *LockedRegion {
    return (*LockedRegion)(C.al_lock_bitmap_region_blocked(bmp.handle, 
    C.int(x), C.int(y), C.int(width), C.int(height), C.int(flags)))
}

func (bmp * Bitmap) Unlock(){
    C.al_unlock_bitmap(bmp.handle)
}

func (bmp * Bitmap) Locked() bool {
    return bool(C.al_is_bitmap_locked(bmp.handle))
}
