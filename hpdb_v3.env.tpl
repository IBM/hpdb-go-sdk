#
# Rename this file to be "hpdb_v3.env" so that it will be found
# by the ExampleServiceIT integration test.
# Otherwise, the tests will be skipped.
# Also, be sure to following the instructions here:
# https://github.ibm.com/CloudEngineering/java-sdk-template/blob/main/README_FIRST.md#integration-tests
# to start up an instance of the Example Service prior to running the integraton test.
#

HPDB_URL=https://HPDB_ENDPOINT/api/v3/ACCOUNT_ID
HPDB_AUTHTYPE=iam
HPDB_APIKEY=YOUR_API_KEY
HPDB_CLUSTER_CRN=HPDB_SERVICE_INSTANCE_CRN
HPDB_NODE_ID=NODE_ID
HPDB_COS_CRN=COS_BUCKET_CRN
HPDB_COS_ENDPOINT=COS_ENDPOINT
HPDB_COS_ACCESS_KEY_ID=COS_HMAC_ACCESS_KEY
HPDB_COS_SECRET_KEY=COS_HMAC_SECRET_KEY
HPDB_COS_FILE=COS_BACKUP_FILE_NAME