# backend-event-tracker-lib

Library for sending event to analytics services used at YoFio.

Currently, it supports:

- Fake service: Only prints the event to the console
- Segment: Sends the event to Segment

## Usage

```go
package main

import (
	"context"
	segment "github.com/yofio-mx/backend-event-tracker-lib"
)

func main() {
	// Create a new event tracker
	tracker, err := segment.NewSegmentTracker(segment.SegmentTrackerConfig{
		APIKey:     "test", // Put a valid API Key
		AppName:    "abc_service",
		AppVersion: "1.0.0",
		AppBuild:   "1",
	})
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = tracker.Close()
	}()

	err = tracker.Track(context.Background(),
		"eventName",
		event, // Define your event that implements the Traceable interface
		WithUserID("123"),
		WithAnonymousID("123"),
	)
	if err != nil {
		panic(err)
	}
}

```

You must send at least the `WithUserID` or `WithAnonymousID` to identify the user.

See the test files for more examples.