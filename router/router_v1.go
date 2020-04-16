package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/group-coldwallet/qywx/controller/v1"
)

func InitV1Router(r *gin.Engine) {
	v1Router := r.Group("/v1")
	{

		v1Wx := v1Router.Group("/wx")
		v1TestRouter := v1Wx.Group("/test")
		{
			v1TestRouter.POST("/text", v1.SendMsgToTest)
		}

		v1ProdRouter := v1Wx.Group("/prod")
		{
			v1ProdAmountRouter := v1ProdRouter.Group("/amount")
			{
				v1ProdAmountRouter.POST("/text", v1.SendMsgToProdAmount)
			}

		}
	}

}
