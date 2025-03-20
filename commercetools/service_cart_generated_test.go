// Automatically generated, do not edit

package commercetools_test

import (
	"context"
	"github.com/labd/commercetools-go-sdk/commercetools"
	"github.com/labd/commercetools-go-sdk/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeneratedCartGetWithCustomerID(t *testing.T) {
	responseData := ` {
	  "type": "Cart",
	  "id": "2a3baa00-44fa-4ab8-bec7-933c31e18dcc",
	  "key": "test-key",
	  "version": 5,
	  "createdAt": "2015-09-22T15:36:17.510Z",
	  "lastModifiedAt": "2015-09-22T15:41:55.816Z",
	  "lineItems": [
	    {
	      "id": "b925a817-d5ad-48bb-a407-29ad8e0649b5",
	      "productId": "9f10dcfb-5cc9-4a18-843a-c07f7e22d01f",
	      "name": {
	        "en": "SAPPHIRE"
	      },
	      "productType": {
	        "typeId": "product-type",
	        "id": "2543e1d8-4915-4f72-a3c9-1df9b1b0082d",
	        "version": 8
	      },
	      "productSlug": {
	        "en": "sapphire1421832124423"
	      },
	      "variant": {
	        "id": 1,
	        "sku": "sku_SAPPHIRE_variant1_1421832124423",
	        "prices": [
	          {
	            "value": {
	              "type": "centPrecision",
	              "fractionDigits": 2,
	              "currencyCode": "EUR",
	              "centAmount": 2800
	            },
	            "id": "8da659ef-9e54-447d-9c36-84912db1848f"
	          }
	        ],
	        "images": [
	          {
	            "url": "https://www.commercetools.com/cli/data/252542005_1.jpg",
	            "dimensions": {
	              "w": 1400,
	              "h": 1400
	            }
	          }
	        ],
	        "attributes": [],
	        "assets": []
	      },
	      "price": {
	        "value": {
	          "type": "centPrecision",
	          "fractionDigits": 2,
	          "currencyCode": "EUR",
	          "centAmount": 2800
	        },
	        "id": "8da659ef-9e54-447d-9c36-84912db1848f"
	      },
	      "quantity": 2,
	      "discountedPricePerQuantity": [],
	      "state": [
	        {
	          "quantity": 2,
	          "state": {
	            "typeId": "state",
	            "id": "7c2e2694-aefe-43d7-888e-6a99514caaca"
	          }
	        }
	      ],
	      "priceMode": "Platform",
	      "lineItemMode": "Standard",
	      "totalPrice": {
	        "type": "centPrecision",
	        "fractionDigits": 2,
	        "currencyCode": "EUR",
	        "centAmount": 5600
	      }
	    }
	  ],
	  "store": {
	    "typeId": "store",
	    "key": "test-key"
	  },
	  "cartState": "Active",
	  "totalPrice": {
	    "type": "centPrecision",
	    "fractionDigits": 2,
	    "currencyCode": "EUR",
	    "centAmount": 5600
	  },
	  "customLineItems": [],
	  "discountCodes": [],
	  "inventoryMode": "None",
	  "taxMode": "Platform",
	  "taxRoundingMode": "HalfEven",
	  "taxCalculationMode": "LineItemLevel",
	  "refusedGifts": [],
	  "origin": "Customer"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	cart, err := client.CartGetWithCustomerID(context.TODO(), "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, cart)
	assert.NotNil(t, cart.Version)
	assert.NotEmpty(t, cart.TaxRoundingMode)
	assert.NotEmpty(t, cart.TaxMode)
	assert.NotEmpty(t, cart.TaxCalculationMode)
	assert.NotEmpty(t, cart.Origin)
	assert.NotEmpty(t, cart.LastModifiedAt)
	assert.NotEmpty(t, cart.Key)
	assert.NotEmpty(t, cart.InventoryMode)
	assert.NotEmpty(t, cart.ID)
	assert.NotEmpty(t, cart.CreatedAt)
	assert.NotEmpty(t, cart.CartState)

}

func TestGeneratedCartGetWithID(t *testing.T) {
	responseData := ` {
	  "type": "Cart",
	  "id": "2a3baa00-44fa-4ab8-bec7-933c31e18dcc",
	  "key": "test-key",
	  "version": 5,
	  "createdAt": "2015-09-22T15:36:17.510Z",
	  "lastModifiedAt": "2015-09-22T15:41:55.816Z",
	  "lineItems": [
	    {
	      "id": "b925a817-d5ad-48bb-a407-29ad8e0649b5",
	      "productId": "9f10dcfb-5cc9-4a18-843a-c07f7e22d01f",
	      "name": {
	        "en": "SAPPHIRE"
	      },
	      "productType": {
	        "typeId": "product-type",
	        "id": "2543e1d8-4915-4f72-a3c9-1df9b1b0082d",
	        "version": 8
	      },
	      "productSlug": {
	        "en": "sapphire1421832124423"
	      },
	      "variant": {
	        "id": 1,
	        "sku": "sku_SAPPHIRE_variant1_1421832124423",
	        "prices": [
	          {
	            "value": {
	              "type": "centPrecision",
	              "fractionDigits": 2,
	              "currencyCode": "EUR",
	              "centAmount": 2800
	            },
	            "id": "8da659ef-9e54-447d-9c36-84912db1848f"
	          }
	        ],
	        "images": [
	          {
	            "url": "https://www.commercetools.com/cli/data/252542005_1.jpg",
	            "dimensions": {
	              "w": 1400,
	              "h": 1400
	            }
	          }
	        ],
	        "attributes": [],
	        "assets": []
	      },
	      "price": {
	        "value": {
	          "type": "centPrecision",
	          "fractionDigits": 2,
	          "currencyCode": "EUR",
	          "centAmount": 2800
	        },
	        "id": "8da659ef-9e54-447d-9c36-84912db1848f"
	      },
	      "quantity": 2,
	      "discountedPricePerQuantity": [],
	      "state": [
	        {
	          "quantity": 2,
	          "state": {
	            "typeId": "state",
	            "id": "7c2e2694-aefe-43d7-888e-6a99514caaca"
	          }
	        }
	      ],
	      "priceMode": "Platform",
	      "lineItemMode": "Standard",
	      "totalPrice": {
	        "type": "centPrecision",
	        "fractionDigits": 2,
	        "currencyCode": "EUR",
	        "centAmount": 5600
	      }
	    }
	  ],
	  "store": {
	    "typeId": "store",
	    "key": "test-key"
	  },
	  "cartState": "Active",
	  "totalPrice": {
	    "type": "centPrecision",
	    "fractionDigits": 2,
	    "currencyCode": "EUR",
	    "centAmount": 5600
	  },
	  "customLineItems": [],
	  "discountCodes": [],
	  "inventoryMode": "None",
	  "taxMode": "Platform",
	  "taxRoundingMode": "HalfEven",
	  "taxCalculationMode": "LineItemLevel",
	  "refusedGifts": [],
	  "origin": "Customer"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	cart, err := client.CartGetWithID(context.TODO(), "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, cart)
	assert.NotNil(t, cart.Version)
	assert.NotEmpty(t, cart.TaxRoundingMode)
	assert.NotEmpty(t, cart.TaxMode)
	assert.NotEmpty(t, cart.TaxCalculationMode)
	assert.NotEmpty(t, cart.Origin)
	assert.NotEmpty(t, cart.LastModifiedAt)
	assert.NotEmpty(t, cart.Key)
	assert.NotEmpty(t, cart.InventoryMode)
	assert.NotEmpty(t, cart.ID)
	assert.NotEmpty(t, cart.CreatedAt)
	assert.NotEmpty(t, cart.CartState)

}

func TestGeneratedCartGetWithKey(t *testing.T) {
	responseData := ` {
	  "type": "Cart",
	  "id": "2a3baa00-44fa-4ab8-bec7-933c31e18dcc",
	  "key": "test-key",
	  "version": 5,
	  "createdAt": "2015-09-22T15:36:17.510Z",
	  "lastModifiedAt": "2015-09-22T15:41:55.816Z",
	  "lineItems": [
	    {
	      "id": "b925a817-d5ad-48bb-a407-29ad8e0649b5",
	      "productId": "9f10dcfb-5cc9-4a18-843a-c07f7e22d01f",
	      "name": {
	        "en": "SAPPHIRE"
	      },
	      "productType": {
	        "typeId": "product-type",
	        "id": "2543e1d8-4915-4f72-a3c9-1df9b1b0082d",
	        "version": 8
	      },
	      "productSlug": {
	        "en": "sapphire1421832124423"
	      },
	      "variant": {
	        "id": 1,
	        "sku": "sku_SAPPHIRE_variant1_1421832124423",
	        "prices": [
	          {
	            "value": {
	              "type": "centPrecision",
	              "fractionDigits": 2,
	              "currencyCode": "EUR",
	              "centAmount": 2800
	            },
	            "id": "8da659ef-9e54-447d-9c36-84912db1848f"
	          }
	        ],
	        "images": [
	          {
	            "url": "https://www.commercetools.com/cli/data/252542005_1.jpg",
	            "dimensions": {
	              "w": 1400,
	              "h": 1400
	            }
	          }
	        ],
	        "attributes": [],
	        "assets": []
	      },
	      "price": {
	        "value": {
	          "type": "centPrecision",
	          "fractionDigits": 2,
	          "currencyCode": "EUR",
	          "centAmount": 2800
	        },
	        "id": "8da659ef-9e54-447d-9c36-84912db1848f"
	      },
	      "quantity": 2,
	      "discountedPricePerQuantity": [],
	      "state": [
	        {
	          "quantity": 2,
	          "state": {
	            "typeId": "state",
	            "id": "7c2e2694-aefe-43d7-888e-6a99514caaca"
	          }
	        }
	      ],
	      "priceMode": "Platform",
	      "lineItemMode": "Standard",
	      "totalPrice": {
	        "type": "centPrecision",
	        "fractionDigits": 2,
	        "currencyCode": "EUR",
	        "centAmount": 5600
	      }
	    }
	  ],
	  "store": {
	    "typeId": "store",
	    "key": "test-key"
	  },
	  "cartState": "Active",
	  "totalPrice": {
	    "type": "centPrecision",
	    "fractionDigits": 2,
	    "currencyCode": "EUR",
	    "centAmount": 5600
	  },
	  "customLineItems": [],
	  "discountCodes": [],
	  "inventoryMode": "None",
	  "taxMode": "Platform",
	  "taxRoundingMode": "HalfEven",
	  "taxCalculationMode": "LineItemLevel",
	  "refusedGifts": [],
	  "origin": "Customer"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	cart, err := client.CartGetWithKey(context.TODO(), "dummy-id")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, cart)
	assert.NotNil(t, cart.Version)
	assert.NotEmpty(t, cart.TaxRoundingMode)
	assert.NotEmpty(t, cart.TaxMode)
	assert.NotEmpty(t, cart.TaxCalculationMode)
	assert.NotEmpty(t, cart.Origin)
	assert.NotEmpty(t, cart.LastModifiedAt)
	assert.NotEmpty(t, cart.Key)
	assert.NotEmpty(t, cart.InventoryMode)
	assert.NotEmpty(t, cart.ID)
	assert.NotEmpty(t, cart.CreatedAt)
	assert.NotEmpty(t, cart.CartState)

}

func TestGeneratedCartDeleteWithID(t *testing.T) {
	responseData := ` {
	  "type": "Cart",
	  "id": "2a3baa00-44fa-4ab8-bec7-933c31e18dcc",
	  "key": "test-key",
	  "version": 5,
	  "createdAt": "2015-09-22T15:36:17.510Z",
	  "lastModifiedAt": "2015-09-22T15:41:55.816Z",
	  "lineItems": [
	    {
	      "id": "b925a817-d5ad-48bb-a407-29ad8e0649b5",
	      "productId": "9f10dcfb-5cc9-4a18-843a-c07f7e22d01f",
	      "name": {
	        "en": "SAPPHIRE"
	      },
	      "productType": {
	        "typeId": "product-type",
	        "id": "2543e1d8-4915-4f72-a3c9-1df9b1b0082d",
	        "version": 8
	      },
	      "productSlug": {
	        "en": "sapphire1421832124423"
	      },
	      "variant": {
	        "id": 1,
	        "sku": "sku_SAPPHIRE_variant1_1421832124423",
	        "prices": [
	          {
	            "value": {
	              "type": "centPrecision",
	              "fractionDigits": 2,
	              "currencyCode": "EUR",
	              "centAmount": 2800
	            },
	            "id": "8da659ef-9e54-447d-9c36-84912db1848f"
	          }
	        ],
	        "images": [
	          {
	            "url": "https://www.commercetools.com/cli/data/252542005_1.jpg",
	            "dimensions": {
	              "w": 1400,
	              "h": 1400
	            }
	          }
	        ],
	        "attributes": [],
	        "assets": []
	      },
	      "price": {
	        "value": {
	          "type": "centPrecision",
	          "fractionDigits": 2,
	          "currencyCode": "EUR",
	          "centAmount": 2800
	        },
	        "id": "8da659ef-9e54-447d-9c36-84912db1848f"
	      },
	      "quantity": 2,
	      "discountedPricePerQuantity": [],
	      "state": [
	        {
	          "quantity": 2,
	          "state": {
	            "typeId": "state",
	            "id": "7c2e2694-aefe-43d7-888e-6a99514caaca"
	          }
	        }
	      ],
	      "priceMode": "Platform",
	      "lineItemMode": "Standard",
	      "totalPrice": {
	        "type": "centPrecision",
	        "fractionDigits": 2,
	        "currencyCode": "EUR",
	        "centAmount": 5600
	      }
	    }
	  ],
	  "store": {
	    "typeId": "store",
	    "key": "test-key"
	  },
	  "cartState": "Active",
	  "totalPrice": {
	    "type": "centPrecision",
	    "fractionDigits": 2,
	    "currencyCode": "EUR",
	    "centAmount": 5600
	  },
	  "customLineItems": [],
	  "discountCodes": [],
	  "inventoryMode": "None",
	  "taxMode": "Platform",
	  "taxRoundingMode": "HalfEven",
	  "taxCalculationMode": "LineItemLevel",
	  "refusedGifts": [],
	  "origin": "Customer"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	cart, err := client.CartDeleteWithID(context.TODO(), "dummy-id", 1, true)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, cart)
	assert.NotNil(t, cart.Version)
	assert.NotEmpty(t, cart.TaxRoundingMode)
	assert.NotEmpty(t, cart.TaxMode)
	assert.NotEmpty(t, cart.TaxCalculationMode)
	assert.NotEmpty(t, cart.Origin)
	assert.NotEmpty(t, cart.LastModifiedAt)
	assert.NotEmpty(t, cart.Key)
	assert.NotEmpty(t, cart.InventoryMode)
	assert.NotEmpty(t, cart.ID)
	assert.NotEmpty(t, cart.CreatedAt)
	assert.NotEmpty(t, cart.CartState)

}

func TestGeneratedCartDeleteWithKey(t *testing.T) {
	responseData := ` {
	  "type": "Cart",
	  "id": "2a3baa00-44fa-4ab8-bec7-933c31e18dcc",
	  "key": "test-key",
	  "version": 5,
	  "createdAt": "2015-09-22T15:36:17.510Z",
	  "lastModifiedAt": "2015-09-22T15:41:55.816Z",
	  "lineItems": [
	    {
	      "id": "b925a817-d5ad-48bb-a407-29ad8e0649b5",
	      "productId": "9f10dcfb-5cc9-4a18-843a-c07f7e22d01f",
	      "name": {
	        "en": "SAPPHIRE"
	      },
	      "productType": {
	        "typeId": "product-type",
	        "id": "2543e1d8-4915-4f72-a3c9-1df9b1b0082d",
	        "version": 8
	      },
	      "productSlug": {
	        "en": "sapphire1421832124423"
	      },
	      "variant": {
	        "id": 1,
	        "sku": "sku_SAPPHIRE_variant1_1421832124423",
	        "prices": [
	          {
	            "value": {
	              "type": "centPrecision",
	              "fractionDigits": 2,
	              "currencyCode": "EUR",
	              "centAmount": 2800
	            },
	            "id": "8da659ef-9e54-447d-9c36-84912db1848f"
	          }
	        ],
	        "images": [
	          {
	            "url": "https://www.commercetools.com/cli/data/252542005_1.jpg",
	            "dimensions": {
	              "w": 1400,
	              "h": 1400
	            }
	          }
	        ],
	        "attributes": [],
	        "assets": []
	      },
	      "price": {
	        "value": {
	          "type": "centPrecision",
	          "fractionDigits": 2,
	          "currencyCode": "EUR",
	          "centAmount": 2800
	        },
	        "id": "8da659ef-9e54-447d-9c36-84912db1848f"
	      },
	      "quantity": 2,
	      "discountedPricePerQuantity": [],
	      "state": [
	        {
	          "quantity": 2,
	          "state": {
	            "typeId": "state",
	            "id": "7c2e2694-aefe-43d7-888e-6a99514caaca"
	          }
	        }
	      ],
	      "priceMode": "Platform",
	      "lineItemMode": "Standard",
	      "totalPrice": {
	        "type": "centPrecision",
	        "fractionDigits": 2,
	        "currencyCode": "EUR",
	        "centAmount": 5600
	      }
	    }
	  ],
	  "store": {
	    "typeId": "store",
	    "key": "test-key"
	  },
	  "cartState": "Active",
	  "totalPrice": {
	    "type": "centPrecision",
	    "fractionDigits": 2,
	    "currencyCode": "EUR",
	    "centAmount": 5600
	  },
	  "customLineItems": [],
	  "discountCodes": [],
	  "inventoryMode": "None",
	  "taxMode": "Platform",
	  "taxRoundingMode": "HalfEven",
	  "taxCalculationMode": "LineItemLevel",
	  "refusedGifts": [],
	  "origin": "Customer"
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	cart, err := client.CartDeleteWithKey(context.TODO(), "dummy-id", 1, true)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, cart)
	assert.NotNil(t, cart.Version)
	assert.NotEmpty(t, cart.TaxRoundingMode)
	assert.NotEmpty(t, cart.TaxMode)
	assert.NotEmpty(t, cart.TaxCalculationMode)
	assert.NotEmpty(t, cart.Origin)
	assert.NotEmpty(t, cart.LastModifiedAt)
	assert.NotEmpty(t, cart.Key)
	assert.NotEmpty(t, cart.InventoryMode)
	assert.NotEmpty(t, cart.ID)
	assert.NotEmpty(t, cart.CreatedAt)
	assert.NotEmpty(t, cart.CartState)

}

func TestGeneratedCartQuery(t *testing.T) {
	responseData := ` {
	  "limit": 20,
	  "offset": 0,
	  "count": 2,
	  "total": 2,
	  "results": [
	    {
	      "type": "Cart",
	      "id": "2a3baa00-44fa-4ab8-bec7-933c31e18dcc",
	      "version": 5,
	      "createdAt": "2015-09-22T15:36:17.510Z",
	      "lastModifiedAt": "2015-09-22T15:41:55.816Z",
	      "lineItems": [
	        {
	          "id": "b925a817-d5ad-48bb-a407-29ad8e0649b5",
	          "productId": "9f10dcfb-5cc9-4a18-843a-c07f7e22d01f",
	          "name": {
	            "en": "SAPPHIRE"
	          },
	          "productType": {
	            "typeId": "product-type",
	            "id": "2543e1d8-4915-4f72-a3c9-1df9b1b0082d",
	            "version": 8
	          },
	          "productSlug": {
	            "en": "sapphire1421832124423"
	          },
	          "variant": {
	            "id": 1,
	            "sku": "sku_SAPPHIRE_variant1_1421832124423",
	            "prices": [
	              {
	                "value": {
	                  "type": "centPrecision",
	                  "fractionDigits": 2,
	                  "currencyCode": "EUR",
	                  "centAmount": 2800
	                },
	                "id": "8da659ef-9e54-447d-9c36-84912db1848f"
	              }
	            ],
	            "images": [
	              {
	                "url": "https://www.commercetools.com/cli/data/252542005_1.jpg",
	                "dimensions": {
	                  "w": 1400,
	                  "h": 1400
	                }
	              }
	            ],
	            "attributes": [],
	            "assets": []
	          },
	          "price": {
	            "value": {
	              "type": "centPrecision",
	              "fractionDigits": 2,
	              "currencyCode": "EUR",
	              "centAmount": 2800
	            },
	            "id": "8da659ef-9e54-447d-9c36-84912db1848f"
	          },
	          "quantity": 2,
	          "discountedPricePerQuantity": [],
	          "state": [
	            {
	              "quantity": 2,
	              "state": {
	                "typeId": "state",
	                "id": "7c2e2694-aefe-43d7-888e-6a99514caaca"
	              }
	            }
	          ],
	          "priceMode": "Platform",
	          "lineItemMode": "Standard",
	          "totalPrice": {
	            "type": "centPrecision",
	            "fractionDigits": 2,
	            "currencyCode": "EUR",
	            "centAmount": 5600
	          }
	        }
	      ],
	      "cartState": "Active",
	      "totalPrice": {
	        "type": "centPrecision",
	        "fractionDigits": 2,
	        "currencyCode": "EUR",
	        "centAmount": 5600
	      },
	      "customLineItems": [],
	      "discountCodes": [],
	      "inventoryMode": "None",
	      "taxMode": "Platform",
	      "taxRoundingMode": "HalfEven",
	      "taxCalculationMode": "LineItemLevel",
	      "refusedGifts": [],
	      "origin": "Customer"
	    },
	    {
	      "type": "Cart",
	      "id": "668e5783-73c8-4f2d-91f4-3c90b872c700",
	      "version": 3,
	      "createdAt": "2015-10-07T07:33:05.894Z",
	      "lastModifiedAt": "2015-10-07T07:33:06.070Z",
	      "lineItems": [
	        {
	          "id": "90dff06c-272e-47fa-b8de-923dce092474",
	          "productId": "7b1203f4-66c0-438c-9a30-f4fb6be79bdf",
	          "name": {
	            "de": "WB ATHLETIC PANZER",
	            "en": "WB ATHLETIC TANK"
	          },
	          "productType": {
	            "typeId": "product-type",
	            "id": "2543e1d8-4915-4f72-a3c9-1df9b1b0082d",
	            "version": 8
	          },
	          "productSlug": {
	            "en": "wb-athletic-tank1421832124574"
	          },
	          "variant": {
	            "id": 1,
	            "sku": "sku_WB_ATHLETIC_TANK_variant1_1421832124574",
	            "prices": [
	              {
	                "value": {
	                  "type": "centPrecision",
	                  "fractionDigits": 2,
	                  "currencyCode": "EUR",
	                  "centAmount": 8400
	                },
	                "id": "37696f7c-8260-4941-a921-68e6aa76b4a3"
	              }
	            ],
	            "images": [
	              {
	                "url": "https://www.commercetools.com/cli/data/253265444_1.jpg",
	                "dimensions": {
	                  "w": 1400,
	                  "h": 1400
	                }
	              }
	            ],
	            "attributes": [],
	            "assets": []
	          },
	          "price": {
	            "value": {
	              "type": "centPrecision",
	              "fractionDigits": 2,
	              "currencyCode": "EUR",
	              "centAmount": 8400
	            },
	            "id": "37696f7c-8260-4941-a921-68e6aa76b4a3"
	          },
	          "quantity": 1,
	          "discountedPricePerQuantity": [],
	          "state": [
	            {
	              "quantity": 1,
	              "state": {
	                "typeId": "state",
	                "id": "7c2e2694-aefe-43d7-888e-6a99514caaca"
	              }
	            }
	          ],
	          "priceMode": "Platform",
	          "lineItemMode": "Standard",
	          "totalPrice": {
	            "type": "centPrecision",
	            "fractionDigits": 2,
	            "currencyCode": "EUR",
	            "centAmount": 8400
	          },
	          "custom": {
	            "type": {
	              "typeId": "type",
	              "id": "3ae9bcca-df23-443e-bd22-0c592f9694fa"
	            },
	            "fields": {
	              "offer_name": "SuperMax"
	            }
	          }
	        }
	      ],
	      "cartState": "Active",
	      "totalPrice": {
	        "type": "centPrecision",
	        "fractionDigits": 2,
	        "currencyCode": "EUR",
	        "centAmount": 8400
	      },
	      "country": "DE",
	      "customLineItems": [],
	      "discountCodes": [],
	      "inventoryMode": "None",
	      "taxMode": "Platform",
	      "taxRoundingMode": "HalfEven",
	      "taxCalculationMode": "LineItemLevel",
	      "refusedGifts": [],
	      "origin": "Customer"
	    }
	  ]
	} `
	client, server := testutil.MockClient(t, responseData, nil, nil)
	defer server.Close()
	input := commercetools.QueryInput{}
	queryResult, err := client.CartQuery(context.TODO(), &input)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, queryResult)
	assert.NotNil(t, queryResult.Total)
	assert.NotNil(t, queryResult.Offset)
	assert.NotNil(t, queryResult.Limit)
	assert.NotNil(t, queryResult.Count)

}
