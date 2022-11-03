package agreement

import (
	"testing"

	ledgerapi "github.com/ptonlix/target-c/chaincode-go/ledger-api"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	assert.Equal(t, "ISSUED", ISSUED.String(), "should return string for issued")
	assert.Equal(t, "TRADING", TRADING.String(), "should return string for issued")
	assert.Equal(t, "FINISH", FINISH.String(), "should return string for issued")
	assert.Equal(t, "UNKNOWN", State(FINISH+1).String(), "should return unknown when not one of constants")
}

func TestCreateCommercialPaperKey(t *testing.T) {
	assert.Equal(t, ledgerapi.MakeKey("100", "someissuer", "somepaper"), CreateAgreementKey("100", "someissuer", "somepaper"), "should return key comprised of passed values")
}

func TestGetState(t *testing.T) {
	cp := new(Agreement)
	cp.state = ISSUED

	assert.Equal(t, ISSUED, cp.GetState(), "should return set state")
}

func TestSetIssued(t *testing.T) {
	cp := new(Agreement)
	cp.SetIssued()
	assert.Equal(t, ISSUED, cp.state, "should set state to issued")
}

func TestSetTrading(t *testing.T) {
	cp := new(Agreement)
	cp.SetTrading()
	assert.Equal(t, TRADING, cp.state, "should set state to trading")
}

func TestSetFinish(t *testing.T) {
	cp := new(Agreement)
	cp.SetFinish()
	assert.Equal(t, FINISH, cp.state, "should set state to finish")
}

func TestIsIssued(t *testing.T) {
	cp := new(Agreement)

	cp.SetIssued()

	assert.True(t, cp.IsIssued(), "should be true when status set to issued")

	cp.SetTrading()
	assert.False(t, cp.IsIssued(), "should be false when status not set to issued")
}

func TestIsTrading(t *testing.T) {
	cp := new(Agreement)

	cp.SetTrading()
	assert.True(t, cp.IsTrading(), "should be true when status set to trading")

	cp.SetFinish()
	assert.False(t, cp.IsTrading(), "should be false when status not set to trading")
}

func TestIsFinish(t *testing.T) {
	cp := new(Agreement)

	cp.SetFinish()
	assert.True(t, cp.IsFinish(), "should be true when status set to trading")

	cp.SetIssued()
	assert.False(t, cp.IsFinish(), "should be false when status not set to trading")
}

func TestGetSplitKey(t *testing.T) {
	cp := new(Agreement)
	cp.IssuerId = "100"
	cp.IssuerOrg = "someissuer"
	cp.AgreementNumber = "somepaper"

	assert.Equal(t, []string{"someissuer", "100", "somepaper"}, cp.GetSplitKey(), "should return issuer and agreement number as split key")
}

func TestSerialize(t *testing.T) {
	cp := new(Agreement)
	cp.AgreementNumber = "somepaper"
	cp.IssuerId = "100"
	cp.IssuerOrg = "someissuer"
	cp.IssueDateTime = "sometime"
	cp.SignerMapList = make(map[string]*Signatory)
	cp.SignerMapList["100"] = &Signatory{Id: "100", Result: false, PublicPem: "123456"}
	cp.SignData = []SignData{{Id: "100", DataHash: "asdqwe", SignDataHash: "qweasd"}}
	cp.state = TRADING

	bytes, err := cp.Serialize()
	//fmt.Println(string(bytes))
	assert.Nil(t, err, "should not error on serialize")
	jsonstr := `{"agreementNumber":"somepaper","issuerOrg":"someissuer","issuerId":"100","issueDateTime":"sometime","signerList":{"100":{"userId":"100","result":false,"publicPem":"123456"}},"signData":[{"userId":"100","dataHash":"asdqwe","signDataHash":"qweasd"}],"currentState":2,"class":"org.targetc-net.agreement","key":"someissuer:100:somepaper"}`
	assert.Equal(t, jsonstr, string(bytes), "should return JSON formatted value")
}

func TestDeserialize(t *testing.T) {
	var cp *Agreement
	var err error

	goodJSON := `{"agreementNumber":"somepaper","issuerOrg":"someissuer","issuerId":"100","issueDateTime":"sometime","signerList":{"100":{"userId":"100","result":false,"publicPem":"123456"}},"signData":[{"userId":"100","dataHash":"asdqwe","signDataHash":"qweasd"}],"currentState":2,"class":"org.targetc-net.agreement","key":"someissuer:100:somepaper"}`
	expectedCp := new(Agreement)
	expectedCp.AgreementNumber = "somepaper"
	expectedCp.IssuerId = "100"
	expectedCp.IssuerOrg = "someissuer"
	expectedCp.IssueDateTime = "sometime"
	expectedCp.SignerMapList = make(map[string]*Signatory)
	expectedCp.SignerMapList["100"] = &Signatory{Id: "100", Result: false, PublicPem: "123456"}
	expectedCp.SignData = []SignData{{Id: "100", DataHash: "asdqwe", SignDataHash: "qweasd"}}
	expectedCp.state = TRADING
	cp = new(Agreement)
	err = Deserialize([]byte(goodJSON), cp)
	assert.Nil(t, err, "should not return error for deserialize")
	assert.Equal(t, expectedCp, cp, "should create expected agreement")

	badJSON := `{"agreementNumber":"somepaper","issuerId":"200","issuerOrg":"someissuer","issueDateTime":"sometime","signerList":{"100":{"userId":"100","result":false,"publicPem":"123456"}},"signData":[{"userId":"100","dataHash":"asdqwe","signDataHash":"qweasd"}],"currentState":"abc","class":"org.targetc-net.agreement","key":"100:someissuer:somepaper"}`
	cp = new(Agreement)
	err = Deserialize([]byte(badJSON), cp)
	assert.EqualError(t, err, "error deserializing agreement. json: cannot unmarshal string into Go struct field jsonAgreement.currentState of type agreement.State", "should return error for bad data")
}
