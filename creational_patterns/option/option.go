package option

import (
	"time"
)

type Server struct {
	addr     string
	port     int
	timeout  time.Duration
	logLevel string
}

func NewServer(addr string, port int, opts ...Option) *Server {
	s := &Server{
		addr:     addr,
		port:     port,
		timeout:  10 * time.Second,
		logLevel: "info",
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

type Option func(*Server)

func WithTimeout(t time.Duration) Option {
	return func(s *Server) {
		s.timeout = t
	}
}

func WithLogLevel(level string) Option {
	return func(s *Server) {
		s.logLevel = level
	}
}
