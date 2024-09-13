package segment

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/analytics-go"
	"time"
)

var (
	ErrAPIKeyNotProvided = fmt.Errorf("API Key or client must be provided")
)

type SegmentTrackerConfig struct {
	// API Key for connecting with Segment
	APIKey string
	// Segment client
	analytics.Client
	// Name of the application to be tracked
	AppName string
	// Version of the application to be tracked
	AppVersion string
	// Build of the application to be tracked
	AppBuild string
	// The flushing interval of the client.
	Interval time.Duration
	// The maximum number of messages to batch before flushing.
	BatchSize int
}

func (config *SegmentTrackerConfig) setDefaults() {
	if config.Interval == 0 {
		config.Interval = analytics.DefaultInterval
	}
	if config.BatchSize == 0 {
		config.BatchSize = analytics.DefaultBatchSize
	}
	if config.AppName == "" {
		config.AppName = "event-tracker-lib"
	}
	if config.AppBuild == "" {
		config.AppBuild = "1"
	}
}

func (config *SegmentTrackerConfig) validate() error {
	if config.APIKey == "" && config.Client == nil {
		return ErrAPIKeyNotProvided
	}
	return nil
}

type segmentTracker struct {
	client analytics.Client
}

func NewSegmentTracker(config SegmentTrackerConfig) (Trackable, error) {
	if config.Client != nil {
		return &segmentTracker{client: config.Client}, nil
	}
	config.setDefaults()
	if err := config.validate(); err != nil {
		log.Error().Err(err).Msg("Error creating segment tracker")
		return nil, err
	}
	client, err := analytics.NewWithConfig(config.APIKey, analytics.Config{
		Interval:  config.Interval,
		BatchSize: config.BatchSize,
		DefaultContext: &analytics.Context{
			App: analytics.AppInfo{
				Name:    config.AppName,
				Version: config.AppVersion,
				Build:   config.AppBuild,
			},
		},
	})
	if err != nil {
		log.Error().Err(err).Msg("Error creating segment client")
		return nil, err
	}
	return &segmentTracker{client: client}, nil
}

func (st *segmentTracker) Track(ctx context.Context, eventName string, event Traceable, opts ...TrackOption) error {
	_logger := log.Ctx(ctx).With().
		Str("EventName", eventName).
		Logger()

	_logger.Debug().
		Msg("Starting to send event information to segment")

	trackOpts := trackOpts{}
	for _, opt := range opts {
		opt(&trackOpts)
	}

	if event.UserProperties() != nil {
		m := analytics.Identify{
			Traits: analytics.Traits(event.UserProperties()),
		}
		if trackOpts.userID != "" {
			m.UserId = trackOpts.userID
		}
		if trackOpts.anonymousID != "" {
			m.AnonymousId = trackOpts.anonymousID
		}
		err := st.client.Enqueue(m)
		if err != nil {
			_logger.Error().Err(err).Msg("Error sending segment identify")
			return err
		}
		_logger.Debug().Msg("Identify message was enqueued successfully to segment")
	}
	if event.EventProperties() != nil {
		m := analytics.Track{
			Event:      eventName,
			Properties: analytics.Properties(event.EventProperties()),
		}
		if trackOpts.userID != "" {
			m.UserId = trackOpts.userID
		}
		if trackOpts.anonymousID != "" {
			m.AnonymousId = trackOpts.anonymousID
		}
		err := st.client.Enqueue(m)
		if err != nil {
			_logger.Error().Err(err).Msg("Error sending segment track")
			return err
		}
		_logger.Debug().Msg("Track message was enqueued successfully to segment")
	}
	return nil
}

func (st *segmentTracker) Close() error {
	return st.client.Close()
}
