package segment

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
	assert.NoError(t, f.Track(ctx, eventName, event, WithUserID("123")))
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
	assert.NoError(t, f.Track(ctx, eventName, event, WithAnonymousID("456")))
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
	assert.NoError(t, f.Track(ctx, eventName, event, WithUserID("123"), WithAnonymousID("123")))
}
