package easclient

import (
	"context"
	"encoding/xml"
)

type StoreStatus struct {
	XMLName  xml.Name `xml:"status"`
	Registry struct {
		Records struct {
			All     int `xml:"all"`
			Indexed int `xml:"indexed"`
		} `xml:"records"`
		Attachments struct {
			All     int `xml:"all"`
			Indexed int `xml:"indexed"`
		} `xml:"attachments"`
	} `xml:"registry"`
	Index struct {
		Documents    int  `xml:"documents"`
		IsCurrent    bool `xml:"isCurrent"`
		HasDeletions bool `xml:"hasDeletions"`
	} `xml:"index"`
}

func (c *StoreClient) GetStoreStatus(ctx context.Context) (*StoreStatus, error) {
	req, err := c.newRequestXML(ctx)
	if err != nil {
		return nil, err
	}

	var result StoreStatus

	req.SetResult(&result)
	res, err := req.Get("/status")
	if err != nil {
		return nil, err
	}

	if _, err := isErrorResponse(res); err != nil {
		return nil, err
	}

	return &result, nil
}
