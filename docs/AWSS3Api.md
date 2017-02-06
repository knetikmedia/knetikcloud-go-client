# \AWSS3Api

All URIs are relative to *https://localhost:8080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSignedS3URL**](AWSS3Api.md#GetSignedS3URL) | **Get** /amazon/s3/signedposturl | Get a signed S3 URL


# **GetSignedS3URL**
> AmazonS3Activity GetSignedS3URL($filename, $contentType)

Get a signed S3 URL

Requires the file name and file content type (i.e., 'video/mpeg')


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filename** | **string**| The file name | [optional] 
 **contentType** | **string**| The content type | [optional] 

### Return type

[**AmazonS3Activity**](AmazonS3Activity.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

