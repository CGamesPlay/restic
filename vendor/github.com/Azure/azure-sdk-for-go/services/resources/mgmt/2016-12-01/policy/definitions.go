package policy

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
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// DefinitionsClient is the to manage and control access to your resources, you can define customized policies and
// assign them at a scope.
type DefinitionsClient struct {
	ManagementClient
}

// NewDefinitionsClient creates an instance of the DefinitionsClient client.
func NewDefinitionsClient(subscriptionID string) DefinitionsClient {
	return NewDefinitionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDefinitionsClientWithBaseURI creates an instance of the DefinitionsClient client.
func NewDefinitionsClientWithBaseURI(baseURI string, subscriptionID string) DefinitionsClient {
	return DefinitionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a policy definition.
//
// policyDefinitionName is the name of the policy definition to create. parameters is the policy definition properties.
func (client DefinitionsClient) CreateOrUpdate(policyDefinitionName string, parameters Definition) (result Definition, err error) {
	req, err := client.CreateOrUpdatePreparer(policyDefinitionName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client DefinitionsClient) CreateOrUpdatePreparer(policyDefinitionName string, parameters Definition) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policyDefinitionName": autorest.Encode("path", policyDefinitionName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client DefinitionsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client DefinitionsClient) CreateOrUpdateResponder(resp *http.Response) (result Definition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateOrUpdateAtManagementGroup creates or updates a policy definition at management group level.
//
// policyDefinitionName is the name of the policy definition to create. parameters is the policy definition properties.
// managementGroupID is the ID of the management group.
func (client DefinitionsClient) CreateOrUpdateAtManagementGroup(policyDefinitionName string, parameters Definition, managementGroupID string) (result Definition, err error) {
	req, err := client.CreateOrUpdateAtManagementGroupPreparer(policyDefinitionName, parameters, managementGroupID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "CreateOrUpdateAtManagementGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateAtManagementGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "CreateOrUpdateAtManagementGroup", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateAtManagementGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "CreateOrUpdateAtManagementGroup", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdateAtManagementGroupPreparer prepares the CreateOrUpdateAtManagementGroup request.
func (client DefinitionsClient) CreateOrUpdateAtManagementGroupPreparer(policyDefinitionName string, parameters Definition, managementGroupID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managementGroupId":    autorest.Encode("path", managementGroupID),
		"policyDefinitionName": autorest.Encode("path", policyDefinitionName),
	}

	const APIVersion = "2016-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Management/managementgroups/{managementGroupId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CreateOrUpdateAtManagementGroupSender sends the CreateOrUpdateAtManagementGroup request. The method will close the
// http.Response Body if it receives an error.
func (client DefinitionsClient) CreateOrUpdateAtManagementGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateOrUpdateAtManagementGroupResponder handles the response to the CreateOrUpdateAtManagementGroup request. The method always
// closes the http.Response Body.
func (client DefinitionsClient) CreateOrUpdateAtManagementGroupResponder(resp *http.Response) (result Definition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a policy definition.
//
// policyDefinitionName is the name of the policy definition to delete.
func (client DefinitionsClient) Delete(policyDefinitionName string) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(policyDefinitionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client DefinitionsClient) DeletePreparer(policyDefinitionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policyDefinitionName": autorest.Encode("path", policyDefinitionName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client DefinitionsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client DefinitionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// DeleteAtManagementGroup deletes a policy definition at management group level.
//
// policyDefinitionName is the name of the policy definition to delete. managementGroupID is the ID of the management
// group.
func (client DefinitionsClient) DeleteAtManagementGroup(policyDefinitionName string, managementGroupID string) (result autorest.Response, err error) {
	req, err := client.DeleteAtManagementGroupPreparer(policyDefinitionName, managementGroupID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "DeleteAtManagementGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteAtManagementGroupSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "DeleteAtManagementGroup", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteAtManagementGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "DeleteAtManagementGroup", resp, "Failure responding to request")
	}

	return
}

// DeleteAtManagementGroupPreparer prepares the DeleteAtManagementGroup request.
func (client DefinitionsClient) DeleteAtManagementGroupPreparer(policyDefinitionName string, managementGroupID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managementGroupId":    autorest.Encode("path", managementGroupID),
		"policyDefinitionName": autorest.Encode("path", policyDefinitionName),
	}

	const APIVersion = "2016-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Management/managementgroups/{managementGroupId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// DeleteAtManagementGroupSender sends the DeleteAtManagementGroup request. The method will close the
// http.Response Body if it receives an error.
func (client DefinitionsClient) DeleteAtManagementGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteAtManagementGroupResponder handles the response to the DeleteAtManagementGroup request. The method always
// closes the http.Response Body.
func (client DefinitionsClient) DeleteAtManagementGroupResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the policy definition.
//
// policyDefinitionName is the name of the policy definition to get.
func (client DefinitionsClient) Get(policyDefinitionName string) (result Definition, err error) {
	req, err := client.GetPreparer(policyDefinitionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client DefinitionsClient) GetPreparer(policyDefinitionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policyDefinitionName": autorest.Encode("path", policyDefinitionName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DefinitionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DefinitionsClient) GetResponder(resp *http.Response) (result Definition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetAtManagementGroup gets the policy definition at management group level.
//
// policyDefinitionName is the name of the policy definition to get. managementGroupID is the ID of the management
// group.
func (client DefinitionsClient) GetAtManagementGroup(policyDefinitionName string, managementGroupID string) (result Definition, err error) {
	req, err := client.GetAtManagementGroupPreparer(policyDefinitionName, managementGroupID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "GetAtManagementGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetAtManagementGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "GetAtManagementGroup", resp, "Failure sending request")
		return
	}

	result, err = client.GetAtManagementGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "GetAtManagementGroup", resp, "Failure responding to request")
	}

	return
}

// GetAtManagementGroupPreparer prepares the GetAtManagementGroup request.
func (client DefinitionsClient) GetAtManagementGroupPreparer(policyDefinitionName string, managementGroupID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managementGroupId":    autorest.Encode("path", managementGroupID),
		"policyDefinitionName": autorest.Encode("path", policyDefinitionName),
	}

	const APIVersion = "2016-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Management/managementgroups/{managementGroupId}/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetAtManagementGroupSender sends the GetAtManagementGroup request. The method will close the
// http.Response Body if it receives an error.
func (client DefinitionsClient) GetAtManagementGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetAtManagementGroupResponder handles the response to the GetAtManagementGroup request. The method always
// closes the http.Response Body.
func (client DefinitionsClient) GetAtManagementGroupResponder(resp *http.Response) (result Definition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetBuiltIn gets the built in policy definition.
//
// policyDefinitionName is the name of the built in policy definition to get.
func (client DefinitionsClient) GetBuiltIn(policyDefinitionName string) (result Definition, err error) {
	req, err := client.GetBuiltInPreparer(policyDefinitionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "GetBuiltIn", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetBuiltInSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "GetBuiltIn", resp, "Failure sending request")
		return
	}

	result, err = client.GetBuiltInResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "GetBuiltIn", resp, "Failure responding to request")
	}

	return
}

// GetBuiltInPreparer prepares the GetBuiltIn request.
func (client DefinitionsClient) GetBuiltInPreparer(policyDefinitionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policyDefinitionName": autorest.Encode("path", policyDefinitionName),
	}

	const APIVersion = "2016-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Authorization/policyDefinitions/{policyDefinitionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetBuiltInSender sends the GetBuiltIn request. The method will close the
// http.Response Body if it receives an error.
func (client DefinitionsClient) GetBuiltInSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetBuiltInResponder handles the response to the GetBuiltIn request. The method always
// closes the http.Response Body.
func (client DefinitionsClient) GetBuiltInResponder(resp *http.Response) (result Definition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all the policy definitions for a subscription.
func (client DefinitionsClient) List() (result DefinitionListResult, err error) {
	req, err := client.ListPreparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client DefinitionsClient) ListPreparer() (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policyDefinitions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client DefinitionsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client DefinitionsClient) ListResponder(resp *http.Response) (result DefinitionListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client DefinitionsClient) ListNextResults(lastResults DefinitionListResult) (result DefinitionListResult, err error) {
	req, err := lastResults.DefinitionListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "policy.DefinitionsClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "policy.DefinitionsClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "List", resp, "Failure responding to next results request")
	}

	return
}

// ListComplete gets all elements from the list without paging.
func (client DefinitionsClient) ListComplete(cancel <-chan struct{}) (<-chan Definition, <-chan error) {
	resultChan := make(chan Definition)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.List()
		if err != nil {
			errChan <- err
			return
		}
		if list.Value != nil {
			for _, item := range *list.Value {
				select {
				case <-cancel:
					return
				case resultChan <- item:
					// Intentionally left blank
				}
			}
		}
		for list.NextLink != nil {
			list, err = client.ListNextResults(list)
			if err != nil {
				errChan <- err
				return
			}
			if list.Value != nil {
				for _, item := range *list.Value {
					select {
					case <-cancel:
						return
					case resultChan <- item:
						// Intentionally left blank
					}
				}
			}
		}
	}()
	return resultChan, errChan
}

// ListBuiltIn gets all the built in policy definitions.
func (client DefinitionsClient) ListBuiltIn() (result DefinitionListResult, err error) {
	req, err := client.ListBuiltInPreparer()
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "ListBuiltIn", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBuiltInSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "ListBuiltIn", resp, "Failure sending request")
		return
	}

	result, err = client.ListBuiltInResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "ListBuiltIn", resp, "Failure responding to request")
	}

	return
}

