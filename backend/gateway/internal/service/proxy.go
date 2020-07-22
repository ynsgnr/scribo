package service

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/ynsgnr/scribo/backend/authenticator/authenticator"
	"github.com/ynsgnr/scribo/backend/common/logger"
)

//Remove filtered data when responding back
type FilterTransport struct{}

func (t FilterTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	resp, err := http.DefaultTransport.RoundTrip(r)
	if err != nil {
		return nil, err
	}
	resp.Header.Del(authenticator.HttpInternalUserIDHeader)
	return resp, nil
}

func (s *service) handleProxy() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			internalUserID string
			externalUserID string
			ok             bool
		)
		internalUserID, ok = r.Context().Value(authenticator.HttpInternalUserIDHeader).(string)
		if !ok {
			internalUserID = ""
		}
		externalUserID, ok = r.Context().Value(authenticator.HttpUserIDHeader).(string)
		if !ok {
			externalUserID = ""
		}
		// parse the url
		urlPath := strings.TrimLeft(r.URL.Path, "/")
		segments := strings.Split(urlPath, "/")
		if len(segments) < 1 {
			logger.Printf(logger.Warning, "handleProxy: service not found for url: %s", r.URL.String())
			w.WriteHeader(http.StatusForbidden)
			return
		}
		urlPath = strings.Replace(r.URL.Path, externalUserID, internalUserID, 1)
		proxyUrl := &url.URL{
			Scheme: "http",
			Host:   segments[0],
		}
		proxy := httputil.ReverseProxy{Director: func(r *http.Request) {
			r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
			r.Host = proxyUrl.Host
			r.URL.Host = proxyUrl.Host
			r.URL.Scheme = proxyUrl.Scheme
		},
			Transport: s.filter,
		}
		logger.Printf(logger.Trace, "customer %s requested %s", internalUserID, urlPath)
		// Note that ServeHttp is non blocking and uses a go routine under the hood
		// TODO add timeout here
		proxy.ServeHTTP(w, r)
	})
}
