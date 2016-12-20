# \StoreCouponsApi

All URIs are relative to *https://localhost:8080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCouponItemUsingPOST**](StoreCouponsApi.md#CreateCouponItemUsingPOST) | **Post** /store/coupons | Create a coupon item
[**CreateCouponTemplateUsingPOST**](StoreCouponsApi.md#CreateCouponTemplateUsingPOST) | **Post** /store/coupons/templates | Create a coupon template
[**DeleteCouponItemUsingDELETE**](StoreCouponsApi.md#DeleteCouponItemUsingDELETE) | **Delete** /store/coupons/{id} | Delete a coupon item
[**DeleteCouponTemplateUsingDELETE**](StoreCouponsApi.md#DeleteCouponTemplateUsingDELETE) | **Delete** /store/coupons/templates/{id} | Delete a coupon template
[**GetCouponItemUsingGET**](StoreCouponsApi.md#GetCouponItemUsingGET) | **Get** /store/coupons/{id} | Get a single coupon item
[**GetCouponTemplateUsingGET**](StoreCouponsApi.md#GetCouponTemplateUsingGET) | **Get** /store/coupons/templates/{id} | Get a single coupon template
[**GetCouponTemplatesUsingGET**](StoreCouponsApi.md#GetCouponTemplatesUsingGET) | **Get** /store/coupons/templates | List and search coupon templates
[**UpdateCouponItemUsingPUT**](StoreCouponsApi.md#UpdateCouponItemUsingPUT) | **Put** /store/coupons/{id} | Update a coupon item
[**UpdateCouponTemplateUsingPUT**](StoreCouponsApi.md#UpdateCouponTemplateUsingPUT) | **Put** /store/coupons/templates/{id} | Update a coupon template


# **CreateCouponItemUsingPOST**
> CouponItem CreateCouponItemUsingPOST($couponItem)

Create a coupon item

SKUs have to be unique in the entire store.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **couponItem** | [**CouponItem**](CouponItem.md)| The coupon item object | [optional] 

### Return type

[**CouponItem**](CouponItem.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCouponTemplateUsingPOST**
> ItemTemplateResource CreateCouponTemplateUsingPOST($couponTemplateResource)

Create a coupon template

Coupon Templates define a type of coupon and the properties they have.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **couponTemplateResource** | [**ItemTemplateResource**](ItemTemplateResource.md)| The new coupon template | [optional] 

### Return type

[**ItemTemplateResource**](ItemTemplateResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCouponItemUsingDELETE**
> DeleteCouponItemUsingDELETE($id)

Delete a coupon item


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| The id of the coupon | 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCouponTemplateUsingDELETE**
> DeleteCouponTemplateUsingDELETE($id, $cascade)

Delete a coupon template


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 
 **cascade** | **string**| force deleting the template if it&#39;s attached to other objects, cascade &#x3D; detach | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCouponItemUsingGET**
> CouponItem GetCouponItemUsingGET($id)

Get a single coupon item


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| The id of the coupon | 

### Return type

[**CouponItem**](CouponItem.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCouponTemplateUsingGET**
> ItemTemplateResource GetCouponTemplateUsingGET($id)

Get a single coupon template

Coupon Templates define a type of coupon and the properties they have.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 

### Return type

[**ItemTemplateResource**](ItemTemplateResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCouponTemplatesUsingGET**
> PageResourceItemTemplateResource GetCouponTemplatesUsingGET($size, $page, $order)

List and search coupon templates


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to id:ASC]

### Return type

[**PageResourceItemTemplateResource**](PageResource«ItemTemplateResource».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCouponItemUsingPUT**
> UpdateCouponItemUsingPUT($id, $couponItem)

Update a coupon item


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| The id of the coupon | 
 **couponItem** | [**CouponItem**](CouponItem.md)| The coupon item object | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCouponTemplateUsingPUT**
> UpdateCouponTemplateUsingPUT($id, $couponTemplateResource)

Update a coupon template


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 
 **couponTemplateResource** | [**ItemTemplateResource**](ItemTemplateResource.md)| The coupon template resource object | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

