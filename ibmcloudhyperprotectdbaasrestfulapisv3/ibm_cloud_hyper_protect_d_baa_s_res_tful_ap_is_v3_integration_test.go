// +build integration

/**
 * (C) Copyright IBM Corp. 2021.
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

package ibmcloudhyperprotectdbaasrestfulapisv3_test

import (
	"fmt"
	"os"

	"github.com/IBM/cloud-go-sdk/ibmcloudhyperprotectdbaasrestfulapisv3"
	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the ibmcloudhyperprotectdbaasrestfulapisv3 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`IbmCloudHyperProtectDBaaSResTfulApIsV3 Integration Tests`, func() {

	const externalConfigFile = "../ibm_cloud_hyper_protect_d_baa_s_res_tful_ap_is_v3.env"

	var (
		err          error
		ibmCloudHyperProtectDBaaSResTfulApIsService *ibmcloudhyperprotectdbaasrestfulapisv3.IbmCloudHyperProtectDBaaSResTfulApIsV3
		serviceURL   string
		config       map[string]string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping tests...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(ibmcloudhyperprotectdbaasrestfulapisv3.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			fmt.Printf("Service URL: %s\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {

			ibmCloudHyperProtectDBaaSResTfulApIsServiceOptions := &ibmcloudhyperprotectdbaasrestfulapisv3.IbmCloudHyperProtectDBaaSResTfulApIsV3Options{}

			ibmCloudHyperProtectDBaaSResTfulApIsService, err = ibmcloudhyperprotectdbaasrestfulapisv3.NewIbmCloudHyperProtectDBaaSResTfulApIsV3UsingExternalConfig(ibmCloudHyperProtectDBaaSResTfulApIsServiceOptions)

			Expect(err).To(BeNil())
			Expect(ibmCloudHyperProtectDBaaSResTfulApIsService).ToNot(BeNil())
			Expect(ibmCloudHyperProtectDBaaSResTfulApIsService.Service.Options.URL).To(Equal(serviceURL))
		})
	})

	Describe(`GetCluster - Get database cluster details`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetCluster(getClusterOptions *GetClusterOptions)`, func() {

			getClusterOptions := &ibmcloudhyperprotectdbaasrestfulapisv3.GetClusterOptions{
				ClusterID: core.StringPtr("testString"),
			}

			cluster, response, err := ibmCloudHyperProtectDBaaSResTfulApIsService.GetCluster(getClusterOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(cluster).ToNot(BeNil())

		})
	})

	Describe(`ListUsers - List database users`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListUsers(listUsersOptions *ListUsersOptions)`, func() {

			listUsersOptions := &ibmcloudhyperprotectdbaasrestfulapisv3.ListUsersOptions{
				ClusterID: core.StringPtr("testString"),
			}

			users, response, err := ibmCloudHyperProtectDBaaSResTfulApIsService.ListUsers(listUsersOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(users).ToNot(BeNil())

		})
	})

	Describe(`GetUser - Get database user details`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetUser(getUserOptions *GetUserOptions)`, func() {

			getUserOptions := &ibmcloudhyperprotectdbaasrestfulapisv3.GetUserOptions{
				ClusterID: core.StringPtr("testString"),
				DbUserID: core.StringPtr("testString"),
			}

			userDetails, response, err := ibmCloudHyperProtectDBaaSResTfulApIsService.GetUser(getUserOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(userDetails).ToNot(BeNil())

		})
	})

	Describe(`ListDatabases - List databases`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDatabases(listDatabasesOptions *ListDatabasesOptions)`, func() {

			listDatabasesOptions := &ibmcloudhyperprotectdbaasrestfulapisv3.ListDatabasesOptions{
				ClusterID: core.StringPtr("testString"),
			}

			databases, response, err := ibmCloudHyperProtectDBaaSResTfulApIsService.ListDatabases(listDatabasesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(databases).ToNot(BeNil())

		})
	})

	Describe(`ScaleResources - Scale resources`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ScaleResources(scaleResourcesOptions *ScaleResourcesOptions)`, func() {

			scaleResourcesResourceModel := &ibmcloudhyperprotectdbaasrestfulapisv3.ScaleResourcesResource{
				Cpu: core.Int64Ptr(int64(2)),
				Memory: core.StringPtr("2GiB"),
				Storage: core.StringPtr("5GiB"),
			}

			scaleResourcesOptions := &ibmcloudhyperprotectdbaasrestfulapisv3.ScaleResourcesOptions{
				ClusterID: core.StringPtr("testString"),
				Resource: scaleResourcesResourceModel,
			}

			scaleResourcesResponse, response, err := ibmCloudHyperProtectDBaaSResTfulApIsService.ScaleResources(scaleResourcesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(scaleResourcesResponse).ToNot(BeNil())

		})
	})

	Describe(`GetConfiguration - Get configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetConfiguration(getConfigurationOptions *GetConfigurationOptions)`, func() {

			getConfigurationOptions := &ibmcloudhyperprotectdbaasrestfulapisv3.GetConfigurationOptions{
				ClusterID: core.StringPtr("testString"),
			}

			configuration, response, err := ibmCloudHyperProtectDBaaSResTfulApIsService.GetConfiguration(getConfigurationOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(configuration).ToNot(BeNil())

		})
	})

	Describe(`UpdateConfiguration - Update configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateConfiguration(updateConfigurationOptions *UpdateConfigurationOptions)`, func() {

			updateConfigurationDataConfigurationModel := &ibmcloudhyperprotectdbaasrestfulapisv3.UpdateConfigurationDataConfiguration{
				DeadlockTimeout: core.Int64Ptr(int64(10000)),
				MaxLocksPerTransaction: core.Int64Ptr(int64(100)),
				SharedBuffers: core.Int64Ptr(int64(256)),
				MaxConnections: core.Int64Ptr(int64(150)),
			}

			updateConfigurationOptions := &ibmcloudhyperprotectdbaasrestfulapisv3.UpdateConfigurationOptions{
				ClusterID: core.StringPtr("testString"),
				XAuthToken: core.StringPtr("testString"),
				Configuration: updateConfigurationDataConfigurationModel,
			}

			updateConfigurationResponse, response, err := ibmCloudHyperProtectDBaaSResTfulApIsService.UpdateConfiguration(updateConfigurationOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(updateConfigurationResponse).ToNot(BeNil())

		})
	})

	Describe(`ListTasks - List tasks`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListTasks(listTasksOptions *ListTasksOptions)`, func() {

			listTasksOptions := &ibmcloudhyperprotectdbaasrestfulapisv3.ListTasksOptions{
				ClusterID: core.StringPtr("testString"),
			}

			tasks, response, err := ibmCloudHyperProtectDBaaSResTfulApIsService.ListTasks(listTasksOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(tasks).ToNot(BeNil())

		})
	})

	Describe(`GetTask - Show task`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetTask(getTaskOptions *GetTaskOptions)`, func() {

			getTaskOptions := &ibmcloudhyperprotectdbaasrestfulapisv3.GetTaskOptions{
				ClusterID: core.StringPtr("testString"),
				TaskID: core.StringPtr("testString"),
			}

			task, response, err := ibmCloudHyperProtectDBaaSResTfulApIsService.GetTask(getTaskOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(task).ToNot(BeNil())

		})
	})

	Describe(`ListNodeLogs - List database log files of a node`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListNodeLogs(listNodeLogsOptions *ListNodeLogsOptions)`, func() {

			listNodeLogsOptions := &ibmcloudhyperprotectdbaasrestfulapisv3.ListNodeLogsOptions{
				NodeID: core.StringPtr("testString"),
			}

			logList, response, err := ibmCloudHyperProtectDBaaSResTfulApIsService.ListNodeLogs(listNodeLogsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(logList).ToNot(BeNil())

		})
	})

	Describe(`GetLog - Get log details`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetLog(getLogOptions *GetLogOptions)`, func() {

			getLogOptions := &ibmcloudhyperprotectdbaasrestfulapisv3.GetLogOptions{
				NodeID: core.StringPtr("testString"),
				LogName: core.StringPtr("testString"),
				Accept: core.StringPtr("application/json"),
			}

			result, response, err := ibmCloudHyperProtectDBaaSResTfulApIsService.GetLog(getLogOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())

		})
	})
})

//
// Utility functions are declared in the unit test file
//
