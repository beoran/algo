package al

/*
#include <stdlib.h>
#include <allegro5/allegro.h>
#include "helpers.h"
*/
import "C"

import "runtime"

// Usful regexp for KATE:  ALLEGRO_([A-Z0-9_]+)(.*) -> \1 = C.ALLEGRO_\1

// Display functions.

// Possible bit combinations for the flags parameter of CreateDisplay.
const (
    WINDOWED                  = C.ALLEGRO_WINDOWED
    FULLSCREEN                = C.ALLEGRO_FULLSCREEN
    OPENGL                    = C.ALLEGRO_OPENGL
    DIRECT3D_INTERNAL         = C.ALLEGRO_DIRECT3D_INTERNAL
    RESIZABLE                 = C.ALLEGRO_RESIZABLE
    FRAMELESS                 = C.ALLEGRO_FRAMELESS
    NOFRAME                   = C.ALLEGRO_NOFRAME
    GENERATE_EXPOSE_EVENTS    = C.ALLEGRO_GENERATE_EXPOSE_EVENTS
    OPENGL_3_0                = C.ALLEGRO_OPENGL_3_0
    OPENGL_FORWARD_COMPATIBLE = C.ALLEGRO_OPENGL_FORWARD_COMPATIBLE
    FULLSCREEN_WINDOW         = C.ALLEGRO_FULLSCREEN_WINDOW
    MINIMIZED                 = C.ALLEGRO_MINIMIZED
)

/* Possible parameters for SetDisplayOption. */
const (
    RED_SIZE               = C.ALLEGRO_RED_SIZE
    GREEN_SIZE             = C.ALLEGRO_GREEN_SIZE
    BLUE_SIZE              = C.ALLEGRO_BLUE_SIZE
    ALPHA_SIZE             = C.ALLEGRO_ALPHA_SIZE
    RED_SHIFT              = C.ALLEGRO_RED_SHIFT
    GREEN_SHIFT            = C.ALLEGRO_GREEN_SHIFT
    BLUE_SHIFT             = C.ALLEGRO_BLUE_SHIFT
    ALPHA_SHIFT            = C.ALLEGRO_ALPHA_SHIFT
    ACC_RED_SIZE           = C.ALLEGRO_ACC_RED_SIZE
    ACC_GREEN_SIZE         = C.ALLEGRO_ACC_GREEN_SIZE
    ACC_BLUE_SIZE          = C.ALLEGRO_ACC_BLUE_SIZE
    ACC_ALPHA_SIZE         = C.ALLEGRO_ACC_ALPHA_SIZE
    STEREO                 = C.ALLEGRO_STEREO
    AUX_BUFFERS            = C.ALLEGRO_AUX_BUFFERS
    COLOR_SIZE             = C.ALLEGRO_COLOR_SIZE
    DEPTH_SIZE             = C.ALLEGRO_DEPTH_SIZE
    STENCIL_SIZE           = C.ALLEGRO_STENCIL_SIZE
    SAMPLE_BUFFERS         = C.ALLEGRO_SAMPLE_BUFFERS
    SAMPLES                = C.ALLEGRO_SAMPLES
    RENDER_METHOD          = C.ALLEGRO_RENDER_METHOD
    FLOAT_COLOR            = C.ALLEGRO_FLOAT_COLOR
    FLOAT_DEPTH            = C.ALLEGRO_FLOAT_DEPTH
    SINGLE_BUFFER          = C.ALLEGRO_SINGLE_BUFFER
    SWAP_METHOD            = C.ALLEGRO_SWAP_METHOD
    COMPATIBLE_DISPLAY     = C.ALLEGRO_COMPATIBLE_DISPLAY
    UPDATE_DISPLAY_REGION  = C.ALLEGRO_UPDATE_DISPLAY_REGION
    VSYNC                  = C.ALLEGRO_VSYNC
    MAX_BITMAP_SIZE        = C.ALLEGRO_MAX_BITMAP_SIZE
    SUPPORT_NPOT_BITMAP    = C.ALLEGRO_SUPPORT_NPOT_BITMAP
    CAN_DRAW_INTO_BITMAP   = C.ALLEGRO_CAN_DRAW_INTO_BITMAP
    SUPPORT_SEPARATE_ALPHA = C.ALLEGRO_SUPPORT_SEPARATE_ALPHA
    DISPLAY_OPTIONS_COUNT  = C.ALLEGRO_DISPLAY_OPTIONS_COUNT
)

