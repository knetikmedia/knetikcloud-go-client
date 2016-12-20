# \InvoicesApi

All URIs are relative to *https://localhost:8080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInvoiceUsingPOST**](InvoicesApi.md#CreateInvoiceUsingPOST) | **Post** /invoices | Create an invoice
[**GetInvoiceHistoryUsingGET**](InvoicesApi.md#GetInvoiceHistoryUsingGET) | **Get** /invoices | Retrieve invoices
[**GetInvoiceUsingGET**](InvoicesApi.md#GetInvoiceUsingGET) | **Get** /invoices/{id} | Retrieve an invoice
[**GetLogsUsingGET**](InvoicesApi.md#GetLogsUsingGET) | **Get** /invoices/{id}/logs | List invoice logs
[**ListFulFillmentStatusesUsingGET**](InvoicesApi.md#ListFulFillmentStatusesUsingGET) | **Get** /invoices/fulfillment-statuses | Lists available fulfillment statuses
[**ListPaymentStatusesUsingGET**](InvoicesApi.md#ListPaymentStatusesUsingGET) | **Get** /invoices/payment-statuses | Lists available payment statuses
[**PayInvoiceUsingPOST**](InvoicesApi.md#PayInvoiceUsingPOST) | **Post** /invoices/{id}/payments | Trigger payment of an invoice
[**SetItemFulfillmentStatusUsingPUT**](InvoicesApi.md#SetItemFulfillmentStatusUsingPUT) | **Put** /invoices/{id}/items/{sku}/fulfillment-status | Set the fulfillment status of an invoice item
[**SetOrderNotesUsingPUT**](InvoicesApi.md#SetOrderNotesUsingPUT) | **Put** /invoices/{id}/order-notes | Set the order notes of an invoice
[**SetPaymentStatusUsingPUT**](InvoicesApi.md#SetPaymentStatusUsingPUT) | **Put** /invoices/{id}/payment-status | Set the payment status of an invoice
[**UpdateBillingInfoUsingPUT**](InvoicesApi.md#UpdateBillingInfoUsingPUT) | **Put** /invoices/{id}/billing-address | Set or update billing info


# **CreateInvoiceUsingPOST**
> []InvoiceResource CreateInvoiceUsingPOST($req)

Create an invoice

Create an invoice(s) by providing a cart GUID. Note that there may be multiple invoices created, one per vendor.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **req** | [**InvoiceCreateRequest**](InvoiceCreateRequest.md)| Invoice to be created | [optional] 

### Return type

[**[]InvoiceResource**](InvoiceResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInvoiceHistoryUsingGET**
> PageInvoiceResource GetInvoiceHistoryUsingGET($filterUser, $filterEmail, $filterFulfillmentStatus, $filterPaymentStatus, $filterItemName, $filterCreatedDate, $size, $page, $order)

Retrieve invoices

Without INVOICES_ADMIN permission the results are automatically filtered for only the logged in user's invoices. It is recomended however that filter_user be added to avoid issues for admin users accidentally getting additional invoices.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterUser** | **int32**| The id of a user to get invoices for. Automtically added if not being called with admin permissions. | [optional] 
 **filterEmail** | **string**| Filters invoices by customer&#39;s email. Admins only. | [optional] 
 **filterFulfillmentStatus** | **string**| Filters invoices by fulfillment status type. Can be a comma separated list of statuses | [optional] 
 **filterPaymentStatus** | **string**| Filters invoices by payment status type. Can be a comma separated list of statuses | [optional] 
 **filterItemName** | **string**| Filters invoices by item name containing the given string | [optional] 
 **filterCreatedDate** | **string**| Filters invoices by creation date. Multiple values possible for range search. Format: filter_created_date&#x3D;OP,ts&amp;... where OP in (GT, LT, GOE, LOE, EQ) and ts is a unix timestamp in seconds. Ex: filter_created_date&#x3D;GT,1452154258,LT,1554254874 | [optional] 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to id:ASC]

### Return type

[**PageInvoiceResource**](Page«InvoiceResource».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInvoiceUsingGET**
> InvoiceResource GetInvoiceUsingGET($id, $postalCode)

Retrieve an invoice

The postal code associated with the invoice may be required for security purposes if the invoice has one set and the config postal_code_required is set to true. Send 'none' if the invoice has no postal code.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| The id of the invoice | 
 **postalCode** | **string**| The postal code of the invoice or &#39;none&#39;. | [optional] 

### Return type

[**InvoiceResource**](InvoiceResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLogsUsingGET**
> PageInvoiceLogEntry GetLogsUsingGET($id, $size, $page)

List invoice logs


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| The id of the invoice | 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]

### Return type

[**PageInvoiceLogEntry**](Page«InvoiceLogEntry».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFulFillmentStatusesUsingGET**
> []string ListFulFillmentStatusesUsingGET()

Lists available fulfillment statuses


### Parameters
This endpoint does not need any parameter.

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPaymentStatusesUsingGET**
> []string ListPaymentStatusesUsingGET()

Lists available payment statuses


### Parameters
This endpoint does not need any parameter.

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PayInvoiceUsingPOST**
> PayInvoiceUsingPOST($id, $request)

Trigger payment of an invoice


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| The id of the invoice | 
 **request** | [**PayBySavedMethodRequest**](PayBySavedMethodRequest.md)| Payment info | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetItemFulfillmentStatusUsingPUT**
> SetItemFulfillmentStatusUsingPUT($id, $sku, $status)

Set the fulfillment status of an invoice item

This allows external fulfillment systems to report success or failure. Fulfillment status changes are restricted by a specific flow determining which status can lead to which.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| The id of the invoice | 
 **sku** | **string**| The sku of an item in the invoice | 
 **status** | **string**| The new fulfillment status for the item. Additional options may be available based on configuration.  Allowable values:  &#39;unfulfilled&#39;, &#39;fulfilled&#39;, &#39;not fulfillable&#39;, &#39;failed&#39;, &#39;processing&#39;, &#39;failed_permanent&#39;, &#39;delayed&#39; | 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetOrderNotesUsingPUT**
> SetOrderNotesUsingPUT($id, $orderNotes)

Set the order notes of an invoice


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| The id of the invoice | 
 **orderNotes** | **string**| Payment status info | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetPaymentStatusUsingPUT**
> SetPaymentStatusUsingPUT($id, $request)

Set the payment status of an invoice

This may trigger fulfillment if setting the status to 'paid'. This is mainly intended to support external payment systems that cannot be incorporated into the payment method system. Payment status changes are restricted by a specific flow determining which status can lead to which.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| The id of the invoice | 
 **request** | [**InvoicePaymentStatusRequest**](InvoicePaymentStatusRequest.md)| Payment status info | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBillingInfoUsingPUT**
> UpdateBillingInfoUsingPUT($id, $billingInfoRequest)

Set or update billing info


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| The id of the invoice | 
 **billingInfoRequest** | [**AddressResource**](AddressResource.md)| Address info | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

