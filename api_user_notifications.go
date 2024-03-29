/*
FastReport Cloud

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gofrcloud

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// UserNotificationsApiService UserNotificationsApi service
type UserNotificationsApiService service

type ApiUserNotificationsClearNotificationsRequest struct {
	ctx context.Context
	ApiService *UserNotificationsApiService
	clearNotificationsVM *ClearNotificationsVM
}

// 
func (r ApiUserNotificationsClearNotificationsRequest) ClearNotificationsVM(clearNotificationsVM ClearNotificationsVM) ApiUserNotificationsClearNotificationsRequest {
	r.clearNotificationsVM = &clearNotificationsVM
	return r
}

func (r ApiUserNotificationsClearNotificationsRequest) Execute() (*http.Response, error) {
	return r.ApiService.UserNotificationsClearNotificationsExecute(r)
}

/*
UserNotificationsClearNotifications Use this endpoint to \"clear\" your notifications

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUserNotificationsClearNotificationsRequest
*/
func (a *UserNotificationsApiService) UserNotificationsClearNotifications(ctx context.Context) ApiUserNotificationsClearNotificationsRequest {
	return ApiUserNotificationsClearNotificationsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *UserNotificationsApiService) UserNotificationsClearNotificationsExecute(r ApiUserNotificationsClearNotificationsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserNotificationsApiService.UserNotificationsClearNotifications")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/manage/v1/notifications"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "text/json", "application/*+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.clearNotificationsVM
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiUserNotificationsGetNotificationsRequest struct {
	ctx context.Context
	ApiService *UserNotificationsApiService
	skip *int32
	take *int32
	subscriptionId *string
}

func (r ApiUserNotificationsGetNotificationsRequest) Skip(skip int32) ApiUserNotificationsGetNotificationsRequest {
	r.skip = &skip
	return r
}

func (r ApiUserNotificationsGetNotificationsRequest) Take(take int32) ApiUserNotificationsGetNotificationsRequest {
	r.take = &take
	return r
}

func (r ApiUserNotificationsGetNotificationsRequest) SubscriptionId(subscriptionId string) ApiUserNotificationsGetNotificationsRequest {
	r.subscriptionId = &subscriptionId
	return r
}

func (r ApiUserNotificationsGetNotificationsRequest) Execute() (*AuditActionsVM, *http.Response, error) {
	return r.ApiService.UserNotificationsGetNotificationsExecute(r)
}

/*
UserNotificationsGetNotifications Use this endpoint to recieve notifications

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUserNotificationsGetNotificationsRequest
*/
func (a *UserNotificationsApiService) UserNotificationsGetNotifications(ctx context.Context) ApiUserNotificationsGetNotificationsRequest {
	return ApiUserNotificationsGetNotificationsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AuditActionsVM
func (a *UserNotificationsApiService) UserNotificationsGetNotificationsExecute(r ApiUserNotificationsGetNotificationsRequest) (*AuditActionsVM, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AuditActionsVM
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserNotificationsApiService.UserNotificationsGetNotifications")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/manage/v1/notifications"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.take != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "take", r.take, "")
	}
	if r.subscriptionId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "subscriptionId", r.subscriptionId, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
