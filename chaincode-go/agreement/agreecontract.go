package agreement

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type Contract struct {
	contractapi.Contract
}

// Instantiate does nothing
func (c *Contract) Instantiate() {
	fmt.Println("Instantiated")
}

type SignerInfo struct {
	Id        string `json:"userId"`
	PublicPem string `json:"publicPem"`
}

// 提交新合同信息并存储到世界状态
func (c *Contract) Issue(ctx TransactionContextInterface, issuerOrg string, issuerId string, agreementNumber string, issueDateTime string, signerList string) (*Agreement, error) {
	// 构造map
	var sl []SignerInfo
	if err := json.Unmarshal([]byte(signerList), &sl); err != nil {
		return nil, err
	}

	mapList := map[string]*Signatory{}
	for _, v := range sl {
		s := Signatory{Id: v.Id, PublicPem: v.PublicPem, Result: false}
		mapList[v.Id] = &s
	}

	agreement := Agreement{AgreementNumber: agreementNumber, IssuerId: issuerId, IssuerOrg: issuerOrg, IssueDateTime: issueDateTime, SignerMapList: mapList, SignData: []SignData{}}
	agreement.SetIssued()

	err := ctx.GetAgreementList().AddAgreement(&agreement)

	if err != nil {
		return nil, err
	}

	return &agreement, nil
}

// 签署方签署合同,记录签署信息并更新世界状态
func (c *Contract) Signature(ctx TransactionContextInterface, issuerOrg string, issuerId string, agreementNumber string, signData string) (*Agreement, error) {
	sd := SignData{}
	if err := json.Unmarshal([]byte(signData), &sd); err != nil {
		return nil, err
	}

	agreement, err := ctx.GetAgreementList().GetAgreement(issuerOrg, issuerId, agreementNumber)

	if err != nil {
		return nil, err
	}

	// 签署合同
	if _, ok := agreement.SignerMapList[sd.Id]; !ok {
		return nil, fmt.Errorf("error Signature agreement. The Signer isn't correct")
	}

	// 校验签名
	publicPem := agreement.SignerMapList[sd.Id].PublicPem
	block, _ := pem.Decode([]byte(publicPem))
	//Step3:将公匙反x509序列化
	pubInterface, _ := x509.ParsePKIXPublicKey(block.Bytes)
	//Step4:执行公匙的类型断言
	publicKey, ok := pubInterface.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("error Signature agreement. The publicPem isn't correct")
	}

	if valid := ecdsa.VerifyASN1(publicKey, []byte(sd.DataHash), []byte(sd.SignDataHash)); !valid {
		return nil, fmt.Errorf("error Signature agreement. The Signature VerifyASN1 Failed")
	}

	agreement.SignerMapList[sd.Id].Result = true
	agreement.SetTrading()

	agreement.SignData = append(agreement.SignData, sd) //保存签署信息

	err = ctx.GetAgreementList().UpdateAgreement(agreement)

	if err != nil {
		return nil, err
	}

	return agreement, nil
}

// 第三方确认（政府或公证处）确认合同已签署完毕，更新世界状态
func (c *Contract) Confirm(ctx TransactionContextInterface, issuerOrg string, issuerId string, agreementNumber string) (*Agreement, error) {
	agreement, err := ctx.GetAgreementList().GetAgreement(issuerOrg, issuerId, agreementNumber)

	if err != nil {
		return nil, err
	}
	// 检查所以签署人是否签署完毕
	for _, v := range agreement.SignerMapList {
		if !v.Result {
			return nil, fmt.Errorf("error Confirm agreement. The Signer Id %s hadn't signed", v.Id)
		}
	}

	agreement.SetFinish() // 确认合同签署完成

	err = ctx.GetAgreementList().UpdateAgreement(agreement)

	if err != nil {
		return nil, err
	}
	return agreement, nil
}

// 判读合同是否存在
func (s *Contract) AgreementExists(ctx TransactionContextInterface, issuerOrg string, issuerId string, agreementNumber string) (bool, error) {

	agreement, err := ctx.GetAgreementList().GetAgreement(issuerOrg, issuerId, agreementNumber)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return agreement != nil, nil
}

// 读取合同信息
func (s *Contract) GetAgreement(ctx TransactionContextInterface, issuerOrg string, issuerId string, agreementNumber string) (*Agreement, error) {

	agreement, err := ctx.GetAgreementList().GetAgreement(issuerOrg, issuerId, agreementNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}

	return agreement, nil
}

// 复合键不可以用Range查询,该接口不可用
func (s *Contract) GetAllAgreement(ctx TransactionContextInterface) ([]*Agreement, error) {

	agreements, err := ctx.GetAgreementList().GetAllAgreement()
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}

	return agreements, nil
}

// 复合键不可以用Range查询,该接口不可用
func (s *Contract) GetAgreementByRange(ctx TransactionContextInterface, startKey, endKey string) ([]*Agreement, error) {
	agreements, err := ctx.GetAgreementList().GetAgreementByRange(startKey, endKey)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}

	return agreements, nil
}

// 通过复合键（组织信息）查询该组织提交的合同信息
func (s *Contract) GetAgreementsByOrg(ctx TransactionContextInterface, issuerOrg string) ([]*Agreement, error) {
	agreements, err := ctx.GetAgreementList().GetAgreementByOrg(issuerOrg)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}

	return agreements, nil
}

// 通过复合键（组织信息和提交人ID）查询该组织提交的合同信息
func (s *Contract) GetAgreementsByOrgAndId(ctx TransactionContextInterface, issuerOrg, issuerId string) ([]*Agreement, error) {
	agreements, err := ctx.GetAgreementList().GetAgreementByOrgAndId(issuerOrg, issuerId)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}

	return agreements, nil
}
