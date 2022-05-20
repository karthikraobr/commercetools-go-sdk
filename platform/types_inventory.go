package platform

// Generated file, please do not change!!!

import (
	"encoding/json"
	"errors"
	"time"
)

type InventoryEntry struct {
	// Platform-generated unique identifier of the InventoryEntry.
	ID             string    `json:"id"`
	Version        int       `json:"version"`
	CreatedAt      time.Time `json:"createdAt"`
	LastModifiedAt time.Time `json:"lastModifiedAt"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	LastModifiedBy *LastModifiedBy `json:"lastModifiedBy,omitempty"`
	// Present on resources created after 1 February 2019 except for [events not tracked](/client-logging#events-tracked).
	CreatedBy *CreatedBy `json:"createdBy,omitempty"`
	// User-defined unique identifier of the InventoryEntry.
	Key *string `json:"key,omitempty"`
	Sku string  `json:"sku"`
	// Connection to a particular supplier.
	SupplyChannel *ChannelReference `json:"supplyChannel,omitempty"`
	// Overall amount of stock.
	// (available + reserved)
	QuantityOnStock int `json:"quantityOnStock"`
	// Available amount of stock.
	// (available means: `quantityOnStock` - reserved quantity)
	AvailableQuantity int `json:"availableQuantity"`
	// The time period in days, that tells how often this inventory entry is restocked.
	RestockableInDays *int `json:"restockableInDays,omitempty"`
	// The date and time of the next restock.
	ExpectedDelivery *time.Time    `json:"expectedDelivery,omitempty"`
	Custom           *CustomFields `json:"custom,omitempty"`
}

type InventoryEntryDraft struct {
	Sku string `json:"sku"`
	// User-defined unique identifier for the InventoryEntry.
	Key               *string                    `json:"key,omitempty"`
	SupplyChannel     *ChannelResourceIdentifier `json:"supplyChannel,omitempty"`
	QuantityOnStock   int                        `json:"quantityOnStock"`
	RestockableInDays *int                       `json:"restockableInDays,omitempty"`
	ExpectedDelivery  *time.Time                 `json:"expectedDelivery,omitempty"`
	// The custom fields.
	Custom *CustomFieldsDraft `json:"custom,omitempty"`
}

/**
*	[Reference](ctp:api:type:Reference) to a [InventoryEntry](ctp:api:type:InventoryEntry).
*
 */
type InventoryEntryReference struct {
	// Platform-generated unique identifier of the referenced [InventoryEntry](ctp:api:type:InventoryEntry).
	ID string `json:"id"`
	// Contains the representation of the expanded InventoryEntry. Only present in responses to requests with [Reference Expansion](/../api/general-concepts#reference-expansion) for InventoryEntries.
	Obj *InventoryEntry `json:"obj,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InventoryEntryReference) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntryReference
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "inventory-entry", Alias: (*Alias)(&obj)})
}

/**
*	[ResourceIdentifier](ctp:api:type:ResourceIdentifier) to a [InventoryEntry](ctp:api:type:InventoryEntry).
*
 */
type InventoryEntryResourceIdentifier struct {
	// Platform-generated unique identifier of the referenced [InventoryEntry](ctp:api:type:InventoryEntry). Either `id` or `key` is required.
	ID *string `json:"id,omitempty"`
	// User-defined unique identifier of the referenced [InventoryEntry](ctp:api:type:InventoryEntry). Either `id` or `key` is required.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InventoryEntryResourceIdentifier) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntryResourceIdentifier
	return json.Marshal(struct {
		Action string `json:"typeId"`
		*Alias
	}{Action: "inventory-entry", Alias: (*Alias)(&obj)})
}

type InventoryEntryUpdate struct {
	Version int                          `json:"version"`
	Actions []InventoryEntryUpdateAction `json:"actions"`
}

// UnmarshalJSON override to deserialize correct attribute types based
// on the discriminator value
func (obj *InventoryEntryUpdate) UnmarshalJSON(data []byte) error {
	type Alias InventoryEntryUpdate
	if err := json.Unmarshal(data, (*Alias)(obj)); err != nil {
		return err
	}
	for i := range obj.Actions {
		var err error
		obj.Actions[i], err = mapDiscriminatorInventoryEntryUpdateAction(obj.Actions[i])
		if err != nil {
			return err
		}
	}
	return nil
}

type InventoryEntryUpdateAction interface{}

