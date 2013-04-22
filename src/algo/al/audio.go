// audio
package al

/*
#cgo pkg-config: allegro_audio-5
#cgo CFLAGS: -I/usr/local/include
#cgo linux LDFLAGS: -lc_nonshared
#include <stdlib.h>
#include <allegro5/allegro.h>
#include <allegro5/allegro_audio.h>
#include "helpers.h"
*/
import "C"
import "runtime"

// User event type emitted when a stream fragment is ready to be
// refilled with more audio data.
const (
	EVENT_AUDIO_STREAM_FRAGMENT   = C.ALLEGRO_EVENT_AUDIO_STREAM_FRAGMENT
	EVENT_AUDIO_STREAM_FINISHED   = C.ALLEGRO_EVENT_AUDIO_STREAM_FINISHED
	EVENT_AUDIO_RECORDER_FRAGMENT = C.ALLEGRO_EVENT_AUDIO_RECORDER_FRAGMENT
)

// Converts wrapper Event pointer to C Allegro audio recorder event
func (self *Event) AUDIO_RECORDER_EVENT() *C.ALLEGRO_AUDIO_RECORDER_EVENT {
	return (*C.ALLEGRO_AUDIO_RECORDER_EVENT)(self.toPointer())
}

/*
TODO?
struct ALLEGRO_AUDIO_RECORDER_EVENT
{
   _AL_EVENT_HEADER(struct ALLEGRO_AUDIO_RECORDER)
   struct ALLEGRO_USER_EVENT_DESCRIPTOR *__internal__descr;
   void *buffer;
   unsigned int samples;
};
*/

type AudioDepth int

const (
	AUDIO_DEPTH_INT8     AudioDepth = C.ALLEGRO_AUDIO_DEPTH_INT8
	AUDIO_DEPTH_INT16    AudioDepth = C.ALLEGRO_AUDIO_DEPTH_INT16
	AUDIO_DEPTH_INT24    AudioDepth = C.ALLEGRO_AUDIO_DEPTH_INT24
	AUDIO_DEPTH_UINT8    AudioDepth = C.ALLEGRO_AUDIO_DEPTH_UINT8
	AUDIO_DEPTH_UINT16   AudioDepth = C.ALLEGRO_AUDIO_DEPTH_UINT16
	AUDIO_DEPTH_UINT24   AudioDepth = C.ALLEGRO_AUDIO_DEPTH_UINT24
	AUDIO_DEPTH_FLOAT32  AudioDepth = C.ALLEGRO_AUDIO_DEPTH_FLOAT32
	AUDIO_DEPTH_UNSIGNED AudioDepth = C.ALLEGRO_AUDIO_DEPTH_UNSIGNED
)

/*
   Speaker configuration (mono, stereo, 2.1, 3, etc). With regards to
   * behavior, most of this code makes no distinction between, say, 4.1 and
   * 5 speaker setups.. they both have 5 "channels". However, users would
   * like the distinction, and later when the higher-level stuff is added,
   * the differences will become more important. (v>>4)+(v&0xF) should yield
   * the total channel count.
*/

type ChannelConf int

const (
	CHANNEL_CONF_1   ChannelConf = C.ALLEGRO_CHANNEL_CONF_1
	CHANNEL_CONF_2   ChannelConf = C.ALLEGRO_CHANNEL_CONF_2
	CHANNEL_CONF_3   ChannelConf = C.ALLEGRO_CHANNEL_CONF_3
	CHANNEL_CONF_4   ChannelConf = C.ALLEGRO_CHANNEL_CONF_4
	CHANNEL_CONF_5_1 ChannelConf = C.ALLEGRO_CHANNEL_CONF_5_1
	CHANNEL_CONF_6_1 ChannelConf = C.ALLEGRO_CHANNEL_CONF_6_1
	CHANNEL_CONF_7_1 ChannelConf = C.ALLEGRO_CHANNEL_CONF_7_1
	MAX_CHANNELS     ChannelConf = C.ALLEGRO_MAX_CHANNELS
)

type PlayMode int

