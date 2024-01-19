package easclient

import (
	"encoding/json"
	"fmt"

	"gopkg.in/resty.v1"
)

const (
	HeaderKeyOtrisErr = "X-Otris-Eas-Error"
)

// errorResponse is the response body of an eas response where
// the HeaderKeyOtrisErr is set. It contains details about why something failed.
type errorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type ResponseErr struct {
	msg string

	Body []byte

	// Response is the original error response from the EAS. This may contain
	// zero values if the EAS did not return a valid error response.
	Response errorResponse
}

func NewResponseErr(msg string, body []byte) *ResponseErr {
	err := &ResponseErr{msg: msg, Body: body}

	// try to parse EAS response. If it fails, we ignore it.
	_ = json.Unmarshal(body, &err.Response)

	return err
}

func (r *ResponseErr) Error() string {
	return r.msg
}

// isErrorResponse checks if the given response is an error.
func isErrorResponse(res *resty.Response) (bool, *ResponseErr) {
	if status := res.StatusCode(); status > 300 {
		return true, NewResponseErr(fmt.Sprintf("unexpected response status %v, expected value < 300", status), res.Body())
	}
	if hasErrHeader(res) {
		return true, NewResponseErr(fmt.Sprintf("eas responded with error header. status %v is considered an error in that case", res.StatusCode()), res.Body())
	}

	return false, nil
}

func hasErrHeader(res *resty.Response) bool {
	return res.Header().Get(HeaderKeyOtrisErr) == "true"
}
