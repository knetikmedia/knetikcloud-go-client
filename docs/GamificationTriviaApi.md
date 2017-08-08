# \GamificationTriviaApi

All URIs are relative to *https://sandbox.knetikcloud.com*

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
> AnswerResource AddQuestionAnswers(ctx, questionId, optional)
Add an answer to a question

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **questionId** | **string**| The id of the question | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **questionId** | **string**| The id of the question | 
 **answer** | [**AnswerResource**](AnswerResource.md)| The new answer | 

### Return type

[**AnswerResource**](AnswerResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddQuestionTag**
> AddQuestionTag(ctx, id, optional)
Add a tag to a question

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **string**| The id of the question | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the question | 
 **tag** | [**StringWrapper**](StringWrapper.md)| The new tag | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddTagToQuestionsBatch**
> int32 AddTagToQuestionsBatch(ctx, optional)
Add a tag to a batch of questions

All questions that dont't have the tag and match filters will have it added. The returned number is the number of questions updated.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tag** | [**StringWrapper**](StringWrapper.md)| The tag to add | 
 **filterSearch** | **string**| Filter for documents whose question, answers or tags contains provided string | 
 **filterIdset** | **string**| Filter for documents whose id is in the comma separated list provided | 
 **filterCategory** | **string**| Filter for questions with specified category, by id | 
 **filterTag** | **string**| Filter for questions with specified tag | 
 **filterTagset** | **string**| Filter for questions with specified tags (separated by comma) | 
 **filterType** | **string**| Filter for questions with specified type | 
 **filterPublished** | **bool**| Filter for questions currenctly published or not | 
 **filterImportId** | **int64**| Filter for questions from a specific import job | 

### Return type

**int32**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateImportJob**
> ImportJobResource CreateImportJob(ctx, optional)
Create an import job

Set up a job to import a set of trivia questions from a cvs file at a remote url. the file will be validated asynchronously but will not be processed until started manually with the process endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**ImportJobResource**](ImportJobResource.md)| The new import job | 

### Return type

