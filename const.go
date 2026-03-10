package livekit

import (
	"github.com/fluxergo/fluxergo/voice"
	"github.com/livekit/protocol/livekit"
)

func audioSourceToLiveKit(s voice.AudioSource) livekit.TrackSource {
	switch s {
	case voice.AudioSourceMicrophone:
		return livekit.TrackSource_MICROPHONE
	case voice.AudioSourceScreenShare:
		return livekit.TrackSource_SCREEN_SHARE_AUDIO
	default:
		return livekit.TrackSource_UNKNOWN
	}
}

func videoSourceToLiveKit(s voice.VideoSource) livekit.TrackSource {
	switch s {
	case voice.VideoSourceCamera:
		return livekit.TrackSource_CAMERA
	case voice.VideoSourceScreenShare:
		return livekit.TrackSource_SCREEN_SHARE
	default:
		return livekit.TrackSource_UNKNOWN
	}
}
