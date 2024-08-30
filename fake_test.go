package main

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func getLoggerCtx(isDebug bool) context.Context {
	w := zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339,
	}
	lv := zerolog.InfoLevel
	if isDebug {
		lv = zerolog.TraceLevel
	}
	l := zerolog.New(w).Level(lv).With().Timestamp().Caller().Stack().Logger()
	ctx := l.WithContext(context.Background())
	return ctx
}

type testEvent struct {
	userProperties  map[string]interface{}
	eventProperties map[string]interface{}
}

func (t testEvent) UserProperties() map[string]interface{} {
	return t.userProperties
}

func (t testEvent) EventProperties() map[string]interface{} {
	return t.eventProperties
}

func getEvent(userProps map[string]interface{}, eventProps map[string]interface{}) Traceable {
	return &testEvent{
		userProperties:  userProps,
		eventProperties: eventProps,
	}
}

func Test_Example_NoOpts(t *testing.T) {
	ctx := getLoggerCtx(false)
	f := NewFakeTracker()
	defer func() {
		_ = f.Close()
	}()

	eventName := "test"
	event := getEvent(map[string]interface{}{
		"key_user": "value_user",
	}, map[string]interface{}{
		"key_event": "value_event",
	})
	assert.NoError(t, f.Track(ctx, eventName, event))
}

func Test_Example_WithOpts(t *testing.T) {
	ctx := getLoggerCtx(true)
	f := NewFakeTracker()
	defer func() {
		_ = f.Close()
	}()

	eventName := "test"
	event := getEvent(map[string]interface{}{
		"key_user": "value_user",
	}, map[string]interface{}{
		"key_event": "value_event",
	})
	err := f.Track(ctx, eventName, event, WithUserID("user_id"), WithAnonymousID("anonymous_id"))
	assert.NoError(t, err)
}
