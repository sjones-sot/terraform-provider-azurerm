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
// ElasticPoolsClient is the the Azure SQL Database management API provides a RESTful set of web services that interact
// with Azure SQL Database services to manage your databases. The API enables you to create, retrieve, update, and
// delete databases.
type ElasticPoolsClient struct {
	BaseClient
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// NewElasticPoolsClient creates an instance of the ElasticPoolsClient client.
func NewElasticPoolsClient(subscriptionID string) ElasticPoolsClient {
	return NewElasticPoolsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// NewElasticPoolsClientWithBaseURI creates an instance of the ElasticPoolsClient client.
func NewElasticPoolsClientWithBaseURI(baseURI string, subscriptionID string) ElasticPoolsClient {
	return ElasticPoolsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// CreateOrUpdate creates a new elastic pool or updates an existing elastic pool.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from
// the Azure Resource Manager API or the portal. serverName is the name of the server. elasticPoolName is the name
// of the elastic pool to be operated on (updated or created). parameters is the required parameters for creating
// or updating an elastic pool.
func (client ElasticPoolsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, parameters ElasticPool) (result ElasticPoolsCreateOrUpdateFuture, err error) {
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, serverName, elasticPoolName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ElasticPoolsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, parameters ElasticPool) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"elasticPoolName":   autorest.Encode("path", elasticPoolName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ElasticPoolsClient) CreateOrUpdateSender(req *http.Request) (future ElasticPoolsCreateOrUpdateFuture, err error) {
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
func (client ElasticPoolsClient) CreateOrUpdateResponder(resp *http.Response) (result ElasticPool, err error) {
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
// Delete deletes the elastic pool.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from
// the Azure Resource Manager API or the portal. serverName is the name of the server. elasticPoolName is the name
// of the elastic pool to be deleted.
func (client ElasticPoolsClient) Delete(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(ctx, resourceGroupName, serverName, elasticPoolName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// DeletePreparer prepares the Delete request.
func (client ElasticPoolsClient) DeletePreparer(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"elasticPoolName":   autorest.Encode("path", elasticPoolName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ElasticPoolsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ElasticPoolsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// Get gets an elastic pool.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from
// the Azure Resource Manager API or the portal. serverName is the name of the server. elasticPoolName is the name
// of the elastic pool to be retrieved.
func (client ElasticPoolsClient) Get(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string) (result ElasticPool, err error) {
	req, err := client.GetPreparer(ctx, resourceGroupName, serverName, elasticPoolName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// GetPreparer prepares the Get request.
func (client ElasticPoolsClient) GetPreparer(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"elasticPoolName":   autorest.Encode("path", elasticPoolName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ElasticPoolsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ElasticPoolsClient) GetResponder(resp *http.Response) (result ElasticPool, err error) {
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
// ListByServer returns a list of elastic pools in a server.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from
// the Azure Resource Manager API or the portal. serverName is the name of the server.
func (client ElasticPoolsClient) ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result ElasticPoolListResult, err error) {
	req, err := client.ListByServerPreparer(ctx, resourceGroupName, serverName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "ListByServer", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByServerSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "ListByServer", resp, "Failure sending request")
		return
	}

	result, err = client.ListByServerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "ListByServer", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// ListByServerPreparer prepares the ListByServer request.
func (client ElasticPoolsClient) ListByServerPreparer(ctx context.Context, resourceGroupName string, serverName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// ListByServerSender sends the ListByServer request. The method will close the
// http.Response Body if it receives an error.
func (client ElasticPoolsClient) ListByServerSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// ListByServerResponder handles the response to the ListByServer request. The method always
// closes the http.Response Body.
func (client ElasticPoolsClient) ListByServerResponder(resp *http.Response) (result ElasticPoolListResult, err error) {
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
// ListMetricDefinitions returns elastic pool metric definitions.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from
// the Azure Resource Manager API or the portal. serverName is the name of the server. elasticPoolName is the name
// of the elastic pool.
func (client ElasticPoolsClient) ListMetricDefinitions(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string) (result MetricDefinitionListResult, err error) {
	req, err := client.ListMetricDefinitionsPreparer(ctx, resourceGroupName, serverName, elasticPoolName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "ListMetricDefinitions", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListMetricDefinitionsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "ListMetricDefinitions", resp, "Failure sending request")
		return
	}

	result, err = client.ListMetricDefinitionsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "ListMetricDefinitions", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// ListMetricDefinitionsPreparer prepares the ListMetricDefinitions request.
func (client ElasticPoolsClient) ListMetricDefinitionsPreparer(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"elasticPoolName":   autorest.Encode("path", elasticPoolName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}/metricDefinitions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// ListMetricDefinitionsSender sends the ListMetricDefinitions request. The method will close the
// http.Response Body if it receives an error.
func (client ElasticPoolsClient) ListMetricDefinitionsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// ListMetricDefinitionsResponder handles the response to the ListMetricDefinitions request. The method always
// closes the http.Response Body.
func (client ElasticPoolsClient) ListMetricDefinitionsResponder(resp *http.Response) (result MetricDefinitionListResult, err error) {
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
// ListMetrics returns elastic pool  metrics.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from
// the Azure Resource Manager API or the portal. serverName is the name of the server. elasticPoolName is the name
// of the elastic pool. filter is an OData filter expression that describes a subset of metrics to return.
func (client ElasticPoolsClient) ListMetrics(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, filter string) (result MetricListResult, err error) {
	req, err := client.ListMetricsPreparer(ctx, resourceGroupName, serverName, elasticPoolName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "ListMetrics", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListMetricsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "ListMetrics", resp, "Failure sending request")
		return
	}

	result, err = client.ListMetricsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "ListMetrics", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// ListMetricsPreparer prepares the ListMetrics request.
func (client ElasticPoolsClient) ListMetricsPreparer(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"elasticPoolName":   autorest.Encode("path", elasticPoolName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"$filter":     autorest.Encode("query", filter),
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}/metrics", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// ListMetricsSender sends the ListMetrics request. The method will close the
// http.Response Body if it receives an error.
func (client ElasticPoolsClient) ListMetricsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// ListMetricsResponder handles the response to the ListMetrics request. The method always
// closes the http.Response Body.
func (client ElasticPoolsClient) ListMetricsResponder(resp *http.Response) (result MetricListResult, err error) {
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
// Update updates an existing elastic pool.
//
// resourceGroupName is the name of the resource group that contains the resource. You can obtain this value from
// the Azure Resource Manager API or the portal. serverName is the name of the server. elasticPoolName is the name
// of the elastic pool to be updated. parameters is the required parameters for updating an elastic pool.
func (client ElasticPoolsClient) Update(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, parameters ElasticPoolUpdate) (result ElasticPoolsUpdateFuture, err error) {
	req, err := client.UpdatePreparer(ctx, resourceGroupName, serverName, elasticPoolName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ElasticPoolsClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// UpdatePreparer prepares the Update request.
func (client ElasticPoolsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, serverName string, elasticPoolName string, parameters ElasticPoolUpdate) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"elasticPoolName":   autorest.Encode("path", elasticPoolName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/elasticPools/{elasticPoolName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql instead.
// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ElasticPoolsClient) UpdateSender(req *http.Request) (future ElasticPoolsUpdateFuture, err error) {
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
func (client ElasticPoolsClient) UpdateResponder(resp *http.Response) (result ElasticPool, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
