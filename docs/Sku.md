# Sku

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalProperties** | [**map[string]Property**](Property.md) | A map of additional properties, keyed on the property name.  Must match the names and types defined in the template for this item type, or be an extra not from the template | [optional] [default to null]
**CurrencyCode** | **string** | The currency code for the SKU, a three letter string (ISO3) | [default to null]
**Description** | **string** | The friendly name of the SKU as it will appear on invoices and reports. Typically represents the option name like red, large, etc | [default to null]
**Inventory** | **int32** | The number of SKUs currently in stock | [optional] [default to null]
**MinInventoryThreshold** | **int32** | Alerts vendor when SKU inventory drops below this value | [optional] [default to null]
**NotAvailable** | **bool** |  | [optional] [default to null]
**NotDisplayable** | **bool** |  | [optional] [default to null]
**OriginalPrice** | **float32** | The base price before any sale | [default to null]
**Price** | **float32** | The current price of the SKU with sales, if any. Set original_price for the base | [optional] [default to null]
**Published** | **bool** | Whether or not the SKU is currently visible to users. This will not block users from purchase. Use start_date or stop_date to prevent purchase. Default: true | [optional] [default to null]
**SaleId** | **int32** | The id of a sale affecting the price, if any | [optional] [default to null]
**SaleName** | **string** | The name of a sale affecting the price, if any | [optional] [default to null]
**Sku** | **string** | The stock keeping unit (SKU), a unique identifier for a given product.  Max 40 characters | [default to null]
**StartDate** | **int64** | The date the sku becomes visible (if published) and available for purchase, unix timestamp in seconds.  If set to null, sku will become available immediately | [optional] [default to null]
**StopDate** | **int64** | The date the sku becomes hidden and unavailable for purchase, unix timestamp in seconds.  If set to null, sku is always available | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


