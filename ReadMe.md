
### 安装caliper并绑定sdk
```shell
cd fabric_wallet
npm init -y
npm install --only=prod \
    @hyperledger/caliper-cli@0.4.2
npx caliper bind \
    --caliper-bind-sut fabric:1.4  --unsafe-perm    
```


###caliper 测试
```shell
cd fabric_wallet

npx caliper launch manager \
    --caliper-workspace . \
    --caliper-benchconfig benchmarks/scenario/simple/config.yaml \
    --caliper-networkconfig networks/fabric/v1/v1.4.1/2org1peergoleveldb/fabric-go.yaml
```

###启动容器

```shell
cd fabric_wallet
cd fixtures
docker-compose up -d
```

###安装链码

```shell
cd initconfig
go mod init chaincode
go mod tidy
go mod vendor
### 命令行编译执行的话，需要将configFile的路径改下，改成../config.yaml
go build 
./initconfig
```


###启动项目
```shell
cd fabric_wallet
go build
./caliper-benchmarks
```

###停止容器，区块链网络环境
```shell
cd fabric_wallet
make clean
```