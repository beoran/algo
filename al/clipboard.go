// clipboard support
package al

/*
#include <stdlib.h>
#include <allegro5/allegro.h>
#include "helpers.h"
*/
import "C"

func (disp * Display) ClipboardText() string {
    return C.GoString(C.al_get_clipboard_text(disp.toC()))
}

func (disp * Display) ClipboardHasText() bool {
    return bool(C.al_clipboard_has_text(disp.toC()))
}

func (disp * Display) SetClipboardText(text string) bool {
    ctext := cstr(text) ; defer cstrFree(ctext)
    return bool(C.al_set_clipboard_text(disp.toC(), ctext))
}

