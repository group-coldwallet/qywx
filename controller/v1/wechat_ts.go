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

var toTest *workwx.Recipient

func init() {
	env := global.WxEnv[global.WXENV_TEST]
	toTest = &workwx.Recipient{
		PartyIDs: env.ToParty,
	}
}

//测试环境
func SendMsgToTest(c *gin.Context) {
	wxTx := new(bo.WxContent)
	c.BindJSON(wxTx)
	if wxTx.Content == "" {
		httpresp.HttpRespErrorOnly(c)
		return
	}
	err := global.TestWarnWx.SendTextMessage(toTest, wxTx.Content, false)
	if err != nil {
		log.Println(fmt.Sprintf("发送错误：%s", err.Error()))
		httpresp.HttpRespErrorOnly(c)
		return
	}
	httpresp.HttpRespOkOnly(c)
}
