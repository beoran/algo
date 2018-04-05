package al

/*
#include <stdlib.h>
#include <allegro5/allegro.h>
#include "helpers.h"
*/
import "C"

type RenderState = C.ALLEGRO_RENDER_STATE

const (
   ALPHA_TEST       = RenderState(C.ALLEGRO_ALPHA_TEST      )
   WRITE_MASK       = RenderState(C.ALLEGRO_WRITE_MASK      )
   DEPTH_TEST       = RenderState(C.ALLEGRO_DEPTH_TEST      )
   DEPTH_FUNCTION   = RenderState(C.ALLEGRO_DEPTH_FUNCTION  )
   ALPHA_FUNCTION   = RenderState(C.ALLEGRO_ALPHA_FUNCTION  )
   ALPHA_TEST_VALUE = RenderState(C.ALLEGRO_ALPHA_TEST_VALUE)
)

func (rs RenderState) toC() C.ALLEGRO_RENDER_STATE {
    return C.ALLEGRO_RENDER_STATE(rs)
}

// not usable as such, cast to int in stead

// type RenderFunction = C.ALLEGRO_RENDER_FUNCTION

const (
   RENDER_NEVER         = int(C.ALLEGRO_RENDER_NEVER        )  
   RENDER_ALWAYS        = int(C.ALLEGRO_RENDER_ALWAYS       ) 
   RENDER_LESS          = int(C.ALLEGRO_RENDER_LESS         )
   RENDER_EQUAL         = int(C.ALLEGRO_RENDER_EQUAL        )
   RENDER_LESS_EQUAL    = int(C.ALLEGRO_RENDER_LESS_EQUAL   )    
   RENDER_GREATER       = int(C.ALLEGRO_RENDER_GREATER      ) 
   RENDER_NOT_EQUAL     = int(C.ALLEGRO_RENDER_NOT_EQUAL    )
   RENDER_GREATER_EQUAL = int(C.ALLEGRO_RENDER_GREATER_EQUAL)
)


// not usable as such, cast to int in stead

// type WriteMaskFlags C.ALLEGRO_WRITE_MASK_FLAGS

const (
   MASK_RED   = int(C.ALLEGRO_MASK_RED  ) 
   MASK_GREEN = int(C.ALLEGRO_MASK_GREEN) 
   MASK_BLUE  = int(C.ALLEGRO_MASK_BLUE ) 
   MASK_ALPHA = int(C.ALLEGRO_MASK_ALPHA) 
   MASK_DEPTH = int(C.ALLEGRO_MASK_DEPTH) 
   MASK_RGB   = int(C.ALLEGRO_MASK_RGB  ) 
   MASK_RGBA  = int(C.ALLEGRO_MASK_RGBA ) 
)

func SetRenderState(state RenderState, value int) { 
    C.al_set_render_state(state.toC(), C.int(value))
}

