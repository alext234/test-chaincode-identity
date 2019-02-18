[![Build Status](https://travis-ci.org/alext234/test-chaincode-identity.svg?branch=master)](https://travis-ci.org/alext234/test-chaincode-identity)

Some experiments with chaincode CID to get client identity and attributes.


## Some notes

- See the output of test logs to know how the output of `GetID, GetMSPID, GetAttributeValue()` unit tests look like `https://travis-ci.org/alext234/test-chaincode-identity`


- The `certWithAttrs` in `main_test.go` is lifted from   `https://github.com/hyperledger/fabric/blob/release-1.4/core/chaincode/shim/ext/cid/cid_test.go#L44`

- How do we generate user certificate with attributes?

When registering a user with CA (e.g. using `fabric-ca-client register`) attributes can be specified.

## TODO

- To try out with a local setup of fabric with multiple users to see how values look like.
