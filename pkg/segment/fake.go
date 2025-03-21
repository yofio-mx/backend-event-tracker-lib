package segment

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/yofio-mx/backend-event-tracker-lib/pkg/track"
)

type fakeTracker struct{}

func NewFakeTracker() track.Trackable {
	return &fakeTracker{}
}

func (f *fakeTracker) Track(ctx context.Context, eventName string, event track.Traceable, opts ...track.TrackOption) error {
	if log.Ctx(ctx).GetLevel() < zerolog.InfoLevel {
		trackOpts := track.Apply(opts...)
		log.Ctx(ctx).Debug().
			Str("eventName", eventName).
			Any("event.userProperties", event.UserProperties()).
			Any("event.eventProperties", event.EventProperties()).
			Any("opts", trackOpts).
			Msg("Fake tracking event")
		return nil
	}
	log.Ctx(ctx).Info().
		Str("eventName", eventName).
		Msg("Fake Enqueueing message")
	return nil
}

func (f *fakeTracker) Close() error {
	return nil
}
