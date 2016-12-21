# \ContentPollsApi

All URIs are relative to *https://localhost:8080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnswerPollUsingPOST**](ContentPollsApi.md#AnswerPollUsingPOST) | **Post** /media/polls/{id}/response | Add your vote to a poll
[**CreatePollTemplateUsingPOST**](ContentPollsApi.md#CreatePollTemplateUsingPOST) | **Post** /media/polls/templates | Create a poll template
[**CreatePollUsingPOST**](ContentPollsApi.md#CreatePollUsingPOST) | **Post** /media/polls | Create a new poll
[**DeletePollTemplateUsingDELETE**](ContentPollsApi.md#DeletePollTemplateUsingDELETE) | **Delete** /media/polls/templates/{id} | Delete a poll template
[**DeletePollUsingDELETE**](ContentPollsApi.md#DeletePollUsingDELETE) | **Delete** /media/polls/{id} | Delete an existing poll
[**GetPollAnswerUsingGET**](ContentPollsApi.md#GetPollAnswerUsingGET) | **Get** /media/polls/{id}/response | Get poll answer
[**GetPollTemplateUsingGET**](ContentPollsApi.md#GetPollTemplateUsingGET) | **Get** /media/polls/templates/{id} | Get a single poll template
[**GetPollTemplatesUsingGET**](ContentPollsApi.md#GetPollTemplatesUsingGET) | **Get** /media/polls/templates | List and search poll templates
[**GetPollUsingGET**](ContentPollsApi.md#GetPollUsingGET) | **Get** /media/polls/{id} | Get a single poll
[**GetPollsUsingGET**](ContentPollsApi.md#GetPollsUsingGET) | **Get** /media/polls | List and search polls
[**UpdatePollTemplateUsingPUT**](ContentPollsApi.md#UpdatePollTemplateUsingPUT) | **Put** /media/polls/templates/{id} | Update a poll template
[**UpdatePollUsingPUT**](ContentPollsApi.md#UpdatePollUsingPUT) | **Put** /media/polls/{id} | Update an existing poll


# **AnswerPollUsingPOST**
> PollResponseResource AnswerPollUsingPOST($id, $answerKey)

Add your vote to a poll


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The poll id | 
 **answerKey** | **string**| The answer key | [optional] 

### Return type

[**PollResponseResource**](PollResponseResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePollTemplateUsingPOST**
> TemplateResource CreatePollTemplateUsingPOST($pollTemplateResource)

Create a poll template

Poll templates define a type of poll and the properties they have


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pollTemplateResource** | [**TemplateResource**](TemplateResource.md)| The poll template resource object | [optional] 

### Return type

[**TemplateResource**](TemplateResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePollUsingPOST**
> PollResource CreatePollUsingPOST($pollResource)

Create a new poll

Polls are blobs of text with titles, a category and assets. Formatting and display of the text is in the hands of the front end.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pollResource** | [**PollResource**](PollResource.md)| The poll object | [optional] 

### Return type

[**PollResource**](PollResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePollTemplateUsingDELETE**
> DeletePollTemplateUsingDELETE($id, $cascade)

Delete a poll template

If cascade = 'detach', it will force delete the template even if it's attached to other objects


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 
 **cascade** | **string**| The value needed to delete used templates | [optional] 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePollUsingDELETE**
> DeletePollUsingDELETE($id)

Delete an existing poll


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The poll id | 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPollAnswerUsingGET**
> PollResponseResource GetPollAnswerUsingGET($id)

Get poll answer


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The poll id | 

### Return type

[**PollResponseResource**](PollResponseResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPollTemplateUsingGET**
> TemplateResource GetPollTemplateUsingGET($id)

Get a single poll template


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 

### Return type

[**TemplateResource**](TemplateResource.md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPollTemplatesUsingGET**
> PageResourceTemplateResource GetPollTemplatesUsingGET($size, $page, $order)

List and search poll templates


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to id:ASC]

### Return type

[**PageResourceTemplateResource**](PageResource«TemplateResource».md)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPollUsingGET**
> PollResource GetPollUsingGET($id)

Get a single poll


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The poll id | 

### Return type

[**PollResource**](PollResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPollsUsingGET**
> PageResourcePollResource GetPollsUsingGET($filterCategory, $filterTagset, $filterText, $size, $page, $order)

List and search polls

Get a list of polls with optional filtering. Assets will not be filled in on the resources returned. Use 'Get a single poll' to retrieve the full resource with assets for a given item as needed.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterCategory** | **string**| Filter for polls from a specific category by id | [optional] 
 **filterTagset** | **string**| Filter for polls with specified tags (separated by comma) | [optional] 
 **filterText** | **string**| Filter for polls whose text contains a string | [optional] 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to id:ASC]

### Return type

[**PageResourcePollResource**](PageResource«PollResource».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePollTemplateUsingPUT**
> UpdatePollTemplateUsingPUT($id, $pollTemplateResource)

Update a poll template


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 
 **pollTemplateResource** | [**TemplateResource**](TemplateResource.md)| The poll template resource object | [optional] 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePollUsingPUT**
> UpdatePollUsingPUT($id, $pollResource)

Update an existing poll


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The poll id | 
 **pollResource** | [**PollResource**](PollResource.md)| The poll object | [optional] 

### Return type

void (empty response body)

### Authorization

[knetik_oauth](../README.md#knetik_oauth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

