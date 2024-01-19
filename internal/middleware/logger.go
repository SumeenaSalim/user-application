package middleware

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// LoggingResponseWriter is a custom ResponseWriter that captures the status code
type LoggingResponseWriter struct {
	http.ResponseWriter
	Status int
}

func (lrw *LoggingResponseWriter) WriteHeader(status int) {
	lrw.Status = status
	lrw.ResponseWriter.WriteHeader(status)
}

// LoggingMiddleware is a middleware for logging incoming HTTP requests
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Capture the request body
		var requestBody bytes.Buffer
		teeBody := io.TeeReader(r.Body, &requestBody)
		r.Body = ioutil.NopCloser(teeBody)

		// Create a custom ResponseWriter
		lrw := &LoggingResponseWriter{w, http.StatusOK}

		startTime := time.Now()

		// Call the next handler in the chain
		next.ServeHTTP(lrw, r)

		// Capture the response status code
		statusCode := lrw.Status

		// Capture the response body
		var responseBody bytes.Buffer
		teeBody = io.TeeReader(r.Body, &responseBody)
		r.Body = ioutil.NopCloser(teeBody)

		// Log information about the request and response
		duration := time.Since(startTime)
		headers := make(map[string]string)
		for key, values := range r.Header {
			headers[key] = strings.Join(values, ", ")
		}

		logrus.WithFields(logrus.Fields{
			"method":       r.Method,
			"uri":          r.RequestURI,
			"status":       statusCode,
			"duration":     duration,
			"headers":      headers,
			"query":        r.URL.Query(),
			"userAgent":    r.UserAgent(),
			"remoteAddr":   r.RemoteAddr,
			"requestBody":  requestBody.String(),  // Add request body to log
			"responseBody": responseBody.String(), // Add response body to log
		}).Info("Request processed")
	})
}
