# ImportJobResource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryId** | **string** | The id of the category to assign all questions in the import to | [default to null]
**CreatedDate** | **int64** | The date the job was created in seconds since unix epoc | [optional] [default to null]
**Id** | **int64** | The id of the job | [optional] [default to null]
**Name** | **string** | A name for this import for later reference | [default to null]
**Output** | [**[]ImportJobOutputResource**](ImportJobOutputResource.md) | Error information from validation | [optional] [default to null]
**RecordCount** | **int64** | The number of questions form the CSV file. Filled in after validation | [optional] [default to null]
**Status** | **string** | The status of the job | [optional] [default to null]
**UpdatedDate** | **int64** | The date the job was last updated in seconds since unix epoc | [optional] [default to null]
**Url** | **string** | The url of a CSV file to pull trivia questions from. Cannot be changed after initial POST | [default to null]
**Vendor** | **string** | The vendor who supplied this set of questions | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


