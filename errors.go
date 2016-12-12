package foursquarego

import (
	"fmt"
)

// APIError is a foursquare error response
// https://developer.foursquare.com/overview/responses
type APIError struct {
	Meta Meta `json:"meta"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("foursquare: %d %v", e.Meta.Code, e.Meta.ErrorDetail)
}

func relevantError(httpError error, apiError APIError) error {
	if httpError != nil {
		return httpError
	}
	return apiError
}
