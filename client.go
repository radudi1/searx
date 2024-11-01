// searx is a simple golang client for searx and searxng search engines
// It relies on searx having json support enabled
package searx

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// SearchOptions contains the search options for a searx
// The names are similar to searx API.
// See  https://docs.searxng.org/dev/search_api.html
// Not all searx API options are currently supported by this package.
type SearchOptions struct {
	Categories []string
	Engines    []string
	Language   string
	TimeRange  string
	SafeSearch uint
}

// Client is a searx client that can perform searches on searx servers.
type Client struct {
	Url string
	SearchOptions
}

// NewClient creates a new searx client with the given url and optional search options.
// If search options are provided they will be used as defaults for each search.
func NewClient(url string, options *SearchOptions) *Client {
	c := Client{
		Url: url,
	}
	if options != nil {
		c.SearchOptions = *options
	}
	return &c
}

// Search performs a search using the given query and options.
// If no options are provided, the default options will be used.
// The function returns the search response or an error.
func (c *Client) Search(query string, options *SearchOptions) (searchResponse SearchResponse, err error) {

	// prepare url
	searxUrl := c.Url + "/search?q=" + url.QueryEscape(query) + "&format=json"
	if options == nil {
		options = &c.SearchOptions
	}
	if len(options.Categories) > 0 {
		searxUrl += "&categories=" + url.QueryEscape(strings.Join(options.Categories, ","))
	}
	if len(options.Engines) > 0 {
		searxUrl += "&engines=" + url.QueryEscape(strings.Join(options.Engines, ","))
	}
	if options.Language != "" {
		searxUrl += "&language=" + url.QueryEscape(options.Language)
	}
	if options.TimeRange != "" {
		searxUrl += "&time_range=" + url.QueryEscape(options.TimeRange)
	}
	if options.SafeSearch > 0 {
		searxUrl += "&safe_search=" + strconv.Itoa(int(options.SafeSearch))
	}

	// make request
	resp, err := http.Get(searxUrl)
	if err != nil {
		return searchResponse, err
	}
	defer resp.Body.Close()

	// parse response
	err = json.NewDecoder(resp.Body).Decode(&searchResponse)
	if err != nil {
		return searchResponse, err
	}

	return searchResponse, nil

}
