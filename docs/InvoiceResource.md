# InvoiceResource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingAddress1** | **string** | Line one of the customer&#39;s billing address | [optional] [default to null]
**BillingAddress2** | **string** | Line two of the customer&#39;s billing address | [optional] [default to null]
**BillingCityName** | **string** | The city for the customer&#39;s billing address | [optional] [default to null]
**BillingCountryName** | **string** | The country for the customer&#39;s billing address | [optional] [default to null]
**BillingFullName** | **string** | The customer&#39;s name for the billing address | [optional] [default to null]
**BillingPostalCode** | **string** | The postal code for the customer&#39;s billing address | [optional] [default to null]
**BillingStateName** | **string** | The state for the customer&#39;s billing address | [optional] [default to null]
**CartId** | **string** | The guid of the cart this invoice came from | [optional] [default to null]
**CreatedDate** | **int64** | The date the invoice was created, unix timestamp in seconds | [optional] [default to null]
**Currency** | **string** | The code for the currency invoice prices are in | [optional] [default to null]
**CurrentFulfillmentStatus** | **string** | The fulfillment status of the invoice | [optional] [default to null]
**CurrentPaymentStatus** | **string** | The payment status of the invoice | [optional] [default to null]
**Discount** | **float64** | The amount of money saved through coupons | [optional] [default to null]
**Email** | **string** | The customer&#39;s email address | [optional] [default to null]
**FedTax** | **float64** | The amount of federal tax added | [optional] [default to null]
**GrandTotal** | **float64** | The final price of the invoice | [optional] [default to null]
**Id** | **int32** | The id of the invoice | [optional] [default to null]
**InvoiceNumber** | **string** | A reference number for the invoice | [optional] [default to null]
**Items** | [**[]InvoiceItemResource**](InvoiceItemResource.md) | A list of items within the invoice | [optional] [default to null]
**NamePrefix** | **string** | The customer&#39;s name prefix | [optional] [default to null]
**OrderNotes** | **string** | Notes about the order | [optional] [default to null]
**ParentInvoiceId** | **int32** | The id of an invoice this is a child of | [optional] [default to null]
**PaymentMethodId** | **int32** | The id of a saved payment method used to pay for the invoice | [optional] [default to null]
**Phone** | **string** | The customer&#39;s phone number | [optional] [default to null]
**PhoneNumber** | **string** | The customer&#39;s phone number | [optional] [default to null]
**Shipping** | **float64** | The shipping cost | [optional] [default to null]
**ShippingAddress1** | **string** | Line one of the customer&#39;s shipping address | [optional] [default to null]
**ShippingAddress2** | **string** | Line two of the customer&#39;s shipping address | [optional] [default to null]
**ShippingCityName** | **string** | The city for the customer&#39;s shipping address | [optional] [default to null]
**ShippingCountryName** | **string** | The country for the customer&#39;s shipping address | [optional] [default to null]
**ShippingFullName** | **string** | The customer&#39;s name for the shipping address | [optional] [default to null]
**ShippingPostalCode** | **string** | The postal code for the customer&#39;s shipping address | [optional] [default to null]
**ShippingStateName** | **string** | The state for the customer&#39;s shipping address | [optional] [default to null]
**Sort** | **int32** | A number to use in sorting items. default 500. | [optional] [default to null]
**StateTax** | **float64** | The amount of state tax added | [optional] [default to null]
**Subtotal** | **float64** | The sum price of all items before shipping, coupons and tax | [optional] [default to null]
**UpdatedDate** | **int64** | The date the invoice was last updated, unix timestamp in seconds | [optional] [default to null]
**User** | [**SimpleUserResource**](SimpleUserResource.md) | The owner of the invoice | [optional] [default to null]
**VendorId** | **int32** | The id of the vendor | [optional] [default to null]
**VendorName** | **string** | The name of the invoice | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