// Constants that determine if a setting is required or not.
const (
    DONTCARE = C.ALLEGRO_DONTCARE
    REQUIRE  = C.ALLEGRO_REQUIRE
    SUGGEST  = C.ALLEGRO_SUGGEST
)

// Display orientations
const (
    DISPLAY_ORIENTATION_0_DEGREES   = C.ALLEGRO_DISPLAY_ORIENTATION_0_DEGREES
    DISPLAY_ORIENTATION_90_DEGREES  = C.ALLEGRO_DISPLAY_ORIENTATION_90_DEGREES
    DISPLAY_ORIENTATION_180_DEGREES = C.ALLEGRO_DISPLAY_ORIENTATION_180_DEGREES
    DISPLAY_ORIENTATION_270_DEGREES = C.ALLEGRO_DISPLAY_ORIENTATION_270_DEGREES
    DISPLAY_ORIENTATION_FACE_UP     = C.ALLEGRO_DISPLAY_ORIENTATION_FACE_UP
    DISPLAY_ORIENTATION_FACE_DOWN   = C.ALLEGRO_DISPLAY_ORIENTATION_FACE_DOWN
)

// Type that wraps a Display (a main window)
type Display struct {
    handle *C.ALLEGRO_DISPLAY
}

// Converts display to C display 
func (disp *Display) toC() *C.ALLEGRO_DISPLAY {
    return disp.handle 
}


// Destroys a display. Use this only when really needed!
func (self *Display) Destroy() {
    if self.handle != nil {
        C.al_destroy_display(self.handle)
    }
    self.handle = nil
}

// Wraps a C Allegro display in a Display. Sets no finalizer.
func wrapDisplayRaw(handle *C.ALLEGRO_DISPLAY) *Display {
    if handle == nil {
        return nil
    }
    return &Display{handle}
}

// Wraps a C Allegro display in a Display. Sets a finalizer that calls Destroy
func wrapDisplay(handle *C.ALLEGRO_DISPLAY) *Display {
    self := wrapDisplayRaw(handle)
    if self != nil {
        runtime.SetFinalizer(self, func(me *Display) { me.Destroy() })
    }
    return self
}

// Display mode info.
type DisplayMode C.ALLEGRO_DISPLAY_MODE

// Converts display mode to C display mode 
func (self *DisplayMode) toC() *C.ALLEGRO_DISPLAY_MODE {
    return (*C.ALLEGRO_DISPLAY_MODE)(self)
}

// Returns the width of the display mode self.
func (self *DisplayMode) Width() int {
    return int(self.width)
}

// Returns the height of the display mode self.
func (self *DisplayMode) Height() int {
    return int(self.height)
}

// Returns the format of the display mode self.
func (self *DisplayMode) Format() int {
    return int(self.format)
}

// Returns the refresh rate of the display mode self.
func (self *DisplayMode) RefreshRate() int {
    return int(self.refresh_rate)
}

// Monitor info
type MonitorInfo C.ALLEGRO_MONITOR_INFO

// Returns the X1 of the monitor info self.
func (self *MonitorInfo) X1() int {
    return int(self.x1)
}

// Returns the Y1 of the monitor info self.
func (self *MonitorInfo) Y1() int {
    return int(self.y1)
}

// Returns the X2 of the monitor info self.
func (self *MonitorInfo) X2() int {
    return int(self.x2)
}

// Returns the Y2 of the monitor info self.
func (self *MonitorInfo) Y2() int {
    return int(self.y2)
}

const (
    DEFAULT_DISPLAY_ADAPTER = C.ALLEGRO_DEFAULT_DISPLAY_ADAPTER
)

// Sets the flags that a display created by CreateDisplay will get after
// this function was called.
func SetNewDisplayFlags(flags int) {
    C.al_set_new_display_flags(C.int(flags))
}

