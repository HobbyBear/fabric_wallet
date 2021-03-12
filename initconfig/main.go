package main

import (
	"fabric_wallet/sdkInit"
	"fmt"
	"os"
)

func main() {

	const (
		configFile  = "./config.yaml"
		initialized = false
		EduCC       = "educc"
	)

	// 获取当前目录，然后获取路径的

	gopath := os.Getenv("GOPATH")

	initInfo := &sdkInit.InitInfo{

		ChannelID:     "kevinkongyixueyuan",
		ChannelConfig: gopath + "/src/fabric_wallet/fixtures/artifacts/channel.tx",

		OrgAdmin:       "Admin",
		OrgName:        "Org1",
		OrdererOrgName: "orderer.kevin.kongyixueyuan.com",

		ChaincodeID:     EduCC,
		ChaincodeGoPath: gopath,
		ChaincodePath:   "fabric_wallet/chaincode/",
		UserName:        "User1",
	}

	sdk, err := sdkInit.SetupSDK(configFile, initialized)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	err = sdkInit.CreateChannel(sdk, initInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	channelClient, err := sdkInit.InstallAndInstantiateCC(sdk, initInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(channelClient)

}
