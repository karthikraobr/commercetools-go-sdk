package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyDiscountCodesKeyByKeyRequestBuilder struct {
	projectKey string
	key        string
	client     *Client
}

/**
*	Deprecated scope: `view_orders:{projectKey}`
 */
func (rb *ByProjectKeyDiscountCodesKeyByKeyRequestBuilder) Get() *ByProjectKeyDiscountCodesKeyByKeyRequestMethodGet {
	return &ByProjectKeyDiscountCodesKeyByKeyRequestMethodGet{
		url:    fmt.Sprintf("/%s/discount-codes/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Checks if a DiscountCode exists for a given `key`. Returns a `200 OK` status if the DiscountCode exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyDiscountCodesKeyByKeyRequestBuilder) Head() *ByProjectKeyDiscountCodesKeyByKeyRequestMethodHead {
	return &ByProjectKeyDiscountCodesKeyByKeyRequestMethodHead{
		url:    fmt.Sprintf("/%s/discount-codes/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Deprecated scope: `manage_orders:{projectKey}`
 */
func (rb *ByProjectKeyDiscountCodesKeyByKeyRequestBuilder) Post(body DiscountCodeUpdate) *ByProjectKeyDiscountCodesKeyByKeyRequestMethodPost {
	return &ByProjectKeyDiscountCodesKeyByKeyRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/discount-codes/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}

/**
*	Deleting a Discount Code produces the [DiscountCodeDeleted](ctp:api:type:DiscountCodeDeletedMessage) Message.
*
*	Deprecated scope: `manage_orders:{projectKey}`
*
 */
func (rb *ByProjectKeyDiscountCodesKeyByKeyRequestBuilder) Delete() *ByProjectKeyDiscountCodesKeyByKeyRequestMethodDelete {
	return &ByProjectKeyDiscountCodesKeyByKeyRequestMethodDelete{
		url:    fmt.Sprintf("/%s/discount-codes/key=%s", rb.projectKey, rb.key),
		client: rb.client,
	}
}
