// +build examples

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
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/hpdb-go-sdk/hpdbv3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the hpdb service.
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
var _ = Describe(`HpdbV3 Examples Tests`, func() {

	const externalConfigFile = "../hpdb_v3.env"

	var (
		hpdbService *hpdbv3.HpdbV3
		config       map[string]string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping examples...")
	}

	var clusterId string;
	var dbType string;

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping examples: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(hpdbv3.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping examples: " + err.Error())
			} else if len(config) == 0 {
				Skip("Unable to load service properties, skipping examples")
			}

			crnSegments := strings.Split(config["CLUSTER_CRN"], ":")
			clusterId = crnSegments[7]
			dbType = crnSegments[4]

			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			hpdbServiceOptions := &hpdbv3.HpdbV3Options{}

			hpdbService, err = hpdbv3.NewHpdbV3UsingExternalConfig(hpdbServiceOptions)

			if err != nil {
				panic(err)
			}

			hpdbService.Service.DisableSSLVerification()

			// end-common

			Expect(hpdbService).ToNot(BeNil())
		})
	})

	Describe(`HpdbV3 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetCluster request example`, func() {
			fmt.Println("\nGetCluster() result:")
			// begin-get_cluster

			getClusterOptions := hpdbService.NewGetClusterOptions(
				clusterId,
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
				clusterId,
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
			userName := "admin"
			if dbType == "hyperp-dbaas-mongodb" {
				userName = "admin.admin"
			}

			getUserOptions := hpdbService.NewGetUserOptions(
				clusterId,
				userName,
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
				clusterId,
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
		It(`GetConfiguration request example`, func() {
			if dbType == "hyperp-dbaas-mongodb" {
				Skip("Skip GetConfiguration test for mongodb clusters")
			}
			fmt.Println("\nGetConfiguration() result:")
			// begin-get_configuration

			getConfigurationOptions := hpdbService.NewGetConfigurationOptions(
				clusterId,
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
		It(`ListTasks request example`, func() {
			fmt.Println("\nListTasks() result:")
			// begin-list_tasks

			listTasksOptions := hpdbService.NewListTasksOptions(
				clusterId,
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

		})
		It(`ListNodeLogs request example`, func() {
			fmt.Println("\nListNodeLogs() result:")
			// begin-list_node_logs

			listNodeLogsOptions := hpdbService.NewListNodeLogsOptions(
				config["NODE_ID"],
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
				config["NODE_ID"],
				"audit.log",
			)

			file, response, err := hpdbService.GetLog(getLogOptions)
			if err != nil {
				panic(err)
			}
			if file != nil {
				defer file.Close()
				outFile, err := os.Create("file.out")
				if err != nil { panic(err) }
				defer outFile.Close()
				_, err = io.Copy(outFile, file)
				if err != nil { panic(err) }
			}

			// end-get_log

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(file).ToNot(BeNil())

		})
	})
})
