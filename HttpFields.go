package gotten

import "github.com/Hexilee/gotten/headers"

type (
	Header string
	HeaderAccept string
	HeaderAcceptEncoding string
	HeaderAllow string
	HeaderAuthorization string
	HeaderContentDisposition string
	HeaderContentEncoding string
	HeaderContentLength string
	HeaderContentType string
	HeaderCookie string
	HeaderSetCookie string
	HeaderIfModifiedSince string
	HeaderLastModified string
	HeaderLocation string
	HeaderUpgrade string
	HeaderVary string
	HeaderWWWAuthenticate string
	HeaderXForwardedFor string
	HeaderXForwardedProto string
	HeaderXForwardedProtocol string
	HeaderXForwardedSsl string
	HeaderXUrlScheme string
	HeaderXHTTPMethodOverride string
	HeaderXRealIP string
	HeaderXRequestID string
	HeaderServer string
	HeaderOrigin string

	// Access control
	HeaderAccessControlRequestMethod string
	HeaderAccessControlRequestHeaders string
	HeaderAccessControlAllowOrigin string
	HeaderAccessControlAllowMethods string
	HeaderAccessControlAllowHeaders string
	HeaderAccessControlAllowCredentials string
	HeaderAccessControlExposeHeaders string
	HeaderAccessControlMaxAge string

	// Security
	HeaderStrictTransportSecurity string
	HeaderXContentTypeOptions string
	HeaderXXSSProtection string
	HeaderXFrameOptions string
	HeaderContentSecurityPolicy string
	HeaderXCSRFToken string
)

func (i HeaderAccept) Key() string {
	return headers.HeaderAccept
}
func (i HeaderAcceptEncoding) Key() string {
	return headers.HeaderAcceptEncoding
}
func (i HeaderAllow) Key() string {
	return headers.HeaderAllow
}
func (i HeaderAuthorization) Key() string {
	return headers.HeaderAuthorization
}
func (i HeaderContentDisposition) Key() string {
	return headers.HeaderContentDisposition
}
func (i HeaderContentEncoding) Key() string {
	return headers.HeaderContentEncoding
}
func (i HeaderContentLength) Key() string {
	return headers.HeaderContentLength
}
func (i HeaderContentType) Key() string {
	return headers.HeaderContentType
}
func (i HeaderCookie) Key() string {
	return headers.HeaderCookie
}
func (i HeaderSetCookie) Key() string {
	return headers.HeaderSetCookie
}
func (i HeaderIfModifiedSince) Key() string {
	return headers.HeaderIfModifiedSince
}
func (i HeaderLastModified) Key() string {
	return headers.HeaderLastModified
}
func (i HeaderLocation) Key() string {
	return headers.HeaderLocation
}
func (i HeaderUpgrade) Key() string {
	return headers.HeaderUpgrade
}
func (i HeaderVary) Key() string {
	return headers.HeaderVary
}
func (i HeaderWWWAuthenticate) Key() string {
	return headers.HeaderWWWAuthenticate
}
func (i HeaderXForwardedFor) Key() string {
	return headers.HeaderXForwardedFor
}
func (i HeaderXForwardedProto) Key() string {
	return headers.HeaderXForwardedProto
}
func (i HeaderXForwardedProtocol) Key() string {
	return headers.HeaderXForwardedProtocol
}
func (i HeaderXForwardedSsl) Key() string {
	return headers.HeaderXForwardedSsl
}
func (i HeaderXUrlScheme) Key() string {
	return headers.HeaderXUrlScheme
}
func (i HeaderXHTTPMethodOverride) Key() string {
	return headers.HeaderXHTTPMethodOverride
}
func (i HeaderXRealIP) Key() string {
	return headers.HeaderXRealIP
}
func (i HeaderXRequestID) Key() string {
	return headers.HeaderXRequestID
}
func (i HeaderServer) Key() string {
	return headers.HeaderServer
}
func (i HeaderOrigin) Key() string {
	return headers.HeaderOrigin
}
func (i HeaderAccessControlRequestMethod) Key() string {
	return headers.HeaderAccessControlRequestMethod
}
func (i HeaderAccessControlRequestHeaders) Key() string {
	return headers.HeaderAccessControlRequestHeaders
}
func (i HeaderAccessControlAllowOrigin) Key() string {
	return headers.HeaderAccessControlAllowOrigin
}
func (i HeaderAccessControlAllowMethods) Key() string {
	return headers.HeaderAccessControlAllowMethods
}
func (i HeaderAccessControlAllowHeaders) Key() string {
	return headers.HeaderAccessControlAllowHeaders
}
func (i HeaderAccessControlAllowCredentials) Key() string {
	return headers.HeaderAccessControlAllowCredentials
}
func (i HeaderAccessControlExposeHeaders) Key() string {
	return headers.HeaderAccessControlExposeHeaders
}
func (i HeaderAccessControlMaxAge) Key() string {
	return headers.HeaderAccessControlMaxAge
}
func (i HeaderStrictTransportSecurity) Key() string {
	return headers.HeaderStrictTransportSecurity
}
func (i HeaderXContentTypeOptions) Key() string {
	return headers.HeaderXContentTypeOptions
}
func (i HeaderXXSSProtection) Key() string {
	return headers.HeaderXXSSProtection
}
func (i HeaderXFrameOptions) Key() string {
	return headers.HeaderXFrameOptions
}
func (i HeaderContentSecurityPolicy) Key() string {
	return headers.HeaderContentSecurityPolicy
}
func (i HeaderXCSRFToken) Key() string {
	return headers.HeaderXCSRFToken
}
