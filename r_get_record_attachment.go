package easclient

import (
	"context"
	"fmt"
	"io"

	"github.com/google/uuid"
)

// GetRecordAttachment retrieves the actual attachment. It is written to the given io.Writer.
//
// This returns the number of bytes written and an error if any.
func (c *StoreClient) GetRecordAttachment(ctx context.Context, writer io.Writer, recordID, attachmentID uuid.UUID) (int64, error) {
	req, err := c.newRequestJSON(ctx)
	if err != nil {
		return 0, err
	}

	// This endpoints responds with the binary attachment. Thus, we do not want to
	// automatically parse the body and must also not set the Accept header.
	req.SetDoNotParseResponse(true)
	req.Header.Del("Accept")

	res, err := req.Get(fmt.Sprintf("/record/%s/attachment/%s", recordID, attachmentID))
	if err != nil {
		return 0, err
	}

	resBody := res.RawBody()
	defer resBody.Close()

	if _, err := isErrorResponse(res); err != nil {
		return 0, err
	}

	n, err := io.Copy(writer, resBody)
	if err != nil {
		return n, fmt.Errorf("failed to copy attachment into byte buffer: %w", err)
	}

	return n, nil
}
