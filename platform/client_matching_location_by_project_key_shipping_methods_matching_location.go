package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyShippingMethodsMatchingLocationRequestBuilder struct {
	projectKey string
	client     *Client
}

/**
*	Retrieves all the ShippingMethods that can ship to the given [Location](/projects/zones#location).
*	ShippingMethods that have a `predicate` defined are automatically disqualified.
*	If the `currency` parameter is given, then the ShippingMethods must also have a rate defined in the specified currency.
*	Each ShippingMethod contains at least one ShippingRate with the flag `isMatching` set to `true`.
*	If the `currency` parameter is given, exactly one ShippingRate will contain it.
*
 */
func (rb *ByProjectKeyShippingMethodsMatchingLocationRequestBuilder) Get() *ByProjectKeyShippingMethodsMatchingLocationRequestMethodGet {
	return &ByProjectKeyShippingMethodsMatchingLocationRequestMethodGet{
		url:    fmt.Sprintf("/%s/shipping-methods/matching-location", rb.projectKey),
		client: rb.client,
	}
}

/**
*	Checks if a ShippingMethod that can ship to the given [Location](ctp:api:type:Location) exists. Returns a `200 OK` status if the ShippingMethod exists or a `404 Not Found` otherwise.
 */
func (rb *ByProjectKeyShippingMethodsMatchingLocationRequestBuilder) Head() *ByProjectKeyShippingMethodsMatchingLocationRequestMethodHead {
	return &ByProjectKeyShippingMethodsMatchingLocationRequestMethodHead{
		url:    fmt.Sprintf("/%s/shipping-methods/matching-location", rb.projectKey),
		client: rb.client,
	}
}
