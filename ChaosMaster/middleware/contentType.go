package middleware

import (
	"mallekoppie/ChaosGenerator/ChaosMaster/util/logger"
	"net/http"
	"strings"
)

func AllowedContentType(inner http.Handler, contentTypeConfig string) http.Handler {

	contentType := strings.ToLower(contentTypeConfig)
	enabled := len(contentTypeConfig) > 0

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if enabled {
			result := r.Header.Get("Content-Type")

			if strings.ToLower(result) == contentType {
				inner.ServeHTTP(w, r)
			} else {
				logger.Error("Media type not allowed: ", result)
				w.WriteHeader(http.StatusUnsupportedMediaType)
			}
		} else {
			inner.ServeHTTP(w, r)
		}
	})
}
