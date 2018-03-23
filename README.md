algo
====

Go language binding to Allegro 5.2.x (under construction).

It presents a more Go-like interface with OOP on most resources after they are created or 
loaded. 

Most resources can be created or loaded in raw form, in which case Destroy() must be called 
on them, or with a finalizer set that calls Destroy automatically. It shoul be safe to call 
Destroy multiple times, and finalizers are not 100% reliable, so it is recommended to call
the Destoy() method manually on any resource that isn't needed anymore.

This wrapper is licensed for free use under the MIT license. 
