package reqdata

import (
	"context"
	"net"
	"net/http"
)

type contextKey string

var (
	ipAddress = contextKey("ipAddress")
)

// ContextFromRequest applies the request data into the context to be later retrieved throughout the application.
func ContextFromRequest(req *http.Request) context.Context {
	ctx := req.Context()

	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		ctx = context.WithValue(ctx, ipAddress, ip)
	}

	return ctx
}

// IPAddressFromContext returns the IP address from an HTTP request if one exists.
func IPAddressFromContext(ctx context.Context) string {
	ipAddress, ok := ctx.Value(ipAddress).(string)
	if !ok {
		return ""
	}

	return ipAddress
}
