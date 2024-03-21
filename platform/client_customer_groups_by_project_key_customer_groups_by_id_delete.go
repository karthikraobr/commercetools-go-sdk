package platform

// Generated file, please do not change!!!

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type ByProjectKeyCustomerGroupsByIDRequestMethodDelete struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyCustomerGroupsByIDRequestMethodDeleteInput
}

func (r *ByProjectKeyCustomerGroupsByIDRequestMethodDelete) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyCustomerGroupsByIDRequestMethodDeleteInput struct {
	Version int
	Expand  []string
}

func (input *ByProjectKeyCustomerGroupsByIDRequestMethodDeleteInput) Values() url.Values {
	values := url.Values{}
	values.Add("version", strconv.Itoa(input.Version))
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyCustomerGroupsByIDRequestMethodDelete) Version(v int) *ByProjectKeyCustomerGroupsByIDRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyCustomerGroupsByIDRequestMethodDeleteInput{}
	}
	rb.params.Version = v
	return rb
}

func (rb *ByProjectKeyCustomerGroupsByIDRequestMethodDelete) Expand(v []string) *ByProjectKeyCustomerGroupsByIDRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyCustomerGroupsByIDRequestMethodDeleteInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyCustomerGroupsByIDRequestMethodDelete) WithQueryParams(input ByProjectKeyCustomerGroupsByIDRequestMethodDeleteInput) *ByProjectKeyCustomerGroupsByIDRequestMethodDelete {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyCustomerGroupsByIDRequestMethodDelete) WithHeaders(headers http.Header) *ByProjectKeyCustomerGroupsByIDRequestMethodDelete {
	rb.headers = headers
	return rb
}
func (rb *ByProjectKeyCustomerGroupsByIDRequestMethodDelete) Execute(ctx context.Context) (result *CustomerGroup, err error) {
	var queryParams url.Values
	if rb.params != nil {
		queryParams = rb.params.Values()
	} else {
		queryParams = url.Values{}
	}
	resp, err := rb.client.delete(
		ctx,
		rb.url,
		queryParams,
		rb.headers,
		nil,
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
		if err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 400:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 401:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 403:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 404:
		return nil, ErrNotFound
	case 500:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 502:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	case 503:
		errorObj := ErrorResponse{}
		err = json.Unmarshal(content, &errorObj)
		if err != nil {
			return nil, err
		}
		return nil, errorObj
	default:
		result := GenericRequestError{
			StatusCode: resp.StatusCode,
			Content:    content,
			Response:   resp,
		}
		return nil, result
	}

}
