# \GamificationTriviaApi

All URIs are relative to *https://devsandbox.knetikcloud.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAnswersUsingPOST**](GamificationTriviaApi.md#AddAnswersUsingPOST) | **Post** /trivia/questions/{question_id}/answers | Add an answer to a question
[**AddTagUsingPOST**](GamificationTriviaApi.md#AddTagUsingPOST) | **Post** /trivia/questions/{id}/tags | Add a tag to a question
[**BatchAddTagUsingPOST**](GamificationTriviaApi.md#BatchAddTagUsingPOST) | **Post** /trivia/questions/tags | Add a tag to a batch of questions
[**BatchRemoveTagUsingDELETE**](GamificationTriviaApi.md#BatchRemoveTagUsingDELETE) | **Delete** /trivia/questions/tags/{tag} | Remove a tag from a batch of questions
[**BulkUpdateUsingPUT**](GamificationTriviaApi.md#BulkUpdateUsingPUT) | **Put** /trivia/questions | Bulk update questions
[**CountQuestionsUsingGET**](GamificationTriviaApi.md#CountQuestionsUsingGET) | **Get** /trivia/questions/count | Count questions based on filters.
[**CreateQuestionTemplateUsingPOST**](GamificationTriviaApi.md#CreateQuestionTemplateUsingPOST) | **Post** /trivia/questions/templates | Create a question template
[**CreateQuestionUsingPOST**](GamificationTriviaApi.md#CreateQuestionUsingPOST) | **Post** /trivia/questions | Create a question
[**CreateUsingPOST**](GamificationTriviaApi.md#CreateUsingPOST) | **Post** /trivia/import | Create an import job
[**DeleteQuestionTemplateUsingDELETE**](GamificationTriviaApi.md#DeleteQuestionTemplateUsingDELETE) | **Delete** /trivia/questions/templates/{id} | Delete a question template
[**DeleteQuestionUsingDELETE**](GamificationTriviaApi.md#DeleteQuestionUsingDELETE) | **Delete** /trivia/questions/{id} | Delete a question
[**DeleteUsingDELETE**](GamificationTriviaApi.md#DeleteUsingDELETE) | **Delete** /trivia/import/{id} | Delete an import job
[**GetAnswerUsingGET**](GamificationTriviaApi.md#GetAnswerUsingGET) | **Get** /trivia/questions/{question_id}/answers/{id} | Get an answer for a question
[**GetAnswersUsingGET**](GamificationTriviaApi.md#GetAnswersUsingGET) | **Get** /trivia/questions/{question_id}/answers | List the answers available for a question
[**GetListUsingGET1**](GamificationTriviaApi.md#GetListUsingGET1) | **Get** /trivia/import | Get a list of import job
[**GetQuestionTemplateUsingGET**](GamificationTriviaApi.md#GetQuestionTemplateUsingGET) | **Get** /trivia/questions/templates/{id} | Get a single question template
[**GetQuestionTemplatesUsingGET**](GamificationTriviaApi.md#GetQuestionTemplatesUsingGET) | **Get** /trivia/questions/templates | List and search question templates
[**GetQuestionUsingGET**](GamificationTriviaApi.md#GetQuestionUsingGET) | **Get** /trivia/questions/{id} | Get a single question
[**GetQuestionsDeltaUsingGET**](GamificationTriviaApi.md#GetQuestionsDeltaUsingGET) | **Get** /trivia/questions/delta | List question deltas in ascending order of updated date
[**GetQuestionsUsingGET**](GamificationTriviaApi.md#GetQuestionsUsingGET) | **Get** /trivia/questions | List and search questions
[**GetTagsUsingGET1**](GamificationTriviaApi.md#GetTagsUsingGET1) | **Get** /trivia/questions/{id}/tags | List the tags for a question
[**GetTagsUsingGET2**](GamificationTriviaApi.md#GetTagsUsingGET2) | **Get** /trivia/tags | List and search tags by the beginning of the string
[**GetUsingGET**](GamificationTriviaApi.md#GetUsingGET) | **Get** /trivia/import/{id} | Get an import job
[**RemoveAnswersUsingDELETE**](GamificationTriviaApi.md#RemoveAnswersUsingDELETE) | **Delete** /trivia/questions/{question_id}/answers/{id} | Remove an answer from a question
[**RemoveTagUsingDELETE**](GamificationTriviaApi.md#RemoveTagUsingDELETE) | **Delete** /trivia/questions/{id}/tags/{tag} | Remove a tag from a question
[**StartProcessUsingPOST**](GamificationTriviaApi.md#StartProcessUsingPOST) | **Post** /trivia/import/{id}/process | Start processing an import job
[**UpdateAnswerUsingPUT**](GamificationTriviaApi.md#UpdateAnswerUsingPUT) | **Put** /trivia/questions/{question_id}/answers/{id} | Update an answer for a question
[**UpdateQuestionTemplateUsingPUT**](GamificationTriviaApi.md#UpdateQuestionTemplateUsingPUT) | **Put** /trivia/questions/templates/{id} | Update a question template
[**UpdateQuestionUsingPUT**](GamificationTriviaApi.md#UpdateQuestionUsingPUT) | **Put** /trivia/questions/{id} | Update a question
[**UpdateUsingPUT**](GamificationTriviaApi.md#UpdateUsingPUT) | **Put** /trivia/import/{id} | Update an import job


# **AddAnswersUsingPOST**
> AnswerResource AddAnswersUsingPOST($questionId, $answer)

Add an answer to a question


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **questionId** | **string**| The id of the question | 
 **answer** | [**AnswerResource**](AnswerResource.md)| The new answer | [optional] 

### Return type

[**AnswerResource**](AnswerResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddTagUsingPOST**
> AddTagUsingPOST($id, $tag)

Add a tag to a question


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the question | 
 **tag** | **string**| The new tag | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BatchAddTagUsingPOST**
> int32 BatchAddTagUsingPOST($tag, $filterSearch, $filterIdset, $filterCategory, $filterTag, $filterTagset, $filterType, $filterPublished, $filterImportId)

Add a tag to a batch of questions

All questions that dont't have the tag and match filters will have it added. The returned number is the number of questions updated.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tag** | **string**| The tag to add | [optional] 
 **filterSearch** | **string**| Filter for documents whose question, answers or tags contains provided string | [optional] 
 **filterIdset** | **string**| Filter for documents whose id is in the comma separated list provided | [optional] 
 **filterCategory** | **string**| Filter for questions with specified category, by id | [optional] 
 **filterTag** | **string**| Filter for questions with specified tag | [optional] 
 **filterTagset** | **string**| Filter for questions with specified tags (separated by comma) | [optional] 
 **filterType** | **string**| Filter for questions with specified type | [optional] 
 **filterPublished** | **bool**| Filter for questions currenctly published or not | [optional] 
 **filterImportId** | **int64**| Filter for questions from a specific import job | [optional] 

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BatchRemoveTagUsingDELETE**
> int32 BatchRemoveTagUsingDELETE($tag, $filterSearch, $filterIdset, $filterCategory, $filterTag, $filterTagset, $filterType, $filterPublished, $filterImportId)

Remove a tag from a batch of questions

ll questions that have the tag and match filters will have it removed. The returned number is the number of questions updated.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tag** | **string**| The tag to remove | 
 **filterSearch** | **string**| Filter for documents whose question, answers or tags contains provided string | [optional] 
 **filterIdset** | **string**| Filter for documents whose id is in the comma separated list provided | [optional] 
 **filterCategory** | **string**| Filter for questions with specified category, by id | [optional] 
 **filterTag** | **string**| Filter for questions with specified tag | [optional] 
 **filterTagset** | **string**| Filter for questions with specified tags (separated by comma) | [optional] 
 **filterType** | **string**| Filter for questions with specified type.  Allowable values: (&#39;TEXT&#39;, &#39;IMAGE&#39;, &#39;VIDEO&#39;, &#39;AUDIO&#39;) | [optional] 
 **filterPublished** | **bool**| Filter for questions currenctly published or not | [optional] 
 **filterImportId** | **int64**| Filter for questions from a specific import job | [optional] 

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BulkUpdateUsingPUT**
> int32 BulkUpdateUsingPUT($question, $filterSearch, $filterIdset, $filterCategory, $filterTagset, $filterType, $filterPublished, $filterImportId)

Bulk update questions

Will update all questions that match filters used (or all questions in system if no filters used). Body should match a question resource with only those properties you wish to set. Null values will be ignored. Returned number is how many were updated.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **question** | [**QuestionResource**](QuestionResource.md)| New values for a set of question fields | [optional] 
 **filterSearch** | **string**| Filter for documents whose question, answers or tags contains provided string | [optional] 
 **filterIdset** | **string**| Filter for documents whose id is in the comma separated list provided | [optional] 
 **filterCategory** | **string**| Filter for questions with specified category, by id | [optional] 
 **filterTagset** | **string**| Filter for questions with specified tags (separated by comma) | [optional] 
 **filterType** | **string**| Filter for questions with specified type.  Allowable values: (&#39;TEXT&#39;, &#39;IMAGE&#39;, &#39;VIDEO&#39;, &#39;AUDIO&#39;) | [optional] 
 **filterPublished** | **bool**| Filter for questions currenctly published or not | [optional] 
 **filterImportId** | **int64**| Filter for questions from a specific import job | [optional] 

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CountQuestionsUsingGET**
> int64 CountQuestionsUsingGET($filterSearch, $filterIdset, $filterCategory, $filterTag, $filterTagset, $filterType, $filterPublished)

Count questions based on filters.

This is also provided by the list endpoint so you don't need to call this for pagination purposes


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterSearch** | **string**| Filter for documents whose question, answers or tags contains provided string | [optional] 
 **filterIdset** | **string**| Filter for documents whose id is in the comma separated list provided | [optional] 
 **filterCategory** | **string**| Filter for questions with specified category, by id | [optional] 
 **filterTag** | **string**| Filter for questions with specified tag | [optional] 
 **filterTagset** | **string**| Filter for questions with specified tags (separated by comma) | [optional] 
 **filterType** | **string**| Filter for questions with specified type.  Allowable values: (&#39;TEXT&#39;, &#39;IMAGE&#39;, &#39;VIDEO&#39;, &#39;AUDIO&#39;) | [optional] 
 **filterPublished** | **bool**| Filter for questions currenctly published or not | [optional] 

### Return type

**int64**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateQuestionTemplateUsingPOST**
> QuestionTemplateResource CreateQuestionTemplateUsingPOST($questionTemplateResource)

Create a question template

Question templates define a type of question and the properties they have


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **questionTemplateResource** | [**QuestionTemplateResource**](QuestionTemplateResource.md)| The question template resource object | [optional] 

### Return type

[**QuestionTemplateResource**](QuestionTemplateResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateQuestionUsingPOST**
> QuestionResource CreateQuestionUsingPOST($question)

Create a question


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **question** | [**QuestionResource**](QuestionResource.md)| The new question | [optional] 

### Return type

[**QuestionResource**](QuestionResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateUsingPOST**
> ImportJobResource CreateUsingPOST($request)

Create an import job

Set up a job to import a set of trivia questions from a cvs file at a remote url. the file will be validated asynchronously but will not be processed until started manually with the process endpoint.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**ImportJobResource**](ImportJobResource.md)| The new import job | [optional] 

### Return type

[**ImportJobResource**](ImportJobResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteQuestionTemplateUsingDELETE**
> DeleteQuestionTemplateUsingDELETE($id, $cascade)

Delete a question template

If cascade = 'detach', it will force delete the template even if it's attached to other objects


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 
 **cascade** | **string**| The value needed to delete used templates | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteQuestionUsingDELETE**
> DeleteQuestionUsingDELETE($id)

Delete a question


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the question | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUsingDELETE**
> DeleteUsingDELETE($id)

Delete an import job

Also deletes all questions that were imported by it


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| The id of the job | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAnswerUsingGET**
> AnswerResource GetAnswerUsingGET($questionId, $id)

Get an answer for a question


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **questionId** | **string**| The id of the question | 
 **id** | **string**| The id of the answer | 

### Return type

[**AnswerResource**](AnswerResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAnswersUsingGET**
> []AnswerResource GetAnswersUsingGET($questionId)

List the answers available for a question


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **questionId** | **string**| The id of the question | 

### Return type

[**[]AnswerResource**](AnswerResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListUsingGET1**
> PageImportJobResource GetListUsingGET1($filterVendor, $filterCategory, $filterName, $filterStatus, $size, $page, $order)

Get a list of import job


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterVendor** | **string**| Filter for jobs by vendor id | [optional] 
 **filterCategory** | **string**| Filter for jobs by category id | [optional] 
 **filterName** | **string**| Filter for jobs which name *STARTS* with the given string | [optional] 
 **filterStatus** | **string**| Filter for jobs that are in a specific set of statuses (comma separated) | [optional] 
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to 1]

### Return type

[**PageImportJobResource**](Page«ImportJobResource».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetQuestionTemplateUsingGET**
> QuestionTemplateResource GetQuestionTemplateUsingGET($id)

Get a single question template


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 

### Return type

[**QuestionTemplateResource**](QuestionTemplateResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetQuestionTemplatesUsingGET**
> PageQuestionTemplateResource GetQuestionTemplatesUsingGET($size, $page, $order)

List and search question templates


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to 1]

### Return type

[**PageQuestionTemplateResource**](Page«QuestionTemplateResource».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetQuestionUsingGET**
> QuestionResource GetQuestionUsingGET($id)

Get a single question


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the question | 

### Return type

[**QuestionResource**](QuestionResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetQuestionsDeltaUsingGET**
> []DeltaResource GetQuestionsDeltaUsingGET($since)

List question deltas in ascending order of updated date

The 'since' parameter is important to avoid getting a full list of all questions. Implementors should make sure they pass the updated date of the last resource loaded, not the date of the last request, in order to avoid gaps


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **since** | **int64**| Timestamp in seconds | [optional] 

### Return type

[**[]DeltaResource**](DeltaResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetQuestionsUsingGET**
> PageQuestionResource GetQuestionsUsingGET($size, $page, $order, $filterSearch, $filterIdset, $filterCategory, $filterTagset, $filterType, $filterPublished, $filterImportId)

List and search questions


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to 1]
 **filterSearch** | **string**| Filter for documents whose question, answers or tags contains provided string | [optional] 
 **filterIdset** | **string**| Filter for documents whose id is in the comma separated list provided | [optional] 
 **filterCategory** | **string**| Filter for questions with specified category, by id | [optional] 
 **filterTagset** | **string**| Filter for questions with specified tags (separated by comma) | [optional] 
 **filterType** | **string**| Filter for questions with specified type.  Allowable values: (&#39;TEXT&#39;, &#39;IMAGE&#39;, &#39;VIDEO&#39;, &#39;AUDIO&#39;) | [optional] 
 **filterPublished** | **bool**| Filter for questions currenctly published or not | [optional] 
 **filterImportId** | **int64**| Filter for questions from a specific import job | [optional] 

### Return type

[**PageQuestionResource**](Page«QuestionResource».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTagsUsingGET1**
> []string GetTagsUsingGET1($id)

List the tags for a question


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the question | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTagsUsingGET2**
> Collectionstring GetTagsUsingGET2($filterSearch, $filterCategory, $filterImportId)

List and search tags by the beginning of the string

For performance reasons, search & category filters are mutually exclusive. If category is specified, search filter will be ignored in order to do fast matches for typeahead.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterSearch** | **string**| Filter for tags starting with the given text | [optional] 
 **filterCategory** | **string**| Filter for tags on questions from a specific category | [optional] 
 **filterImportId** | **int64**| Filter for tags on questions from a specific import job | [optional] 

### Return type

[**Collectionstring**](Collection«string».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsingGET**
> ImportJobResource GetUsingGET($id)

Get an import job


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| The id of the job | 

### Return type

[**ImportJobResource**](ImportJobResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveAnswersUsingDELETE**
> RemoveAnswersUsingDELETE($questionId, $id)

Remove an answer from a question


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **questionId** | **string**| The id of the question | 
 **id** | **string**| The id of the answer | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveTagUsingDELETE**
> RemoveTagUsingDELETE($id, $tag)

Remove a tag from a question


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the question | 
 **tag** | **string**| The tag to remove | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartProcessUsingPOST**
> ImportJobResource StartProcessUsingPOST($id, $publishNow)

Start processing an import job

Will process the CSV file and add new questions asynchronously. The status of the job must be 'VALID'.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| The id of the job | 
 **publishNow** | **bool**| Whether the new questions should be published live immediately | 

### Return type

[**ImportJobResource**](ImportJobResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAnswerUsingPUT**
> UpdateAnswerUsingPUT($questionId, $id, $answer)

Update an answer for a question


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **questionId** | **string**| The id of the question | 
 **id** | **string**| The id of the answer | 
 **answer** | [**AnswerResource**](AnswerResource.md)| The updated answer | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateQuestionTemplateUsingPUT**
> UpdateQuestionTemplateUsingPUT($id, $questionTemplateResource)

Update a question template


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 
 **questionTemplateResource** | [**QuestionTemplateResource**](QuestionTemplateResource.md)| The question template resource object | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateQuestionUsingPUT**
> QuestionResource UpdateQuestionUsingPUT($id, $question)

Update a question


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the question | 
 **question** | [**QuestionResource**](QuestionResource.md)| The updated question | [optional] 

### Return type

[**QuestionResource**](QuestionResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUsingPUT**
> ImportJobResource UpdateUsingPUT($id, $request)

Update an import job

Changes should be made before process is started for there to be any effect.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| The id of the job | 
 **request** | [**ImportJobResource**](ImportJobResource.md)| The updated job | [optional] 

### Return type

[**ImportJobResource**](ImportJobResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

