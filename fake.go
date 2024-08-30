package main

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
		log.Ctx(ctx).Debug().
			Str("eventName", eventName).
			Any("event.userProperties", event.UserProperties()).
			Any("event.eventProperties", event.EventProperties()).
			Any("opts", opts).
			Msg("Tracking event")
		return nil
	}
	log.Ctx(ctx).Info().
		Str("eventName", eventName).
		Msg("Enqueueing message")
	return nil
}

func (f *fakeTracker) Close() error {
	return nil
}
