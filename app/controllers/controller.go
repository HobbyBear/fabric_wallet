package controllers

import (
	"caliper-benchmarks/app/service"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

type App struct {
	Setup *service.ServiceSetup
}


func (a *App) Query(c *gin.Context) {

	walletId := c.Param("wallet_id")

	result, err := a.Setup.FindWalletInfoByEntityID(walletId)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":"success",
		"data":result,
	})
}

func (a *App) ExecTransaction(c *gin.Context) {
	var reqBody struct {
		FromWallet string  `json:"from_wallet"`
		ToWallet   string  `json:"to_wallet"`
		Account    int64 `json:"account"`
	}

	err := c.BindJSON(&reqBody)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "参数错误",
		})
		return
	}
	var resp struct {
		TransactionId string `json:"transaction_id"`
	}
	resp.TransactionId = primitive.NewObjectID().Hex()

	fromWallet, err := a.Setup.FindWalletInfoByEntityID(reqBody.FromWallet)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "参数错误",
		})
		return
	}

	toWallet, err := a.Setup.FindWalletInfoByEntityID(reqBody.ToWallet)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "参数错误",
		})
		return
	}

	if fromWallet.Amount-reqBody.Account < 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  fmt.Sprintf("%s的账户余额不足", reqBody.FromWallet),
		})
		return
	}

	_, err = a.Setup.SaveTransaction(service.Transaction{
		Id:         resp.TransactionId,
		CreateTime: time.Now().Unix(),
		FromWallet: reqBody.FromWallet,
		ToWallet:   reqBody.ToWallet,
		Amount:     reqBody.Account,
	})

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "执行交易失败",
		})
		return
	}

	fromWallet.Amount -= reqBody.Account

	toWallet.Amount += reqBody.Account

	a.Setup.ModifyEdu(*fromWallet)

	a.Setup.ModifyEdu(*toWallet)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "交易成功",
		"data": resp.TransactionId,
	})

}
func (a *App) QueryTransaction(c *gin.Context) {

	transactionId := c.Param("transaction_id")

	data, err := a.Setup.FindTransactionInfoByEntityID(transactionId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "参数错误",
		})
		return
	}

	var resp service.TransactionCC
	json.Unmarshal(data, &resp)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": service.Transaction{
			Id:         resp.Id,
			CreateTime: resp.CreateTime,
			FromWallet: resp.FromWallet,
			ToWallet:   resp.ToWallet,
			Amount:     service.Decry(resp.Amount),
		},
	})

}
