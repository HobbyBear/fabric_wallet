/**
  @Author : hanxiaodong
*/

package service

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	paillier "github.com/Roasbeef/go-go-gadget-paillier"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)


var (
	PrivateKey *paillier.PrivateKey
)

func init() {
	PrivateKey, _ = paillier.GenerateKey(rand.Reader, 128)
}

func (t *ServiceSetup) SaveWallet(wallet Wallet) (string, error) {

	eventID := "eventAddWallet"
	reg, notifier := regitserEvent(t.Client, t.ChaincodeID, eventID)
	defer t.Client.UnregisterChaincodeEvent(reg)

	// 将edu对象序列化成为字节数组
	b, err := json.Marshal(WalletCC{
		Id:     wallet.Id,
		Amount: Encry(wallet.Amount),
	})
	if err != nil {
		return "", fmt.Errorf("指定的edu对象序列化时发生错误")
	}

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "addWallet", Args: [][]byte{b, []byte(eventID)}}
	respone, err := t.Client.Execute(req)
	if err != nil {
		return "", err
	}

	err = eventResult(notifier, eventID)
	if err != nil {
		return "", err
	}

	return string(respone.TransactionID), nil
}

func (t *ServiceSetup) FindWalletInfoByEntityID(entityID string) (*Wallet, error) {

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "queryWalletInfoByWalletID", Args: [][]byte{[]byte(entityID)}}
	respone, err := t.Client.Query(req)
	if err != nil {
		return nil, err
	}
	var walletCC WalletCC
	json.Unmarshal(respone.Payload, &walletCC)

	return &Wallet{
		Id:     walletCC.Id,
		Amount: Decry(walletCC.Amount),
	}, nil
}

func (t *ServiceSetup) ModifyEdu(edu Wallet) (string, error) {

	eventID := "eventModifyEdu"
	reg, notifier := regitserEvent(t.Client, t.ChaincodeID, eventID)
	defer t.Client.UnregisterChaincodeEvent(reg)

	// 将edu对象序列化成为字节数组
	b, err := json.Marshal(WalletCC{
		Id:     edu.Id,
		Amount: Encry(edu.Amount),
	})
	if err != nil {
		return "", fmt.Errorf("指定的edu对象序列化时发生错误")
	}

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "updateWallet", Args: [][]byte{b, []byte(eventID)}}
	respone, err := t.Client.Execute(req)
	if err != nil {
		return "", err
	}

	err = eventResult(notifier, eventID)
	if err != nil {
		return "", err
	}

	return string(respone.TransactionID), nil
}

func (t *ServiceSetup) SaveTransaction(transaction Transaction) (string, error) {

	eventID := "eventAddTransaction"
	reg, notifier := regitserEvent(t.Client, t.ChaincodeID, eventID)
	defer t.Client.UnregisterChaincodeEvent(reg)

	// 将edu对象序列化成为字节数组
	encryAmount := Encry(transaction.Amount)
	b, err := json.Marshal(TransactionCC{
		Id:         transaction.Id,
		CreateTime: transaction.CreateTime,
		FromWallet: transaction.FromWallet,
		ToWallet:   transaction.ToWallet,
		Amount:     encryAmount,
	})
	if err != nil {
		return "", fmt.Errorf("指定的edu对象序列化时发生错误")
	}

	fmt.Println(fmt.Sprintf("交易:%s,交易的加密金额为：%s", transaction.Id, encryAmount))

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "addTransaction", Args: [][]byte{b, []byte(eventID)}}
	respone, err := t.Client.Execute(req)
	if err != nil {
		return "", err
	}

	err = eventResult(notifier, eventID)
	if err != nil {
		return "", err
	}

	return string(respone.TransactionID), nil
}

func (t *ServiceSetup) FindTransactionInfoByEntityID(entityID string) ([]byte, error) {

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "queryTransactionInfoByTransactionID", Args: [][]byte{[]byte(entityID)}}
	respone, err := t.Client.Query(req)
	if err != nil {
		return []byte{0x00}, err
	}

	return respone.Payload, nil
}