[**ImportJobResource**](ImportJobResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateQuestion**
> QuestionResource CreateQuestion(ctx, optional)
Create a question

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **question** | [**QuestionResource**](QuestionResource.md)| The new question | 

### Return type

[**QuestionResource**](QuestionResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateQuestionTemplate**
> QuestionTemplateResource CreateQuestionTemplate(ctx, optional)
Create a question template

Question templates define a type of question and the properties they have

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **questionTemplateResource** | [**QuestionTemplateResource**](QuestionTemplateResource.md)| The question template resource object | 

### Return type

[**QuestionTemplateResource**](QuestionTemplateResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteImportJob**
> DeleteImportJob(ctx, id)
Delete an import job

Also deletes all questions that were imported by it

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **int64**| The id of the job | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteQuestion**
> DeleteQuestion(ctx, id)
Delete a question

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **string**| The id of the question | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteQuestionAnswers**
> DeleteQuestionAnswers(ctx, questionId, id)
Remove an answer from a question

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **questionId** | **string**| The id of the question | 
  **id** | **string**| The id of the answer | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteQuestionTemplate**
> DeleteQuestionTemplate(ctx, id, optional)
Delete a question template

If cascade = 'detach', it will force delete the template even if it's attached to other objects

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **string**| The id of the template | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 
 **cascade** | **string**| The value needed to delete used templates | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetImportJob**
> ImportJobResource GetImportJob(ctx, id)
Get an import job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
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
> PageResourceImportJobResource GetImportJobs(ctx, optional)
Get a list of import job

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterVendor** | **string**| Filter for jobs by vendor id | 
 **filterCategory** | **string**| Filter for jobs by category id | 
 **filterName** | **string**| Filter for jobs which name *STARTS* with the given string | 
 **filterStatus** | **string**| Filter for jobs that are in a specific set of statuses (comma separated) | 
 **size** | **int32**| The number of objects returned per page | [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [default to id:ASC]

### Return type

[**PageResourceImportJobResource**](PageResource«ImportJobResource».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetQuestion**
> QuestionResource GetQuestion(ctx, id)
Get a single question

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
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
> AnswerResource GetQuestionAnswer(ctx, questionId, id)
Get an answer for a question

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
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
> []AnswerResource GetQuestionAnswers(ctx, questionId)
List the answers available for a question

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
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
> []DeltaResource GetQuestionDeltas(ctx, optional)
List question deltas in ascending order of updated date

The 'since' parameter is important to avoid getting a full list of all questions. Implementors should make sure they pass the updated date of the last resource loaded, not the date of the last request, in order to avoid gaps

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **since** | **int64**| Timestamp in seconds | 

### Return type

[**[]DeltaResource**](DeltaResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetQuestionTags**
> []string GetQuestionTags(ctx, id)
List the tags for a question

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
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
> QuestionTemplateResource GetQuestionTemplate(ctx, id)
Get a single question template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
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
> PageResourceQuestionTemplateResource GetQuestionTemplates(ctx, optional)
List and search question templates

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32**| The number of objects returned per page | [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [default to id:ASC]

### Return type

[**PageResourceQuestionTemplateResource**](PageResource«QuestionTemplateResource».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetQuestions**
> PageResourceQuestionResource GetQuestions(ctx, optional)
List and search questions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32**| The number of objects returned per page | [default to 25]
 **page** | **int32**| The number of the page returned, starting with 1 | [default to 1]
 **order** | **string**| A comma separated list of sorting requirements in priority order, each entry matching PROPERTY_NAME:[ASC|DESC] | [default to id:ASC]
 **filterSearch** | **string**| Filter for documents whose question, answers or tags contains provided string | 
 **filterIdset** | **string**| Filter for documents whose id is in the comma separated list provided | 
 **filterCategory** | **string**| Filter for questions with specified category, by id | 
 **filterTagset** | **string**| Filter for questions with specified tags (separated by comma) | 
 **filterTag** | **string**| Filter for questions with specified tag | 
 **filterType** | **string**| Filter for questions with specified type.  Allowable values: (&#39;TEXT&#39;, &#39;IMAGE&#39;, &#39;VIDEO&#39;, &#39;AUDIO&#39;) | 
 **filterPublished** | **bool**| Filter for questions currenctly published or not | 
 **filterImportId** | **int64**| Filter for questions from a specific import job | 

### Return type

[**PageResourceQuestionResource**](PageResource«QuestionResource».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetQuestionsCount**
> int64 GetQuestionsCount(ctx, optional)
Count questions based on filters

This is also provided by the list endpoint so you don't need to call this for pagination purposes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterSearch** | **string**| Filter for documents whose question, answers or tags contains provided string | 
 **filterIdset** | **string**| Filter for documents whose id is in the comma separated list provided | 
 **filterCategory** | **string**| Filter for questions with specified category, by id | 
 **filterTag** | **string**| Filter for questions with specified tag | 
 **filterTagset** | **string**| Filter for questions with specified tags (separated by comma) | 
 **filterType** | **string**| Filter for questions with specified type.  Allowable values: (&#39;TEXT&#39;, &#39;IMAGE&#39;, &#39;VIDEO&#39;, &#39;AUDIO&#39;) | 
 **filterPublished** | **bool**| Filter for questions currenctly published or not | 

### Return type

**int64**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProcessImportJob**
> ImportJobResource ProcessImportJob(ctx, id, publishNow)
Start processing an import job

Will process the CSV file and add new questions asynchronously. The status of the job must be 'VALID'.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
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
> RemoveQuestionTag(ctx, id, tag)
Remove a tag from a question

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **string**| The id of the question | 
  **tag** | **string**| The tag to remove | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveTagToQuestionsBatch**
> int32 RemoveTagToQuestionsBatch(ctx, tag, optional)
Remove a tag from a batch of questions

ll questions that have the tag and match filters will have it removed. The returned number is the number of questions updated.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **tag** | **string**| The tag to remove | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tag** | **string**| The tag to remove | 
 **filterSearch** | **string**| Filter for documents whose question, answers or tags contains provided string | 
 **filterIdset** | **string**| Filter for documents whose id is in the comma separated list provided | 
 **filterCategory** | **string**| Filter for questions with specified category, by id | 
 **filterTag** | **string**| Filter for questions with specified tag | 
 **filterTagset** | **string**| Filter for questions with specified tags (separated by comma) | 
 **filterType** | **string**| Filter for questions with specified type.  Allowable values: (&#39;TEXT&#39;, &#39;IMAGE&#39;, &#39;VIDEO&#39;, &#39;AUDIO&#39;) | 
 **filterPublished** | **bool**| Filter for questions currenctly published or not | 
 **filterImportId** | **int64**| Filter for questions from a specific import job | 

### Return type

**int32**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchQuestionTags**
> Collectionstring SearchQuestionTags(ctx, optional)
List and search tags by the beginning of the string

For performance reasons, search & category filters are mutually exclusive. If category is specified, search filter will be ignored in order to do fast matches for typeahead.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterSearch** | **string**| Filter for tags starting with the given text | 
 **filterCategory** | **string**| Filter for tags on questions from a specific category | 
 **filterImportId** | **int64**| Filter for tags on questions from a specific import job | 

### Return type

[**Collectionstring**](Collection«string».md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateImportJob**
> ImportJobResource UpdateImportJob(ctx, id, optional)
Update an import job

Changes should be made before process is started for there to be any effect.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **int64**| The id of the job | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int64**| The id of the job | 
 **request** | [**ImportJobResource**](ImportJobResource.md)| The updated job | 

### Return type

[**ImportJobResource**](ImportJobResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateQuestion**
> QuestionResource UpdateQuestion(ctx, id, optional)
Update a question

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **string**| The id of the question | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the question | 
 **question** | [**QuestionResource**](QuestionResource.md)| The updated question | 

### Return type

[**QuestionResource**](QuestionResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateQuestionAnswer**
> UpdateQuestionAnswer(ctx, questionId, id, optional)
Update an answer for a question

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **questionId** | **string**| The id of the question | 
  **id** | **string**| The id of the answer | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **questionId** | **string**| The id of the question | 
 **id** | **string**| The id of the answer | 
 **answer** | [**AnswerResource**](AnswerResource.md)| The updated answer | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateQuestionTemplate**
> QuestionTemplateResource UpdateQuestionTemplate(ctx, id, optional)
Update a question template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **string**| The id of the template | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the template | 
 **questionTemplateResource** | [**QuestionTemplateResource**](QuestionTemplateResource.md)| The question template resource object | 

### Return type

[**QuestionTemplateResource**](QuestionTemplateResource.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateQuestionsInBulk**
> int32 UpdateQuestionsInBulk(ctx, optional)
Bulk update questions

Will update all questions that match filters used (or all questions in system if no filters used). Body should match a question resource with only those properties you wish to set. Null values will be ignored. Returned number is how many were updated.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **question** | [**QuestionResource**](QuestionResource.md)| New values for a set of question fields | 
 **filterSearch** | **string**| Filter for documents whose question, answers or tags contains provided string | 
 **filterIdset** | **string**| Filter for documents whose id is in the comma separated list provided | 
 **filterCategory** | **string**| Filter for questions with specified category, by id | 
 **filterTagset** | **string**| Filter for questions with specified tags (separated by comma) | 
 **filterType** | **string**| Filter for questions with specified type.  Allowable values: (&#39;TEXT&#39;, &#39;IMAGE&#39;, &#39;VIDEO&#39;, &#39;AUDIO&#39;) | 
 **filterPublished** | **bool**| Filter for questions currenctly published or not | 
 **filterImportId** | **int64**| Filter for questions from a specific import job | 

### Return type

**int32**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

