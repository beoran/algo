package al

/*
#include <stdlib.h>
#include <stdint.h>
*/
import "C"
import "unsafe"
import "runtime"
import "fmt"

// Helper functions for working with C easier

// Calls C malloc
func malloc(size int) unsafe.Pointer {
    return (unsafe.Pointer(C.calloc(C.size_t(size), C.size_t(1))))
}

// Calls C free
func free(ptr unsafe.Pointer) {
    C.free(ptr)
}

// Allocates a string with the given byte length
// don't forget a call to defer cstrFree() ! 
func cstrNew(size int) *C.char {
    return (*C.char)(malloc((size)))
}

// free is a method on C char * strings to method to free the associated memory 
func cstrFree(self *C.char) {
    free(unsafe.Pointer(self))
}

// Coverts a string to a C string. This allocates memory, 
// so don't forget to add a "defer cstrFree(cstr)"
func cstr(self string) *C.char {
    return C.CString(self)
}

// Shorthand for C.GoString. Yes, it's just laziness. :)
func gostr(cstr *C.char) string {
    return C.GoString(cstr)
}

// Converts an int pointer to a C.int pointer.
func cintptr(ptr *int) *C.int {
    return (*C.int)(unsafe.Pointer(ptr))
}

/*
// Converts a byte pointer to a C.Uchar8 pointer.
func cbyteptr(ptr * uint8)  (*C.Uint8)  { 
  return (*C.Uint8)(unsafe.Pointer(ptr))
}
*/

// Converts ints to bools.
func i2b(res int) bool {
    if res != 0 {
        return true
    }
    return false
}

// Converts bools to ints.
func b2i(res bool) int {
    if res {
        return 1
    }
    return 0
}

// Interface for destructable objects
type Destroyer interface {
    Destroy()
}

// Sets up a automatic finalizer for destructable objects
// that will call Destroy using runtime.SetFinalizer
// when the garbage collecter cleans up self.
// self may also be nil in which case the destructor is NOT set up
func SelfDestruct(self Destroyer) {
    if self == nil {
        return
    }
    clean := func(me Destroyer) {
        fmt.Printf("Finalizing %#v.\n", me)
        me.Destroy()
    }
    runtime.SetFinalizer(self, clean)
}

// this is too for laziness, but it's quite handy
func cf(f float32) C.float {
    return C.float(f)
}

// this is too for laziness, but it's quite handy
func ci(f int) C.int {
    return C.int(f)
}

// this is too for laziness, but it's quite handy
func cd(f float64) C.double {
    return C.double(f)
}


// this is too for laziness, but it's quite handy
func cui16(f int) C.uint16_t {
    return C.uint16_t(f)
}

//Converts an array of C strings to a slice of Go strings
func GoStrings(argc C.int, argv **C.char) []string {
    length := int(argc)
    tmpslice := (*[1 << 30]*C.char)(unsafe.Pointer(argv))[:length:length]
    gostrings := make([]string, length)
    for i, s := range tmpslice {
        gostrings[i] = C.GoString(s)
    }
    return gostrings
}

//Converts an array of go strings to an array of C strings and a length 
func CStrings(args []string) (argc C.int, argv **C.char) {
    length := len(args)
    argv = (**C.char)(malloc(length * int(unsafe.Sizeof(*argv))))
    tmpslice := (*[1 << 30]*C.char)(unsafe.Pointer(argv))[:length:length]
    for i, s := range args {
        tmpslice[i] = cstr(s)
    }
    argc = C.int(length)
    return argc, argv
}

// frees the data allocated by Cstrings
func CStringsFree(argc C.int, argv **C.char) {
    length := int(argc)
    tmpslice := (*[1 << 30]*C.char)(unsafe.Pointer(argv))[:length:length]
    for _, s := range tmpslice {
        cstrFree(s)
    }
    free(unsafe.Pointer(argv))
}




/* This is the usual boilerplate for wrapping C types through a handle

type XXX struct {
    handle * C.YYY
}

// Converts a zzz to it's underlying C pointer
func (self * XXX) toC() *C.YYY {
    return (*C.YYY)(self.handle)
}

// Destroys the zzz.
func (self *XXX) Destroy() {
    if self.handle != nil {
        C.al_destroy_zzz(self.toC())
    }
    self.handle = nil
}

// Wraps a C zzz into a go zzz
func wrapXXXRaw(data *C.YYY) *XXX {
    if data == nil {
        return nil
    }
    return &XXX{data}
}

// Sets up a finalizer for this XXX that calls Destroy()
func (self *XXX) SetDestroyFinalizer() *XXX {
    if self != nil {
        runtime.SetFinalizer(self, func(me *XXX) { me.Destroy() })
    }
    return self
}

// Wraps a C zzz into a go zzz and sets up a finalizer that calls Destroy()
func wrapXXX(data *C.YYY) *XXX {
    self := wrapXXXRaw(data)
    return self.SetDestroyFinalizer()
}


*/
