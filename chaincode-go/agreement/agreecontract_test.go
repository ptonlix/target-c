package agreement

import (
	"errors"
	"testing"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// #########
// HELPERS
// #########
type MockAgreementList struct {
	mock.Mock
}

func (mpl *MockAgreementList) AddAgreement(agree *Agreement) error {
	args := mpl.Called(agree)

	return args.Error(0)
}

func (mpl *MockAgreementList) GetAgreement(issuerOrg string, issuerId string, agreementnumber string) (*Agreement, error) {
	args := mpl.Called(issuerOrg, issuerId, agreementnumber)

	return args.Get(0).(*Agreement), args.Error(1)
}

func (mpl *MockAgreementList) UpdateAgreement(agree *Agreement) error {
	args := mpl.Called(agree)

	return args.Error(0)
}

func (mpl *MockAgreementList) GetAllAgreement() ([]*Agreement, error) {
	args := mpl.Called()

	return args.Get(0).([]*Agreement), args.Error(1)
}

func (mpl *MockAgreementList) GetAgreementByRange(startKey, endKey string) ([]*Agreement, error) {
	args := mpl.Called(startKey, endKey)

	return args.Get(0).([]*Agreement), args.Error(1)
}

func (mpl *MockAgreementList) GetAgreementByOrg(issuerOrg string) ([]*Agreement, error) {
	args := mpl.Called(issuerOrg)

	return args.Get(0).([]*Agreement), args.Error(1)
}

func (mpl *MockAgreementList) GetAgreementByOrgAndId(issuerOrg, issuerId string) ([]*Agreement, error) {
	args := mpl.Called(issuerOrg, issuerId)

	return args.Get(0).([]*Agreement), args.Error(1)
}

type MockTransactionContext struct {
	contractapi.TransactionContext
	agreementList *MockAgreementList
}

func (mtc *MockTransactionContext) GetAgreementList() ListInterface {
	return mtc.agreementList
}

func resetAgree(agree *Agreement) {
	agree.IssuerOrg = "someowner"
	agree.SetTrading()
}

// #########
// TESTS
// #########

func TestIssue(t *testing.T) {
	var agree *Agreement
	var err error

	mpl := new(MockAgreementList)
	ctx := new(MockTransactionContext)
	ctx.agreementList = mpl

	contract := new(Contract)

	var sentAgree *Agreement

	mpl.On("AddAgreement", mock.MatchedBy(func(agree *Agreement) bool { sentAgree = agree; return agree.IssuerOrg == "someissuer" })).Return(nil)
	mpl.On("AddAgreement", mock.MatchedBy(func(agree *Agreement) bool { sentAgree = agree; return agree.IssuerOrg == "someotherissuer" })).Return(errors.New("AddAgreement error"))

	expectedAgree := Agreement{AgreementNumber: "somepaper", IssuerId: "someissuerid", IssuerOrg: "someissuer", IssueDateTime: "someissuedate",
		SignerMapList: map[string]*Signatory{"100": &Signatory{Id: "100", Result: false, PublicPem: "123456"}}, SignData: []SignData{},
		state: ISSUED}

	agree, err = contract.Issue(ctx, "someissuer", "someissuerid", "somepaper", "someissuedate", `[{"userId":"100", "publicPem":"123456"}]`)
	assert.Nil(t, err, "should not error when add agreement does not error")
	assert.Equal(t, sentAgree, agree, "should send the same agreement as it returns to add agreement")
	assert.Equal(t, expectedAgree, *agree, "should correctly configure agreement")

	agree, err = contract.Issue(ctx, "someotherissuer", "someissuerid", "somepaper", "someissuedate", `[{"userId":"100", "publicPem":"123456"}]`)
	assert.EqualError(t, err, "AddAgreement error", "should return error when add agreement fails")
	assert.Nil(t, agree, "should not return agreement when fails")
}

func TestSignature(t *testing.T) {
	// var agree *Agreement
	// var err error

	// mpl := new(MockAgreementList)
	// ctx := new(MockTransactionContext)
	// ctx.agreementList = mpl

	// contract := new(Contract)

	// wsAgree := new(Agreement)
	// resetAgree(wsAgree)

	// var sentAgree *Agreement
	// shouldError := false

	// mpl.On("GetAgreement", "someissuerid", "someissuer", "somepaper").Return(wsAgree, nil)
	// mpl.On("UpdateAgreement", mock.MatchedBy(func(agree *Agreement) bool { sentAgree = agree; return !shouldError })).Return(nil)

	// expectedAgree := Agreement{AgreementNumber: "somepaper", IssuerId: "someissuerid", Issuer: "someissuer", IssueDateTime: "someissuedate",
	// 	SignerMapList: map[string]*Signatory{"100": &Signatory{Id: "100", Result: false, PublicPem: "123456"}},
	// 	state:         ISSUED}

	// agree, err = contract.Signature(ctx, "someissuerid", "someissuer", "somepaper", `{"userId":"200", "dataHash":"123456", "signDataHash":"asdqw"}`)
	// assert.Nil(t, err, "should not error when add agreement does not error")
	// assert.Equal(t, sentAgree, agree, "should send the same agreement as it returns to add agreement")
}
