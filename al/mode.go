// acodec
package al

/*
#cgo pkg-config: allegro-5
#cgo CFLAGS: -I/usr/local/include
#cgo linux LDFLAGS: -lc_nonshared
#include <stdlib.h>
#include <allegro5/allegro.h>
#include <allegro5/fullscreen_mode.h>
#include "helpers.h"
*/
import "C"
import "fmt"

type DisplayMode C.ALLEGRO_DISPLAY_MODE 

func (dm * DisplayMode) toC() * C.ALLEGRO_DISPLAY_MODE {
    return (*C.ALLEGRO_DISPLAY_MODE)(dm) 
}

func (dm * DisplayMode) Width() int {
    return int(dm.width)
}

func (dm * DisplayMode) Height() int {
    return int(dm.height)
}

func (dm * DisplayMode) Format() int {
    return int(dm.format)
}

func (dm * DisplayMode) RefreshRate() int {
    return int(dm.refresh_rate)
}

func (dm * DisplayMode) String() string {
    return fmt.Sprintf("(%d %d) %d %d hz", dm.Width(), dm.Height(), dm.Format(), dm.RefreshRate() ) 
}

func NumDisplayModes() int {
    return int(C.al_get_num_display_modes())
}

func FindDisplayMode(index int) (disp * DisplayMode) {
    disp = &DisplayMode{}
    if nil == C.al_get_display_mode(C.int(index), disp.toC()) {
        return nil
    }
    return disp
}

func DisplayModes() (modes [] *DisplayMode) {
    count := NumDisplayModes()
    modes = make([]*DisplayMode, count)
    for i:=0 ; i < count ; i ++ {
        modes[i] = FindDisplayMode(i)
    }
    return modes
}



/* vim: set ts=8 sts=3 sw=3 et: */
