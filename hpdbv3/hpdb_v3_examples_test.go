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
				"9cebab98-afeb-4886-9a29-8e741716e7ff",
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
				"9cebab98-afeb-4886-9a29-8e741716e7ff",
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
				"9cebab98-afeb-4886-9a29-8e741716e7ff",
				"admin.admin",
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
				"9cebab98-afeb-4886-9a29-8e741716e7ff",
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

			scaleResourcesOptions := hpdbService.NewScaleResourcesOptions(
				"9cebab98-afeb-4886-9a29-8e741716e7ff",
			)

			cpuNumber := 2
			memory := "2gib"
			storage := "5gib"

			var resource hpdbv3.Resources
			resource.Cpu = &cpuNumber
			resource.Memory = &memory
			resource.Storage = &storage
			scaleResourcesOptions.SetResource(&resource)

			taskID, response, err := hpdbService.ScaleResources(scaleResourcesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(taskID, "", "  ")
			fmt.Println(string(b))

			// end-scale_resources

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(taskID).ToNot(BeNil())

		})
		It(`GetConfiguration request example`, func() {
			fmt.Println("\nGetConfiguration() result:")
			// begin-get_configuration

			getConfigurationOptions := hpdbService.NewGetConfigurationOptions(
				"9cebab98-afeb-4886-9a29-8e741716e7ff",
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
		It(`UpdateConfiguration request example`, func() {
			fmt.Println("\nUpdateConfiguration() result:")
			// begin-update_configuration

			updateConfigurationOptions := hpdbService.NewUpdateConfigurationOptions(
				"9cebab98-afeb-4886-9a29-8e741716e7ff",
			)

			var timeout int64 = 12000
			var config hpdbv3.Configurations
			config.DeadlockTimeout = &timeout
			updateConfigurationOptions.SetConfiguration(&config)

			taskID, response, err := hpdbService.UpdateConfiguration(updateConfigurationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(taskID, "", "  ")
			fmt.Println(string(b))

			// end-update_configuration

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(taskID).ToNot(BeNil())

		})
		It(`ListTasks request example`, func() {
			fmt.Println("\nListTasks() result:")
			// begin-list_tasks

			listTasksOptions := hpdbService.NewListTasksOptions(
				"9cebab98-afeb-4886-9a29-8e741716e7ff",
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
				"9cebab98-afeb-4886-9a29-8e741716e7ff",
				"c1a15760-a4f2-11ec-b00a-7f684d1dd53",
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
		It(`ListBackups request example`, func() {
			fmt.Println("\nListBackups() result:")
			// begin-list_backups

			listBackupsOptions := hpdbService.NewListBackupsOptions(
				"9cebab98-afeb-4886-9a29-8e741716e7ff",
			)

			listBackupsResponse, response, err := hpdbService.ListBackups(listBackupsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(listBackupsResponse, "", "  ")
			fmt.Println(string(b))

			// end-list_backups

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listBackupsResponse).ToNot(BeNil())

		})
		It(`EnableCosBackup request example`, func() {
			fmt.Println("\nEnableCosBackup() result:")
			// begin-enable_cos_backup

			enableCosBackupOptions := hpdbService.NewEnableCosBackupOptions(
				"9cebab98-afeb-4886-9a29-8e741716e7ff",
			)
			enableCosBackupOptions.SetBucketInstanceCrn("crn:v1:staging:public:cloud-object-storage:global:a/7dd00b8ebb54466fa06fe5936913d169:09dc7766-326e-4765-bb1e-1e76189fff12:bucket:files")
			enableCosBackupOptions.SetCosEndpoint("s3.us-west.cloud-object-storage.test.appdomain.cloud")

			var cosKey hpdbv3.CosHmacKeys
			keyId := "COS_ACCESS_KEY_ID"
			key := "COS_SECRET_KEY"
			cosKey.AccessKeyID = &keyId
			cosKey.SecretAccessKey = &key
			enableCosBackupOptions.SetCosHmacKeys(&cosKey)

			taskID, response, err := hpdbService.EnableCosBackup(enableCosBackupOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(taskID, "", "  ")
			fmt.Println(string(b))

			// end-enable_cos_backup

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(taskID).ToNot(BeNil())

		})
		It(`DisableCosBackup request example`, func() {
			fmt.Println("\nDisableCosBackup() result:")
			// begin-disable_cos_backup

			disableCosBackupOptions := hpdbService.NewDisableCosBackupOptions(
				"9cebab98-afeb-4886-9a29-8e741716e7ff",
			)

			taskID, response, err := hpdbService.DisableCosBackup(disableCosBackupOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(taskID, "", "  ")
			fmt.Println(string(b))

			// end-disable_cos_backup

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(taskID).ToNot(BeNil())

		})
		It(`GetCosBackupConfig request example`, func() {
			fmt.Println("\nGetCosBackupConfig() result:")
			// begin-get_cos_backup_config

			getCosBackupConfigOptions := hpdbService.NewGetCosBackupConfigOptions(
				"9cebab98-afeb-4886-9a29-8e741716e7ff",
			)

			getCosBackupConfigResponse, response, err := hpdbService.GetCosBackupConfig(getCosBackupConfigOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(getCosBackupConfigResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_cos_backup_config

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getCosBackupConfigResponse).ToNot(BeNil())

		})
		It(`GetBackupConfig request example`, func() {
			fmt.Println("\nGetBackupConfig() result:")
			// begin-get_backup_config

			getBackupConfigOptions := hpdbService.NewGetBackupConfigOptions(
				"9cebab98-afeb-4886-9a29-8e741716e7ff",
			)

			getBackupConfigResponse, response, err := hpdbService.GetBackupConfig(getBackupConfigOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(getBackupConfigResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_backup_config

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getBackupConfigResponse).ToNot(BeNil())

		})
		It(`UpdateBackupConfig request example`, func() {
			fmt.Println("\nUpdateBackupConfig() result:")
			// begin-update_backup_config

			updateBackupConfigOptions := hpdbService.NewUpdateBackupConfigOptions(
				"9cebab98-afeb-4886-9a29-8e741716e7ff",
			)

			var cos hpdbv3.CosBackupConfig
			var schedule hpdbv3.BackupSchedule

			frequency := "frequency"
			value := "1w"

			schedule.Type = &frequency
			schedule.Value = &value

			cos.Schedule = &schedule

			updateBackupConfigOptions.SetCos(&cos)

			taskID, response, err := hpdbService.UpdateBackupConfig(updateBackupConfigOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(taskID, "", "  ")
			fmt.Println(string(b))

			// end-update_backup_config

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(taskID).ToNot(BeNil())

		})
		It(`Restore request example`, func() {
			fmt.Println("\nRestore() result:")
			// begin-restore

			restoreOptions := hpdbService.NewRestoreOptions(
				"9cebab98-afeb-4886-9a29-8e741716e7ff",
			)

      restoreOptions.SetSourceType("cos")
			var cosKey hpdbv3.CosHmacKeys
			keyId := "COS_ACCESS_KEY_ID"
			key := "COS_SECRET_KEY"
			cosKey.AccessKeyID = &keyId
			cosKey.SecretAccessKey = &key
			restoreOptions.SetCosHmacKeys(&cosKey)
			restoreOptions.SetBucketInstanceCrn(config["crn:v1:staging:public:cloud-object-storage:global:a/7dd00b8ebb54466fa06fe5936913d169:09dc7766-326e-4765-bb1e-1e76189fff12:bucket:files"])
			restoreOptions.SetCosEndpoint("s3.us-west.cloud-object-storage.test.appdomain.cloud")
			restoreOptions.SetBackupFile(config["archive-2022-03-16-140004Z.tar"])

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

		})
		It(`ListNodeLogs request example`, func() {
			fmt.Println("\nListNodeLogs() result:")
			// begin-list_node_logs

			listNodeLogsOptions := hpdbService.NewListNodeLogsOptions(
				"452ebc6007955ba275cfbbe0f2a78e40",
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
				"452ebc6007955ba275cfbbe0f2a78e40",
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
