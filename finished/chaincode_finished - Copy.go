/*
Copyright IBM Corp 2016 All Rights Reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
		 http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"errors"
	"fmt"
    "strconv"	
	
	

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init resets all the things
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	err := stub.PutState("hello_world", []byte(args[0]))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Invoke isur entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {
		return t.Init(stub, "init", args)
	} else if function == "write" {
		return t.write(stub, args)
	}else if function == "add" {
		return t.write(stub, args)
	}
	fmt.Println("invoke did not find func: " + function)

	return nil, errors.New("Received unknown function invocation: " + function)
}

// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	if function == "read" { //read a variable
		return t.read(stub, args)
	}
	fmt.Println("query did not find func: " + function)

	return nil, errors.New("Received unknown function query: " + function)
}

// write - invoke function to write key/value pair
func (t *SimpleChaincode) write(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var key, value string
	var err error
	fmt.Println("running write()")

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the key and value to set")
	}

	key = args[0] //rename for funsies
	value = args[1]
	err = stub.PutState(key, []byte(value)) //write the variable into the chaincode state
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// read - query function to read key/value pair
func (t *SimpleChaincode) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var key, jsonResp string
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
	}

	key = args[0]
	valAsbytes, err := stub.GetState(key)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return nil, errors.New(jsonResp)
	}

	return valAsbytes, nil
}

func (t *SimpleChaincode) add(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var jsonResp string
	var aVal,bVal int
	var sumVal int
	var err error

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
	}

	aVal = strconv.Atoi(args[0])
	bVal = strconv.Atoi(args[1])
    sumVal = strconv.Atoi(args[0])

	err = stub.PutState(A, []byte(strconv.Itoa(aVal)))
	if err != nil {
		return nil, err
	}

	err = stub.PutState(B, []byte(strconv.Itoa(bVal)))
	if err != nil {
		return nil, err
	}

	err = stub.PutState(SUM, []byte(strconv.Itoa(sumVal)))
	if err != nil {
		return nil, err
	}


//addition 

	Avalbytes, err := stub.GetState(A)
	if err != nil {
		return nil, errors.New("Failed to get state")
	}
	if Avalbytes == nil {
		return nil, errors.New("Entity not found")
	}
	Avalue, _ = strconv.Atoi(string(Avalbytes))

	Bvalbytes, err := stub.GetState(B)
	if err != nil {
		return nil, errors.New("Failed to get state")
	}
	if Bvalbytes == nil {
		return nil, errors.New("Entity not found")
	}
	Bvalue, _ = strconv.Atoi(string(Bvalbytes))

	Sumvalbytes, err := stub.GetState(SUM)
	if err != nil {
		return nil, errors.New("Failed to get state")
	}
	if Sumvalbytes == nil {
		return nil, errors.New("Entity not found")
	}
	Sumvalue, _ = strconv.Atoi(string(Sumvalbytes))





    sumVal = Avalue +Bvalue
	fmt.Printf("Aval = %d, Bval = %d , sumVal = %d\n", Avalue, Bvalue,sumVal)
	

	err = stub.PutState(A, []byte(strconv.Itoa(Avalue)))
	if err != nil {
		return nil, err
	}

	err = stub.PutState(B, []byte(strconv.Itoa(Bvalue)))
	if err != nil {
		return nil, err
	}

	err = stub.PutState(SUM, []byte(strconv.Itoa(sumVal)))
	if err != nil {
		return nil, err
	}
	
	
	

	if err != nil {
		jsonResp = "{\"Error\":\"Failed to add two numbers }"
		return nil, errors.New(jsonResp)
	}

	return sumVal, nil
}


