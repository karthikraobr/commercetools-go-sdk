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

type ByProjectKeyInStoreKeyByStoreKeyMeCartsByIDRequestMethodDelete struct {
	url     string
	client  *Client
	headers http.Header
	params  *ByProjectKeyInStoreKeyByStoreKeyMeCartsByIDRequestMethodDeleteInput
}

func (r *ByProjectKeyInStoreKeyByStoreKeyMeCartsByIDRequestMethodDelete) Dump() map[string]interface{} {
	return map[string]interface{}{
		"url":    r.url,
		"params": r.params,
	}
}

type ByProjectKeyInStoreKeyByStoreKeyMeCartsByIDRequestMethodDeleteInput struct {
	Version int
	Expand  []string
}

func (input *ByProjectKeyInStoreKeyByStoreKeyMeCartsByIDRequestMethodDeleteInput) Values() url.Values {
	values := url.Values{}
	values.Add("version", strconv.Itoa(input.Version))
	for _, v := range input.Expand {
		values.Add("expand", fmt.Sprintf("%v", v))
	}
	return values
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyMeCartsByIDRequestMethodDelete) Version(v int) *ByProjectKeyInStoreKeyByStoreKeyMeCartsByIDRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyMeCartsByIDRequestMethodDeleteInput{}
	}
	rb.params.Version = v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyMeCartsByIDRequestMethodDelete) Expand(v []string) *ByProjectKeyInStoreKeyByStoreKeyMeCartsByIDRequestMethodDelete {
	if rb.params == nil {
		rb.params = &ByProjectKeyInStoreKeyByStoreKeyMeCartsByIDRequestMethodDeleteInput{}
	}
	rb.params.Expand = v
	return rb
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyMeCartsByIDRequestMethodDelete) WithQueryParams(input ByProjectKeyInStoreKeyByStoreKeyMeCartsByIDRequestMethodDeleteInput) *ByProjectKeyInStoreKeyByStoreKeyMeCartsByIDRequestMethodDelete {
	rb.params = &input
	return rb
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeCartsByIDRequestMethodDelete) WithHeaders(headers http.Header) *ByProjectKeyInStoreKeyByStoreKeyMeCartsByIDRequestMethodDelete {
	rb.headers = headers
	return rb
}

/**
*	Deletes the Cart for a given `id` in a Store. Returns a `200 OK` status if successful.
*
*	A [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error is returned in the following scenarios:
*
*	- If no Cart exists in the Store for the given `id`.
*	- If the Cart exists in the Project but does not belong to a Store, or the Cart's `store` field references a different Store.
*	- If the Cart exists in the Project but does not have either a `customerId` that matches the [customer:{id}](/scopes#composable-commerce-oauth) scope, or an `anonymousId` that matches the [anonymous_id:{id}](/scopes#composable-commerce-oauth) scope.
*
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyMeCartsByIDRequestMethodDelete) Execute(ctx context.Context) (result *Cart, err error) {
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
