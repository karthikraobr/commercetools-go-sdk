package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyTaxCategoriesRequestBuilder struct {
	projectKey string
	client     *Client
}

func (rb *ByProjectKeyTaxCategoriesRequestBuilder) WithKey(key string) *ByProjectKeyTaxCategoriesKeyByKeyRequestBuilder {
	return &ByProjectKeyTaxCategoriesKeyByKeyRequestBuilder{
		key:        key,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyTaxCategoriesRequestBuilder) WithId(id string) *ByProjectKeyTaxCategoriesByIDRequestBuilder {
	return &ByProjectKeyTaxCategoriesByIDRequestBuilder{
		id:         id,
		projectKey: rb.projectKey,
		client:     rb.client,
	}
}
func (rb *ByProjectKeyTaxCategoriesRequestBuilder) Get() *ByProjectKeyTaxCategoriesRequestMethodGet {
	return &ByProjectKeyTaxCategoriesRequestMethodGet{
		url:    fmt.Sprintf("/%s/tax-categories", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks if a TaxCategory exists for a given Query Predicate. Returns a `200 OK` status if any TaxCategories match the Query Predicate, or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyTaxCategoriesRequestBuilder) Head() *ByProjectKeyTaxCategoriesRequestMethodHead {
	return &ByProjectKeyTaxCategoriesRequestMethodHead{
		url:    fmt.Sprintf("/%s/tax-categories", rb.projectKey),
		client: rb.client,
	}
}

func (rb *ByProjectKeyTaxCategoriesRequestBuilder) Post(body TaxCategoryDraft) *ByProjectKeyTaxCategoriesRequestMethodPost {
	return &ByProjectKeyTaxCategoriesRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/tax-categories", rb.projectKey),
		client: rb.client,
	}
}
