package easclient

import (
	"fmt"

	"gopkg.in/resty.v1"
)

const (
	HeaderKeyOtrisErr = "X-Otris-Eas-Error"
)

// isErrorResponse checks if the given response is an error.
// Returns true and an error object if that is the case.
func isErrorResponse(res *resty.Response) (bool, error) {
	if status := res.StatusCode(); status > 300 {
		return true, fmt.Errorf("unexpected response status %v, expected value < 300", status)
	}
	if hasErrHeader(res) {
		// TODO: unmarshal body to error type?
		return true, fmt.Errorf("eas responded with error header. status %v is considered an error in that case", res.StatusCode())
	}

	return false, nil
}

func hasErrHeader(res *resty.Response) bool {
	return res.Header().Get(HeaderKeyOtrisErr) == "true"
}
