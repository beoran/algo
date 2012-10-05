package al

/*
#include <stdlib.h>
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
	return (*C.char)(malloc(size))
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
