# \Templates_PropertiesApi

All URIs are relative to *https://jsapi-integration.us-east-1.elasticbeanstalk.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTemplatePropertyType**](Templates_PropertiesApi.md#GetTemplatePropertyType) | **Get** /templates/properties/{type} | Get details for a template property type
[**GetTemplatePropertyTypes**](Templates_PropertiesApi.md#GetTemplatePropertyTypes) | **Get** /templates/properties | List template property types


# **GetTemplatePropertyType**
> PropertyFieldListResource GetTemplatePropertyType(ctx, ctx, type_)
Get details for a template property type

<b>Permissions Needed:</b> ANY

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **type_** | **string**| type | 

### Return type

[**PropertyFieldListResource**](PropertyFieldListResource.md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTemplatePropertyTypes**
> []PropertyFieldListResource GetTemplatePropertyTypes(ctx, ctx, )
List template property types

<b>Permissions Needed:</b> ANY

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]PropertyFieldListResource**](PropertyFieldListResource.md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

