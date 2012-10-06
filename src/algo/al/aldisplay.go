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
   WINDOWED = C.ALLEGRO_WINDOWED
   FULLSCREEN = C.ALLEGRO_FULLSCREEN
   OPENGL = C.ALLEGRO_OPENGL
   DIRECT3D_INTERNAL = C.ALLEGRO_DIRECT3D_INTERNAL
   RESIZABLE = C.ALLEGRO_RESIZABLE
   FRAMELESS = C.ALLEGRO_FRAMELESS
   NOFRAME = C.ALLEGRO_NOFRAME
   GENERATE_EXPOSE_EVENTS = C.ALLEGRO_GENERATE_EXPOSE_EVENTS
   OPENGL_3_0 = C.ALLEGRO_OPENGL_3_0
   OPENGL_FORWARD_COMPATIBLE = C.ALLEGRO_OPENGL_FORWARD_COMPATIBLE
   FULLSCREEN_WINDOW = C.ALLEGRO_FULLSCREEN_WINDOW
   MINIMIZED = C.ALLEGRO_MINIMIZED
)

/* Possible parameters for SetDisplayOption. */
const (
   RED_SIZE = C.ALLEGRO_RED_SIZE
   GREEN_SIZE = C.ALLEGRO_GREEN_SIZE
   BLUE_SIZE = C.ALLEGRO_BLUE_SIZE
   ALPHA_SIZE = C.ALLEGRO_ALPHA_SIZE
   RED_SHIFT = C.ALLEGRO_RED_SHIFT
   GREEN_SHIFT = C.ALLEGRO_GREEN_SHIFT
   BLUE_SHIFT = C.ALLEGRO_BLUE_SHIFT
   ALPHA_SHIFT = C.ALLEGRO_ALPHA_SHIFT
   ACC_RED_SIZE = C.ALLEGRO_ACC_RED_SIZE
   ACC_GREEN_SIZE = C.ALLEGRO_ACC_GREEN_SIZE
   ACC_BLUE_SIZE = C.ALLEGRO_ACC_BLUE_SIZE
   ACC_ALPHA_SIZE = C.ALLEGRO_ACC_ALPHA_SIZE
   STEREO = C.ALLEGRO_STEREO
   AUX_BUFFERS = C.ALLEGRO_AUX_BUFFERS
   COLOR_SIZE = C.ALLEGRO_COLOR_SIZE
   DEPTH_SIZE = C.ALLEGRO_DEPTH_SIZE
   STENCIL_SIZE = C.ALLEGRO_STENCIL_SIZE
   SAMPLE_BUFFERS = C.ALLEGRO_SAMPLE_BUFFERS
   SAMPLES = C.ALLEGRO_SAMPLES
   RENDER_METHOD = C.ALLEGRO_RENDER_METHOD
   FLOAT_COLOR = C.ALLEGRO_FLOAT_COLOR
   FLOAT_DEPTH = C.ALLEGRO_FLOAT_DEPTH
   SINGLE_BUFFER = C.ALLEGRO_SINGLE_BUFFER
   SWAP_METHOD = C.ALLEGRO_SWAP_METHOD
   COMPATIBLE_DISPLAY = C.ALLEGRO_COMPATIBLE_DISPLAY
   UPDATE_DISPLAY_REGION = C.ALLEGRO_UPDATE_DISPLAY_REGION
   VSYNC = C.ALLEGRO_VSYNC
   MAX_BITMAP_SIZE = C.ALLEGRO_MAX_BITMAP_SIZE
   SUPPORT_NPOT_BITMAP = C.ALLEGRO_SUPPORT_NPOT_BITMAP
   CAN_DRAW_INTO_BITMAP = C.ALLEGRO_CAN_DRAW_INTO_BITMAP
   SUPPORT_SEPARATE_ALPHA = C.ALLEGRO_SUPPORT_SEPARATE_ALPHA
   DISPLAY_OPTIONS_COUNT = C.ALLEGRO_DISPLAY_OPTIONS_COUNT
)

// Constants thatdetermine if a setting is required or not.
const (
   DONTCARE = C.ALLEGRO_DONTCARE
   REQUIRE = C.ALLEGRO_REQUIRE
   SUGGEST = C.ALLEGRO_SUGGEST
)

