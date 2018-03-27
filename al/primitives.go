// acodec
package al

/*
#cgo pkg-config: allegro_primitives-5
#cgo CFLAGS: -I/usr/local/include
#cgo linux LDFLAGS: -lc_nonshared
#include <stdlib.h>
#include <allegro5/allegro.h>
#include <allegro5/allegro_primitives.h>
#include "helpers.h"
*/
import "C"
import "runtime"
import "unsafe"

type PrimType C.ALLEGRO_PRIM_TYPE


const (
  PRIM_LINE_LIST        = PrimType(C.ALLEGRO_PRIM_LINE_LIST)
  PRIM_LINE_STRIP       = PrimType(C.ALLEGRO_PRIM_LINE_STRIP)
  PRIM_LINE_LOOP        = PrimType(C.ALLEGRO_PRIM_LINE_LOOP)
  PRIM_TRIANGLE_LIST    = PrimType(C.ALLEGRO_PRIM_TRIANGLE_LIST)
  PRIM_TRIANGLE_STRIP   = PrimType(C.ALLEGRO_PRIM_TRIANGLE_STRIP)
  PRIM_TRIANGLE_FAN     = PrimType(C.ALLEGRO_PRIM_TRIANGLE_FAN)
  PRIM_POINT_LIST       = PrimType(C.ALLEGRO_PRIM_POINT_LIST)
  PRIM_NUM_TYPES        = PrimType(C.ALLEGRO_PRIM_NUM_TYPES)
)

const PRIM_MAX_USER_ATTR = C.ALLEGRO_PRIM_MAX_USER_ATTR

type PrimAttr C.ALLEGRO_PRIM_ATTR

const (
   PRIM_POSITION        = PrimAttr(C.ALLEGRO_PRIM_POSITION)
   PRIM_COLOR_ATTR      = PrimAttr(C.ALLEGRO_PRIM_COLOR_ATTR)
   PRIM_TEX_COORD       = PrimAttr(C.ALLEGRO_PRIM_TEX_COORD)
   PRIM_TEX_COORD_PIXEL = PrimAttr(C.ALLEGRO_PRIM_TEX_COORD_PIXEL)
   PRIM_USER_ATTR       = PrimAttr(C.ALLEGRO_PRIM_USER_ATTR)
   PRIM_ATTR_NUM        = PrimAttr(C.ALLEGRO_PRIM_ATTR_NUM)
)


type PrimStorage C.ALLEGRO_PRIM_STORAGE

const (
   PRIM_FLOAT_2             = PrimStorage(C.ALLEGRO_PRIM_FLOAT_2)
   PRIM_FLOAT_3             = PrimStorage(C.ALLEGRO_PRIM_FLOAT_3)
   PRIM_SHORT_2             = PrimStorage(C.ALLEGRO_PRIM_SHORT_2)
   PRIM_FLOAT_1             = PrimStorage(C.ALLEGRO_PRIM_FLOAT_1)
   PRIM_FLOAT_4             = PrimStorage(C.ALLEGRO_PRIM_FLOAT_4)
   PRIM_UBYTE_4             = PrimStorage(C.ALLEGRO_PRIM_UBYTE_4)
   PRIM_SHORT_4             = PrimStorage(C.ALLEGRO_PRIM_SHORT_4)
   PRIM_NORMALIZED_UBYTE_4  = PrimStorage(C.ALLEGRO_PRIM_NORMALIZED_UBYTE_4)
   PRIM_NORMALIZED_SHORT_2  = PrimStorage(C.ALLEGRO_PRIM_NORMALIZED_SHORT_2)
   PRIM_NORMALIZED_SHORT_4  = PrimStorage(C.ALLEGRO_PRIM_NORMALIZED_SHORT_4)
   PRIM_NORMALIZED_USHORT_2 = PrimStorage(C.ALLEGRO_PRIM_NORMALIZED_USHORT_2)
   PRIM_NORMALIZED_USHORT_4 = PrimStorage(C.ALLEGRO_PRIM_NORMALIZED_USHORT_4)
   PRIM_HALF_FLOAT_2        = PrimStorage(C.ALLEGRO_PRIM_HALF_FLOAT_2)
   PRIM_HALF_FLOAT_4        = PrimStorage(C.ALLEGRO_PRIM_HALF_FLOAT_4)
)

