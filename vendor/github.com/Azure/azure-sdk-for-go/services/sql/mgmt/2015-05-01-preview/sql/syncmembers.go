package sql

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
	"net/http"
)

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// SyncMembersClient is the the Azure SQL Database management API provides a RESTful set of web services that interact
// with Azure SQL Database services to manage your databases. The API enables you to create, retrieve, update, and
// delete databases.
type SyncMembersClient struct {
	BaseClient
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// NewSyncMembersClient creates an instance of the SyncMembersClient client.
func NewSyncMembersClient(subscriptionID string) SyncMembersClient {
	return NewSyncMembersClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// NewSyncMembersClientWithBaseURI creates an instance of the SyncMembersClient client.
func NewSyncMembersClientWithBaseURI(baseURI string, subscriptionID string) SyncMembersClient {
	return SyncMembersClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// CreateOrUpdate creates or updates a sync member.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from
// the Azure Resource Manager API or the portal. serverName is the name of the server. databaseName is the name of
// the database on which the sync group is hosted. syncGroupName is the name of the sync group on which the sync
// member is hosted. syncMemberName is the name of the sync member. parameters is the requested sync member
// resource state.
func (client SyncMembersClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, syncMemberName string, parameters SyncMember) (result SyncMembersCreateOrUpdateFuture, err error) {
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serverName, databaseName, syncGroupName, syncMemberName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncMembersClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncMembersClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client SyncMembersClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, syncMemberName string, parameters SyncMember) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":      autorest.Encode("path", databaseName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"syncGroupName":     autorest.Encode("path", syncGroupName),
		"syncMemberName":    autorest.Encode("path", syncMemberName),
	}

	const APIVersion = "2015-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/syncGroups/{syncGroupName}/syncMembers/{syncMemberName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client SyncMembersClient) CreateOrUpdateSender(req *http.Request) (future SyncMembersCreateOrUpdateFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted))
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client SyncMembersClient) CreateOrUpdateResponder(resp *http.Response) (result SyncMember, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// Delete deletes a sync member.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from
// the Azure Resource Manager API or the portal. serverName is the name of the server. databaseName is the name of
// the database on which the sync group is hosted. syncGroupName is the name of the sync group on which the sync
// member is hosted. syncMemberName is the name of the sync member.
func (client SyncMembersClient) Delete(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, syncMemberName string) (result SyncMembersDeleteFuture, err error) {
	req, err := client.DeletePreparer(ctx, resourceGroupName, serverName, databaseName, syncGroupName, syncMemberName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncMembersClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncMembersClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// DeletePreparer prepares the Delete request.
func (client SyncMembersClient) DeletePreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, syncMemberName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":      autorest.Encode("path", databaseName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"syncGroupName":     autorest.Encode("path", syncGroupName),
		"syncMemberName":    autorest.Encode("path", syncMemberName),
	}

	const APIVersion = "2015-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/syncGroups/{syncGroupName}/syncMembers/{syncMemberName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SyncMembersClient) DeleteSender(req *http.Request) (future SyncMembersDeleteFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent))
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SyncMembersClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// Get gets a sync member.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from
// the Azure Resource Manager API or the portal. serverName is the name of the server. databaseName is the name of
// the database on which the sync group is hosted. syncGroupName is the name of the sync group on which the sync
// member is hosted. syncMemberName is the name of the sync member.
func (client SyncMembersClient) Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, syncMemberName string) (result SyncMember, err error) {
	req, err := client.GetPreparer(ctx, resourceGroupName, serverName, databaseName, syncGroupName, syncMemberName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncMembersClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.SyncMembersClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncMembersClient", "Get", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// GetPreparer prepares the Get request.
func (client SyncMembersClient) GetPreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, syncMemberName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":      autorest.Encode("path", databaseName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"syncGroupName":     autorest.Encode("path", syncGroupName),
		"syncMemberName":    autorest.Encode("path", syncMemberName),
	}

	const APIVersion = "2015-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/syncGroups/{syncGroupName}/syncMembers/{syncMemberName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SyncMembersClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SyncMembersClient) GetResponder(resp *http.Response) (result SyncMember, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// ListBySyncGroup lists sync members in the given sync group.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from
// the Azure Resource Manager API or the portal. serverName is the name of the server. databaseName is the name of
// the database on which the sync group is hosted. syncGroupName is the name of the sync group.
func (client SyncMembersClient) ListBySyncGroup(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string) (result SyncMemberListResultPage, err error) {
	result.fn = client.listBySyncGroupNextResults
	req, err := client.ListBySyncGroupPreparer(ctx, resourceGroupName, serverName, databaseName, syncGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncMembersClient", "ListBySyncGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySyncGroupSender(req)
	if err != nil {
		result.smlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.SyncMembersClient", "ListBySyncGroup", resp, "Failure sending request")
		return
	}

	result.smlr, err = client.ListBySyncGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncMembersClient", "ListBySyncGroup", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// ListBySyncGroupPreparer prepares the ListBySyncGroup request.
func (client SyncMembersClient) ListBySyncGroupPreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":      autorest.Encode("path", databaseName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"syncGroupName":     autorest.Encode("path", syncGroupName),
	}

	const APIVersion = "2015-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/syncGroups/{syncGroupName}/syncMembers", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// ListBySyncGroupSender sends the ListBySyncGroup request. The method will close the
// http.Response Body if it receives an error.
func (client SyncMembersClient) ListBySyncGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// ListBySyncGroupResponder handles the response to the ListBySyncGroup request. The method always
// closes the http.Response Body.
func (client SyncMembersClient) ListBySyncGroupResponder(resp *http.Response) (result SyncMemberListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySyncGroupNextResults retrieves the next set of results, if any.
func (client SyncMembersClient) listBySyncGroupNextResults(lastResults SyncMemberListResult) (result SyncMemberListResult, err error) {
	req, err := lastResults.syncMemberListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.SyncMembersClient", "listBySyncGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySyncGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.SyncMembersClient", "listBySyncGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySyncGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncMembersClient", "listBySyncGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// ListBySyncGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client SyncMembersClient) ListBySyncGroupComplete(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string) (result SyncMemberListResultIterator, err error) {
	result.page, err = client.ListBySyncGroup(ctx, resourceGroupName, serverName, databaseName, syncGroupName)
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// ListMemberSchemas gets a sync member database schema.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from
// the Azure Resource Manager API or the portal. serverName is the name of the server. databaseName is the name of
// the database on which the sync group is hosted. syncGroupName is the name of the sync group on which the sync
// member is hosted. syncMemberName is the name of the sync member.
func (client SyncMembersClient) ListMemberSchemas(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, syncMemberName string) (result SyncFullSchemaPropertiesListResultPage, err error) {
	result.fn = client.listMemberSchemasNextResults
	req, err := client.ListMemberSchemasPreparer(ctx, resourceGroupName, serverName, databaseName, syncGroupName, syncMemberName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncMembersClient", "ListMemberSchemas", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListMemberSchemasSender(req)
	if err != nil {
		result.sfsplr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.SyncMembersClient", "ListMemberSchemas", resp, "Failure sending request")
		return
	}

	result.sfsplr, err = client.ListMemberSchemasResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncMembersClient", "ListMemberSchemas", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// ListMemberSchemasPreparer prepares the ListMemberSchemas request.
func (client SyncMembersClient) ListMemberSchemasPreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, syncMemberName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":      autorest.Encode("path", databaseName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"syncGroupName":     autorest.Encode("path", syncGroupName),
		"syncMemberName":    autorest.Encode("path", syncMemberName),
	}

	const APIVersion = "2015-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/syncGroups/{syncGroupName}/syncMembers/{syncMemberName}/schemas", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// ListMemberSchemasSender sends the ListMemberSchemas request. The method will close the
// http.Response Body if it receives an error.
func (client SyncMembersClient) ListMemberSchemasSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// ListMemberSchemasResponder handles the response to the ListMemberSchemas request. The method always
// closes the http.Response Body.
func (client SyncMembersClient) ListMemberSchemasResponder(resp *http.Response) (result SyncFullSchemaPropertiesListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listMemberSchemasNextResults retrieves the next set of results, if any.
func (client SyncMembersClient) listMemberSchemasNextResults(lastResults SyncFullSchemaPropertiesListResult) (result SyncFullSchemaPropertiesListResult, err error) {
	req, err := lastResults.syncFullSchemaPropertiesListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.SyncMembersClient", "listMemberSchemasNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListMemberSchemasSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.SyncMembersClient", "listMemberSchemasNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListMemberSchemasResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncMembersClient", "listMemberSchemasNextResults", resp, "Failure responding to next results request")
	}
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// ListMemberSchemasComplete enumerates all values, automatically crossing page boundaries as required.
func (client SyncMembersClient) ListMemberSchemasComplete(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, syncMemberName string) (result SyncFullSchemaPropertiesListResultIterator, err error) {
	result.page, err = client.ListMemberSchemas(ctx, resourceGroupName, serverName, databaseName, syncGroupName, syncMemberName)
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// RefreshMemberSchema refreshes a sync member database schema.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from
// the Azure Resource Manager API or the portal. serverName is the name of the server. databaseName is the name of
// the database on which the sync group is hosted. syncGroupName is the name of the sync group on which the sync
// member is hosted. syncMemberName is the name of the sync member.
func (client SyncMembersClient) RefreshMemberSchema(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, syncMemberName string) (result SyncMembersRefreshMemberSchemaFuture, err error) {
	req, err := client.RefreshMemberSchemaPreparer(ctx, resourceGroupName, serverName, databaseName, syncGroupName, syncMemberName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncMembersClient", "RefreshMemberSchema", nil, "Failure preparing request")
		return
	}

	result, err = client.RefreshMemberSchemaSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncMembersClient", "RefreshMemberSchema", result.Response(), "Failure sending request")
		return
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// RefreshMemberSchemaPreparer prepares the RefreshMemberSchema request.
func (client SyncMembersClient) RefreshMemberSchemaPreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, syncMemberName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":      autorest.Encode("path", databaseName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"syncGroupName":     autorest.Encode("path", syncGroupName),
		"syncMemberName":    autorest.Encode("path", syncMemberName),
	}

	const APIVersion = "2015-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/syncGroups/{syncGroupName}/syncMembers/{syncMemberName}/refreshSchema", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// RefreshMemberSchemaSender sends the RefreshMemberSchema request. The method will close the
// http.Response Body if it receives an error.
func (client SyncMembersClient) RefreshMemberSchemaSender(req *http.Request) (future SyncMembersRefreshMemberSchemaFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted))
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// RefreshMemberSchemaResponder handles the response to the RefreshMemberSchema request. The method always
// closes the http.Response Body.
func (client SyncMembersClient) RefreshMemberSchemaResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// Update updates an existing sync member.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from
// the Azure Resource Manager API or the portal. serverName is the name of the server. databaseName is the name of
// the database on which the sync group is hosted. syncGroupName is the name of the sync group on which the sync
// member is hosted. syncMemberName is the name of the sync member. parameters is the requested sync member
// resource state.
func (client SyncMembersClient) Update(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, syncMemberName string, parameters SyncMember) (result SyncMembersUpdateFuture, err error) {
	req, err := client.UpdatePreparer(ctx, resourceGroupName, serverName, databaseName, syncGroupName, syncMemberName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncMembersClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.SyncMembersClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// UpdatePreparer prepares the Update request.
func (client SyncMembersClient) UpdatePreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string, syncGroupName string, syncMemberName string, parameters SyncMember) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":      autorest.Encode("path", databaseName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"syncGroupName":     autorest.Encode("path", syncGroupName),
		"syncMemberName":    autorest.Encode("path", syncMemberName),
	}

	const APIVersion = "2015-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/syncGroups/{syncGroupName}/syncMembers/{syncMemberName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client SyncMembersClient) UpdateSender(req *http.Request) (future SyncMembersUpdateFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted))
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client SyncMembersClient) UpdateResponder(resp *http.Response) (result SyncMember, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
