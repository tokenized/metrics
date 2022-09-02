package metrics

import (
	"context"
	"os"
	"time"

	"github.com/tokenized/logger"
)

func init() {
	metricsEnabled = !(os.Getenv("METRICS_DISABLED") == "true")
}

var (
	// metricType is a field written into every elapsed time metrics message.
	metricType = logger.String("type", "metric_elapsed")

	metricsEnabled = true
)

const (
	// fieldElapsed is the name of the "elapsed" field written into metrics messages.
	fieldElapsed = "elapsed"
)

// Elapsed is intended to be used to log general timing metrics for a function.
//
// Example
//
//     defer metrics.Elapsed(ctx, time.Now(), "your.func.Name")
func Elapsed(ctx context.Context, start time.Time, name string) {
	if !metricsEnabled {
		return
	}

	// nanoseconds(billion) 1e9, milliseconds(thousand) 1e3, so divide nanoseconds by 1e6 for
	// milliseconds.
	elapsed := float64(time.Since(start).Nanoseconds()) / 1e6

	logger.InfoWithFields(ctx, []logger.Field{
		metricType,
		logger.Formatter(fieldElapsed, "%06f", elapsed), // use %06f so it is fixed width
	}, name)
}

func Disable() {
	metricsEnabled = false
}

func Enable() {
	metricsEnabled = true
}
