package al

/*
#include <stdlib.h>
#include <allegro5/allegro.h>
#include "helpers.h"
#include "callbacks.h"
*/
import "C"


type Transform C.ALLEGRO_TRANSFORM

func wrapTransformRaw(ctrans * C.ALLEGRO_TRANSFORM) * Transform {
    return (*Transform)(ctrans)
}

func (trans * Transform) toC() * C.ALLEGRO_TRANSFORM {
    return (* C.ALLEGRO_TRANSFORM)(trans)
}


func (trans * Transform) Init(matrix [4][4]float32) {
    for i:=0; i <4; i ++ {
        for j:=0; j < 4; j ++ {
            trans.Put(i, j, matrix[i][j])
        }
    }
}

func (trans * Transform) Matrix() (matrix [4][4]float32) {    
    for i:=0; i <4; i ++ {
        for j:=0; j < 4; j ++ {
            matrix[i][j] = trans.Get(i,j)
        }
    }
    return matrix
}

func (trans * Transform) Get(i, j int) float32 {
    return float32(trans.m[i][j])
}

func (trans * Transform) Put(i, j int, v float32) {
    trans.m[i][j] = C.float(v)
}

func CreateIdentityTransform() * Transform {
    trans := &Transform{}
    return trans.Identity()
}


func CreateTransform(x, y, sx, sy, theta float32) * Transform {
    trans := &Transform{}
    return trans.Build(x, y, sx, sy, theta)
}

func (trans * Transform) Identity() * Transform {
    C.al_identity_transform(trans.toC())
    return trans
}


func (trans * Transform) Build(x, y, sx, sy, theta float32) * Transform {
    C.al_build_transform(trans.toC(), C.float(x), C.float(y), C.float(sx), C.float(sy), C.float(theta))
    return trans
}

func (trans * Transform) Use() {
    C.al_use_transform(trans.toC())
}

func (trans * Transform) UseProjection() {
    C.al_use_projection_transform(trans.toC())
}

func (trans * Transform) BuildCamera(px, py, pz, 
    lx, ly, lz, ux, uy, uz float32) * Transform {
    cpx, cpy, cpz := cf3(px, py, pz)
    clx, cly, clz := cf3(lx, ly, lz)
    cux, cuy, cuz := cf3(ux, uy, uz)
    C.al_build_camera_transform(trans.toC(), cpx, cpy, cpz, clx, cly, clz, cux, cuy, cuz)
    return trans
}


func (trans * Transform) Translate(x, y float32) * Transform {
    C.al_translate_transform(trans.toC(), C.float(x), C.float(y))
    return trans
}

func (trans * Transform) TranslateInt(x, y int) * Transform {
    C.al_translate_transform(trans.toC(), C.float(x), C.float(y))
    return trans
}


func (trans * Transform) Translate3D(x, y, z float32) * Transform {
    C.al_translate_transform_3d(trans.toC(), C.float(x), C.float(y), C.float(z))
    return trans
}


func (trans * Transform) Rotate(theta float32) * Transform {
    C.al_rotate_transform(trans.toC(), C.float(theta))
    return trans
}

func (trans * Transform) Rotate3D(x, y, z, angle float32) * Transform {
    C.al_rotate_transform_3d(trans.toC(), C.float(x), C.float(y), C.float(z), C.float(angle))
    return trans
}


func (trans * Transform) Scale(x, y float32) * Transform {
    C.al_scale_transform(trans.toC(), C.float(x), C.float(y))
    return trans
}

func (trans * Transform) ScaleInt(x, y int) * Transform {
    C.al_scale_transform(trans.toC(), C.float(x), C.float(y))
    return trans
}


func (trans * Transform) Scale3D(x, y, z float32) * Transform {
    C.al_scale_transform_3d(trans.toC(), C.float(x), C.float(y), C.float(z))
    return trans
}

func (trans * Transform) Coordinates(x, y float32) (float32, float32) {
    cx, cy := cf2(x, y)
    C.al_transform_coordinates(trans.toC(), &cx, &cy)    
    return float32(cx), float32(cy)
}

func (trans * Transform) Coordinates3D(x, y, z float32) (float32, float32, float32) {
    cx, cy, cz := cf3(x, y, z)
    C.al_transform_coordinates_3d(trans.toC(), &cx, &cy, &cz)    
    return float32(cx), float32(cy), float32(cz)
}


func (trans * Transform) Compose(other * Transform)  * Transform {
    C.al_compose_transform(trans.toC(), other.toC())    
    return trans
}


func (trans * Transform) CheckInverse(tolerance float32) bool {
    return 1 == (C.al_check_inverse(trans.toC(), C.float(tolerance))) 
}


func (trans * Transform) Invert()  * Transform {
    if trans.CheckInverse(1.0e-7) { 
        C.al_invert_transform(trans.toC())
        return trans
    } else {
        return nil
    }
}
func CurrentTransform() * Transform {
    return wrapTransformRaw(C.al_get_current_transform())
}

func CurrentProjectionTransform() * Transform {
    return wrapTransformRaw(C.al_get_current_projection_transform())
}

func (trans * Transform) Orthographic(l, t, n, r, b, f float32) * Transform {
    cl, ct, cn := cf3(l,t,n)
    cr, cb, cf := cf3(r,b,f)
    C.al_orthographic_transform(trans.toC(), cl, ct, cn, cr, cb, cf)
    return trans
}

func (trans * Transform) Perspective(l, t, n, r, b, f float32) * Transform {
    cl, ct, cn := cf3(l,t,n)
    cr, cb, cf := cf3(r,b,f)
    C.al_perspective_transform(trans.toC(), cl, ct, cn, cr, cb, cf)
    return trans
}

func (trans * Transform) HorizontalShear(theta float32) * Transform {
    C.al_horizontal_shear_transform(trans.toC(), C.float(theta))
    return trans
}

func (trans * Transform) VerticalShear(theta float32) * Transform {
    C.al_horizontal_shear_transform(trans.toC(), C.float(theta))
    return trans
}

