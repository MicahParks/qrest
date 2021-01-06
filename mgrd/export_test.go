package mgrd

import (
	"net/http"
)

func InjectErr(mgrd *Mgrd, err error) {
	mgrd.errCh <- err
}

func (d *Mgrd) BuildRoutes() {
	d.buildRoutes()
}

func (d *Mgrd) Server() *http.Server {
	return d.server
}
