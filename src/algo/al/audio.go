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
import "unsafe"

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

type AudioStream struct {
	handle *C.ALLEGRO_AUDIO_STREAM
}

type Mixer struct {
	handle *C.ALLEGRO_MIXER
}

type Voice struct {
	handle *C.ALLEGRO_VOICE
}

type AudioRecorder struct {
	handle *C.ALLEGRO_AUDIO_RECORDER
}

func (self *Sample) toC() *C.ALLEGRO_SAMPLE {
	return (*C.ALLEGRO_SAMPLE)(self.handle)
}

func (self *Sample) Destroy() {
	if self.handle != nil {
		C.al_destroy_sample(self.toC())
	}
	self.handle = nil
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

func createSample(data []byte, samples uint, freq uint, depth AudioDepth, chan_conf ChannelConf) *C.ALLEGRO_SAMPLE {
	// don't let allegro free the data, it's owned by Go.
	// XXX: copy data here in stead of using the go data???
	return C.al_create_sample(unsafe.Pointer(&data[0]), C.uint(samples),
		C.uint(freq), C.ALLEGRO_AUDIO_DEPTH(depth),
		C.ALLEGRO_CHANNEL_CONF(chan_conf), b2cb(false))
}

// Creates a Sample with a Destroy finalizer set. BEWARE! data must be big enough fort he params
// or this will crash spectacularly. Also, make sure that data doesn't get collected by the GC, or 
// trouble will arise. 
func CreateSample(data []byte, samples uint, freq uint, depth AudioDepth, chan_conf ChannelConf) *Sample {
	return wrapSample(createSample(data, samples, freq, depth, chan_conf))
}

// Returns the frequency of the sample
func (self *Sample) Frequency() uint {
	return uint(C.al_get_sample_frequency(self.handle))
}

// Returns the length of the sample
func (self *Sample) Length() uint {
	return uint(C.al_get_sample_length(self.handle))
}

// Returns the depth of the sample 
func (self *Sample) Depth() uint {
	return uint(C.al_get_sample_depth(self.handle))
}

// returns the amount of channels the sample has
func (self *Sample) Channels() uint {
	return uint(C.al_get_sample_channels(self.handle))
}

// Returns the raw data pointer of the sample data. 
func (self *Sample) DataRaw() unsafe.Pointer {
	return (C.al_get_sample_data(self.handle))
}

// Returns the frequency of the sample instance
func (self *SampleInstance) Frequency() uint {
	return uint(C.al_get_sample_instance_frequency(self.handle))
}

// Returns the length of the sample instance
func (self *SampleInstance) Length() uint {
	return uint(C.al_get_sample_instance_length(self.handle))
}

// Returns the position of the sample instance
func (self *SampleInstance) Position() uint {
	return uint(C.al_get_sample_instance_position(self.handle))
}

// Returns the of the sample instance
func (self *SampleInstance) Speed() float32 {
	return float32(C.al_get_sample_instance_speed(self.handle))
}

// Returns the of the sample instance
func (self *SampleInstance) Gain() float32 {
	return float32(C.al_get_sample_instance_gain(self.handle))
}

// Returns the of the sample instance
func (self *SampleInstance) Pan() float32 {
	return float32(C.al_get_sample_instance_pan(self.handle))
}

// Returns the of the sample instance
func (self *SampleInstance) Time() float32 {
	return float32(C.al_get_sample_instance_time(self.handle))
}

// Returns the depth of the sample instance
func (self *SampleInstance) Depth() AudioDepth {
	return AudioDepth(C.al_get_sample_instance_depth(self.handle))
}

// Returns the channel configuration of the sample instance
func (self *SampleInstance) Channels() ChannelConf {
	return ChannelConf(C.al_get_sample_instance_channels(self.handle))
}

// Returns the play mode of the sample instance
func (self *SampleInstance) Playmode() PlayMode {
	return PlayMode(C.al_get_sample_instance_playmode(self.handle))
}

// Returns wheter or not the sample instance is playing
func (self *SampleInstance) Playing() bool {
	return cb2b(C.al_get_sample_instance_playing(self.handle))
}

// Returns wheter or not the sample instance is attached
func (self *SampleInstance) Attached() bool {
	return cb2b(C.al_get_sample_instance_attached(self.handle))
}

// Sets the position of the sample instance.
func (self *SampleInstance) SetPosition(val uint) bool {
	return cb2b(C.al_set_sample_instance_position(self.handle, C.uint(val)))
}

// Sets the length of the sample instance.
func (self *SampleInstance) SetLength(val uint) bool {
	return cb2b(C.al_set_sample_instance_length(self.handle, C.uint(val)))
}

// Sets the speed of the sample instance.
func (self *SampleInstance) SetSpeed(val float32) bool {
	return cb2b(C.al_set_sample_instance_speed(self.handle, C.float(val)))
}

// Sets the gain of the sample instance.
func (self *SampleInstance) SetGain(val float32) bool {
	return cb2b(C.al_set_sample_instance_gain(self.handle, C.float(val)))
}

// Sets the pan of the sample instance.
func (self *SampleInstance) SetPan(val float32) bool {
	return cb2b(C.al_set_sample_instance_pan(self.handle, C.float(val)))
}

// Sets the play mode of the sample instance.
func (self *SampleInstance) SetPlaymode(val PlayMode) bool {
	return cb2b(C.al_set_sample_instance_playmode(self.handle, C.ALLEGRO_PLAYMODE(val)))
}

// Sets the play status of the sample instance.
func (self *SampleInstance) SetPlaying(val bool) bool {
	return cb2b(C.al_set_sample_instance_playing(self.handle, b2cb(val)))
}

// Detaches the sample instance from it's player
func (self *SampleInstance) Detach() bool {
	return cb2b(C.al_detach_sample_instance(self.handle))
}

// Sets the sample data to use for the sample instance
func (self *SampleInstance) SetSample(val *Sample) bool {
	return cb2b(C.al_set_sample(self.handle, val.handle))
}

// Gets the RAW sample data that was linked to the sample instance
func (self *SampleInstance) SampleRaw() *Sample {
	return wrapSampleRaw(C.al_get_sample(self.handle))
}

// Plays the sample instance 
func (self *SampleInstance) Play() bool {
	return cb2b(C.al_play_sample_instance(self.handle))
}

// Stops the sample instannce playback
func (self *SampleInstance) Stop() bool {
	return cb2b(C.al_stop_sample_instance(self.handle))
}

func (self *AudioStream) toC() *C.ALLEGRO_AUDIO_STREAM {
	return (*C.ALLEGRO_AUDIO_STREAM)(self.handle)
}

// Destroys the audio stream.
func (self *AudioStream) Destroy() {
	if self.handle != nil {
		C.al_destroy_audio_stream(self.toC())
	}
	self.handle = nil
}

func wrapAudioStreamRaw(data *C.ALLEGRO_AUDIO_STREAM) *AudioStream {
	if data == nil {
		return nil
	}
	return &AudioStream{data}
}

func wrapAudioStream(data *C.ALLEGRO_AUDIO_STREAM) *AudioStream {
	self := wrapAudioStreamRaw(data)
	if self != nil {
		runtime.SetFinalizer(self, func(me *AudioStream) { me.Destroy() })
	}
	return self
}

// Creates an audio stream, with finalizer installed.
func CreateAudioStream(bufc, samples, freq uint, depth AudioDepth, chan_conf ChannelConf) *AudioStream {
	return wrapAudioStream(C.al_create_audio_stream(C.size_t(bufc), C.uint(samples),
		C.uint(freq), C.ALLEGRO_AUDIO_DEPTH(depth), C.ALLEGRO_CHANNEL_CONF(chan_conf)))
}

// Creates an audio stream, with NO finalizer installed.
func CreateAudioStreamRaw(bufc, samples, freq uint, depth AudioDepth, chan_conf ChannelConf) *AudioStream {
	return wrapAudioStreamRaw(C.al_create_audio_stream(C.size_t(bufc), C.uint(samples),
		C.uint(freq), C.ALLEGRO_AUDIO_DEPTH(depth), C.ALLEGRO_CHANNEL_CONF(chan_conf)))
}

// Drains all data from an audio stream
func (self *AudioStream) Drain() {
	C.al_drain_audio_stream(self.handle)
}

// Returns the frequency of the audio stream
func (self *AudioStream) Frequency() uint {
	return uint(C.al_get_audio_stream_frequency(self.handle))
}

// Returns the length of the audio stream
func (self *AudioStream) Length() uint {
	return uint(C.al_get_audio_stream_length(self.handle))
}

// Returns the speed of the audio stream
func (self *AudioStream) Speed() float32 {
	return float32(C.al_get_audio_stream_speed(self.handle))
}

// Returns the amount of fragments of the audio stream
func (self *AudioStream) Fragments() float32 {
	return float32(C.al_get_audio_stream_fragments(self.handle))
}

// Returns the amount of available fragments of the audio stream
func (self *AudioStream) AvailableFragments() float32 {
	return float32(C.al_get_available_audio_stream_fragments(self.handle))
}

// Returns the gain of the audio stream
func (self *AudioStream) Gain() float32 {
	return float32(C.al_get_audio_stream_gain(self.handle))
}

// Returns the pan of the audio stream
func (self *AudioStream) Pan() float32 {
	return float32(C.al_get_audio_stream_pan(self.handle))
}

// Returns the depth of the audio stream
func (self *AudioStream) Depth() AudioDepth {
	return AudioDepth(C.al_get_audio_stream_depth(self.handle))
}

// Returns the channel configuration of the audio stream
func (self *AudioStream) Channels() ChannelConf {
	return ChannelConf(C.al_get_audio_stream_channels(self.handle))
}

// Returns the play mode of the audio stream
func (self *AudioStream) Playmode() PlayMode {
	return PlayMode(C.al_get_audio_stream_playmode(self.handle))
}

// Returns wheter or not the audio stream is playing
func (self *AudioStream) Playing() bool {
	return cb2b(C.al_get_audio_stream_playing(self.handle))
}

// Returns wheter or not the audio stream is attached
func (self *AudioStream) Attached() bool {
	return cb2b(C.al_get_audio_stream_attached(self.handle))
}

// Returns an unsafe pointer to the audio stream's fragment
func (self *AudioStream) Fragment() unsafe.Pointer {
	return C.al_get_audio_stream_fragment(self.handle)
}

// Sets the speed of the audio stream.
func (self *AudioStream) SetSpeed(val float32) bool {
	return cb2b(C.al_set_audio_stream_speed(self.handle, C.float(val)))
}

// Sets the gain of the audio stream.
func (self *AudioStream) SetGain(val float32) bool {
	return cb2b(C.al_set_audio_stream_gain(self.handle, C.float(val)))
}

// Sets the pan of the audio stream.
func (self *AudioStream) SetPan(val float32) bool {
	return cb2b(C.al_set_audio_stream_pan(self.handle, C.float(val)))
}

// Sets the play mode of the audio stream.
func (self *AudioStream) SetPlaymode(val PlayMode) bool {
	return cb2b(C.al_set_audio_stream_playmode(self.handle, C.ALLEGRO_PLAYMODE(val)))
}

// Sets the play status of the audio stream.
func (self *AudioStream) SetPlaying(val bool) bool {
	return cb2b(C.al_set_audio_stream_playing(self.handle, b2cb(val)))
}

// Detaches the audio stream from it's player
func (self *AudioStream) Detach() bool {
	return cb2b(C.al_detach_sample_instance(self.handle))
}

// Sets an unsafe pointer as the the audio stream's fragment
func (self *AudioStream) SetFragment(ptr unsafe.Pointer) bool {
	return cb2b(C.al_set_audio_stream_fragment(self.handle, ptr))
}

// Rewinds the audio stream
func (self *AudioStream) Rewind() bool {
	return cb2b(C.al_rewind_audio_stream(self.handle))
}

// Seeks to a position in the audio stream, expressed in seconds.
func (self *AudioStream) SeekSeconds(secs float64) bool {
	return cb2b(C.al_seek_audio_stream_secs(self.handle, C.double(secs)))
}

// Gets the position in the audio stream, expressed in seconds.
func (self *AudioStream) PositionSeconds() (secs float64) {
	return float64(C.al_get_audio_stream_position_secs(self.handle))
}

// Gets the length of the audio stream, expressed in seconds.
func (self *AudioStream) LengthSeconds() (secs float64) {
	return float64(C.al_get_audio_stream_length_secs(self.handle))
}

// Sets up a loop in the audio stream, expressed in seconds.
func (self *AudioStream) LoopSeconds(start, end float64) bool {
	return cb2b(C.al_set_audio_stream_loop_secs(self.handle, C.double(start), C.double(end)))
}

// Returns the event source to use to listen to events on the audio stream.
func (self *AudioStream) Eventsource() *EventSource {
	return wrapEventSourceRaw(C.al_get_audio_stream_event_source(self.handle))
}

// Converts a mixer to it's underlying C pointer
func (self *Mixer) toC() *C.ALLEGRO_MIXER {
	return (*C.ALLEGRO_MIXER)(self.handle)
}

// Destroys the mixer.
func (self *Mixer) Destroy() {
	if self.handle != nil {
		C.al_destroy_mixer(self.toC())
	}
	self.handle = nil
}

// Wraps a C mixer into a go mixer
func wrapMixerRaw(data *C.ALLEGRO_MIXER) *Mixer {
	if data == nil {
		return nil
	}
	return &Mixer{data}
}

// Wraps a C mixer into a go mixer and sets up a finalizer that calls Destroy()
func wrapMixer(data *C.ALLEGRO_MIXER) *Mixer {
	self := wrapMixerRaw(data)
	if self != nil {
		runtime.SetFinalizer(self, func(me *Mixer) { me.Destroy() })
	}
	return self
}

func createMixer(freq uint, depth AudioDepth, chan_conf ChannelConf) *C.ALLEGRO_MIXER {
	return C.al_create_mixer(C.uint(freq), C.ALLEGRO_AUDIO_DEPTH(depth), C.ALLEGRO_CHANNEL_CONF(chan_conf))
}

// Creates a new mixer with no finaliser attached to it.
func CreateMixerRaw(freq uint, depth AudioDepth, chan_conf ChannelConf) *Mixer {
	return wrapMixerRaw(createMixer(freq, depth, chan_conf))
}

// Creates a new mixer with a finaliser that will call Destroy attached to it.
func CreateMixer(freq uint, depth AudioDepth, chan_conf ChannelConf) *Mixer {
	return wrapMixer(createMixer(freq, depth, chan_conf))
}

// Attaches a sample instance to the given mixer. 
func (mixer *Mixer) AttachSampleInstance(stream *SampleInstance) bool {
	return cb2b(C.al_attach_sample_instance_to_mixer(stream.handle, mixer.handle))
}

// Attaches the sample instance to the given mixer. 
func (stream *SampleInstance) Attach(mixer *Mixer) bool {
	return cb2b(C.al_attach_sample_instance_to_mixer(stream.handle, mixer.handle))
}

// Attaches an audio stream  to the given mixer. 
func (mixer *Mixer) AttachAudioStream(stream *AudioStream) bool {
	return cb2b(C.al_attach_audio_stream_to_mixer(stream.handle, mixer.handle))
}

// Attaches the audio stream to the given mixer. 
func (stream *AudioStream) Attach(mixer *Mixer) bool {
	return cb2b(C.al_attach_audio_stream_to_mixer(stream.handle, mixer.handle))
}

// Attaches a mixer to the given mixer. 
func (mixer *Mixer) AttachMixer(stream *Mixer) bool {
	return cb2b(C.al_attach_mixer_to_mixer(stream.handle, mixer.handle))
}

// Attaches the given mixer to the latter mixer. 
func (stream *Mixer) Attach(mixer *Mixer) bool {
	return cb2b(C.al_attach_mixer_to_mixer(stream.handle, mixer.handle))
}

/*
TODO:
ALLEGRO_KCM_AUDIO_FUNC(bool, al_set_mixer_postprocess_callback, (
      ALLEGRO_MIXER *mixer,
      void (*cb)(void *buf, unsigned int samples, void *data),
      void *data));
*/

// Returns the frequency of the mixer
func (self *Mixer) Frequency() uint {
	return uint(C.al_get_mixer_frequency(self.handle))
}

// Returns the gain of the mixer
func (self *Mixer) Gain() float32 {
	return float32(C.al_get_mixer_gain(self.handle))
}

// Returns the depth of the mixer
func (self *Mixer) Depth() AudioDepth {
	return AudioDepth(C.al_get_mixer_depth(self.handle))
}

// Returns the channel configuration of the mixer
func (self *Mixer) Channels() ChannelConf {
	return ChannelConf(C.al_get_mixer_channels(self.handle))
}

// Returns the quality of the mixer
func (self *Mixer) Quality() MixerQuality {
	return MixerQuality(C.al_get_mixer_quality(self.handle))
}

// Returns wheter or not the mixer is playing
func (self *Mixer) Playing() bool {
	return cb2b(C.al_get_mixer_playing(self.handle))
}

// Returns wheter or not the mixer is attached
func (self *Mixer) Attached() bool {
	return cb2b(C.al_get_mixer_attached(self.handle))
}

// Sets the frequency of the mixer.
func (self *Mixer) SetFrequency(val uint) bool {
	return cb2b(C.al_set_mixer_frequency(self.handle, C.uint(val)))
}

// Sets the quality of the mixer.
func (self *Mixer) SetQuality(val MixerQuality) bool {
	return cb2b(C.al_set_mixer_quality(self.handle, C.ALLEGRO_MIXER_QUALITY(val)))
}

// Sets the gain of the mixer.
func (self *Mixer) SetGain(val float32) bool {
	return cb2b(C.al_set_mixer_gain(self.handle, C.float(val)))
}

// Sets the play status of the mixer.
func (self *Mixer) SetPlaying(val bool) bool {
	return cb2b(C.al_set_mixer_playing(self.handle, b2cb(val)))
}

// Detaches the mixer from it's player
func (self *Mixer) Detach() bool {
	return cb2b(C.al_detach_mixer(self.handle))
}

// Converts a voice to it's underlying C pointer
func (self *Voice) toC() *C.ALLEGRO_VOICE {
	return (*C.ALLEGRO_VOICE)(self.handle)
}

// Destroys the voice.
func (self *Voice) Destroy() {
	if self.handle != nil {
		C.al_destroy_voice(self.toC())
	}
	self.handle = nil
}

// Wraps a C voice into a go mixer
func wrapVoiceRaw(data *C.ALLEGRO_VOICE) *Voice {
	if data == nil {
		return nil
	}
	return &Voice{data}
}

// Wraps a C voice into a go mixer and sets up a finalizer that calls Destroy()
func wrapVoice(data *C.ALLEGRO_VOICE) *Voice {
	self := wrapVoiceRaw(data)
	if self != nil {
		runtime.SetFinalizer(self, func(me *Voice) { me.Destroy() })
	}
	return self
}

// creates a C voice 
func createVoice(freq uint, depth AudioDepth, chan_conf ChannelConf) *C.ALLEGRO_VOICE {
	return C.al_create_voice(C.uint(freq), C.ALLEGRO_AUDIO_DEPTH(depth), C.ALLEGRO_CHANNEL_CONF(chan_conf))
}

// Creates a voice
func CreateVoiceRaw(freq uint, depth AudioDepth, chan_conf ChannelConf) *Voice {
	return wrapVoiceRaw(createVoice(freq, depth, chan_conf))
}

// Creates a voice and setsup a finalizer on it that calls Destroy
func CreateVoice(freq uint, depth AudioDepth, chan_conf ChannelConf) *Voice {
	return wrapVoice(createVoice(freq, depth, chan_conf))
}

// Attaches a sample instance to the given voice. 
func (voice *Voice) AttachSampleInstance(stream *SampleInstance) bool {
	return cb2b(C.al_attach_sample_instance_to_voice(stream.handle, voice.handle))
}

// Attaches the sample instance to the given voice. 
func (stream *SampleInstance) AttachToVoice(voice *Voice) bool {
	return cb2b(C.al_attach_sample_instance_to_voice(stream.handle, voice.handle))
}

// Attaches an audio stream  to the given voice. 
func (voice *Voice) AttachAudioStream(stream *AudioStream) bool {
	return cb2b(C.al_attach_audio_stream_to_voice(stream.handle, voice.handle))
}

// Attaches the audio stream to the given voice. 
func (stream *AudioStream) AttachToVoice(voice *Voice) bool {
	return cb2b(C.al_attach_audio_stream_to_voice(stream.handle, voice.handle))
}

// Attaches the given mixer to the voice.
func (mixer *Mixer) AttachToVoice(voice *Voice) bool {
	return cb2b(C.al_attach_mixer_to_voice(mixer.handle, voice.handle))
}

// Attaches the given voice to the mixer.
func (voice *Voice) AttachMixer(mixer *Mixer) bool {
	return cb2b(C.al_attach_mixer_to_voice(mixer.handle, voice.handle))
}

// Detaches the voice.
func (voice *Voice) Detach() {
	C.al_detach_voice(voice.handle)
}

/*

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
      float32 gain, float32 pan, float32 speed, ALLEGRO_PLAYMODE loop, ALLEGRO_SAMPLE_ID *ret_id));
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