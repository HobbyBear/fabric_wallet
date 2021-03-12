package main

import (
	"encoding/json"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
)

func putWallet(stub shim.ChaincodeStubInterface, wallet  WalletCC) ([]byte, bool) {

	b, err := json.Marshal(wallet)
	if err != nil {
		return nil, false
	}

	err = stub.PutState(wallet.Id, b)
	if err != nil {
		return nil, false
	}

	return b, true
}

func putTransaction(stub shim.ChaincodeStubInterface, trans TransactionCC) ([]byte, bool) {

	b, err := json.Marshal(trans)
	if err != nil {
		return nil, false
	}

	err = stub.PutState(trans.Id, b)
	if err != nil {
		return nil, false
	}

	return b, true
}

func (t *EducationChaincode) addTransaction(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 2 {
		return shim.Error("给定的参数个数不符合要求")
	}

	var edu TransactionCC
	err := json.Unmarshal([]byte(args[0]), &edu)
	if err != nil {
		return shim.Error("反序列化信息时发生错误")
	}

	// 查重: 结算序号必须唯一
	_, exist := GetTransactionInfo(stub, edu.Id)
	if exist {
		return shim.Error("要添加的身份证号码已存在")
	}

	_, bl := putTransaction(stub, edu)
	if !bl {
		return shim.Error("保存信息时发生错误")
	}

	err = stub.SetEvent(args[1], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("信息添加成功"))
}

// 根据身份证号码查询信息状态
// args: entityID
func GetWalletInfo(stub shim.ChaincodeStubInterface, walletId string) (WalletCC, bool) {
	var edu WalletCC
	b, err := stub.GetState(walletId)
	if err != nil {
		return edu, false
	}

	if b == nil {
		return edu, false
	}

	// 对查询到的状态进行反序列
	err = json.Unmarshal(b, &edu)
	if err != nil {
		return edu, false
	}

	// 返回结果
	return edu, true
}

// 查询交易信息
// args: entityID
func GetTransactionInfo(stub shim.ChaincodeStubInterface, transactionId string) (TransactionCC, bool) {
	var edu TransactionCC
	b, err := stub.GetState(transactionId)
	if err != nil {
		return edu, false
	}

	if b == nil {
		return edu, false
	}

	// 对查询到的状态进行反序列化
	err = json.Unmarshal(b, &edu)
	if err != nil {
		return edu, false
	}

	// 返回结果
	return edu, true
}

func (t *EducationChaincode) addWallet(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 2 {
		return shim.Error("给定的参数个数不符合要求")
	}

	var edu WalletCC
	err := json.Unmarshal([]byte(args[0]), &edu)
	if err != nil {
		return shim.Error("反序列化信息时发生错误")
	}

	// 查重: 结算序号必须唯一
	_, exist := GetWalletInfo(stub, edu.Id)
	if exist {
		return shim.Error("要添加的身份证号码已存在")
	}

	_, bl := putWallet(stub, edu)
	if !bl {
		return shim.Error("保存信息时发生错误")
	}

	err = stub.SetEvent(args[1], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("信息添加成功"))
}

// 根据身份证号码查询详情（溯源）
// args: entityID
func (t *EducationChaincode) queryWalletInfoByWalletID(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("给定的参数个数不符合要求")
	}

	// 根据身份证号码查询edu状态
	b, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("根据身份证号码查询信息失败")
	}

	if b == nil {
		return shim.Error("根据身份证号码没有查询到相关的信息")
	}

	// 对查询到的状态进行反序列化
	var walletCC WalletCC
	err = json.Unmarshal(b, &walletCC)
	if err != nil {
		return shim.Error("反序列化edu信息失败")
	}

	// 返回
	result, err := json.Marshal(walletCC)
	if err != nil {
		return shim.Error("序列化edu信息时发生错误")
	}
	return shim.Success(result)
}

func (t *EducationChaincode) queryTransactionInfoByTransactionID(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("给定的参数个数不符合要求")
	}

	// 根据身份证号码查询edu状态
	b, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("根据身份证号码查询信息失败")
	}

	if b == nil {
		return shim.Error("根据身份证号码没有查询到相关的信息")
	}

	// 对查询到的状态进行反序列化
	var transactionCC TransactionCC
	err = json.Unmarshal(b, &transactionCC)
	if err != nil {
		return shim.Error("反序列化edu信息失败")
	}

	// 返回
	result, err := json.Marshal(transactionCC)
	if err != nil {
		return shim.Error("序列化edu信息时发生错误")
	}
	return shim.Success(result)
}

// 根据身份证号更新信息
// args: educationObject
func (t *EducationChaincode) updateWallet(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2 {
		return shim.Error("给定的参数个数不符合要求")
	}

	var info WalletCC
	err := json.Unmarshal([]byte(args[0]), &info)
	if err != nil {
		return shim.Error("反序列化edu信息失败")
	}

	// 根据身份证号码查询信息
	result, bl := GetWalletInfo(stub, info.Id)
	if !bl {
		return shim.Error("根据身份证号码查询信息时发生错误")
	}

	result = WalletCC{
		Id:     info.Id,
		Amount: info.Amount,
	}

	_, bl = putWallet(stub, result)
	if !bl {
		return shim.Error("保存信息信息时发生错误")
	}

	err = stub.SetEvent(args[1], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("信息更新成功"))
}
