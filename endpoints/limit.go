package endpoints

import (
	"github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"

	"github.com/MicahParks/qrest/backend"
	"github.com/MicahParks/qrest/restapi/operations/api"
)

// HandleGroupLimitRead TODO
func HandleGroupLimitRead(logger *zap.SugaredLogger, quotaManager *backend.QuotaManager) api.GroupLimitReadHandlerFunc {
	return func(params api.GroupLimitReadParams) middleware.Responder {

	}
}
