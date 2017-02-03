package containerregistry

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
// Code generated by Microsoft (R) AutoRest Code Generator 0.17.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// RegistriesClient is the client for the Registries methods of the
// Containerregistry service.
type RegistriesClient struct {
	ManagementClient
}

// NewRegistriesClient creates an instance of the RegistriesClient client.
func NewRegistriesClient(subscriptionID string) RegistriesClient {
	return NewRegistriesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRegistriesClientWithBaseURI creates an instance of the RegistriesClient
// client.
func NewRegistriesClientWithBaseURI(baseURI string, subscriptionID string) RegistriesClient {
	return RegistriesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CheckNameAvailability checks whether the container registry name is
// available for use. The name must contain only alphanumeric characters, be
// globally unique, and between 5 and 60 characters in length.
//
// registryNameCheckRequest is the object containing information for the
// availability request.
func (client RegistriesClient) CheckNameAvailability(registryNameCheckRequest RegistryNameCheckRequest) (result RegistryNameStatus, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: registryNameCheckRequest,
			Constraints: []validation.Constraint{{Target: "registryNameCheckRequest.Name", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "registryNameCheckRequest.Type", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "containerregistry.RegistriesClient", "CheckNameAvailability")
	}

	req, err := client.CheckNameAvailabilityPreparer(registryNameCheckRequest)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "CheckNameAvailability", nil, "Failure preparing request")
	}

	resp, err := client.CheckNameAvailabilitySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "CheckNameAvailability", resp, "Failure sending request")
	}

	result, err = client.CheckNameAvailabilityResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "CheckNameAvailability", resp, "Failure responding to request")
	}

	return
}

// CheckNameAvailabilityPreparer prepares the CheckNameAvailability request.
func (client RegistriesClient) CheckNameAvailabilityPreparer(registryNameCheckRequest RegistryNameCheckRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.ContainerRegistry/checkNameAvailability", pathParameters),
		autorest.WithJSON(registryNameCheckRequest),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CheckNameAvailabilitySender sends the CheckNameAvailability request. The method will close the
// http.Response Body if it receives an error.
func (client RegistriesClient) CheckNameAvailabilitySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CheckNameAvailabilityResponder handles the response to the CheckNameAvailability request. The method always
// closes the http.Response Body.
func (client RegistriesClient) CheckNameAvailabilityResponder(resp *http.Response) (result RegistryNameStatus, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateOrUpdate creates or updates a container registry with the specified
// parameters.
//
// resourceGroupName is the name of the resource group to which the container
// registry belongs. registryName is the name of the container registry.
// registry is the parameters for creating or updating a container registry.
func (client RegistriesClient) CreateOrUpdate(resourceGroupName string, registryName string, registry Registry) (result Registry, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: registry,
			Constraints: []validation.Constraint{{Target: "registry.RegistryProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "registry.RegistryProperties.StorageAccount", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "registry.RegistryProperties.StorageAccount.Name", Name: validation.Null, Rule: true, Chain: nil},
						{Target: "registry.RegistryProperties.StorageAccount.AccessKey", Name: validation.Null, Rule: true, Chain: nil},
					}},
				}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "containerregistry.RegistriesClient", "CreateOrUpdate")
	}

	req, err := client.CreateOrUpdatePreparer(resourceGroupName, registryName, registry)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "CreateOrUpdate", nil, "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "CreateOrUpdate", resp, "Failure sending request")
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client RegistriesClient) CreateOrUpdatePreparer(resourceGroupName string, registryName string, registry Registry) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"registryName":      autorest.Encode("path", registryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}", pathParameters),
		autorest.WithJSON(registry),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client RegistriesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client RegistriesClient) CreateOrUpdateResponder(resp *http.Response) (result Registry, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a container registry.
//
// resourceGroupName is the name of the resource group to which the container
// registry belongs. registryName is the name of the container registry.
func (client RegistriesClient) Delete(resourceGroupName string, registryName string) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(resourceGroupName, registryName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "Delete", nil, "Failure preparing request")
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "Delete", resp, "Failure sending request")
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client RegistriesClient) DeletePreparer(resourceGroupName string, registryName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"registryName":      autorest.Encode("path", registryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client RegistriesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client RegistriesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetCredentials gets the administrator login credentials for the specified
// container registry.
//
// resourceGroupName is the name of the resource group to which the container
// registry belongs. registryName is the name of the container registry.
func (client RegistriesClient) GetCredentials(resourceGroupName string, registryName string) (result RegistryCredentials, err error) {
	req, err := client.GetCredentialsPreparer(resourceGroupName, registryName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "GetCredentials", nil, "Failure preparing request")
	}

	resp, err := client.GetCredentialsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "GetCredentials", resp, "Failure sending request")
	}

	result, err = client.GetCredentialsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "GetCredentials", resp, "Failure responding to request")
	}

	return
}

