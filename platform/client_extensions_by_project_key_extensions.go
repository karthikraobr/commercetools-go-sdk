package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyExtensionsRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyExtensionsRequestBuilder) WithKey(key string) *ByProjectKeyExtensionsKeyByKeyRequestBuilder {
	return &ByProjectKeyExtensionsKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyExtensionsRequestBuilder) WithId(id string) *ByProjectKeyExtensionsByIDRequestBuilder {
	return &ByProjectKeyExtensionsByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyExtensionsRequestBuilder) Get() *ByProjectKeyExtensionsRequestMethodGet {
	return &ByProjectKeyExtensionsRequestMethodGet{
		url:    fmt.Sprintf("/%s/extensions", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks if an Extension exists for a given Query Predicate. Returns a `200 OK` status if any Extensions match the Query Predicate or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyExtensionsRequestBuilder) Head() *ByProjectKeyExtensionsRequestMethodHead {
	return &ByProjectKeyExtensionsRequestMethodHead{
		url:    fmt.Sprintf("/%s/extensions", rb.projectKey),
		client: rb.client,
	}
}

func (rb *ByProjectKeyExtensionsRequestBuilder) Post(body ExtensionDraft) *ByProjectKeyExtensionsRequestMethodPost {
	return &ByProjectKeyExtensionsRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/extensions", rb.projectKey),
		client: rb.client,
	}
}
