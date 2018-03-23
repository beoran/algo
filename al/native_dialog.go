package al

// Native dialogs extension

/*
#cgo pkg-config: allegro_dialog-5
#cgo CFLAGS: -I/usr/local/include
#cgo linux LDFLAGS: -lc_nonshared
#include <stdlib.h>
#include <stdint.h>
#include <allegro5/allegro.h>
#include <allegro5/allegro_native_dialog.h>
#include "helpers.h"
*/
import "C"

import "unsafe"
import "runtime"

const (
    FILECHOOSER_FILE_MUST_EXIST = 1
    FILECHOOSER_SAVE            = 2
    FILECHOOSER_FOLDER          = 4
    FILECHOOSER_PICTURES        = 8
    FILECHOOSER_SHOW_HIDDEN     = 16
    FILECHOOSER_MULTIPLE        = 32

    MESSAGEBOX_WARN      = 1 << 0
    MESSAGEBOX_ERROR     = 1 << 1
    MESSAGEBOX_OK_CANCEL = 1 << 2
    MESSAGEBOX_YES_NO    = 1 << 3
    MESSAGEBOX_QUESTION  = 1 << 4

    TEXTLOG_NO_CLOSE  = 1 << 0
    TEXTLOG_MONOSPACE = 1 << 1

    EVENT_NATIVE_DIALOG_CLOSE = 600
    EVENT_MENU_CLICK          = 601

    MENU_ITEM_ENABLED  = 0
    MENU_ITEM_CHECKBOX = 1
    MENU_ITEM_CHECKED  = 2
    MENU_ITEM_DISABLED = 4
)

type FileChooser struct {
    handle *C.ALLEGRO_FILECHOOSER
}

// Converts a file chooser to it's underlying C pointer
func (self *FileChooser) toC() *C.ALLEGRO_FILECHOOSER {
    return (*C.ALLEGRO_FILECHOOSER)(self.handle)
}

// Destroys the file chooser.
func (self *FileChooser) Destroy() {
    if self.handle != nil {
        C.al_destroy_native_file_dialog(self.toC())
    }
    self.handle = nil
}

// Wraps a C file chooser into a go file chooser
func wrapFileChooserRaw(data *C.ALLEGRO_FILECHOOSER) *FileChooser {
    if data == nil {
        return nil
    }
    return &FileChooser{data}
}

// Sets up a finalizer for this FileChooser that calls Destroy()
func (self *FileChooser) SetDestroyFinalizer() *FileChooser {
    if self != nil {
        runtime.SetFinalizer(self, func(me *FileChooser) { me.Destroy() })
    }
    return self
}

// Wraps a C voice into a go mixer and sets up a finalizer that calls Destroy()
func wrapFileChooser(data *C.ALLEGRO_FILECHOOSER) *FileChooser {
    self := wrapFileChooserRaw(data)
    return self.SetDestroyFinalizer()
}

type TextLog struct {
    handle *C.ALLEGRO_TEXTLOG
}

// Converts a native_text_log to it's underlying C pointer
func (self *TextLog) toC() *C.ALLEGRO_TEXTLOG {
    return (*C.ALLEGRO_TEXTLOG)(self.handle)
}

// Closes the native_text_log.
func (self *TextLog) Close() {
    if self.handle != nil {
        C.al_close_native_text_log(self.toC())
    }
    self.handle = nil
}

// Wraps a C native_text_log into a go native_text_log
func wrapTextLogRaw(data *C.ALLEGRO_TEXTLOG) *TextLog {
    if data == nil {
        return nil
    }
    return &TextLog{data}
}

// Sets up a finalizer for this TextLog that calls Destroy()
func (self *TextLog) SetDestroyFinalizer() *TextLog {
    if self != nil {
        runtime.SetFinalizer(self, func(me *TextLog) { me.Close() })
    }
    return self
}

// Wraps a C native_text_log into a go native_text_log and sets up a finalizer that calls Destroy()
func wrapTextLog(data *C.ALLEGRO_TEXTLOG) *TextLog {
    self := wrapTextLogRaw(data)
    return self.SetDestroyFinalizer()
}

type Menu struct {
    handle *C.ALLEGRO_MENU
}

// Converts a menu to it's underlying C pointer
func (self *Menu) toC() *C.ALLEGRO_MENU {
    return (*C.ALLEGRO_MENU)(self.handle)
}

// Destroys the menu.
func (self *Menu) Destroy() {
    if self.handle != nil {
        C.al_destroy_menu(self.toC())
    }
    self.handle = nil
}

// Wraps a C menu into a go menu
func wrapMenuRaw(data *C.ALLEGRO_MENU) *Menu {
    if data == nil {
        return nil
    }
    return &Menu{data}
}

