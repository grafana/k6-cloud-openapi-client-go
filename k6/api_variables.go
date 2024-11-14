/*


HTTP API for interacting with k6 Cloud.

API version: 0.0.0
Contact: info@grafana.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package k6

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// VariablesAPIService VariablesAPI service
type VariablesAPIService service

type ApiVariablesCreateRequest struct {
	ctx                             context.Context
	ApiService                      *VariablesAPIService
	createStackEnvVarRequestRequest *CreateStackEnvVarRequestRequest
}

func (r ApiVariablesCreateRequest) CreateStackEnvVarRequestRequest(createStackEnvVarRequestRequest CreateStackEnvVarRequestRequest) ApiVariablesCreateRequest {
	r.createStackEnvVarRequestRequest = &createStackEnvVarRequestRequest
	return r
}

func (r ApiVariablesCreateRequest) Execute() (*StackEnvVar, *http.Response, error) {
	return r.ApiService.VariablesCreateExecute(r)
}

/*
VariablesCreate Method for VariablesCreate

Create a new environment variable in the stack. Requires admin/owner permissions.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiVariablesCreateRequest
*/
func (a *VariablesAPIService) VariablesCreate(ctx context.Context) ApiVariablesCreateRequest {
	return ApiVariablesCreateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return StackEnvVar
func (a *VariablesAPIService) VariablesCreateExecute(r ApiVariablesCreateRequest) (*StackEnvVar, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *StackEnvVar
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VariablesAPIService.VariablesCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/variables"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createStackEnvVarRequestRequest == nil {
		return localVarReturnValue, nil, reportError("createStackEnvVarRequestRequest is required and must be specified")
	}

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
	localVarPostBody = r.createStackEnvVarRequestRequest
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

type ApiVariablesDestroyRequest struct {
	ctx        context.Context
	ApiService *VariablesAPIService
	xStackId   *string
	id         int32
}

// Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack.
func (r ApiVariablesDestroyRequest) XStackId(xStackId string) ApiVariablesDestroyRequest {
	r.xStackId = &xStackId
	return r
}

func (r ApiVariablesDestroyRequest) Execute() (*http.Response, error) {
	return r.ApiService.VariablesDestroyExecute(r)
}

/*
VariablesDestroy Method for VariablesDestroy

Delete an environment variable. Requires admin/owner permissions.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Id or name of the variable.
	@return ApiVariablesDestroyRequest
*/
func (a *VariablesAPIService) VariablesDestroy(ctx context.Context, id int32) ApiVariablesDestroyRequest {
	return ApiVariablesDestroyRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *VariablesAPIService) VariablesDestroyExecute(r ApiVariablesDestroyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VariablesAPIService.VariablesDestroy")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/variables/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xStackId == nil {
		return nil, reportError("xStackId is required and must be specified")
	}

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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Stack-Id", r.xStackId, "simple", "")
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

type ApiVariablesListRequest struct {
	ctx        context.Context
	ApiService *VariablesAPIService
	xStackId   *string
	count      *bool
	orderby    *string
	skip       *int32
	top        *int32
}

// Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack.
func (r ApiVariablesListRequest) XStackId(xStackId string) ApiVariablesListRequest {
	r.xStackId = &xStackId
	return r
}

// Include collection length in the response object as &#39;@count&#39;.
func (r ApiVariablesListRequest) Count(count bool) ApiVariablesListRequest {
	r.count = &count
	return r
}

// Comma separated list of fields to use when ordering the results. The default order is ascending and can be reversed by appending &#x60;desc&#x60; specifier. Available fields: - name - created - updated
func (r ApiVariablesListRequest) Orderby(orderby string) ApiVariablesListRequest {
	r.orderby = &orderby
	return r
}

// The initial index from which to return the results.
func (r ApiVariablesListRequest) Skip(skip int32) ApiVariablesListRequest {
	r.skip = &skip
	return r
}

// Number of results to return per page.
func (r ApiVariablesListRequest) Top(top int32) ApiVariablesListRequest {
	r.top = &top
	return r
}

func (r ApiVariablesListRequest) Execute() (*PaginatedStackEnvVarList, *http.Response, error) {
	return r.ApiService.VariablesListExecute(r)
}

/*
VariablesList Method for VariablesList

List all environment variables in the stack. The values are accessible only by
the stack owner and admins.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiVariablesListRequest
*/
func (a *VariablesAPIService) VariablesList(ctx context.Context) ApiVariablesListRequest {
	return ApiVariablesListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return PaginatedStackEnvVarList
func (a *VariablesAPIService) VariablesListExecute(r ApiVariablesListRequest) (*PaginatedStackEnvVarList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PaginatedStackEnvVarList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VariablesAPIService.VariablesList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/variables"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xStackId == nil {
		return localVarReturnValue, nil, reportError("xStackId is required and must be specified")
	}

	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "$count", r.count, "form", "")
	}
	if r.orderby != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "$orderby", r.orderby, "form", "")
	} else {
		var defaultValue string = "name"
		r.orderby = &defaultValue
	}
	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "$skip", r.skip, "form", "")
	}
	if r.top != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "$top", r.top, "form", "")
	} else {
		var defaultValue int32 = 1000
		r.top = &defaultValue
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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Stack-Id", r.xStackId, "simple", "")
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

type ApiVariablesPartialUpdateRequest struct {
	ctx                       context.Context
	ApiService                *VariablesAPIService
	xStackId                  *string
	id                        int32
	patchedStackEnvVarRequest *PatchedStackEnvVarRequest
}

// Numeric ID of the Grafana stack representing the request scope. - If the API is called with a *Personal API token*, the user must be a member of the specified stack. - If the API is called with a *Grafana Stack API token*, the value must be the ID of the corresponding stack.
func (r ApiVariablesPartialUpdateRequest) XStackId(xStackId string) ApiVariablesPartialUpdateRequest {
	r.xStackId = &xStackId
	return r
}

func (r ApiVariablesPartialUpdateRequest) PatchedStackEnvVarRequest(patchedStackEnvVarRequest PatchedStackEnvVarRequest) ApiVariablesPartialUpdateRequest {
	r.patchedStackEnvVarRequest = &patchedStackEnvVarRequest
	return r
}

func (r ApiVariablesPartialUpdateRequest) Execute() (*http.Response, error) {
	return r.ApiService.VariablesPartialUpdateExecute(r)
}

/*
VariablesPartialUpdate Method for VariablesPartialUpdate

Update an environment variable. Requires admin/owner permissions.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Id or name of the variable.
	@return ApiVariablesPartialUpdateRequest
*/
func (a *VariablesAPIService) VariablesPartialUpdate(ctx context.Context, id int32) ApiVariablesPartialUpdateRequest {
	return ApiVariablesPartialUpdateRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *VariablesAPIService) VariablesPartialUpdateExecute(r ApiVariablesPartialUpdateRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VariablesAPIService.VariablesPartialUpdate")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/variables/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xStackId == nil {
		return nil, reportError("xStackId is required and must be specified")
	}

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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Stack-Id", r.xStackId, "simple", "")
	// body params
	localVarPostBody = r.patchedStackEnvVarRequest
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
