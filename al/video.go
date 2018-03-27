// Video extension
package al

/*
#cgo pkg-config: allegro_video-5
#cgo CFLAGS: -I/usr/local/include
#cgo linux LDFLAGS: -lc_nonshared
#include <stdlib.h>
#include <allegro5/allegro.h>
#include <allegro5/allegro_video.h>
#include "helpers.h"
*/
import "C"

import "runtime"
// import "unsafe"

type VideoEventType C.enum_ALLEGRO_VIDEO_EVENT_TYPE

const (
   EVENT_VIDEO_FRAME_SHOW       = VideoEventType(C.ALLEGRO_EVENT_VIDEO_FRAME_SHOW)
   EVENT_VIDEO_FRAME_FINISHED   = VideoEventType(C.ALLEGRO_EVENT_VIDEO_FINISHED)
   eVENT_VIDEO_FRAME_FINISHED   = VideoEventType(C._ALLEGRO_EVENT_VIDEO_SEEK)
)

type VideoPositionType C.ALLEGRO_VIDEO_POSITION_TYPE

const (
   VIDEO_POSITION_ACTUAL       = VideoPositionType(C.ALLEGRO_VIDEO_POSITION_ACTUAL)
   VIDEO_POSITION_VIDEO_DECODE = VideoPositionType(C.ALLEGRO_VIDEO_POSITION_VIDEO_DECODE)
   VIDEO_POSITION_AUDIO_DECODE = VideoPositionType(C.ALLEGRO_VIDEO_POSITION_AUDIO_DECODE)
)

type Video struct {
    handle * C.ALLEGRO_VIDEO
}

// Converts a video to it's underlying C pointer
func (self * Video) toC() *C.ALLEGRO_VIDEO {
    return (*C.ALLEGRO_VIDEO)(self.handle)
}

// Destroys the video.
func (self *Video) Destroy() {
    if self.handle != nil {
        C.al_close_video(self.toC())
    }
    self.handle = nil
}

// Wraps a C video into a go video
func wrapVideoRaw(data *C.ALLEGRO_VIDEO) *Video {
    if data == nil {
        return nil
    }
    return &Video{data}
}

// Sets up a finalizer for this Video that calls Destroy()
func (self *Video) SetDestroyFinalizer() *Video {
    if self != nil {
        runtime.SetFinalizer(self, func(me *Video) { me.Destroy() })
    }
    return self
}

// Wraps a C video into a go video and sets up a finalizer that calls Destroy()
func wrapVideo(data *C.ALLEGRO_VIDEO) *Video {
    self := wrapVideoRaw(data)
    return self.SetDestroyFinalizer()
}


func OpenVideo(filename string) * Video {
    cfilename := cstr(filename) ; defer cstrFree(cfilename)
    return wrapVideo(C.al_open_video(cfilename)) 
}

func (video * Video) Close() {
    video.Destroy()
}

func (video * Video) Start(mixer * Mixer) {
    C.al_start_video(video.toC(), mixer.toC())
}

func (video * Video) StartWithVoice(voice * Voice) {
    C.al_start_video_with_voice(video.toC(), voice.toC())
}

func (video * Video) EventSource() * EventSource {
    return wrapEventSourceRaw(C.al_get_video_event_source(video.toC()))
} 

func (video * Video) SetPlaying(playing bool) {
    C.al_set_video_playing(video.toC(), C.bool(playing))
} 

func (video * Video) Playing() bool {
    return bool(C.al_is_video_playing(video.toC()))
} 

func (video * Video) AudioRate() float64 {
    return float64(C.al_get_video_audio_rate(video.toC()))
}

func (video * Video) FPS() float64 {
    return float64(C.al_get_video_fps(video.toC()))
}

func (video * Video) ScaledWidth() float32 {
    return float32(C.al_get_video_scaled_width(video.toC()))
} 

func (video * Video) ScaledHeight() float32 {
    return float32(C.al_get_video_scaled_height(video.toC()))
}

func (video * Video) Frame() * Bitmap {
    return wrapBitmapRaw(C.al_get_video_frame(video.toC()))
}

func (vpt VideoPositionType) toC() C.ALLEGRO_VIDEO_POSITION_TYPE {
    return C.ALLEGRO_VIDEO_POSITION_TYPE(vpt)
}

func (video * Video) Position(vpt VideoPositionType) float64 {
    return float64(C.al_get_video_position(video.toC(), vpt.toC()))    
}


func (video * Video) ActualPosition(vpt VideoPositionType) float64 {
    return video.Position(VIDEO_POSITION_ACTUAL) 
}

func (video * Video) AudioDecodePosition(vpt VideoPositionType) float64 {
    return video.Position(VIDEO_POSITION_AUDIO_DECODE) 
}

func (video * Video) VideoDecodePosition(vpt VideoPositionType) float64 {
    return video.Position(VIDEO_POSITION_VIDEO_DECODE) 
}

func (video * Video) Seek(position_in_seconds float64) bool {
    return bool(C.al_seek_video(video.toC(), C.double(position_in_seconds)))    
}

func InitVideoAddon() bool {
    return bool(C.al_init_video_addon())
}

func ShutdownVideoAddon() {
    C.al_shutdown_video_addon()
}

func VideoVersion() uint32 {
    return uint32(C.al_get_allegro_video_version())
}
