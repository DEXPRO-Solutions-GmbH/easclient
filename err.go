package easclient

import (
	"fmt"

	"gopkg.in/resty.v1"
)

const (
	HeaderKeyOtrisErr = "X-Otris-Eas-Error"
)

type ResponseErr struct {
	msg string

	Body []byte
}

func (r *ResponseErr) Error() string {
	return r.msg
}

// isErrorResponse checks if the given response is an error.
func isErrorResponse(res *resty.Response) (bool, *ResponseErr) {
	if status := res.StatusCode(); status > 300 {
		return true, &ResponseErr{
			msg:  fmt.Sprintf("unexpected response status %v, expected value < 300", status),
			Body: res.Body(),
		}
	}
	if hasErrHeader(res) {
		return true, &ResponseErr{
			msg:  fmt.Sprintf("eas responded with error header. status %v is considered an error in that case", res.StatusCode()),
			Body: res.Body(),
		}
	}

	return false, nil
}

func hasErrHeader(res *resty.Response) bool {
	return res.Header().Get(HeaderKeyOtrisErr) == "true"
}
