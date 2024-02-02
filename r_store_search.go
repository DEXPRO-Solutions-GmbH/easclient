package easclient

import (
	"context"
	"encoding/xml"
	"fmt"
	"net/url"
	"strconv"
	"time"

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
	q := map[string]string{
		"query":        request.Query,
		"itemsPerPage": strconv.Itoa(request.ItemsPerPage),
		"startIndex":   strconv.Itoa(request.StartIndex),
		"sort":         request.Sort,
		"sortOrder":    request.SortOrder,
	}

	// delete zero values which would result in an invalid request
	if q["itemsPerPage"] == "0" {
		delete(q, "itemsPerPage")
	}
	if q["startIndex"] == "0" {
		delete(q, "startIndex")
	}

	return q
}

type Link struct {
	Type string `xml:"type,attr"`
	Href string `xml:"href,attr"`
}

type SearchResponse struct {
	XMLName xml.Name               `xml:"rss"`
	Version string                 `xml:"version,attr"`
	Channel *SearchResponseChannel `xml:"channel"`
}

type SearchResponseChannel struct {
	Title        string `xml:"title"`
	Link         string `xml:"link"`
	Description  string `xml:"description"`
	TotalResults int    `xml:"totalResults"` // TODO: Assert in unmarshal test
	ItemsPerPage int    `xml:"itemsPerPage"` // TODO: Assert in unmarshal test
	StartIndex   int    `xml:"startIndex"`   // TODO: Assert in unmarshal test
	Query        struct {
		Role        string `xml:"role,attr"`
		SearchTerms string `xml:"searchTerms,attr"`
		StartPage   int    `xml:"startPage,attr"`
	} `xml:"Query"`
	Topn             int                   `xml:"topn"`
	EffectiveResults int                   `xml:"effectiveResults"`
	NextPage         string                `xml:"nextPage"`
	Items            []*SearchResponseItem `xml:"item"`
}

type SearchResponseItem struct {
	Title                  string         `xml:"title"`
	Link                   string         `xml:"link"`
	Description            string         `xml:"description"`
	Score                  float64        `xml:"score"`
	ExplainLink            Link           `xml:"explainLink"`
	VersionLink            Link           `xml:"versionLink"`
	HistoryLink            Link           `xml:"historyLink"`
	VerifyLink             Link           `xml:"verifyLink"`
	DocumentType           string         `xml:"documentType"`
	Fields                 []*RecordField `xml:"field"` // TODO: Assert and check in get attachment response if this is the correct way to handle recurring fields
	MasterId               uuid.UUID      `xml:"masterId"`
	ArchiveDateTime        time.Time      `xml:"archiveDateTime"`
	ID                     uuid.UUID      `xml:"id"`
	Version                string         `xml:"version"`
	ArchiverLogin          string         `xml:"archiverLogin"`
	Archiver               string         `xml:"archiver"`
	InitialArchiver        string         `xml:"initialArchiver"`
	InitialArchiverLogin   string         `xml:"initialArchiverLogin"`
	InitialArchiveDateTime time.Time      `xml:"initialArchiveDateTime"`
}

// SearchQuery is similar to Search but expects a URL from which SearchRequest is parsed via SearchRequestFromURL.
func (c *StoreClient) SearchQuery(ctx context.Context, url string) (*SearchResponseChannel, error) {
	request, err := SearchRequestFromURL(url)
	if err != nil {
		return nil, fmt.Errorf("failed to parse search request: %w", err)
	}

	return c.Search(ctx, request)
}

func (c *StoreClient) Search(ctx context.Context, request *SearchRequest) (*SearchResponseChannel, error) {
	req, err := c.newRequestXML(ctx)
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

	return result.Channel, nil
}
