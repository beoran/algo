#include <stdio.h>
#include <stdlib.h>
#include <allegro5/allegro.h>
#include "helpers.h"
#include "callbacks.h"


int take_callback(function_pointer cb, void * data) {
    go_take_callback_function * fb = (go_take_callback_function *) cb;
    int res;
    puts("Calling Callback");
    res = fb(data);
    puts("Callback Called");
    return res;
}



