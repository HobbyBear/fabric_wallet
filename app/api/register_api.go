package api

import (
	"fabric_wallet/app/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterApi(e *gin.Engine, app controllers.App) {

	e.GET("/query/:wallet_id",app.Query)

	e.GET("/query_transaction/:transaction_id",app.QueryTransaction)

	e.POST("exec_transaction",app.ExecTransaction)

}