// Display orientations
const (
   DISPLAY_ORIENTATION_0_DEGREES = C.ALLEGRO_DISPLAY_ORIENTATION_0_DEGREES
   DISPLAY_ORIENTATION_90_DEGREES = C.ALLEGRO_DISPLAY_ORIENTATION_90_DEGREES
   DISPLAY_ORIENTATION_180_DEGREES = C.ALLEGRO_DISPLAY_ORIENTATION_180_DEGREES
   DISPLAY_ORIENTATION_270_DEGREES = C.ALLEGRO_DISPLAY_ORIENTATION_270_DEGREES
   DISPLAY_ORIENTATION_FACE_UP = C.ALLEGRO_DISPLAY_ORIENTATION_FACE_UP
   DISPLAY_ORIENTATION_FACE_DOWN = C.ALLEGRO_DISPLAY_ORIENTATION_FACE_DOWN
)

// Type that wraps a Display (a main window)
type Display struct {
  handle * C.ALLEGRO_DISPLAY
}

// Destroys a display. Use this only when really needed!
func (self * Display) Destroy() {
  if(self.handle != nil) {
    C.al_destroy_display(self.handle)
  }
  self.handle = nil
}

// Wraps a C Allegro display in a Display. Sets no finalizer.
func wrapDisplayRaw(handle * C.ALLEGRO_DISPLAY) * Display {
  if handle == nil { return nil } 
  return &Display{handle}
}

// Wraps a C Allegro display in a Display. Sets a finalizer that calls Destroy
func wrapDisplay(handle * C.ALLEGRO_DISPLAY) * Display {
  self := wrapDisplayRaw(handle)
  if self != nil {
    runtime.SetFinalizer(self, func (me * Display) { me.Destroy() })
  }
  return self
}


// Display mode info.
type DisplayMode C.ALLEGRO_DISPLAY_MODE 

// Returns the width of the display mode self.
func (self * DisplayMode) Width() int {
  return int(self.width)
}

// Returns the height of the display mode self.
func (self * DisplayMode) Height() int {
  return int(self.height)
}

// Returns the format of the display mode self.
func (self * DisplayMode) Format() int {
  return int(self.format)
}

// Returns the refresh rate of the display mode self.
func (self * DisplayMode) RefreshRate() int {
  return int(self.refresh_rate)
}

// Monitor info
type MonitorInfo C.ALLEGRO_MONITOR_INFO

// Returns the X1 of the monitor info self.
func (self * MonitorInfo) X1() int {
  return int(self.x1)
}

// Returns the Y1 of the monitor info self.
func (self * MonitorInfo) Y1() int {
  return int(self.y1)
}

// Returns the X2 of the monitor info self.
func (self * MonitorInfo) X2() int {
  return int(self.x2)
}

