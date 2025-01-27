package gateway

import (
	"github.com/gorilla/mux"
	"github.com/prysmaticlabs/prysm/api/gateway/apimiddleware"
)

type Option func(g *Gateway) error

func (g *Gateway) SetRouter(r *mux.Router) *Gateway {
	g.cfg.router = r
	return g
}

func WithPbHandlers(handlers []*PbMux) Option {
	return func(g *Gateway) error {
		g.cfg.pbHandlers = handlers
		return nil
	}
}

func WithMuxHandler(m MuxHandler) Option {
	return func(g *Gateway) error {
		g.cfg.muxHandler = m
		return nil
	}
}

func WithGatewayAddr(addr string) Option {
	return func(g *Gateway) error {
		g.cfg.gatewayAddr = addr
		return nil
	}
}

func WithRemoteAddr(addr string) Option {
	return func(g *Gateway) error {
		g.cfg.remoteAddr = addr
		return nil
	}
}

// WithRouter allows adding a custom mux router to the gateway.
func WithRouter(r *mux.Router) Option {
	return func(g *Gateway) error {
		g.cfg.router = r
		return nil
	}
}

// WithAllowedOrigins allows adding a set of allowed origins to the gateway.
func WithAllowedOrigins(origins []string) Option {
	return func(g *Gateway) error {
		g.cfg.allowedOrigins = origins
		return nil
	}
}

// WithRemoteCert allows adding a custom certificate to the gateway,
func WithRemoteCert(cert string) Option {
	return func(g *Gateway) error {
		g.cfg.remoteCert = cert
		return nil
	}
}

// WithMaxCallRecvMsgSize allows specifying the maximum allowed gRPC message size.
func WithMaxCallRecvMsgSize(size uint64) Option {
	return func(g *Gateway) error {
		g.cfg.maxCallRecvMsgSize = size
		return nil
	}
}

// WithApiMiddleware allows adding an API middleware proxy to the gateway.
func WithApiMiddleware(endpointFactory apimiddleware.EndpointFactory) Option {
	return func(g *Gateway) error {
		g.cfg.apiMiddlewareEndpointFactory = endpointFactory
		return nil
	}
}
