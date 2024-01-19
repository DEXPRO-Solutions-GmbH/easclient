package easclient

import (
	"context"
	"time"
)

type StoreStatus struct {
	Registry struct {
		AllRecords         int `json:"allRecords"`
		IndexedRecords     int `json:"indexedRecords"`
		AllAttachments     int `json:"allAttachments"`
		IndexedAttachments int `json:"indexedAttachments"`
	} `json:"registry"`
	Index struct {
		Documents    int  `json:"documents"`
		IsCurrent    bool `json:"isCurrent"`
		HasDeletions bool `json:"hasDeletions"`
		Records      int  `json:"records"`
		Attachments  int  `json:"attachments"`
	} `json:"index"`
	Capacity struct {
		Maximum     int64     `json:"maximum"`
		Utilized    float64   `json:"utilized"`
		GrowthRate  float64   `json:"growthRate"`
		ExpectedEnd time.Time `json:"expectedEnd"`
		Lifetime    int       `json:"lifetime"`
	} `json:"capacity"`
	Periods []struct {
		Start    string `json:"start"`
		End      string `json:"end"`
		Registry struct {
			AllRecords         int `json:"allRecords"`
			IndexedRecords     int `json:"indexedRecords"`
			AllAttachments     int `json:"allAttachments"`
			IndexedAttachments int `json:"indexedAttachments"`
		} `json:"registry"`
		Index struct {
			Records     int `json:"records"`
			Attachments int `json:"attachments"`
		} `json:"index"`
		Capacity struct {
			Utilized float64 `json:"utilized"`
		} `json:"capacity"`
	} `json:"periods"`
}

func (c *StoreClient) GetStoreStatus(ctx context.Context) (*StoreStatus, error) {
	req, err := c.newRequest(ctx)
	if err != nil {
		return nil, err
	}

	type Res struct {
		Status *StoreStatus `json:"status"`
	}

	var result Res

	req.SetResult(&result)
	res, err := req.Get("/status")
	if err != nil {
		return nil, err
	}

	if _, err := isErrorResponse(res); err != nil {
		return nil, err
	}

	return result.Status, nil
}