type LineJoin C.ALLEGRO_LINE_JOIN
const(
   LINE_JOIN_NONE   = LineJoin(C.ALLEGRO_LINE_JOIN_NONE)
   LINE_JOIN_BEVEL  = LineJoin(C.ALLEGRO_LINE_JOIN_BEVEL)
   LINE_JOIN_ROUND  = LineJoin(C.ALLEGRO_LINE_JOIN_ROUND)
   LINE_JOIN_MITER  = LineJoin(C.ALLEGRO_LINE_JOIN_MITER)
   LINE_JOIN_MITRE  = LineJoin(C.ALLEGRO_LINE_JOIN_MITRE)
)

type LineCap C.ALLEGRO_LINE_CAP

const(
   LINE_CAP_NONE        = LineCap(C.ALLEGRO_LINE_CAP_NONE)
   LINE_CAP_SQUARE      = LineCap(C.ALLEGRO_LINE_CAP_SQUARE)
   LINE_CAP_ROUND       = LineCap(C.ALLEGRO_LINE_CAP_ROUND)
   LINE_CAP_TRIANGLE    = LineCap(C.ALLEGRO_LINE_CAP_TRIANGLE)
   LINE_CAP_CLOSED      = LineCap(C.ALLEGRO_LINE_CAP_CLOSED)
)


type PrimBufferFlags C.ALLEGRO_PRIM_BUFFER_FLAGS
const (
   PRIM_BUFFER_STREAM       = PrimBufferFlags(C.ALLEGRO_PRIM_BUFFER_STREAM)
   PRIM_BUFFER_STATIC       = PrimBufferFlags(C.ALLEGRO_PRIM_BUFFER_STATIC)
   PRIM_BUFFER_DYNAMIC      = PrimBufferFlags(C.ALLEGRO_PRIM_BUFFER_DYNAMIC)
   PRIM_BUFFER_READWRITE    = PrimBufferFlags(C.ALLEGRO_PRIM_BUFFER_READWRITE)
)

const VERTEX_CACHE_SIZE = C.ALLEGRO_VERTEX_CACHE_SIZE

const PRIM_QUALITY = C.ALLEGRO_PRIM_QUALITY


type VertexElement = C.ALLEGRO_VERTEX_ELEMENT

// Convert from C
func wrapVertexElement(elem C.ALLEGRO_VERTEX_ELEMENT) VertexElement {
    return VertexElement(elem)
}

// Convert to C
func (elem VertexElement) toC() C.ALLEGRO_VERTEX_ELEMENT {
    return C.ALLEGRO_VERTEX_ELEMENT(elem)
}

// Convert to C
func (elem * VertexElement) toCPointer() * C.ALLEGRO_VERTEX_ELEMENT {
    return (* C.ALLEGRO_VERTEX_ELEMENT)(elem)
}

func (elem * VertexElement) Attribute() PrimAttr {
    return PrimAttr(elem.toC().attribute)
}

func (elem * VertexElement) Storage() PrimStorage {
    return PrimStorage(elem.toC().storage)
}

func (elem * VertexElement) Offset() int {
    return int(elem.toC().offset)
}

func (elem * VertexElement) SetAttribute(attribute PrimAttr) {
    // celem := elem.toC()
    elem.attribute = C.int(attribute)
}

func (elem * VertexElement) SetStorage(storage PrimStorage) {
     elem.storage = C.int(storage)
}

func (elem * VertexElement) SetOffset(offset int) {
    elem.offset = C.int(offset)
}

func NewVertexElement(attr PrimAttr, store PrimStorage, offset int) *VertexElement {
    res := new(VertexElement)
    res.SetAttribute(attr)
    res.SetStorage(store)
    res.SetOffset(offset)
    return res
}


