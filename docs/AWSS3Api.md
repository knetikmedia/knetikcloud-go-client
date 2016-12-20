# \AWSS3Api

All URIs are relative to *https://localhost:8080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PreSignUrlUsingGET**](AWSS3Api.md#PreSignUrlUsingGET) | **Get** /amazon/s3/signedposturl | Get a signed S3 URL


# **PreSignUrlUsingGET**
> AmazonS3Activity PreSignUrlUsingGET($filename, $contentType)

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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