// ListBuiltInPreparer prepares the ListBuiltIn request.
func (client DefinitionsClient) ListBuiltInPreparer() (*http.Request, error) {
	const APIVersion = "2016-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Authorization/policyDefinitions"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListBuiltInSender sends the ListBuiltIn request. The method will close the
// http.Response Body if it receives an error.
func (client DefinitionsClient) ListBuiltInSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListBuiltInResponder handles the response to the ListBuiltIn request. The method always
// closes the http.Response Body.
func (client DefinitionsClient) ListBuiltInResponder(resp *http.Response) (result DefinitionListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListBuiltInNextResults retrieves the next set of results, if any.
func (client DefinitionsClient) ListBuiltInNextResults(lastResults DefinitionListResult) (result DefinitionListResult, err error) {
	req, err := lastResults.DefinitionListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "policy.DefinitionsClient", "ListBuiltIn", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListBuiltInSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "policy.DefinitionsClient", "ListBuiltIn", resp, "Failure sending next results request")
	}

	result, err = client.ListBuiltInResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "ListBuiltIn", resp, "Failure responding to next results request")
	}

	return
}

// ListBuiltInComplete gets all elements from the list without paging.
func (client DefinitionsClient) ListBuiltInComplete(cancel <-chan struct{}) (<-chan Definition, <-chan error) {
	resultChan := make(chan Definition)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.ListBuiltIn()
		if err != nil {
			errChan <- err
			return
		}
		if list.Value != nil {
			for _, item := range *list.Value {
				select {
				case <-cancel:
					return
				case resultChan <- item:
					// Intentionally left blank
				}
			}
		}
		for list.NextLink != nil {
			list, err = client.ListBuiltInNextResults(list)
			if err != nil {
				errChan <- err
				return
			}
			if list.Value != nil {
				for _, item := range *list.Value {
					select {
					case <-cancel:
						return
					case resultChan <- item:
						// Intentionally left blank
					}
				}
			}
		}
	}()
	return resultChan, errChan
}

