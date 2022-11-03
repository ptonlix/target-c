package agreement

import (
	"testing"

	ledgerapi "github.com/ptonlix/target-c/chaincode-go/ledger-api"
	"github.com/stretchr/testify/assert"
)

func TestGetPaperList(t *testing.T) {
	var tc *TransactionContext
	var expectedAgreementList *list

	tc = new(TransactionContext)
	expectedAgreementList = newList(tc)
	actualList := tc.GetAgreementList().(*list)
	assert.Equal(t, expectedAgreementList.stateList.(*ledgerapi.StateList).Name, actualList.stateList.(*ledgerapi.StateList).Name, "should configure paper list when one not already configured")

	tc = new(TransactionContext)
	expectedAgreementList = new(list)
	expectedStateList := new(ledgerapi.StateList)
	expectedStateList.Ctx = tc
	expectedStateList.Name = "existing paper list"
	expectedAgreementList.stateList = expectedStateList
	tc.agreementList = expectedAgreementList
	assert.Equal(t, expectedAgreementList, tc.GetAgreementList(), "should return set paper list when already set")
}
