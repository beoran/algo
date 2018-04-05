// blender support
package al

/*
#include <stdlib.h>
#include <allegro5/allegro.h>
#include "helpers.h"
*/
import "C"

// import "runtime"
// import "unsafe"

type BlendMode C.enum_ALLEGRO_BLEND_MODE

/*
 * Blending modes
 */
const (
   ZERO               = BlendMode(C.ALLEGRO_ZERO)
   ONE                = BlendMode(C.ALLEGRO_ONE)
   ALPHA              = BlendMode(C.ALLEGRO_ALPHA)
   INVERSE_ALPHA      = BlendMode(C.ALLEGRO_INVERSE_ALPHA)
   SRC_COLOR          = BlendMode(C.ALLEGRO_SRC_COLOR)
   DEST_COLOR         = BlendMode(C.ALLEGRO_DEST_COLOR)  
   INVERSE_SRC_COLOR  = BlendMode(C.ALLEGRO_INVERSE_SRC_COLOR)
   INVERSE_DEST_COLOR = BlendMode(C.ALLEGRO_INVERSE_DEST_COLOR)
   CONST_COLOR        = BlendMode(C.ALLEGRO_CONST_COLOR)
   INVERSE_CONST_COLOR= BlendMode(C.ALLEGRO_INVERSE_CONST_COLOR)
   NUM_BLEND_MODES    = BlendMode(C.ALLEGRO_NUM_BLEND_MODES)
)

func (bm BlendMode) toC() C.enum_ALLEGRO_BLEND_MODE {
    return C.enum_ALLEGRO_BLEND_MODE(bm)
}

type BlendOperations C.enum_ALLEGRO_BLEND_OPERATIONS

func (bo BlendOperations) toC() C.enum_ALLEGRO_BLEND_MODE {
    return C.enum_ALLEGRO_BLEND_OPERATIONS(bo)
}

const(
   ADD                  = BlendOperations(C.ALLEGRO_ADD)     
   SRC_MINUS_DEST       = BlendOperations(C.ALLEGRO_SRC_MINUS_DEST)     
   DEST_MINUS_SRC       = BlendOperations(C.ALLEGRO_DEST_MINUS_SRC)
   NUM_BLEND_OPERATIONS = BlendOperations(C.ALLEGRO_NUM_BLEND_OPERATIONS)
)


func SetBlender(op BlendOperations, src, dest BlendMode) {
    C.al_set_blender(C.int(op), C.int(src), C.int(dest))
}

func SetBlendColor(color Color) {
    C.al_set_blend_color(color.toC())
}

func GetBlender() (op BlendOperations, src, dest BlendMode) {
    var cop, csrc, cdest C.int
    C.al_get_blender(&cop, &csrc, &cdest)
    op   = BlendOperations(cop)
    src  = BlendMode(csrc)
    dest = BlendMode(cdest)
    return op, src, dest
}

func GetBlendColor() Color {
    return wrapColor(C.al_get_blend_color())
}


func SetSeparateBlender(op BlendOperations, src, dest BlendMode,
    alpha_op BlendOperations, alpha_src, alpha_dest BlendMode) {
    C.al_set_separate_blender(C.int(op), C.int(src), C.int(dest),
        C.int(alpha_op), C.int(alpha_src), C.int(alpha_dest))
}

func GetSeparateBlender() (op BlendOperations, src, dest BlendMode, 
    alpha_op BlendOperations, alpha_src, alpha_dest BlendMode) {
    var cop, csrc, cdest, calpha_op, calpha_src, calpha_dest C.int
    C.al_get_separate_blender(&cop, &csrc, &cdest,
        &calpha_op, &calpha_src, &calpha_dest)
    op          = BlendOperations(cop)
    src         = BlendMode(csrc)
    dest        = BlendMode(cdest)
    alpha_op    = BlendOperations(calpha_op)
    alpha_src   = BlendMode(calpha_src)
    alpha_dest  = BlendMode(calpha_dest)
    
    return op, src, dest, alpha_op, alpha_src, alpha_dest
}