type VertexDecl struct { 
    handle * C.ALLEGRO_VERTEX_DECL
}

// returns low level handle for the vertex declaration
func (vd *VertexDecl) toC() * C.ALLEGRO_VERTEX_DECL {
    return vd.handle
}

// Wraps a C Allegro vertex declaration into a VertexDecl Sets no finalizer.
func wrapVertexDeclRaw(handle *C.ALLEGRO_VERTEX_DECL) *VertexDecl {
    if handle == nil {
        return nil
    }
    return &VertexDecl{handle}
}

// Wraps a C Allegro vertex declaration. Sets a finalizer that calls Destroy.
func wrapVertexDecl(handle *C.ALLEGRO_VERTEX_DECL) *VertexDecl {
    vd := wrapVertexDeclRaw(handle)
    if vd != nil {
        runtime.SetFinalizer(vd, func(me * VertexDecl) { me.Destroy() })
    }
    return vd
}


func CreateVertexDecl(gelements []VertexElement, stride int) * VertexDecl {
    raw := C.al_create_vertex_decl(&gelements[0], C.int(stride))
    return wrapVertexDecl(raw)
}

func (vd * VertexDecl) Destroy() {
    if vd.handle != nil {
        C.al_destroy_vertex_decl(vd.handle)
    }
    vd.handle  = nil
}



type Vertex = C.ALLEGRO_VERTEX

type VertexBuffer struct {
    handle * C.ALLEGRO_VERTEX_BUFFER
}

type IndexBuffer struct { 
    handle * C.ALLEGRO_INDEX_BUFFER
}

func GetPrimitivesVersion() uint32 {
    return uint32(C.al_get_allegro_primitives_version())
}

func InitPrimitivesAddon() bool {
    return bool(C.al_init_primitives_addon())
}

func ShutdownPrimitivesAddon() {
    C.al_shutdown_primitives_addon()
}

/*
The generic case is risky due to the GC kicking in, and hard to square with the go type system .For now only support drawing ALLEGRO_VERTICES.
func DrawPrimGeneric(vertices []interface{}, decl * VertexDecl, texture * Bitmap, 
    start, end int, ptype PrimType) int {
    va := make([]unsafe.Pointer, len(vertices))
    for i:= 0 ; i < len(vertices); index ++
}
*/

func DrawPrim(v []Vertex, texture * Bitmap, start, end int, ptype PrimType) int {
    return int(C.al_draw_prim(unsafe.Pointer(&v[0]), nil,
                texture.toC(), C.int(start), C.int(end), C.int(ptype))) 
}

func DrawIndexedPrim(v []Vertex, indices []int, texture * Bitmap, start, end int, ptype PrimType) int {
    return int(C.al_draw_indexed_prim(unsafe.Pointer(&v[0]), nil,
                texture.toC(), (* C.int)(unsafe.Pointer(&indices[0])), 
                C.int(len(indices)), C.int(ptype))) 
}

func DrawVertexBuffer(vb * VertexBuffer, texture * Bitmap, start, end int, ptype PrimType) int {
    return int(C.al_draw_vertex_buffer(vb.toC(),
                texture.toC(), C.int(start), C.int(end), C.int(ptype)))
}

func DrawIndexedBuffer(vb * VertexBuffer, texture * Bitmap, ib * IndexBuffer, start, end int, ptype PrimType) int {
    return int(C.al_draw_indexed_buffer(vb.toC(),
                texture.toC(), ib.toC(), C.int(start), C.int(end), C.int(ptype)))     
}

// returns low level handle for the vertex buffer
func (vb *VertexBuffer) toC() * C.ALLEGRO_VERTEX_BUFFER {
    return vb.handle
}

// Wraps a C Allegro vertex buffer. Sets no finalizer.
func wrapVertexBufferRaw(handle *C.ALLEGRO_VERTEX_BUFFER) *VertexBuffer {
    if handle == nil {
        return nil
    }
    return &VertexBuffer{handle}
}

