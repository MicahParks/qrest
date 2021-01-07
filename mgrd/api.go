package mgrd

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (d *Mgrd) buildRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/", d.help)
	d.server.Handler = r
}

func (d *Mgrd) help(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "rest API for quota service")
}

type quotaGroupJSON struct {
	Name      string   `json:"name"`
	MaxMemory uint64   `json:"max-memory"`
	Snaps     []string `json:"snaps,omitempty"`
	Groups    []string `json:"groups,omitempty"`
}
