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
	"strings"
	"testing"
	"time"

	"github.com/IBM/hpdb-go-sdk/hpdbv3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

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

func TestHpdbV3(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "HpdbV3 Suite")
}
