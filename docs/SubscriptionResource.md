# SubscriptionResource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalProperties** | [**map[string]Property**](Property.md) | The additional properties for the subscription | [optional] [default to null]
**Availability** | **string** | Who can purchase this subscription | [optional] [default to null]
**Behaviors** | [**[]Behavior**](Behavior.md) | The behaviors linked to the item, describing various options and interactions. May not be included in item lists | [optional] [default to null]
**Category** | **string** | The category of the subscription | [optional] [default to null]
**ConsolidationDayOfMonth** | **int32** | The day of the month 1..31 this subscription will renew | [optional] [default to null]
**CreatedDate** | **int64** | The date the item was created, unix timestamp in seconds | [optional] [default to null]
**Displayable** | **bool** | Whether or not the item is currently visible to users. Does not block purchase; Use store_start or store_end to block purchase.  Default &#x3D; true | [optional] [default to null]
**GeoCountryList** | **[]string** | The geo country list for the subscription | [optional] [default to null]
**GeoPolicyType** | **string** | The geo policy type for the subscription | [optional] [default to null]
**Id** | **int32** | The id of the item | [optional] [default to null]
**LongDescription** | **string** | A long description of the subscription | [optional] [default to null]
**Name** | **string** | The name of the item | [default to null]
**Plans** | [**[]SubscriptionPlanResource**](SubscriptionPlanResource.md) | The billing options for this subscription | [optional] [default to null]
**ShortDescription** | **string** | A short description of the subscription.  Max 255 characters | [optional] [default to null]
**Sort** | **int32** | A number to use in sorting items.  Default 500 | [optional] [default to null]
**StoreEnd** | **int64** | Used to schedule removal from store.  Null means the subscription will never be removed | [optional] [default to null]
**StoreStart** | **int64** | Used to schedule appearance in store.  Null means the subscription will appear now | [optional] [default to null]
**Tags** | **[]string** | The tags for the subscription | [optional] [default to null]
**Template** | **string** | The template being used | [optional] [default to null]
**UniqueKey** | **string** | The unique key of the subscription | [optional] [default to null]
**UpdatedDate** | **int64** | The date the item was last updated | [optional] [default to null]
**VendorId** | **int32** | The id of the vendor | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


