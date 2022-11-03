package main

import (
	"fmt"

	"github.com/ptonlix/target-c/chaincode-go/agreement"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	contract := new(agreement.Contract)
	contract.TransactionContextHandler = new(agreement.TransactionContext)
	contract.Name = "org.targetc-net.agreement"
	contract.Info.Version = "1.0.0"

	chaincode, err := contractapi.NewChaincode(contract)

	if err != nil {
		panic(fmt.Sprintf("Error creating chaincode. %s", err.Error()))
	}

	chaincode.Info.Title = "AgreementChaincode"
	chaincode.Info.Version = "1.0.0"

	err = chaincode.Start()

	if err != nil {
		panic(fmt.Sprintf("Error starting chaincode. %s", err.Error()))
	}

}
