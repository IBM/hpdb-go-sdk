// +build examples

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

package hpdbv3_test

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/hpdb-go-sdk/hpdbv3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the IBM Cloud Hyper Protect DBaaS RESTful APIs service.
//
// The following configuration properties are assumed to be defined:
// HPDB_URL=<service base url>
// HPDB_AUTH_TYPE=iam
// HPDB_APIKEY=<IAM apikey>
// HPDB_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
const externalConfigFile = "../hpdb.env"

var (
	hpdbService  *hpdbv3.HPDBV3
	config       map[string]string
	configLoaded bool = false
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping tests...")
	}
}

var _ = Describe(`HPDBV3 Examples Tests`, func() {
	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(hpdbv3.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}

			configLoaded = len(config) > 0
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			hpdbServiceOptions := &hpdbv3.HPDBV3Options{}

			hpdbService, err = hpdbv3.NewHPDBV3UsingExternalConfig(hpdbServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(hpdbService).ToNot(BeNil())
		})
	})

	Describe(`HPDBV3 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetCluster request example`, func() {
			fmt.Println("\nGetCluster() result:")
			// begin-get_cluster

			getClusterOptions := hpdbService.NewGetClusterOptions(
				"a958e854-ab46-42d0-9b49-5aef714a36b3",
			)

			cluster, response, err := hpdbService.GetCluster(getClusterOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(cluster, "", "  ")
			fmt.Println(string(b))

			// end-get_cluster

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(cluster).ToNot(BeNil())

		})
		It(`ListUsers request example`, func() {
			fmt.Println("\nListUsers() result:")
			// begin-list_users

			listUsersOptions := hpdbService.NewListUsersOptions(
				"a958e854-ab46-42d0-9b49-5aef714a36b3",
			)

			users, response, err := hpdbService.ListUsers(listUsersOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(users, "", "  ")
			fmt.Println(string(b))

			// end-list_users

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(users).ToNot(BeNil())

		})
		It(`GetUser request example`, func() {
			fmt.Println("\nGetUser() result:")
			// begin-get_user

			getUserOptions := hpdbService.NewGetUserOptions(
				"a958e854-ab46-42d0-9b49-5aef714a36b3",
				"admin",
			)

			userDetails, response, err := hpdbService.GetUser(getUserOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(userDetails, "", "  ")
			fmt.Println(string(b))

			// end-get_user

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(userDetails).ToNot(BeNil())

		})
		It(`ListDatabases request example`, func() {
			fmt.Println("\nListDatabases() result:")
			// begin-list_databases

			listDatabasesOptions := hpdbService.NewListDatabasesOptions(
				"a958e854-ab46-42d0-9b49-5aef714a36b3",
			)

			databases, response, err := hpdbService.ListDatabases(listDatabasesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(databases, "", "  ")
			fmt.Println(string(b))

			// end-list_databases

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(databases).ToNot(BeNil())

		})
		It(`ScaleResources request example`, func() {
			fmt.Println("\nScaleResources() result:")
			// begin-scale_resources

			scaleResourcesResourceModel := &hpdbv3.ScaleResourcesResource{
				Cpu:     core.Int64Ptr(int64(1)),
				Memory:  core.StringPtr("2GiB"),
				Storage: core.StringPtr("5GiB"),
			}

			scaleResourcesOptions := &hpdbv3.ScaleResourcesOptions{
				ClusterID: core.StringPtr("a958e854-ab46-42d0-9b49-5aef714a36b3"),
				Resource:  scaleResourcesResourceModel,
			}

			scaleResourcesResponse, response, err := hpdbService.ScaleResources(scaleResourcesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(scaleResourcesResponse, "", "  ")
			fmt.Println(string(b))

			// end-scale_resources

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(scaleResourcesResponse).ToNot(BeNil())

		})
		It(`GetConfiguration request example`, func() {
			fmt.Println("\nGetConfiguration() result:")
			// begin-get_configuration

			getConfigurationOptions := hpdbService.NewGetConfigurationOptions(
				"a958e854-ab46-42d0-9b49-5aef714a36b3",
			)

			configuration, response, err := hpdbService.GetConfiguration(getConfigurationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(configuration, "", "  ")
			fmt.Println(string(b))

			// end-get_configuration

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(configuration).ToNot(BeNil())

		})
		/* UpdateConfiguration should run after ScaleResource completed */
		/*
			It(`UpdateConfiguration request example`, func() {
				fmt.Println("\nUpdateConfiguration() result:")
				// begin-update_configuration

				updateConfigurationDataConfigurationModel := &hpdbv3.UpdateConfigurationDataConfiguration{
					DeadlockTimeout:        core.Int64Ptr(int64(10000)),
					MaxLocksPerTransaction: core.Int64Ptr(int64(100)),
					SharedBuffers:          core.Int64Ptr(int64(256)),
					MaxConnections:         core.Int64Ptr(int64(201)),
				}

				updateConfigurationOptions := &hpdbv3.UpdateConfigurationOptions{
					ClusterID:     core.StringPtr("a958e854-ab46-42d0-9b49-5aef714a36b3"),
					Configuration: updateConfigurationDataConfigurationModel,
				}

				updateConfigurationResponse, response, err := hpdbService.UpdateConfiguration(updateConfigurationOptions)
				if err != nil {
					panic(err)
				}
				b, _ := json.MarshalIndent(updateConfigurationResponse, "", "  ")
				fmt.Println(string(b))

				// end-update_configuration

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(202))
				Expect(updateConfigurationResponse).ToNot(BeNil())

			})
		*/
		It(`ListTasks request example`, func() {
			fmt.Println("\nListTasks() result:")
			// begin-list_tasks

			listTasksOptions := hpdbService.NewListTasksOptions(
				"a958e854-ab46-42d0-9b49-5aef714a36b3",
			)

			tasks, response, err := hpdbService.ListTasks(listTasksOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(tasks, "", "  ")
			fmt.Println(string(b))

			// end-list_tasks

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(tasks).ToNot(BeNil())

		})
		It(`GetTask request example`, func() {
			fmt.Println("\nGetTask() result:")
			// begin-get_task

			getTaskOptions := hpdbService.NewGetTaskOptions(
				"a958e854-ab46-42d0-9b49-5aef714a36b3",
				"732fc8e0-da37-11eb-9433-755fe141f81f",
			)

			task, response, err := hpdbService.GetTask(getTaskOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(task, "", "  ")
			fmt.Println(string(b))

			// end-get_task

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(task).ToNot(BeNil())

		})
		It(`ListNodeLogs request example`, func() {
			fmt.Println("\nListNodeLogs() result:")
			// begin-list_node_logs

			listNodeLogsOptions := hpdbService.NewListNodeLogsOptions(
				"c5ff2d841c7e6a11de3cbaa2b992d712",
			)

			logList, response, err := hpdbService.ListNodeLogs(listNodeLogsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(logList, "", "  ")
			fmt.Println(string(b))

			// end-list_node_logs

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(logList).ToNot(BeNil())

		})
		It(`GetLog request example`, func() {
			fmt.Println("\nGetLog() result:")
			// begin-get_log

			getLogOptions := hpdbService.NewGetLogOptions(
				"c5ff2d841c7e6a11de3cbaa2b992d712",
				"audit.log",
			)

			getLogOptions.SetAccept("application/json")

			file, response, err := hpdbService.GetLog(getLogOptions)
			if err != nil {
				panic(err)
			}
			if file != nil {
				defer file.Close()
				outFile, err := os.Create("file.out")
				if err != nil {
					panic(err)
				}
				defer outFile.Close()
				_, err = io.Copy(outFile, file)
				if err != nil {
					panic(err)
				}
			}

			// end-get_log

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(file).ToNot(BeNil())

		})
	})
})
