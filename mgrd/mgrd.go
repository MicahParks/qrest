package mgrd

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/mvo5/qrest-skeleton/backend"
)

type Mgrd struct {
	server *http.Server

	errCh  chan error
	quitCh chan struct{}

	backend *backend.QuotaManager
}

// New creates a new Mgrd for the given backend. It will listen on
// the given addr (e.g. ":8080")
func New(addr string, backend *backend.QuotaManager) *Mgrd {
	d := &Mgrd{
		server: &http.Server{
			Addr: addr,
		},
		errCh:   make(chan error, 2),
		backend: backend,
	}
	d.buildRoutes()
	return d
}

// Start starts the daemon in a go-routine.
// Use "Wait()" to wait until it's finished.
func (r *Mgrd) Start() error {
	if r.quitCh != nil {
		return fmt.Errorf("mgrd already started")
	}

	// allow graceful shutdown of server
	r.quitCh = make(chan struct{})
	go func() {
		select {
		case <-r.quitCh:
			ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
			if err := r.server.Shutdown(ctx); err != nil {
				r.errCh <- err
			}
		}
	}()
	// run server, calling Listen here ensures the listen is ready
	// when Start() returns, doing ListenAndServe() in the go routing
	// does not
	ln, err := net.Listen("tcp", r.server.Addr)
	if err != nil {
		return err

	}
	go func() {
		if err := r.server.Serve(ln); err != nil && err != http.ErrServerClosed {
			r.errCh <- err
		}
		close(r.errCh)
	}()

	return nil
}

// Stop asks the daemon to stop gracefully
func (r *Mgrd) Stop() {
	close(r.quitCh)
}

// Wait waits until the daemon stops and returns any errors that happend
// during the daemon lifetime
func (r *Mgrd) Wait() error {
	var errs []error
	for err := range r.errCh {
		errs = append(errs, err)
	}
	switch len(errs) {
	case 0:
		return nil
	case 1:
		return errs[0]
	default:
		return fmt.Errorf("mgrd errored: %v", errs)
	}
}