const (
	PLAYMODE_ONCE  PlayMode = C.ALLEGRO_PLAYMODE_ONCE
	PLAYMODE_LOOP  PlayMode = C.ALLEGRO_PLAYMODE_LOOP
	PLAYMODE_BIDIR PlayMode = C.ALLEGRO_PLAYMODE_BIDIR
)

type MixerQuality int

const (
	MIXER_QUALITY_POINT  MixerQuality = C.ALLEGRO_MIXER_QUALITY_POINT
	MIXER_QUALITY_LINEAR MixerQuality = C.ALLEGRO_MIXER_QUALITY_LINEAR
	MIXER_QUALITY_CUBIC  MixerQuality = C.ALLEGRO_MIXER_QUALITY_CUBIC
)

const AUDIO_PAN_NONE = -1000.0

type AudioEventType int

const (
	ALLEGRO_EVENT_AUDIO_ROUTE_CHANGE = C.ALLEGRO_EVENT_AUDIO_ROUTE_CHANGE
	EVENT_AUDIO_INTERRUPTION         = C.ALLEGRO_EVENT_AUDIO_INTERRUPTION
	EVENT_AUDIO_END_INTERRUPTION     = C.ALLEGRO_EVENT_AUDIO_END_INTERRUPTION
)

type Sample struct {
	handle *C.ALLEGRO_SAMPLE
}

type SampleId C.ALLEGRO_SAMPLE_ID

type SampleInstance struct {
	handle *C.ALLEGRO_SAMPLE_INSTANCE
}

type AudioStream C.ALLEGRO_AUDIO_STREAM
type Mixer C.ALLEGRO_MIXER
type Voice C.ALLEGRO_VOICE
type AudioRecorder C.ALLEGRO_AUDIO_RECORDER

func (self *Sample) toC() *C.ALLEGRO_SAMPLE {
	return (*C.ALLEGRO_SAMPLE)(self.handle)
}

func (self *Sample) Destroy() {
	C.al_destroy_sample(self.toC())
}

func wrapSampleRaw(sample *C.ALLEGRO_SAMPLE) *Sample {
	if sample == nil {
		return nil
	}
	return &Sample{sample}
}

func wrapSample(sample *C.ALLEGRO_SAMPLE) *Sample {
	self := wrapSampleRaw(sample)
	if self != nil {
		runtime.SetFinalizer(self, func(me *Sample) { me.Destroy() })
	}
	return self
}

