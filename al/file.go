// acodec
package al

/*
#cgo pkg-config: allegro-5
#cgo CFLAGS: -I/usr/local/include
#cgo linux LDFLAGS: -lc_nonshared
#include <stdlib.h>
#include <allegro5/allegro.h>
#include <allegro5/file.h>
#include "helpers.h"
#include "callbacks.h"
*/
import "C"

import "errors"
import "runtime"
import "unsafe"

// Allegro's own file for cross platform and physfs reasons.
type File struct {
    handle *C.ALLEGRO_FILE
}

// Closes the Allegro file
func (self *File) Close() {
    if self.handle != nil {
        C.al_fclose(self.handle)
    }
    self.handle = nil
}

// Returns the low level handle of the file 
func (self *File) toC() * C.ALLEGRO_FILE {
    return self.handle
}

// Wraps an ALLEGRO_FILE into a File
func wrapFileRaw(file *C.ALLEGRO_FILE) *File {
    if file == nil {
        return nil
    }
    return &File{file}
}

// Opens an Allegro File
func openFile(filename, mode string) *C.ALLEGRO_FILE {
    cfilename := cstr(filename)
    defer cstrFree(cfilename)
    cmode := cstr(mode)
    defer cstrFree(cmode)
    return C.al_fopen(cfilename, cmode)
}

// Sets up a finalizer for this File that calls Close()
func (self *File) SetCloseFinalizer() *File {
    if self != nil {
        runtime.SetFinalizer(self, func(me *File) { me.Close() })
    }
    return self
}

// Wraps a file and sets up a finalizer that calls Destroy()
func wrapFile(data *C.ALLEGRO_FILE) *File {
    self := wrapFileRaw(data)
    return self.SetCloseFinalizer()
}

// Opens a file with no finalizer set
func OpenFileRaw(filename, mode string) *File {
    self := openFile(filename, mode)
    return wrapFileRaw(self)
}

// Opens a file with a Close finalizer set
func OpenFile(filename, mode string) *File {
    self := OpenFileRaw(filename, mode)
    return self.SetCloseFinalizer()
}

/* File interface is as good as useless in GO, too many
 * callbacks. */

type Seek C.enum_ALLEGRO_SEEK

const (
   SEEK_SET = Seek(C.ALLEGRO_SEEK_SET)
   SEEK_CUR = Seek(C.ALLEGRO_SEEK_CUR)
   SEEK_END = Seek(C.ALLEGRO_SEEK_END)
)

func (f * File) Read(p []byte) (n int, err error) {
    err = nil
    size := C.al_fread(f.toC(), unsafe.Pointer(&p[0]), C.size_t(len(p)))
    n = int(size)
    erri := C.al_ferror(f.toC()) 
    if erri != 0 {
        mesg := C.GoString(C.al_ferrmsg(f.toC()))
        err = errors.New("Could not read file: " + mesg)
    } 
    return n, err
}

/* No Allegro File IO (yet) since that is redundant with Go's std libs. 
 * Likeswise I'm skipping fixed point math and 
 * the FS entry.
 */



