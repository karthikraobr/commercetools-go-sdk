package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyCartsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyCartsRequestBuilder) Replicate() *ByProjectKeyCartsReplicateRequestBuilder {
	return &ByProjectKeyCartsReplicateRequestBuilder{
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyCartsRequestBuilder) WithCustomerId(customerId string) *ByProjectKeyCartsCustomerIdByCustomerIdRequestBuilder {
	return &ByProjectKeyCartsCustomerIdByCustomerIdRequestBuilder{
		customerId: customerId,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyCartsRequestBuilder) WithKey(key string) *ByProjectKeyCartsKeyByKeyRequestBuilder {
	return &ByProjectKeyCartsKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyCartsRequestBuilder) WithId(id string) *ByProjectKeyCartsByIDRequestBuilder {
	return &ByProjectKeyCartsByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyCartsRequestBuilder) Get() *ByProjectKeyCartsRequestMethodGet {
	return &ByProjectKeyCartsRequestMethodGet{
		url:    fmt.Sprintf("/%s/carts", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks if a Cart exists for a given Query Predicate. Returns a `200 OK` status if any Carts match the Query Predicate, or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyCartsRequestBuilder) Head() *ByProjectKeyCartsRequestMethodHead {
	return &ByProjectKeyCartsRequestMethodHead{
		url:    fmt.Sprintf("/%s/carts", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Creating a Cart fails with an [InvalidOperation](ctp:api:type:InvalidOperationError) error if the
*	[ShippingMethod](ctp:api:type:ShippingMethod) referenced in the CartDraft
*	has a `predicate` that does not match the Cart.
*
 */
func (rb *ByProjectKeyCartsRequestBuilder) Post(body CartDraft) *ByProjectKeyCartsRequestMethodPost {
	return &ByProjectKeyCartsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/carts", rb.projectKey),
		client: rb.client,
	}
}
