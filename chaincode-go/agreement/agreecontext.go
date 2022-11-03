package agreement

import "github.com/hyperledger/fabric-contract-api-go/contractapi"

type TransactionContextInterface interface {
	contractapi.TransactionContextInterface
	GetAgreementList() ListInterface
}

type TransactionContext struct {
	contractapi.TransactionContext
	agreementList *list
}

func (tc *TransactionContext) GetAgreementList() ListInterface {
	if tc.agreementList == nil {
		tc.agreementList = newList(tc)
	}

	return tc.agreementList
}
