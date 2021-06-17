package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/group-coldwallet/qywx/controller/v1"
	"github.com/group-coldwallet/qywx/global"
)

func InitV1Router(r *gin.Engine) {
	v1Router := r.Group("/v1")
	{

		v1Wx := v1Router.Group("/wx")

		if global.Cfg.Env != "prod" {
			//测试环境
			v1TestRouter := v1Wx.Group("/test")
			{
				v1TestRouter.POST("/text", v1.SendMsgToTest)
			}
		} else {
			//生产环境
			v1ProdRouter := v1Wx.Group("/prod")
			{
				v1ProdAmountRouter := v1ProdRouter.Group("/amount")
				{
					v1ProdAmountRouter.POST("/text", v1.SendMsgToProdAmount)
				}

			}
		}

	}

}
