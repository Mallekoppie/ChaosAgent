# \DefaultApi

All URIs are relative to *http://localhost/Integration/ChaosAgent/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTests**](DefaultApi.md#AddTests) | **Put** /configure | Configure the tests that will be executed
[**GetTestStatus**](DefaultApi.md#GetTestStatus) | **Get** /test | Get test run status
[**IsAlive**](DefaultApi.md#IsAlive) | **Get** /health | Synthetic transaction
[**StartTestRun**](DefaultApi.md#StartTestRun) | **Post** /test | Start a new test run
[**StopTestRun**](DefaultApi.md#StopTestRun) | **Delete** /test | Stop a test run
[**UpdateTestRun**](DefaultApi.md#UpdateTestRun) | **Put** /test | Change the parameters of a test that is being executed


# **AddTests**
> AddTests(ctx, testCollection)
Configure the tests that will be executed

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **testCollection** | [**TestCollection**](TestCollection.md)| Pet object that needs to be added to the store | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTestStatus**
> TestStatus GetTestStatus(ctx, )
Get test run status

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**TestStatus**](TestStatus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IsAlive**
> IsAlive(ctx, )
Synthetic transaction

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartTestRun**
> StartTestRun(ctx, testParameters)
Start a new test run

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **testParameters** | [**TestParameters**](TestParameters.md)| Pet object that needs to be added to the store | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StopTestRun**
> StopTestRun(ctx, testName)
Stop a test run

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **testName** | **string**| Name of the test run that must be stopped | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTestRun**
> UpdateTestRun(ctx, optional)
Change the parameters of a test that is being executed

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **testParameters** | [**TestParameters**](TestParameters.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

