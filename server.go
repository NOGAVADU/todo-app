package todo

import (
	"context"
	"net/http"
	"time"
)

type Sever struct {
	httpServer *http.Server
}

func (s *Sever) Start(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Sever) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
