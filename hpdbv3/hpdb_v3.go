/**
 * (C) Copyright IBM Corp. 2021,2022.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/*
 * IBM OpenAPI SDK Code Generator Version: 3.51.0-5b8b699d-20220613-200818
 */

// Package hpdbv3 : Operations and models for the HpdbV3 service
package hpdbv3

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/IBM/hpdb-go-sdk/common"
)

// HpdbV3 : The DBaaS RESTful APIs are used to manage the database cluster, the database itself, and database users.
//
// API Version: 3
type HpdbV3 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://dbaas900.hyperp-dbaas.cloud.ibm.com/api/v3/unknown"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "hpdb"

const ParameterizedServiceURL = "https://dbaas900.hyperp-dbaas.cloud.ibm.com/api/v3/{account_id}"

var defaultUrlVariables = map[string]string{
	"account_id": "unknown",
}

// HpdbV3Options : Service options
type HpdbV3Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewHpdbV3UsingExternalConfig : constructs an instance of HpdbV3 with passed in options and external configuration.
func NewHpdbV3UsingExternalConfig(options *HpdbV3Options) (hpdb *HpdbV3, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	hpdb, err = NewHpdbV3(options)
	if err != nil {
		return
	}

	err = hpdb.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = hpdb.Service.SetServiceURL(options.URL)
	}
	return
}

