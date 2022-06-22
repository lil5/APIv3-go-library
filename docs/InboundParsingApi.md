# \InboundParsingApi

All URIs are relative to *https://api.sendinblue.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInboundEmailEvents**](InboundParsingApi.md#GetInboundEmailEvents) | **Get** /inbound/events | Get the list of all the events for the received emails.
[**GetInboundEmailEventsByUuid**](InboundParsingApi.md#GetInboundEmailEventsByUuid) | **Get** /inbound/events/{uuid} | Fetch all events history for one particular received email.


# **GetInboundEmailEvents**
> GetInboundEmailEvents GetInboundEmailEvents(ctx, optional)
Get the list of all the events for the received emails.

This endpoint will show the list of all the events for the received emails.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetInboundEmailEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetInboundEmailEventsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sender** | **optional.String**| Email address of the sender. | 
 **startDate** | **optional.String**| Mandatory if endDate is used. Starting date (YYYY-MM-DD) from which you want to fetch the list. Maximum time period that can be selected is one month. | 
 **endDate** | **optional.String**| Mandatory if startDate is used. Ending date (YYYY-MM-DD) till which you want to fetch the list. Maximum time period that can be selected is one month. | 
 **limit** | **optional.Int64**| Number of documents returned per page | [default to 100]
 **offset** | **optional.Int64**| Index of the first document on the page | [default to 0]
 **sort** | **optional.String**| Sort the results in the ascending/descending order of record creation | [default to desc]

### Return type

[**GetInboundEmailEvents**](GetInboundEmailEvents.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInboundEmailEventsByUuid**
> GetInboundEmailEventsByUuid GetInboundEmailEventsByUuid(ctx, uuid)
Fetch all events history for one particular received email.

This endpoint will show the list of all events history for one particular received email.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | **string**| UUID to fetch events specific to recieved email | 

### Return type

[**GetInboundEmailEventsByUuid**](GetInboundEmailEventsByUuid.md)

### Authorization

[api-key](../README.md#api-key), [partner-key](../README.md#partner-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

