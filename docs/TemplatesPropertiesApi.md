# \TemplatesPropertiesApi

All URIs are relative to *https://sandbox.knetikcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTemplatePropertyType**](TemplatesPropertiesApi.md#GetTemplatePropertyType) | **Get** /templates/properties/{type} | Get details for a template property type
[**GetTemplatePropertyTypes**](TemplatesPropertiesApi.md#GetTemplatePropertyTypes) | **Get** /templates/properties | List template property types


# **GetTemplatePropertyType**
> PropertyFieldListResource GetTemplatePropertyType($type_)

Get details for a template property type


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string**| type | 

### Return type

[**PropertyFieldListResource**](PropertyFieldListResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTemplatePropertyTypes**
> []PropertyFieldListResource GetTemplatePropertyTypes()

List template property types


### Parameters
This endpoint does not need any parameter.

### Return type

[**[]PropertyFieldListResource**](PropertyFieldListResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