// NewHpdbV3 : constructs an instance of HpdbV3 with passed in options.
func NewHpdbV3(options *HpdbV3Options) (service *HpdbV3, err error) {
	serviceOptions := &core.ServiceOptions{
		URL:           DefaultServiceURL,
		Authenticator: options.Authenticator,
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = baseService.SetServiceURL(options.URL)
		if err != nil {
			return
		}
	}

	service = &HpdbV3{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "hpdb" suitable for processing requests.
func (hpdb *HpdbV3) Clone() *HpdbV3 {
	if core.IsNil(hpdb) {
		return nil
	}
	clone := *hpdb
	clone.Service = hpdb.Service.Clone()
	return &clone
}

// ConstructServiceURL constructs a service URL from the parameterized URL.
func ConstructServiceURL(providedUrlVariables map[string]string) (string, error) {
	return core.ConstructServiceURL(ParameterizedServiceURL, defaultUrlVariables, providedUrlVariables)
}

// SetServiceURL sets the service URL
func (hpdb *HpdbV3) SetServiceURL(url string) error {
	return hpdb.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (hpdb *HpdbV3) GetServiceURL() string {
	return hpdb.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (hpdb *HpdbV3) SetDefaultHeaders(headers http.Header) {
	hpdb.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (hpdb *HpdbV3) SetEnableGzipCompression(enableGzip bool) {
	hpdb.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (hpdb *HpdbV3) GetEnableGzipCompression() bool {
	return hpdb.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (hpdb *HpdbV3) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	hpdb.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (hpdb *HpdbV3) DisableRetries() {
	hpdb.Service.DisableRetries()
}

// GetCluster : Get database cluster details
// Get the detailed information of the specific database cluster that is indicated by its ID.
func (hpdb *HpdbV3) GetCluster(getClusterOptions *GetClusterOptions) (result *Cluster, response *core.DetailedResponse, err error) {
	return hpdb.GetClusterWithContext(context.Background(), getClusterOptions)
}

// GetClusterWithContext is an alternate form of the GetCluster method which supports a Context parameter
func (hpdb *HpdbV3) GetClusterWithContext(ctx context.Context, getClusterOptions *GetClusterOptions) (result *Cluster, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getClusterOptions, "getClusterOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getClusterOptions, "getClusterOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"cluster_id": *getClusterOptions.ClusterID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = hpdb.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(hpdb.Service.Options.URL, `/clusters/{cluster_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getClusterOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("hpdb", "V3", "GetCluster")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = hpdb.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCluster)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListUsers : List database users
// List the information about all the users in the specified database cluster that is indicated by its ID.
func (hpdb *HpdbV3) ListUsers(listUsersOptions *ListUsersOptions) (result *Users, response *core.DetailedResponse, err error) {
	return hpdb.ListUsersWithContext(context.Background(), listUsersOptions)
}

// ListUsersWithContext is an alternate form of the ListUsers method which supports a Context parameter
func (hpdb *HpdbV3) ListUsersWithContext(ctx context.Context, listUsersOptions *ListUsersOptions) (result *Users, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listUsersOptions, "listUsersOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listUsersOptions, "listUsersOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"cluster_id": *listUsersOptions.ClusterID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = hpdb.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(hpdb.Service.Options.URL, `/clusters/{cluster_id}/users`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listUsersOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("hpdb", "V3", "ListUsers")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = hpdb.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalUsers)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetUser : Get database user details
// Get the detailed information about the user of a specified database cluster that is indicated by its ID.
func (hpdb *HpdbV3) GetUser(getUserOptions *GetUserOptions) (result *UserDetails, response *core.DetailedResponse, err error) {
	return hpdb.GetUserWithContext(context.Background(), getUserOptions)
}

// GetUserWithContext is an alternate form of the GetUser method which supports a Context parameter
func (hpdb *HpdbV3) GetUserWithContext(ctx context.Context, getUserOptions *GetUserOptions) (result *UserDetails, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getUserOptions, "getUserOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getUserOptions, "getUserOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"cluster_id": *getUserOptions.ClusterID,
		"db_user_id": *getUserOptions.DbUserID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = hpdb.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(hpdb.Service.Options.URL, `/clusters/{cluster_id}/users/{db_user_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getUserOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("hpdb", "V3", "GetUser")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = hpdb.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalUserDetails)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListDatabases : List databases
// Get a list of all databases in a specified database cluster that is indicated by its ID.
func (hpdb *HpdbV3) ListDatabases(listDatabasesOptions *ListDatabasesOptions) (result *Databases, response *core.DetailedResponse, err error) {
	return hpdb.ListDatabasesWithContext(context.Background(), listDatabasesOptions)
}

// ListDatabasesWithContext is an alternate form of the ListDatabases method which supports a Context parameter
func (hpdb *HpdbV3) ListDatabasesWithContext(ctx context.Context, listDatabasesOptions *ListDatabasesOptions) (result *Databases, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listDatabasesOptions, "listDatabasesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listDatabasesOptions, "listDatabasesOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"cluster_id": *listDatabasesOptions.ClusterID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = hpdb.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(hpdb.Service.Options.URL, `/clusters/{cluster_id}/databases`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listDatabasesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("hpdb", "V3", "ListDatabases")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = hpdb.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDatabases)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ScaleResources : Scale resources
// Scale resources in a specified cluster that is indicated by its ID.
func (hpdb *HpdbV3) ScaleResources(scaleResourcesOptions *ScaleResourcesOptions) (result *TaskID, response *core.DetailedResponse, err error) {
	return hpdb.ScaleResourcesWithContext(context.Background(), scaleResourcesOptions)
}

// ScaleResourcesWithContext is an alternate form of the ScaleResources method which supports a Context parameter
func (hpdb *HpdbV3) ScaleResourcesWithContext(ctx context.Context, scaleResourcesOptions *ScaleResourcesOptions) (result *TaskID, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(scaleResourcesOptions, "scaleResourcesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(scaleResourcesOptions, "scaleResourcesOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"cluster_id": *scaleResourcesOptions.ClusterID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = hpdb.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(hpdb.Service.Options.URL, `/clusters/{cluster_id}/resource`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range scaleResourcesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("hpdb", "V3", "ScaleResources")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if scaleResourcesOptions.Resource != nil {
		body["resource"] = scaleResourcesOptions.Resource
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = hpdb.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTaskID)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetConfiguration : Get configuration
// Get database configuration in a specified cluster that is indicated by its ID.
func (hpdb *HpdbV3) GetConfiguration(getConfigurationOptions *GetConfigurationOptions) (result *Configuration, response *core.DetailedResponse, err error) {
	return hpdb.GetConfigurationWithContext(context.Background(), getConfigurationOptions)
}

// GetConfigurationWithContext is an alternate form of the GetConfiguration method which supports a Context parameter
func (hpdb *HpdbV3) GetConfigurationWithContext(ctx context.Context, getConfigurationOptions *GetConfigurationOptions) (result *Configuration, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getConfigurationOptions, "getConfigurationOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getConfigurationOptions, "getConfigurationOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"cluster_id": *getConfigurationOptions.ClusterID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = hpdb.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(hpdb.Service.Options.URL, `/clusters/{cluster_id}/configuration`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getConfigurationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("hpdb", "V3", "GetConfiguration")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = hpdb.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalConfiguration)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateConfiguration : Update configuration
// Update database configuration in a specified cluster that is indicated by its ID.
func (hpdb *HpdbV3) UpdateConfiguration(updateConfigurationOptions *UpdateConfigurationOptions) (result *TaskID, response *core.DetailedResponse, err error) {
	return hpdb.UpdateConfigurationWithContext(context.Background(), updateConfigurationOptions)
}

// UpdateConfigurationWithContext is an alternate form of the UpdateConfiguration method which supports a Context parameter
func (hpdb *HpdbV3) UpdateConfigurationWithContext(ctx context.Context, updateConfigurationOptions *UpdateConfigurationOptions) (result *TaskID, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateConfigurationOptions, "updateConfigurationOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateConfigurationOptions, "updateConfigurationOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"cluster_id": *updateConfigurationOptions.ClusterID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = hpdb.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(hpdb.Service.Options.URL, `/clusters/{cluster_id}/configuration`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateConfigurationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("hpdb", "V3", "UpdateConfiguration")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateConfigurationOptions.Configuration != nil {
		body["configuration"] = updateConfigurationOptions.Configuration
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = hpdb.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTaskID)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListTasks : List tasks
// List tasks running or recently run on a specified cluster that is indicated by its ID.
func (hpdb *HpdbV3) ListTasks(listTasksOptions *ListTasksOptions) (result *Tasks, response *core.DetailedResponse, err error) {
	return hpdb.ListTasksWithContext(context.Background(), listTasksOptions)
}

// ListTasksWithContext is an alternate form of the ListTasks method which supports a Context parameter
func (hpdb *HpdbV3) ListTasksWithContext(ctx context.Context, listTasksOptions *ListTasksOptions) (result *Tasks, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listTasksOptions, "listTasksOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listTasksOptions, "listTasksOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"cluster_id": *listTasksOptions.ClusterID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = hpdb.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(hpdb.Service.Options.URL, `/clusters/{cluster_id}/tasks`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listTasksOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("hpdb", "V3", "ListTasks")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = hpdb.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTasks)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetTask : Show task
// Show task information of a specified task ID in a cluster.
func (hpdb *HpdbV3) GetTask(getTaskOptions *GetTaskOptions) (result *Task, response *core.DetailedResponse, err error) {
	return hpdb.GetTaskWithContext(context.Background(), getTaskOptions)
}

// GetTaskWithContext is an alternate form of the GetTask method which supports a Context parameter
func (hpdb *HpdbV3) GetTaskWithContext(ctx context.Context, getTaskOptions *GetTaskOptions) (result *Task, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getTaskOptions, "getTaskOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getTaskOptions, "getTaskOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"cluster_id": *getTaskOptions.ClusterID,
		"task_id": *getTaskOptions.TaskID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = hpdb.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(hpdb.Service.Options.URL, `/clusters/{cluster_id}/tasks/{task_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getTaskOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("hpdb", "V3", "GetTask")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = hpdb.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTask)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListBackups : List backups
// Get the list of backups.
func (hpdb *HpdbV3) ListBackups(listBackupsOptions *ListBackupsOptions) (result *ListBackupsResponse, response *core.DetailedResponse, err error) {
	return hpdb.ListBackupsWithContext(context.Background(), listBackupsOptions)
}

// ListBackupsWithContext is an alternate form of the ListBackups method which supports a Context parameter
func (hpdb *HpdbV3) ListBackupsWithContext(ctx context.Context, listBackupsOptions *ListBackupsOptions) (result *ListBackupsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listBackupsOptions, "listBackupsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listBackupsOptions, "listBackupsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"cluster_id": *listBackupsOptions.ClusterID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = hpdb.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(hpdb.Service.Options.URL, `/clusters/{cluster_id}/backups`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listBackupsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("hpdb", "V3", "ListBackups")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = hpdb.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListBackupsResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// EnableCosBackup : Enable backup to COS
// Enable backup to COS.
func (hpdb *HpdbV3) EnableCosBackup(enableCosBackupOptions *EnableCosBackupOptions) (result *TaskID, response *core.DetailedResponse, err error) {
	return hpdb.EnableCosBackupWithContext(context.Background(), enableCosBackupOptions)
}

// EnableCosBackupWithContext is an alternate form of the EnableCosBackup method which supports a Context parameter
func (hpdb *HpdbV3) EnableCosBackupWithContext(ctx context.Context, enableCosBackupOptions *EnableCosBackupOptions) (result *TaskID, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(enableCosBackupOptions, "enableCosBackupOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(enableCosBackupOptions, "enableCosBackupOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"cluster_id": *enableCosBackupOptions.ClusterID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = hpdb.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(hpdb.Service.Options.URL, `/clusters/{cluster_id}/backups/cos/enable`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range enableCosBackupOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("hpdb", "V3", "EnableCosBackup")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if enableCosBackupOptions.CosHmacKeys != nil {
		body["cos_hmac_keys"] = enableCosBackupOptions.CosHmacKeys
	}
	if enableCosBackupOptions.CosEndpoint != nil {
		body["cos_endpoint"] = enableCosBackupOptions.CosEndpoint
	}
	if enableCosBackupOptions.BucketInstanceCrn != nil {
		body["bucket_instance_crn"] = enableCosBackupOptions.BucketInstanceCrn
	}
	if enableCosBackupOptions.Schedule != nil {
		body["schedule"] = enableCosBackupOptions.Schedule
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = hpdb.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTaskID)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DisableCosBackup : Disable backup to COS
// Disable backup to COS.
func (hpdb *HpdbV3) DisableCosBackup(disableCosBackupOptions *DisableCosBackupOptions) (result *TaskID, response *core.DetailedResponse, err error) {
	return hpdb.DisableCosBackupWithContext(context.Background(), disableCosBackupOptions)
}

// DisableCosBackupWithContext is an alternate form of the DisableCosBackup method which supports a Context parameter
func (hpdb *HpdbV3) DisableCosBackupWithContext(ctx context.Context, disableCosBackupOptions *DisableCosBackupOptions) (result *TaskID, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(disableCosBackupOptions, "disableCosBackupOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(disableCosBackupOptions, "disableCosBackupOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"cluster_id": *disableCosBackupOptions.ClusterID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = hpdb.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(hpdb.Service.Options.URL, `/clusters/{cluster_id}/backups/cos/disable`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range disableCosBackupOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("hpdb", "V3", "DisableCosBackup")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = hpdb.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTaskID)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetCosBackupConfig : Get backup configuration (Deprecated)
// Get backup configuration.
func (hpdb *HpdbV3) GetCosBackupConfig(getCosBackupConfigOptions *GetCosBackupConfigOptions) (result *GetCosBackupConfigResponse, response *core.DetailedResponse, err error) {
	return hpdb.GetCosBackupConfigWithContext(context.Background(), getCosBackupConfigOptions)
}

// GetCosBackupConfigWithContext is an alternate form of the GetCosBackupConfig method which supports a Context parameter
func (hpdb *HpdbV3) GetCosBackupConfigWithContext(ctx context.Context, getCosBackupConfigOptions *GetCosBackupConfigOptions) (result *GetCosBackupConfigResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getCosBackupConfigOptions, "getCosBackupConfigOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getCosBackupConfigOptions, "getCosBackupConfigOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"cluster_id": *getCosBackupConfigOptions.ClusterID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = hpdb.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(hpdb.Service.Options.URL, `/clusters/{cluster_id}/backups/cos/configuration`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getCosBackupConfigOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("hpdb", "V3", "GetCosBackupConfig")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = hpdb.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetCosBackupConfigResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetBackupConfig : Get backup configuration
// Get backup configuration.
func (hpdb *HpdbV3) GetBackupConfig(getBackupConfigOptions *GetBackupConfigOptions) (result *GetBackupConfigResponse, response *core.DetailedResponse, err error) {
	return hpdb.GetBackupConfigWithContext(context.Background(), getBackupConfigOptions)
}

// GetBackupConfigWithContext is an alternate form of the GetBackupConfig method which supports a Context parameter
func (hpdb *HpdbV3) GetBackupConfigWithContext(ctx context.Context, getBackupConfigOptions *GetBackupConfigOptions) (result *GetBackupConfigResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getBackupConfigOptions, "getBackupConfigOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getBackupConfigOptions, "getBackupConfigOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"cluster_id": *getBackupConfigOptions.ClusterID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = hpdb.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(hpdb.Service.Options.URL, `/clusters/{cluster_id}/backups/configuration`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getBackupConfigOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("hpdb", "V3", "GetBackupConfig")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = hpdb.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetBackupConfigResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateBackupConfig : Update backup configuration
// Update backup configuration.
func (hpdb *HpdbV3) UpdateBackupConfig(updateBackupConfigOptions *UpdateBackupConfigOptions) (result *TaskID, response *core.DetailedResponse, err error) {
	return hpdb.UpdateBackupConfigWithContext(context.Background(), updateBackupConfigOptions)
}

// UpdateBackupConfigWithContext is an alternate form of the UpdateBackupConfig method which supports a Context parameter
func (hpdb *HpdbV3) UpdateBackupConfigWithContext(ctx context.Context, updateBackupConfigOptions *UpdateBackupConfigOptions) (result *TaskID, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateBackupConfigOptions, "updateBackupConfigOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateBackupConfigOptions, "updateBackupConfigOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"cluster_id": *updateBackupConfigOptions.ClusterID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = hpdb.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(hpdb.Service.Options.URL, `/clusters/{cluster_id}/backups/configuration`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateBackupConfigOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("hpdb", "V3", "UpdateBackupConfig")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateBackupConfigOptions.Cos != nil {
		body["cos"] = updateBackupConfigOptions.Cos
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = hpdb.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTaskID)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// Restore : Restore from backup
// Restore from a backup file.
func (hpdb *HpdbV3) Restore(restoreOptions *RestoreOptions) (result *TaskID, response *core.DetailedResponse, err error) {
	return hpdb.RestoreWithContext(context.Background(), restoreOptions)
}

// RestoreWithContext is an alternate form of the Restore method which supports a Context parameter
func (hpdb *HpdbV3) RestoreWithContext(ctx context.Context, restoreOptions *RestoreOptions) (result *TaskID, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(restoreOptions, "restoreOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(restoreOptions, "restoreOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"cluster_id": *restoreOptions.ClusterID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = hpdb.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(hpdb.Service.Options.URL, `/clusters/{cluster_id}/restore`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range restoreOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("hpdb", "V3", "Restore")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if restoreOptions.SourceType != nil {
		body["source_type"] = restoreOptions.SourceType
	}
	if restoreOptions.CosHmacKeys != nil {
		body["cos_hmac_keys"] = restoreOptions.CosHmacKeys
	}
	if restoreOptions.CosEndpoint != nil {
		body["cos_endpoint"] = restoreOptions.CosEndpoint
	}
	if restoreOptions.BucketInstanceCrn != nil {
		body["bucket_instance_crn"] = restoreOptions.BucketInstanceCrn
	}
	if restoreOptions.BackupFile != nil {
		body["backup_file"] = restoreOptions.BackupFile
	}
	if restoreOptions.BackupID != nil {
		body["backup_id"] = restoreOptions.BackupID
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = hpdb.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTaskID)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListNodeLogs : List database log files of a node
// List the latest log files of the node that is indicated by its ID.
func (hpdb *HpdbV3) ListNodeLogs(listNodeLogsOptions *ListNodeLogsOptions) (result *LogList, response *core.DetailedResponse, err error) {
	return hpdb.ListNodeLogsWithContext(context.Background(), listNodeLogsOptions)
}

// ListNodeLogsWithContext is an alternate form of the ListNodeLogs method which supports a Context parameter
func (hpdb *HpdbV3) ListNodeLogsWithContext(ctx context.Context, listNodeLogsOptions *ListNodeLogsOptions) (result *LogList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listNodeLogsOptions, "listNodeLogsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listNodeLogsOptions, "listNodeLogsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"node_id": *listNodeLogsOptions.NodeID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = hpdb.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(hpdb.Service.Options.URL, `/nodes/{node_id}/logs`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listNodeLogsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("hpdb", "V3", "ListNodeLogs")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = hpdb.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalLogList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetLog : Get log details
// Get the content of the specified log file of the node that is indicated by its ID.
func (hpdb *HpdbV3) GetLog(getLogOptions *GetLogOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	return hpdb.GetLogWithContext(context.Background(), getLogOptions)
}

// GetLogWithContext is an alternate form of the GetLog method which supports a Context parameter
func (hpdb *HpdbV3) GetLogWithContext(ctx context.Context, getLogOptions *GetLogOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getLogOptions, "getLogOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getLogOptions, "getLogOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"node_id": *getLogOptions.NodeID,
		"log_name": *getLogOptions.LogName,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = hpdb.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(hpdb.Service.Options.URL, `/nodes/{node_id}/logs/{log_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getLogOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("hpdb", "V3", "GetLog")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getLogOptions.Accept != nil {
		builder.AddHeader("Accept", fmt.Sprint(*getLogOptions.Accept))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = hpdb.Service.Request(request, &result)

	return
}

// Access : The privileges on databases.
type Access struct {
	// Name of the database to which the privilege is applied.
	Db *string `json:"db,omitempty"`

	// Privileges that are applied to the specified database. If called API is listing all users, this will be returned
	// only for MongoDB. For MongoDB, privileges are associated with MongoDB built-in roles. Common ones are read,
	// readWrite, dbAdmin, userAdmin, and clusterAdmin. For more information, see MongoDB documentation
	// (https://docs.mongodb.com/manual/reference/built-in-roles/). For PostgreSQL, privileges are assigned by the GRANT
	// command. For example, on databases, privileges can be CARETE, CONNECT, TEMP, TEMPORARY. For more information, see
	// PostgreSQL documentaton (https://www.postgresql.org/docs/current/static/sql-grant.html).
	Privileges []string `json:"privileges,omitempty"`
}

// UnmarshalAccess unmarshals an instance of Access from the specified map of raw messages.
func UnmarshalAccess(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Access)
	err = core.UnmarshalPrimitive(m, "db", &obj.Db)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "privileges", &obj.Privileges)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Backup : The backup of a datbase cluster.
type Backup struct {
	// The ID of the backup.
	ID *string `json:"id,omitempty"`

	// The type of the backup.
	Type *string `json:"type,omitempty"`

	// The date and time when the backup is created.
	CreatedAt *string `json:"created_at,omitempty"`
}

// UnmarshalBackup unmarshals an instance of Backup from the specified map of raw messages.
func UnmarshalBackup(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Backup)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// BackupSchedule : BackupSchedule struct
type BackupSchedule struct {
	// The type of backup schedule.
	Type *string `json:"type,omitempty"`

	// Available values are '1h', '2h', '4h', '8h', '1d', '2d', '1w'. 'h' or 'H' means hour. 'd' or 'D' means day. 'w' or
	// 'W' means week.
	Value *string `json:"value,omitempty"`
}

// UnmarshalBackupSchedule unmarshals an instance of BackupSchedule from the specified map of raw messages.
func UnmarshalBackupSchedule(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BackupSchedule)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Cluster : An object which shows detailed information about a cluster which has id, name, state, three replicas and so on.
type Cluster struct {
	// The ID of the cluster object.
	ID *string `json:"id,omitempty"`

	// The CRN of the database cluster (service instance).
	Crn *string `json:"crn,omitempty"`

	// The status of database cluster monitoring.
	CustomerMonitoringStatus *string `json:"customer_monitoring_status,omitempty"`

	// The status of backup to COS.
	IsCosBackupEnabled *bool `json:"is_cos_backup_enabled,omitempty"`

	// The region of the cluster.
	Region *string `json:"region,omitempty"`

	// The name of the cluster.
	Name *string `json:"name,omitempty"`

	// The current state of the cluster.
	State *string `json:"state,omitempty"`

	// The reason why the cluster entered the failed state.
	Reason *string `json:"reason,omitempty"`

	// The type of the database cluster; currently 'mongodb' and 'postgresql' are supported.
	DbType *string `json:"db_type,omitempty"`

	// The version of the database cluster.
	DbVersion *string `json:"db_version,omitempty"`

	// The public endpoint of cluster.
	PublicEndpoint *string `json:"public_endpoint,omitempty"`

	// The private endpoint of cluster.
	PrivateEndpoint *string `json:"private_endpoint,omitempty"`

	// The type of private endpoint.
	PrivateEndpointType *string `json:"private_endpoint_type,omitempty"`

	// The plan ID of the database cluster.
	PlanID *string `json:"plan_id,omitempty"`

	// Time interval since last activation.
	LastActive *int64 `json:"last_active,omitempty"`

	// IBM Cloud Logging service url for DBA.
	LogURL *string `json:"log_url,omitempty"`

	// IBM Cloud Monitoring service url.
	MetricURL *string `json:"metric_url,omitempty"`

	// The number of replicas of the cluster to be created; currently only 3 is supported.
	ReplicaCount *int64 `json:"replica_count,omitempty"`

	// The IBM Cloud account ID.
	UserID *string `json:"user_id,omitempty"`

	// The resources required by the cluster to be created.
	Resource *ClusterResource `json:"resource,omitempty"`

	// The external key information.
	ExternalKey *ClusterExternalKey `json:"external_key,omitempty"`

	// The details of nodes that were created.
	Nodes []Node `json:"nodes,omitempty"`

	// The UTC time when the cluster object was created.
	CreatedAt *string `json:"created_at,omitempty"`

	// The UTC time when the cluster object is updated.
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// UnmarshalCluster unmarshals an instance of Cluster from the specified map of raw messages.
func UnmarshalCluster(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Cluster)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "crn", &obj.Crn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "customer_monitoring_status", &obj.CustomerMonitoringStatus)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "is_cos_backup_enabled", &obj.IsCosBackupEnabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "region", &obj.Region)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "reason", &obj.Reason)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "db_type", &obj.DbType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "db_version", &obj.DbVersion)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "public_endpoint", &obj.PublicEndpoint)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "private_endpoint", &obj.PrivateEndpoint)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "private_endpoint_type", &obj.PrivateEndpointType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "plan_id", &obj.PlanID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last_active", &obj.LastActive)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "log_url", &obj.LogURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "metric_url", &obj.MetricURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "replica_count", &obj.ReplicaCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "user_id", &obj.UserID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "resource", &obj.Resource, UnmarshalClusterResource)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "external_key", &obj.ExternalKey, UnmarshalClusterExternalKey)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "nodes", &obj.Nodes, UnmarshalNode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_at", &obj.UpdatedAt)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ClusterExternalKey : The external key information.
type ClusterExternalKey struct {
	// The CRN (Cloud Resource Name) of the KMS (key management service) instance.
	KmsInstance *string `json:"kms_instance" validate:"required"`

	// The ID of the root key of the KMS instance.
	KmsKey *string `json:"kms_key" validate:"required"`
}

// UnmarshalClusterExternalKey unmarshals an instance of ClusterExternalKey from the specified map of raw messages.
func UnmarshalClusterExternalKey(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ClusterExternalKey)
	err = core.UnmarshalPrimitive(m, "kms_instance", &obj.KmsInstance)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "kms_key", &obj.KmsKey)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ClusterResource : The resources required by the cluster to be created.
type ClusterResource struct {
	// The CPU number.
	Cpu *int64 `json:"cpu" validate:"required"`

	// The memory size in units MB, MiB, GB, GiB, TB or TiB.
	Memory *string `json:"memory" validate:"required"`

	// The storage size in units MB, MiB, GB, GiB, TB or TiB.
	Storage *string `json:"storage" validate:"required"`

	// The used storate size in units MB, MiB, GB, GiB, TB or TiB.
	StorageUsed *string `json:"storage_used,omitempty"`
}

// UnmarshalClusterResource unmarshals an instance of ClusterResource from the specified map of raw messages.
func UnmarshalClusterResource(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ClusterResource)
	err = core.UnmarshalPrimitive(m, "cpu", &obj.Cpu)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "memory", &obj.Memory)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "storage", &obj.Storage)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "storage_used", &obj.StorageUsed)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Configuration : Database configuration.
type Configuration struct {
	// Parameter infomation.
	Configuration *ConfigurationItem `json:"configuration,omitempty"`
}

// UnmarshalConfiguration unmarshals an instance of Configuration from the specified map of raw messages.
func UnmarshalConfiguration(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Configuration)
	err = core.UnmarshalModel(m, "configuration", &obj.Configuration, UnmarshalConfigurationItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigurationItem : Parameter infomation.
type ConfigurationItem struct {
	// Integer type parameter.
	DeadlockTimeout *IntegerType `json:"deadlock_timeout" validate:"required"`

	// Integer type parameter.
	MaxLocksPerTransaction *IntegerType `json:"max_locks_per_transaction" validate:"required"`

	// Integer type parameter.
	SharedBuffers *IntegerType `json:"shared_buffers" validate:"required"`

	// Integer type parameter.
	MaxConnections *IntegerType `json:"max_connections" validate:"required"`
}

// UnmarshalConfigurationItem unmarshals an instance of ConfigurationItem from the specified map of raw messages.
func UnmarshalConfigurationItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigurationItem)
	err = core.UnmarshalModel(m, "deadlock_timeout", &obj.DeadlockTimeout, UnmarshalIntegerType)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "max_locks_per_transaction", &obj.MaxLocksPerTransaction, UnmarshalIntegerType)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "shared_buffers", &obj.SharedBuffers, UnmarshalIntegerType)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "max_connections", &obj.MaxConnections, UnmarshalIntegerType)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Configurations : Object of information about configuration.
type Configurations struct {
	// Value of deadlock_timeout to be updated.
	DeadlockTimeout *int64 `json:"deadlock_timeout,omitempty"`

	// Value of max_locks_per_transaction to be updated.
	MaxLocksPerTransaction *int64 `json:"max_locks_per_transaction,omitempty"`

	// Value of shared_buffers to be updated.
	SharedBuffers *int64 `json:"shared_buffers,omitempty"`

	// Value of max_connections to be updated.
	MaxConnections *int64 `json:"max_connections,omitempty"`
}

// UnmarshalConfigurations unmarshals an instance of Configurations from the specified map of raw messages.
func UnmarshalConfigurations(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Configurations)
	err = core.UnmarshalPrimitive(m, "deadlock_timeout", &obj.DeadlockTimeout)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "max_locks_per_transaction", &obj.MaxLocksPerTransaction)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "shared_buffers", &obj.SharedBuffers)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "max_connections", &obj.MaxConnections)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CosBackupConfig : CosBackupConfig struct
type CosBackupConfig struct {
	CosHmacKeys *CosHmacKeys `json:"cos_hmac_keys,omitempty"`

	// Endpoint of IBM Cloud Object Storage (required). In the IBM Cloud Object Storage web UI, select your bucket and you
	// can find Endpoints in the Configuration tab. Only public endpoints are supported for now. Support for private
	// endpoints will come soon.
	CosEndpoint *string `json:"cos_endpoint,omitempty"`

	// COS bucket CRN (required). In the IBM Cloud Object Storage web UI, select your bucket and you can find Bucket
	// instance CRN in the Configurations tab.
	BucketInstanceCrn *string `json:"bucket_instance_crn,omitempty"`

	Schedule *BackupSchedule `json:"schedule,omitempty"`
}

// UnmarshalCosBackupConfig unmarshals an instance of CosBackupConfig from the specified map of raw messages.
func UnmarshalCosBackupConfig(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CosBackupConfig)
	err = core.UnmarshalModel(m, "cos_hmac_keys", &obj.CosHmacKeys, UnmarshalCosHmacKeys)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cos_endpoint", &obj.CosEndpoint)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "bucket_instance_crn", &obj.BucketInstanceCrn)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "schedule", &obj.Schedule, UnmarshalBackupSchedule)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CosHmacKeys : CosHmacKeys struct
type CosHmacKeys struct {
	// COS HMAC access key ID.
	AccessKeyID *string `json:"access_key_id,omitempty"`

	// COS HMAC secret access key.
	SecretAccessKey *string `json:"secret_access_key,omitempty"`
}

// UnmarshalCosHmacKeys unmarshals an instance of CosHmacKeys from the specified map of raw messages.
func UnmarshalCosHmacKeys(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CosHmacKeys)
	err = core.UnmarshalPrimitive(m, "access_key_id", &obj.AccessKeyID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "secret_access_key", &obj.SecretAccessKey)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Database : Database information.
type Database struct {
	// Name of the database.
	Name *string `json:"name,omitempty"`

	// Total size of the database files on disk, in bytes.
	SizeOnDisk *int64 `json:"size_on_disk,omitempty"`
}

// UnmarshalDatabase unmarshals an instance of Database from the specified map of raw messages.
func UnmarshalDatabase(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Database)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "size_on_disk", &obj.SizeOnDisk)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Databases : Information about all databases in a cluster.
type Databases struct {
	// The result of the sum of all the size_on_disk fields, in bytes.
	TotalSize *int64 `json:"total_size,omitempty"`

	// A list of databases.
	Databases []Database `json:"databases,omitempty"`
}

// UnmarshalDatabases unmarshals an instance of Databases from the specified map of raw messages.
func UnmarshalDatabases(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Databases)
	err = core.UnmarshalPrimitive(m, "total_size", &obj.TotalSize)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "databases", &obj.Databases, UnmarshalDatabase)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DisableCosBackupOptions : The DisableCosBackup options.
type DisableCosBackupOptions struct {
	// The ID of a cluster object.
	ClusterID *string `json:"cluster_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDisableCosBackupOptions : Instantiate DisableCosBackupOptions
func (*HpdbV3) NewDisableCosBackupOptions(clusterID string) *DisableCosBackupOptions {
	return &DisableCosBackupOptions{
		ClusterID: core.StringPtr(clusterID),
	}
}

// SetClusterID : Allow user to set ClusterID
func (_options *DisableCosBackupOptions) SetClusterID(clusterID string) *DisableCosBackupOptions {
	_options.ClusterID = core.StringPtr(clusterID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DisableCosBackupOptions) SetHeaders(param map[string]string) *DisableCosBackupOptions {
	options.Headers = param
	return options
}

// EnableCosBackupOptions : The EnableCosBackup options.
type EnableCosBackupOptions struct {
	// The ID of a cluster object.
	ClusterID *string `json:"cluster_id" validate:"required,ne="`

	CosHmacKeys *CosHmacKeys `json:"cos_hmac_keys,omitempty"`

	// Endpoint of IBM Cloud Object Storage (required). In the IBM Cloud Object Storage web UI, select your bucket and you
	// can find Endpoints in the Configuration tab. Only public endpoints are supported for now. Support for private
	// endpoints will come soon.
	CosEndpoint *string `json:"cos_endpoint,omitempty"`

	// COS bucket CRN (required). In the IBM Cloud Object Storage web UI, select your bucket and you can find Bucket
	// instance CRN in the Configurations tab.
	BucketInstanceCrn *string `json:"bucket_instance_crn,omitempty"`

	Schedule *BackupSchedule `json:"schedule,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewEnableCosBackupOptions : Instantiate EnableCosBackupOptions
func (*HpdbV3) NewEnableCosBackupOptions(clusterID string) *EnableCosBackupOptions {
	return &EnableCosBackupOptions{
		ClusterID: core.StringPtr(clusterID),
	}
}

// SetClusterID : Allow user to set ClusterID
func (_options *EnableCosBackupOptions) SetClusterID(clusterID string) *EnableCosBackupOptions {
	_options.ClusterID = core.StringPtr(clusterID)
	return _options
}

// SetCosHmacKeys : Allow user to set CosHmacKeys
func (_options *EnableCosBackupOptions) SetCosHmacKeys(cosHmacKeys *CosHmacKeys) *EnableCosBackupOptions {
	_options.CosHmacKeys = cosHmacKeys
	return _options
}

// SetCosEndpoint : Allow user to set CosEndpoint
func (_options *EnableCosBackupOptions) SetCosEndpoint(cosEndpoint string) *EnableCosBackupOptions {
	_options.CosEndpoint = core.StringPtr(cosEndpoint)
	return _options
}

// SetBucketInstanceCrn : Allow user to set BucketInstanceCrn
func (_options *EnableCosBackupOptions) SetBucketInstanceCrn(bucketInstanceCrn string) *EnableCosBackupOptions {
	_options.BucketInstanceCrn = core.StringPtr(bucketInstanceCrn)
	return _options
}

// SetSchedule : Allow user to set Schedule
func (_options *EnableCosBackupOptions) SetSchedule(schedule *BackupSchedule) *EnableCosBackupOptions {
	_options.Schedule = schedule
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *EnableCosBackupOptions) SetHeaders(param map[string]string) *EnableCosBackupOptions {
	options.Headers = param
	return options
}

// GetBackupConfigOptions : The GetBackupConfig options.
type GetBackupConfigOptions struct {
	// The ID of a cluster object.
	ClusterID *string `json:"cluster_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetBackupConfigOptions : Instantiate GetBackupConfigOptions
func (*HpdbV3) NewGetBackupConfigOptions(clusterID string) *GetBackupConfigOptions {
	return &GetBackupConfigOptions{
		ClusterID: core.StringPtr(clusterID),
	}
}

// SetClusterID : Allow user to set ClusterID
func (_options *GetBackupConfigOptions) SetClusterID(clusterID string) *GetBackupConfigOptions {
	_options.ClusterID = core.StringPtr(clusterID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetBackupConfigOptions) SetHeaders(param map[string]string) *GetBackupConfigOptions {
	options.Headers = param
	return options
}

// GetBackupConfigResponse : GetBackupConfigResponse struct
type GetBackupConfigResponse struct {
	Cos *GetBackupConfigResponseCos `json:"cos,omitempty"`
}

// UnmarshalGetBackupConfigResponse unmarshals an instance of GetBackupConfigResponse from the specified map of raw messages.
func UnmarshalGetBackupConfigResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetBackupConfigResponse)
	err = core.UnmarshalModel(m, "cos", &obj.Cos, UnmarshalGetBackupConfigResponseCos)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetBackupConfigResponseCos : GetBackupConfigResponseCos struct
type GetBackupConfigResponseCos struct {
	// Endpoint of IBM Cloud Object Storage (required). In the IBM Cloud Object Storage web UI, select your bucket and you
	// can find Endpoints in the Configuration tab. Only public endpoints are supported for now. Support for private
	// endpoints will come soon.
	CosEndpoint *string `json:"cos_endpoint,omitempty"`

	// COS bucket CRN (required). In the IBM Cloud Object Storage web UI, select your bucket and you can find Bucket
	// instance CRN in the Configurations tab.
	BucketInstanceCrn *string `json:"bucket_instance_crn,omitempty"`

	Schedule *BackupSchedule `json:"schedule,omitempty"`
}

// UnmarshalGetBackupConfigResponseCos unmarshals an instance of GetBackupConfigResponseCos from the specified map of raw messages.
func UnmarshalGetBackupConfigResponseCos(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetBackupConfigResponseCos)
	err = core.UnmarshalPrimitive(m, "cos_endpoint", &obj.CosEndpoint)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "bucket_instance_crn", &obj.BucketInstanceCrn)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "schedule", &obj.Schedule, UnmarshalBackupSchedule)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetClusterOptions : The GetCluster options.
type GetClusterOptions struct {
	// The ID of a cluster object.
	ClusterID *string `json:"cluster_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetClusterOptions : Instantiate GetClusterOptions
func (*HpdbV3) NewGetClusterOptions(clusterID string) *GetClusterOptions {
	return &GetClusterOptions{
		ClusterID: core.StringPtr(clusterID),
	}
}

// SetClusterID : Allow user to set ClusterID
func (_options *GetClusterOptions) SetClusterID(clusterID string) *GetClusterOptions {
	_options.ClusterID = core.StringPtr(clusterID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetClusterOptions) SetHeaders(param map[string]string) *GetClusterOptions {
	options.Headers = param
	return options
}

// GetConfigurationOptions : The GetConfiguration options.
type GetConfigurationOptions struct {
	// The ID of a cluster object.
	ClusterID *string `json:"cluster_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetConfigurationOptions : Instantiate GetConfigurationOptions
func (*HpdbV3) NewGetConfigurationOptions(clusterID string) *GetConfigurationOptions {
	return &GetConfigurationOptions{
		ClusterID: core.StringPtr(clusterID),
	}
}

// SetClusterID : Allow user to set ClusterID
func (_options *GetConfigurationOptions) SetClusterID(clusterID string) *GetConfigurationOptions {
	_options.ClusterID = core.StringPtr(clusterID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetConfigurationOptions) SetHeaders(param map[string]string) *GetConfigurationOptions {
	options.Headers = param
	return options
}

// GetCosBackupConfigOptions : The GetCosBackupConfig options.
type GetCosBackupConfigOptions struct {
	// The ID of a cluster object.
	ClusterID *string `json:"cluster_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetCosBackupConfigOptions : Instantiate GetCosBackupConfigOptions
func (*HpdbV3) NewGetCosBackupConfigOptions(clusterID string) *GetCosBackupConfigOptions {
	return &GetCosBackupConfigOptions{
		ClusterID: core.StringPtr(clusterID),
	}
}

// SetClusterID : Allow user to set ClusterID
func (_options *GetCosBackupConfigOptions) SetClusterID(clusterID string) *GetCosBackupConfigOptions {
	_options.ClusterID = core.StringPtr(clusterID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetCosBackupConfigOptions) SetHeaders(param map[string]string) *GetCosBackupConfigOptions {
	options.Headers = param
	return options
}

// GetCosBackupConfigResponse : GetCosBackupConfigResponse struct
type GetCosBackupConfigResponse struct {
	// Endpoint of IBM Cloud Object Storage (required). In the IBM Cloud Object Storage web UI, select your bucket and you
	// can find Endpoints in the Configuration tab. Only public endpoints are supported for now. Support for private
	// endpoints will come soon.
	CosEndpoint *string `json:"cos_endpoint,omitempty"`

	// COS bucket CRN (required). In the IBM Cloud Object Storage web UI, select your bucket and you can find Bucket
	// instance CRN in the Configurations tab.
	BucketInstanceCrn *string `json:"bucket_instance_crn,omitempty"`
}

// UnmarshalGetCosBackupConfigResponse unmarshals an instance of GetCosBackupConfigResponse from the specified map of raw messages.
func UnmarshalGetCosBackupConfigResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetCosBackupConfigResponse)
	err = core.UnmarshalPrimitive(m, "cos_endpoint", &obj.CosEndpoint)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "bucket_instance_crn", &obj.BucketInstanceCrn)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetLogOptions : The GetLog options.
type GetLogOptions struct {
	// The ID of an node object.
	NodeID *string `json:"node_id" validate:"required,ne="`

	// The name of the log file.
	LogName *string `json:"log_name" validate:"required,ne="`

	// The type of the response: application/json or application/x-download.
	Accept *string `json:"Accept,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetLogOptions : Instantiate GetLogOptions
func (*HpdbV3) NewGetLogOptions(nodeID string, logName string) *GetLogOptions {
	return &GetLogOptions{
		NodeID: core.StringPtr(nodeID),
		LogName: core.StringPtr(logName),
	}
}

// SetNodeID : Allow user to set NodeID
func (_options *GetLogOptions) SetNodeID(nodeID string) *GetLogOptions {
	_options.NodeID = core.StringPtr(nodeID)
	return _options
}

// SetLogName : Allow user to set LogName
func (_options *GetLogOptions) SetLogName(logName string) *GetLogOptions {
	_options.LogName = core.StringPtr(logName)
	return _options
}

// SetAccept : Allow user to set Accept
func (_options *GetLogOptions) SetAccept(accept string) *GetLogOptions {
	_options.Accept = core.StringPtr(accept)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetLogOptions) SetHeaders(param map[string]string) *GetLogOptions {
	options.Headers = param
	return options
}

// GetTaskOptions : The GetTask options.
type GetTaskOptions struct {
	// The ID of a cluster object.
	ClusterID *string `json:"cluster_id" validate:"required,ne="`

	// The ID of a task object.
	TaskID *string `json:"task_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetTaskOptions : Instantiate GetTaskOptions
func (*HpdbV3) NewGetTaskOptions(clusterID string, taskID string) *GetTaskOptions {
	return &GetTaskOptions{
		ClusterID: core.StringPtr(clusterID),
		TaskID: core.StringPtr(taskID),
	}
}

// SetClusterID : Allow user to set ClusterID
func (_options *GetTaskOptions) SetClusterID(clusterID string) *GetTaskOptions {
	_options.ClusterID = core.StringPtr(clusterID)
	return _options
}

// SetTaskID : Allow user to set TaskID
func (_options *GetTaskOptions) SetTaskID(taskID string) *GetTaskOptions {
	_options.TaskID = core.StringPtr(taskID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetTaskOptions) SetHeaders(param map[string]string) *GetTaskOptions {
	options.Headers = param
	return options
}

// GetUserOptions : The GetUser options.
type GetUserOptions struct {
	// The ID of a cluster object.
	ClusterID *string `json:"cluster_id" validate:"required,ne="`

	// The ID of the user about which you want to get information. For MongoDB, it should be
	// 'authentication_database.username'; for example: 'mydb.syrena'. For PostgreSQL, it should be only 'username'; for
	// example: 'syrena'.
	DbUserID *string `json:"db_user_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetUserOptions : Instantiate GetUserOptions
func (*HpdbV3) NewGetUserOptions(clusterID string, dbUserID string) *GetUserOptions {
	return &GetUserOptions{
		ClusterID: core.StringPtr(clusterID),
		DbUserID: core.StringPtr(dbUserID),
	}
}

// SetClusterID : Allow user to set ClusterID
func (_options *GetUserOptions) SetClusterID(clusterID string) *GetUserOptions {
	_options.ClusterID = core.StringPtr(clusterID)
	return _options
}

// SetDbUserID : Allow user to set DbUserID
func (_options *GetUserOptions) SetDbUserID(dbUserID string) *GetUserOptions {
	_options.DbUserID = core.StringPtr(dbUserID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetUserOptions) SetHeaders(param map[string]string) *GetUserOptions {
	options.Headers = param
	return options
}

// IntegerType : Integer type parameter.
type IntegerType struct {
	// Default value of the parameter.
	Default *int64 `json:"default,omitempty"`

	// The description of the parameter.
	Description *string `json:"description,omitempty"`

	// The max value of the parameter.
	Max *int64 `json:"max,omitempty"`

	// The minimum value of the parameter.
	Min *int64 `json:"min,omitempty"`

	// Whether to restart the database server when the value of the parameter is changed.
	RequiresRestart *bool `json:"requires_restart,omitempty"`

	// Type of the value of the parameter.
	Type *string `json:"type,omitempty"`

	// The current value of the parameter.
	Value *int64 `json:"value,omitempty"`
}

// UnmarshalIntegerType unmarshals an instance of IntegerType from the specified map of raw messages.
func UnmarshalIntegerType(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(IntegerType)
	err = core.UnmarshalPrimitive(m, "default", &obj.Default)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "max", &obj.Max)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "min", &obj.Min)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "requires_restart", &obj.RequiresRestart)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListBackupsOptions : The ListBackups options.
type ListBackupsOptions struct {
	// The ID of a cluster object.
	ClusterID *string `json:"cluster_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListBackupsOptions : Instantiate ListBackupsOptions
func (*HpdbV3) NewListBackupsOptions(clusterID string) *ListBackupsOptions {
	return &ListBackupsOptions{
		ClusterID: core.StringPtr(clusterID),
	}
}

// SetClusterID : Allow user to set ClusterID
func (_options *ListBackupsOptions) SetClusterID(clusterID string) *ListBackupsOptions {
	_options.ClusterID = core.StringPtr(clusterID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListBackupsOptions) SetHeaders(param map[string]string) *ListBackupsOptions {
	options.Headers = param
	return options
}

// ListBackupsResponse : ListBackupsResponse struct
type ListBackupsResponse struct {
	// The list of backups.
	Backups []Backup `json:"backups,omitempty"`
}

// UnmarshalListBackupsResponse unmarshals an instance of ListBackupsResponse from the specified map of raw messages.
func UnmarshalListBackupsResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListBackupsResponse)
	err = core.UnmarshalModel(m, "backups", &obj.Backups, UnmarshalBackup)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListDatabasesOptions : The ListDatabases options.
type ListDatabasesOptions struct {
	// The ID of a cluster object.
	ClusterID *string `json:"cluster_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListDatabasesOptions : Instantiate ListDatabasesOptions
func (*HpdbV3) NewListDatabasesOptions(clusterID string) *ListDatabasesOptions {
	return &ListDatabasesOptions{
		ClusterID: core.StringPtr(clusterID),
	}
}

// SetClusterID : Allow user to set ClusterID
func (_options *ListDatabasesOptions) SetClusterID(clusterID string) *ListDatabasesOptions {
	_options.ClusterID = core.StringPtr(clusterID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListDatabasesOptions) SetHeaders(param map[string]string) *ListDatabasesOptions {
	options.Headers = param
	return options
}

// ListNodeLogsOptions : The ListNodeLogs options.
type ListNodeLogsOptions struct {
	// The ID of an node object.
	NodeID *string `json:"node_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListNodeLogsOptions : Instantiate ListNodeLogsOptions
func (*HpdbV3) NewListNodeLogsOptions(nodeID string) *ListNodeLogsOptions {
	return &ListNodeLogsOptions{
		NodeID: core.StringPtr(nodeID),
	}
}

// SetNodeID : Allow user to set NodeID
func (_options *ListNodeLogsOptions) SetNodeID(nodeID string) *ListNodeLogsOptions {
	_options.NodeID = core.StringPtr(nodeID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListNodeLogsOptions) SetHeaders(param map[string]string) *ListNodeLogsOptions {
	options.Headers = param
	return options
}

// ListTasksOptions : The ListTasks options.
type ListTasksOptions struct {
	// The ID of a cluster object.
	ClusterID *string `json:"cluster_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListTasksOptions : Instantiate ListTasksOptions
func (*HpdbV3) NewListTasksOptions(clusterID string) *ListTasksOptions {
	return &ListTasksOptions{
		ClusterID: core.StringPtr(clusterID),
	}
}

// SetClusterID : Allow user to set ClusterID
func (_options *ListTasksOptions) SetClusterID(clusterID string) *ListTasksOptions {
	_options.ClusterID = core.StringPtr(clusterID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListTasksOptions) SetHeaders(param map[string]string) *ListTasksOptions {
	options.Headers = param
	return options
}

// ListUsersOptions : The ListUsers options.
type ListUsersOptions struct {
	// The ID of a cluster object.
	ClusterID *string `json:"cluster_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListUsersOptions : Instantiate ListUsersOptions
func (*HpdbV3) NewListUsersOptions(clusterID string) *ListUsersOptions {
	return &ListUsersOptions{
		ClusterID: core.StringPtr(clusterID),
	}
}

// SetClusterID : Allow user to set ClusterID
func (_options *ListUsersOptions) SetClusterID(clusterID string) *ListUsersOptions {
	_options.ClusterID = core.StringPtr(clusterID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListUsersOptions) SetHeaders(param map[string]string) *ListUsersOptions {
	options.Headers = param
	return options
}

// Log : Log file name object.
type Log struct {
	// The file name of log file.
	Filename *string `json:"filename,omitempty"`

	// The size of log file, in byte.
	Size *int64 `json:"size,omitempty"`

	// The last modified date of log file.
	LastModified *string `json:"last_modified,omitempty"`
}

// UnmarshalLog unmarshals an instance of Log from the specified map of raw messages.
func UnmarshalLog(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Log)
	err = core.UnmarshalPrimitive(m, "filename", &obj.Filename)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "size", &obj.Size)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last_modified", &obj.LastModified)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// LogList : Logs object.
type LogList struct {
	// Log file list.
	Logs []Log `json:"logs,omitempty"`
}

// UnmarshalLogList unmarshals an instance of LogList from the specified map of raw messages.
func UnmarshalLogList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(LogList)
	err = core.UnmarshalModel(m, "logs", &obj.Logs, UnmarshalLog)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Node : An node information.
type Node struct {
	// The node ID.
	ID *string `json:"id,omitempty"`

	// The state of the node, such as PRIMARY or SECONDARY.
	ReplicaState *string `json:"replica_state,omitempty"`

	// The replication lag of current member in seconds.
	ReplicationLag *int64 `json:"replication_lag,omitempty"`

	// The state of the node, such as RUNNING or DELETED.
	NodeState *string `json:"node_state,omitempty"`

	// The reason why the node entered the failed state.
	Reason *string `json:"reason,omitempty"`

	// The reason why the node entered the stopped state. The possible values are '', 'EXTERNAL_KEY_DELETED',
	// 'EXTERNAL_KEY_UNKNOWN', and 'UNKNOWN'. '' means that the node was not stopped. 'EXTERNAL_KEY_DELETED' means that the
	// node was stopped because the external key was deleted. 'EXTERNAL_KEY_UNKNOWN' means that the node was stopped
	// because of the external key with unknown reason. 'UNKNOWN' means that the node was stopped with unknown reason.
	StoppedReason *string `json:"stopped_reason,omitempty"`

	// The LPAR name and port. The format is: name-port.
	Name *string `json:"name,omitempty"`

	// The UTC time when each node was created.
	CreatedAt *string `json:"created_at,omitempty"`

	// The UTC time when each node was updated.
	UpdatedAt *string `json:"updated_at,omitempty"`

	// True if the monitoring service is enabled on this node.
	IsMetricEnabled *bool `json:"is_metric_enabled,omitempty"`

	// True if the logging service is enabled on this node.
	IsLoggingEnabled *bool `json:"is_logging_enabled,omitempty"`

	// The IBM Cloud account ID.
	UserID *string `json:"user_id,omitempty"`
}

// UnmarshalNode unmarshals an instance of Node from the specified map of raw messages.
func UnmarshalNode(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Node)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "replica_state", &obj.ReplicaState)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "replication_lag", &obj.ReplicationLag)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "node_state", &obj.NodeState)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "reason", &obj.Reason)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "stopped_reason", &obj.StoppedReason)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_at", &obj.UpdatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "is_metric_enabled", &obj.IsMetricEnabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "is_logging_enabled", &obj.IsLoggingEnabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "user_id", &obj.UserID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Resources : Object of information about resources.
type Resources struct {
	// Number of CPUs. Allowed values are 1, 2, 3, 4, 5, 6, 8, 9, 12, and 16.
	Cpu *int64 `json:"cpu,omitempty"`

	// Size of memory. Allowed values are '2GiB', '3GiB', '4GiB', '5GiB', '8GiB', '12GiB', '16GiB', '24GiB', '32GiB',
	// '64GiB', '96GiB', and '128GiB'.
	Memory *string `json:"memory,omitempty"`

	// Size of storage. Allowed values are '5GiB', '10GiB', '16GiB', '24GiB', '32GiB', '64GiB', '128GiB', '160GiB',
	// '256GiB', '512GiB', '640GiB', and '1280GiB'.
	Storage *string `json:"storage,omitempty"`
}

// UnmarshalResources unmarshals an instance of Resources from the specified map of raw messages.
func UnmarshalResources(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Resources)
	err = core.UnmarshalPrimitive(m, "cpu", &obj.Cpu)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "memory", &obj.Memory)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "storage", &obj.Storage)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RestoreOptions : The Restore options.
type RestoreOptions struct {
	// The ID of a cluster object.
	ClusterID *string `json:"cluster_id" validate:"required,ne="`

	// Backup source type. Available values are "cos" and "default".
	SourceType *string `json:"source_type,omitempty"`

	CosHmacKeys *CosHmacKeys `json:"cos_hmac_keys,omitempty"`

	// Endpoint of IBM Cloud Object Storage (required). In the IBM Cloud Object Storage web UI, select your bucket and you
	// can find Endpoints in the Configuration tab. Only public endpoints are supported for now. Support for private
	// endpoints will come soon.
	CosEndpoint *string `json:"cos_endpoint,omitempty"`

	// COS bucket CRN (required for source_type cos). In the IBM Cloud Object Storage web UI, select your bucket and you
	// can find Bucket instance CRN in the Configurations tab.
	BucketInstanceCrn *string `json:"bucket_instance_crn,omitempty"`

	// The COS backup file to be restored (required for source_type cos). The backup file name identifies the date and time
	// yyyy-mm-dd-hhmmssZ (UTC) when the backup was generated.
	BackupFile *string `json:"backup_file,omitempty"`

	// The ID of the backup to be restored (required for source_type default).
	BackupID *string `json:"backup_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewRestoreOptions : Instantiate RestoreOptions
func (*HpdbV3) NewRestoreOptions(clusterID string) *RestoreOptions {
	return &RestoreOptions{
		ClusterID: core.StringPtr(clusterID),
	}
}

// SetClusterID : Allow user to set ClusterID
func (_options *RestoreOptions) SetClusterID(clusterID string) *RestoreOptions {
	_options.ClusterID = core.StringPtr(clusterID)
	return _options
}

// SetSourceType : Allow user to set SourceType
func (_options *RestoreOptions) SetSourceType(sourceType string) *RestoreOptions {
	_options.SourceType = core.StringPtr(sourceType)
	return _options
}

// SetCosHmacKeys : Allow user to set CosHmacKeys
func (_options *RestoreOptions) SetCosHmacKeys(cosHmacKeys *CosHmacKeys) *RestoreOptions {
	_options.CosHmacKeys = cosHmacKeys
	return _options
}

// SetCosEndpoint : Allow user to set CosEndpoint
func (_options *RestoreOptions) SetCosEndpoint(cosEndpoint string) *RestoreOptions {
	_options.CosEndpoint = core.StringPtr(cosEndpoint)
	return _options
}

// SetBucketInstanceCrn : Allow user to set BucketInstanceCrn
func (_options *RestoreOptions) SetBucketInstanceCrn(bucketInstanceCrn string) *RestoreOptions {
	_options.BucketInstanceCrn = core.StringPtr(bucketInstanceCrn)
	return _options
}

// SetBackupFile : Allow user to set BackupFile
func (_options *RestoreOptions) SetBackupFile(backupFile string) *RestoreOptions {
	_options.BackupFile = core.StringPtr(backupFile)
	return _options
}

// SetBackupID : Allow user to set BackupID
func (_options *RestoreOptions) SetBackupID(backupID string) *RestoreOptions {
	_options.BackupID = core.StringPtr(backupID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *RestoreOptions) SetHeaders(param map[string]string) *RestoreOptions {
	options.Headers = param
	return options
}

// ScaleResourcesOptions : The ScaleResources options.
type ScaleResourcesOptions struct {
	// The ID of a cluster object.
	ClusterID *string `json:"cluster_id" validate:"required,ne="`

	// Object of information about resources.
	Resource *Resources `json:"resource,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewScaleResourcesOptions : Instantiate ScaleResourcesOptions
func (*HpdbV3) NewScaleResourcesOptions(clusterID string) *ScaleResourcesOptions {
	return &ScaleResourcesOptions{
		ClusterID: core.StringPtr(clusterID),
	}
}

// SetClusterID : Allow user to set ClusterID
func (_options *ScaleResourcesOptions) SetClusterID(clusterID string) *ScaleResourcesOptions {
	_options.ClusterID = core.StringPtr(clusterID)
	return _options
}

// SetResource : Allow user to set Resource
func (_options *ScaleResourcesOptions) SetResource(resource *Resources) *ScaleResourcesOptions {
	_options.Resource = resource
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ScaleResourcesOptions) SetHeaders(param map[string]string) *ScaleResourcesOptions {
	options.Headers = param
	return options
}

// Task : The task details object.
type Task struct {
	// The ID of the task object.
	ID *string `json:"id,omitempty"`

	// Task type.
	Type *string `json:"type,omitempty"`

	// The UTC time when the task started.
	StartedAt *string `json:"started_at,omitempty"`

	// The UTC time when the task finished.
	FinishedAt *string `json:"finished_at,omitempty"`

	// The reason why the task entered the failed state.
	Reason *string `json:"reason,omitempty"`

	// The state of the task.
	State *string `json:"state,omitempty"`

	// Information about a task on each node.
	Nodes []TaskNode `json:"nodes,omitempty"`

	// The parameters for the task.
	Spec interface{} `json:"spec,omitempty"`
}

// UnmarshalTask unmarshals an instance of Task from the specified map of raw messages.
func UnmarshalTask(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Task)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "started_at", &obj.StartedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "finished_at", &obj.FinishedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "reason", &obj.Reason)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "nodes", &obj.Nodes, UnmarshalTaskNode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "spec", &obj.Spec)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TaskID : The ID of a task.
type TaskID struct {
	// The ID of the relevant task.
	TaskID *string `json:"task_id,omitempty"`
}

// UnmarshalTaskID unmarshals an instance of TaskID from the specified map of raw messages.
func UnmarshalTaskID(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TaskID)
	err = core.UnmarshalPrimitive(m, "task_id", &obj.TaskID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TaskItem : Task information.
type TaskItem struct {
	// The task ID.
	ID *string `json:"id,omitempty"`

	// The task type.
	Type *string `json:"type,omitempty"`

	// The task state which can be "RUNNING", "SUCCEEDED" or "FAILED".
	State *string `json:"state,omitempty"`

	// The reason why the task entered the failed state.
	Reason *string `json:"reason,omitempty"`

	// The UTC time when the task started.
	StartedAt *string `json:"started_at,omitempty"`

	// The UTC time when the task finished.
	FinishedAt *string `json:"finished_at,omitempty"`
}

// UnmarshalTaskItem unmarshals an instance of TaskItem from the specified map of raw messages.
func UnmarshalTaskItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TaskItem)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "reason", &obj.Reason)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "started_at", &obj.StartedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "finished_at", &obj.FinishedAt)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TaskNode : Information about a task on each node.
type TaskNode struct {
	// The node ID.
	ID *string `json:"id,omitempty"`

	// The state of the task on a node which can be "RUNNING", "SUCCEEDED" or "FAILED".
	State *string `json:"state,omitempty"`

	// The reason why the task entered the failed state.
	Reason *string `json:"reason,omitempty"`

	// The UTC time when the task started on the node.
	StartedAt *string `json:"started_at,omitempty"`

	// The UTC time when the task finished on the node.
	FinishedAt *string `json:"finished_at,omitempty"`
}

