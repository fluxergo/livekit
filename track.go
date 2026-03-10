package livekit

import (
	"time"

	lksdk "github.com/livekit/server-sdk-go/v2"
	"github.com/pion/webrtc/v4/pkg/media"
)

type trackWriter struct {
	track          *lksdk.LocalTrack
	sampleDuration time.Duration
}

func (t *trackWriter) Write(p []byte) (int, error) {
	err := t.track.WriteSample(media.Sample{
		Data:     p,
		Duration: t.sampleDuration,
	}, nil)

	return len(p), err
}

func (t *trackWriter) Close() error {
	return t.track.Close()
}
