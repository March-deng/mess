package web

import (
	"context"
	"errors"
	"net/http"
)

type mhttp struct {
	server *http.Server
}

func newMhttp(addr string, handler http.Handler) *mhttp {
	return &mhttp{
		server: &http.Server{
			Addr:    addr,
			Handler: handler,
		},
	}
}

func (m *mhttp) registerServer(server *http.Server) {
	m.server = server
	return
}

func (m *mhttp) shutDown() error {
	if m.server != nil {
		return m.server.Shutdown(context.Background())
	}
	return nil
}

func (m *mhttp) start() error {
	if m.server == nil {
		return errors.New("server not registerd yet")
	}
	if err := m.server.ListenAndServe(); err == http.ErrServerClosed {
		return errors.New("http server closed")
	}
	return nil
}

//复用一些没有被明确指明的配置，将原先的http server关闭并开启新的http server
func (m *mhttp) restart() error {
	return nil
}
