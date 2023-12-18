package easclient

import (
	"context"
	"fmt"
)

type SearchRequest struct {
	Query string `json:"query"`
}

func (request SearchRequest) ToQuery() map[string]string {
	return map[string]string{
		"query": request.Query,
	}
}

type Link struct {
	Type  string `json:"type"`
	Title string `json:"title"`
	Href  string `json:"href"`
}

type SearchResult struct {
	Title            bool         `json:"title"`
	Score            float64      `json:"score"`
	Id               string       `json:"id"`
	FileLink         Link         `json:"fileLink"`
	ExplainLink      Link         `json:"explainLink"`
	CheckVersionLink Link         `json:"checkVersionLink"`
	HistoryLink      Link         `json:"historyLink"`
	VerifyLink       Link         `json:"verifyLink"`
	HeaderFields     HeaderFields `json:"headerFields"`
	RecordFields     RecordFields `json:"recordFields"`
}

type SearchResponse struct {
	Query            string          `json:"query"`
	TotalHits        int             `json:"totalHits"`
	ItemsPerPage     int             `json:"itemsPerPage"`
	StartIndex       int             `json:"startIndex"`
	Topn             int             `json:"topn"`
	EffectiveResults int             `json:"effectiveResults"`
	Result           []*SearchResult `json:"result"`
}

func (c *StoreClient) Search(ctx context.Context, request *SearchRequest) (*SearchResponse, error) {
	req, err := c.newRequest(ctx)
	if err != nil {
		return nil, err
	}

	var result SearchResponse

	req.SetResult(&result)

	req.SetQueryParams(request.ToQuery())
	res, err := req.Get("")
	if err != nil {
		return nil, err
	}

	if status := res.StatusCode(); status != 200 {
		return nil, fmt.Errorf("unexpected response status %v", status)
	}

	return &result, nil
}
