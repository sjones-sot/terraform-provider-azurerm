package graphrbac

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// ObjectsClient is the the Graph RBAC Management Client
type ObjectsClient struct {
	BaseClient
}

// NewObjectsClient creates an instance of the ObjectsClient client.
func NewObjectsClient(tenantID string) ObjectsClient {
	return NewObjectsClientWithBaseURI(DefaultBaseURI, tenantID)
}

// NewObjectsClientWithBaseURI creates an instance of the ObjectsClient client.
func NewObjectsClientWithBaseURI(baseURI string, tenantID string) ObjectsClient {
	return ObjectsClient{NewWithBaseURI(baseURI, tenantID)}
}

// GetCurrentUser gets the details for the currently logged-in user.
func (client ObjectsClient) GetCurrentUser(ctx context.Context) (result AADObject, err error) {
	req, err := client.GetCurrentUserPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ObjectsClient", "GetCurrentUser", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetCurrentUserSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "graphrbac.ObjectsClient", "GetCurrentUser", resp, "Failure sending request")
		return
	}

	result, err = client.GetCurrentUserResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ObjectsClient", "GetCurrentUser", resp, "Failure responding to request")
	}

	return
}

// GetCurrentUserPreparer prepares the GetCurrentUser request.
func (client ObjectsClient) GetCurrentUserPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"tenantID": autorest.Encode("path", client.TenantID),
	}

	const APIVersion = "1.6"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{tenantID}/me", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetCurrentUserSender sends the GetCurrentUser request. The method will close the
// http.Response Body if it receives an error.
func (client ObjectsClient) GetCurrentUserSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetCurrentUserResponder handles the response to the GetCurrentUser request. The method always
// closes the http.Response Body.
func (client ObjectsClient) GetCurrentUserResponder(resp *http.Response) (result AADObject, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetObjectsByObjectIds gets AD group membership for the specified AD object IDs.
// Parameters:
// parameters - objects filtering parameters.
func (client ObjectsClient) GetObjectsByObjectIds(ctx context.Context, parameters GetObjectsParameters) (result GetObjectsResultPage, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.IncludeDirectoryObjectReferences", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("graphrbac.ObjectsClient", "GetObjectsByObjectIds", err.Error())
	}

	result.fn = func(lastResult GetObjectsResult) (GetObjectsResult, error) {
		if lastResult.OdataNextLink == nil || len(to.String(lastResult.OdataNextLink)) < 1 {
			return GetObjectsResult{}, nil
		}
		return client.GetObjectsByObjectIdsNext(ctx, *lastResult.OdataNextLink)
	}
	req, err := client.GetObjectsByObjectIdsPreparer(ctx, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ObjectsClient", "GetObjectsByObjectIds", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetObjectsByObjectIdsSender(req)
	if err != nil {
		result.gor.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "graphrbac.ObjectsClient", "GetObjectsByObjectIds", resp, "Failure sending request")
		return
	}

	result.gor, err = client.GetObjectsByObjectIdsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ObjectsClient", "GetObjectsByObjectIds", resp, "Failure responding to request")
	}

	return
}

// GetObjectsByObjectIdsPreparer prepares the GetObjectsByObjectIds request.
func (client ObjectsClient) GetObjectsByObjectIdsPreparer(ctx context.Context, parameters GetObjectsParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"tenantID": autorest.Encode("path", client.TenantID),
	}

	const APIVersion = "1.6"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{tenantID}/getObjectsByObjectIds", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetObjectsByObjectIdsSender sends the GetObjectsByObjectIds request. The method will close the
// http.Response Body if it receives an error.
func (client ObjectsClient) GetObjectsByObjectIdsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetObjectsByObjectIdsResponder handles the response to the GetObjectsByObjectIds request. The method always
// closes the http.Response Body.
func (client ObjectsClient) GetObjectsByObjectIdsResponder(resp *http.Response) (result GetObjectsResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetObjectsByObjectIdsComplete enumerates all values, automatically crossing page boundaries as required.
func (client ObjectsClient) GetObjectsByObjectIdsComplete(ctx context.Context, parameters GetObjectsParameters) (result GetObjectsResultIterator, err error) {
	result.page, err = client.GetObjectsByObjectIds(ctx, parameters)
	return
}

// GetObjectsByObjectIdsNext gets AD group membership for the specified AD object IDs.
// Parameters:
// nextLink - next link for the list operation.
func (client ObjectsClient) GetObjectsByObjectIdsNext(ctx context.Context, nextLink string) (result GetObjectsResult, err error) {
	req, err := client.GetObjectsByObjectIdsNextPreparer(ctx, nextLink)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ObjectsClient", "GetObjectsByObjectIdsNext", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetObjectsByObjectIdsNextSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "graphrbac.ObjectsClient", "GetObjectsByObjectIdsNext", resp, "Failure sending request")
		return
	}

	result, err = client.GetObjectsByObjectIdsNextResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "graphrbac.ObjectsClient", "GetObjectsByObjectIdsNext", resp, "Failure responding to request")
	}

	return
}

// GetObjectsByObjectIdsNextPreparer prepares the GetObjectsByObjectIdsNext request.
func (client ObjectsClient) GetObjectsByObjectIdsNextPreparer(ctx context.Context, nextLink string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"nextLink": nextLink,
		"tenantID": autorest.Encode("path", client.TenantID),
	}

	const APIVersion = "1.6"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{tenantID}/{nextLink}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetObjectsByObjectIdsNextSender sends the GetObjectsByObjectIdsNext request. The method will close the
// http.Response Body if it receives an error.
func (client ObjectsClient) GetObjectsByObjectIdsNextSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetObjectsByObjectIdsNextResponder handles the response to the GetObjectsByObjectIdsNext request. The method always
// closes the http.Response Body.
func (client ObjectsClient) GetObjectsByObjectIdsNextResponder(resp *http.Response) (result GetObjectsResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
