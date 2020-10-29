package metrics

import (
	"context"
	"os"
	"time"

	"github.com/tokenized/pkg/logger"
	"go.uber.org/zap"
)

func init() {
	metricsEnabled = !(os.Getenv("METRICS_DISABLED") == "true")
}

var (
	// metricType is a field written into every elapsed time metrics message.
	metricType = zap.String("type", "metric_elapsed")

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

	// get elapsed time in milliseconds
	f := float64(time.Since(start).Nanoseconds()) / float64(time.Millisecond)

	logger.GetLogger(ctx).Info(name, metricType, zap.Float64(fieldElapsed, f))
}
