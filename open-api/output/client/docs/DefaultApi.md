# \DefaultApi

All URIs are relative to *http://localhost/hello*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TestGet**](DefaultApi.md#TestGet) | **Get** /test | 
[**TestPost**](DefaultApi.md#TestPost) | **Post** /test | 



## TestGet

> TestGetResponse TestGet(ctx, )



Test Get

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**TestGetResponse**](TestGetResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestPost

> TestPostResponse TestPost(ctx, testPost)



Test Post

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**testPost** | [**TestPost**](TestPost.md)| Test Body | 

### Return type

[**TestPostResponse**](TestPostResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

