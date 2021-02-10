package endpoints

import (
	"fmt"
)

// groupNotFound TODO
func groupNotFound(groupName string) (code int, message string) {
	return 400, fmt.Sprintf("The quota-group \"%s\" was not found.", groupName)
}
