package mainthread

import "runtime"

func Run(Main func()) {
    runtime.LockOSThread()
    Main()
    runtime.UnlockOSThread()
}