// Wraps a C Allegro vertex buffer. Sets a finalizer that calls Destroy.
func wrapVertexBuffer(handle *C.ALLEGRO_VERTEX_BUFFER) *VertexBuffer {
    vd := wrapVertexBufferRaw(handle)
    if vd != nil {
        runtime.SetFinalizer(vd, func(me * VertexBuffer) { me.Destroy() })
    }
    return vd
}

func CreateVertexBuffer(initial []Vertex, flags PrimBufferFlags) * VertexBuffer {
    size := len(initial)
    raw := C.al_create_vertex_buffer(nil, unsafe.Pointer(&initial[0]), C.int(size), C.int(flags))
    return wrapVertexBuffer(raw)
}

func (vb * VertexBuffer) Destroy() {
    if vb.handle != nil {
        C.al_destroy_vertex_buffer(vb.handle)
    }
    vb.handle  = nil
}

func (vb * VertexBuffer) Get(offset, length int) []Vertex {
    res := make([]Vertex, length)
    ptr := C.al_lock_vertex_buffer(vb.toC(), C.int(offset), C.int(length), C.ALLEGRO_LOCK_READONLY)
    if (ptr == nil) {
        res = nil
    } else { 
        for i := 0; i < length; i ++ {
            res[i] = *((*C.ALLEGRO_VERTEX)(unsafe.Pointer(uintptr(ptr) + uintptr(i * C.sizeof_ALLEGRO_VERTEX))))
        }
        C.al_unlock_vertex_buffer(vb.toC())
    }
    return res
}

func (vb * VertexBuffer) Set(offset int, v []Vertex) int {
    length  := len(v)
    res     := length
    ptr     := C.al_lock_vertex_buffer(vb.toC(), C.int(offset), C.int(length), C.ALLEGRO_LOCK_READWRITE)
    if (ptr == nil) {
        res = 0
    } else { 
        for i := 0; i < length; i ++ {
            *((*C.ALLEGRO_VERTEX)(unsafe.Pointer(uintptr(ptr) + uintptr(i * C.sizeof_ALLEGRO_VERTEX)))) = v[i]
        }
        C.al_unlock_vertex_buffer(vb.toC())  
    }
    return res
}

// returns low level handle for the vertex buffer
func (ib *IndexBuffer) toC() * C.ALLEGRO_INDEX_BUFFER {
    return ib.handle
}

// Wraps a C Allegro vertex buffer. Sets no finalizer.
func wrapIndexBufferRaw(handle *C.ALLEGRO_INDEX_BUFFER) *IndexBuffer {
    if handle == nil {
        return nil
    }
    return &IndexBuffer{handle}
}

// Wraps a C Allegro vertex buffer. Sets a finalizer that calls Destroy.
func wrapIndexBuffer(handle *C.ALLEGRO_INDEX_BUFFER) *IndexBuffer {
    ib := wrapIndexBufferRaw(handle)
    if ib != nil {
        runtime.SetFinalizer(ib, func(me * IndexBuffer) { me.Destroy() })
    }
    return ib
}

func CreateIndexBuffer(initial []int32, flags PrimBufferFlags) * IndexBuffer {
    size := len(initial)
    raw := C.al_create_index_buffer(4, unsafe.Pointer(&initial[0]), C.int(size), C.int(flags))
    return wrapIndexBuffer(raw)
}

func (ib * IndexBuffer) Destroy() {
    if ib.handle != nil {
        C.al_destroy_index_buffer(ib.handle)
    }
    ib.handle  = nil
}

func (ib * IndexBuffer) Get(offset, length int) []int32 {
    res := make([]int32, length)
    ptr := C.al_lock_index_buffer(ib.toC(), C.int(offset), C.int(length), C.ALLEGRO_LOCK_READONLY)
    if (ptr == nil) {
        res = nil
    } else { 
        for i := 0; i < length; i ++ {
            res[i] = int32(*((*C.int)(unsafe.Pointer(uintptr(ptr) + uintptr(i * C.sizeof_int)))))
        }
        C.al_unlock_index_buffer(ib.toC())
    }
    return res
}

