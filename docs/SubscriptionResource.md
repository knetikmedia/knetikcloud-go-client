# SubscriptionResource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalProperties** | [**map[string]Property**](Property.md) | A map of item additional properties, keyed on the property name. Must match the names and types defined in the template for this item type. | [optional] [default to null]
**Availability** | **string** | Who can purchase this subscription | [optional] [default to null]
**Behaviors** | [**[]Behavior**](Behavior.md) | The behaviors linked to the item, describing various options and interactions. May not be included in item lists | [optional] [default to null]
**Category** | **string** | A category for filtering items | [optional] [default to null]
**ConsolidationDayOfMonth** | **int32** | The day of the month 1..31 this subscription will renew | [optional] [default to null]
**CreatedDate** | **int64** | The date the item was created, unix timestamp in seconds | [optional] [default to null]
**GeoCountryList** | **[]string** | A list of country iso3 codes to include in the blacklist/whitelist geo policy | [optional] [default to null]
**GeoPolicyType** | **string** | Whether to use the geo_country_list as a black list or white list for item geographical availability | [optional] [default to null]
**Id** | **int32** | The id of the item | [optional] [default to null]
**LongDescription** | **string** | A long description of the subscription | [optional] [default to null]
**Name** | **string** | The name of the item | [default to null]
**Plans** | [**[]SubscriptionPlanResource**](SubscriptionPlanResource.md) | The billing options for this subscription | [optional] [default to null]
**ShortDescription** | **string** | A short description of the subscription.  Max 255 characters | [optional] [default to null]
**Sort** | **int32** | A number to use in sorting items.  Default 500 | [optional] [default to null]
**StoreEnd** | **int64** | Used to schedule removal from store.  Null means the subscription will never be removed | [optional] [default to null]
**StoreStart** | **int64** | Used to schedule appearance in store.  Null means the subscription will appear now | [optional] [default to null]
**Tags** | **[]string** | List of tags used for filtering items | [optional] [default to null]
**Template** | **string** | An item template this item is validated against. May be null and no validation of additional properties will be done. | [optional] [default to null]
**UniqueKey** | **string** | The unique key for the item | [optional] [default to null]
**UpdatedDate** | **int64** | The date the item was last updated | [optional] [default to null]
**VendorId** | **int32** | The vendor who provides the item | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


