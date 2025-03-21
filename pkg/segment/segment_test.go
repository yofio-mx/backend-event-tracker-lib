package segment

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yofio-mx/backend-event-tracker-lib/pkg/track"
)

func Test_Example_UserID(t *testing.T) {
	ctx := getLoggerCtx(true)
	f, err := NewSegmentTracker(SegmentTrackerConfig{
		APIKey: "test", // Put a valid API Key
	})
	assert.NoError(t, err)
	defer func() {
		_ = f.Close()
	}()

	eventName := "test"
	event := getEvent(map[string]interface{}{
		"key_user": "value_user",
	}, map[string]interface{}{
		"key_event": "value_event",
	})
	assert.NoError(t, f.Track(ctx, eventName, event, track.WithUserID("123")))
}

func Test_Example_AnonID(t *testing.T) {
	ctx := getLoggerCtx(true)
	f, err := NewSegmentTracker(SegmentTrackerConfig{
		APIKey: "test", // Put a valid API Key
	})
	assert.NoError(t, err)
	defer func() {
		_ = f.Close()
	}()

	eventName := "test"
	event := getEvent(map[string]interface{}{
		"key_user": "value_user",
	}, map[string]interface{}{
		"key_event": "value_event",
	})
	assert.NoError(t, f.Track(ctx, eventName, event, track.WithAnonymousID("456")))
}

func Test_Example_UserAndAnonID(t *testing.T) {
	ctx := getLoggerCtx(true)
	f, err := NewSegmentTracker(SegmentTrackerConfig{
		APIKey: "test", // Put a valid API Key
	})
	assert.NoError(t, err)
	defer func() {
		_ = f.Close()
	}()

	eventName := "test"
	event := getEvent(map[string]interface{}{
		"key_user": "value_user",
	}, map[string]interface{}{
		"key_event": "value_event",
	})
	assert.NoError(t, f.Track(ctx, eventName, event, track.WithUserID("123"), track.WithAnonymousID("123")))
}

func Test_Example_Email(t *testing.T) {
	ctx := getLoggerCtx(true)
	f, err := NewSegmentTracker(SegmentTrackerConfig{
		APIKey: "test",
	})
	assert.NoError(t, err)
	defer func() {
		_ = f.Close()
	}()

	eventName := "test"
	event := getEvent(map[string]interface{}{
		"key_user": "value_user",
	}, map[string]interface{}{
		"key_event": "value_event",
	})
	assert.NoError(t, f.Track(ctx, eventName, event, track.WithEmail("example@yofio.com")))
}
