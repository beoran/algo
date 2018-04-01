// TODO configuration support
package al

/*
#include <stdlib.h>
#include <allegro5/allegro.h>
#include "helpers.h"
#include "callbacks.h"
*/
import "C"

func CPUCount() int {
    return int(C.al_get_cpu_count())
}

func RamSize() int { 
    return int(C.al_get_ram_size())
}

