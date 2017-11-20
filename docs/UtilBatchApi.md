# \UtilBatchApi

All URIs are relative to *https://sandbox.knetikcloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBatch**](UtilBatchApi.md#GetBatch) | **Get** /batch/{token} | Get batch result with token
[**SendBatch**](UtilBatchApi.md#SendBatch) | **Post** /batch | Request to run API call given the method, content type, path url, and body of request


# **GetBatch**
> []BatchReturn GetBatch(ctx, ctx, token)
Get batch result with token

Tokens expire in 24 hours

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **token** | **string**| token | 

### Return type

[**[]BatchReturn**](BatchReturn.md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendBatch**
> []BatchReturn SendBatch(ctx, ctx, optional)
Request to run API call given the method, content type, path url, and body of request

Should the request take longer than one of the alloted timeout parameters, a token will be returned instead, which can be used on the token endpoint in this service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batch** | [**Batch**](Batch.md)| The batch object | 

### Return type

[**[]BatchReturn**](BatchReturn.md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