// Sets up a finalizer for this Menu that calls Destroy()
func (self *Menu) SetDestroyFinalizer() *Menu {
    if self != nil {
        runtime.SetFinalizer(self, func(me *Menu) { me.Destroy() })
    }
    return self
}

// Wraps a C menu into a go menu and sets up a finalizer that calls Destroy()
func wrapMenu(data *C.ALLEGRO_MENU) *Menu {
    self := wrapMenuRaw(data)
    return self.SetDestroyFinalizer()
}

type MenuInfo C.ALLEGRO_MENU_INFO

func makeMenuInfo(text *string, id, flags int, icon *Bitmap) C.ALLEGRO_MENU_INFO {
    res := C.ALLEGRO_MENU_INFO{}
    if text == nil {
        res.caption = nil
    } else {
        bytes := []byte(*text)
        res.caption = (*C.char)(unsafe.Pointer(&bytes[0]))
    }
    res.id = cui16(id)
    res.flags = ci(flags)
    res.icon = icon.handle
    return res
}

/// Formats a menuinfo element for an element of the menu.
func MakeMenuInfo(text *string, id, flags int, icon *Bitmap) MenuInfo {
    return (MenuInfo)(makeMenuInfo(text, id, flags, icon))
}

// Returns a menuinfo that is a separator
func MenuSeparator() MenuInfo {
    return MakeMenuInfo(nil, -1, 0, nil)
}

// Returns a menuinfo that is the start of the menu
func StartOfMenu(caption string, id int) MenuInfo {
    return MakeMenuInfo(&caption, id, 0, nil)
}

// Returns a menuinfo that is the end of the menu
func EndOfMenu(caption string, id int) MenuInfo {
    return MakeMenuInfo(nil, 0, 0, nil)
}

// Starts the native dialog addon
func InitNativeDialogAddon() bool {
    return cb2b(C.al_init_native_dialog_addon())
}

// Stops the native dialog addon
func ShutdownNativeDialogAddon() {
    C.al_shutdown_native_dialog_addon()
}

// Creates a native file dialog.
func CreateNativeFileDialogRaw(path, title, patterns string, mode int) *FileChooser {
    return nil
    //return wrapFileChooser()
}