// UnmarshalTaskNode unmarshals an instance of TaskNode from the specified map of raw messages.
func UnmarshalTaskNode(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TaskNode)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "reason", &obj.Reason)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "started_at", &obj.StartedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "finished_at", &obj.FinishedAt)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Tasks : List of tasks running or recently run on a specified cluster.
type Tasks struct {
	// An array of tasks.
	Tasks []TaskItem `json:"tasks,omitempty"`
}

// UnmarshalTasks unmarshals an instance of Tasks from the specified map of raw messages.
func UnmarshalTasks(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Tasks)
	err = core.UnmarshalModel(m, "tasks", &obj.Tasks, UnmarshalTaskItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdateBackupConfigOptions : The UpdateBackupConfig options.
type UpdateBackupConfigOptions struct {
	// The ID of a cluster object.
	ClusterID *string `json:"cluster_id" validate:"required,ne="`

	Cos *CosBackupConfig `json:"cos,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateBackupConfigOptions : Instantiate UpdateBackupConfigOptions
func (*HpdbV3) NewUpdateBackupConfigOptions(clusterID string) *UpdateBackupConfigOptions {
	return &UpdateBackupConfigOptions{
		ClusterID: core.StringPtr(clusterID),
	}
}

// SetClusterID : Allow user to set ClusterID
func (_options *UpdateBackupConfigOptions) SetClusterID(clusterID string) *UpdateBackupConfigOptions {
	_options.ClusterID = core.StringPtr(clusterID)
	return _options
}

// SetCos : Allow user to set Cos
func (_options *UpdateBackupConfigOptions) SetCos(cos *CosBackupConfig) *UpdateBackupConfigOptions {
	_options.Cos = cos
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateBackupConfigOptions) SetHeaders(param map[string]string) *UpdateBackupConfigOptions {
	options.Headers = param
	return options
}

// UpdateConfigurationOptions : The UpdateConfiguration options.
type UpdateConfigurationOptions struct {
	// The ID of a cluster object.
	ClusterID *string `json:"cluster_id" validate:"required,ne="`

	// Object of information about configuration.
	Configuration *Configurations `json:"configuration,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateConfigurationOptions : Instantiate UpdateConfigurationOptions
func (*HpdbV3) NewUpdateConfigurationOptions(clusterID string) *UpdateConfigurationOptions {
	return &UpdateConfigurationOptions{
		ClusterID: core.StringPtr(clusterID),
	}
}

// SetClusterID : Allow user to set ClusterID
func (_options *UpdateConfigurationOptions) SetClusterID(clusterID string) *UpdateConfigurationOptions {
	_options.ClusterID = core.StringPtr(clusterID)
	return _options
}

// SetConfiguration : Allow user to set Configuration
func (_options *UpdateConfigurationOptions) SetConfiguration(configuration *Configurations) *UpdateConfigurationOptions {
	_options.Configuration = configuration
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateConfigurationOptions) SetHeaders(param map[string]string) *UpdateConfigurationOptions {
	options.Headers = param
	return options
}

// User : Object of information about user.
type User struct {
	// Name of the user.
	Name *string `json:"name,omitempty"`

	// This field is only for MongoDB. Name of authentication database of the user. For more information about
	// authentication database, see MongoDB documentation
	// (https://docs.mongodb.com/manual/core/security-users/#user-authentication-database).
	AuthDb *string `json:"auth_db,omitempty"`

	// This field is only for PostgreSQL. For more information about role attributes, see PostgreSQL documentation
	// (https://www.postgresql.org/docs/12/role-attributes.html).
	RoleAttributes []string `json:"role_attributes,omitempty"`
}

// UnmarshalUser unmarshals an instance of User from the specified map of raw messages.
func UnmarshalUser(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(User)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "auth_db", &obj.AuthDb)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "role_attributes", &obj.RoleAttributes)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UserDetails : A user information object received.
type UserDetails struct {
	// Name of the user.
	Name *string `json:"name,omitempty"`

	// This field is only for MongoDB. Name of authentication database of the user. For more information about
	// authentication database, see MongoDB documentation
	// (https://docs.mongodb.com/manual/core/security-users/#user-authentication-database).
	AuthDb *string `json:"auth_db,omitempty"`

	// Database access for the user you want to create.
	DbAccess []Access `json:"db_access,omitempty"`

	// This field is only for PostgreSQL. For more information about role attributes, see PostgreSQL documentation
	// (https://www.postgresql.org/docs/12/role-attributes.html).
	RoleAttributes []string `json:"role_attributes,omitempty"`
}

// UnmarshalUserDetails unmarshals an instance of UserDetails from the specified map of raw messages.
func UnmarshalUserDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UserDetails)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "auth_db", &obj.AuthDb)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "db_access", &obj.DbAccess, UnmarshalAccess)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "role_attributes", &obj.RoleAttributes)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Users : Object of information about users.
type Users struct {
	// User list.
	Users []User `json:"users,omitempty"`
}

// UnmarshalUsers unmarshals an instance of Users from the specified map of raw messages.
func UnmarshalUsers(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Users)
	err = core.UnmarshalModel(m, "users", &obj.Users, UnmarshalUser)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
