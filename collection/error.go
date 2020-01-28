package collection

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ErrResponse is the error that occurred as a result of non-200 error code in
// the response to the HTTP client's request.
type ErrResponse struct {
	Code    int    `json:"code"`
	Message string `json:"error"`
}

func (e *ErrResponse) Error() string {
	return fmt.Sprintf(
		"code: %d, %s",
		e.Code,
		e.Message,
	)
}

// checkResponse decorates any response into an error if response's status is
// not within 200-299 range.
func checkResponse(resp *http.Response) error {
	if c := resp.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	var errResp ErrResponse

	body, _ := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(body, &errResp); err == nil {
		return &errResp
	}

	// In case the decoding fails convert the entire body as a string.
	return &ErrResponse{
		Code:    resp.StatusCode,
		Message: string(body),
	}
}
