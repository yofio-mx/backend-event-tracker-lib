package segment

import (
	"context"
	"io"
)

type Traceable interface {
	UserProperties() map[string]interface{}
	EventProperties() map[string]interface{}
}

type Trackable interface {
	io.Closer
	Track(ctx context.Context, eventName string, t Traceable, opts ...TrackOption) error
}

type TrackOption func(*trackOpts)

func WithUserID(userID string) TrackOption {
	return func(o *trackOpts) {
		o.userID = userID
	}
}

func WithAnonymousID(anonymousID string) TrackOption {
	return func(o *trackOpts) {
		o.anonymousID = anonymousID
	}
}

type trackOpts struct {
	anonymousID string
	userID      string
}