func mapDiscriminatorInventoryEntryUpdateAction(input interface{}) (InventoryEntryUpdateAction, error) {
	var discriminator string
	if data, ok := input.(map[string]interface{}); ok {
		discriminator, ok = data["action"].(string)
		if !ok {
			return nil, errors.New("error processing discriminator field 'action'")
		}
	} else {
		return nil, errors.New("invalid data")
	}

	switch discriminator {
	case "addQuantity":
		obj := InventoryEntryAddQuantityAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "changeQuantity":
		obj := InventoryEntryChangeQuantityAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "removeQuantity":
		obj := InventoryEntryRemoveQuantityAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomField":
		obj := InventoryEntrySetCustomFieldAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setCustomType":
		obj := InventoryEntrySetCustomTypeAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setExpectedDelivery":
		obj := InventoryEntrySetExpectedDeliveryAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setKey":
		obj := InventoryEntrySetKeyAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setRestockableInDays":
		obj := InventoryEntrySetRestockableInDaysAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	case "setSupplyChannel":
		obj := InventoryEntrySetSupplyChannelAction{}
		if err := decodeStruct(input, &obj); err != nil {
			return nil, err
		}
		return obj, nil
	}
	return nil, nil
}

type InventoryPagedQueryResponse struct {
	// Number of [results requested](/../api/general-concepts#limit).
	Limit int  `json:"limit"`
	Count int  `json:"count"`
	Total *int `json:"total,omitempty"`
	// Number of [elements skipped](/../api/general-concepts#offset).
	Offset  int              `json:"offset"`
	Results []InventoryEntry `json:"results"`
}

type InventoryEntryAddQuantityAction struct {
	Quantity int `json:"quantity"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InventoryEntryAddQuantityAction) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntryAddQuantityAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "addQuantity", Alias: (*Alias)(&obj)})
}

type InventoryEntryChangeQuantityAction struct {
	Quantity int `json:"quantity"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InventoryEntryChangeQuantityAction) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntryChangeQuantityAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "changeQuantity", Alias: (*Alias)(&obj)})
}

type InventoryEntryRemoveQuantityAction struct {
	Quantity int `json:"quantity"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InventoryEntryRemoveQuantityAction) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntryRemoveQuantityAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "removeQuantity", Alias: (*Alias)(&obj)})
}

type InventoryEntrySetCustomFieldAction struct {
	// Name of the [Custom Field](/../api/projects/custom-fields).
	Name string `json:"name"`
	// If `value` is absent or `null`, this field will be removed if it exists.
	// Trying to remove a field that does not exist will fail with an [InvalidOperation](/../api/errors#general-400-invalid-operation) error.
	// If `value` is provided, it is set for the field defined by `name`.
	Value interface{} `json:"value,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InventoryEntrySetCustomFieldAction) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntrySetCustomFieldAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomField", Alias: (*Alias)(&obj)})
}

type InventoryEntrySetCustomTypeAction struct {
	// Defines the [Type](ctp:api:type:Type) that extends the InventoryEntry with [Custom Fields](/../api/projects/custom-fields).
	// If absent, any existing Type and Custom Fields are removed from the InventoryEntry.
	Type *TypeResourceIdentifier `json:"type,omitempty"`
	// Sets the [Custom Fields](/../api/projects/custom-fields) fields for the InventoryEntry.
	Fields *FieldContainer `json:"fields,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InventoryEntrySetCustomTypeAction) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntrySetCustomTypeAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setCustomType", Alias: (*Alias)(&obj)})
}

type InventoryEntrySetExpectedDeliveryAction struct {
	ExpectedDelivery *time.Time `json:"expectedDelivery,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InventoryEntrySetExpectedDeliveryAction) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntrySetExpectedDeliveryAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setExpectedDelivery", Alias: (*Alias)(&obj)})
}

type InventoryEntrySetKeyAction struct {
	// Value to set. If empty, any existing value will be removed.
	Key *string `json:"key,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InventoryEntrySetKeyAction) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntrySetKeyAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setKey", Alias: (*Alias)(&obj)})
}

type InventoryEntrySetRestockableInDaysAction struct {
	RestockableInDays *int `json:"restockableInDays,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InventoryEntrySetRestockableInDaysAction) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntrySetRestockableInDaysAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setRestockableInDays", Alias: (*Alias)(&obj)})
}

type InventoryEntrySetSupplyChannelAction struct {
	// If absent, the supply channel is removed.
	// This action will fail if an entry with the combination of sku and supplyChannel already exists.
	SupplyChannel *ChannelResourceIdentifier `json:"supplyChannel,omitempty"`
}

// MarshalJSON override to set the discriminator value or remove
// optional nil slices
func (obj InventoryEntrySetSupplyChannelAction) MarshalJSON() ([]byte, error) {
	type Alias InventoryEntrySetSupplyChannelAction
	return json.Marshal(struct {
		Action string `json:"action"`
		*Alias
	}{Action: "setSupplyChannel", Alias: (*Alias)(&obj)})
}
