package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/group-coldwallet/qywx/global"
	"github.com/group-coldwallet/qywx/model/bo"
	"github.com/group-coldwallet/qywx/pkg/httpresp"
	"github.com/xen0n/go-workwx"
	"log"
)

var toProd *workwx.Recipient

func init() {
	env := global.WxEnv[global.WXENV_PROD_AMOUNT]
	toProd = &workwx.Recipient{
		PartyIDs: env.ToParty,
	}
}

//生产环境大金额高进
func SendMsgToProdAmount(c *gin.Context) {
	wxTx := new(bo.WxContent)
	c.BindJSON(wxTx)
	if wxTx.Content == "" {
		httpresp.HttpRespErrorOnly(c)
		return
	}

	log.Println(global.ProdLargeAmountWX == nil)
	log.Println(toProd == nil)
	err := global.ProdLargeAmountWX.SendTextMessage(toProd, wxTx.Content, false)
	if err != nil {
		log.Println(fmt.Sprintf("发送错误：%s", err.Error()))
		httpresp.HttpRespErrorOnly(c)
		return
	}
	httpresp.HttpRespOkOnly(c)
}
