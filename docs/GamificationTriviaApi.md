# \GamificationTriviaApi

All URIs are relative to *https://localhost:8080/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddQuestionAnswers**](GamificationTriviaApi.md#AddQuestionAnswers) | **Post** /trivia/questions/{question_id}/answers | Add an answer to a question
[**AddQuestionTag**](GamificationTriviaApi.md#AddQuestionTag) | **Post** /trivia/questions/{id}/tags | Add a tag to a question
[**AddTagToQuestionsBatch**](GamificationTriviaApi.md#AddTagToQuestionsBatch) | **Post** /trivia/questions/tags | Add a tag to a batch of questions
[**CreateImportJob**](GamificationTriviaApi.md#CreateImportJob) | **Post** /trivia/import | Create an import job
[**CreateQuestion**](GamificationTriviaApi.md#CreateQuestion) | **Post** /trivia/questions | Create a question
[**CreateQuestionTemplate**](GamificationTriviaApi.md#CreateQuestionTemplate) | **Post** /trivia/questions/templates | Create a question template
[**DeleteImportJob**](GamificationTriviaApi.md#DeleteImportJob) | **Delete** /trivia/import/{id} | Delete an import job
[**DeleteQuestion**](GamificationTriviaApi.md#DeleteQuestion) | **Delete** /trivia/questions/{id} | Delete a question
[**DeleteQuestionAnswers**](GamificationTriviaApi.md#DeleteQuestionAnswers) | **Delete** /trivia/questions/{question_id}/answers/{id} | Remove an answer from a question
[**DeleteQuestionTemplate**](GamificationTriviaApi.md#DeleteQuestionTemplate) | **Delete** /trivia/questions/templates/{id} | Delete a question template
[**GetImportJob**](GamificationTriviaApi.md#GetImportJob) | **Get** /trivia/import/{id} | Get an import job
[**GetImportJobs**](GamificationTriviaApi.md#GetImportJobs) | **Get** /trivia/import | Get a list of import job
[**GetQuestion**](GamificationTriviaApi.md#GetQuestion) | **Get** /trivia/questions/{id} | Get a single question
[**GetQuestionAnswer**](GamificationTriviaApi.md#GetQuestionAnswer) | **Get** /trivia/questions/{question_id}/answers/{id} | Get an answer for a question
[**GetQuestionAnswers**](GamificationTriviaApi.md#GetQuestionAnswers) | **Get** /trivia/questions/{question_id}/answers | List the answers available for a question
[**GetQuestionDeltas**](GamificationTriviaApi.md#GetQuestionDeltas) | **Get** /trivia/questions/delta | List question deltas in ascending order of updated date
[**GetQuestionTags**](GamificationTriviaApi.md#GetQuestionTags) | **Get** /trivia/questions/{id}/tags | List the tags for a question
[**GetQuestionTemplate**](GamificationTriviaApi.md#GetQuestionTemplate) | **Get** /trivia/questions/templates/{id} | Get a single question template
[**GetQuestionTemplates**](GamificationTriviaApi.md#GetQuestionTemplates) | **Get** /trivia/questions/templates | List and search question templates
[**GetQuestions**](GamificationTriviaApi.md#GetQuestions) | **Get** /trivia/questions | List and search questions
[**GetQuestionsCount**](GamificationTriviaApi.md#GetQuestionsCount) | **Get** /trivia/questions/count | Count questions based on filters
[**ProcessImportJob**](GamificationTriviaApi.md#ProcessImportJob) | **Post** /trivia/import/{id}/process | Start processing an import job
[**RemoveQuestionTag**](GamificationTriviaApi.md#RemoveQuestionTag) | **Delete** /trivia/questions/{id}/tags/{tag} | Remove a tag from a question
[**RemoveTagToQuestionsBatch**](GamificationTriviaApi.md#RemoveTagToQuestionsBatch) | **Delete** /trivia/questions/tags/{tag} | Remove a tag from a batch of questions
[**SearchQuestionTags**](GamificationTriviaApi.md#SearchQuestionTags) | **Get** /trivia/tags | List and search tags by the beginning of the string
[**UpdateImportJob**](GamificationTriviaApi.md#UpdateImportJob) | **Put** /trivia/import/{id} | Update an import job
[**UpdateQuestion**](GamificationTriviaApi.md#UpdateQuestion) | **Put** /trivia/questions/{id} | Update a question
[**UpdateQuestionAnswer**](GamificationTriviaApi.md#UpdateQuestionAnswer) | **Put** /trivia/questions/{question_id}/answers/{id} | Update an answer for a question
[**UpdateQuestionTemplate**](GamificationTriviaApi.md#UpdateQuestionTemplate) | **Put** /trivia/questions/templates/{id} | Update a question template
[**UpdateQuestionsInBulk**](GamificationTriviaApi.md#UpdateQuestionsInBulk) | **Put** /trivia/questions | Bulk update questions


# **AddQuestionAnswers**
> AnswerResource AddQuestionAnswers($questionId, $answer)

Add an answer to a question


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **questionId** | **string**| The id of the question | 
 **answer** | [**AnswerResource**](AnswerResource.md)| The new answer | [optional] 

### Return type

[**AnswerResource**](AnswerResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddQuestionTag**
> AddQuestionTag($id, $tag)

Add a tag to a question


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the question | 
 **tag** | **string**| The new tag | [optional] 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddTagToQuestionsBatch**
> int32 AddTagToQuestionsBatch($tag, $filterSearch, $filterIdset, $filterCategory, $filterTag, $filterTagset, $filterType, $filterPublished, $filterImportId)

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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateImportJob**
> ImportJobResource CreateImportJob($request)

Create an import job

Set up a job to import a set of trivia questions from a cvs file at a remote url. the file will be validated asynchronously but will not be processed until started manually with the process endpoint.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**ImportJobResource**](ImportJobResource.md)| The new import job | [optional] 

### Return type

[**ImportJobResource**](ImportJobResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateQuestion**
> QuestionResource CreateQuestion($question)

Create a question


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **question** | [**QuestionResource**](QuestionResource.md)| The new question | [optional] 

### Return type

[**QuestionResource**](QuestionResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateQuestionTemplate**
> QuestionTemplateResource CreateQuestionTemplate($questionTemplateResource)

Create a question template

Question templates define a type of question and the properties they have


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **questionTemplateResource** | [**QuestionTemplateResource**](QuestionTemplateResource.md)| The question template resource object | [optional] 

### Return type

[**QuestionTemplateResource**](QuestionTemplateResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteImportJob**
> DeleteImportJob($id)

Delete an import job

Also deletes all questions that were imported by it


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| The id of the job | 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteQuestion**
> DeleteQuestion($id)

Delete a question


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the question | 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteQuestionAnswers**
> DeleteQuestionAnswers($questionId, $id)

Remove an answer from a question


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **questionId** | **string**| The id of the question | 
 **id** | **string**| The id of the answer | 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteQuestionTemplate**
> DeleteQuestionTemplate($id, $cascade)

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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetImportJob**
> ImportJobResource GetImportJob($id)

Get an import job


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| The id of the job | 

### Return type

[**ImportJobResource**](ImportJobResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetImportJobs**
> PageResourceImportJobResource GetImportJobs($filterVendor, $filterCategory, $filterName, $filterStatus, $size, $page, $order)

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
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to id:ASC]

### Return type

[**PageResourceImportJobResource**](PageResource«ImportJobResource».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetQuestion**
> QuestionResource GetQuestion($id)

Get a single question


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the question | 

### Return type

[**QuestionResource**](QuestionResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetQuestionAnswer**
> AnswerResource GetQuestionAnswer($questionId, $id)

Get an answer for a question


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **questionId** | **string**| The id of the question | 
 **id** | **string**| The id of the answer | 

### Return type

[**AnswerResource**](AnswerResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetQuestionAnswers**
> []AnswerResource GetQuestionAnswers($questionId)

List the answers available for a question


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **questionId** | **string**| The id of the question | 

### Return type

[**[]AnswerResource**](AnswerResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetQuestionDeltas**
> []DeltaResource GetQuestionDeltas($since)

List question deltas in ascending order of updated date

The 'since' parameter is important to avoid getting a full list of all questions. Implementors should make sure they pass the updated date of the last resource loaded, not the date of the last request, in order to avoid gaps


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **since** | **int64**| Timestamp in seconds | [optional] 

### Return type

[**[]DeltaResource**](DeltaResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetQuestionTags**
> []string GetQuestionTags($id)

List the tags for a question


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the question | 

### Return type

**[]string**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetQuestionTemplate**
> QuestionTemplateResource GetQuestionTemplate($id)

Get a single question template


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 

### Return type

[**QuestionTemplateResource**](QuestionTemplateResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetQuestionTemplates**
> PageResourceQuestionTemplateResource GetQuestionTemplates($size, $page, $order)

List and search question templates


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to id:ASC]

### Return type

[**PageResourceQuestionTemplateResource**](PageResource«QuestionTemplateResource».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetQuestions**
> PageResourceQuestionResource GetQuestions($size, $page, $order, $filterSearch, $filterIdset, $filterCategory, $filterTagset, $filterTag, $filterType, $filterPublished, $filterImportId)

List and search questions


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32**| The number of objects returned per page | [optional] [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [optional] [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [optional] [default to id:ASC]
 **filterSearch** | **string**| Filter for documents whose question, answers or tags contains provided string | [optional] 
 **filterIdset** | **string**| Filter for documents whose id is in the comma separated list provided | [optional] 
 **filterCategory** | **string**| Filter for questions with specified category, by id | [optional] 
 **filterTagset** | **string**| Filter for questions with specified tags (separated by comma) | [optional] 
 **filterTag** | **string**| Filter for questions with specified tag | [optional] 
 **filterType** | **string**| Filter for questions with specified type.  Allowable values: (&#39;TEXT&#39;, &#39;IMAGE&#39;, &#39;VIDEO&#39;, &#39;AUDIO&#39;) | [optional] 
 **filterPublished** | **bool**| Filter for questions currenctly published or not | [optional] 
 **filterImportId** | **int64**| Filter for questions from a specific import job | [optional] 

### Return type

[**PageResourceQuestionResource**](PageResource«QuestionResource».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetQuestionsCount**
> int64 GetQuestionsCount($filterSearch, $filterIdset, $filterCategory, $filterTag, $filterTagset, $filterType, $filterPublished)

Count questions based on filters

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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProcessImportJob**
> ImportJobResource ProcessImportJob($id, $publishNow)

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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveQuestionTag**
> RemoveQuestionTag($id, $tag)

Remove a tag from a question


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the question | 
 **tag** | **string**| The tag to remove | 

### Return type

void (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveTagToQuestionsBatch**
> int32 RemoveTagToQuestionsBatch($tag, $filterSearch, $filterIdset, $filterCategory, $filterTag, $filterTagset, $filterType, $filterPublished, $filterImportId)

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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchQuestionTags**
> Collectionstring SearchQuestionTags($filterSearch, $filterCategory, $filterImportId)

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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateImportJob**
> ImportJobResource UpdateImportJob($id, $request)

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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateQuestion**
> QuestionResource UpdateQuestion($id, $question)

Update a question


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the question | 
 **question** | [**QuestionResource**](QuestionResource.md)| The updated question | [optional] 

### Return type

[**QuestionResource**](QuestionResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateQuestionAnswer**
> UpdateQuestionAnswer($questionId, $id, $answer)

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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateQuestionTemplate**
> QuestionTemplateResource UpdateQuestionTemplate($id, $questionTemplateResource)

Update a question template


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 
 **questionTemplateResource** | [**QuestionTemplateResource**](QuestionTemplateResource.md)| The question template resource object | [optional] 

### Return type

[**QuestionTemplateResource**](QuestionTemplateResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateQuestionsInBulk**
> int32 UpdateQuestionsInBulk($question, $filterSearch, $filterIdset, $filterCategory, $filterTagset, $filterType, $filterPublished, $filterImportId)

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

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

