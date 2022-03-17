// +build cosbackup

/**
 * (C) Copyright IBM Corp. 2022.
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
	"os"
	"strings"
	"time"

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

func getTaskStatus(status chan string, hpdbService *hpdbv3.HpdbV3, taskId string, clusterId string) {
	fmt.Println("getTaskStatus: Started")
	getTaskOptions := hpdbService.NewGetTaskOptions(clusterId, taskId)
	fmt.Println("getting task status...")
	task, _, err := hpdbService.GetTask(getTaskOptions)
	fmt.Println("task status:", *task.State)
	if err != nil {
		panic(err)
	}
	for strings.ToLower(*(task.State)) == "running" {
		time.Sleep(30 * time.Second)
		fmt.Println("getting task status...")
		task, _, err = hpdbService.GetTask(getTaskOptions)
		fmt.Println("task status:", *task.State)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("getTaskStatus: Finished")
	if strings.ToLower(*task.State) != "succeeded" {
		b, _ := json.MarshalIndent(task, "", "  ")
		fmt.Println(string(b))
	}
	status <- *(task.State)
}

var _ = Describe(`HpdbV3 COS Backup/restore Tests`, func() {

	const externalConfigFile = "../hpdb_v3.env"

	var (
		hpdbService *hpdbv3.HpdbV3
		config       map[string]string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping examples...")
	}

	var clusterId string;

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

	Describe(`COS backup/restore tests`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Test enableCosBackup if COS backup is disabled`, func() {
			getClusterOptions := hpdbService.NewGetClusterOptions(
				clusterId,
			)

			cluster, response, err := hpdbService.GetCluster(getClusterOptions)
			if (*(cluster.IsCosBackupEnabled) == true) {
				Skip("COS backup is enabled, so skip enableCosBackup test")
			}

			fmt.Println("\nEnableCosBackup() result:")
			// begin-enable_backup

			enableCosBackupOptions := hpdbService.NewEnableCosBackupOptions(
				clusterId,
			)

			enableCosBackupOptions.SetBucketInstanceCrn(config["COS_CRN"])
			enableCosBackupOptions.SetCosEndpoint(config["COS_ENDPOINT"])

			var cosKey hpdbv3.CosHmacKeys
			keyId := config["COS_ACCESS_KEY_ID"]
			key := config["COS_SECRET_KEY"]
			cosKey.AccessKeyID = &keyId
			cosKey.SecretAccessKey = &key
			enableCosBackupOptions.SetCosHmacKeys(&cosKey)

			taskID, response, err := hpdbService.EnableCosBackup(enableCosBackupOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(taskID, "", "  ")
			fmt.Println(string(b))

			// end-enable_backup

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(taskID).ToNot(BeNil())

			status := make(chan string)
			go getTaskStatus(status, hpdbService, *taskID.TaskID, clusterId)
			state := <- status
			fmt.Println(state)
			Expect(strings.ToLower(state)).To(Equal("succeeded"))
		})
		It(`Test getCosBackupConfig if COS backup is enabled`, func() {
			getClusterOptions := hpdbService.NewGetClusterOptions(
				clusterId,
			)

			cluster, response, err := hpdbService.GetCluster(getClusterOptions)
			if (cluster.IsCosBackupEnabled == nil || *(cluster.IsCosBackupEnabled) != true) {
				Skip("COS backup is disabled, so skip getCosBackupConfig test")
			}

			fmt.Println("\nGetCosBackupConfig() result:")
			// begin-get_backup_config

			getCosBackupConfigOptions := hpdbService.NewGetCosBackupConfigOptions(
				clusterId,
			)

			getCosBackupConfigResponse, response, err := hpdbService.GetCosBackupConfig(getCosBackupConfigOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(getCosBackupConfigResponse, "", "  ")
			fmt.Println(string(b))
			fmt.Println(response.StatusCode)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getCosBackupConfigResponse).ToNot(BeNil())
			// end-get_backup_config

		})
		It(`Test restore if backup file exists`, func() {
			getClusterOptions := hpdbService.NewGetClusterOptions(
				clusterId,
			)

			cluster, response, err := hpdbService.GetCluster(getClusterOptions)
			if (cluster.IsCosBackupEnabled == nil || *(cluster.IsCosBackupEnabled) != true || config["COS_FILE"] == "") {
				Skip("COS backup is disabled or no backup file found, so skip restore test")
			}

			fmt.Println("\nRestore() result:")
			// begin-restore

			restoreOptions := hpdbService.NewRestoreOptions(
				clusterId,
			)
			restoreOptions.SetSourceType("cos")
			var cosKey hpdbv3.CosHmacKeys
			keyId := config["COS_ACCESS_KEY_ID"]
			key := config["COS_SECRET_KEY"]
			cosKey.AccessKeyID = &keyId
			cosKey.SecretAccessKey = &key
			restoreOptions.SetCosHmacKeys(&cosKey)
			restoreOptions.SetBucketInstanceCrn(config["COS_CRN"])
			restoreOptions.SetCosEndpoint(config["COS_ENDPOINT"])
			restoreOptions.SetBackupFile(config["COS_FILE"])

			taskID, response, err := hpdbService.Restore(restoreOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(taskID, "", "  ")
			fmt.Println(string(b))

			// end-restore

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(taskID).ToNot(BeNil())

			status := make(chan string)
			go getTaskStatus(status, hpdbService, *taskID.TaskID, clusterId)
			state := <- status
			fmt.Println(state)
			Expect(strings.ToLower(state)).To(Equal("succeeded"))

		})
		It(`Test disableCosBackup if COS backup is enabled`, func() {
			getClusterOptions := hpdbService.NewGetClusterOptions(
				clusterId,
			)

			cluster, response, err := hpdbService.GetCluster(getClusterOptions)
			if (*(cluster.IsCosBackupEnabled) == false) {
				Skip("COS backup is disabled, so skip disableCosBackup test")
			}

			fmt.Println("\nDisableCosBackup() result:")
			// begin-enable_backup

			disableCosBackupOptions := hpdbService.NewDisableCosBackupOptions(
				clusterId,
			)

			taskID, response, err := hpdbService.DisableCosBackup(disableCosBackupOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(taskID, "", "  ")
			fmt.Println(string(b))

			// end-enable_backup

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(taskID).ToNot(BeNil())

			// query disable result
			status := make(chan string)
			go getTaskStatus(status, hpdbService, *taskID.TaskID, clusterId)
			state := <- status
			fmt.Println(state)
			Expect(strings.ToLower(state)).To(Equal("succeeded"))
		})
	})
})
