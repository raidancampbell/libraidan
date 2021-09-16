// Package rprof contains functions useful in runtime profiling of golang programs
package rprof

import (
	"context"
	"fmt"
	"net/http"
	"runtime/pprof"
)

// LabelMiddleware will add a pprof "URL" label containing the incoming HTTP requests's method and URL Path
func LabelMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		labels := pprof.Labels("URL", fmt.Sprintf("%s %s", r.Method, r.URL.Path))
		pprof.Do(r.Context(), labels, func(ctx context.Context) {
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	})
}
