package jobs

import (
	"time"

	"github.com/sentinel-official/sentinel-go-sdk/libs/geoip"
)

// Default values for GeoIPJob
const (
	DefaultGeoIPJobInterval = 6 * time.Hour // Default interval for running the GeoIP job
	DefaultGeoIPJobName     = "GeoIP"       // Default name for the GeoIP job
)

// GeoIPJob is a job that fetches IP location data at a specified interval.
type GeoIPJob struct {
	geoip.Resolver                             // Embedded geoip.Resolver to resolve IP location
	handler        func(*geoip.Location) error // Handler function to process location data
	interval       time.Duration               // Interval at which the job runs
	name           string                      // Name of the GeoIP job
}

// NewGeoIPJob creates a new instance of GeoIPJob with default settings.
func NewGeoIPJob(handler func(*geoip.Location) error) *GeoIPJob {
	return &GeoIPJob{
		Resolver: geoip.NewDefaultResolver(),
		handler:  handler,
		interval: DefaultGeoIPJobInterval,
		name:     DefaultGeoIPJobName,
	}
}

// Run executes the GeoIP job to fetch IP location data and prints it.
func (j *GeoIPJob) Run() error {
	location, err := j.Resolve("")
	if err != nil {
		return err
	}

	// Call the handler function with the location data
	if j.handler != nil {
		if err := j.handler(location); err != nil {
			return err
		}
	}

	return nil
}

// Name returns the name of the GeoIP job.
func (j *GeoIPJob) Name() string {
	return j.name
}

// Interval returns the interval duration at which the GeoIP job should be executed.
func (j *GeoIPJob) Interval() time.Duration {
	return j.interval
}

// OnError handles errors that occur during the job's execution.
func (j *GeoIPJob) OnError(_ error) bool {
	// Do not stop job on error
	return false
}
