package authenticator

const (
	ServiceName = "authenticator"
	Version     = "1"

	HttpAuthHeader           = "Authorization"
	HttpForwardedHeader      = "X-FORWARDED-FOR"
	HttpInternalUserIDHeader = "I-User-ID"
	HttpUserIDHeader         = "User"
	HttpAuthType             = "Bearer"
)
