package track

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

func WithEmail(email string) TrackOption {
	return func(o *trackOpts) {
		o.email = email
	}
}

type trackOpts struct {
	anonymousID string
	userID      string
	email       string
}

func (o *trackOpts) GetAnonymousID() string {
	return o.anonymousID
}

func (o *trackOpts) GetUserID() string {
	return o.userID
}

func (o *trackOpts) GetEmail() string {
	return o.email
}

func Apply(options ...TrackOption) trackOpts {
	opts := trackOpts{}
	for _, option := range options {
		option(&opts)
	}
	return opts
}
