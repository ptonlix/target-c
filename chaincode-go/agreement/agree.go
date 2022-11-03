package agreement

import (
	"encoding/json"
	"fmt"

	ledgerapi "github.com/ptonlix/target-c/chaincode-go/ledger-api"
)

// 状态标志
type State uint

const (
	// 合同创建成功
	ISSUED State = iota + 1
	// 合同正在签署
	TRADING
	// 合同签署完成
	FINISH
)

func (state State) String() string {
	names := []string{"ISSUED", "TRADING", "FINISH"}

	if state < ISSUED || state > FINISH {
		return "UNKNOWN"
	}

	return names[state-1]
}

// CreateAgreementKey creates a key for agreement
func CreateAgreementKey(issuerOrg string, issuerId string, agreementNumber string) string {
	return ledgerapi.MakeKey(issuerOrg, issuerId, agreementNumber)
}

// Used for managing the fact status is private but want it in world state
type agreementAlias Agreement
type jsonAgreement struct {
	*agreementAlias
	State State  `json:"currentState"`
	Class string `json:"class"`
	Key   string `json:"key"`
}

type Signatory struct {
	Id        string `json:"userId"`
	Result    bool   `json:"result"`
	PublicPem string `json:"publicPem"`
}

type SignData struct {
	Id           string `json:"userId"`
	DataHash     string `json:"dataHash"`
	SignDataHash string `json:"signDataHash"`
}

// CommercialPaper defines a commercial paper
type Agreement struct {
	AgreementNumber string                `json:"agreementNumber"`
	IssuerOrg       string                `json:"issuerOrg"`
	IssuerId        string                `json:"issuerId"`
	IssueDateTime   string                `json:"issueDateTime"`
	SignerMapList   map[string]*Signatory `json:"signerList"`
	SignData        []SignData            `json:"signData"`
	state           State                 `metadata:"currentState"`
	class           string                `metadata:"class"`
	key             string                `metadata:"key"`
}

// UnmarshalJSON special handler for managing JSON marshalling
func (cp *Agreement) UnmarshalJSON(data []byte) error {
	jcp := jsonAgreement{agreementAlias: (*agreementAlias)(cp)}

	err := json.Unmarshal(data, &jcp)

	if err != nil {
		return err
	}

	cp.state = jcp.State

	return nil
}

// MarshalJSON special handler for managing JSON marshalling
func (cp Agreement) MarshalJSON() ([]byte, error) {
	jcp := jsonAgreement{agreementAlias: (*agreementAlias)(&cp), State: cp.state, Class: "org.targetc-net.agreement",
		Key: ledgerapi.MakeKey(cp.IssuerOrg, cp.IssuerId, cp.AgreementNumber)}

	return json.Marshal(&jcp)
}

// GetState returns the state
func (cp *Agreement) GetState() State {
	return cp.state
}

// SetIssued returns the state to issued
func (cp *Agreement) SetIssued() {
	cp.state = ISSUED
}

// SetTrading sets the state to trading
func (cp *Agreement) SetTrading() {
	cp.state = TRADING
}

// SetFinish sets the state to Finish
func (cp *Agreement) SetFinish() {
	cp.state = FINISH
}

// IsIssued returns true if state is issued
func (cp *Agreement) IsIssued() bool {
	return cp.state == ISSUED
}

// IsTrading returns true if state is trading
func (cp *Agreement) IsTrading() bool {
	return cp.state == TRADING
}

// IsRedeemed returns true if state is redeemed
func (cp *Agreement) IsFinish() bool {
	return cp.state == FINISH
}

// GetSplitKey returns values which should be used to form key
func (cp *Agreement) GetSplitKey() []string {
	return []string{cp.IssuerOrg, cp.IssuerId, cp.AgreementNumber}
}

// Serialize formats the agreement as JSON bytes
func (cp *Agreement) Serialize() ([]byte, error) {
	return json.Marshal(cp)
}

// Deserialize formats the agreement from JSON bytes
func Deserialize(bytes []byte, cp *Agreement) error {
	err := json.Unmarshal(bytes, cp)

	if err != nil {
		return fmt.Errorf("error deserializing agreement. %s", err.Error())
	}

	return nil
}
