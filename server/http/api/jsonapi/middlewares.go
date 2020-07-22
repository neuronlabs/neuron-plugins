package jsonapi

import (
	"fmt"
	"net/http"

	"github.com/neuronlabs/neuron-plugins/codec/jsonapi"
	"github.com/neuronlabs/neuron-plugins/server/http/httputil"
)

// MidAccept creates a middleware that requires provided accept
func MidAccept(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		parsed := httputil.ParseAcceptHeader(req.Header)
		for _, qv := range parsed {
			if qv.Value == jsonapi.MimeType {
				next.ServeHTTP(rw, req)
				return
			}
		}
		err := httputil.ErrUnsupportedHeader()
		err.Detail = fmt.Sprintf("header Accept doesn't contain '%s' mime type", jsonapi.MimeType)

		rw.WriteHeader(http.StatusNotAcceptable)
		jsonapi.Codec().MarshalErrors(rw, err)
	})
}

// MidAccept creates a middleware that requires provided accept
func MidContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		ct := req.Header.Get("Content-Type")
		if ct == jsonapi.MimeType {
			next.ServeHTTP(rw, req)
			return
		}
		rw.WriteHeader(http.StatusUnsupportedMediaType)
	})
}
