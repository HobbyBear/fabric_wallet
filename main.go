/**
  author: kevin
*/

package main

import (
	"caliper-benchmarks/app/api"
	"caliper-benchmarks/app/controllers"
	"caliper-benchmarks/app/service"
	"caliper-benchmarks/sdkInit"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
)

const (
	configFile  = "config.yaml"
	initialized = false
	EduCC       = "educc"
)

func main() {

	gopath := os.Getenv("GOPATH")

	initInfo := &sdkInit.InitInfo{

		ChannelID:     "kevinkongyixueyuan",
		ChannelConfig: gopath + "/src/fabric_source/fixtures/artifacts/channel.tx",

		OrgAdmin:       "Admin",
		OrgName:        "Org1",
		OrdererOrgName: "orderer.kevin.kongyixueyuan.com",

		ChaincodeID:     EduCC,
		ChaincodeGoPath: gopath,
		ChaincodePath:   "fabric_source/chaincode/",
		UserName:        "User1",
	}

	sdk, err := sdkInit.SetupSDK(configFile, initialized)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	defer sdk.Close()

	clientChannelContext := sdk.ChannelContext(initInfo.ChannelID, fabsdk.WithUser(initInfo.UserName), fabsdk.WithOrg(initInfo.OrgName))
	// returns a Client instance. Channel client can query chaincode, execute chaincode and register/unregister for chaincode events on specific channel.
	channelClient, err := channel.New(clientChannelContext)
	if err != nil {
		fmt.Printf("创建应用通道客户端失败: %v", err)
	}


	//===========================================//

	serviceSetup := service.ServiceSetup{
		ChaincodeID: EduCC,
		Client:      channelClient,
	}

	edu := service.Wallet{
		Id:    primitive.NewObjectID().Hex() ,
		Amount: 10,
	}

	edu2 := service.Wallet{
		Id:    primitive.NewObjectID().Hex() ,
		Amount: 10,
	}

	_, err = serviceSetup.SaveWallet(edu)
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = serviceSetup.SaveWallet(edu2)
	if err != nil {
		fmt.Println(err.Error())
	}

	//根据身份证号码查询信息
	result, err := serviceSetup.FindWalletInfoByEntityID(edu.Id)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("根据身份证号码查询信息成功：")
		fmt.Println(result)
	}

	// 根据身份证号码查询信息
	result, err = serviceSetup.FindWalletInfoByEntityID(edu2.Id)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("根据身份证号码查询信息成功：")
		fmt.Println(result)
	}

	//===========================================//

	app := controllers.App{
		Setup: &serviceSetup,
	}
	e := gin.Default()
	api.RegisterApi(e, app)
	e.Run(":8080")

}
