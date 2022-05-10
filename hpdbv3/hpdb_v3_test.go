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
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/hpdb-go-sdk/hpdbv3"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`HpdbV3`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(hpdbService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(hpdbService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
				URL: "https://hpdbv3/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(hpdbService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"HPDB_URL": "https://hpdbv3/api",
				"HPDB_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				hpdbService, serviceErr := hpdbv3.NewHpdbV3UsingExternalConfig(&hpdbv3.HpdbV3Options{
				})
				Expect(hpdbService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := hpdbService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != hpdbService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(hpdbService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(hpdbService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				hpdbService, serviceErr := hpdbv3.NewHpdbV3UsingExternalConfig(&hpdbv3.HpdbV3Options{
					URL: "https://testService/api",
				})
				Expect(hpdbService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := hpdbService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != hpdbService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(hpdbService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(hpdbService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				hpdbService, serviceErr := hpdbv3.NewHpdbV3UsingExternalConfig(&hpdbv3.HpdbV3Options{
				})
				err := hpdbService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := hpdbService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != hpdbService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(hpdbService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(hpdbService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"HPDB_URL": "https://hpdbv3/api",
				"HPDB_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			hpdbService, serviceErr := hpdbv3.NewHpdbV3UsingExternalConfig(&hpdbv3.HpdbV3Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(hpdbService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"HPDB_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			hpdbService, serviceErr := hpdbv3.NewHpdbV3UsingExternalConfig(&hpdbv3.HpdbV3Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(hpdbService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = hpdbv3.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`Parameterized URL tests`, func() {
		It(`Format parameterized URL with all default values`, func() {
			constructedURL, err := hpdbv3.ConstructServiceURL(nil)
			Expect(constructedURL).To(Equal("https://dbaas900.hyperp-dbaas.cloud.ibm.com/api/v3/unknown"))
			Expect(constructedURL).ToNot(BeNil())
			Expect(err).To(BeNil())
		})
		It(`Return an error if a provided variable name is invalid`, func() {
			var providedUrlVariables = map[string]string{
				"invalid_variable_name": "value",
			}
			constructedURL, err := hpdbv3.ConstructServiceURL(providedUrlVariables)
			Expect(constructedURL).To(Equal(""))
			Expect(err).ToNot(BeNil())
		})
	})
	Describe(`GetCluster(getClusterOptions *GetClusterOptions) - Operation response error`, func() {
		getClusterPath := "/clusters/9cebab98-afeb-4886-9a29-8e741716e7ff"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getClusterPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCluster with error: Operation response processing error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the GetClusterOptions model
				getClusterOptionsModel := new(hpdbv3.GetClusterOptions)
				getClusterOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				getClusterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := hpdbService.GetCluster(getClusterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				hpdbService.EnableRetries(0, 0)
				result, response, operationErr = hpdbService.GetCluster(getClusterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCluster(getClusterOptions *GetClusterOptions)`, func() {
		getClusterPath := "/clusters/9cebab98-afeb-4886-9a29-8e741716e7ff"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getClusterPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "9cebab98-afeb-4886-9a29-8e741716e7ff", "crn": "crn:v1:staging:public:hyperp-dbaas-mongodb:us-east:a/23a24a3e3fe7a115473f07be1c44bdb5:9cebab98-afeb-4886-9a29-8e741716e7ff::", "customer_monitoring_status": "enabled", "is_cos_backup_enabled": true, "region": "us-east", "name": "cluster01", "state": "PROVISIONED", "reason": "Reason", "db_type": "mongodb", "db_version": "mongodb-enterprise 4.4.6", "public_endpoint": "dbaas226.hyperp-dbaas.cloud.ibm.com:21089,dbaas227.hyperp-dbaas.cloud.ibm.com:21451,dbaas228.hyperp-dbaas.cloud.ibm.com:20290", "private_endpoint": "dbaas905.private.hyperp-dbaas.cloud.ibm.com:29494", "private_endpoint_type": "vpe", "plan_id": "1563f33f-874f-4d83-b6de-09384cdfb1ef", "last_active": 1647843913002, "log_url": "LogURL", "metric_url": "MetricURL", "replica_count": 3, "user_id": "23a24a3e3fe7a115473f07be1c44bdb5", "resource": {"cpu": 1, "memory": "2gib", "storage": "5gib", "storage_used": "0.19gib"}, "external_key": {"kms_instance": "KmsInstance", "kms_key": "KmsKey"}, "nodes": [{"id": "c5ff2d841c7e6a11de3cbaa2b992d712", "replica_state": "PRIMARY", "replication_lag": 0, "node_state": "RUNNING", "reason": "Reason", "stopped_reason": "EXTERNAL_KEY_DELETED", "name": "dbaas55-29247", "created_at": "2021-06-29T07:46:56Z", "updated_at": "2021-06-29T07:48:11Z", "is_metric_enabled": false, "is_logging_enabled": false, "user_id": "23a24a3e3fe7a115473f07be1c44bdb5"}], "created_at": "2021-06-29T07:46:51Z", "updated_at": "2021-06-29T07:48:11Z"}`)
				}))
			})
			It(`Invoke GetCluster successfully with retries`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())
				hpdbService.EnableRetries(0, 0)

				// Construct an instance of the GetClusterOptions model
				getClusterOptionsModel := new(hpdbv3.GetClusterOptions)
				getClusterOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				getClusterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := hpdbService.GetClusterWithContext(ctx, getClusterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				hpdbService.DisableRetries()
				result, response, operationErr := hpdbService.GetCluster(getClusterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = hpdbService.GetClusterWithContext(ctx, getClusterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getClusterPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "9cebab98-afeb-4886-9a29-8e741716e7ff", "crn": "crn:v1:staging:public:hyperp-dbaas-mongodb:us-east:a/23a24a3e3fe7a115473f07be1c44bdb5:9cebab98-afeb-4886-9a29-8e741716e7ff::", "customer_monitoring_status": "enabled", "is_cos_backup_enabled": true, "region": "us-east", "name": "cluster01", "state": "PROVISIONED", "reason": "Reason", "db_type": "mongodb", "db_version": "mongodb-enterprise 4.4.6", "public_endpoint": "dbaas226.hyperp-dbaas.cloud.ibm.com:21089,dbaas227.hyperp-dbaas.cloud.ibm.com:21451,dbaas228.hyperp-dbaas.cloud.ibm.com:20290", "private_endpoint": "dbaas905.private.hyperp-dbaas.cloud.ibm.com:29494", "private_endpoint_type": "vpe", "plan_id": "1563f33f-874f-4d83-b6de-09384cdfb1ef", "last_active": 1647843913002, "log_url": "LogURL", "metric_url": "MetricURL", "replica_count": 3, "user_id": "23a24a3e3fe7a115473f07be1c44bdb5", "resource": {"cpu": 1, "memory": "2gib", "storage": "5gib", "storage_used": "0.19gib"}, "external_key": {"kms_instance": "KmsInstance", "kms_key": "KmsKey"}, "nodes": [{"id": "c5ff2d841c7e6a11de3cbaa2b992d712", "replica_state": "PRIMARY", "replication_lag": 0, "node_state": "RUNNING", "reason": "Reason", "stopped_reason": "EXTERNAL_KEY_DELETED", "name": "dbaas55-29247", "created_at": "2021-06-29T07:46:56Z", "updated_at": "2021-06-29T07:48:11Z", "is_metric_enabled": false, "is_logging_enabled": false, "user_id": "23a24a3e3fe7a115473f07be1c44bdb5"}], "created_at": "2021-06-29T07:46:51Z", "updated_at": "2021-06-29T07:48:11Z"}`)
				}))
			})
			It(`Invoke GetCluster successfully`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := hpdbService.GetCluster(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetClusterOptions model
				getClusterOptionsModel := new(hpdbv3.GetClusterOptions)
				getClusterOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				getClusterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = hpdbService.GetCluster(getClusterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetCluster with error: Operation validation and request error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the GetClusterOptions model
				getClusterOptionsModel := new(hpdbv3.GetClusterOptions)
				getClusterOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				getClusterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := hpdbService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := hpdbService.GetCluster(getClusterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetClusterOptions model with no property values
				getClusterOptionsModelNew := new(hpdbv3.GetClusterOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = hpdbService.GetCluster(getClusterOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetCluster successfully`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the GetClusterOptions model
				getClusterOptionsModel := new(hpdbv3.GetClusterOptions)
				getClusterOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				getClusterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := hpdbService.GetCluster(getClusterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListUsers(listUsersOptions *ListUsersOptions) - Operation response error`, func() {
		listUsersPath := "/clusters/9cebab98-afeb-4886-9a29-8e741716e7ff/users"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listUsersPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListUsers with error: Operation response processing error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the ListUsersOptions model
				listUsersOptionsModel := new(hpdbv3.ListUsersOptions)
				listUsersOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				listUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := hpdbService.ListUsers(listUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				hpdbService.EnableRetries(0, 0)
				result, response, operationErr = hpdbService.ListUsers(listUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListUsers(listUsersOptions *ListUsersOptions)`, func() {
		listUsersPath := "/clusters/9cebab98-afeb-4886-9a29-8e741716e7ff/users"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listUsersPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"users": [{"name": "admin", "auth_db": "admin", "role_attributes": ["CREATEDB"]}]}`)
				}))
			})
			It(`Invoke ListUsers successfully with retries`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())
				hpdbService.EnableRetries(0, 0)

				// Construct an instance of the ListUsersOptions model
				listUsersOptionsModel := new(hpdbv3.ListUsersOptions)
				listUsersOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				listUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := hpdbService.ListUsersWithContext(ctx, listUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				hpdbService.DisableRetries()
				result, response, operationErr := hpdbService.ListUsers(listUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = hpdbService.ListUsersWithContext(ctx, listUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listUsersPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"users": [{"name": "admin", "auth_db": "admin", "role_attributes": ["CREATEDB"]}]}`)
				}))
			})
			It(`Invoke ListUsers successfully`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := hpdbService.ListUsers(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListUsersOptions model
				listUsersOptionsModel := new(hpdbv3.ListUsersOptions)
				listUsersOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				listUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = hpdbService.ListUsers(listUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListUsers with error: Operation validation and request error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the ListUsersOptions model
				listUsersOptionsModel := new(hpdbv3.ListUsersOptions)
				listUsersOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				listUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := hpdbService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := hpdbService.ListUsers(listUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListUsersOptions model with no property values
				listUsersOptionsModelNew := new(hpdbv3.ListUsersOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = hpdbService.ListUsers(listUsersOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListUsers successfully`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the ListUsersOptions model
				listUsersOptionsModel := new(hpdbv3.ListUsersOptions)
				listUsersOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				listUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := hpdbService.ListUsers(listUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetUser(getUserOptions *GetUserOptions) - Operation response error`, func() {
		getUserPath := "/clusters/9cebab98-afeb-4886-9a29-8e741716e7ff/users/admin.admin"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getUserPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetUser with error: Operation response processing error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the GetUserOptions model
				getUserOptionsModel := new(hpdbv3.GetUserOptions)
				getUserOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				getUserOptionsModel.DbUserID = core.StringPtr("admin.admin")
				getUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := hpdbService.GetUser(getUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				hpdbService.EnableRetries(0, 0)
				result, response, operationErr = hpdbService.GetUser(getUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetUser(getUserOptions *GetUserOptions)`, func() {
		getUserPath := "/clusters/9cebab98-afeb-4886-9a29-8e741716e7ff/users/admin.admin"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getUserPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "admin", "auth_db": "admin", "db_access": [{"db": "admin", "privileges": ["readWrite"]}], "role_attributes": ["CREATEDB"]}`)
				}))
			})
			It(`Invoke GetUser successfully with retries`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())
				hpdbService.EnableRetries(0, 0)

				// Construct an instance of the GetUserOptions model
				getUserOptionsModel := new(hpdbv3.GetUserOptions)
				getUserOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				getUserOptionsModel.DbUserID = core.StringPtr("admin.admin")
				getUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := hpdbService.GetUserWithContext(ctx, getUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				hpdbService.DisableRetries()
				result, response, operationErr := hpdbService.GetUser(getUserOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = hpdbService.GetUserWithContext(ctx, getUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getUserPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "admin", "auth_db": "admin", "db_access": [{"db": "admin", "privileges": ["readWrite"]}], "role_attributes": ["CREATEDB"]}`)
				}))
			})
			It(`Invoke GetUser successfully`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := hpdbService.GetUser(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetUserOptions model
				getUserOptionsModel := new(hpdbv3.GetUserOptions)
				getUserOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				getUserOptionsModel.DbUserID = core.StringPtr("admin.admin")
				getUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = hpdbService.GetUser(getUserOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetUser with error: Operation validation and request error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the GetUserOptions model
				getUserOptionsModel := new(hpdbv3.GetUserOptions)
				getUserOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				getUserOptionsModel.DbUserID = core.StringPtr("admin.admin")
				getUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := hpdbService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := hpdbService.GetUser(getUserOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetUserOptions model with no property values
				getUserOptionsModelNew := new(hpdbv3.GetUserOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = hpdbService.GetUser(getUserOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetUser successfully`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the GetUserOptions model
				getUserOptionsModel := new(hpdbv3.GetUserOptions)
				getUserOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				getUserOptionsModel.DbUserID = core.StringPtr("admin.admin")
				getUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := hpdbService.GetUser(getUserOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListDatabases(listDatabasesOptions *ListDatabasesOptions) - Operation response error`, func() {
		listDatabasesPath := "/clusters/9cebab98-afeb-4886-9a29-8e741716e7ff/databases"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDatabasesPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListDatabases with error: Operation response processing error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the ListDatabasesOptions model
				listDatabasesOptionsModel := new(hpdbv3.ListDatabasesOptions)
				listDatabasesOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				listDatabasesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := hpdbService.ListDatabases(listDatabasesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				hpdbService.EnableRetries(0, 0)
				result, response, operationErr = hpdbService.ListDatabases(listDatabasesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListDatabases(listDatabasesOptions *ListDatabasesOptions)`, func() {
		listDatabasesPath := "/clusters/9cebab98-afeb-4886-9a29-8e741716e7ff/databases"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDatabasesPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_size": 8084615, "databases": [{"name": "admin", "size_on_disk": 8084615}]}`)
				}))
			})
			It(`Invoke ListDatabases successfully with retries`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())
				hpdbService.EnableRetries(0, 0)

				// Construct an instance of the ListDatabasesOptions model
				listDatabasesOptionsModel := new(hpdbv3.ListDatabasesOptions)
				listDatabasesOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				listDatabasesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := hpdbService.ListDatabasesWithContext(ctx, listDatabasesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				hpdbService.DisableRetries()
				result, response, operationErr := hpdbService.ListDatabases(listDatabasesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = hpdbService.ListDatabasesWithContext(ctx, listDatabasesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDatabasesPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_size": 8084615, "databases": [{"name": "admin", "size_on_disk": 8084615}]}`)
				}))
			})
			It(`Invoke ListDatabases successfully`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := hpdbService.ListDatabases(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListDatabasesOptions model
				listDatabasesOptionsModel := new(hpdbv3.ListDatabasesOptions)
				listDatabasesOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				listDatabasesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = hpdbService.ListDatabases(listDatabasesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListDatabases with error: Operation validation and request error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the ListDatabasesOptions model
				listDatabasesOptionsModel := new(hpdbv3.ListDatabasesOptions)
				listDatabasesOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				listDatabasesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := hpdbService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := hpdbService.ListDatabases(listDatabasesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListDatabasesOptions model with no property values
				listDatabasesOptionsModelNew := new(hpdbv3.ListDatabasesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = hpdbService.ListDatabases(listDatabasesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListDatabases successfully`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the ListDatabasesOptions model
				listDatabasesOptionsModel := new(hpdbv3.ListDatabasesOptions)
				listDatabasesOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				listDatabasesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := hpdbService.ListDatabases(listDatabasesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ScaleResources(scaleResourcesOptions *ScaleResourcesOptions) - Operation response error`, func() {
		scaleResourcesPath := "/clusters/9cebab98-afeb-4886-9a29-8e741716e7ff/resource"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(scaleResourcesPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ScaleResources with error: Operation response processing error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the Resources model
				resourcesModel := new(hpdbv3.Resources)
				resourcesModel.Cpu = core.Int64Ptr(int64(2))
				resourcesModel.Memory = core.StringPtr("2GiB")
				resourcesModel.Storage = core.StringPtr("5GiB")

				// Construct an instance of the ScaleResourcesOptions model
				scaleResourcesOptionsModel := new(hpdbv3.ScaleResourcesOptions)
				scaleResourcesOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				scaleResourcesOptionsModel.Resource = resourcesModel
				scaleResourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := hpdbService.ScaleResources(scaleResourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				hpdbService.EnableRetries(0, 0)
				result, response, operationErr = hpdbService.ScaleResources(scaleResourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ScaleResources(scaleResourcesOptions *ScaleResourcesOptions)`, func() {
		scaleResourcesPath := "/clusters/9cebab98-afeb-4886-9a29-8e741716e7ff/resource"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(scaleResourcesPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"task_id": "1e902f30-da1b-11eb-9433-755fe141f81f"}`)
				}))
			})
			It(`Invoke ScaleResources successfully with retries`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())
				hpdbService.EnableRetries(0, 0)

				// Construct an instance of the Resources model
				resourcesModel := new(hpdbv3.Resources)
				resourcesModel.Cpu = core.Int64Ptr(int64(2))
				resourcesModel.Memory = core.StringPtr("2GiB")
				resourcesModel.Storage = core.StringPtr("5GiB")

				// Construct an instance of the ScaleResourcesOptions model
				scaleResourcesOptionsModel := new(hpdbv3.ScaleResourcesOptions)
				scaleResourcesOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				scaleResourcesOptionsModel.Resource = resourcesModel
				scaleResourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := hpdbService.ScaleResourcesWithContext(ctx, scaleResourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				hpdbService.DisableRetries()
				result, response, operationErr := hpdbService.ScaleResources(scaleResourcesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = hpdbService.ScaleResourcesWithContext(ctx, scaleResourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(scaleResourcesPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"task_id": "1e902f30-da1b-11eb-9433-755fe141f81f"}`)
				}))
			})
			It(`Invoke ScaleResources successfully`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := hpdbService.ScaleResources(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Resources model
				resourcesModel := new(hpdbv3.Resources)
				resourcesModel.Cpu = core.Int64Ptr(int64(2))
				resourcesModel.Memory = core.StringPtr("2GiB")
				resourcesModel.Storage = core.StringPtr("5GiB")

				// Construct an instance of the ScaleResourcesOptions model
				scaleResourcesOptionsModel := new(hpdbv3.ScaleResourcesOptions)
				scaleResourcesOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				scaleResourcesOptionsModel.Resource = resourcesModel
				scaleResourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = hpdbService.ScaleResources(scaleResourcesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ScaleResources with error: Operation validation and request error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the Resources model
				resourcesModel := new(hpdbv3.Resources)
				resourcesModel.Cpu = core.Int64Ptr(int64(2))
				resourcesModel.Memory = core.StringPtr("2GiB")
				resourcesModel.Storage = core.StringPtr("5GiB")

				// Construct an instance of the ScaleResourcesOptions model
				scaleResourcesOptionsModel := new(hpdbv3.ScaleResourcesOptions)
				scaleResourcesOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				scaleResourcesOptionsModel.Resource = resourcesModel
				scaleResourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := hpdbService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := hpdbService.ScaleResources(scaleResourcesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ScaleResourcesOptions model with no property values
				scaleResourcesOptionsModelNew := new(hpdbv3.ScaleResourcesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = hpdbService.ScaleResources(scaleResourcesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(202)
				}))
			})
			It(`Invoke ScaleResources successfully`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the Resources model
				resourcesModel := new(hpdbv3.Resources)
				resourcesModel.Cpu = core.Int64Ptr(int64(2))
				resourcesModel.Memory = core.StringPtr("2GiB")
				resourcesModel.Storage = core.StringPtr("5GiB")

				// Construct an instance of the ScaleResourcesOptions model
				scaleResourcesOptionsModel := new(hpdbv3.ScaleResourcesOptions)
				scaleResourcesOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				scaleResourcesOptionsModel.Resource = resourcesModel
				scaleResourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := hpdbService.ScaleResources(scaleResourcesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetConfiguration(getConfigurationOptions *GetConfigurationOptions) - Operation response error`, func() {
		getConfigurationPath := "/clusters/9cebab98-afeb-4886-9a29-8e741716e7ff/configuration"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getConfigurationPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetConfiguration with error: Operation response processing error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the GetConfigurationOptions model
				getConfigurationOptionsModel := new(hpdbv3.GetConfigurationOptions)
				getConfigurationOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				getConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := hpdbService.GetConfiguration(getConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				hpdbService.EnableRetries(0, 0)
				result, response, operationErr = hpdbService.GetConfiguration(getConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetConfiguration(getConfigurationOptions *GetConfigurationOptions)`, func() {
		getConfigurationPath := "/clusters/9cebab98-afeb-4886-9a29-8e741716e7ff/configuration"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getConfigurationPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"configuration": {"deadlock_timeout": {"default": 7, "description": "Description", "max": 3, "min": 3, "requires_restart": false, "type": "Type", "value": 5}, "max_locks_per_transaction": {"default": 7, "description": "Description", "max": 3, "min": 3, "requires_restart": false, "type": "Type", "value": 5}, "shared_buffers": {"default": 7, "description": "Description", "max": 3, "min": 3, "requires_restart": false, "type": "Type", "value": 5}, "max_connections": {"default": 7, "description": "Description", "max": 3, "min": 3, "requires_restart": false, "type": "Type", "value": 5}}}`)
				}))
			})
			It(`Invoke GetConfiguration successfully with retries`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())
				hpdbService.EnableRetries(0, 0)

				// Construct an instance of the GetConfigurationOptions model
				getConfigurationOptionsModel := new(hpdbv3.GetConfigurationOptions)
				getConfigurationOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				getConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := hpdbService.GetConfigurationWithContext(ctx, getConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				hpdbService.DisableRetries()
				result, response, operationErr := hpdbService.GetConfiguration(getConfigurationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = hpdbService.GetConfigurationWithContext(ctx, getConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getConfigurationPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"configuration": {"deadlock_timeout": {"default": 7, "description": "Description", "max": 3, "min": 3, "requires_restart": false, "type": "Type", "value": 5}, "max_locks_per_transaction": {"default": 7, "description": "Description", "max": 3, "min": 3, "requires_restart": false, "type": "Type", "value": 5}, "shared_buffers": {"default": 7, "description": "Description", "max": 3, "min": 3, "requires_restart": false, "type": "Type", "value": 5}, "max_connections": {"default": 7, "description": "Description", "max": 3, "min": 3, "requires_restart": false, "type": "Type", "value": 5}}}`)
				}))
			})
			It(`Invoke GetConfiguration successfully`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := hpdbService.GetConfiguration(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetConfigurationOptions model
				getConfigurationOptionsModel := new(hpdbv3.GetConfigurationOptions)
				getConfigurationOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				getConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = hpdbService.GetConfiguration(getConfigurationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetConfiguration with error: Operation validation and request error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the GetConfigurationOptions model
				getConfigurationOptionsModel := new(hpdbv3.GetConfigurationOptions)
				getConfigurationOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				getConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := hpdbService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := hpdbService.GetConfiguration(getConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetConfigurationOptions model with no property values
				getConfigurationOptionsModelNew := new(hpdbv3.GetConfigurationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = hpdbService.GetConfiguration(getConfigurationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetConfiguration successfully`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the GetConfigurationOptions model
				getConfigurationOptionsModel := new(hpdbv3.GetConfigurationOptions)
				getConfigurationOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				getConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := hpdbService.GetConfiguration(getConfigurationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateConfiguration(updateConfigurationOptions *UpdateConfigurationOptions) - Operation response error`, func() {
		updateConfigurationPath := "/clusters/9cebab98-afeb-4886-9a29-8e741716e7ff/configuration"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateConfigurationPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateConfiguration with error: Operation response processing error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the Configurations model
				configurationsModel := new(hpdbv3.Configurations)
				configurationsModel.DeadlockTimeout = core.Int64Ptr(int64(10000))
				configurationsModel.MaxLocksPerTransaction = core.Int64Ptr(int64(100))
				configurationsModel.SharedBuffers = core.Int64Ptr(int64(256))
				configurationsModel.MaxConnections = core.Int64Ptr(int64(150))

				// Construct an instance of the UpdateConfigurationOptions model
				updateConfigurationOptionsModel := new(hpdbv3.UpdateConfigurationOptions)
				updateConfigurationOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				updateConfigurationOptionsModel.Configuration = configurationsModel
				updateConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := hpdbService.UpdateConfiguration(updateConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				hpdbService.EnableRetries(0, 0)
				result, response, operationErr = hpdbService.UpdateConfiguration(updateConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateConfiguration(updateConfigurationOptions *UpdateConfigurationOptions)`, func() {
		updateConfigurationPath := "/clusters/9cebab98-afeb-4886-9a29-8e741716e7ff/configuration"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateConfigurationPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"task_id": "1e902f30-da1b-11eb-9433-755fe141f81f"}`)
				}))
			})
			It(`Invoke UpdateConfiguration successfully with retries`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())
				hpdbService.EnableRetries(0, 0)

				// Construct an instance of the Configurations model
				configurationsModel := new(hpdbv3.Configurations)
				configurationsModel.DeadlockTimeout = core.Int64Ptr(int64(10000))
				configurationsModel.MaxLocksPerTransaction = core.Int64Ptr(int64(100))
				configurationsModel.SharedBuffers = core.Int64Ptr(int64(256))
				configurationsModel.MaxConnections = core.Int64Ptr(int64(150))

				// Construct an instance of the UpdateConfigurationOptions model
				updateConfigurationOptionsModel := new(hpdbv3.UpdateConfigurationOptions)
				updateConfigurationOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				updateConfigurationOptionsModel.Configuration = configurationsModel
				updateConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := hpdbService.UpdateConfigurationWithContext(ctx, updateConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				hpdbService.DisableRetries()
				result, response, operationErr := hpdbService.UpdateConfiguration(updateConfigurationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = hpdbService.UpdateConfigurationWithContext(ctx, updateConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateConfigurationPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"task_id": "1e902f30-da1b-11eb-9433-755fe141f81f"}`)
				}))
			})
			It(`Invoke UpdateConfiguration successfully`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := hpdbService.UpdateConfiguration(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Configurations model
				configurationsModel := new(hpdbv3.Configurations)
				configurationsModel.DeadlockTimeout = core.Int64Ptr(int64(10000))
				configurationsModel.MaxLocksPerTransaction = core.Int64Ptr(int64(100))
				configurationsModel.SharedBuffers = core.Int64Ptr(int64(256))
				configurationsModel.MaxConnections = core.Int64Ptr(int64(150))

				// Construct an instance of the UpdateConfigurationOptions model
				updateConfigurationOptionsModel := new(hpdbv3.UpdateConfigurationOptions)
				updateConfigurationOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				updateConfigurationOptionsModel.Configuration = configurationsModel
				updateConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = hpdbService.UpdateConfiguration(updateConfigurationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateConfiguration with error: Operation validation and request error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the Configurations model
				configurationsModel := new(hpdbv3.Configurations)
				configurationsModel.DeadlockTimeout = core.Int64Ptr(int64(10000))
				configurationsModel.MaxLocksPerTransaction = core.Int64Ptr(int64(100))
				configurationsModel.SharedBuffers = core.Int64Ptr(int64(256))
				configurationsModel.MaxConnections = core.Int64Ptr(int64(150))

				// Construct an instance of the UpdateConfigurationOptions model
				updateConfigurationOptionsModel := new(hpdbv3.UpdateConfigurationOptions)
				updateConfigurationOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				updateConfigurationOptionsModel.Configuration = configurationsModel
				updateConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := hpdbService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := hpdbService.UpdateConfiguration(updateConfigurationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateConfigurationOptions model with no property values
				updateConfigurationOptionsModelNew := new(hpdbv3.UpdateConfigurationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = hpdbService.UpdateConfiguration(updateConfigurationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(202)
				}))
			})
			It(`Invoke UpdateConfiguration successfully`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the Configurations model
				configurationsModel := new(hpdbv3.Configurations)
				configurationsModel.DeadlockTimeout = core.Int64Ptr(int64(10000))
				configurationsModel.MaxLocksPerTransaction = core.Int64Ptr(int64(100))
				configurationsModel.SharedBuffers = core.Int64Ptr(int64(256))
				configurationsModel.MaxConnections = core.Int64Ptr(int64(150))

				// Construct an instance of the UpdateConfigurationOptions model
				updateConfigurationOptionsModel := new(hpdbv3.UpdateConfigurationOptions)
				updateConfigurationOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				updateConfigurationOptionsModel.Configuration = configurationsModel
				updateConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := hpdbService.UpdateConfiguration(updateConfigurationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListTasks(listTasksOptions *ListTasksOptions) - Operation response error`, func() {
		listTasksPath := "/clusters/9cebab98-afeb-4886-9a29-8e741716e7ff/tasks"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTasksPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListTasks with error: Operation response processing error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the ListTasksOptions model
				listTasksOptionsModel := new(hpdbv3.ListTasksOptions)
				listTasksOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				listTasksOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := hpdbService.ListTasks(listTasksOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				hpdbService.EnableRetries(0, 0)
				result, response, operationErr = hpdbService.ListTasks(listTasksOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListTasks(listTasksOptions *ListTasksOptions)`, func() {
		listTasksPath := "/clusters/9cebab98-afeb-4886-9a29-8e741716e7ff/tasks"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTasksPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"tasks": [{"id": "1e902f30-da1b-11eb-9433-755fe141f81f", "type": "resource_scale", "state": "SUCCEEDED", "reason": "Reason", "started_at": "2021-07-01T03:19:17Z", "finished_at": "2021-07-01T03:21:13Z"}]}`)
				}))
			})
			It(`Invoke ListTasks successfully with retries`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())
				hpdbService.EnableRetries(0, 0)

				// Construct an instance of the ListTasksOptions model
				listTasksOptionsModel := new(hpdbv3.ListTasksOptions)
				listTasksOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				listTasksOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := hpdbService.ListTasksWithContext(ctx, listTasksOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				hpdbService.DisableRetries()
				result, response, operationErr := hpdbService.ListTasks(listTasksOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = hpdbService.ListTasksWithContext(ctx, listTasksOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTasksPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"tasks": [{"id": "1e902f30-da1b-11eb-9433-755fe141f81f", "type": "resource_scale", "state": "SUCCEEDED", "reason": "Reason", "started_at": "2021-07-01T03:19:17Z", "finished_at": "2021-07-01T03:21:13Z"}]}`)
				}))
			})
			It(`Invoke ListTasks successfully`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := hpdbService.ListTasks(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListTasksOptions model
				listTasksOptionsModel := new(hpdbv3.ListTasksOptions)
				listTasksOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				listTasksOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = hpdbService.ListTasks(listTasksOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListTasks with error: Operation validation and request error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the ListTasksOptions model
				listTasksOptionsModel := new(hpdbv3.ListTasksOptions)
				listTasksOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				listTasksOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := hpdbService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := hpdbService.ListTasks(listTasksOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListTasksOptions model with no property values
				listTasksOptionsModelNew := new(hpdbv3.ListTasksOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = hpdbService.ListTasks(listTasksOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListTasks successfully`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the ListTasksOptions model
				listTasksOptionsModel := new(hpdbv3.ListTasksOptions)
				listTasksOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				listTasksOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := hpdbService.ListTasks(listTasksOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTask(getTaskOptions *GetTaskOptions) - Operation response error`, func() {
		getTaskPath := "/clusters/9cebab98-afeb-4886-9a29-8e741716e7ff/tasks/c1a15760-a4f2-11ec-b00a-7f684d1dd53"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTaskPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetTask with error: Operation response processing error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the GetTaskOptions model
				getTaskOptionsModel := new(hpdbv3.GetTaskOptions)
				getTaskOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				getTaskOptionsModel.TaskID = core.StringPtr("c1a15760-a4f2-11ec-b00a-7f684d1dd53")
				getTaskOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := hpdbService.GetTask(getTaskOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				hpdbService.EnableRetries(0, 0)
				result, response, operationErr = hpdbService.GetTask(getTaskOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTask(getTaskOptions *GetTaskOptions)`, func() {
		getTaskPath := "/clusters/9cebab98-afeb-4886-9a29-8e741716e7ff/tasks/c1a15760-a4f2-11ec-b00a-7f684d1dd53"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTaskPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "1e902f30-da1b-11eb-9433-755fe141f81f", "type": "resource_scale", "started_at": "2021-07-01T03:19:17Z", "finished_at": "2021-07-01T03:21:13Z", "reason": "Reason", "state": "SUCCEEDED", "nodes": [{"id": "c5ff2d841c7e6a11de3cbaa2b992d712", "state": "SUCCEEDED", "reason": "Reason", "started_at": "2021-07-01T03:20:36Z", "finished_at": "2021-07-01T03:20:52Z"}], "spec": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke GetTask successfully with retries`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())
				hpdbService.EnableRetries(0, 0)

				// Construct an instance of the GetTaskOptions model
				getTaskOptionsModel := new(hpdbv3.GetTaskOptions)
				getTaskOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				getTaskOptionsModel.TaskID = core.StringPtr("c1a15760-a4f2-11ec-b00a-7f684d1dd53")
				getTaskOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := hpdbService.GetTaskWithContext(ctx, getTaskOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				hpdbService.DisableRetries()
				result, response, operationErr := hpdbService.GetTask(getTaskOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = hpdbService.GetTaskWithContext(ctx, getTaskOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTaskPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "1e902f30-da1b-11eb-9433-755fe141f81f", "type": "resource_scale", "started_at": "2021-07-01T03:19:17Z", "finished_at": "2021-07-01T03:21:13Z", "reason": "Reason", "state": "SUCCEEDED", "nodes": [{"id": "c5ff2d841c7e6a11de3cbaa2b992d712", "state": "SUCCEEDED", "reason": "Reason", "started_at": "2021-07-01T03:20:36Z", "finished_at": "2021-07-01T03:20:52Z"}], "spec": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke GetTask successfully`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := hpdbService.GetTask(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTaskOptions model
				getTaskOptionsModel := new(hpdbv3.GetTaskOptions)
				getTaskOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				getTaskOptionsModel.TaskID = core.StringPtr("c1a15760-a4f2-11ec-b00a-7f684d1dd53")
				getTaskOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = hpdbService.GetTask(getTaskOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetTask with error: Operation validation and request error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the GetTaskOptions model
				getTaskOptionsModel := new(hpdbv3.GetTaskOptions)
				getTaskOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				getTaskOptionsModel.TaskID = core.StringPtr("c1a15760-a4f2-11ec-b00a-7f684d1dd53")
				getTaskOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := hpdbService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := hpdbService.GetTask(getTaskOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetTaskOptions model with no property values
				getTaskOptionsModelNew := new(hpdbv3.GetTaskOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = hpdbService.GetTask(getTaskOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetTask successfully`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the GetTaskOptions model
				getTaskOptionsModel := new(hpdbv3.GetTaskOptions)
				getTaskOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				getTaskOptionsModel.TaskID = core.StringPtr("c1a15760-a4f2-11ec-b00a-7f684d1dd53")
				getTaskOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := hpdbService.GetTask(getTaskOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListBackups(listBackupsOptions *ListBackupsOptions) - Operation response error`, func() {
		listBackupsPath := "/clusters/9cebab98-afeb-4886-9a29-8e741716e7ff/backups"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listBackupsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListBackups with error: Operation response processing error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the ListBackupsOptions model
				listBackupsOptionsModel := new(hpdbv3.ListBackupsOptions)
				listBackupsOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				listBackupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := hpdbService.ListBackups(listBackupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				hpdbService.EnableRetries(0, 0)
				result, response, operationErr = hpdbService.ListBackups(listBackupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListBackups(listBackupsOptions *ListBackupsOptions)`, func() {
		listBackupsPath := "/clusters/9cebab98-afeb-4886-9a29-8e741716e7ff/backups"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listBackupsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"backups": [{"id": "dbaas23-20022-20220410-140000F", "type": "scheduled", "created_at": "2022-04-11T14:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke ListBackups successfully with retries`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())
				hpdbService.EnableRetries(0, 0)

				// Construct an instance of the ListBackupsOptions model
				listBackupsOptionsModel := new(hpdbv3.ListBackupsOptions)
				listBackupsOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				listBackupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := hpdbService.ListBackupsWithContext(ctx, listBackupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				hpdbService.DisableRetries()
				result, response, operationErr := hpdbService.ListBackups(listBackupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = hpdbService.ListBackupsWithContext(ctx, listBackupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listBackupsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"backups": [{"id": "dbaas23-20022-20220410-140000F", "type": "scheduled", "created_at": "2022-04-11T14:00:00.000Z"}]}`)
				}))
			})
			It(`Invoke ListBackups successfully`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := hpdbService.ListBackups(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListBackupsOptions model
				listBackupsOptionsModel := new(hpdbv3.ListBackupsOptions)
				listBackupsOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				listBackupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = hpdbService.ListBackups(listBackupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListBackups with error: Operation validation and request error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the ListBackupsOptions model
				listBackupsOptionsModel := new(hpdbv3.ListBackupsOptions)
				listBackupsOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				listBackupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := hpdbService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := hpdbService.ListBackups(listBackupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListBackupsOptions model with no property values
				listBackupsOptionsModelNew := new(hpdbv3.ListBackupsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = hpdbService.ListBackups(listBackupsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListBackups successfully`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the ListBackupsOptions model
				listBackupsOptionsModel := new(hpdbv3.ListBackupsOptions)
				listBackupsOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				listBackupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := hpdbService.ListBackups(listBackupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`EnableCosBackup(enableCosBackupOptions *EnableCosBackupOptions) - Operation response error`, func() {
		enableCosBackupPath := "/clusters/9cebab98-afeb-4886-9a29-8e741716e7ff/backups/cos/enable"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(enableCosBackupPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke EnableCosBackup with error: Operation response processing error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the CosHmacKeys model
				cosHmacKeysModel := new(hpdbv3.CosHmacKeys)
				cosHmacKeysModel.AccessKeyID = core.StringPtr("e4465d50a56f401a81d275f55c57bc2f")
				cosHmacKeysModel.SecretAccessKey = core.StringPtr("0c29c8299dbba1d6xx7191d6xx5ce3e7eb601fa4bd9f5a")

				// Construct an instance of the EnableCosBackupOptions model
				enableCosBackupOptionsModel := new(hpdbv3.EnableCosBackupOptions)
				enableCosBackupOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				enableCosBackupOptionsModel.CosHmacKeys = cosHmacKeysModel
				enableCosBackupOptionsModel.CosEndpoint = core.StringPtr("s3.us-west.cloud-object-storage.test.appdomain.cloud")
				enableCosBackupOptionsModel.BucketInstanceCrn = core.StringPtr("crn:v1:staging:public:cloud-object-storage:global:a/23a24a3e3fe7a115473f07be1c44bdb5:3ac008bb-16a8-48b2-b8f7-25a990c14b41:bucket:cloud-object-storage-gc-cos-standard-luu")
				enableCosBackupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := hpdbService.EnableCosBackup(enableCosBackupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				hpdbService.EnableRetries(0, 0)
				result, response, operationErr = hpdbService.EnableCosBackup(enableCosBackupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`EnableCosBackup(enableCosBackupOptions *EnableCosBackupOptions)`, func() {
		enableCosBackupPath := "/clusters/9cebab98-afeb-4886-9a29-8e741716e7ff/backups/cos/enable"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(enableCosBackupPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"task_id": "1e902f30-da1b-11eb-9433-755fe141f81f"}`)
				}))
			})
			It(`Invoke EnableCosBackup successfully with retries`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())
				hpdbService.EnableRetries(0, 0)

				// Construct an instance of the CosHmacKeys model
				cosHmacKeysModel := new(hpdbv3.CosHmacKeys)
				cosHmacKeysModel.AccessKeyID = core.StringPtr("e4465d50a56f401a81d275f55c57bc2f")
				cosHmacKeysModel.SecretAccessKey = core.StringPtr("0c29c8299dbba1d6xx7191d6xx5ce3e7eb601fa4bd9f5a")

				// Construct an instance of the EnableCosBackupOptions model
				enableCosBackupOptionsModel := new(hpdbv3.EnableCosBackupOptions)
				enableCosBackupOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				enableCosBackupOptionsModel.CosHmacKeys = cosHmacKeysModel
				enableCosBackupOptionsModel.CosEndpoint = core.StringPtr("s3.us-west.cloud-object-storage.test.appdomain.cloud")
				enableCosBackupOptionsModel.BucketInstanceCrn = core.StringPtr("crn:v1:staging:public:cloud-object-storage:global:a/23a24a3e3fe7a115473f07be1c44bdb5:3ac008bb-16a8-48b2-b8f7-25a990c14b41:bucket:cloud-object-storage-gc-cos-standard-luu")
				enableCosBackupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := hpdbService.EnableCosBackupWithContext(ctx, enableCosBackupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				hpdbService.DisableRetries()
				result, response, operationErr := hpdbService.EnableCosBackup(enableCosBackupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = hpdbService.EnableCosBackupWithContext(ctx, enableCosBackupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(enableCosBackupPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"task_id": "1e902f30-da1b-11eb-9433-755fe141f81f"}`)
				}))
			})
			It(`Invoke EnableCosBackup successfully`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := hpdbService.EnableCosBackup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CosHmacKeys model
				cosHmacKeysModel := new(hpdbv3.CosHmacKeys)
				cosHmacKeysModel.AccessKeyID = core.StringPtr("e4465d50a56f401a81d275f55c57bc2f")
				cosHmacKeysModel.SecretAccessKey = core.StringPtr("0c29c8299dbba1d6xx7191d6xx5ce3e7eb601fa4bd9f5a")

				// Construct an instance of the EnableCosBackupOptions model
				enableCosBackupOptionsModel := new(hpdbv3.EnableCosBackupOptions)
				enableCosBackupOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				enableCosBackupOptionsModel.CosHmacKeys = cosHmacKeysModel
				enableCosBackupOptionsModel.CosEndpoint = core.StringPtr("s3.us-west.cloud-object-storage.test.appdomain.cloud")
				enableCosBackupOptionsModel.BucketInstanceCrn = core.StringPtr("crn:v1:staging:public:cloud-object-storage:global:a/23a24a3e3fe7a115473f07be1c44bdb5:3ac008bb-16a8-48b2-b8f7-25a990c14b41:bucket:cloud-object-storage-gc-cos-standard-luu")
				enableCosBackupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = hpdbService.EnableCosBackup(enableCosBackupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke EnableCosBackup with error: Operation validation and request error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the CosHmacKeys model
				cosHmacKeysModel := new(hpdbv3.CosHmacKeys)
				cosHmacKeysModel.AccessKeyID = core.StringPtr("e4465d50a56f401a81d275f55c57bc2f")
				cosHmacKeysModel.SecretAccessKey = core.StringPtr("0c29c8299dbba1d6xx7191d6xx5ce3e7eb601fa4bd9f5a")

				// Construct an instance of the EnableCosBackupOptions model
				enableCosBackupOptionsModel := new(hpdbv3.EnableCosBackupOptions)
				enableCosBackupOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				enableCosBackupOptionsModel.CosHmacKeys = cosHmacKeysModel
				enableCosBackupOptionsModel.CosEndpoint = core.StringPtr("s3.us-west.cloud-object-storage.test.appdomain.cloud")
				enableCosBackupOptionsModel.BucketInstanceCrn = core.StringPtr("crn:v1:staging:public:cloud-object-storage:global:a/23a24a3e3fe7a115473f07be1c44bdb5:3ac008bb-16a8-48b2-b8f7-25a990c14b41:bucket:cloud-object-storage-gc-cos-standard-luu")
				enableCosBackupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := hpdbService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := hpdbService.EnableCosBackup(enableCosBackupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the EnableCosBackupOptions model with no property values
				enableCosBackupOptionsModelNew := new(hpdbv3.EnableCosBackupOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = hpdbService.EnableCosBackup(enableCosBackupOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(202)
				}))
			})
			It(`Invoke EnableCosBackup successfully`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the CosHmacKeys model
				cosHmacKeysModel := new(hpdbv3.CosHmacKeys)
				cosHmacKeysModel.AccessKeyID = core.StringPtr("e4465d50a56f401a81d275f55c57bc2f")
				cosHmacKeysModel.SecretAccessKey = core.StringPtr("0c29c8299dbba1d6xx7191d6xx5ce3e7eb601fa4bd9f5a")

				// Construct an instance of the EnableCosBackupOptions model
				enableCosBackupOptionsModel := new(hpdbv3.EnableCosBackupOptions)
				enableCosBackupOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				enableCosBackupOptionsModel.CosHmacKeys = cosHmacKeysModel
				enableCosBackupOptionsModel.CosEndpoint = core.StringPtr("s3.us-west.cloud-object-storage.test.appdomain.cloud")
				enableCosBackupOptionsModel.BucketInstanceCrn = core.StringPtr("crn:v1:staging:public:cloud-object-storage:global:a/23a24a3e3fe7a115473f07be1c44bdb5:3ac008bb-16a8-48b2-b8f7-25a990c14b41:bucket:cloud-object-storage-gc-cos-standard-luu")
				enableCosBackupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := hpdbService.EnableCosBackup(enableCosBackupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DisableCosBackup(disableCosBackupOptions *DisableCosBackupOptions) - Operation response error`, func() {
		disableCosBackupPath := "/clusters/9cebab98-afeb-4886-9a29-8e741716e7ff/backups/cos/disable"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(disableCosBackupPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DisableCosBackup with error: Operation response processing error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the DisableCosBackupOptions model
				disableCosBackupOptionsModel := new(hpdbv3.DisableCosBackupOptions)
				disableCosBackupOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				disableCosBackupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := hpdbService.DisableCosBackup(disableCosBackupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				hpdbService.EnableRetries(0, 0)
				result, response, operationErr = hpdbService.DisableCosBackup(disableCosBackupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DisableCosBackup(disableCosBackupOptions *DisableCosBackupOptions)`, func() {
		disableCosBackupPath := "/clusters/9cebab98-afeb-4886-9a29-8e741716e7ff/backups/cos/disable"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(disableCosBackupPath))
					Expect(req.Method).To(Equal("POST"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"task_id": "1e902f30-da1b-11eb-9433-755fe141f81f"}`)
				}))
			})
			It(`Invoke DisableCosBackup successfully with retries`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())
				hpdbService.EnableRetries(0, 0)

				// Construct an instance of the DisableCosBackupOptions model
				disableCosBackupOptionsModel := new(hpdbv3.DisableCosBackupOptions)
				disableCosBackupOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				disableCosBackupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := hpdbService.DisableCosBackupWithContext(ctx, disableCosBackupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				hpdbService.DisableRetries()
				result, response, operationErr := hpdbService.DisableCosBackup(disableCosBackupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = hpdbService.DisableCosBackupWithContext(ctx, disableCosBackupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(disableCosBackupPath))
					Expect(req.Method).To(Equal("POST"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"task_id": "1e902f30-da1b-11eb-9433-755fe141f81f"}`)
				}))
			})
			It(`Invoke DisableCosBackup successfully`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := hpdbService.DisableCosBackup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DisableCosBackupOptions model
				disableCosBackupOptionsModel := new(hpdbv3.DisableCosBackupOptions)
				disableCosBackupOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				disableCosBackupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = hpdbService.DisableCosBackup(disableCosBackupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DisableCosBackup with error: Operation validation and request error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the DisableCosBackupOptions model
				disableCosBackupOptionsModel := new(hpdbv3.DisableCosBackupOptions)
				disableCosBackupOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				disableCosBackupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := hpdbService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := hpdbService.DisableCosBackup(disableCosBackupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DisableCosBackupOptions model with no property values
				disableCosBackupOptionsModelNew := new(hpdbv3.DisableCosBackupOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = hpdbService.DisableCosBackup(disableCosBackupOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(202)
				}))
			})
			It(`Invoke DisableCosBackup successfully`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the DisableCosBackupOptions model
				disableCosBackupOptionsModel := new(hpdbv3.DisableCosBackupOptions)
				disableCosBackupOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				disableCosBackupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := hpdbService.DisableCosBackup(disableCosBackupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCosBackupConfig(getCosBackupConfigOptions *GetCosBackupConfigOptions) - Operation response error`, func() {
		getCosBackupConfigPath := "/clusters/9cebab98-afeb-4886-9a29-8e741716e7ff/backups/cos/configuration"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCosBackupConfigPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCosBackupConfig with error: Operation response processing error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the GetCosBackupConfigOptions model
				getCosBackupConfigOptionsModel := new(hpdbv3.GetCosBackupConfigOptions)
				getCosBackupConfigOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				getCosBackupConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := hpdbService.GetCosBackupConfig(getCosBackupConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				hpdbService.EnableRetries(0, 0)
				result, response, operationErr = hpdbService.GetCosBackupConfig(getCosBackupConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCosBackupConfig(getCosBackupConfigOptions *GetCosBackupConfigOptions)`, func() {
		getCosBackupConfigPath := "/clusters/9cebab98-afeb-4886-9a29-8e741716e7ff/backups/cos/configuration"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCosBackupConfigPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"cos_endpoint": "CosEndpoint", "bucket_instance_crn": "BucketInstanceCrn"}`)
				}))
			})
			It(`Invoke GetCosBackupConfig successfully with retries`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())
				hpdbService.EnableRetries(0, 0)

				// Construct an instance of the GetCosBackupConfigOptions model
				getCosBackupConfigOptionsModel := new(hpdbv3.GetCosBackupConfigOptions)
				getCosBackupConfigOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				getCosBackupConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := hpdbService.GetCosBackupConfigWithContext(ctx, getCosBackupConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				hpdbService.DisableRetries()
				result, response, operationErr := hpdbService.GetCosBackupConfig(getCosBackupConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = hpdbService.GetCosBackupConfigWithContext(ctx, getCosBackupConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCosBackupConfigPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"cos_endpoint": "CosEndpoint", "bucket_instance_crn": "BucketInstanceCrn"}`)
				}))
			})
			It(`Invoke GetCosBackupConfig successfully`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := hpdbService.GetCosBackupConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCosBackupConfigOptions model
				getCosBackupConfigOptionsModel := new(hpdbv3.GetCosBackupConfigOptions)
				getCosBackupConfigOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				getCosBackupConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = hpdbService.GetCosBackupConfig(getCosBackupConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetCosBackupConfig with error: Operation validation and request error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the GetCosBackupConfigOptions model
				getCosBackupConfigOptionsModel := new(hpdbv3.GetCosBackupConfigOptions)
				getCosBackupConfigOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				getCosBackupConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := hpdbService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := hpdbService.GetCosBackupConfig(getCosBackupConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetCosBackupConfigOptions model with no property values
				getCosBackupConfigOptionsModelNew := new(hpdbv3.GetCosBackupConfigOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = hpdbService.GetCosBackupConfig(getCosBackupConfigOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetCosBackupConfig successfully`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the GetCosBackupConfigOptions model
				getCosBackupConfigOptionsModel := new(hpdbv3.GetCosBackupConfigOptions)
				getCosBackupConfigOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				getCosBackupConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := hpdbService.GetCosBackupConfig(getCosBackupConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetBackupConfig(getBackupConfigOptions *GetBackupConfigOptions) - Operation response error`, func() {
		getBackupConfigPath := "/clusters/9cebab98-afeb-4886-9a29-8e741716e7ff/backups/configuration"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBackupConfigPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetBackupConfig with error: Operation response processing error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the GetBackupConfigOptions model
				getBackupConfigOptionsModel := new(hpdbv3.GetBackupConfigOptions)
				getBackupConfigOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				getBackupConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := hpdbService.GetBackupConfig(getBackupConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				hpdbService.EnableRetries(0, 0)
				result, response, operationErr = hpdbService.GetBackupConfig(getBackupConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetBackupConfig(getBackupConfigOptions *GetBackupConfigOptions)`, func() {
		getBackupConfigPath := "/clusters/9cebab98-afeb-4886-9a29-8e741716e7ff/backups/configuration"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBackupConfigPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"cos": {"cos_endpoint": "CosEndpoint", "bucket_instance_crn": "BucketInstanceCrn", "schedule": {"type": "frequency", "interval": "PT8H"}}, "status": {"state": "State", "created_at": "CreatedAt", "updated_at": "2022-04-11T14:00:00.000Z"}}`)
				}))
			})
			It(`Invoke GetBackupConfig successfully with retries`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())
				hpdbService.EnableRetries(0, 0)

				// Construct an instance of the GetBackupConfigOptions model
				getBackupConfigOptionsModel := new(hpdbv3.GetBackupConfigOptions)
				getBackupConfigOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				getBackupConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := hpdbService.GetBackupConfigWithContext(ctx, getBackupConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				hpdbService.DisableRetries()
				result, response, operationErr := hpdbService.GetBackupConfig(getBackupConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = hpdbService.GetBackupConfigWithContext(ctx, getBackupConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBackupConfigPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"cos": {"cos_endpoint": "CosEndpoint", "bucket_instance_crn": "BucketInstanceCrn", "schedule": {"type": "frequency", "interval": "PT8H"}}, "status": {"state": "State", "created_at": "CreatedAt", "updated_at": "2022-04-11T14:00:00.000Z"}}`)
				}))
			})
			It(`Invoke GetBackupConfig successfully`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := hpdbService.GetBackupConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetBackupConfigOptions model
				getBackupConfigOptionsModel := new(hpdbv3.GetBackupConfigOptions)
				getBackupConfigOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				getBackupConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = hpdbService.GetBackupConfig(getBackupConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetBackupConfig with error: Operation validation and request error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the GetBackupConfigOptions model
				getBackupConfigOptionsModel := new(hpdbv3.GetBackupConfigOptions)
				getBackupConfigOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				getBackupConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := hpdbService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := hpdbService.GetBackupConfig(getBackupConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetBackupConfigOptions model with no property values
				getBackupConfigOptionsModelNew := new(hpdbv3.GetBackupConfigOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = hpdbService.GetBackupConfig(getBackupConfigOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetBackupConfig successfully`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the GetBackupConfigOptions model
				getBackupConfigOptionsModel := new(hpdbv3.GetBackupConfigOptions)
				getBackupConfigOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				getBackupConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := hpdbService.GetBackupConfig(getBackupConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateBackupConfig(updateBackupConfigOptions *UpdateBackupConfigOptions) - Operation response error`, func() {
		updateBackupConfigPath := "/clusters/9cebab98-afeb-4886-9a29-8e741716e7ff/backups/configuration"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateBackupConfigPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateBackupConfig with error: Operation response processing error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the CosHmacKeys model
				cosHmacKeysModel := new(hpdbv3.CosHmacKeys)
				cosHmacKeysModel.AccessKeyID = core.StringPtr("e4465d50a56f401a81d275f55c57bc2f")
				cosHmacKeysModel.SecretAccessKey = core.StringPtr("0c29c8299dbba1d6xx7191d6xx5ce3e7eb601fa4bd9f5a")

				// Construct an instance of the BackupSchedule model
				backupScheduleModel := new(hpdbv3.BackupSchedule)
				backupScheduleModel.Type = core.StringPtr("frequency")
				backupScheduleModel.Interval = core.StringPtr("PT8H")

				// Construct an instance of the BackupConfigCos model
				backupConfigCosModel := new(hpdbv3.BackupConfigCos)
				backupConfigCosModel.CosHmacKeys = cosHmacKeysModel
				backupConfigCosModel.CosEndpoint = core.StringPtr("s3.us-west.cloud-object-storage.test.appdomain.cloud")
				backupConfigCosModel.BucketInstanceCrn = core.StringPtr("crn:v1:staging:public:cloud-object-storage:global:a/23a24a3e3fe7a115473f07be1c44bdb5:3ac008bb-16a8-48b2-b8f7-25a990c14b41:bucket:cloud-object-storage-gc-cos-standard-luu")
				backupConfigCosModel.Schedule = backupScheduleModel

				// Construct an instance of the UpdateBackupConfigOptions model
				updateBackupConfigOptionsModel := new(hpdbv3.UpdateBackupConfigOptions)
				updateBackupConfigOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				updateBackupConfigOptionsModel.Cos = backupConfigCosModel
				updateBackupConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := hpdbService.UpdateBackupConfig(updateBackupConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				hpdbService.EnableRetries(0, 0)
				result, response, operationErr = hpdbService.UpdateBackupConfig(updateBackupConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateBackupConfig(updateBackupConfigOptions *UpdateBackupConfigOptions)`, func() {
		updateBackupConfigPath := "/clusters/9cebab98-afeb-4886-9a29-8e741716e7ff/backups/configuration"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateBackupConfigPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"task_id": "1e902f30-da1b-11eb-9433-755fe141f81f"}`)
				}))
			})
			It(`Invoke UpdateBackupConfig successfully with retries`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())
				hpdbService.EnableRetries(0, 0)

				// Construct an instance of the CosHmacKeys model
				cosHmacKeysModel := new(hpdbv3.CosHmacKeys)
				cosHmacKeysModel.AccessKeyID = core.StringPtr("e4465d50a56f401a81d275f55c57bc2f")
				cosHmacKeysModel.SecretAccessKey = core.StringPtr("0c29c8299dbba1d6xx7191d6xx5ce3e7eb601fa4bd9f5a")

				// Construct an instance of the BackupSchedule model
				backupScheduleModel := new(hpdbv3.BackupSchedule)
				backupScheduleModel.Type = core.StringPtr("frequency")
				backupScheduleModel.Interval = core.StringPtr("PT8H")

				// Construct an instance of the BackupConfigCos model
				backupConfigCosModel := new(hpdbv3.BackupConfigCos)
				backupConfigCosModel.CosHmacKeys = cosHmacKeysModel
				backupConfigCosModel.CosEndpoint = core.StringPtr("s3.us-west.cloud-object-storage.test.appdomain.cloud")
				backupConfigCosModel.BucketInstanceCrn = core.StringPtr("crn:v1:staging:public:cloud-object-storage:global:a/23a24a3e3fe7a115473f07be1c44bdb5:3ac008bb-16a8-48b2-b8f7-25a990c14b41:bucket:cloud-object-storage-gc-cos-standard-luu")
				backupConfigCosModel.Schedule = backupScheduleModel

				// Construct an instance of the UpdateBackupConfigOptions model
				updateBackupConfigOptionsModel := new(hpdbv3.UpdateBackupConfigOptions)
				updateBackupConfigOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				updateBackupConfigOptionsModel.Cos = backupConfigCosModel
				updateBackupConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := hpdbService.UpdateBackupConfigWithContext(ctx, updateBackupConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				hpdbService.DisableRetries()
				result, response, operationErr := hpdbService.UpdateBackupConfig(updateBackupConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = hpdbService.UpdateBackupConfigWithContext(ctx, updateBackupConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateBackupConfigPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"task_id": "1e902f30-da1b-11eb-9433-755fe141f81f"}`)
				}))
			})
			It(`Invoke UpdateBackupConfig successfully`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := hpdbService.UpdateBackupConfig(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CosHmacKeys model
				cosHmacKeysModel := new(hpdbv3.CosHmacKeys)
				cosHmacKeysModel.AccessKeyID = core.StringPtr("e4465d50a56f401a81d275f55c57bc2f")
				cosHmacKeysModel.SecretAccessKey = core.StringPtr("0c29c8299dbba1d6xx7191d6xx5ce3e7eb601fa4bd9f5a")

				// Construct an instance of the BackupSchedule model
				backupScheduleModel := new(hpdbv3.BackupSchedule)
				backupScheduleModel.Type = core.StringPtr("frequency")
				backupScheduleModel.Interval = core.StringPtr("PT8H")

				// Construct an instance of the BackupConfigCos model
				backupConfigCosModel := new(hpdbv3.BackupConfigCos)
				backupConfigCosModel.CosHmacKeys = cosHmacKeysModel
				backupConfigCosModel.CosEndpoint = core.StringPtr("s3.us-west.cloud-object-storage.test.appdomain.cloud")
				backupConfigCosModel.BucketInstanceCrn = core.StringPtr("crn:v1:staging:public:cloud-object-storage:global:a/23a24a3e3fe7a115473f07be1c44bdb5:3ac008bb-16a8-48b2-b8f7-25a990c14b41:bucket:cloud-object-storage-gc-cos-standard-luu")
				backupConfigCosModel.Schedule = backupScheduleModel

				// Construct an instance of the UpdateBackupConfigOptions model
				updateBackupConfigOptionsModel := new(hpdbv3.UpdateBackupConfigOptions)
				updateBackupConfigOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				updateBackupConfigOptionsModel.Cos = backupConfigCosModel
				updateBackupConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = hpdbService.UpdateBackupConfig(updateBackupConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateBackupConfig with error: Operation validation and request error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the CosHmacKeys model
				cosHmacKeysModel := new(hpdbv3.CosHmacKeys)
				cosHmacKeysModel.AccessKeyID = core.StringPtr("e4465d50a56f401a81d275f55c57bc2f")
				cosHmacKeysModel.SecretAccessKey = core.StringPtr("0c29c8299dbba1d6xx7191d6xx5ce3e7eb601fa4bd9f5a")

				// Construct an instance of the BackupSchedule model
				backupScheduleModel := new(hpdbv3.BackupSchedule)
				backupScheduleModel.Type = core.StringPtr("frequency")
				backupScheduleModel.Interval = core.StringPtr("PT8H")

				// Construct an instance of the BackupConfigCos model
				backupConfigCosModel := new(hpdbv3.BackupConfigCos)
				backupConfigCosModel.CosHmacKeys = cosHmacKeysModel
				backupConfigCosModel.CosEndpoint = core.StringPtr("s3.us-west.cloud-object-storage.test.appdomain.cloud")
				backupConfigCosModel.BucketInstanceCrn = core.StringPtr("crn:v1:staging:public:cloud-object-storage:global:a/23a24a3e3fe7a115473f07be1c44bdb5:3ac008bb-16a8-48b2-b8f7-25a990c14b41:bucket:cloud-object-storage-gc-cos-standard-luu")
				backupConfigCosModel.Schedule = backupScheduleModel

				// Construct an instance of the UpdateBackupConfigOptions model
				updateBackupConfigOptionsModel := new(hpdbv3.UpdateBackupConfigOptions)
				updateBackupConfigOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				updateBackupConfigOptionsModel.Cos = backupConfigCosModel
				updateBackupConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := hpdbService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := hpdbService.UpdateBackupConfig(updateBackupConfigOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateBackupConfigOptions model with no property values
				updateBackupConfigOptionsModelNew := new(hpdbv3.UpdateBackupConfigOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = hpdbService.UpdateBackupConfig(updateBackupConfigOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(202)
				}))
			})
			It(`Invoke UpdateBackupConfig successfully`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the CosHmacKeys model
				cosHmacKeysModel := new(hpdbv3.CosHmacKeys)
				cosHmacKeysModel.AccessKeyID = core.StringPtr("e4465d50a56f401a81d275f55c57bc2f")
				cosHmacKeysModel.SecretAccessKey = core.StringPtr("0c29c8299dbba1d6xx7191d6xx5ce3e7eb601fa4bd9f5a")

				// Construct an instance of the BackupSchedule model
				backupScheduleModel := new(hpdbv3.BackupSchedule)
				backupScheduleModel.Type = core.StringPtr("frequency")
				backupScheduleModel.Interval = core.StringPtr("PT8H")

				// Construct an instance of the BackupConfigCos model
				backupConfigCosModel := new(hpdbv3.BackupConfigCos)
				backupConfigCosModel.CosHmacKeys = cosHmacKeysModel
				backupConfigCosModel.CosEndpoint = core.StringPtr("s3.us-west.cloud-object-storage.test.appdomain.cloud")
				backupConfigCosModel.BucketInstanceCrn = core.StringPtr("crn:v1:staging:public:cloud-object-storage:global:a/23a24a3e3fe7a115473f07be1c44bdb5:3ac008bb-16a8-48b2-b8f7-25a990c14b41:bucket:cloud-object-storage-gc-cos-standard-luu")
				backupConfigCosModel.Schedule = backupScheduleModel

				// Construct an instance of the UpdateBackupConfigOptions model
				updateBackupConfigOptionsModel := new(hpdbv3.UpdateBackupConfigOptions)
				updateBackupConfigOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				updateBackupConfigOptionsModel.Cos = backupConfigCosModel
				updateBackupConfigOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := hpdbService.UpdateBackupConfig(updateBackupConfigOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Restore(restoreOptions *RestoreOptions) - Operation response error`, func() {
		restorePath := "/clusters/9cebab98-afeb-4886-9a29-8e741716e7ff/restore"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(restorePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke Restore with error: Operation response processing error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the CosHmacKeys model
				cosHmacKeysModel := new(hpdbv3.CosHmacKeys)
				cosHmacKeysModel.AccessKeyID = core.StringPtr("e4465d50a56f401a81d275f55c57bc2f")
				cosHmacKeysModel.SecretAccessKey = core.StringPtr("0c29c8299dbba1d6xx7191d6xx5ce3e7eb601fa4bd9f5a")

				// Construct an instance of the RestoreOptions model
				restoreOptionsModel := new(hpdbv3.RestoreOptions)
				restoreOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				restoreOptionsModel.SourceType = core.StringPtr("cos")
				restoreOptionsModel.CosHmacKeys = cosHmacKeysModel
				restoreOptionsModel.CosEndpoint = core.StringPtr("s3.us-west.cloud-object-storage.test.appdomain.cloud")
				restoreOptionsModel.BucketInstanceCrn = core.StringPtr("crn:v1:staging:public:cloud-object-storage:global:a/23a24a3e3fe7a115473f07be1c44bdb5:3ac008bb-16a8-48b2-b8f7-25a990c14b41:bucket:cloud-object-storage-gc-cos-standard-luu")
				restoreOptionsModel.BackupFile = core.StringPtr("archive-2022-03-16-140004Z.tar")
				restoreOptionsModel.BackupID = core.StringPtr("dbaas23-20022-20220410-140000F")
				restoreOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := hpdbService.Restore(restoreOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				hpdbService.EnableRetries(0, 0)
				result, response, operationErr = hpdbService.Restore(restoreOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Restore(restoreOptions *RestoreOptions)`, func() {
		restorePath := "/clusters/9cebab98-afeb-4886-9a29-8e741716e7ff/restore"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(restorePath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"task_id": "1e902f30-da1b-11eb-9433-755fe141f81f"}`)
				}))
			})
			It(`Invoke Restore successfully with retries`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())
				hpdbService.EnableRetries(0, 0)

				// Construct an instance of the CosHmacKeys model
				cosHmacKeysModel := new(hpdbv3.CosHmacKeys)
				cosHmacKeysModel.AccessKeyID = core.StringPtr("e4465d50a56f401a81d275f55c57bc2f")
				cosHmacKeysModel.SecretAccessKey = core.StringPtr("0c29c8299dbba1d6xx7191d6xx5ce3e7eb601fa4bd9f5a")

				// Construct an instance of the RestoreOptions model
				restoreOptionsModel := new(hpdbv3.RestoreOptions)
				restoreOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				restoreOptionsModel.SourceType = core.StringPtr("cos")
				restoreOptionsModel.CosHmacKeys = cosHmacKeysModel
				restoreOptionsModel.CosEndpoint = core.StringPtr("s3.us-west.cloud-object-storage.test.appdomain.cloud")
				restoreOptionsModel.BucketInstanceCrn = core.StringPtr("crn:v1:staging:public:cloud-object-storage:global:a/23a24a3e3fe7a115473f07be1c44bdb5:3ac008bb-16a8-48b2-b8f7-25a990c14b41:bucket:cloud-object-storage-gc-cos-standard-luu")
				restoreOptionsModel.BackupFile = core.StringPtr("archive-2022-03-16-140004Z.tar")
				restoreOptionsModel.BackupID = core.StringPtr("dbaas23-20022-20220410-140000F")
				restoreOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := hpdbService.RestoreWithContext(ctx, restoreOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				hpdbService.DisableRetries()
				result, response, operationErr := hpdbService.Restore(restoreOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = hpdbService.RestoreWithContext(ctx, restoreOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(restorePath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"task_id": "1e902f30-da1b-11eb-9433-755fe141f81f"}`)
				}))
			})
			It(`Invoke Restore successfully`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := hpdbService.Restore(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CosHmacKeys model
				cosHmacKeysModel := new(hpdbv3.CosHmacKeys)
				cosHmacKeysModel.AccessKeyID = core.StringPtr("e4465d50a56f401a81d275f55c57bc2f")
				cosHmacKeysModel.SecretAccessKey = core.StringPtr("0c29c8299dbba1d6xx7191d6xx5ce3e7eb601fa4bd9f5a")

				// Construct an instance of the RestoreOptions model
				restoreOptionsModel := new(hpdbv3.RestoreOptions)
				restoreOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				restoreOptionsModel.SourceType = core.StringPtr("cos")
				restoreOptionsModel.CosHmacKeys = cosHmacKeysModel
				restoreOptionsModel.CosEndpoint = core.StringPtr("s3.us-west.cloud-object-storage.test.appdomain.cloud")
				restoreOptionsModel.BucketInstanceCrn = core.StringPtr("crn:v1:staging:public:cloud-object-storage:global:a/23a24a3e3fe7a115473f07be1c44bdb5:3ac008bb-16a8-48b2-b8f7-25a990c14b41:bucket:cloud-object-storage-gc-cos-standard-luu")
				restoreOptionsModel.BackupFile = core.StringPtr("archive-2022-03-16-140004Z.tar")
				restoreOptionsModel.BackupID = core.StringPtr("dbaas23-20022-20220410-140000F")
				restoreOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = hpdbService.Restore(restoreOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke Restore with error: Operation validation and request error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the CosHmacKeys model
				cosHmacKeysModel := new(hpdbv3.CosHmacKeys)
				cosHmacKeysModel.AccessKeyID = core.StringPtr("e4465d50a56f401a81d275f55c57bc2f")
				cosHmacKeysModel.SecretAccessKey = core.StringPtr("0c29c8299dbba1d6xx7191d6xx5ce3e7eb601fa4bd9f5a")

				// Construct an instance of the RestoreOptions model
				restoreOptionsModel := new(hpdbv3.RestoreOptions)
				restoreOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				restoreOptionsModel.SourceType = core.StringPtr("cos")
				restoreOptionsModel.CosHmacKeys = cosHmacKeysModel
				restoreOptionsModel.CosEndpoint = core.StringPtr("s3.us-west.cloud-object-storage.test.appdomain.cloud")
				restoreOptionsModel.BucketInstanceCrn = core.StringPtr("crn:v1:staging:public:cloud-object-storage:global:a/23a24a3e3fe7a115473f07be1c44bdb5:3ac008bb-16a8-48b2-b8f7-25a990c14b41:bucket:cloud-object-storage-gc-cos-standard-luu")
				restoreOptionsModel.BackupFile = core.StringPtr("archive-2022-03-16-140004Z.tar")
				restoreOptionsModel.BackupID = core.StringPtr("dbaas23-20022-20220410-140000F")
				restoreOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := hpdbService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := hpdbService.Restore(restoreOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the RestoreOptions model with no property values
				restoreOptionsModelNew := new(hpdbv3.RestoreOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = hpdbService.Restore(restoreOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(202)
				}))
			})
			It(`Invoke Restore successfully`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the CosHmacKeys model
				cosHmacKeysModel := new(hpdbv3.CosHmacKeys)
				cosHmacKeysModel.AccessKeyID = core.StringPtr("e4465d50a56f401a81d275f55c57bc2f")
				cosHmacKeysModel.SecretAccessKey = core.StringPtr("0c29c8299dbba1d6xx7191d6xx5ce3e7eb601fa4bd9f5a")

				// Construct an instance of the RestoreOptions model
				restoreOptionsModel := new(hpdbv3.RestoreOptions)
				restoreOptionsModel.ClusterID = core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")
				restoreOptionsModel.SourceType = core.StringPtr("cos")
				restoreOptionsModel.CosHmacKeys = cosHmacKeysModel
				restoreOptionsModel.CosEndpoint = core.StringPtr("s3.us-west.cloud-object-storage.test.appdomain.cloud")
				restoreOptionsModel.BucketInstanceCrn = core.StringPtr("crn:v1:staging:public:cloud-object-storage:global:a/23a24a3e3fe7a115473f07be1c44bdb5:3ac008bb-16a8-48b2-b8f7-25a990c14b41:bucket:cloud-object-storage-gc-cos-standard-luu")
				restoreOptionsModel.BackupFile = core.StringPtr("archive-2022-03-16-140004Z.tar")
				restoreOptionsModel.BackupID = core.StringPtr("dbaas23-20022-20220410-140000F")
				restoreOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := hpdbService.Restore(restoreOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListNodeLogs(listNodeLogsOptions *ListNodeLogsOptions) - Operation response error`, func() {
		listNodeLogsPath := "/nodes/452ebc6007955ba275cfbbe0f2a78e40/logs"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listNodeLogsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListNodeLogs with error: Operation response processing error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the ListNodeLogsOptions model
				listNodeLogsOptionsModel := new(hpdbv3.ListNodeLogsOptions)
				listNodeLogsOptionsModel.NodeID = core.StringPtr("452ebc6007955ba275cfbbe0f2a78e40")
				listNodeLogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := hpdbService.ListNodeLogs(listNodeLogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				hpdbService.EnableRetries(0, 0)
				result, response, operationErr = hpdbService.ListNodeLogs(listNodeLogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListNodeLogs(listNodeLogsOptions *ListNodeLogsOptions)`, func() {
		listNodeLogsPath := "/nodes/452ebc6007955ba275cfbbe0f2a78e40/logs"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listNodeLogsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"logs": [{"filename": "postgresql.log", "size": 26369, "last_modified": "2021-06-29 07:55:19"}]}`)
				}))
			})
			It(`Invoke ListNodeLogs successfully with retries`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())
				hpdbService.EnableRetries(0, 0)

				// Construct an instance of the ListNodeLogsOptions model
				listNodeLogsOptionsModel := new(hpdbv3.ListNodeLogsOptions)
				listNodeLogsOptionsModel.NodeID = core.StringPtr("452ebc6007955ba275cfbbe0f2a78e40")
				listNodeLogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := hpdbService.ListNodeLogsWithContext(ctx, listNodeLogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				hpdbService.DisableRetries()
				result, response, operationErr := hpdbService.ListNodeLogs(listNodeLogsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = hpdbService.ListNodeLogsWithContext(ctx, listNodeLogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listNodeLogsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"logs": [{"filename": "postgresql.log", "size": 26369, "last_modified": "2021-06-29 07:55:19"}]}`)
				}))
			})
			It(`Invoke ListNodeLogs successfully`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := hpdbService.ListNodeLogs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListNodeLogsOptions model
				listNodeLogsOptionsModel := new(hpdbv3.ListNodeLogsOptions)
				listNodeLogsOptionsModel.NodeID = core.StringPtr("452ebc6007955ba275cfbbe0f2a78e40")
				listNodeLogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = hpdbService.ListNodeLogs(listNodeLogsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListNodeLogs with error: Operation validation and request error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the ListNodeLogsOptions model
				listNodeLogsOptionsModel := new(hpdbv3.ListNodeLogsOptions)
				listNodeLogsOptionsModel.NodeID = core.StringPtr("452ebc6007955ba275cfbbe0f2a78e40")
				listNodeLogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := hpdbService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := hpdbService.ListNodeLogs(listNodeLogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListNodeLogsOptions model with no property values
				listNodeLogsOptionsModelNew := new(hpdbv3.ListNodeLogsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = hpdbService.ListNodeLogs(listNodeLogsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListNodeLogs successfully`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the ListNodeLogsOptions model
				listNodeLogsOptionsModel := new(hpdbv3.ListNodeLogsOptions)
				listNodeLogsOptionsModel.NodeID = core.StringPtr("452ebc6007955ba275cfbbe0f2a78e40")
				listNodeLogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := hpdbService.ListNodeLogs(listNodeLogsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetLog(getLogOptions *GetLogOptions)`, func() {
		getLogPath := "/nodes/452ebc6007955ba275cfbbe0f2a78e40/logs/audit.log"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLogPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "application/json")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke GetLog successfully with retries`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())
				hpdbService.EnableRetries(0, 0)

				// Construct an instance of the GetLogOptions model
				getLogOptionsModel := new(hpdbv3.GetLogOptions)
				getLogOptionsModel.NodeID = core.StringPtr("452ebc6007955ba275cfbbe0f2a78e40")
				getLogOptionsModel.LogName = core.StringPtr("audit.log")
				getLogOptionsModel.Accept = core.StringPtr("application/json")
				getLogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := hpdbService.GetLogWithContext(ctx, getLogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				hpdbService.DisableRetries()
				result, response, operationErr := hpdbService.GetLog(getLogOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = hpdbService.GetLogWithContext(ctx, getLogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLogPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "application/json")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke GetLog successfully`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := hpdbService.GetLog(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetLogOptions model
				getLogOptionsModel := new(hpdbv3.GetLogOptions)
				getLogOptionsModel.NodeID = core.StringPtr("452ebc6007955ba275cfbbe0f2a78e40")
				getLogOptionsModel.LogName = core.StringPtr("audit.log")
				getLogOptionsModel.Accept = core.StringPtr("application/json")
				getLogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = hpdbService.GetLog(getLogOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetLog with error: Operation validation and request error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the GetLogOptions model
				getLogOptionsModel := new(hpdbv3.GetLogOptions)
				getLogOptionsModel.NodeID = core.StringPtr("452ebc6007955ba275cfbbe0f2a78e40")
				getLogOptionsModel.LogName = core.StringPtr("audit.log")
				getLogOptionsModel.Accept = core.StringPtr("application/json")
				getLogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := hpdbService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := hpdbService.GetLog(getLogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetLogOptions model with no property values
				getLogOptionsModelNew := new(hpdbv3.GetLogOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = hpdbService.GetLog(getLogOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetLog successfully`, func() {
				hpdbService, serviceErr := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the GetLogOptions model
				getLogOptionsModel := new(hpdbv3.GetLogOptions)
				getLogOptionsModel.NodeID = core.StringPtr("452ebc6007955ba275cfbbe0f2a78e40")
				getLogOptionsModel.LogName = core.StringPtr("audit.log")
				getLogOptionsModel.Accept = core.StringPtr("application/json")
				getLogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := hpdbService.GetLog(getLogOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())


				// Verify empty byte buffer.
				Expect(result).ToNot(BeNil())
				buffer, operationErr := ioutil.ReadAll(result)
				Expect(operationErr).To(BeNil())
				Expect(buffer).ToNot(BeNil())
				Expect(len(buffer)).To(Equal(0))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			hpdbService, _ := hpdbv3.NewHpdbV3(&hpdbv3.HpdbV3Options{
				URL:           "http://hpdbv3modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewDisableCosBackupOptions successfully`, func() {
				// Construct an instance of the DisableCosBackupOptions model
				clusterID := "9cebab98-afeb-4886-9a29-8e741716e7ff"
				disableCosBackupOptionsModel := hpdbService.NewDisableCosBackupOptions(clusterID)
				disableCosBackupOptionsModel.SetClusterID("9cebab98-afeb-4886-9a29-8e741716e7ff")
				disableCosBackupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(disableCosBackupOptionsModel).ToNot(BeNil())
				Expect(disableCosBackupOptionsModel.ClusterID).To(Equal(core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")))
				Expect(disableCosBackupOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewEnableCosBackupOptions successfully`, func() {
				// Construct an instance of the CosHmacKeys model
				cosHmacKeysModel := new(hpdbv3.CosHmacKeys)
				Expect(cosHmacKeysModel).ToNot(BeNil())
				cosHmacKeysModel.AccessKeyID = core.StringPtr("e4465d50a56f401a81d275f55c57bc2f")
				cosHmacKeysModel.SecretAccessKey = core.StringPtr("0c29c8299dbba1d6xx7191d6xx5ce3e7eb601fa4bd9f5a")
				Expect(cosHmacKeysModel.AccessKeyID).To(Equal(core.StringPtr("e4465d50a56f401a81d275f55c57bc2f")))
				Expect(cosHmacKeysModel.SecretAccessKey).To(Equal(core.StringPtr("0c29c8299dbba1d6xx7191d6xx5ce3e7eb601fa4bd9f5a")))

				// Construct an instance of the EnableCosBackupOptions model
				clusterID := "9cebab98-afeb-4886-9a29-8e741716e7ff"
				enableCosBackupOptionsModel := hpdbService.NewEnableCosBackupOptions(clusterID)
				enableCosBackupOptionsModel.SetClusterID("9cebab98-afeb-4886-9a29-8e741716e7ff")
				enableCosBackupOptionsModel.SetCosHmacKeys(cosHmacKeysModel)
				enableCosBackupOptionsModel.SetCosEndpoint("s3.us-west.cloud-object-storage.test.appdomain.cloud")
				enableCosBackupOptionsModel.SetBucketInstanceCrn("crn:v1:staging:public:cloud-object-storage:global:a/23a24a3e3fe7a115473f07be1c44bdb5:3ac008bb-16a8-48b2-b8f7-25a990c14b41:bucket:cloud-object-storage-gc-cos-standard-luu")
				enableCosBackupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(enableCosBackupOptionsModel).ToNot(BeNil())
				Expect(enableCosBackupOptionsModel.ClusterID).To(Equal(core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")))
				Expect(enableCosBackupOptionsModel.CosHmacKeys).To(Equal(cosHmacKeysModel))
				Expect(enableCosBackupOptionsModel.CosEndpoint).To(Equal(core.StringPtr("s3.us-west.cloud-object-storage.test.appdomain.cloud")))
				Expect(enableCosBackupOptionsModel.BucketInstanceCrn).To(Equal(core.StringPtr("crn:v1:staging:public:cloud-object-storage:global:a/23a24a3e3fe7a115473f07be1c44bdb5:3ac008bb-16a8-48b2-b8f7-25a990c14b41:bucket:cloud-object-storage-gc-cos-standard-luu")))
				Expect(enableCosBackupOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetBackupConfigOptions successfully`, func() {
				// Construct an instance of the GetBackupConfigOptions model
				clusterID := "9cebab98-afeb-4886-9a29-8e741716e7ff"
				getBackupConfigOptionsModel := hpdbService.NewGetBackupConfigOptions(clusterID)
				getBackupConfigOptionsModel.SetClusterID("9cebab98-afeb-4886-9a29-8e741716e7ff")
				getBackupConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getBackupConfigOptionsModel).ToNot(BeNil())
				Expect(getBackupConfigOptionsModel.ClusterID).To(Equal(core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")))
				Expect(getBackupConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetClusterOptions successfully`, func() {
				// Construct an instance of the GetClusterOptions model
				clusterID := "9cebab98-afeb-4886-9a29-8e741716e7ff"
				getClusterOptionsModel := hpdbService.NewGetClusterOptions(clusterID)
				getClusterOptionsModel.SetClusterID("9cebab98-afeb-4886-9a29-8e741716e7ff")
				getClusterOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getClusterOptionsModel).ToNot(BeNil())
				Expect(getClusterOptionsModel.ClusterID).To(Equal(core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")))
				Expect(getClusterOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetConfigurationOptions successfully`, func() {
				// Construct an instance of the GetConfigurationOptions model
				clusterID := "9cebab98-afeb-4886-9a29-8e741716e7ff"
				getConfigurationOptionsModel := hpdbService.NewGetConfigurationOptions(clusterID)
				getConfigurationOptionsModel.SetClusterID("9cebab98-afeb-4886-9a29-8e741716e7ff")
				getConfigurationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getConfigurationOptionsModel).ToNot(BeNil())
				Expect(getConfigurationOptionsModel.ClusterID).To(Equal(core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")))
				Expect(getConfigurationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCosBackupConfigOptions successfully`, func() {
				// Construct an instance of the GetCosBackupConfigOptions model
				clusterID := "9cebab98-afeb-4886-9a29-8e741716e7ff"
				getCosBackupConfigOptionsModel := hpdbService.NewGetCosBackupConfigOptions(clusterID)
				getCosBackupConfigOptionsModel.SetClusterID("9cebab98-afeb-4886-9a29-8e741716e7ff")
				getCosBackupConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCosBackupConfigOptionsModel).ToNot(BeNil())
				Expect(getCosBackupConfigOptionsModel.ClusterID).To(Equal(core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")))
				Expect(getCosBackupConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetLogOptions successfully`, func() {
				// Construct an instance of the GetLogOptions model
				nodeID := "452ebc6007955ba275cfbbe0f2a78e40"
				logName := "audit.log"
				getLogOptionsModel := hpdbService.NewGetLogOptions(nodeID, logName)
				getLogOptionsModel.SetNodeID("452ebc6007955ba275cfbbe0f2a78e40")
				getLogOptionsModel.SetLogName("audit.log")
				getLogOptionsModel.SetAccept("application/json")
				getLogOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getLogOptionsModel).ToNot(BeNil())
				Expect(getLogOptionsModel.NodeID).To(Equal(core.StringPtr("452ebc6007955ba275cfbbe0f2a78e40")))
				Expect(getLogOptionsModel.LogName).To(Equal(core.StringPtr("audit.log")))
				Expect(getLogOptionsModel.Accept).To(Equal(core.StringPtr("application/json")))
				Expect(getLogOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetTaskOptions successfully`, func() {
				// Construct an instance of the GetTaskOptions model
				clusterID := "9cebab98-afeb-4886-9a29-8e741716e7ff"
				taskID := "c1a15760-a4f2-11ec-b00a-7f684d1dd53"
				getTaskOptionsModel := hpdbService.NewGetTaskOptions(clusterID, taskID)
				getTaskOptionsModel.SetClusterID("9cebab98-afeb-4886-9a29-8e741716e7ff")
				getTaskOptionsModel.SetTaskID("c1a15760-a4f2-11ec-b00a-7f684d1dd53")
				getTaskOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTaskOptionsModel).ToNot(BeNil())
				Expect(getTaskOptionsModel.ClusterID).To(Equal(core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")))
				Expect(getTaskOptionsModel.TaskID).To(Equal(core.StringPtr("c1a15760-a4f2-11ec-b00a-7f684d1dd53")))
				Expect(getTaskOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetUserOptions successfully`, func() {
				// Construct an instance of the GetUserOptions model
				clusterID := "9cebab98-afeb-4886-9a29-8e741716e7ff"
				dbUserID := "admin.admin"
				getUserOptionsModel := hpdbService.NewGetUserOptions(clusterID, dbUserID)
				getUserOptionsModel.SetClusterID("9cebab98-afeb-4886-9a29-8e741716e7ff")
				getUserOptionsModel.SetDbUserID("admin.admin")
				getUserOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getUserOptionsModel).ToNot(BeNil())
				Expect(getUserOptionsModel.ClusterID).To(Equal(core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")))
				Expect(getUserOptionsModel.DbUserID).To(Equal(core.StringPtr("admin.admin")))
				Expect(getUserOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListBackupsOptions successfully`, func() {
				// Construct an instance of the ListBackupsOptions model
				clusterID := "9cebab98-afeb-4886-9a29-8e741716e7ff"
				listBackupsOptionsModel := hpdbService.NewListBackupsOptions(clusterID)
				listBackupsOptionsModel.SetClusterID("9cebab98-afeb-4886-9a29-8e741716e7ff")
				listBackupsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listBackupsOptionsModel).ToNot(BeNil())
				Expect(listBackupsOptionsModel.ClusterID).To(Equal(core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")))
				Expect(listBackupsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListDatabasesOptions successfully`, func() {
				// Construct an instance of the ListDatabasesOptions model
				clusterID := "9cebab98-afeb-4886-9a29-8e741716e7ff"
				listDatabasesOptionsModel := hpdbService.NewListDatabasesOptions(clusterID)
				listDatabasesOptionsModel.SetClusterID("9cebab98-afeb-4886-9a29-8e741716e7ff")
				listDatabasesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listDatabasesOptionsModel).ToNot(BeNil())
				Expect(listDatabasesOptionsModel.ClusterID).To(Equal(core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")))
				Expect(listDatabasesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListNodeLogsOptions successfully`, func() {
				// Construct an instance of the ListNodeLogsOptions model
				nodeID := "452ebc6007955ba275cfbbe0f2a78e40"
				listNodeLogsOptionsModel := hpdbService.NewListNodeLogsOptions(nodeID)
				listNodeLogsOptionsModel.SetNodeID("452ebc6007955ba275cfbbe0f2a78e40")
				listNodeLogsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listNodeLogsOptionsModel).ToNot(BeNil())
				Expect(listNodeLogsOptionsModel.NodeID).To(Equal(core.StringPtr("452ebc6007955ba275cfbbe0f2a78e40")))
				Expect(listNodeLogsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListTasksOptions successfully`, func() {
				// Construct an instance of the ListTasksOptions model
				clusterID := "9cebab98-afeb-4886-9a29-8e741716e7ff"
				listTasksOptionsModel := hpdbService.NewListTasksOptions(clusterID)
				listTasksOptionsModel.SetClusterID("9cebab98-afeb-4886-9a29-8e741716e7ff")
				listTasksOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listTasksOptionsModel).ToNot(BeNil())
				Expect(listTasksOptionsModel.ClusterID).To(Equal(core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")))
				Expect(listTasksOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListUsersOptions successfully`, func() {
				// Construct an instance of the ListUsersOptions model
				clusterID := "9cebab98-afeb-4886-9a29-8e741716e7ff"
				listUsersOptionsModel := hpdbService.NewListUsersOptions(clusterID)
				listUsersOptionsModel.SetClusterID("9cebab98-afeb-4886-9a29-8e741716e7ff")
				listUsersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listUsersOptionsModel).ToNot(BeNil())
				Expect(listUsersOptionsModel.ClusterID).To(Equal(core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")))
				Expect(listUsersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRestoreOptions successfully`, func() {
				// Construct an instance of the CosHmacKeys model
				cosHmacKeysModel := new(hpdbv3.CosHmacKeys)
				Expect(cosHmacKeysModel).ToNot(BeNil())
				cosHmacKeysModel.AccessKeyID = core.StringPtr("e4465d50a56f401a81d275f55c57bc2f")
				cosHmacKeysModel.SecretAccessKey = core.StringPtr("0c29c8299dbba1d6xx7191d6xx5ce3e7eb601fa4bd9f5a")
				Expect(cosHmacKeysModel.AccessKeyID).To(Equal(core.StringPtr("e4465d50a56f401a81d275f55c57bc2f")))
				Expect(cosHmacKeysModel.SecretAccessKey).To(Equal(core.StringPtr("0c29c8299dbba1d6xx7191d6xx5ce3e7eb601fa4bd9f5a")))

				// Construct an instance of the RestoreOptions model
				clusterID := "9cebab98-afeb-4886-9a29-8e741716e7ff"
				restoreOptionsModel := hpdbService.NewRestoreOptions(clusterID)
				restoreOptionsModel.SetClusterID("9cebab98-afeb-4886-9a29-8e741716e7ff")
				restoreOptionsModel.SetSourceType("cos")
				restoreOptionsModel.SetCosHmacKeys(cosHmacKeysModel)
				restoreOptionsModel.SetCosEndpoint("s3.us-west.cloud-object-storage.test.appdomain.cloud")
				restoreOptionsModel.SetBucketInstanceCrn("crn:v1:staging:public:cloud-object-storage:global:a/23a24a3e3fe7a115473f07be1c44bdb5:3ac008bb-16a8-48b2-b8f7-25a990c14b41:bucket:cloud-object-storage-gc-cos-standard-luu")
				restoreOptionsModel.SetBackupFile("archive-2022-03-16-140004Z.tar")
				restoreOptionsModel.SetBackupID("dbaas23-20022-20220410-140000F")
				restoreOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(restoreOptionsModel).ToNot(BeNil())
				Expect(restoreOptionsModel.ClusterID).To(Equal(core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")))
				Expect(restoreOptionsModel.SourceType).To(Equal(core.StringPtr("cos")))
				Expect(restoreOptionsModel.CosHmacKeys).To(Equal(cosHmacKeysModel))
				Expect(restoreOptionsModel.CosEndpoint).To(Equal(core.StringPtr("s3.us-west.cloud-object-storage.test.appdomain.cloud")))
				Expect(restoreOptionsModel.BucketInstanceCrn).To(Equal(core.StringPtr("crn:v1:staging:public:cloud-object-storage:global:a/23a24a3e3fe7a115473f07be1c44bdb5:3ac008bb-16a8-48b2-b8f7-25a990c14b41:bucket:cloud-object-storage-gc-cos-standard-luu")))
				Expect(restoreOptionsModel.BackupFile).To(Equal(core.StringPtr("archive-2022-03-16-140004Z.tar")))
				Expect(restoreOptionsModel.BackupID).To(Equal(core.StringPtr("dbaas23-20022-20220410-140000F")))
				Expect(restoreOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewScaleResourcesOptions successfully`, func() {
				// Construct an instance of the Resources model
				resourcesModel := new(hpdbv3.Resources)
				Expect(resourcesModel).ToNot(BeNil())
				resourcesModel.Cpu = core.Int64Ptr(int64(2))
				resourcesModel.Memory = core.StringPtr("2GiB")
				resourcesModel.Storage = core.StringPtr("5GiB")
				Expect(resourcesModel.Cpu).To(Equal(core.Int64Ptr(int64(2))))
				Expect(resourcesModel.Memory).To(Equal(core.StringPtr("2GiB")))
				Expect(resourcesModel.Storage).To(Equal(core.StringPtr("5GiB")))

				// Construct an instance of the ScaleResourcesOptions model
				clusterID := "9cebab98-afeb-4886-9a29-8e741716e7ff"
				scaleResourcesOptionsModel := hpdbService.NewScaleResourcesOptions(clusterID)
				scaleResourcesOptionsModel.SetClusterID("9cebab98-afeb-4886-9a29-8e741716e7ff")
				scaleResourcesOptionsModel.SetResource(resourcesModel)
				scaleResourcesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(scaleResourcesOptionsModel).ToNot(BeNil())
				Expect(scaleResourcesOptionsModel.ClusterID).To(Equal(core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")))
				Expect(scaleResourcesOptionsModel.Resource).To(Equal(resourcesModel))
				Expect(scaleResourcesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateBackupConfigOptions successfully`, func() {
				// Construct an instance of the CosHmacKeys model
				cosHmacKeysModel := new(hpdbv3.CosHmacKeys)
				Expect(cosHmacKeysModel).ToNot(BeNil())
				cosHmacKeysModel.AccessKeyID = core.StringPtr("e4465d50a56f401a81d275f55c57bc2f")
				cosHmacKeysModel.SecretAccessKey = core.StringPtr("0c29c8299dbba1d6xx7191d6xx5ce3e7eb601fa4bd9f5a")
				Expect(cosHmacKeysModel.AccessKeyID).To(Equal(core.StringPtr("e4465d50a56f401a81d275f55c57bc2f")))
				Expect(cosHmacKeysModel.SecretAccessKey).To(Equal(core.StringPtr("0c29c8299dbba1d6xx7191d6xx5ce3e7eb601fa4bd9f5a")))

				// Construct an instance of the BackupSchedule model
				backupScheduleModel := new(hpdbv3.BackupSchedule)
				Expect(backupScheduleModel).ToNot(BeNil())
				backupScheduleModel.Type = core.StringPtr("frequency")
				backupScheduleModel.Interval = core.StringPtr("PT8H")
				Expect(backupScheduleModel.Type).To(Equal(core.StringPtr("frequency")))
				Expect(backupScheduleModel.Interval).To(Equal(core.StringPtr("PT8H")))

				// Construct an instance of the BackupConfigCos model
				backupConfigCosModel := new(hpdbv3.BackupConfigCos)
				Expect(backupConfigCosModel).ToNot(BeNil())
				backupConfigCosModel.CosHmacKeys = cosHmacKeysModel
				backupConfigCosModel.CosEndpoint = core.StringPtr("s3.us-west.cloud-object-storage.test.appdomain.cloud")
				backupConfigCosModel.BucketInstanceCrn = core.StringPtr("crn:v1:staging:public:cloud-object-storage:global:a/23a24a3e3fe7a115473f07be1c44bdb5:3ac008bb-16a8-48b2-b8f7-25a990c14b41:bucket:cloud-object-storage-gc-cos-standard-luu")
				backupConfigCosModel.Schedule = backupScheduleModel
				Expect(backupConfigCosModel.CosHmacKeys).To(Equal(cosHmacKeysModel))
				Expect(backupConfigCosModel.CosEndpoint).To(Equal(core.StringPtr("s3.us-west.cloud-object-storage.test.appdomain.cloud")))
				Expect(backupConfigCosModel.BucketInstanceCrn).To(Equal(core.StringPtr("crn:v1:staging:public:cloud-object-storage:global:a/23a24a3e3fe7a115473f07be1c44bdb5:3ac008bb-16a8-48b2-b8f7-25a990c14b41:bucket:cloud-object-storage-gc-cos-standard-luu")))
				Expect(backupConfigCosModel.Schedule).To(Equal(backupScheduleModel))

				// Construct an instance of the UpdateBackupConfigOptions model
				clusterID := "9cebab98-afeb-4886-9a29-8e741716e7ff"
				updateBackupConfigOptionsModel := hpdbService.NewUpdateBackupConfigOptions(clusterID)
				updateBackupConfigOptionsModel.SetClusterID("9cebab98-afeb-4886-9a29-8e741716e7ff")
				updateBackupConfigOptionsModel.SetCos(backupConfigCosModel)
				updateBackupConfigOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateBackupConfigOptionsModel).ToNot(BeNil())
				Expect(updateBackupConfigOptionsModel.ClusterID).To(Equal(core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")))
				Expect(updateBackupConfigOptionsModel.Cos).To(Equal(backupConfigCosModel))
				Expect(updateBackupConfigOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateConfigurationOptions successfully`, func() {
				// Construct an instance of the Configurations model
				configurationsModel := new(hpdbv3.Configurations)
				Expect(configurationsModel).ToNot(BeNil())
				configurationsModel.DeadlockTimeout = core.Int64Ptr(int64(10000))
				configurationsModel.MaxLocksPerTransaction = core.Int64Ptr(int64(100))
				configurationsModel.SharedBuffers = core.Int64Ptr(int64(256))
				configurationsModel.MaxConnections = core.Int64Ptr(int64(150))
				Expect(configurationsModel.DeadlockTimeout).To(Equal(core.Int64Ptr(int64(10000))))
				Expect(configurationsModel.MaxLocksPerTransaction).To(Equal(core.Int64Ptr(int64(100))))
				Expect(configurationsModel.SharedBuffers).To(Equal(core.Int64Ptr(int64(256))))
				Expect(configurationsModel.MaxConnections).To(Equal(core.Int64Ptr(int64(150))))

				// Construct an instance of the UpdateConfigurationOptions model
				clusterID := "9cebab98-afeb-4886-9a29-8e741716e7ff"
				updateConfigurationOptionsModel := hpdbService.NewUpdateConfigurationOptions(clusterID)
				updateConfigurationOptionsModel.SetClusterID("9cebab98-afeb-4886-9a29-8e741716e7ff")
				updateConfigurationOptionsModel.SetConfiguration(configurationsModel)
				updateConfigurationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateConfigurationOptionsModel).ToNot(BeNil())
				Expect(updateConfigurationOptionsModel.ClusterID).To(Equal(core.StringPtr("9cebab98-afeb-4886-9a29-8e741716e7ff")))
				Expect(updateConfigurationOptionsModel.Configuration).To(Equal(configurationsModel))
				Expect(updateConfigurationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("This is a test")
			Expect(mockByteArray).ToNot(BeNil())
		})
		It(`Invoke CreateMockUUID() successfully`, func() {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})
		It(`Invoke CreateMockReader() successfully`, func() {
			mockReader := CreateMockReader("This is a test.")
			Expect(mockReader).ToNot(BeNil())
		})
		It(`Invoke CreateMockDate() successfully`, func() {
			mockDate := CreateMockDate("2019-01-01")
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime("2019-01-01T12:00:00.000Z")
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(mockData string) *[]byte {
	ba := make([]byte, 0)
	ba = append(ba, mockData...)
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate(mockData string) *strfmt.Date {
	d, err := core.ParseDate(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func CreateMockDateTime(mockData string) *strfmt.DateTime {
	d, err := core.ParseDateTime(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func SetTestEnvironment(testEnvironment map[string]string) {
	for key, value := range testEnvironment {
		os.Setenv(key, value)
	}
}

func ClearTestEnvironment(testEnvironment map[string]string) {
	for key := range testEnvironment {
		os.Unsetenv(key)
	}
}