// Returns the Y2 of the monitor info self.
func (self * MonitorInfo) Y2() int {
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
func CreateDisplay(width, height int) (*Display) {
  return wrapDisplay(C.al_create_display(C.int(width), C.int(height)))
}

// Resizes the display.
func (self * Display) Resize(width, height int) bool {
  return bool(C.al_resize_display(self.handle, C.int(width), C.int(height)))
}

// Updates the display to the physical scree no any changes become visible
func FlipDisplay() {
  C.al_flip_display()
}

// Same as FlipDisplay, for mere consistency
func (self * Display) Flip() {
  C.al_flip_display()
}


// Color type
type Color C.ALLEGRO_COLOR

// Convert to C
func (self Color) toC() C.ALLEGRO_COLOR {
  return C.ALLEGRO_COLOR(self)
}


// Creates a new color 
func CreateColor(r, g, b, a float32) Color {
  return Color{C.float(r),C.float(g),C.float(b),C.float(a)}
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
  C.al_clear_to_color(color.toC());
}

// Draws a pixel on the active display at the given location 
// with the given color
func DrawPixel(x, y float32, color Color) {
  C.al_draw_pixel(C.float(x), C.float(y), C.ALLEGRO_COLOR(color))
}


/*
AL_FUNC(void, al_set_new_display_refresh_rate, (int refresh_rate));
AL_FUNC(void, al_set_new_display_flags, (int flags));
AL_FUNC(int,  al_get_new_display_refresh_rate, (void));
AL_FUNC(int,  al_get_new_display_flags, (void));

AL_FUNC(int, al_get_display_width,  (ALLEGRO_DISPLAY *display));
AL_FUNC(int, al_get_display_height, (ALLEGRO_DISPLAY *display));
AL_FUNC(int, al_get_display_format, (ALLEGRO_DISPLAY *display));
AL_FUNC(int, al_get_display_refresh_rate, (ALLEGRO_DISPLAY *display));
AL_FUNC(int, al_get_display_flags,  (ALLEGRO_DISPLAY *display));
AL_FUNC(bool, al_set_display_flag, (ALLEGRO_DISPLAY *display, int flag, bool onoff));
AL_FUNC(bool, al_toggle_display_flag, (ALLEGRO_DISPLAY *display, int flag, bool onoff));

AL_FUNC(ALLEGRO_DISPLAY*, al_create_display, (int w, int h));
AL_FUNC(void,             al_destroy_display, (ALLEGRO_DISPLAY *display));
AL_FUNC(ALLEGRO_DISPLAY*, al_get_current_display, (void));
AL_FUNC(void,            al_set_target_bitmap, (ALLEGRO_BITMAP *bitmap));
AL_FUNC(void,            al_set_target_backbuffer, (ALLEGRO_DISPLAY *display));
AL_FUNC(ALLEGRO_BITMAP*, al_get_backbuffer,    (ALLEGRO_DISPLAY *display));
AL_FUNC(ALLEGRO_BITMAP*, al_get_target_bitmap, (void));

AL_FUNC(bool, al_acknowledge_resize, (ALLEGRO_DISPLAY *display));
AL_FUNC(bool, al_resize_display,     (ALLEGRO_DISPLAY *display, int width, int height));
AL_FUNC(void, al_flip_display,       (void));
AL_FUNC(void, al_update_display_region, (int x, int y, int width, int height));
AL_FUNC(bool, al_is_compatible_bitmap, (ALLEGRO_BITMAP *bitmap));

AL_FUNC(int, al_get_num_display_modes, (void));
AL_FUNC(ALLEGRO_DISPLAY_MODE*, al_get_display_mode, (int index,
        ALLEGRO_DISPLAY_MODE *mode));

AL_FUNC(bool, al_wait_for_vsync, (void));

AL_FUNC(ALLEGRO_EVENT_SOURCE *, al_get_display_event_source, (ALLEGRO_DISPLAY *display));


AL_FUNC(void, al_clear_to_color, (ALLEGRO_COLOR color));
AL_FUNC(void, al_draw_pixel, (float x, float y, ALLEGRO_COLOR color));

AL_FUNC(void, al_set_display_icon, (ALLEGRO_DISPLAY *display, ALLEGRO_BITMAP *icon));


AL_FUNC(int, al_get_num_video_adapters, (void));
AL_FUNC(bool, al_get_monitor_info, (int adapter, ALLEGRO_MONITOR_INFO *info));
AL_FUNC(int, al_get_new_display_adapter, (void));
AL_FUNC(void, al_set_new_display_adapter, (int adapter));
AL_FUNC(void, al_set_new_window_position, (int x, int y));
AL_FUNC(void, al_get_new_window_position, (int *x, int *y));
AL_FUNC(void, al_set_window_position, (ALLEGRO_DISPLAY *display, int x, int y));
AL_FUNC(void, al_get_window_position, (ALLEGRO_DISPLAY *display, int *x, int *y));

AL_FUNC(void, al_set_window_title, (ALLEGRO_DISPLAY *display, const char *title));


AL_FUNC(void, al_set_new_display_option, (int option, int value, int importance));
AL_FUNC(int, al_get_new_display_option, (int option, int *importance));
AL_FUNC(void, al_reset_new_display_options, (void));
AL_FUNC(int, al_get_display_option, (ALLEGRO_DISPLAY *display, int option));


AL_FUNC(void, al_hold_bitmap_drawing, (bool hold));
AL_FUNC(bool, al_is_bitmap_drawing_held, (void));

*/


