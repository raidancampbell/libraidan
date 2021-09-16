// Package rprof contains functions useful in runtime profiling of golang programs
package rprof

import (
	"fmt"
	"net/http"
	"runtime/pprof"
)

// LabelMiddleware will add a pprof "URL" label containing the incoming HTTP requests's method and URL Path
func LabelMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := pprof.WithLabels(r.Context(), pprof.Labels("URL", fmt.Sprintf("%s %s", r.Method, r.URL.Path)))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
