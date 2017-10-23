# PaymentMethodTypeResource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The id of the payment method type | [default to null]
**InvoiceProcessingHours** | **int32** | The maximum timelimit in hours for an invoice in the processing status while waiting on this payment method type. Defaults to the global config invoice.processing_expiration_hours if null | [optional] [default to null]
**Name** | **string** | The name of the payment method type | [default to null]
**SupportsCapture** | **bool** | Whether the payment handler supports the authorize and capture flow | [optional] [default to null]
**SupportsPartial** | **bool** | Whether the payment handler supports paying for part of an invoice, rather than the full grand_total | [optional] [default to null]
**SupportsRebill** | **bool** | Whether the payment handler supports rebilling the method later (for saved payments or subscriptions) | [optional] [default to null]
**SupportsRefunds** | **bool** | Whether the payment handler supports refunding | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


