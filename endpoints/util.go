package endpoints

import (
	"fmt"
)

// groupNotFound creates an HTTP status code and message for when a group is not found.
func groupNotFound(groupName string) (code int, message string) {
	return 400, fmt.Sprintf("The quota-group \"%s\" was not found.", groupName)
}
