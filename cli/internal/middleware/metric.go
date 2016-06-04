package middleware

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"gopkg.in/inconshreveable/log15.v2"

	"github.com/prometheus/client_golang/prometheus"
	"sourcegraph.com/sourcegraph/sourcegraph/pkg/httputil/httpctx"
	"sourcegraph.com/sourcegraph/sourcegraph/pkg/repotrackutil"
	"sourcegraph.com/sourcegraph/sourcegraph/pkg/statsutil"
	"sourcegraph.com/sourcegraph/sourcegraph/pkg/traceutil"
)

var metricLabels = []string{"route", "method", "code", "repo"}
var requestDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Namespace: "src",
	Subsystem: "http",
	Name:      "request_duration_seconds",
	Help:      "The HTTP request latencies in seconds.",
	Buckets:   statsutil.UserLatencyBuckets,
}, metricLabels)
var requestHeartbeat = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: "src",
	Subsystem: "http",
	Name:      "requests_last_timestamp_unixtime",
	Help:      "Last time a request finished for a http endpoint.",
}, metricLabels)

func init() {
	prometheus.MustRegister(requestDuration)
	prometheus.MustRegister(requestHeartbeat)
}

// Metrics captures and exports metrics to prometheus for our HTTP handlers
func Metrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rwIntercept := &ResponseWriterStatusIntercept{ResponseWriter: rw}
		next.ServeHTTP(rwIntercept, r)

		// If we have an error, name is an empty string which
		// indicates to httptrace to use a fallback value
		name, _ := httpctx.RouteNameOrError(r)
		// If the code is zero, the inner Handler never explicitly called
		// WriterHeader. We can assume the response code is 200 in such a case
		code := rwIntercept.Code
		if code == 0 {
			code = 200
		}
		duration := time.Now().Sub(start)
		labels := prometheus.Labels{
			"route":  name,
			"method": strings.ToLower(r.Method),
			"code":   strconv.Itoa(code),
			"repo":   repotrackutil.GetTrackedRepoPath(r.URL.Path),
		}
		requestDuration.With(labels).Observe(duration.Seconds())
		requestHeartbeat.With(labels).Set(float64(time.Now().Unix()))

		log15.Debug("TRACE HTTP", "method", r.Method, "URL", r.URL.String(), "routename", name, "spanID", traceutil.SpanID(r), "code", code, "RemoteAddr", r.RemoteAddr, "UserAgent", r.UserAgent(), "duration", duration)
	})
}

// ResponseWriterStatusIntercept implements the http.ResponseWriter interface
// so we can intercept the status that we can otherwise not access
type ResponseWriterStatusIntercept struct {
	http.ResponseWriter
	Code int
}

// WriteHeader saves the code and then delegates to http.ResponseWriter
func (r *ResponseWriterStatusIntercept) WriteHeader(code int) {
	r.Code = code
	r.ResponseWriter.WriteHeader(code)
}

var _ http.ResponseWriter = (*ResponseWriterStatusIntercept)(nil)
