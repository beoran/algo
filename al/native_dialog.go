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

void al_append_native_text_log_wrapper(ALLEGRO_TEXTLOG * tl, const char * text) {
    al_append_native_text_log(tl, "%s", text);
}

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

func (info * MenuInfo) toC() * C.ALLEGRO_MENU_INFO {
    return (*C.ALLEGRO_MENU_INFO)(info)
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
func CreateNativeFileDialog(path, title, patterns string, mode int) *FileChooser {    
    cpath    := cstr(path)      ; defer cstrFree(cpath)
    ctitle   := cstr(title)     ; defer cstrFree(ctitle)
    cpatterns:= cstr(patterns)  ; defer cstrFree(cpatterns)
    
    return wrapFileChooser(C.al_create_native_file_dialog(cpath, ctitle, cpatterns, C.int(mode)))
}

func (dialog * FileChooser) Show(display * Display) bool {
    return bool(C.al_show_native_file_dialog(display.toC(), dialog.toC()))
}

func (display * Display) ShowNativeFileDialog(dialog * FileChooser) bool {
    return bool(C.al_show_native_file_dialog(display.toC(), dialog.toC()))
}

func (dialog *  FileChooser) Count() int {
    return int(C.al_get_native_file_dialog_count(dialog.toC()))
}

func (dialog * FileChooser) Path(index int) string {
    return C.GoString(C.al_get_native_file_dialog_path(dialog.toC(), C.size_t(index)))
}


func (display * Display) ShowNativeMessageBox(title, heading, text, buttons string, flags int) int {
    ctitle   := cstr(title)     ; defer cstrFree(ctitle)
    cheading := cstr(heading)   ; defer cstrFree(cheading)
    ctext    := cstr(text)      ; defer cstrFree(ctext)
    cbuttons := cstr(buttons)   ; defer cstrFree(cbuttons)
    
    return int(C.al_show_native_message_box(display.toC(), ctitle, cheading, ctext, cbuttons, C.int(flags)))
}

// Creates a native text log.
func CreateNativeTextLog(title string, flags int) * TextLog {    
    ctitle   := cstr(title) ; defer cstrFree(ctitle)
    
    return wrapTextLog(C.al_open_native_text_log(ctitle, C.int(flags)))
}

func (log * TextLog) Append(text string) {
    ctext    := cstr(text) ; defer cstrFree(ctext)
    C.al_append_native_text_log_wrapper(log.toC(), ctext)
}

func (log * TextLog) EventSource() * EventSource {
    return wrapEventSourceRaw(C.al_get_native_text_log_event_source(log.toC()))
}

func CreateMenu() * Menu {
    return wrapMenu(C.al_create_menu())
}

func CreatePopupMenu() * Menu {
    return wrapMenu(C.al_create_popup_menu())
}

func BuildMenu(info * MenuInfo) * Menu {
    return wrapMenu(C.al_build_menu(info.toC()))
}

func (menu * Menu) AppendItem(title string, id, flags int, icon * Bitmap, submenu * Menu) int {
    ctitle   := cstr(title)     ; defer cstrFree(ctitle)
    return int(C.al_append_menu_item(menu.toC(), ctitle, C.uint16_t(id), C.int(flags), icon.toC(), submenu.toC()))
}

func (menu * Menu) InsertItem(pos int, title string, id, flags int, icon * Bitmap, submenu * Menu) int {
    ctitle   := cstr(title)     ; defer cstrFree(ctitle)
    return int(C.al_insert_menu_item(menu.toC(), C.int(pos), ctitle, C.uint16_t(id), C.int(flags), icon.toC(), submenu.toC()))
}

func (menu * Menu) RemoveItem(position int) bool {    
    return bool(C.al_remove_menu_item(menu.toC(), C.int(position)))
}

func (menu * Menu) Clone() * Menu {
    return wrapMenu(C.al_clone_menu(menu.toC()))
}

func (menu * Menu) CloneForPopup() * Menu {
    return wrapMenu(C.al_clone_menu_for_popup(menu.toC()))
}

func (menu * Menu) Caption(pos int) string {
    return C.GoString(C.al_get_menu_item_caption(menu.toC(), C.int(pos)))
}

func (menu * Menu) SetCaption(pos int, caption string) {
    ccaption  := cstr(caption)     ; defer cstrFree(ccaption)
    C.al_set_menu_item_caption(menu.toC(), C.int(pos), ccaption)
}

func (menu * Menu) Flags(pos int) int {
    return int(C.al_get_menu_item_flags(menu.toC(), C.int(pos)))
}

func (menu * Menu) SetFlags(pos int, flags int) {    
    C.al_set_menu_item_flags(menu.toC(), C.int(pos), C.int(flags))
}

func (menu * Menu) Icon(pos int) * Bitmap {
    return wrapBitmapRaw(C.al_get_menu_item_icon(menu.toC(), C.int(pos)))
}

func (menu * Menu) SetIcon(pos int, icon * Bitmap) {    
    C.al_set_menu_item_icon(menu.toC(), C.int(pos), icon.toC())
}


func (menu * Menu) Find(id int) * Menu {
    res     := C.al_find_menu(menu.toC(), C.uint16_t(id))
    return wrapMenuRaw(res)
} 


func (menu * Menu) FindItem(id int) (ok bool, found * Menu, index int) {
    var cmenu * C.ALLEGRO_MENU = nil
    var cindex C.int = -1
    res     := C.al_find_menu_item(menu.toC(), C.uint16_t(id), &cmenu, &cindex)
    ok      = bool(res)
    found   = wrapMenuRaw(cmenu)
    index   = int(cindex)
    return ok, menu, index
}


func DefaultMenuEventSource() * EventSource {
    return wrapEventSourceRaw(C.al_get_default_menu_event_source())
}

func (menu * Menu) EnableEventSource() * EventSource {
    return wrapEventSourceRaw(C.al_enable_menu_event_source(menu.toC()))
}


func (menu * Menu) DisableEventSource() {
    C.al_disable_menu_event_source(menu.toC())
}

func (disp * Display) Menu() * Menu {
    return wrapMenuRaw(C.al_get_display_menu(disp.toC()))
}

func (disp * Display) SetMenu(menu * Menu) bool {
    return bool(C.al_set_display_menu(disp.toC(), menu.toC()))
}

func (disp * Display) PopupMenu(menu * Menu) bool {
    return bool(C.al_popup_menu(menu.toC(), disp.toC()))
}

func (menu * Menu) Popup(disp * Display) bool {
    return bool(C.al_popup_menu(menu.toC(), disp.toC()))
}


func (disp * Display) RemoveMenu() * Menu {
    return wrapMenuRaw(C.al_remove_display_menu(disp.toC()))
}

func NativeDialogVersion() uint32 {
    return uint32(C.al_get_allegro_native_dialog_version())
}