// Creates a new dosplay with the given size. Influenced by SetNewDisplayFlags.
func CreateDisplay(width, height int) *Display {
    return wrapDisplay(C.al_create_display(C.int(width), C.int(height)))
}

// Resizes the display.
func (self *Display) Resize(width, height int) bool {
    return bool(C.al_resize_display(self.handle, C.int(width), C.int(height)))
}

// Updates the display to the physical scree no any changes become visible
func FlipDisplay() {
    C.al_flip_display()
}

// Same as FlipDisplay, for mere consistency
func (self *Display) Flip() {
    C.al_flip_display()
}

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

// Fills the current active display with a color
func ClearToColor(color Color) {
    C.al_clear_to_color(color.toC())
}

// Draws a pixel on the active display at the given location 
// with the given color
func DrawPixel(x, y float32, color Color) {
    C.al_draw_pixel(C.float(x), C.float(y), C.ALLEGRO_COLOR(color))
}

// Sets the refresh rate that the display should have after CreateDisplay().
func SetNewDisplayRefreshRate(refresh_rate int) {
    C.al_set_new_display_refresh_rate(C.int(refresh_rate))
}

// Gets the refresh rate that the display should have after CreateDisplay().
func NewDisplayRefreshRate() int {
    return int(C.al_get_new_display_refresh_rate())
}

// Gets the display flags that the display should have after CreateDisplay().
func NewDisplayFlags() int {
    return int(C.al_get_new_display_flags())
}

// Gets the width of the display  in pixels
func (self *Display) Width() int {
    return int(C.al_get_display_width(self.handle))
}

// Gets the height of the display in pixels
func (self *Display) Height() int {
    return int(C.al_get_display_height(self.handle))
}

// Gets the refresh rate of the display
func (self *Display) RefreshRate() int {
    return int(C.al_get_display_refresh_rate(self.handle))
}

// Gets the display flags of the display
func (self *Display) DisplayFlags() int {
    return int(C.al_get_display_flags(self.handle))
}

// Sets a dispay flag on the display
func (self *Display) SetDisplayFlag(flag int, onoff bool) bool {
    return cb2b(C.al_set_display_flag(self.handle, C.int(flag), b2cb(onoff)))
}

// Returns the current display 
func CurrentDisplay() *Display {
    return wrapDisplayRaw(C.al_get_current_display())
}

// Sets the target C bitmap of allegro drawing 
func setTargetCBitmap(bmp *C.ALLEGRO_BITMAP) {
    C.al_set_target_bitmap(bmp)
}

// Sets the target bitmap of the allegro drawing
func SetTargetBitmap(bmp Bitmap) {
    setTargetCBitmap(bmp.handle)
}

// Sets the target C backbuffer of allegro drawing 
func setTargetCBackbuffer(display *C.ALLEGRO_DISPLAY) {
    C.al_set_target_backbuffer(display)
}

// Sets the target backbuffer of allegro drawing 
func SetTargetBackbuffer(display *Display) {
    setTargetCBackbuffer(display.handle)
}

// Gets the backbuffer bitmap of the display 
func (self *Display) Backbuffer() *Bitmap {
    return wrapBitmapRaw(C.al_get_backbuffer(self.handle))
}

// Gets the target bitmap of allegro drawing
func TargetBitmap() *Bitmap {
    return wrapBitmapRaw(C.al_get_target_bitmap())
}

// Must be called to aknowledge a RESIZE event
func (self *Display) AcknowledgeResize() bool {
    return cb2b(C.al_acknowledge_resize(self.handle))
}

// Updates a region of the display (not the whole display like flip does)
func UpdateDisplayRegion(x, y, width, height int) {
    C.al_update_display_region(C.int(x), C.int(y), C.int(width), C.int(height))
}

// Returns true of the bitmap is compatible with the currebt display, false if not. 
func (bitmap *Bitmap) IsCompatibleBitmap() bool {
    return cb2b(C.al_is_compatible_bitmap(bitmap.handle))
}