// ListByManagementGroup gets all the policy definitions for a subscription at management group level.
//
// managementGroupID is the ID of the management group.
func (client DefinitionsClient) ListByManagementGroup(managementGroupID string) (result DefinitionListResult, err error) {
	req, err := client.ListByManagementGroupPreparer(managementGroupID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "ListByManagementGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByManagementGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "ListByManagementGroup", resp, "Failure sending request")
		return
	}

	result, err = client.ListByManagementGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "ListByManagementGroup", resp, "Failure responding to request")
	}

	return
}

// ListByManagementGroupPreparer prepares the ListByManagementGroup request.
func (client DefinitionsClient) ListByManagementGroupPreparer(managementGroupID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"managementGroupId": autorest.Encode("path", managementGroupID),
	}

	const APIVersion = "2016-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Management/managementgroups/{managementGroupId}/providers/Microsoft.Authorization/policyDefinitions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByManagementGroupSender sends the ListByManagementGroup request. The method will close the
// http.Response Body if it receives an error.
func (client DefinitionsClient) ListByManagementGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListByManagementGroupResponder handles the response to the ListByManagementGroup request. The method always
// closes the http.Response Body.
func (client DefinitionsClient) ListByManagementGroupResponder(resp *http.Response) (result DefinitionListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByManagementGroupNextResults retrieves the next set of results, if any.
func (client DefinitionsClient) ListByManagementGroupNextResults(lastResults DefinitionListResult) (result DefinitionListResult, err error) {
	req, err := lastResults.DefinitionListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "policy.DefinitionsClient", "ListByManagementGroup", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByManagementGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "policy.DefinitionsClient", "ListByManagementGroup", resp, "Failure sending next results request")
	}

	result, err = client.ListByManagementGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "ListByManagementGroup", resp, "Failure responding to next results request")
	}

	return
}

// ListByManagementGroupComplete gets all elements from the list without paging.
func (client DefinitionsClient) ListByManagementGroupComplete(managementGroupID string, cancel <-chan struct{}) (<-chan Definition, <-chan error) {
	resultChan := make(chan Definition)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.ListByManagementGroup(managementGroupID)
		if err != nil {
			errChan <- err
			return
		}
		if list.Value != nil {
			for _, item := range *list.Value {
				select {
				case <-cancel:
					return
				case resultChan <- item:
					// Intentionally left blank
				}
			}
		}
		for list.NextLink != nil {
			list, err = client.ListByManagementGroupNextResults(list)
			if err != nil {
				errChan <- err
				return
			}
			if list.Value != nil {
				for _, item := range *list.Value {
					select {
					case <-cancel:
						return
					case resultChan <- item:
						// Intentionally left blank
					}
				}
			}
		}
	}()
	return resultChan, errChan
}