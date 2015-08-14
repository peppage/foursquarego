package foursquarego

import (
	"fmt"
	"net/http"
	"net/url"
)

const (
	InvalidAuth       = 401
	ParamError        = 400
	EndPointError     = 404
	NotAuthorized     = 403
	RateLimitExceeded = 403
	Depreciated       = 200
)

type ApiError struct {
	StatusCode int
	Meta       Meta
	Url        *url.URL
}

func newApiError(resp *http.Response, meta Meta) *ApiError {
	return &ApiError{
		StatusCode: resp.StatusCode,
		Meta:       meta,
		Url:        resp.Request.URL,
	}
}

func (aerr ApiError) Error() string {
	return fmt.Sprintf("Get %s returned status %d. Code: %d, %s: %s", aerr.Url, aerr.StatusCode, aerr.Meta.Code, aerr.Meta.ErrorDetail, aerr.Meta.ErrorType)
}
