/**
  @Author : hanxiaodong
*/
package service

import (
	"fmt"
	paillier "github.com/Roasbeef/go-go-gadget-paillier"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"math/big"
	"time"
)

type Wallet struct {
	Id     string `json:"id"`
	Amount int64  `json:"amount"`
}

type Transaction struct {
	Id         string `json:"id"`
	CreateTime int64  `json:"create_time"`
	FromWallet string `json:"from_wallet"`
	ToWallet   string `json:"to_wallet"`
	Amount     int64  `json:"amount"`
}

type ServiceSetup struct {
	ChaincodeID string
	Client      *channel.Client
}

func regitserEvent(client *channel.Client, chaincodeID, eventID string) (fab.Registration, <-chan *fab.CCEvent) {

	reg, notifier, err := client.RegisterChaincodeEvent(chaincodeID, eventID)
	if err != nil {
		fmt.Println("注册链码事件失败: %s", err)
	}
	return reg, notifier
}

func eventResult(notifier <-chan *fab.CCEvent, eventID string) error {
	select {
	case ccEvent := <-notifier:
		fmt.Printf("接收到链码事件: %v\n", ccEvent)
	case <-time.After(time.Second * 20):
		return fmt.Errorf("不能根据指定的事件ID接收到相应的链码事件(%s)", eventID)
	}
	return nil
}

type WalletCC struct {
	Id     string `json:"id"`
	Amount []byte `json:"amount"`
}

type TransactionCC struct {
	Id         string `json:"id"`
	CreateTime int64  `json:"create_time"`
	FromWallet string `json:"from_wallet"`
	ToWallet   string `json:"to_wallet"`
	Amount     []byte `json:"amount"`
}

func Int64ToByte(num int64) []byte {
	return new(big.Int).SetInt64(num).Bytes()
}

func ByteToInt64(bytes []byte) int64 {
	return new(big.Int).SetBytes(bytes).Int64()
}

func Encry(num int64) []byte {
	data, _ := paillier.Encrypt(&PrivateKey.PublicKey, Int64ToByte(num))
	return data
}

func EncryString(account string) []byte {
	data, _ := paillier.Encrypt(&PrivateKey.PublicKey, []byte(account))
	return data
}

func Decry(bytes []byte) int64 {
	data, _ := paillier.Decrypt(PrivateKey, bytes)
	return ByteToInt64(data)
}

func DecryString(bytes []byte) string {
	data, _ := paillier.Decrypt(PrivateKey, bytes)
	return string(data)
}
