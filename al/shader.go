package al

/*
#include <stdlib.h>
#include <allegro5/allegro.h>
#include "helpers.h"
*/
import "C"

// import "unsafe"
import "runtime"

type ShaderType C.ALLEGRO_SHADER_TYPE

const (
   VertexShader = ShaderType(C.ALLEGRO_VERTEX_SHADER)
   PixelShader  = ShaderType(C.ALLEGRO_PIXEL_SHADER)
)

func (st ShaderType) toC() C.ALLEGRO_SHADER_TYPE {
    return (C.ALLEGRO_SHADER_TYPE)(st)
}

type ShaderPlatform C.ALLEGRO_SHADER_PLATFORM

const (
   ShaderAuto = ShaderPlatform(C.ALLEGRO_SHADER_AUTO)
   ShaderGLSL = ShaderPlatform(C.ALLEGRO_SHADER_GLSL)
   ShaderHLSL = ShaderPlatform(C.ALLEGRO_SHADER_HLSL)
)

func (sp ShaderPlatform) toC() C.ALLEGRO_SHADER_PLATFORM {
    return (C.ALLEGRO_SHADER_PLATFORM)(sp)
}


/* Shader variable names */
const  (
    ShaderVarColor          = "al_color"
    ShaderVarPos            = "al_pos"
    ShaderVarProjviewMatrix = "al_projview_matrix"
    ShaderVarTex            = "al_tex"
    ShaderVarTextcoord      = "al_texcoord"
    ShaderVarTextMatrix     = "al_tex_matrix"
    ShaderVarUserAttr       = "al_user_attr_"
    ShaderVarUseTex         = "al_use_tex"
    ShaderVarUseTexMatrix   = "al_use_tex_matrix"
)


type Shader struct {
    handle * C.ALLEGRO_SHADER
}

// Converts a shader to it's underlying C pointer
func (self * Shader) toC() *C.ALLEGRO_SHADER {
    return (*C.ALLEGRO_SHADER)(self.handle)
}

// Destroys the shader.
func (self *Shader) Destroy() {
    if self.handle != nil {
        C.al_destroy_shader(self.toC())
    }
    self.handle = nil
}

// Wraps a C shader into a go shader
func wrapShaderRaw(data *C.ALLEGRO_SHADER) *Shader {
    if data == nil {
        return nil
    }
    return &Shader{data}
}

// Sets up a finalizer for this Shader that calls Destroy()
func (self *Shader) SetDestroyFinalizer() *Shader {
    if self != nil {
        runtime.SetFinalizer(self, func(me *Shader) { me.Destroy() })
    }
    return self
}

// Wraps a C shader into a go shader and sets up a finalizer that calls Destroy()
func wrapShader(data *C.ALLEGRO_SHADER) *Shader {
    self := wrapShaderRaw(data)
    return self.SetDestroyFinalizer()
}


func (sp ShaderPlatform) Create() *Shader {
    return wrapShader(C.al_create_shader(sp.toC()))
}

func (sh * Shader) AttachSource(shatype ShaderType, source string) bool {
    csource := cstr(source); defer cstrFree(csource)
    return bool(C.al_attach_shader_source(sh.toC(), shatype.toC(), csource))
}

func (sh * Shader) AttachSourceFile(shatype ShaderType, source string) bool {
    csource := cstr(source); defer cstrFree(csource)
    return bool(C.al_attach_shader_source_file(sh.toC(), shatype.toC(), csource))
}

func (sh * Shader) Build() bool {
    return bool(C.al_build_shader(sh.toC()))
}

func (sh * Shader) Log() string {
    return C.GoString(C.al_get_shader_log(sh.toC()))
}

func (sh * Shader) Platform() ShaderPlatform {
    return ShaderPlatform(C.al_get_shader_platform(sh.toC()))
}

func (sh * Shader) Use() bool { 
    return bool(C.al_use_shader(sh.toC()))
}

func SetShaderSampler(name string, bitmap * Bitmap, unit int) bool {
    cname := cstr(name); defer cstrFree(cname) 
    return bool(C.al_set_shader_sampler(cname, bitmap.toC(), C.int(unit)))
}

/*
func (sh * Shader) SetMatrix(name string, matrix * Matrix) bool {
    cname := cstr(name); defer cstrFree(cname) 
    return bool(C.al_set_shader_matrix(sh.toC(), cname, matrix.toC()))
}
*/ 

func SetShaderInt(name string, i int) bool {
    cname := cstr(name); defer cstrFree(cname) 
    return bool(C.al_set_shader_int(cname, C.int(i)))
}

func SetShaderFloat(name string, f float32) bool {
    cname := cstr(name); defer cstrFree(cname) 
    return bool(C.al_set_shader_float(cname, C.float(f)))
}

func SetShaderBool(name string, b bool) bool {
    cname := cstr(name); defer cstrFree(cname) 
    return bool(C.al_set_shader_bool(cname, C.bool(b)))
}

func (sh * Shader) SetIntVector(name string, nc int, i []int) bool {
    cname := cstr(name); defer cstrFree(cname)
    csize, cvec := CInts(i) ; defer CIntsFree(csize, cvec) 
    /* XXX, I doubt this will work for nc > 1 ... */
    return bool(C.al_set_shader_int_vector(cname, C.int(nc), cvec, csize / C.int(nc)))
}

func (sh * Shader) SetFloatVector(name string, nc int, f []float32) bool {
    cname := cstr(name); defer cstrFree(cname) 
    csize, cvec := CFloats(f) ; defer CFloatsFree(csize, cvec) 
    /* XXX, I doubt this will work for nc > 1 ... */
    return bool(C.al_set_shader_float_vector(cname, C.int(nc), cvec, csize / C.int(nc)))
}


func GetDefaultShaderSource(pla ShaderPlatform, ty ShaderType) string {
    return C.GoString(C.al_get_default_shader_source(pla.toC(), ty.toC()))
}


