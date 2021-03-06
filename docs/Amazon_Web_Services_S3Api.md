# \Amazon_Web_Services_S3Api

All URIs are relative to *https://jsapi-integration.us-east-1.elasticbeanstalk.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDownloadURL**](Amazon_Web_Services_S3Api.md#GetDownloadURL) | **Get** /amazon/s3/downloadurl | Get a temporary signed S3 URL for download
[**GetSignedS3URL**](Amazon_Web_Services_S3Api.md#GetSignedS3URL) | **Get** /amazon/s3/signedposturl | Get a signed S3 URL for upload


# **GetDownloadURL**
> string GetDownloadURL(ctx, ctx, optional)
Get a temporary signed S3 URL for download

To give access to files in your own S3 account, you will need to grant KnetikcCloud access to the file by adjusting your bucket policy accordingly. See S3 documentation for details. <br><br><b>Permissions Needed:</b> S3_ADMIN

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
 **bucket** | **string**| S3 bucket name | 
 **path** | **string**| The path to the file relative to the bucket (the s3 object key) | 
 **expiration** | **int32**| The number of seconds this URL will be valid. Default to 60 | [default to 60]

### Return type

**string**

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSignedS3URL**
> AmazonS3Activity GetSignedS3URL(ctx, ctx, optional)
Get a signed S3 URL for upload

Requires the file name and file content type (i.e., 'video/mpeg'). Make a PUT to the resulting url to upload the file and use the cdn_url to retrieve it after. <br><br><b>Permissions Needed:</b> S3_USER or S3_ADMIN

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
 **filename** | **string**| The file name | 
 **contentType** | **string**| The content type | 

### Return type

[**AmazonS3Activity**](AmazonS3Activity.md)

### Authorization

[oauth2_client_credentials_grant](../README.md#oauth2_client_credentials_grant), [oauth2_password_grant](../README.md#oauth2_password_grant)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

