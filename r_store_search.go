package easclient

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/google/uuid"
)

type SearchRequest struct {
	Query string `json:"query"`

	ItemsPerPage int `json:"itemsPerPage"`
	StartIndex   int `json:"startIndex"`

	// Sort is the field to sort by
	Sort string `json:"sort"`
	// SortOrder is the order to sort by, either "asc" or "desc"
	SortOrder string `json:"sortOrder"`
}

func SearchRequestFromURL(s string) (*SearchRequest, error) {
	ur, err := url.Parse(s)
	if err != nil {
		return nil, err
	}

	vals := ur.Query()

	req := &SearchRequest{
		Query:        vals.Get("query"),
		ItemsPerPage: 0,
		StartIndex:   0,
		Sort:         vals.Get("sort"),
		SortOrder:    vals.Get("sortOrder"),
	}

	if vals.Has("itemsPerPage") {
		req.ItemsPerPage, err = strconv.Atoi(vals.Get("itemsPerPage"))
		if err != nil {
			return nil, err
		}
	}
	if vals.Has("startIndex") {
		req.StartIndex, err = strconv.Atoi(vals.Get("startIndex"))
		if err != nil {
			return nil, err
		}
	}

	return req, nil
}

func (request SearchRequest) ToQuery() map[string]string {
	return map[string]string{
		"query":        request.Query,
		"itemsPerPage": strconv.Itoa(request.ItemsPerPage),
		"startIndex":   strconv.Itoa(request.StartIndex),
		"sort":         request.Sort,
		"sortOrder":    request.SortOrder,
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
	Id               uuid.UUID    `json:"id"`
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

// SearchQuery is similar to Search but expects a URL from which SearchRequest is parsed via SearchRequestFromURL.
func (c *StoreClient) SearchQuery(ctx context.Context, url string) (*SearchResponse, error) {
	request, err := SearchRequestFromURL(url)
	if err != nil {
		return nil, fmt.Errorf("failed to parse search request: %w", err)
	}

	return c.Search(ctx, request)
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

	if _, err := isErrorResponse(res); err != nil {
		return nil, err
	}

	return &result, nil
}
