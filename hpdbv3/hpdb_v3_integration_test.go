// +build integration

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

package hpdbv3_test

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/hpdb-go-sdk/hpdbv3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the hpdbv3 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`HpdbV3 Integration Tests`, func() {

	const externalConfigFile = "../hpdb_v3.env"

	var (
		err          error
		hpdbService *hpdbv3.HpdbV3
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
			config, err = core.GetServiceProperties(hpdbv3.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			fmt.Fprintf(GinkgoWriter, "Service URL: %v\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {

			hpdbServiceOptions := &hpdbv3.HpdbV3Options{}

			hpdbService, err = hpdbv3.NewHpdbV3UsingExternalConfig(hpdbServiceOptions)

			Expect(err).To(BeNil())
			Expect(hpdbService).ToNot(BeNil())
			Expect(hpdbService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			hpdbService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`GetCluster - Get database cluster details`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetCluster(getClusterOptions *GetClusterOptions)`, func() {

			getClusterOptions := &hpdbv3.GetClusterOptions{
				ClusterID: core.StringPtr("a958e854-ab46-42d0-9b49-5aef714a36b3"),
			}

			cluster, response, err := hpdbService.GetCluster(getClusterOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(cluster).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 400
			// 401
			// 404
			// 500
			//
		})
	})

	Describe(`ListUsers - List database users`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListUsers(listUsersOptions *ListUsersOptions)`, func() {

			listUsersOptions := &hpdbv3.ListUsersOptions{
				ClusterID: core.StringPtr("a958e854-ab46-42d0-9b49-5aef714a36b3"),
			}

			users, response, err := hpdbService.ListUsers(listUsersOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(users).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 404
			// 500
			//
		})
	})

	Describe(`GetUser - Get database user details`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetUser(getUserOptions *GetUserOptions)`, func() {

			getUserOptions := &hpdbv3.GetUserOptions{
				ClusterID: core.StringPtr("a958e854-ab46-42d0-9b49-5aef714a36b3"),
				DbUserID: core.StringPtr("admin"),
			}

			userDetails, response, err := hpdbService.GetUser(getUserOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(userDetails).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 404
			// 500
			//
		})
	})

	Describe(`ListDatabases - List databases`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDatabases(listDatabasesOptions *ListDatabasesOptions)`, func() {

			listDatabasesOptions := &hpdbv3.ListDatabasesOptions{
				ClusterID: core.StringPtr("a958e854-ab46-42d0-9b49-5aef714a36b3"),
			}

			databases, response, err := hpdbService.ListDatabases(listDatabasesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(databases).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 404
			// 500
			//
		})
	})
/*
	Describe(`ScaleResources - Scale resources`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ScaleResources(scaleResourcesOptions *ScaleResourcesOptions)`, func() {

			scaleResourcesResourceModel := &hpdbv3.ScaleResourcesResource{
				Cpu: core.Int64Ptr(int64(2)),
				Memory: core.StringPtr("2GiB"),
				Storage: core.StringPtr("5GiB"),
			}

			scaleResourcesOptions := &hpdbv3.ScaleResourcesOptions{
				ClusterID: core.StringPtr("a958e854-ab46-42d0-9b49-5aef714a36b3"),
				Resource: scaleResourcesResourceModel,
			}

			scaleResourcesResponse, response, err := hpdbService.ScaleResources(scaleResourcesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(scaleResourcesResponse).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 404
			// 500
			//
		})
	})
*/
	Describe(`GetConfiguration - Get configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetConfiguration(getConfigurationOptions *GetConfigurationOptions)`, func() {

			getConfigurationOptions := &hpdbv3.GetConfigurationOptions{
				ClusterID: core.StringPtr("a958e854-ab46-42d0-9b49-5aef714a36b3"),
			}

			configuration, response, err := hpdbService.GetConfiguration(getConfigurationOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(configuration).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 404
			// 500
			//
		})
	})
/*
	Describe(`UpdateConfiguration - Update configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateConfiguration(updateConfigurationOptions *UpdateConfigurationOptions)`, func() {

			updateConfigurationDataConfigurationModel := &hpdbv3.UpdateConfigurationDataConfiguration{
				DeadlockTimeout: core.Int64Ptr(int64(10000)),
				MaxLocksPerTransaction: core.Int64Ptr(int64(100)),
				SharedBuffers: core.Int64Ptr(int64(256)),
				MaxConnections: core.Int64Ptr(int64(150)),
			}

			updateConfigurationOptions := &hpdbv3.UpdateConfigurationOptions{
				ClusterID: core.StringPtr("testString"),
				Configuration: updateConfigurationDataConfigurationModel,
			}

			updateConfigurationResponse, response, err := hpdbService.UpdateConfiguration(updateConfigurationOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(updateConfigurationResponse).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 400
			// 401
			// 500
			//
		})
	})
*/
	Describe(`ListTasks - List tasks`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListTasks(listTasksOptions *ListTasksOptions)`, func() {

			listTasksOptions := &hpdbv3.ListTasksOptions{
				ClusterID: core.StringPtr("a958e854-ab46-42d0-9b49-5aef714a36b3"),
			}

			tasks, response, err := hpdbService.ListTasks(listTasksOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(tasks).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 404
			// 500
			//
		})
	})
/*
	Describe(`GetTask - Show task`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetTask(getTaskOptions *GetTaskOptions)`, func() {

			getTaskOptions := &hpdbv3.GetTaskOptions{
				ClusterID: core.StringPtr("a958e854-ab46-42d0-9b49-5aef714a36b3"),
				TaskID: core.StringPtr("732fc8e0-da37-11eb-9433-755fe141f81f"),
			}

			task, response, err := hpdbService.GetTask(getTaskOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(task).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 404
			// 500
			//
		})
	})
*/
	Describe(`ListNodeLogs - List database log files of a node`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListNodeLogs(listNodeLogsOptions *ListNodeLogsOptions)`, func() {

			listNodeLogsOptions := &hpdbv3.ListNodeLogsOptions{
				NodeID: core.StringPtr("c5ff2d841c7e6a11de3cbaa2b992d712"),
			}

			logList, response, err := hpdbService.ListNodeLogs(listNodeLogsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(logList).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 404
			// 500
			//
		})
	})

	Describe(`GetLog - Get log details`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetLog(getLogOptions *GetLogOptions)`, func() {

			getLogOptions := &hpdbv3.GetLogOptions{
				NodeID: core.StringPtr("c5ff2d841c7e6a11de3cbaa2b992d712"),
				LogName: core.StringPtr("audit.log"),
				Accept: core.StringPtr("application/json"),
			}

			result, response, err := hpdbService.GetLog(getLogOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(result).ToNot(BeNil())

			//
			// The following status codes aren't covered by tests.
			// Please provide integration tests for these too.
			//
			// 401
			// 404
			// 500
			//
		})
	})
})

//
// Utility functions are declared in the unit test file
//