// Returns the number of display modes available to Allegro
func NumDisplayModes() int {
    return int(C.al_get_num_display_modes())
}

// Returns the index'th display mode. Pass in a DisplayMode struct to store the display
// mode info in. 
func (self *DisplayMode) Get(index int) *DisplayMode {
    return (*DisplayMode)(C.al_get_display_mode(C.int(index), self.toC()))
}

// Gets display mode info for the index'th display mode
func GetDisplayMode(index int) *DisplayMode {
    var mode DisplayMode
    if (&mode).Get(index) != nil {
        return &mode
    }
    return nil
}

// Waits for the vertical retrace of the monitor to lessen tearing.
func WaitForVsync() {
    C.al_wait_for_vsync()
}

// Gets the event source of the display to registeron an event queue 
// with RegisterEventSource.
func (self *Display) GetEventSource() *EventSource {
    return (*EventSource)(C.al_get_display_event_source(self.handle))
}

// Sets the display icon the window manager should use for the display window
func (self *Display) SetDisplayIcon(bitmap *Bitmap) {
    C.al_set_display_icon(self.handle, bitmap.handle)
}

// Gets the number of available video adapters (I.E. grapic cards)
func NumVideoAdapters() int {
    return int(C.al_get_num_video_adapters())
}

// Converts a monitor info pointer to a C  * ALLEGRO_MONITOR_INFO
func (self *MonitorInfo) toC() *C.ALLEGRO_MONITOR_INFO {
    return (*C.ALLEGRO_MONITOR_INFO)(self)
}

// Gets the monitor info for the index'th video adapter
func (self *MonitorInfo) Get(index int) bool {
    return cb2b(C.al_get_monitor_info(C.int(index), self.toC()))
}

// Gets the monitor info for the index'th video adapter
func GetMonitorInfo(index int) *MonitorInfo {
    var info MonitorInfo
    if (&info).Get(index) {
        return &info
    }
    return nil
}

// Returns the number of the display adapter where new dsplays will be created
func NewDisplayAdapter() int {
    return int(C.al_get_new_display_adapter())
}

// Sets the number of the display adapter where new dsplays will be created
func SetNewDisplayAdapter(adapter int) {
    C.al_set_new_display_adapter(C.int(adapter))
}

// Returns the position where new windowed displays will be created
func NewWindowPosition() (x, y int) {
    var cx, cy C.int
    C.al_get_new_window_position(&cx, &cy)
    return int(cx), int(cy)
}

// Sets the position where new windowed displays will be created
func SetNewWindowPosition(x, y int) {
    C.al_set_new_window_position(C.int(x), C.int(y))
}

// Returns the current position of the windowed display
func (self *Display) WindowPosition() (x, y int) {
    var cx, cy C.int
    C.al_get_window_position(self.handle, &cx, &cy)
    return int(cx), int(cy)
}

// Sets the position where new windowed displays will be created
func (self *Display) SetWindowPosition(x, y int) {
    C.al_set_window_position(self.handle, C.int(x), C.int(y))
}

// Sets the title of the windowed display
func (self *Display) SetTitle(str string) {
    cstr := cstr(str)
    defer cstrFree(cstr)
    C.al_set_window_title(self.handle, cstr)
}

// Sets a display option to be used when a new display is created
func SetNewDisplayOption(option, value, importance int) {
    C.al_set_new_display_option(C.int(option), C.int(value), C.int(importance))
}

// Resets all display oprions for new displays to their default values.
func ResetNewDisplayOptions() {
    C.al_reset_new_display_options()
}

// Gets the display option of this display
func (self *Display) DisplayOption(option int) int {
    return int(C.al_get_display_option(self.handle, C.int(option)))
}

// Allows to speed up drawing by holding the display . Only bitmap functions and font 
// drawing, as well as tranformations shouldbe done until the hold is released
func HoldBitmapDrawing(hold bool) {
    C.al_hold_bitmap_drawing(b2cb(hold))
}

// Returns whether or not the bitmap drawing was held
func IsBitmapDrawingHeld() bool {
    return cb2b(C.al_is_bitmap_drawing_held())
}
