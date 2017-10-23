# AmazonS3Activity

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | S3 action (i.e., &#39;PUT&#39;) associated with the activity | [optional] [default to null]
**CdnUrl** | **string** | URL for accessing the resource. Will use a CDN if configured, or direct to S3 if not | [optional] [default to null]
**CreatedDate** | **int64** | Date the resource was created in S3 | [optional] [default to null]
**Filename** | **string** | Name of the file being processed as a resource in S3 | [optional] [default to null]
**Id** | **int64** | Unique id of the S3 activity | [optional] [default to null]
**ObjectKey** | **string** | S3 object key for the resource | [optional] [default to null]
**Url** | **string** | URL that one can PUT the file to, to upload it to S3 | [optional] [default to null]
**UserId** | **int32** | The id of the user that created this S3 activity | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