/*
TODO:
ALLEGRO_DIALOG_FUNC(ALLEGRO_FILECHOOSER *, al_create_native_file_dialog, (char const *initial_path,
   char const *title, char const *patterns, int mode));
ALLEGRO_DIALOG_FUNC(bool, al_show_native_file_dialog, (ALLEGRO_DISPLAY *display, ALLEGRO_FILECHOOSER *dialog));
ALLEGRO_DIALOG_FUNC(int, al_get_native_file_dialog_count, (const ALLEGRO_FILECHOOSER *dialog));
ALLEGRO_DIALOG_FUNC(const char *, al_get_native_file_dialog_path, (const ALLEGRO_FILECHOOSER *dialog,
   size_t index));
ALLEGRO_DIALOG_FUNC(void, al_destroy_native_file_dialog, (ALLEGRO_FILECHOOSER *dialog));

ALLEGRO_DIALOG_FUNC(int, al_show_native_message_box, (ALLEGRO_DISPLAY *display, char const *title,
   char const *heading, char const *text, char const *buttons, int flags));

ALLEGRO_DIALOG_FUNC(ALLEGRO_TEXTLOG *, al_open_native_text_log, (char const *title, int flags));
ALLEGRO_DIALOG_FUNC(void, al_close_native_text_log, (ALLEGRO_TEXTLOG *textlog));
ALLEGRO_DIALOG_FUNC(void, al_append_native_text_log, (ALLEGRO_TEXTLOG *textlog, char const *format, ...));
ALLEGRO_DIALOG_FUNC(ALLEGRO_EVENT_SOURCE *, al_get_native_text_log_event_source, (ALLEGRO_TEXTLOG *textlog));


ALLEGRO_DIALOG_FUNC(ALLEGRO_MENU *, al_create_menu, (void));
ALLEGRO_DIALOG_FUNC(ALLEGRO_MENU *, al_create_popup_menu, (void));
ALLEGRO_DIALOG_FUNC(ALLEGRO_MENU *, al_build_menu, (ALLEGRO_MENU_INFO *info));
ALLEGRO_DIALOG_FUNC(int, al_append_menu_item, (ALLEGRO_MENU *parent, char const *title, int id, int flags,
   ALLEGRO_BITMAP *icon, ALLEGRO_MENU *submenu));
ALLEGRO_DIALOG_FUNC(int, al_insert_menu_item, (ALLEGRO_MENU *parent, int pos, char const *title, int id,
   int flags, ALLEGRO_BITMAP *icon, ALLEGRO_MENU *submenu));
ALLEGRO_DIALOG_FUNC(bool, al_remove_menu_item, (ALLEGRO_MENU *menu, int pos));
ALLEGRO_DIALOG_FUNC(ALLEGRO_MENU *, al_clone_menu, (ALLEGRO_MENU *menu));
ALLEGRO_DIALOG_FUNC(ALLEGRO_MENU *, al_clone_menu_for_popup, (ALLEGRO_MENU *menu));
ALLEGRO_DIALOG_FUNC(void, al_destroy_menu, (ALLEGRO_MENU *menu));


ALLEGRO_DIALOG_FUNC(const char *, al_get_menu_item_caption, (ALLEGRO_MENU *menu, int pos));
ALLEGRO_DIALOG_FUNC(void, al_set_menu_item_caption, (ALLEGRO_MENU *menu, int pos, const char *caption));
ALLEGRO_DIALOG_FUNC(int, al_get_menu_item_flags, (ALLEGRO_MENU *menu, int pos));
ALLEGRO_DIALOG_FUNC(void, al_set_menu_item_flags, (ALLEGRO_MENU *menu, int pos, int flags));
ALLEGRO_DIALOG_FUNC(int, al_toggle_menu_item_flags, (ALLEGRO_MENU *menu, int pos, int flags));
ALLEGRO_DIALOG_FUNC(ALLEGRO_BITMAP *, al_get_menu_item_icon, (ALLEGRO_MENU *menu, int pos));
ALLEGRO_DIALOG_FUNC(void, al_set_menu_item_icon, (ALLEGRO_MENU *menu, int pos, ALLEGRO_BITMAP *icon));

ALLEGRO_DIALOG_FUNC(ALLEGRO_MENU *, al_find_menu, (ALLEGRO_MENU *haystack, int id));
ALLEGRO_DIALOG_FUNC(bool, al_find_menu_item, (ALLEGRO_MENU *haystack, int id, ALLEGRO_MENU **menu, int *index));

ALLEGRO_DIALOG_FUNC(ALLEGRO_EVENT_SOURCE *, al_get_default_menu_event_source, (void));
ALLEGRO_DIALOG_FUNC(ALLEGRO_EVENT_SOURCE *, al_enable_menu_event_source, (ALLEGRO_MENU *menu));
ALLEGRO_DIALOG_FUNC(void, al_disable_menu_event_source, (ALLEGRO_MENU *menu));

ALLEGRO_DIALOG_FUNC(ALLEGRO_MENU *, al_get_display_menu, (ALLEGRO_DISPLAY *display));
ALLEGRO_DIALOG_FUNC(bool, al_set_display_menu, (ALLEGRO_DISPLAY *display, ALLEGRO_MENU *menu));
ALLEGRO_DIALOG_FUNC(bool, al_popup_menu, (ALLEGRO_MENU *popup, ALLEGRO_DISPLAY *display));
ALLEGRO_DIALOG_FUNC(ALLEGRO_MENU *, al_remove_display_menu, (ALLEGRO_DISPLAY *display));

ALLEGRO_DIALOG_FUNC(uint32_t, al_get_allegro_native_dialog_version, (void));

enum {
   ALLEGRO_FILECHOOSER_FILE_MUST_EXIST = 1,
   ALLEGRO_FILECHOOSER_SAVE            = 2,
   ALLEGRO_FILECHOOSER_FOLDER          = 4,
   ALLEGRO_FILECHOOSER_PICTURES        = 8,
   ALLEGRO_FILECHOOSER_SHOW_HIDDEN     = 16,
   ALLEGRO_FILECHOOSER_MULTIPLE        = 32
};

enum {
   ALLEGRO_MESSAGEBOX_WARN             = 1<<0,
   ALLEGRO_MESSAGEBOX_ERROR            = 1<<1,
   ALLEGRO_MESSAGEBOX_OK_CANCEL        = 1<<2,
   ALLEGRO_MESSAGEBOX_YES_NO           = 1<<3,
   ALLEGRO_MESSAGEBOX_QUESTION         = 1<<4
};

enum {
   ALLEGRO_TEXTLOG_NO_CLOSE            = 1<<0,
   ALLEGRO_TEXTLOG_MONOSPACE           = 1<<1
};

enum {
   ALLEGRO_EVENT_NATIVE_DIALOG_CLOSE   = 600,
   ALLEGRO_EVENT_MENU_CLICK            = 601
};

enum {
   ALLEGRO_MENU_ITEM_ENABLED            = 0,
   ALLEGRO_MENU_ITEM_CHECKBOX           = 1,
   ALLEGRO_MENU_ITEM_CHECKED            = 2,
   ALLEGRO_MENU_ITEM_DISABLED           = 4
};

*/