func (ib * IndexBuffer) Set(offset int, v []int32) int {
    length  := len(v)
    res     := length
    ptr     := C.al_lock_index_buffer(ib.toC(), C.int(offset), C.int(length), C.ALLEGRO_LOCK_READWRITE)
    if (ptr == nil) {
        res = 0
    } else { 
        for i := 0; i < length; i ++ {
            *((*C.int)(unsafe.Pointer(uintptr(ptr) + uintptr(i * C.sizeof_int)))) = C.int(v[i])
        }
        C.al_unlock_index_buffer(ib.toC())  
    }
    return res
}


/* Callback primitives not implemented due to being a PITA. */

func DrawLine(x1, y1, x2, y2 float32, color Color, thickness float32) {
    C.al_draw_line(C.float(x1), C.float(y1), 
                   C.float(x2), C.float(y2), 
                   color.toC(), C.float(thickness))
} 

func DrawTriangle(x1, y1, x2, y2, x3, y3 float32, color Color, thickness float32) {
    C.al_draw_triangle(C.float(x1), C.float(y1), 
                   C.float(x2), C.float(y2), 
                   C.float(x3), C.float(y3),
                   color.toC(), C.float(thickness))
}

func DrawRectangle(x1, y1, x2, y2 float32, color Color, thickness float32) {
    C.al_draw_rectangle(C.float(x1), C.float(y1), 
                   C.float(x2), C.float(y2), 
                   color.toC(), C.float(thickness))
} 


func DrawRoundedRectangle(x1, y1, x2, y2, rx, ry float32, color Color, thickness float32) {
    C.al_draw_rounded_rectangle(C.float(x1), C.float(y1), 
                   C.float(x2), C.float(y2), 
                   C.float(rx), C.float(ry),
                   color.toC(), C.float(thickness))
}

func DrawCircle(cx, cy, r float32, color Color, thickness float32) {
    C.al_draw_circle(C.float(cx), C.float(cy), 
                   C.float(r), 
                   color.toC(), C.float(thickness))
}

func DrawEllipse(cx, cy, rx, ry float32, color Color, thickness float32) {
    C.al_draw_ellipse(C.float(cx), C.float(cy), 
                   C.float(rx), C.float(ry), 
                   color.toC(), C.float(thickness))
}

func DrawPieSlice(cx, cy, r, start, delta float32, color Color, thickness float32) {
    C.al_draw_pieslice(C.float(cx), C.float(cy), 
                   C.float(r),
                   C.float(start), C.float(delta),
                   color.toC(), C.float(thickness))
}

func DrawPolygon(vertices []float32, join LineJoin,  color Color, thickness float32, miter_limit float32) {
    C.al_draw_polygon((*C.float)(&vertices[0]), C.int(len(vertices)/2), C.int(join), 
                        color.toC(), C.float(thickness),
                        C.float(miter_limit))
}


func DrawArc(cx, cy, r, start, delta float32, color Color, thickness float32) {
    C.al_draw_arc(C.float(cx), C.float(cy), 
                   C.float(r),
                   C.float(start), C.float(delta),
                   color.toC(), C.float(thickness))
}



func DrawEllipticalArc(cx, cy, rx, ry, start, delta float32, color Color, thickness float32) {
    C.al_draw_elliptical_arc(C.float(cx), C.float(cy), 
                   C.float(rx), C.float(ry), 
                   C.float(start), C.float(delta),
                   color.toC(), C.float(thickness))
}


func DrawSpline(points [8]float32, color Color, thickness float32) {
    C.al_draw_spline((*C.float)(&points[0]),
                   color.toC(), C.float(thickness))
}


func DrawRibbon(points []float32, stride int, color Color, thickness float32) {
    C.al_draw_ribbon((*C.float)(&points[0]), C.int(stride),
                   color.toC(), C.float(thickness), C.int(len(points)/2))
}

