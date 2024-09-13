package segment

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type fakeTracker struct{}

func NewFakeTracker() Trackable {
	return &fakeTracker{}
}

func (f *fakeTracker) Track(ctx context.Context, eventName string, event Traceable, opts ...TrackOption) error {
	if log.Ctx(ctx).GetLevel() < zerolog.InfoLevel {
		trackOpts := trackOpts{}
		for _, opt := range opts {
			opt(&trackOpts)
		}
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
