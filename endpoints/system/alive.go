package system

import (
	"github.com/go-openapi/runtime/middleware"

	"github.com/mvo5/qrest-skeleton/restapi/operations/system"
)

// HandleAlive creates and /api/alive endpoint handler via a closure.
func HandleAlive() system.AliveHandlerFunc {
	return func(params system.AliveParams) middleware.Responder {
		return &system.AliveOK{}
	}
}
