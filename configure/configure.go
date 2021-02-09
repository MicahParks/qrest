package configure

import (
	"github.com/MicahParks/qrest/backend"
)

// Configure TODO
func Configure() *backend.QuotaManager {
	return backend.NewQuotaManager()
}
