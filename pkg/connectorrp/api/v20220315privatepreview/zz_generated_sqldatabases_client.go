//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package v20220315privatepreview

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// SQLDatabasesClient contains the methods for the SQLDatabases group.
// Don't use this type directly, use NewSQLDatabasesClient() instead.
type SQLDatabasesClient struct {
	con *connection
	subscriptionID string
}

// NewSQLDatabasesClient creates a new instance of SQLDatabasesClient with the specified values.
func NewSQLDatabasesClient(con *connection, subscriptionID string) *SQLDatabasesClient {
	return &SQLDatabasesClient{con: con, subscriptionID: subscriptionID}
}

// CreateOrUpdate - Creates or updates a SQLDatabase resource
// If the operation fails it returns the *ErrorResponse error type.
func (client *SQLDatabasesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, sqlDatabaseName string, sqlDatabaseParameters SQLDatabaseResource, options *SQLDatabasesCreateOrUpdateOptions) (SQLDatabasesCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, sqlDatabaseName, sqlDatabaseParameters, options)
	if err != nil {
		return SQLDatabasesCreateOrUpdateResponse{}, err
	}
	resp, err := 	client.con.Pipeline().Do(req)
	if err != nil {
		return SQLDatabasesCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return SQLDatabasesCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SQLDatabasesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, sqlDatabaseName string, sqlDatabaseParameters SQLDatabaseResource, options *SQLDatabasesCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Applications.Connector/sqlDatabases/{sqlDatabaseName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlDatabaseName == "" {
		return nil, errors.New("parameter sqlDatabaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlDatabaseName}", url.PathEscape(sqlDatabaseName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, sqlDatabaseParameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *SQLDatabasesClient) createOrUpdateHandleResponse(resp *http.Response) (SQLDatabasesCreateOrUpdateResponse, error) {
	result := SQLDatabasesCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLDatabaseResource); err != nil {
		return SQLDatabasesCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *SQLDatabasesClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
		errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Delete - Deletes an existing sqlDatabase resource
// If the operation fails it returns the *ErrorResponse error type.
func (client *SQLDatabasesClient) Delete(ctx context.Context, resourceGroupName string, sqlDatabaseName string, options *SQLDatabasesDeleteOptions) (SQLDatabasesDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, sqlDatabaseName, options)
	if err != nil {
		return SQLDatabasesDeleteResponse{}, err
	}
	resp, err := 	client.con.Pipeline().Do(req)
	if err != nil {
		return SQLDatabasesDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return SQLDatabasesDeleteResponse{}, client.deleteHandleError(resp)
	}
	return SQLDatabasesDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SQLDatabasesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, sqlDatabaseName string, options *SQLDatabasesDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Applications.Connector/sqlDatabases/{sqlDatabaseName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlDatabaseName == "" {
		return nil, errors.New("parameter sqlDatabaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlDatabaseName}", url.PathEscape(sqlDatabaseName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *SQLDatabasesClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
		errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Retrieves information about a sqlDatabase resource
// If the operation fails it returns the *ErrorResponse error type.
func (client *SQLDatabasesClient) Get(ctx context.Context, resourceGroupName string, sqlDatabaseName string, options *SQLDatabasesGetOptions) (SQLDatabasesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, sqlDatabaseName, options)
	if err != nil {
		return SQLDatabasesGetResponse{}, err
	}
	resp, err := 	client.con.Pipeline().Do(req)
	if err != nil {
		return SQLDatabasesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SQLDatabasesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SQLDatabasesClient) getCreateRequest(ctx context.Context, resourceGroupName string, sqlDatabaseName string, options *SQLDatabasesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Applications.Connector/sqlDatabases/{sqlDatabaseName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlDatabaseName == "" {
		return nil, errors.New("parameter sqlDatabaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlDatabaseName}", url.PathEscape(sqlDatabaseName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SQLDatabasesClient) getHandleResponse(resp *http.Response) (SQLDatabasesGetResponse, error) {
	result := SQLDatabasesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLDatabaseResource); err != nil {
		return SQLDatabasesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *SQLDatabasesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
		errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Lists information about all sqlDatabase resources in the given subscription and resource group
// If the operation fails it returns the *ErrorResponse error type.
func (client *SQLDatabasesClient) List(resourceGroupName string, options *SQLDatabasesListOptions) (*SQLDatabasesListPager) {
	return &SQLDatabasesListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp SQLDatabasesListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SQLDatabaseList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *SQLDatabasesClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *SQLDatabasesListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Applications.Connector/sqlDatabases"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SQLDatabasesClient) listHandleResponse(resp *http.Response) (SQLDatabasesListResponse, error) {
	result := SQLDatabasesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLDatabaseList); err != nil {
		return SQLDatabasesListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *SQLDatabasesClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
		errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListBySubscription - Lists information about all sqlDatabase resources in the given subscription
// If the operation fails it returns the *ErrorResponse error type.
func (client *SQLDatabasesClient) ListBySubscription(options *SQLDatabasesListBySubscriptionOptions) (*SQLDatabasesListBySubscriptionPager) {
	return &SQLDatabasesListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp SQLDatabasesListBySubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SQLDatabaseList.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *SQLDatabasesClient) listBySubscriptionCreateRequest(ctx context.Context, options *SQLDatabasesListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Applications.Connector/sqlDatabases"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-15-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *SQLDatabasesClient) listBySubscriptionHandleResponse(resp *http.Response) (SQLDatabasesListBySubscriptionResponse, error) {
	result := SQLDatabasesListBySubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLDatabaseList); err != nil {
		return SQLDatabasesListBySubscriptionResponse{}, err
	}
	return result, nil
}

// listBySubscriptionHandleError handles the ListBySubscription error response.
func (client *SQLDatabasesClient) listBySubscriptionHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
		errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

