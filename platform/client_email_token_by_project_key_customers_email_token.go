package platform

// Generated file, please do not change!!!

import (
	"fmt"
)

type ByProjectKeyCustomersEmailTokenRequestBuilder struct {
	projectKey string
	client     *Client
}

/**
*	Produces the [CustomerEmailTokenCreated](ctp:api:type:CustomerEmailTokenCreatedMessage) Message.
 */
func (rb *ByProjectKeyCustomersEmailTokenRequestBuilder) Post(body CustomerCreateEmailToken) *ByProjectKeyCustomersEmailTokenRequestMethodPost {
	return &ByProjectKeyCustomersEmailTokenRequestMethodPost{
		body:   body,
		url:    fmt.Sprintf("/%s/customers/email-token", rb.projectKey),
		client: rb.client,
	}
}
