[![Build Status](https://travis-ci.com/IBM/hpdb-go-sdk.svg?branch=main)](https://travis-ci.com/IBM/hpdb-go-sdk)
[![semantic-release](https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg)](https://github.com/semantic-release/semantic-release)

# IBM Hyper Protect DBaaS Services Go SDK 0.2.0
Go client library to interact with the various [IBM Hyper Protect DBaaS Services APIs](https://cloud.ibm.com/apidocs/hyperp-dbaas/hyperp-dbaas-v3).

Disclaimer: this SDK is being released initially as a **pre-release** version.
Changes might occur which impact applications that use this SDK.

## Table of Contents
<!--
  The TOC below is generated using the `markdown-toc` node package.

      https://github.com/jonschlinkert/markdown-toc

  You should regenerate the TOC after making changes to this file.

      npx markdown-toc -i README.md
  -->

<!-- toc -->

- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
  * [Go modules](#go-modules)
  * [`go get` command](#go-get-command)
- [Using the SDK](#using-the-sdk)
  * [Sample Code](#sample-code)
- [Questions](#questions)
- [Issues](#issues)
- [Open source @ IBM](#open-source--ibm)
- [Contributing](#contributing)
- [License](#license)

<!-- tocstop -->

## Overview

The IBM Hyper Protect DBaaS Services Go SDK allows developers to programmatically interact with the following IBM Cloud services:

Service Name | Package name 
--- | --- 
[IBM Hyper Protect DBaaS](https://cloud.ibm.com/apidocs/hyperp-dbaas/hyperp-dbaas-v3) | hpdbv3

## Prerequisites

[ibm-cloud-onboarding]: https://cloud.ibm.com/registration

* An [IBM Cloud][ibm-cloud-onboarding] account.
* An IAM API key to allow the SDK to access your account. Create one [here](https://cloud.ibm.com/iam/apikeys).
* Go version 1.14 or above.

## Installation
The current version of this SDK: 0.2.0

### Go modules  
If your application uses Go modules for dependency management (recommended), just add an import for each service 
that you will use in your application.  
Here is an example:

```go
import (
	"github.com/IBM/hpdb-go-sdk/hpdbv3"
)
```
Next, run `go build` or `go mod tidy` to download and install the new dependencies and update your application's
`go.mod` file.  

In the example above, the `hpdbv3` part of the import path is the package name
associated with the IBM Hyper Protect DBaaS service.

### `go get` command  
Alternatively, you can use the `go get` command to download and install the appropriate packages needed by your application:
```
go get -u github.com/IBM/hpdb-go-sdk/hpdbv3
```
Be sure to use the appropriate package name from the service table above for the services used by your application.

## Using the SDK
For general SDK usage information, please see [this link](https://github.com/IBM/ibm-cloud-sdk-common/blob/main/README.md)

IBM Hyper Protect DBaaS features supported by this SDK are as follows.

| Features | Go Function |
|----|----|
| Show database cluster details | GetCluster |
| List all databases | ListDatabases |
| List all database users | ListUsers |
| Show the details of a database user | GetUser |
| List all log files of a database node | ListNodeLogs |
| Download a log file | GetLog |
| Scale resources in a specified cluster | ScaleResources |
| List all tasks | ListTasks  |
| Show the details of a task | GetTask |
| Get database configurations (only for postgresql) | GetConfiguration |
| Update database configurations (only for postgresql) | UpdateConfiguration |
| Enable backups to COS | EnableCosBackup |
| Disable backups to COS | DisableCosBackup |
| Show COS configuration (Deprecated) | GetCosBackupConfig |
| Show backup configuration | GetBackupConfig |
| Update backup configuration | UpdateBackupConfig |
| Restore DB from backup file | Restore |

Please run following command for the detailed usage.

```
go doc -all github.com/IBM/hpdb-go-sdk/hpdbv3
```

### Sample Code

Here's a sample on getting information of a database cluster. You can find more sample code at [IBM Hyper Protect DBaaS API Documentation](https://cloud.ibm.com/apidocs/hyperp-dbaas/hyperp-dbaas-v3)

```
package main

import (
	"fmt"
	"strings"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/hpdb-go-sdk/hpdbv3"
)

func main() {
	const dbClusterCRN = "crn:v1:bluemix:public:hyperp-dbaas-postgresql:eu-de:a/e530ebd25f5ab6b0cf1e889593015f7a:e3dd2973-0e15-4fe3-8a26-f567a76e0b29::"
	const hpdbEndpoint = "dbaas902.hyperp-dbaas.cloud.ibm.com:20000"
	const apiKey = "API_KEY"

	crnSegments := strings.Split(dbClusterCRN, ":")
	accountID := strings.TrimPrefix(crnSegments[6], "a/")
	clusterID := crnSegments[7]

	authenticator := &core.IamAuthenticator{
		ApiKey: apiKey,
	}

	options := &hpdbv3.HpdbV3Options{
		Authenticator: authenticator,
		URL:           fmt.Sprintf("https://%s/api/v3/%s", hpdbEndpoint, accountID),
	}

	hpdb, err := hpdbv3.NewHpdbV3(options)
	if err != nil {
		panic(err)
	}

	hpdb.Service.DisableSSLVerification()

	getClusterOpts := hpdb.NewGetClusterOptions(clusterID)
	cluster, _, err := hpdb.GetCluster(getClusterOpts)
	if err != nil {
		panic(err)
	}
	fmt.Println("Cluster status: ", *cluster.State)
}
```

`dbClusterCRN` is the CRN of your IBM Hyper Protect DBaaS service instance.

`hpdbEndpoint` is the endpoint of IBM Hyper Protect DBaaS service. Different regions have different endpoints. You can find the list [here](https://cloud.ibm.com/docs/hyper-protect-dbaas-for-mongodb?topic=hyper-protect-dbaas-for-mongodb-api-setup#gen_inst_mgr_apis)



## Questions

If you are having difficulties using this SDK or have a question about the IBM Cloud services,
please ask a question at 
[Stack Overflow](http://stackoverflow.com/questions/ask?tags=ibm-cloud).

## Issues
If you encounter an issue with the project, you are welcome to submit a
[bug report](https://github.com/IBM/hpdb-go-sdk/issues).
Before that, please search for similar issues. It's possible that someone has already reported the problem.

## Open source @ IBM
Find more open source projects on the [IBM Github Page](http://ibm.github.io/)

## Contributing
See [CONTRIBUTING](CONTRIBUTING.md).

## License

This SDK project is released under the Apache 2.0 license.
The license's full text can be found in [LICENSE](LICENSE).