func DrawFilledTriangle(x1, y1, x2, y2, x3, y3 float32, color Color) {
    C.al_draw_filled_triangle(C.float(x1), C.float(y1), 
                   C.float(x2), C.float(y2), 
                   C.float(x3), C.float(y3),
                   color.toC())
}

func DrawFilledRectangle(x1, y1, x2, y2 float32, color Color) {
    C.al_draw_filled_rectangle(C.float(x1), C.float(y1), 
                   C.float(x2), C.float(y2), 
                   color.toC())
} 


func DrawFilledRoundedRectangle(x1, y1, x2, y2, rx, ry float32, color Color) {
    C.al_draw_filled_rounded_rectangle(C.float(x1), C.float(y1), 
                   C.float(x2), C.float(y2), 
                   C.float(rx), C.float(ry),
                   color.toC())
}

func DrawFilledCircle(cx, cy, r float32, color Color) {
    C.al_draw_filled_circle(C.float(cx), C.float(cy), 
                   C.float(r), 
                   color.toC())
}

func DrawFilledEllipse(cx, cy, rx, ry float32, color Color) {
    C.al_draw_filled_ellipse(C.float(cx), C.float(cy), 
                   C.float(rx), C.float(ry), 
                   color.toC())
}

func DrawFilledPieSlice(cx, cy, r, start, delta float32, color Color) {
    C.al_draw_filled_pieslice(C.float(cx), C.float(cy), 
                   C.float(r),
                   C.float(start), C.float(delta),
                   color.toC())
}

func DrawFilledPolygon(vertices []float32, join LineJoin,  color Color) {
    C.al_draw_filled_polygon((*C.float)(&vertices[0]), C.int(len(vertices)/2), color.toC())
}

func (vert * Vertex) Init(x, y, z, u, v float32, color Color) {
    vert.x      = C.float(x)
    vert.y      = C.float(y)
    vert.z      = C.float(z)
    vert.u      = C.float(u)
    vert.v      = C.float(v)
    vert.color  = color.toC()
}


var later = `
/*
*High level primitives
*/

ALLEGRO_PRIM_FUNC(void, al_draw_filled_triangle, (float x1, float y1, float x2, float y2, float x3, float y3, ALLEGRO_COLOR color));
ALLEGRO_PRIM_FUNC(void, al_draw_filled_rectangle, (float x1, float y1, float x2, float y2, ALLEGRO_COLOR color));
ALLEGRO_PRIM_FUNC(void, al_draw_filled_ellipse, (float cx, float cy, float rx, float ry, ALLEGRO_COLOR color));
ALLEGRO_PRIM_FUNC(void, al_draw_filled_circle, (float cx, float cy, float r, ALLEGRO_COLOR color));
ALLEGRO_PRIM_FUNC(void, al_draw_filled_pieslice, (float cx, float cy, float r, float start_theta, float delta_theta, ALLEGRO_COLOR color));
ALLEGRO_PRIM_FUNC(void, al_draw_filled_rounded_rectangle, (float x1, float y1, float x2, float y2, float rx, float ry, ALLEGRO_COLOR color));

ALLEGRO_PRIM_FUNC(void, al_draw_polyline, (const float* vertices, int vertex_stride, int vertex_count, int join_style, int cap_style, ALLEGRO_COLOR color, float thickness, float miter_limit));

ALLEGRO_PRIM_FUNC(void, al_draw_polygon, (const float* vertices, int vertex_count, int join_style, ALLEGRO_COLOR color, float thickness, float miter_limit));
ALLEGRO_PRIM_FUNC(void, al_draw_filled_polygon, (const float* vertices, int vertex_count, ALLEGRO_COLOR color));
ALLEGRO_PRIM_FUNC(void, al_draw_filled_polygon_with_holes, (const float* vertices, const int* vertex_counts, ALLEGRO_COLOR color));


#ifdef __cplusplus
}
#endif

#endif
`

