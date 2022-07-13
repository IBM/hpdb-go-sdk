// +build config

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

var _ = Describe(`HpdbV3 Configuration Update Tests`, func() {

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

	Describe(`Configuration Update tests`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Test updateConfiguration`, func() {
			var timeout int64

			getConfigurationOptions := hpdbService.NewGetConfigurationOptions(
				clusterId,
			)

			configuration, response, err := hpdbService.GetConfiguration(getConfigurationOptions)

			timeout = *(configuration.Configuration.DeadlockTimeout.Value) + 1

			fmt.Println("Current DeadlockTimeout is ", *(configuration.Configuration.DeadlockTimeout.Value))
			fmt.Println("Update DeadlockTimeout to ", timeout)

			// begin-enable_backup

			updateConfigurationOptions := hpdbService.NewUpdateConfigurationOptions(
				clusterId,
			)

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
	})
})