/*

ALLEGRO_KCM_AUDIO_FUNC(ALLEGRO_SAMPLE *, al_create_sample, (void *buf,
      unsigned int samples, unsigned int freq, AudioDepth depth,
      ALLEGRO_CHANNEL_CONF chan_conf, bool free_buf));
ALLEGRO_KCM_AUDIO_FUNC(void, al_destroy_sample, (ALLEGRO_SAMPLE *spl));



ALLEGRO_KCM_AUDIO_FUNC(ALLEGRO_SAMPLE_INSTANCE*, al_create_sample_instance, (
      ALLEGRO_SAMPLE *data));
ALLEGRO_KCM_AUDIO_FUNC(void, al_destroy_sample_instance, (
      ALLEGRO_SAMPLE_INSTANCE *spl));

ALLEGRO_KCM_AUDIO_FUNC(unsigned int, al_get_sample_frequency, (const ALLEGRO_SAMPLE *spl));
ALLEGRO_KCM_AUDIO_FUNC(unsigned int, al_get_sample_length, (const ALLEGRO_SAMPLE *spl));
ALLEGRO_KCM_AUDIO_FUNC(ALLEGRO_AUDIO_DEPTH, al_get_sample_depth, (const ALLEGRO_SAMPLE *spl));
ALLEGRO_KCM_AUDIO_FUNC(ALLEGRO_CHANNEL_CONF, al_get_sample_channels, (const ALLEGRO_SAMPLE *spl));
ALLEGRO_KCM_AUDIO_FUNC(void *, al_get_sample_data, (const ALLEGRO_SAMPLE *spl));

ALLEGRO_KCM_AUDIO_FUNC(unsigned int, al_get_sample_instance_frequency, (const ALLEGRO_SAMPLE_INSTANCE *spl));
ALLEGRO_KCM_AUDIO_FUNC(unsigned int, al_get_sample_instance_length, (const ALLEGRO_SAMPLE_INSTANCE *spl));
ALLEGRO_KCM_AUDIO_FUNC(unsigned int, al_get_sample_instance_position, (const ALLEGRO_SAMPLE_INSTANCE *spl));

ALLEGRO_KCM_AUDIO_FUNC(float, al_get_sample_instance_speed, (const ALLEGRO_SAMPLE_INSTANCE *spl));
ALLEGRO_KCM_AUDIO_FUNC(float, al_get_sample_instance_gain, (const ALLEGRO_SAMPLE_INSTANCE *spl));
ALLEGRO_KCM_AUDIO_FUNC(float, al_get_sample_instance_pan, (const ALLEGRO_SAMPLE_INSTANCE *spl));
ALLEGRO_KCM_AUDIO_FUNC(float, al_get_sample_instance_time, (const ALLEGRO_SAMPLE_INSTANCE *spl));

ALLEGRO_KCM_AUDIO_FUNC(ALLEGRO_AUDIO_DEPTH, al_get_sample_instance_depth, (const ALLEGRO_SAMPLE_INSTANCE *spl));
ALLEGRO_KCM_AUDIO_FUNC(ALLEGRO_CHANNEL_CONF, al_get_sample_instance_channels, (const ALLEGRO_SAMPLE_INSTANCE *spl));
ALLEGRO_KCM_AUDIO_FUNC(ALLEGRO_PLAYMODE, al_get_sample_instance_playmode, (const ALLEGRO_SAMPLE_INSTANCE *spl));

ALLEGRO_KCM_AUDIO_FUNC(bool, al_get_sample_instance_playing, (const ALLEGRO_SAMPLE_INSTANCE *spl));
ALLEGRO_KCM_AUDIO_FUNC(bool, al_get_sample_instance_attached, (const ALLEGRO_SAMPLE_INSTANCE *spl));

ALLEGRO_KCM_AUDIO_FUNC(bool, al_set_sample_instance_position, (ALLEGRO_SAMPLE_INSTANCE *spl, unsigned int val));
ALLEGRO_KCM_AUDIO_FUNC(bool, al_set_sample_instance_length, (ALLEGRO_SAMPLE_INSTANCE *spl, unsigned int val));

ALLEGRO_KCM_AUDIO_FUNC(bool, al_set_sample_instance_speed, (ALLEGRO_SAMPLE_INSTANCE *spl, float val));
ALLEGRO_KCM_AUDIO_FUNC(bool, al_set_sample_instance_gain, (ALLEGRO_SAMPLE_INSTANCE *spl, float val));
ALLEGRO_KCM_AUDIO_FUNC(bool, al_set_sample_instance_pan, (ALLEGRO_SAMPLE_INSTANCE *spl, float val));

ALLEGRO_KCM_AUDIO_FUNC(bool, al_set_sample_instance_playmode, (ALLEGRO_SAMPLE_INSTANCE *spl, ALLEGRO_PLAYMODE val));

ALLEGRO_KCM_AUDIO_FUNC(bool, al_set_sample_instance_playing, (ALLEGRO_SAMPLE_INSTANCE *spl, bool val));
ALLEGRO_KCM_AUDIO_FUNC(bool, al_detach_sample_instance, (ALLEGRO_SAMPLE_INSTANCE *spl));

ALLEGRO_KCM_AUDIO_FUNC(bool, al_set_sample, (ALLEGRO_SAMPLE_INSTANCE *spl, ALLEGRO_SAMPLE *data));
ALLEGRO_KCM_AUDIO_FUNC(ALLEGRO_SAMPLE *, al_get_sample, (ALLEGRO_SAMPLE_INSTANCE *spl));
ALLEGRO_KCM_AUDIO_FUNC(bool, al_play_sample_instance, (ALLEGRO_SAMPLE_INSTANCE *spl));
ALLEGRO_KCM_AUDIO_FUNC(bool, al_stop_sample_instance, (ALLEGRO_SAMPLE_INSTANCE *spl));



ALLEGRO_KCM_AUDIO_FUNC(ALLEGRO_AUDIO_STREAM*, al_create_audio_stream, (size_t buffer_count,
      unsigned int samples, unsigned int freq,
      AudioDepth depth, ALLEGRO_CHANNEL_CONF chan_conf));
ALLEGRO_KCM_AUDIO_FUNC(void, al_destroy_audio_stream, (ALLEGRO_AUDIO_STREAM *stream));
ALLEGRO_KCM_AUDIO_FUNC(void, al_drain_audio_stream, (ALLEGRO_AUDIO_STREAM *stream));

ALLEGRO_KCM_AUDIO_FUNC(unsigned int, al_get_audio_stream_frequency, (const ALLEGRO_AUDIO_STREAM *stream));
ALLEGRO_KCM_AUDIO_FUNC(unsigned int, al_get_audio_stream_length, (const ALLEGRO_AUDIO_STREAM *stream));
ALLEGRO_KCM_AUDIO_FUNC(unsigned int, al_get_audio_stream_fragments, (const ALLEGRO_AUDIO_STREAM *stream));
ALLEGRO_KCM_AUDIO_FUNC(unsigned int, al_get_available_audio_stream_fragments, (const ALLEGRO_AUDIO_STREAM *stream));

ALLEGRO_KCM_AUDIO_FUNC(float, al_get_audio_stream_speed, (const ALLEGRO_AUDIO_STREAM *stream));
ALLEGRO_KCM_AUDIO_FUNC(float, al_get_audio_stream_gain, (const ALLEGRO_AUDIO_STREAM *stream));
ALLEGRO_KCM_AUDIO_FUNC(float, al_get_audio_stream_pan, (const ALLEGRO_AUDIO_STREAM *stream));

ALLEGRO_KCM_AUDIO_FUNC(ALLEGRO_CHANNEL_CONF, al_get_audio_stream_channels, (const ALLEGRO_AUDIO_STREAM *stream));
ALLEGRO_KCM_AUDIO_FUNC(ALLEGRO_AUDIO_DEPTH, al_get_audio_stream_depth, (const ALLEGRO_AUDIO_STREAM *stream));
ALLEGRO_KCM_AUDIO_FUNC(ALLEGRO_PLAYMODE, al_get_audio_stream_playmode, (const ALLEGRO_AUDIO_STREAM *stream));

ALLEGRO_KCM_AUDIO_FUNC(bool, al_get_audio_stream_playing, (const ALLEGRO_AUDIO_STREAM *spl));
ALLEGRO_KCM_AUDIO_FUNC(bool, al_get_audio_stream_attached, (const ALLEGRO_AUDIO_STREAM *spl));

ALLEGRO_KCM_AUDIO_FUNC(void *, al_get_audio_stream_fragment, (const ALLEGRO_AUDIO_STREAM *stream));

ALLEGRO_KCM_AUDIO_FUNC(bool, al_set_audio_stream_speed, (ALLEGRO_AUDIO_STREAM *stream, float val));
ALLEGRO_KCM_AUDIO_FUNC(bool, al_set_audio_stream_gain, (ALLEGRO_AUDIO_STREAM *stream, float val));
ALLEGRO_KCM_AUDIO_FUNC(bool, al_set_audio_stream_pan, (ALLEGRO_AUDIO_STREAM *stream, float val));

ALLEGRO_KCM_AUDIO_FUNC(bool, al_set_audio_stream_playmode, (ALLEGRO_AUDIO_STREAM *stream, ALLEGRO_PLAYMODE val));

ALLEGRO_KCM_AUDIO_FUNC(bool, al_set_audio_stream_playing, (ALLEGRO_AUDIO_STREAM *stream, bool val));
ALLEGRO_KCM_AUDIO_FUNC(bool, al_detach_audio_stream, (ALLEGRO_AUDIO_STREAM *stream));
ALLEGRO_KCM_AUDIO_FUNC(bool, al_set_audio_stream_fragment, (ALLEGRO_AUDIO_STREAM *stream, void *val));

ALLEGRO_KCM_AUDIO_FUNC(bool, al_rewind_audio_stream, (ALLEGRO_AUDIO_STREAM *stream));
ALLEGRO_KCM_AUDIO_FUNC(bool, al_seek_audio_stream_secs, (ALLEGRO_AUDIO_STREAM *stream, double time));
ALLEGRO_KCM_AUDIO_FUNC(double, al_get_audio_stream_position_secs, (ALLEGRO_AUDIO_STREAM *stream));
ALLEGRO_KCM_AUDIO_FUNC(double, al_get_audio_stream_length_secs, (ALLEGRO_AUDIO_STREAM *stream));
ALLEGRO_KCM_AUDIO_FUNC(bool, al_set_audio_stream_loop_secs, (ALLEGRO_AUDIO_STREAM *stream, double start, double end));

ALLEGRO_KCM_AUDIO_FUNC(ALLEGRO_EVENT_SOURCE *, al_get_audio_stream_event_source, (ALLEGRO_AUDIO_STREAM *stream));


ALLEGRO_KCM_AUDIO_FUNC(ALLEGRO_MIXER*, al_create_mixer, (unsigned int freq,
      AudioDepth depth, ALLEGRO_CHANNEL_CONF chan_conf));
ALLEGRO_KCM_AUDIO_FUNC(void, al_destroy_mixer, (ALLEGRO_MIXER *mixer));
ALLEGRO_KCM_AUDIO_FUNC(bool, al_attach_sample_instance_to_mixer, (
   ALLEGRO_SAMPLE_INSTANCE *stream, ALLEGRO_MIXER *mixer));
ALLEGRO_KCM_AUDIO_FUNC(bool, al_attach_audio_stream_to_mixer, (ALLEGRO_AUDIO_STREAM *stream,
   ALLEGRO_MIXER *mixer));
ALLEGRO_KCM_AUDIO_FUNC(bool, al_attach_mixer_to_mixer, (ALLEGRO_MIXER *stream,
   ALLEGRO_MIXER *mixer));
ALLEGRO_KCM_AUDIO_FUNC(bool, al_set_mixer_postprocess_callback, (
      ALLEGRO_MIXER *mixer,
      void (*cb)(void *buf, unsigned int samples, void *data),
      void *data));

ALLEGRO_KCM_AUDIO_FUNC(unsigned int, al_get_mixer_frequency, (const ALLEGRO_MIXER *mixer));
ALLEGRO_KCM_AUDIO_FUNC(ALLEGRO_CHANNEL_CONF, al_get_mixer_channels, (const ALLEGRO_MIXER *mixer));
ALLEGRO_KCM_AUDIO_FUNC(ALLEGRO_AUDIO_DEPTH, al_get_mixer_depth, (const ALLEGRO_MIXER *mixer));
ALLEGRO_KCM_AUDIO_FUNC(ALLEGRO_MIXER_QUALITY, al_get_mixer_quality, (const ALLEGRO_MIXER *mixer));
ALLEGRO_KCM_AUDIO_FUNC(float, al_get_mixer_gain, (const ALLEGRO_MIXER *mixer));
ALLEGRO_KCM_AUDIO_FUNC(bool, al_get_mixer_playing, (const ALLEGRO_MIXER *mixer));
ALLEGRO_KCM_AUDIO_FUNC(bool, al_get_mixer_attached, (const ALLEGRO_MIXER *mixer));
ALLEGRO_KCM_AUDIO_FUNC(bool, al_set_mixer_frequency, (ALLEGRO_MIXER *mixer, unsigned int val));
ALLEGRO_KCM_AUDIO_FUNC(bool, al_set_mixer_quality, (ALLEGRO_MIXER *mixer, ALLEGRO_MIXER_QUALITY val));
ALLEGRO_KCM_AUDIO_FUNC(bool, al_set_mixer_gain, (ALLEGRO_MIXER *mixer, float gain));
ALLEGRO_KCM_AUDIO_FUNC(bool, al_set_mixer_playing, (ALLEGRO_MIXER *mixer, bool val));
ALLEGRO_KCM_AUDIO_FUNC(bool, al_detach_mixer, (ALLEGRO_MIXER *mixer));


ALLEGRO_KCM_AUDIO_FUNC(ALLEGRO_VOICE*, al_create_voice, (unsigned int freq,
      AudioDepth depth,
      ALLEGRO_CHANNEL_CONF chan_conf));
ALLEGRO_KCM_AUDIO_FUNC(void, al_destroy_voice, (ALLEGRO_VOICE *voice));
ALLEGRO_KCM_AUDIO_FUNC(bool, al_attach_sample_instance_to_voice, (
   ALLEGRO_SAMPLE_INSTANCE *stream, ALLEGRO_VOICE *voice));
ALLEGRO_KCM_AUDIO_FUNC(bool, al_attach_audio_stream_to_voice, (
   ALLEGRO_AUDIO_STREAM *stream, ALLEGRO_VOICE *voice ));
ALLEGRO_KCM_AUDIO_FUNC(bool, al_attach_mixer_to_voice, (ALLEGRO_MIXER *mixer,
   ALLEGRO_VOICE *voice));
ALLEGRO_KCM_AUDIO_FUNC(void, al_detach_voice, (ALLEGRO_VOICE *voice));

ALLEGRO_KCM_AUDIO_FUNC(unsigned int, al_get_voice_frequency, (const ALLEGRO_VOICE *voice));
ALLEGRO_KCM_AUDIO_FUNC(unsigned int, al_get_voice_position, (const ALLEGRO_VOICE *voice));
ALLEGRO_KCM_AUDIO_FUNC(ALLEGRO_CHANNEL_CONF, al_get_voice_channels, (const ALLEGRO_VOICE *voice));
ALLEGRO_KCM_AUDIO_FUNC(ALLEGRO_AUDIO_DEPTH, al_get_voice_depth, (const ALLEGRO_VOICE *voice));
ALLEGRO_KCM_AUDIO_FUNC(bool, al_get_voice_playing, (const ALLEGRO_VOICE *voice));
ALLEGRO_KCM_AUDIO_FUNC(bool, al_set_voice_position, (ALLEGRO_VOICE *voice, unsigned int val));
ALLEGRO_KCM_AUDIO_FUNC(bool, al_set_voice_playing, (ALLEGRO_VOICE *voice, bool val));


ALLEGRO_KCM_AUDIO_FUNC(bool, al_install_audio, (void));
ALLEGRO_KCM_AUDIO_FUNC(void, al_uninstall_audio, (void));
ALLEGRO_KCM_AUDIO_FUNC(bool, al_is_audio_installed, (void));
ALLEGRO_KCM_AUDIO_FUNC(uint32_t, al_get_allegro_audio_version, (void));

ALLEGRO_KCM_AUDIO_FUNC(size_t, al_get_channel_count, (ALLEGRO_CHANNEL_CONF conf));
ALLEGRO_KCM_AUDIO_FUNC(size_t, al_get_audio_depth_size, (ALLEGRO_AUDIO_DEPTH conf));


ALLEGRO_KCM_AUDIO_FUNC(bool, al_reserve_samples, (int reserve_samples));
ALLEGRO_KCM_AUDIO_FUNC(ALLEGRO_MIXER *, al_get_default_mixer, (void));
ALLEGRO_KCM_AUDIO_FUNC(bool, al_set_default_mixer, (ALLEGRO_MIXER *mixer));
ALLEGRO_KCM_AUDIO_FUNC(bool, al_restore_default_mixer, (void));
ALLEGRO_KCM_AUDIO_FUNC(bool, al_play_sample, (ALLEGRO_SAMPLE *data,
      float gain, float pan, float speed, ALLEGRO_PLAYMODE loop, ALLEGRO_SAMPLE_ID *ret_id));
ALLEGRO_KCM_AUDIO_FUNC(void, al_stop_sample, (ALLEGRO_SAMPLE_ID *spl_id));
ALLEGRO_KCM_AUDIO_FUNC(void, al_stop_samples, (void));


ALLEGRO_KCM_AUDIO_FUNC(bool, al_register_sample_loader, (const char *ext,
	ALLEGRO_SAMPLE *(*loader)(const char *filename)));
ALLEGRO_KCM_AUDIO_FUNC(bool, al_register_sample_saver, (const char *ext,
	bool (*saver)(const char *filename, ALLEGRO_SAMPLE *spl)));
ALLEGRO_KCM_AUDIO_FUNC(bool, al_register_audio_stream_loader, (const char *ext,
	ALLEGRO_AUDIO_STREAM *(*stream_loader)(const char *filename,
	    size_t buffer_count, unsigned int samples)));

ALLEGRO_KCM_AUDIO_FUNC(bool, al_register_sample_loader_f, (const char *ext,
	ALLEGRO_SAMPLE *(*loader)(ALLEGRO_FILE *fp)));
ALLEGRO_KCM_AUDIO_FUNC(bool, al_register_sample_saver_f, (const char *ext,
	bool (*saver)(ALLEGRO_FILE *fp, ALLEGRO_SAMPLE *spl)));
ALLEGRO_KCM_AUDIO_FUNC(bool, al_register_audio_stream_loader_f, (const char *ext,
	ALLEGRO_AUDIO_STREAM *(*stream_loader)(ALLEGRO_FILE *fp,
	    size_t buffer_count, unsigned int samples)));

ALLEGRO_KCM_AUDIO_FUNC(ALLEGRO_SAMPLE *, al_load_sample, (const char *filename));
ALLEGRO_KCM_AUDIO_FUNC(bool, al_save_sample, (const char *filename,
	ALLEGRO_SAMPLE *spl));
ALLEGRO_KCM_AUDIO_FUNC(ALLEGRO_AUDIO_STREAM *, al_load_audio_stream, (const char *filename,
	size_t buffer_count, unsigned int samples));

ALLEGRO_KCM_AUDIO_FUNC(ALLEGRO_SAMPLE *, al_load_sample_f, (ALLEGRO_FILE* fp, const char *ident));
ALLEGRO_KCM_AUDIO_FUNC(bool, al_save_sample_f, (ALLEGRO_FILE* fp, const char *ident,
	ALLEGRO_SAMPLE *spl));
ALLEGRO_KCM_AUDIO_FUNC(ALLEGRO_AUDIO_STREAM *, al_load_audio_stream_f, (ALLEGRO_FILE* fp, const char *ident,
	size_t buffer_count, unsigned int samples));

ALLEGRO_KCM_AUDIO_FUNC(ALLEGRO_EVENT_SOURCE *, al_get_audio_event_source, (void));


ALLEGRO_KCM_AUDIO_FUNC(ALLEGRO_AUDIO_RECORDER *, al_create_audio_recorder, (size_t fragment_count,
   unsigned int samples, unsigned int freq, AudioDepth depth, ALLEGRO_CHANNEL_CONF chan_conf));
ALLEGRO_KCM_AUDIO_FUNC(bool, al_start_audio_recorder, (ALLEGRO_AUDIO_RECORDER *r));
ALLEGRO_KCM_AUDIO_FUNC(void, al_stop_audio_recorder, (ALLEGRO_AUDIO_RECORDER *r));
ALLEGRO_KCM_AUDIO_FUNC(bool, al_is_audio_recorder_recording, (ALLEGRO_AUDIO_RECORDER *r));
ALLEGRO_KCM_AUDIO_FUNC(ALLEGRO_EVENT_SOURCE *, al_get_audio_recorder_event_source,
   (ALLEGRO_AUDIO_RECORDER *r));
ALLEGRO_KCM_AUDIO_FUNC(ALLEGRO_AUDIO_RECORDER_EVENT *, al_get_audio_recorder_event, (ALLEGRO_EVENT *event));
ALLEGRO_KCM_AUDIO_FUNC(void, al_destroy_audio_recorder, (ALLEGRO_AUDIO_RECORDER *r));
*/
