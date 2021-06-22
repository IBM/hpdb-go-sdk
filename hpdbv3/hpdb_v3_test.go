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

var _ = Describe(`HPDBV3`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(hpdbService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(hpdbService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
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
				"HPDB_URL":       "https://hpdbv3/api",
				"HPDB_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				hpdbService, serviceErr := hpdbv3.NewHPDBV3UsingExternalConfig(&hpdbv3.HPDBV3Options{})
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
				hpdbService, serviceErr := hpdbv3.NewHPDBV3UsingExternalConfig(&hpdbv3.HPDBV3Options{
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
				hpdbService, serviceErr := hpdbv3.NewHPDBV3UsingExternalConfig(&hpdbv3.HPDBV3Options{})
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
				"HPDB_URL":       "https://hpdbv3/api",
				"HPDB_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			hpdbService, serviceErr := hpdbv3.NewHPDBV3UsingExternalConfig(&hpdbv3.HPDBV3Options{})

			It(`Instantiate service client with error`, func() {
				Expect(hpdbService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"HPDB_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			hpdbService, serviceErr := hpdbv3.NewHPDBV3UsingExternalConfig(&hpdbv3.HPDBV3Options{
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
		getClusterPath := "/clusters/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getClusterPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCluster with error: Operation response processing error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the GetClusterOptions model
				getClusterOptionsModel := new(hpdbv3.GetClusterOptions)
				getClusterOptionsModel.ClusterID = core.StringPtr("testString")
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
		getClusterPath := "/clusters/testString"
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
					fmt.Fprintf(res, "%s", `{"id": "7f7aebb87d4fa262a18c8faca505b92a", "region": "eu-de", "name": "replica1", "state": "PROVISIONED", "reason": "Reason", "db_type": "mongodb", "public_endpoint": "dbaas510.hyperp-dbaas.cloud.ibm.com:24110", "private_endpoint": "dbaas510.private.hyperp-dbaas.cloud.ibm.com:24110", "log_url": "https://logging.eu-fra.bluemix.net/app/#/kibana5/discover?_g=()&_a=(columns:!(_source),interval:auto,query:(query_string:(analyze_wildcard:!t,query:'cluster_id_str:3d60c20a4709526ad299236881f9d54c')),sort:!('@timestamp',desc))", "metric_url": "https://metrics.stage1.ng.bluemix.net", "replica_count": 3, "resource": {"cpu": 2, "memory": "20 GiB", "storage": "20 TiB"}, "external_key": {"kms_instance": "crn:v1:staging:public:kms:us-south:a/23a24a3e3fe7a115473f07be1c44bdb5:9eeb285a-88e4-4378-b7cf-dbdcd97b5e4e::", "kms_key": "95d5e441-27e7-4715-a8a8-357871722585"}, "nodes": [{"id": "0286c5b91f9f079d9f1df91fceb391f9", "replica_state": "PRIMARY", "replication_lag": 0, "node_state": "RUNNING", "reason": "Reason", "stopped_reason": "EXTERNAL_KEY_DELETED", "name": "dbaas23:27019", "created_at": "2017-06-22T19:10:51Z", "updated_at": "2017-06-22T19:10:51Z", "is_metric_enabled": true, "is_logging_enabled": false}], "created_at": "2017-06-22T19:10:51Z", "updated_at": "2017-06-22T19:10:51Z"}`)
				}))
			})
			It(`Invoke GetCluster successfully with retries`, func() {
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())
				hpdbService.EnableRetries(0, 0)

				// Construct an instance of the GetClusterOptions model
				getClusterOptionsModel := new(hpdbv3.GetClusterOptions)
				getClusterOptionsModel.ClusterID = core.StringPtr("testString")
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
					fmt.Fprintf(res, "%s", `{"id": "7f7aebb87d4fa262a18c8faca505b92a", "region": "eu-de", "name": "replica1", "state": "PROVISIONED", "reason": "Reason", "db_type": "mongodb", "public_endpoint": "dbaas510.hyperp-dbaas.cloud.ibm.com:24110", "private_endpoint": "dbaas510.private.hyperp-dbaas.cloud.ibm.com:24110", "log_url": "https://logging.eu-fra.bluemix.net/app/#/kibana5/discover?_g=()&_a=(columns:!(_source),interval:auto,query:(query_string:(analyze_wildcard:!t,query:'cluster_id_str:3d60c20a4709526ad299236881f9d54c')),sort:!('@timestamp',desc))", "metric_url": "https://metrics.stage1.ng.bluemix.net", "replica_count": 3, "resource": {"cpu": 2, "memory": "20 GiB", "storage": "20 TiB"}, "external_key": {"kms_instance": "crn:v1:staging:public:kms:us-south:a/23a24a3e3fe7a115473f07be1c44bdb5:9eeb285a-88e4-4378-b7cf-dbdcd97b5e4e::", "kms_key": "95d5e441-27e7-4715-a8a8-357871722585"}, "nodes": [{"id": "0286c5b91f9f079d9f1df91fceb391f9", "replica_state": "PRIMARY", "replication_lag": 0, "node_state": "RUNNING", "reason": "Reason", "stopped_reason": "EXTERNAL_KEY_DELETED", "name": "dbaas23:27019", "created_at": "2017-06-22T19:10:51Z", "updated_at": "2017-06-22T19:10:51Z", "is_metric_enabled": true, "is_logging_enabled": false}], "created_at": "2017-06-22T19:10:51Z", "updated_at": "2017-06-22T19:10:51Z"}`)
				}))
			})
			It(`Invoke GetCluster successfully`, func() {
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
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
				getClusterOptionsModel.ClusterID = core.StringPtr("testString")
				getClusterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = hpdbService.GetCluster(getClusterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetCluster with error: Operation validation and request error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the GetClusterOptions model
				getClusterOptionsModel := new(hpdbv3.GetClusterOptions)
				getClusterOptionsModel.ClusterID = core.StringPtr("testString")
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
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the GetClusterOptions model
				getClusterOptionsModel := new(hpdbv3.GetClusterOptions)
				getClusterOptionsModel.ClusterID = core.StringPtr("testString")
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
		listUsersPath := "/clusters/testString/users"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listUsersPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListUsers with error: Operation response processing error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the ListUsersOptions model
				listUsersOptionsModel := new(hpdbv3.ListUsersOptions)
				listUsersOptionsModel.ClusterID = core.StringPtr("testString")
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
		listUsersPath := "/clusters/testString/users"
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
					fmt.Fprintf(res, "%s", `{"users": [{"name": "tom", "auth_db": "mydb", "role_attributes": ["CREATEDB"]}]}`)
				}))
			})
			It(`Invoke ListUsers successfully with retries`, func() {
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())
				hpdbService.EnableRetries(0, 0)

				// Construct an instance of the ListUsersOptions model
				listUsersOptionsModel := new(hpdbv3.ListUsersOptions)
				listUsersOptionsModel.ClusterID = core.StringPtr("testString")
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
					fmt.Fprintf(res, "%s", `{"users": [{"name": "tom", "auth_db": "mydb", "role_attributes": ["CREATEDB"]}]}`)
				}))
			})
			It(`Invoke ListUsers successfully`, func() {
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
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
				listUsersOptionsModel.ClusterID = core.StringPtr("testString")
				listUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = hpdbService.ListUsers(listUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListUsers with error: Operation validation and request error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the ListUsersOptions model
				listUsersOptionsModel := new(hpdbv3.ListUsersOptions)
				listUsersOptionsModel.ClusterID = core.StringPtr("testString")
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
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the ListUsersOptions model
				listUsersOptionsModel := new(hpdbv3.ListUsersOptions)
				listUsersOptionsModel.ClusterID = core.StringPtr("testString")
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
		getUserPath := "/clusters/testString/users/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getUserPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetUser with error: Operation response processing error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the GetUserOptions model
				getUserOptionsModel := new(hpdbv3.GetUserOptions)
				getUserOptionsModel.ClusterID = core.StringPtr("testString")
				getUserOptionsModel.DbUserID = core.StringPtr("testString")
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
		getUserPath := "/clusters/testString/users/testString"
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
					fmt.Fprintf(res, "%s", `{"name": "syrena", "auth_db": "mydb", "db_access": [{"db": "test", "privileges": ["readWrite"]}], "role_attributes": ["CREATEDB"]}`)
				}))
			})
			It(`Invoke GetUser successfully with retries`, func() {
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())
				hpdbService.EnableRetries(0, 0)

				// Construct an instance of the GetUserOptions model
				getUserOptionsModel := new(hpdbv3.GetUserOptions)
				getUserOptionsModel.ClusterID = core.StringPtr("testString")
				getUserOptionsModel.DbUserID = core.StringPtr("testString")
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
					fmt.Fprintf(res, "%s", `{"name": "syrena", "auth_db": "mydb", "db_access": [{"db": "test", "privileges": ["readWrite"]}], "role_attributes": ["CREATEDB"]}`)
				}))
			})
			It(`Invoke GetUser successfully`, func() {
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
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
				getUserOptionsModel.ClusterID = core.StringPtr("testString")
				getUserOptionsModel.DbUserID = core.StringPtr("testString")
				getUserOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = hpdbService.GetUser(getUserOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetUser with error: Operation validation and request error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the GetUserOptions model
				getUserOptionsModel := new(hpdbv3.GetUserOptions)
				getUserOptionsModel.ClusterID = core.StringPtr("testString")
				getUserOptionsModel.DbUserID = core.StringPtr("testString")
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
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the GetUserOptions model
				getUserOptionsModel := new(hpdbv3.GetUserOptions)
				getUserOptionsModel.ClusterID = core.StringPtr("testString")
				getUserOptionsModel.DbUserID = core.StringPtr("testString")
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
		listDatabasesPath := "/clusters/testString/databases"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDatabasesPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListDatabases with error: Operation response processing error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the ListDatabasesOptions model
				listDatabasesOptionsModel := new(hpdbv3.ListDatabasesOptions)
				listDatabasesOptionsModel.ClusterID = core.StringPtr("testString")
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
		listDatabasesPath := "/clusters/testString/databases"
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
					fmt.Fprintf(res, "%s", `{"total_size": 251658240, "databases": [{"name": "test", "size_on_disk": 83886080}]}`)
				}))
			})
			It(`Invoke ListDatabases successfully with retries`, func() {
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())
				hpdbService.EnableRetries(0, 0)

				// Construct an instance of the ListDatabasesOptions model
				listDatabasesOptionsModel := new(hpdbv3.ListDatabasesOptions)
				listDatabasesOptionsModel.ClusterID = core.StringPtr("testString")
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
					fmt.Fprintf(res, "%s", `{"total_size": 251658240, "databases": [{"name": "test", "size_on_disk": 83886080}]}`)
				}))
			})
			It(`Invoke ListDatabases successfully`, func() {
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
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
				listDatabasesOptionsModel.ClusterID = core.StringPtr("testString")
				listDatabasesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = hpdbService.ListDatabases(listDatabasesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListDatabases with error: Operation validation and request error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the ListDatabasesOptions model
				listDatabasesOptionsModel := new(hpdbv3.ListDatabasesOptions)
				listDatabasesOptionsModel.ClusterID = core.StringPtr("testString")
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
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the ListDatabasesOptions model
				listDatabasesOptionsModel := new(hpdbv3.ListDatabasesOptions)
				listDatabasesOptionsModel.ClusterID = core.StringPtr("testString")
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
		scaleResourcesPath := "/clusters/testString/resource"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(scaleResourcesPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ScaleResources with error: Operation response processing error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the ScaleResourcesResource model
				scaleResourcesResourceModel := new(hpdbv3.ScaleResourcesResource)
				scaleResourcesResourceModel.Cpu = core.Int64Ptr(int64(2))
				scaleResourcesResourceModel.Memory = core.StringPtr("2GiB")
				scaleResourcesResourceModel.Storage = core.StringPtr("5GiB")

				// Construct an instance of the ScaleResourcesOptions model
				scaleResourcesOptionsModel := new(hpdbv3.ScaleResourcesOptions)
				scaleResourcesOptionsModel.ClusterID = core.StringPtr("testString")
				scaleResourcesOptionsModel.Resource = scaleResourcesResourceModel
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
		scaleResourcesPath := "/clusters/testString/resource"
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
					fmt.Fprintf(res, "%s", `{"task_id": "b62026a0-b5ff-4f9a-8780-ecf28dd32c45"}`)
				}))
			})
			It(`Invoke ScaleResources successfully with retries`, func() {
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())
				hpdbService.EnableRetries(0, 0)

				// Construct an instance of the ScaleResourcesResource model
				scaleResourcesResourceModel := new(hpdbv3.ScaleResourcesResource)
				scaleResourcesResourceModel.Cpu = core.Int64Ptr(int64(2))
				scaleResourcesResourceModel.Memory = core.StringPtr("2GiB")
				scaleResourcesResourceModel.Storage = core.StringPtr("5GiB")

				// Construct an instance of the ScaleResourcesOptions model
				scaleResourcesOptionsModel := new(hpdbv3.ScaleResourcesOptions)
				scaleResourcesOptionsModel.ClusterID = core.StringPtr("testString")
				scaleResourcesOptionsModel.Resource = scaleResourcesResourceModel
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
					fmt.Fprintf(res, "%s", `{"task_id": "b62026a0-b5ff-4f9a-8780-ecf28dd32c45"}`)
				}))
			})
			It(`Invoke ScaleResources successfully`, func() {
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
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

				// Construct an instance of the ScaleResourcesResource model
				scaleResourcesResourceModel := new(hpdbv3.ScaleResourcesResource)
				scaleResourcesResourceModel.Cpu = core.Int64Ptr(int64(2))
				scaleResourcesResourceModel.Memory = core.StringPtr("2GiB")
				scaleResourcesResourceModel.Storage = core.StringPtr("5GiB")

				// Construct an instance of the ScaleResourcesOptions model
				scaleResourcesOptionsModel := new(hpdbv3.ScaleResourcesOptions)
				scaleResourcesOptionsModel.ClusterID = core.StringPtr("testString")
				scaleResourcesOptionsModel.Resource = scaleResourcesResourceModel
				scaleResourcesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = hpdbService.ScaleResources(scaleResourcesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ScaleResources with error: Operation validation and request error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the ScaleResourcesResource model
				scaleResourcesResourceModel := new(hpdbv3.ScaleResourcesResource)
				scaleResourcesResourceModel.Cpu = core.Int64Ptr(int64(2))
				scaleResourcesResourceModel.Memory = core.StringPtr("2GiB")
				scaleResourcesResourceModel.Storage = core.StringPtr("5GiB")

				// Construct an instance of the ScaleResourcesOptions model
				scaleResourcesOptionsModel := new(hpdbv3.ScaleResourcesOptions)
				scaleResourcesOptionsModel.ClusterID = core.StringPtr("testString")
				scaleResourcesOptionsModel.Resource = scaleResourcesResourceModel
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
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the ScaleResourcesResource model
				scaleResourcesResourceModel := new(hpdbv3.ScaleResourcesResource)
				scaleResourcesResourceModel.Cpu = core.Int64Ptr(int64(2))
				scaleResourcesResourceModel.Memory = core.StringPtr("2GiB")
				scaleResourcesResourceModel.Storage = core.StringPtr("5GiB")

				// Construct an instance of the ScaleResourcesOptions model
				scaleResourcesOptionsModel := new(hpdbv3.ScaleResourcesOptions)
				scaleResourcesOptionsModel.ClusterID = core.StringPtr("testString")
				scaleResourcesOptionsModel.Resource = scaleResourcesResourceModel
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
		getConfigurationPath := "/clusters/testString/configuration"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getConfigurationPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetConfiguration with error: Operation response processing error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the GetConfigurationOptions model
				getConfigurationOptionsModel := new(hpdbv3.GetConfigurationOptions)
				getConfigurationOptionsModel.ClusterID = core.StringPtr("testString")
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
		getConfigurationPath := "/clusters/testString/configuration"
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
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())
				hpdbService.EnableRetries(0, 0)

				// Construct an instance of the GetConfigurationOptions model
				getConfigurationOptionsModel := new(hpdbv3.GetConfigurationOptions)
				getConfigurationOptionsModel.ClusterID = core.StringPtr("testString")
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
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
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
				getConfigurationOptionsModel.ClusterID = core.StringPtr("testString")
				getConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = hpdbService.GetConfiguration(getConfigurationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetConfiguration with error: Operation validation and request error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the GetConfigurationOptions model
				getConfigurationOptionsModel := new(hpdbv3.GetConfigurationOptions)
				getConfigurationOptionsModel.ClusterID = core.StringPtr("testString")
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
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the GetConfigurationOptions model
				getConfigurationOptionsModel := new(hpdbv3.GetConfigurationOptions)
				getConfigurationOptionsModel.ClusterID = core.StringPtr("testString")
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
		updateConfigurationPath := "/clusters/testString/configuration"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateConfigurationPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["X-Auth-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateConfiguration with error: Operation response processing error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the UpdateConfigurationDataConfiguration model
				updateConfigurationDataConfigurationModel := new(hpdbv3.UpdateConfigurationDataConfiguration)
				updateConfigurationDataConfigurationModel.DeadlockTimeout = core.Int64Ptr(int64(10000))
				updateConfigurationDataConfigurationModel.MaxLocksPerTransaction = core.Int64Ptr(int64(100))
				updateConfigurationDataConfigurationModel.SharedBuffers = core.Int64Ptr(int64(256))
				updateConfigurationDataConfigurationModel.MaxConnections = core.Int64Ptr(int64(150))

				// Construct an instance of the UpdateConfigurationOptions model
				updateConfigurationOptionsModel := new(hpdbv3.UpdateConfigurationOptions)
				updateConfigurationOptionsModel.ClusterID = core.StringPtr("testString")
				updateConfigurationOptionsModel.XAuthToken = core.StringPtr("testString")
				updateConfigurationOptionsModel.Configuration = updateConfigurationDataConfigurationModel
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
		updateConfigurationPath := "/clusters/testString/configuration"
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

					Expect(req.Header["X-Auth-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"task_id": "8254d870-b485-11ea-92be-757a89e2da77"}`)
				}))
			})
			It(`Invoke UpdateConfiguration successfully with retries`, func() {
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())
				hpdbService.EnableRetries(0, 0)

				// Construct an instance of the UpdateConfigurationDataConfiguration model
				updateConfigurationDataConfigurationModel := new(hpdbv3.UpdateConfigurationDataConfiguration)
				updateConfigurationDataConfigurationModel.DeadlockTimeout = core.Int64Ptr(int64(10000))
				updateConfigurationDataConfigurationModel.MaxLocksPerTransaction = core.Int64Ptr(int64(100))
				updateConfigurationDataConfigurationModel.SharedBuffers = core.Int64Ptr(int64(256))
				updateConfigurationDataConfigurationModel.MaxConnections = core.Int64Ptr(int64(150))

				// Construct an instance of the UpdateConfigurationOptions model
				updateConfigurationOptionsModel := new(hpdbv3.UpdateConfigurationOptions)
				updateConfigurationOptionsModel.ClusterID = core.StringPtr("testString")
				updateConfigurationOptionsModel.XAuthToken = core.StringPtr("testString")
				updateConfigurationOptionsModel.Configuration = updateConfigurationDataConfigurationModel
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

					Expect(req.Header["X-Auth-Token"]).ToNot(BeNil())
					Expect(req.Header["X-Auth-Token"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"task_id": "8254d870-b485-11ea-92be-757a89e2da77"}`)
				}))
			})
			It(`Invoke UpdateConfiguration successfully`, func() {
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
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

				// Construct an instance of the UpdateConfigurationDataConfiguration model
				updateConfigurationDataConfigurationModel := new(hpdbv3.UpdateConfigurationDataConfiguration)
				updateConfigurationDataConfigurationModel.DeadlockTimeout = core.Int64Ptr(int64(10000))
				updateConfigurationDataConfigurationModel.MaxLocksPerTransaction = core.Int64Ptr(int64(100))
				updateConfigurationDataConfigurationModel.SharedBuffers = core.Int64Ptr(int64(256))
				updateConfigurationDataConfigurationModel.MaxConnections = core.Int64Ptr(int64(150))

				// Construct an instance of the UpdateConfigurationOptions model
				updateConfigurationOptionsModel := new(hpdbv3.UpdateConfigurationOptions)
				updateConfigurationOptionsModel.ClusterID = core.StringPtr("testString")
				updateConfigurationOptionsModel.XAuthToken = core.StringPtr("testString")
				updateConfigurationOptionsModel.Configuration = updateConfigurationDataConfigurationModel
				updateConfigurationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = hpdbService.UpdateConfiguration(updateConfigurationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateConfiguration with error: Operation validation and request error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the UpdateConfigurationDataConfiguration model
				updateConfigurationDataConfigurationModel := new(hpdbv3.UpdateConfigurationDataConfiguration)
				updateConfigurationDataConfigurationModel.DeadlockTimeout = core.Int64Ptr(int64(10000))
				updateConfigurationDataConfigurationModel.MaxLocksPerTransaction = core.Int64Ptr(int64(100))
				updateConfigurationDataConfigurationModel.SharedBuffers = core.Int64Ptr(int64(256))
				updateConfigurationDataConfigurationModel.MaxConnections = core.Int64Ptr(int64(150))

				// Construct an instance of the UpdateConfigurationOptions model
				updateConfigurationOptionsModel := new(hpdbv3.UpdateConfigurationOptions)
				updateConfigurationOptionsModel.ClusterID = core.StringPtr("testString")
				updateConfigurationOptionsModel.XAuthToken = core.StringPtr("testString")
				updateConfigurationOptionsModel.Configuration = updateConfigurationDataConfigurationModel
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
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the UpdateConfigurationDataConfiguration model
				updateConfigurationDataConfigurationModel := new(hpdbv3.UpdateConfigurationDataConfiguration)
				updateConfigurationDataConfigurationModel.DeadlockTimeout = core.Int64Ptr(int64(10000))
				updateConfigurationDataConfigurationModel.MaxLocksPerTransaction = core.Int64Ptr(int64(100))
				updateConfigurationDataConfigurationModel.SharedBuffers = core.Int64Ptr(int64(256))
				updateConfigurationDataConfigurationModel.MaxConnections = core.Int64Ptr(int64(150))

				// Construct an instance of the UpdateConfigurationOptions model
				updateConfigurationOptionsModel := new(hpdbv3.UpdateConfigurationOptions)
				updateConfigurationOptionsModel.ClusterID = core.StringPtr("testString")
				updateConfigurationOptionsModel.XAuthToken = core.StringPtr("testString")
				updateConfigurationOptionsModel.Configuration = updateConfigurationDataConfigurationModel
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
		listTasksPath := "/clusters/testString/tasks"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTasksPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListTasks with error: Operation response processing error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the ListTasksOptions model
				listTasksOptionsModel := new(hpdbv3.ListTasksOptions)
				listTasksOptionsModel.ClusterID = core.StringPtr("testString")
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
		listTasksPath := "/clusters/testString/tasks"
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
					fmt.Fprintf(res, "%s", `{"tasks": [{"id": "0286c5b91f9f079d9f1df91fceb391f9", "type": "configuration_update", "state": "RUNNING", "reason": "Reason", "started_at": "2020-06-22T19:10:51Z", "finished_at": "2020-06-22T19:11:51Z"}]}`)
				}))
			})
			It(`Invoke ListTasks successfully with retries`, func() {
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())
				hpdbService.EnableRetries(0, 0)

				// Construct an instance of the ListTasksOptions model
				listTasksOptionsModel := new(hpdbv3.ListTasksOptions)
				listTasksOptionsModel.ClusterID = core.StringPtr("testString")
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
					fmt.Fprintf(res, "%s", `{"tasks": [{"id": "0286c5b91f9f079d9f1df91fceb391f9", "type": "configuration_update", "state": "RUNNING", "reason": "Reason", "started_at": "2020-06-22T19:10:51Z", "finished_at": "2020-06-22T19:11:51Z"}]}`)
				}))
			})
			It(`Invoke ListTasks successfully`, func() {
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
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
				listTasksOptionsModel.ClusterID = core.StringPtr("testString")
				listTasksOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = hpdbService.ListTasks(listTasksOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListTasks with error: Operation validation and request error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the ListTasksOptions model
				listTasksOptionsModel := new(hpdbv3.ListTasksOptions)
				listTasksOptionsModel.ClusterID = core.StringPtr("testString")
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
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the ListTasksOptions model
				listTasksOptionsModel := new(hpdbv3.ListTasksOptions)
				listTasksOptionsModel.ClusterID = core.StringPtr("testString")
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
		getTaskPath := "/clusters/testString/tasks/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTaskPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetTask with error: Operation response processing error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the GetTaskOptions model
				getTaskOptionsModel := new(hpdbv3.GetTaskOptions)
				getTaskOptionsModel.ClusterID = core.StringPtr("testString")
				getTaskOptionsModel.TaskID = core.StringPtr("testString")
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
		getTaskPath := "/clusters/testString/tasks/testString"
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
					fmt.Fprintf(res, "%s", `{"id": "8254d870-b485-11ea-92be-757a89e2da77", "type": "configuration_update", "started_at": "2017-06-22T19:10:51Z", "finished_at": "2017-06-22T19:10:51Z", "reason": "Reason", "nodes": [{"id": "0286c5b91f9f079d9f1df91fceb391f9", "state": "RUNNING", "reason": "Reason", "started_at": "2020-06-22T19:10:51Z", "finished_at": "2020-06-22T19:11:51Z"}], "spec": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke GetTask successfully with retries`, func() {
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())
				hpdbService.EnableRetries(0, 0)

				// Construct an instance of the GetTaskOptions model
				getTaskOptionsModel := new(hpdbv3.GetTaskOptions)
				getTaskOptionsModel.ClusterID = core.StringPtr("testString")
				getTaskOptionsModel.TaskID = core.StringPtr("testString")
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
					fmt.Fprintf(res, "%s", `{"id": "8254d870-b485-11ea-92be-757a89e2da77", "type": "configuration_update", "started_at": "2017-06-22T19:10:51Z", "finished_at": "2017-06-22T19:10:51Z", "reason": "Reason", "nodes": [{"id": "0286c5b91f9f079d9f1df91fceb391f9", "state": "RUNNING", "reason": "Reason", "started_at": "2020-06-22T19:10:51Z", "finished_at": "2020-06-22T19:11:51Z"}], "spec": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke GetTask successfully`, func() {
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
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
				getTaskOptionsModel.ClusterID = core.StringPtr("testString")
				getTaskOptionsModel.TaskID = core.StringPtr("testString")
				getTaskOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = hpdbService.GetTask(getTaskOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetTask with error: Operation validation and request error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the GetTaskOptions model
				getTaskOptionsModel := new(hpdbv3.GetTaskOptions)
				getTaskOptionsModel.ClusterID = core.StringPtr("testString")
				getTaskOptionsModel.TaskID = core.StringPtr("testString")
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
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the GetTaskOptions model
				getTaskOptionsModel := new(hpdbv3.GetTaskOptions)
				getTaskOptionsModel.ClusterID = core.StringPtr("testString")
				getTaskOptionsModel.TaskID = core.StringPtr("testString")
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
	Describe(`ListNodeLogs(listNodeLogsOptions *ListNodeLogsOptions) - Operation response error`, func() {
		listNodeLogsPath := "/nodes/testString/logs"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listNodeLogsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListNodeLogs with error: Operation response processing error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the ListNodeLogsOptions model
				listNodeLogsOptionsModel := new(hpdbv3.ListNodeLogsOptions)
				listNodeLogsOptionsModel.NodeID = core.StringPtr("testString")
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
		listNodeLogsPath := "/nodes/testString/logs"
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
					fmt.Fprintf(res, "%s", `{"logs": [{"filename": "mongod.log.20180604-1528094566", "size": 531015965, "last_modified": "2018-06-04 06:42:46"}]}`)
				}))
			})
			It(`Invoke ListNodeLogs successfully with retries`, func() {
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())
				hpdbService.EnableRetries(0, 0)

				// Construct an instance of the ListNodeLogsOptions model
				listNodeLogsOptionsModel := new(hpdbv3.ListNodeLogsOptions)
				listNodeLogsOptionsModel.NodeID = core.StringPtr("testString")
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
					fmt.Fprintf(res, "%s", `{"logs": [{"filename": "mongod.log.20180604-1528094566", "size": 531015965, "last_modified": "2018-06-04 06:42:46"}]}`)
				}))
			})
			It(`Invoke ListNodeLogs successfully`, func() {
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
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
				listNodeLogsOptionsModel.NodeID = core.StringPtr("testString")
				listNodeLogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = hpdbService.ListNodeLogs(listNodeLogsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListNodeLogs with error: Operation validation and request error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the ListNodeLogsOptions model
				listNodeLogsOptionsModel := new(hpdbv3.ListNodeLogsOptions)
				listNodeLogsOptionsModel.NodeID = core.StringPtr("testString")
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
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the ListNodeLogsOptions model
				listNodeLogsOptionsModel := new(hpdbv3.ListNodeLogsOptions)
				listNodeLogsOptionsModel.NodeID = core.StringPtr("testString")
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
		getLogPath := "/nodes/testString/logs/testString"
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
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())
				hpdbService.EnableRetries(0, 0)

				// Construct an instance of the GetLogOptions model
				getLogOptionsModel := new(hpdbv3.GetLogOptions)
				getLogOptionsModel.NodeID = core.StringPtr("testString")
				getLogOptionsModel.LogName = core.StringPtr("testString")
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
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
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
				getLogOptionsModel.NodeID = core.StringPtr("testString")
				getLogOptionsModel.LogName = core.StringPtr("testString")
				getLogOptionsModel.Accept = core.StringPtr("application/json")
				getLogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = hpdbService.GetLog(getLogOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetLog with error: Operation validation and request error`, func() {
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the GetLogOptions model
				getLogOptionsModel := new(hpdbv3.GetLogOptions)
				getLogOptionsModel.NodeID = core.StringPtr("testString")
				getLogOptionsModel.LogName = core.StringPtr("testString")
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
				hpdbService, serviceErr := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(hpdbService).ToNot(BeNil())

				// Construct an instance of the GetLogOptions model
				getLogOptionsModel := new(hpdbv3.GetLogOptions)
				getLogOptionsModel.NodeID = core.StringPtr("testString")
				getLogOptionsModel.LogName = core.StringPtr("testString")
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
			hpdbService, _ := hpdbv3.NewHPDBV3(&hpdbv3.HPDBV3Options{
				URL:           "http://hpdbv3modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewGetClusterOptions successfully`, func() {
				// Construct an instance of the GetClusterOptions model
				clusterID := "testString"
				getClusterOptionsModel := hpdbService.NewGetClusterOptions(clusterID)
				getClusterOptionsModel.SetClusterID("testString")
				getClusterOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getClusterOptionsModel).ToNot(BeNil())
				Expect(getClusterOptionsModel.ClusterID).To(Equal(core.StringPtr("testString")))
				Expect(getClusterOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetConfigurationOptions successfully`, func() {
				// Construct an instance of the GetConfigurationOptions model
				clusterID := "testString"
				getConfigurationOptionsModel := hpdbService.NewGetConfigurationOptions(clusterID)
				getConfigurationOptionsModel.SetClusterID("testString")
				getConfigurationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getConfigurationOptionsModel).ToNot(BeNil())
				Expect(getConfigurationOptionsModel.ClusterID).To(Equal(core.StringPtr("testString")))
				Expect(getConfigurationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetLogOptions successfully`, func() {
				// Construct an instance of the GetLogOptions model
				nodeID := "testString"
				logName := "testString"
				getLogOptionsModel := hpdbService.NewGetLogOptions(nodeID, logName)
				getLogOptionsModel.SetNodeID("testString")
				getLogOptionsModel.SetLogName("testString")
				getLogOptionsModel.SetAccept("application/json")
				getLogOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getLogOptionsModel).ToNot(BeNil())
				Expect(getLogOptionsModel.NodeID).To(Equal(core.StringPtr("testString")))
				Expect(getLogOptionsModel.LogName).To(Equal(core.StringPtr("testString")))
				Expect(getLogOptionsModel.Accept).To(Equal(core.StringPtr("application/json")))
				Expect(getLogOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetTaskOptions successfully`, func() {
				// Construct an instance of the GetTaskOptions model
				clusterID := "testString"
				taskID := "testString"
				getTaskOptionsModel := hpdbService.NewGetTaskOptions(clusterID, taskID)
				getTaskOptionsModel.SetClusterID("testString")
				getTaskOptionsModel.SetTaskID("testString")
				getTaskOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTaskOptionsModel).ToNot(BeNil())
				Expect(getTaskOptionsModel.ClusterID).To(Equal(core.StringPtr("testString")))
				Expect(getTaskOptionsModel.TaskID).To(Equal(core.StringPtr("testString")))
				Expect(getTaskOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetUserOptions successfully`, func() {
				// Construct an instance of the GetUserOptions model
				clusterID := "testString"
				dbUserID := "testString"
				getUserOptionsModel := hpdbService.NewGetUserOptions(clusterID, dbUserID)
				getUserOptionsModel.SetClusterID("testString")
				getUserOptionsModel.SetDbUserID("testString")
				getUserOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getUserOptionsModel).ToNot(BeNil())
				Expect(getUserOptionsModel.ClusterID).To(Equal(core.StringPtr("testString")))
				Expect(getUserOptionsModel.DbUserID).To(Equal(core.StringPtr("testString")))
				Expect(getUserOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListDatabasesOptions successfully`, func() {
				// Construct an instance of the ListDatabasesOptions model
				clusterID := "testString"
				listDatabasesOptionsModel := hpdbService.NewListDatabasesOptions(clusterID)
				listDatabasesOptionsModel.SetClusterID("testString")
				listDatabasesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listDatabasesOptionsModel).ToNot(BeNil())
				Expect(listDatabasesOptionsModel.ClusterID).To(Equal(core.StringPtr("testString")))
				Expect(listDatabasesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListNodeLogsOptions successfully`, func() {
				// Construct an instance of the ListNodeLogsOptions model
				nodeID := "testString"
				listNodeLogsOptionsModel := hpdbService.NewListNodeLogsOptions(nodeID)
				listNodeLogsOptionsModel.SetNodeID("testString")
				listNodeLogsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listNodeLogsOptionsModel).ToNot(BeNil())
				Expect(listNodeLogsOptionsModel.NodeID).To(Equal(core.StringPtr("testString")))
				Expect(listNodeLogsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListTasksOptions successfully`, func() {
				// Construct an instance of the ListTasksOptions model
				clusterID := "testString"
				listTasksOptionsModel := hpdbService.NewListTasksOptions(clusterID)
				listTasksOptionsModel.SetClusterID("testString")
				listTasksOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listTasksOptionsModel).ToNot(BeNil())
				Expect(listTasksOptionsModel.ClusterID).To(Equal(core.StringPtr("testString")))
				Expect(listTasksOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListUsersOptions successfully`, func() {
				// Construct an instance of the ListUsersOptions model
				clusterID := "testString"
				listUsersOptionsModel := hpdbService.NewListUsersOptions(clusterID)
				listUsersOptionsModel.SetClusterID("testString")
				listUsersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listUsersOptionsModel).ToNot(BeNil())
				Expect(listUsersOptionsModel.ClusterID).To(Equal(core.StringPtr("testString")))
				Expect(listUsersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewScaleResourcesOptions successfully`, func() {
				// Construct an instance of the ScaleResourcesResource model
				scaleResourcesResourceModel := new(hpdbv3.ScaleResourcesResource)
				Expect(scaleResourcesResourceModel).ToNot(BeNil())
				scaleResourcesResourceModel.Cpu = core.Int64Ptr(int64(2))
				scaleResourcesResourceModel.Memory = core.StringPtr("2GiB")
				scaleResourcesResourceModel.Storage = core.StringPtr("5GiB")
				Expect(scaleResourcesResourceModel.Cpu).To(Equal(core.Int64Ptr(int64(2))))
				Expect(scaleResourcesResourceModel.Memory).To(Equal(core.StringPtr("2GiB")))
				Expect(scaleResourcesResourceModel.Storage).To(Equal(core.StringPtr("5GiB")))

				// Construct an instance of the ScaleResourcesOptions model
				clusterID := "testString"
				scaleResourcesOptionsModel := hpdbService.NewScaleResourcesOptions(clusterID)
				scaleResourcesOptionsModel.SetClusterID("testString")
				scaleResourcesOptionsModel.SetResource(scaleResourcesResourceModel)
				scaleResourcesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(scaleResourcesOptionsModel).ToNot(BeNil())
				Expect(scaleResourcesOptionsModel.ClusterID).To(Equal(core.StringPtr("testString")))
				Expect(scaleResourcesOptionsModel.Resource).To(Equal(scaleResourcesResourceModel))
				Expect(scaleResourcesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateConfigurationOptions successfully`, func() {
				// Construct an instance of the UpdateConfigurationDataConfiguration model
				updateConfigurationDataConfigurationModel := new(hpdbv3.UpdateConfigurationDataConfiguration)
				Expect(updateConfigurationDataConfigurationModel).ToNot(BeNil())
				updateConfigurationDataConfigurationModel.DeadlockTimeout = core.Int64Ptr(int64(10000))
				updateConfigurationDataConfigurationModel.MaxLocksPerTransaction = core.Int64Ptr(int64(100))
				updateConfigurationDataConfigurationModel.SharedBuffers = core.Int64Ptr(int64(256))
				updateConfigurationDataConfigurationModel.MaxConnections = core.Int64Ptr(int64(150))
				Expect(updateConfigurationDataConfigurationModel.DeadlockTimeout).To(Equal(core.Int64Ptr(int64(10000))))
				Expect(updateConfigurationDataConfigurationModel.MaxLocksPerTransaction).To(Equal(core.Int64Ptr(int64(100))))
				Expect(updateConfigurationDataConfigurationModel.SharedBuffers).To(Equal(core.Int64Ptr(int64(256))))
				Expect(updateConfigurationDataConfigurationModel.MaxConnections).To(Equal(core.Int64Ptr(int64(150))))

				// Construct an instance of the UpdateConfigurationOptions model
				clusterID := "testString"
				xAuthToken := "testString"
				updateConfigurationOptionsModel := hpdbService.NewUpdateConfigurationOptions(clusterID, xAuthToken)
				updateConfigurationOptionsModel.SetClusterID("testString")
				updateConfigurationOptionsModel.SetXAuthToken("testString")
				updateConfigurationOptionsModel.SetConfiguration(updateConfigurationDataConfigurationModel)
				updateConfigurationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateConfigurationOptionsModel).ToNot(BeNil())
				Expect(updateConfigurationOptionsModel.ClusterID).To(Equal(core.StringPtr("testString")))
				Expect(updateConfigurationOptionsModel.XAuthToken).To(Equal(core.StringPtr("testString")))
				Expect(updateConfigurationOptionsModel.Configuration).To(Equal(updateConfigurationDataConfigurationModel))
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
