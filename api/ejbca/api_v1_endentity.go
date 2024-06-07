/*
Copyright 2024 Keyfactor

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

EJBCA REST Interface

API reference documentation.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ejbca

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// V1EndentityApiService V1EndentityApi service
type V1EndentityApiService service

type ApiAddRequest struct {
	ctx                     context.Context
	ApiService              *V1EndentityApiService
	addEndEntityRestRequest *AddEndEntityRestRequest
}

// request
func (r ApiAddRequest) AddEndEntityRestRequest(addEndEntityRestRequest AddEndEntityRestRequest) ApiAddRequest {
	r.addEndEntityRestRequest = &addEndEntityRestRequest
	return r
}

func (r ApiAddRequest) Execute() (*http.Response, error) {
	return r.ApiService.AddExecute(r)
}

/*
Add Add new end entity, if it does not exist

Register new end entity based on provided registration data

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddRequest
*/
func (a *V1EndentityApiService) Add(ctx context.Context) ApiAddRequest {
	return ApiAddRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *V1EndentityApiService) AddExecute(r ApiAddRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath := "/ejbca/ejbca-rest-api"

	localVarPath := localBasePath + "/v1/endentity"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.addEndEntityRestRequest
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
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiDeleteRequest struct {
	ctx           context.Context
	ApiService    *V1EndentityApiService
	endentityName string
}

func (r ApiDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteExecute(r)
}

/*
Delete Deletes end entity

Deletes specified end entity and keeps certificate information untouched, if end entity does not exist success is still returned

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param endentityName Name of the end entity
	@return ApiDeleteRequest
*/
func (a *V1EndentityApiService) Delete(ctx context.Context, endentityName string) ApiDeleteRequest {
	return ApiDeleteRequest{
		ApiService:    a,
		ctx:           ctx,
		endentityName: endentityName,
	}
}

// Execute executes the request
func (a *V1EndentityApiService) DeleteExecute(r ApiDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath := "/ejbca/ejbca-rest-api"

	localVarPath := localBasePath + "/v1/endentity/{endentity_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"endentity_name"+"}", url.PathEscape(parameterValueToString(r.endentityName, "endentityName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
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
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiRevokeRequest struct {
	ctx                            context.Context
	ApiService                     *V1EndentityApiService
	endentityName                  string
	endEntityRevocationRestRequest *EndEntityRevocationRestRequest
}

// request
func (r ApiRevokeRequest) EndEntityRevocationRestRequest(endEntityRevocationRestRequest EndEntityRevocationRestRequest) ApiRevokeRequest {
	r.endEntityRevocationRestRequest = &endEntityRevocationRestRequest
	return r
}

func (r ApiRevokeRequest) Execute() (*http.Response, error) {
	return r.ApiService.RevokeExecute(r)
}

/*
Revoke Revokes all end entity certificates

Revokes all certificates associated with given end entity name with specified reason code (see RFC 5280 Section 5.3.1), and optionally deletes the end entity

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param endentityName Name of the end entity
	@return ApiRevokeRequest
*/
func (a *V1EndentityApiService) Revoke(ctx context.Context, endentityName string) ApiRevokeRequest {
	return ApiRevokeRequest{
		ApiService:    a,
		ctx:           ctx,
		endentityName: endentityName,
	}
}

// Execute executes the request
func (a *V1EndentityApiService) RevokeExecute(r ApiRevokeRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath := "/ejbca/ejbca-rest-api"

	localVarPath := localBasePath + "/v1/endentity/{endentity_name}/revoke"
	localVarPath = strings.Replace(localVarPath, "{"+"endentity_name"+"}", url.PathEscape(parameterValueToString(r.endentityName, "endentityName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.endEntityRevocationRestRequest
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
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiSearchRequest struct {
	ctx                          context.Context
	ApiService                   *V1EndentityApiService
	searchEndEntitiesRestRequest *SearchEndEntitiesRestRequest
}

// Maximum number of results and collection of search criterias.
func (r ApiSearchRequest) SearchEndEntitiesRestRequest(searchEndEntitiesRestRequest SearchEndEntitiesRestRequest) ApiSearchRequest {
	r.searchEndEntitiesRestRequest = &searchEndEntitiesRestRequest
	return r
}

func (r ApiSearchRequest) Execute() (*SearchEndEntitiesRestResponse, *http.Response, error) {
	return r.ApiService.SearchExecute(r)
}

/*
Search Searches for end entity confirming given criteria.

Insert as many search criteria as needed. A reference about allowed values for criteria could be found below, under SearchEndEntityCriteriaRestRequest model.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiSearchRequest
*/
func (a *V1EndentityApiService) Search(ctx context.Context) ApiSearchRequest {
	return ApiSearchRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SearchEndEntitiesRestResponse
func (a *V1EndentityApiService) SearchExecute(r ApiSearchRequest) (*SearchEndEntitiesRestResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SearchEndEntitiesRestResponse
	)

	localBasePath := "/ejbca/ejbca-rest-api"

	localVarPath := localBasePath + "/v1/endentity/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	localVarPostBody = r.searchEndEntitiesRestRequest
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

type ApiSetstatusRequest struct {
	ctx                           context.Context
	ApiService                    *V1EndentityApiService
	endentityName                 string
	setEndEntityStatusRestRequest *SetEndEntityStatusRestRequest
}

// request
func (r ApiSetstatusRequest) SetEndEntityStatusRestRequest(setEndEntityStatusRestRequest SetEndEntityStatusRestRequest) ApiSetstatusRequest {
	r.setEndEntityStatusRestRequest = &setEndEntityStatusRestRequest
	return r
}

func (r ApiSetstatusRequest) Execute() (*http.Response, error) {
	return r.ApiService.SetstatusExecute(r)
}

/*
Setstatus Edits end entity setting new status

Edit status, password and token type of related end entity

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param endentityName Name of the end entity to edit status for
	@return ApiSetstatusRequest
*/
func (a *V1EndentityApiService) Setstatus(ctx context.Context, endentityName string) ApiSetstatusRequest {
	return ApiSetstatusRequest{
		ApiService:    a,
		ctx:           ctx,
		endentityName: endentityName,
	}
}

// Execute executes the request
func (a *V1EndentityApiService) SetstatusExecute(r ApiSetstatusRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath := "/ejbca/ejbca-rest-api"

	localVarPath := localBasePath + "/v1/endentity/{endentity_name}/setstatus"
	localVarPath = strings.Replace(localVarPath, "{"+"endentity_name"+"}", url.PathEscape(parameterValueToString(r.endentityName, "endentityName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.setEndEntityStatusRestRequest
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
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiStatus6Request struct {
	ctx        context.Context
	ApiService *V1EndentityApiService
}

func (r ApiStatus6Request) Execute() (*RestResourceStatusRestResponse, *http.Response, error) {
	return r.ApiService.Status6Execute(r)
}

/*
Status6 Get the status of this REST Resource

Returns status, API version and EJBCA version.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiStatus6Request
*/
func (a *V1EndentityApiService) Status6(ctx context.Context) ApiStatus6Request {
	return ApiStatus6Request{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return RestResourceStatusRestResponse
func (a *V1EndentityApiService) Status6Execute(r ApiStatus6Request) (*RestResourceStatusRestResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *RestResourceStatusRestResponse
	)

	localBasePath := "/ejbca/ejbca-rest-api"

	localVarPath := localBasePath + "/v1/endentity/status"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
