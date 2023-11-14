package connect

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ConnectorsDraftsByIDRequestMethodGet struct {
	url     string
	client  *Client
	headers http.Header
}

func (r *ConnectorsDraftsByIDRequestMethodGet) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *ConnectorsDraftsByIDRequestMethodGet) WithHeaders(headers http.Header) *ConnectorsDraftsByIDRequestMethodGet {
	rb.headers = headers
	return rb
}
func (rb *ConnectorsDraftsByIDRequestMethodGet) Execute(ctx context.Context) (result *ConnectorStaged, err error) {
	queryParams := url.Values{}
	resp, err := rb.client.get(
		ctx,
		rb.url,
		queryParams,
		rb.headers,
	)

	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	switch resp.StatusCode {
	case 200:
		err = json.Unmarshal(content, &result)
		return result, nil
	case 404:
		return nil, ErrNotFound
	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return nil, result
	}

}
