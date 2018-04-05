// threads and tls support
package al

/*
#include <stdlib.h>
#include <allegro5/allegro.h>
#include "helpers.h"
#include "callbacks.h"
*/
import "C"
import "runtime"
import "unsafe"


type Thread struct {
    handle * C.ALLEGRO_THREAD
}

// Converts a thread to it's underlying C pointer
func (self * Thread) toC() *C.ALLEGRO_THREAD {
    return (*C.ALLEGRO_THREAD)(self.handle)
}

// Destroys the thread.
func (self *Thread) Destroy() {
    if self.handle != nil {
        C.al_destroy_thread(self.toC())
    }
    self.handle = nil
}

// Wraps a C thread into a go thread
func wrapThreadRaw(data *C.ALLEGRO_THREAD) *Thread {
    if data == nil {
        return nil
    }
    return &Thread{data}
}

// Sets up a finalizer for this Thread that calls Destroy()
func (self *Thread) SetDestroyFinalizer() *Thread {
    if self != nil {
        runtime.SetFinalizer(self, func(me *Thread) { me.Destroy() })
    }
    return self
}

// Wraps a C thread into a go thread and sets up a finalizer that calls Destroy()
func wrapThread(data *C.ALLEGRO_THREAD) *Thread {
    self := wrapThreadRaw(data)
    return self.SetDestroyFinalizer()
}

type Cond struct {
    handle * C.ALLEGRO_COND
}

// Converts a cond to it's underlying C pointer
func (self * Cond) toC() *C.ALLEGRO_COND {
    return (*C.ALLEGRO_COND)(self.handle)
}

// Destroys the cond.
func (self *Cond) Destroy() {
    if self.handle != nil {
        C.al_destroy_cond(self.toC())
    }
    self.handle = nil
}

// Wraps a C cond into a go cond
func wrapCondRaw(data *C.ALLEGRO_COND) *Cond {
    if data == nil {
        return nil
    }
    return &Cond{data}
}

// Sets up a finalizer for this Cond that calls Destroy()
func (self *Cond) SetDestroyFinalizer() *Cond {
    if self != nil {
        runtime.SetFinalizer(self, func(me *Cond) { me.Destroy() })
    }
    return self
}

// Wraps a C cond into a go cond and sets up a finalizer that calls Destroy()
func wrapCond(data *C.ALLEGRO_COND) *Cond {
    self := wrapCondRaw(data)
    return self.SetDestroyFinalizer()
}

type Mutex struct {
    handle * C.ALLEGRO_MUTEX
}

// Converts a mutex to it's underlying C pointer
func (self * Mutex) toC() *C.ALLEGRO_MUTEX {
    return (*C.ALLEGRO_MUTEX)(self.handle)
}

// Destroys the mutex.
func (self *Mutex) Destroy() {
    if self.handle != nil {
        C.al_destroy_mutex(self.toC())
    }
    self.handle = nil
}

// Wraps a C mutex into a go mutex
func wrapMutexRaw(data *C.ALLEGRO_MUTEX) *Mutex {
    if data == nil {
        return nil
    }
    return &Mutex{data}
}

// Sets up a finalizer for this Mutex that calls Destroy()
func (self *Mutex) SetDestroyFinalizer() *Mutex {
    if self != nil {
        runtime.SetFinalizer(self, func(me *Mutex) { me.Destroy() })
    }
    return self
}

// Wraps a C mutex into a go mutex and sets up a finalizer that calls Destroy()
func wrapMutex(data *C.ALLEGRO_MUTEX) *Mutex {
    self := wrapMutexRaw(data)
    return self.SetDestroyFinalizer()
}


func CreateThread(fn ThreadCallbackFunction, data unsafe.Pointer) *Thread {
    cbd := & threadCallbackData{fn , data}
    ct := C.al_create_thread((*[0]byte)(C.go_create_thread_callback), unsafe.Pointer(cbd))
    return wrapThread(ct)
}

func (thread * Thread) Start() {
    C.al_start_thread(thread.toC())
}

func (thread * Thread) Join() (data interface {}) {
    gdata:= make([]byte, 64)
    gptr := unsafe.Pointer(&gdata)
    cptr := &gptr 
    /* XXX: I am not sure this hack will work. */
    C.al_join_thread(thread.toC(), cptr)
    return gdata
}

func (thread * Thread) ShouldStop() (bool) {
    return bool(C.al_get_thread_should_stop(thread.toC())) 
}

func (thread * Thread) SetShouldStop() {
    C.al_set_thread_should_stop(thread.toC()) 
}

func RunDetachedThread(fn ThreadCallbackFunction, data unsafe.Pointer) {
    cbd := & threadCallbackData{fn , data}
    C.al_run_detached_thread((*[0]byte)(C.go_create_thread_callback), unsafe.Pointer(cbd))
}

func CreateMutex() * Mutex {
    return wrapMutex(C.al_create_mutex())
}

func CreateMutexRecursive() * Mutex {
    return wrapMutex(C.al_create_mutex_recursive())
}

func (mutex * Mutex) Lock() {
    C.al_lock_mutex(mutex.toC())
} 

func (mutex * Mutex) Unlock() {
    C.al_unlock_mutex(mutex.toC())
} 

func CreateCond() * Cond {
    return wrapCond(C.al_create_cond())
}

func (cond * Cond) Wait(mutex * Mutex) {
    C.al_wait_cond(cond.toC(), mutex.toC())
} 

func (cond * Cond) WaitUntil(mutex * Mutex, timeout * Timeout) {
    C.al_wait_cond_until(cond.toC(), mutex.toC(), timeout.toC())
} 

func (cond * Cond) Broadcast(mutex * Mutex) {
    C.al_broadcast_cond(cond.toC())
} 

func (cond * Cond) Signal(mutex * Mutex) {
    C.al_signal_cond(cond.toC())
} 



const (
    STATE_NEW_DISPLAY_PARAMETERS= C.ALLEGRO_STATE_NEW_DISPLAY_PARAMETERS
    STATE_NEW_BITMAP_PARAMETERS = C.ALLEGRO_STATE_NEW_BITMAP_PARAMETERS 
    STATE_DISPLAY               = C.ALLEGRO_STATE_DISPLAY               
    STATE_TARGET_BITMAP         = C.ALLEGRO_STATE_TARGET_BITMAP         
    STATE_BLENDER               = C.ALLEGRO_STATE_BLENDER               
    STATE_NEW_FILE_INTERFACE    = C.ALLEGRO_STATE_NEW_FILE_INTERFACE    
    STATE_TRANSFORM             = C.ALLEGRO_STATE_TRANSFORM             
    STATE_PROJECTION_TRANSFORM  = C.ALLEGRO_STATE_PROJECTION_TRANSFORM  
    STATE_BITMAP                = C.ALLEGRO_STATE_BITMAP                
    STATE_ALL                   = C.ALLEGRO_STATE_ALL                   
)



type State C.ALLEGRO_STATE;

func StoreState(flags int) * State {
    state := &C.ALLEGRO_STATE{}
    C.al_store_state(state, C.int(flags))
    return (*State)(state)
} 

func (state * State) Restore() {
    cstate := (*C.ALLEGRO_STATE)(state)
    C.al_restore_state(cstate);
}


