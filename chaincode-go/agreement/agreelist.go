package agreement

import ledgerapi "github.com/ptonlix/target-c/chaincode-go/ledger-api"

type ListInterface interface {
	AddAgreement(*Agreement) error
	GetAgreement(string, string, string) (*Agreement, error)
	UpdateAgreement(*Agreement) error
	GetAllAgreement() ([]*Agreement, error)
	GetAgreementByRange(string, string) ([]*Agreement, error)
	GetAgreementByOrg(string) ([]*Agreement, error)
	GetAgreementByOrgAndId(string, string) ([]*Agreement, error)
}

type list struct {
	stateList ledgerapi.StateListInterface
}

func (cpl *list) AddAgreement(agree *Agreement) error {
	return cpl.stateList.AddState(agree)
}

func (cpl *list) GetAgreement(issuerOrg string, issuerId string, agreementNumber string) (*Agreement, error) {
	cp := new(Agreement)

	err := cpl.stateList.GetState(CreateAgreementKey(issuerOrg, issuerId, agreementNumber), cp)

	if err != nil {
		return nil, err
	}

	return cp, nil
}

func (cpl *list) UpdateAgreement(agree *Agreement) error {
	return cpl.stateList.UpdateState(agree)
}

func (cpl *list) GetAllAgreement() ([]*Agreement, error) {

	cp := new(Agreement)
	cplist, err := cpl.stateList.GetStateByRange(cp, "", "")
	if err != nil {
		return nil, err
	}

	agreements := []*Agreement{}
	for _, v := range cplist {
		agreements = append(agreements, v.(*Agreement))
	}
	return agreements, nil
}

func (cpl *list) GetAgreementByRange(startKey, endKey string) ([]*Agreement, error) {

	cp := new(Agreement)
	cplist, err := cpl.stateList.GetStateByRange(cp, startKey, endKey)
	if err != nil {
		return nil, err
	}

	agreements := []*Agreement{}
	for _, v := range cplist {
		agreements = append(agreements, v.(*Agreement))
	}
	return agreements, nil
}

func (cpl *list) GetAgreementByOrg(issuerOrg string) ([]*Agreement, error) {

	cp := new(Agreement)
	cplist, err := cpl.stateList.GetStateByPartialCompositeKey(cp, issuerOrg)
	if err != nil {
		return nil, err
	}

	agreements := []*Agreement{}
	for _, v := range cplist {
		agreements = append(agreements, v.(*Agreement))
	}
	return agreements, nil
}

func (cpl *list) GetAgreementByOrgAndId(issuerOrg, issuerId string) ([]*Agreement, error) {

	cp := new(Agreement)
	cplist, err := cpl.stateList.GetStateByPartialCompositeKey(cp, issuerOrg, issuerId)
	if err != nil {
		return nil, err
	}

	agreements := []*Agreement{}
	for _, v := range cplist {
		agreements = append(agreements, v.(*Agreement))
	}
	return agreements, nil
}

// NewList create a new list from context
func newList(ctx TransactionContextInterface) *list {
	stateList := new(ledgerapi.StateList)
	stateList.Ctx = ctx
	stateList.Name = "org.targetc-net.agreementlist"
	stateList.Deserialize = func(bytes []byte, state ledgerapi.StateInterface) error {
		return Deserialize(bytes, state.(*Agreement))
	}

	list := new(list)
	list.stateList = stateList

	return list
}
