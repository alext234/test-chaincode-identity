package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/chaincode/shim/ext/cid"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type CC struct {
}


func (t *CC) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Chaincode Init")
	return shim.Success(nil)
}

func (t *CC) testGetID(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	
	
	invokeID, err := cid.GetID(stub)
	if err != nil {
		fmt.Println("unable to get ID ", err)
		return shim.Error(err.Error())
	}
	fmt.Println("invokeID = ", invokeID)
	
	
	
	mspID, err := cid.GetMSPID(stub)
	if err != nil {
		fmt.Println("unable to get MSPID ", err)
		return shim.Error(err.Error())
	}
	fmt.Println("MSP ID = ", mspID)
	
	
	attrValue, found, err := cid.GetAttributeValue(stub, "attr1")
	if err != nil {
		fmt.Println("unable to get attribute value ", err)
		return shim.Error(err.Error())
	}
	
	if found {
		fmt.Println("attribute value found; value = ", attrValue)
		} else {
			fmt.Println("attribute value not found")
		}
		
		return shim.Success(nil)
}
	
func (t *CC) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Printf("Chaincode Invoke; function='%s'\n", function)
	if function == "testGetID" {
		return t.testGetID(stub, args)
	}
	
	return shim.Error("Invalid invoke function name")
}
	
func main() {
	err := shim.Start(new(CC))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}
