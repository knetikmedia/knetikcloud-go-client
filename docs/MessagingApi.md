# \MessagingApi

All URIs are relative to *https://localhost:8080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SendRawEmailUsingPOST**](MessagingApi.md#SendRawEmailUsingPOST) | **Post** /messaging/raw-email | Send a raw email to one or more users
[**SendRawSMSUsingPOST**](MessagingApi.md#SendRawSMSUsingPOST) | **Post** /messaging/raw-sms | Send a raw SMS
[**SendRawSMSUsingPOST1**](MessagingApi.md#SendRawSMSUsingPOST1) | **Post** /messaging/templated-sms | Send a new templated SMS
[**SendTemplateEmailUsingPOST**](MessagingApi.md#SendTemplateEmailUsingPOST) | **Post** /messaging/templated-email | Send a templated email to one or more users


# **SendRawEmailUsingPOST**
> SendRawEmailUsingPOST($rawEmailResource)

Send a raw email to one or more users


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rawEmailResource** | [**RawEmailResource**](RawEmailResource.md)| The new raw email to be sent | [optional] 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendRawSMSUsingPOST**
> SendRawSMSUsingPOST($rawSMSResource)

Send a raw SMS

Sends a raw SMS text message to one or more users. User's without registered mobile numbers will be skipped.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rawSMSResource** | [**RawSmsResource**](RawSmsResource.md)| The new raw SMS to be sent | [optional] 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendRawSMSUsingPOST1**
> SendRawSMSUsingPOST1($templateSMSResource)

Send a new templated SMS

Sends a templated SMS text message to one or more users. User's without registered mobile numbers will be skipped.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **templateSMSResource** | [**TemplateSmsResource**](TemplateSmsResource.md)| The new template SMS to be sent | [optional] 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendTemplateEmailUsingPOST**
> SendTemplateEmailUsingPOST($messageResource)

Send a templated email to one or more users


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **messageResource** | [**TemplateEmailResource**](TemplateEmailResource.md)| The new template email to be sent | [optional] 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

