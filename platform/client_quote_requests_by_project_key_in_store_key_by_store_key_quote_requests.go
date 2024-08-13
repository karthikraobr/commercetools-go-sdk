package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsRequestBuilder struct {
	projectKey string
	storeKey   string
	client     *Client
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsRequestBuilder) WithKey(key string) *ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsKeyByKeyRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsRequestBuilder) WithId(id string) *ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsByIDRequestBuilder {
	return &ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		storeKey:   rb.storeKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsRequestBuilder) Get() *ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsRequestMethodGet {
	return &ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsRequestMethodGet{
		url:    fmt.Sprintf("/%s/in-store/key=%s/quote-requests", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}

/**
*	Checks if a QuoteRequest exists for a given Query Predicate. Returns a `200 OK` status if any QuoteRequests match the Query Predicate or a [ResourceNotFound](ctp:api:type:ResourceNotFoundError) error otherwise.
 */
func (rb *ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsRequestBuilder) Head() *ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsRequestMethodHead {
	return &ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsRequestMethodHead{
		url:    fmt.Sprintf("/%s/in-store/key=%s/quote-requests", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}

func (rb *ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsRequestBuilder) Post(body QuoteRequestDraft) *ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsRequestMethodPost {
	return &ByProjectKeyInStoreKeyByStoreKeyQuoteRequestsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/in-store/key=%s/quote-requests", rb.projectKey, rb.storeKey),
		client: rb.client,
	}
}
