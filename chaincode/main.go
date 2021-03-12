/**
  @Author : hanxiaodong
*/

package main

import (
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
)

type EducationChaincode struct {
}

func (t *EducationChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {

	return shim.Success(nil)
}

func (t *EducationChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	// 获取用户意图
	fun, args := stub.GetFunctionAndParameters()

	if fun == "addWallet" {
		return t.addWallet(stub, args) // 添加信息
	}  else if fun == "queryWalletInfoByWalletID" {
		return t.queryWalletInfoByWalletID(stub, args) // 根据身份证号码及姓名查询详情
	} else if fun == "updateWallet" {
		return t.updateWallet(stub, args) // 根据证书编号更新信息
	}else if fun == "addTransaction" {
		return t.addTransaction(stub, args) // 根据证书编号更新信息
	}else if fun == "queryTransactionInfoByTransactionID" {
		return t.queryTransactionInfoByTransactionID(stub, args) // 根据证书编号更新信息
	}
	return shim.Error("指定的函数名称错误")

}

func main() {
	err := shim.Start(new(EducationChaincode))
	if err != nil {
		fmt.Printf("启动EducationChaincode时发生错误: %s", err)
	}
}
