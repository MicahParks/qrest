package configure

import (
	"go.uber.org/zap"

	"github.com/MicahParks/qrest/backend"
)

// Configure TODO
func Configure() (quotaManager *backend.QuotaManager, logger *zap.SugaredLogger, err error) {

	// Create a logger.
	zapLogger, err := zap.NewDevelopment() // TODO Make NewProduction
	if err != nil {
		return nil, nil, err
	}
	logger = zapLogger.Sugar()

	// Create the quota manager.
	return backend.NewQuotaManager(), logger, nil
}