// GetCredentialsPreparer prepares the GetCredentials request.
func (client RegistriesClient) GetCredentialsPreparer(resourceGroupName string, registryName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"registryName":      autorest.Encode("path", registryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/getCredentials", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetCredentialsSender sends the GetCredentials request. The method will close the
// http.Response Body if it receives an error.
func (client RegistriesClient) GetCredentialsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetCredentialsResponder handles the response to the GetCredentials request. The method always
// closes the http.Response Body.
func (client RegistriesClient) GetCredentialsResponder(resp *http.Response) (result RegistryCredentials, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetProperties gets the properties of the specified container registry.
//
// resourceGroupName is the name of the resource group to which the container
// registry belongs. registryName is the name of the container registry.
func (client RegistriesClient) GetProperties(resourceGroupName string, registryName string) (result Registry, err error) {
	req, err := client.GetPropertiesPreparer(resourceGroupName, registryName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "GetProperties", nil, "Failure preparing request")
	}

	resp, err := client.GetPropertiesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "GetProperties", resp, "Failure sending request")
	}

	result, err = client.GetPropertiesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "GetProperties", resp, "Failure responding to request")
	}

	return
}

// GetPropertiesPreparer prepares the GetProperties request.
func (client RegistriesClient) GetPropertiesPreparer(resourceGroupName string, registryName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"registryName":      autorest.Encode("path", registryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetPropertiesSender sends the GetProperties request. The method will close the
// http.Response Body if it receives an error.
func (client RegistriesClient) GetPropertiesSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetPropertiesResponder handles the response to the GetProperties request. The method always
// closes the http.Response Body.
func (client RegistriesClient) GetPropertiesResponder(resp *http.Response) (result Registry, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all the available container registries under the specified
// subscription.
func (client RegistriesClient) List() (result RegistryListResult, err error) {
	req, err := client.ListPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "List", nil, "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "List", resp, "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client RegistriesClient) ListPreparer() (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.ContainerRegistry/registries", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client RegistriesClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client RegistriesClient) ListResponder(resp *http.Response) (result RegistryListResult, err error) {
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
func (client RegistriesClient) ListNextResults(lastResults RegistryListResult) (result RegistryListResult, err error) {
	req, err := lastResults.RegistryListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "List", resp, "Failure responding to next results request")
	}

	return
}

// ListByResourceGroup lists all the available container registries under the
// specified resource group.
//
// resourceGroupName is the name of the resource group to which the container
// registry belongs.
func (client RegistriesClient) ListByResourceGroup(resourceGroupName string) (result RegistryListResult, err error) {
	req, err := client.ListByResourceGroupPreparer(resourceGroupName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "ListByResourceGroup", nil, "Failure preparing request")
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "ListByResourceGroup", resp, "Failure sending request")
	}

	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client RegistriesClient) ListByResourceGroupPreparer(resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client RegistriesClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client RegistriesClient) ListByResourceGroupResponder(resp *http.Response) (result RegistryListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroupNextResults retrieves the next set of results, if any.
func (client RegistriesClient) ListByResourceGroupNextResults(lastResults RegistryListResult) (result RegistryListResult, err error) {
	req, err := lastResults.RegistryListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "ListByResourceGroup", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "ListByResourceGroup", resp, "Failure sending next results request")
	}

	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "ListByResourceGroup", resp, "Failure responding to next results request")
	}

	return
}

// RegenerateCredentials regenerates the administrator login credentials for
// the specified container registry.
//
// resourceGroupName is the name of the resource group to which the container
// registry belongs. registryName is the name of the container registry.
func (client RegistriesClient) RegenerateCredentials(resourceGroupName string, registryName string) (result RegistryCredentials, err error) {
	req, err := client.RegenerateCredentialsPreparer(resourceGroupName, registryName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "RegenerateCredentials", nil, "Failure preparing request")
	}

	resp, err := client.RegenerateCredentialsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "RegenerateCredentials", resp, "Failure sending request")
	}

	result, err = client.RegenerateCredentialsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "RegenerateCredentials", resp, "Failure responding to request")
	}

	return
}

// RegenerateCredentialsPreparer prepares the RegenerateCredentials request.
func (client RegistriesClient) RegenerateCredentialsPreparer(resourceGroupName string, registryName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"registryName":      autorest.Encode("path", registryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/regenerateCredentials", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// RegenerateCredentialsSender sends the RegenerateCredentials request. The method will close the
// http.Response Body if it receives an error.
func (client RegistriesClient) RegenerateCredentialsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// RegenerateCredentialsResponder handles the response to the RegenerateCredentials request. The method always
// closes the http.Response Body.
func (client RegistriesClient) RegenerateCredentialsResponder(resp *http.Response) (result RegistryCredentials, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update updates a container registry with the specified parameters.
//
// resourceGroupName is the name of the resource group to which the container
// registry belongs. registryName is the name of the container registry.
// registryUpdateParameters is the parameters for updating a container
// registry.
func (client RegistriesClient) Update(resourceGroupName string, registryName string, registryUpdateParameters RegistryUpdateParameters) (result Registry, err error) {
	req, err := client.UpdatePreparer(resourceGroupName, registryName, registryUpdateParameters)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "Update", nil, "Failure preparing request")
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "Update", resp, "Failure sending request")
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.RegistriesClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client RegistriesClient) UpdatePreparer(resourceGroupName string, registryName string, registryUpdateParameters RegistryUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"registryName":      autorest.Encode("path", registryName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": client.APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}", pathParameters),
		autorest.WithJSON(registryUpdateParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client RegistriesClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client RegistriesClient) UpdateResponder(resp *http.Response) (result Registry, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
