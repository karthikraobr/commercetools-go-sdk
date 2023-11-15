package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyBusinessUnitsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyBusinessUnitsRequestBuilder) WithKey(key string) *ByProjectKeyBusinessUnitsKeyByKeyRequestBuilder {
	return &ByProjectKeyBusinessUnitsKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyBusinessUnitsRequestBuilder) WithId(id string) *ByProjectKeyBusinessUnitsByIDRequestBuilder {
	return &ByProjectKeyBusinessUnitsByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyBusinessUnitsRequestBuilder) Get() *ByProjectKeyBusinessUnitsRequestMethodGet {
	return &ByProjectKeyBusinessUnitsRequestMethodGet{
		url:    fmt.Sprintf("/%s/business-units", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks if a BusinessUnit exists for a given Query Predicate. Returns a `200 OK` status if any BusinessUnits match the Query Predicate or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyBusinessUnitsRequestBuilder) Head() *ByProjectKeyBusinessUnitsRequestMethodHead {
	return &ByProjectKeyBusinessUnitsRequestMethodHead{
		url:    fmt.Sprintf("/%s/business-units", rb.projectKey),
		client: rb.client,
	}
}

func (rb *ByProjectKeyBusinessUnitsRequestBuilder) Post(body BusinessUnitDraft) *ByProjectKeyBusinessUnitsRequestMethodPost {
	return &ByProjectKeyBusinessUnitsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/business-units", rb.projectKey),
		client: rb.client,
	}
}
