# InventorySubscriptionResource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillDate** | **int64** | The date the subscription will be billed | [optional] [default to null]
**Credit** | **float32** | A credit of money already applied to a subscription for the next bill, or a debt if negative | [optional] [default to null]
**CreditLog** | [**[]SubscriptionCreditResource**](SubscriptionCreditResource.md) | A record of past and present credit/debt changes | [optional] [default to null]
**GraceEnd** | **int64** | The date the grace period ends | [optional] [default to null]
**InventoryId** | **int32** | The id of the inventory | [optional] [default to null]
**InventoryStatus** | **string** | The inventory status object | [optional] [default to null]
**ItemId** | **int32** | The id of the item | [optional] [default to null]
**PaymentMethod** | [***PaymentMethodResource**](PaymentMethodResource.md) | The payment method object | [optional] [default to null]
**PriceOverride** | **float32** | The recurring price that has been set to override the base price. Null if not overriding | [optional] [default to null]
**PriceOverrideReason** | **string** | An explanation for the reason the price is being overridden | [optional] [default to null]
**RecurringPrice** | **float32** | The default recurring price | [optional] [default to null]
**Sku** | **string** | The recurring sku of the subscription | [optional] [default to null]
**StartDate** | **int64** | The date the subscription will start | [optional] [default to null]
**SubscriptionStatus** | **int32** | The status of the subscription | [optional] [default to null]
**User** | [***SimpleUserResource**](SimpleUserResource.md) | The user | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


