# \StoreShoppingCartsApi

All URIs are relative to *https://localhost:8080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDiscountUsingPOST**](StoreShoppingCartsApi.md#AddDiscountUsingPOST) | **Post** /carts/{id}/discounts | Adds a coupon to the cart
[**AddItemUsingPOST**](StoreShoppingCartsApi.md#AddItemUsingPOST) | **Post** /carts/{id}/items | Add an item to the cart
[**AssignCartUsingPUT**](StoreShoppingCartsApi.md#AssignCartUsingPUT) | **Put** /carts/{id}/owner | Sets the owner of a cart if none is set already
[**CheckShippableUsingGET**](StoreShoppingCartsApi.md#CheckShippableUsingGET) | **Get** /carts/{id}/shippable | Returns whether a cart requires shipping
[**CreateCartUsingPOST**](StoreShoppingCartsApi.md#CreateCartUsingPOST) | **Post** /carts | Creates a new cart from scratch
[**GetCartUsingGET**](StoreShoppingCartsApi.md#GetCartUsingGET) | **Get** /carts/{id} | Returns the cart with the given GUID
[**GetCountriesUsingGET**](StoreShoppingCartsApi.md#GetCountriesUsingGET) | **Get** /carts/{id}/countries | Get the list of available shipping countries per vendor
[**ModifyShippingAddressUsingPUT**](StoreShoppingCartsApi.md#ModifyShippingAddressUsingPUT) | **Put** /carts/{id}/shipping-address | Modifies or sets the order shipping address
[**RemoveDiscountUsingDELETE**](StoreShoppingCartsApi.md#RemoveDiscountUsingDELETE) | **Delete** /carts/{id}/discounts/{code} | Removes a coupon from the cart
[**SearchCartsUsingGET**](StoreShoppingCartsApi.md#SearchCartsUsingGET) | **Get** /carts | Get a list of carts
[**SetCartCurrencyUsingPUT**](StoreShoppingCartsApi.md#SetCartCurrencyUsingPUT) | **Put** /carts/{id}/currency | Sets the currency to use for the cart
[**UpdateItemUsingPUT**](StoreShoppingCartsApi.md#UpdateItemUsingPUT) | **Put** /carts/{id}/items | Changes the quantity of an item already in the cart


# **AddDiscountUsingPOST**
> AddDiscountUsingPOST($id, $skuRequest)

Adds a coupon to the cart


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the cart | 
 **skuRequest** | [**SkuRequest**](SkuRequest.md)| The request of the sku | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddItemUsingPOST**
> AddItemUsingPOST($id, $cartItemRequest)

Add an item to the cart

Currently, carts cannot contain virtual and real currency items at the same time. Furthermore, the API only support a single virtual item at the moment


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the cart | 
 **cartItemRequest** | [**CartItemRequest**](CartItemRequest.md)| The cart item request object | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AssignCartUsingPUT**
> AssignCartUsingPUT($id, $userId)

Sets the owner of a cart if none is set already


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the cart | 
 **userId** | **int32**| The id of the user | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CheckShippableUsingGET**
> CartShippableResponse CheckShippableUsingGET($id)

Returns whether a cart requires shipping


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the cart | 

### Return type

[**CartShippableResponse**](CartShippableResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCartUsingPOST**
> string CreateCartUsingPOST($owner, $currencyCode)

Creates a new cart from scratch

You don't have to have a user to create a cart but the API requires authentication to checkout


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **int32**| Set the owner of a cart. If not specified, defaults to the calling user&#39;s id. If specified and is not the calling user&#39;s id, SHOPPING_CARTS_ADMIN permission is required | [optional] 
 **currencyCode** | **string**| Set the currency for the cart, by currency code. May be disallowed by site settings. | [optional] 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCartUsingGET**
> Cart GetCartUsingGET($id)

Returns the cart with the given GUID


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the cart | 

### Return type

[**Cart**](Cart.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCountriesUsingGET**
> SampleCountriesResponse GetCountriesUsingGET($id)

Get the list of available shipping countries per vendor

Since a cart can have multiple vendors with different shipping options, the countries are broken down by vendors. Please see notes about the response object as the fields are variable.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the cart | 

### Return type

[**SampleCountriesResponse**](SampleCountriesResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ModifyShippingAddressUsingPUT**
> ModifyShippingAddressUsingPUT($id, $cartShippingAddressRequest)

Modifies or sets the order shipping address


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the cart | 
 **cartShippingAddressRequest** | [**CartShippingAddressRequest**](CartShippingAddressRequest.md)| The cart shipping address request object | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveDiscountUsingDELETE**
> RemoveDiscountUsingDELETE($id, $code)

Removes a coupon from the cart


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the cart | 
 **code** | **string**| The SKU code of the coupon to remove | 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchCartsUsingGET**
> PageCartSummary SearchCartsUsingGET($filterOwnerId, $size, $page, $order)

Get a list of carts


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterOwnerId** | **int32**| Filter by the id of the owner | [optional] 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to id:ASC]

### Return type

[**PageCartSummary**](Page«CartSummary».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetCartCurrencyUsingPUT**
> SetCartCurrencyUsingPUT($id, $currencyCode)

Sets the currency to use for the cart

May be disallowed by site settings.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the cart | 
 **currencyCode** | **string**| The code of the currency | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateItemUsingPUT**
> UpdateItemUsingPUT($id, $cartItemRequest)

Changes the quantity of an item already in the cart

A quantity of zero will remove the item from the cart altogether.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the cart | 
 **cartItemRequest** | [**CartItemRequest**](CartItemRequest.md)| The cart item request object | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

