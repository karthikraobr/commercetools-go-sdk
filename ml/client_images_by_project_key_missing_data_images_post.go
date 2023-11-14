package ml

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ByProjectKeyMissingDataImagesRequestMethodPost struct {
	body    MissingImagesSearchRequest
	url     string
	client  *Client
	headers http.Header
}

func (r *ByProjectKeyMissingDataImagesRequestMethodPost) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url": r.url,
	}
}

func (rb *ByProjectKeyMissingDataImagesRequestMethodPost) WithHeaders(headers http.Header) *ByProjectKeyMissingDataImagesRequestMethodPost {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyMissingDataImagesRequestMethodPost) Execute(ctx context.Context) (result *TaskToken, err error) {
	data, err := serializeInput(rb.body)
	if err != nil {
		return nil, err
	}
	queryParams := url.Values{}
	resp, err := rb.client.post(
		ctx,
		rb.url,
		queryParams,
		rb.headers,
		data,
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
	case 202:
		err = json.Unmarshal(content, &result)
		if err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return nil, result
	}

}
