package global

import (
	"github.com/xen0n/go-workwx"
	"log"
)

var TestWarnWx *workwx.WorkwxApp        //测试环境
var ProdLargeAmountWX *workwx.WorkwxApp //生产大金额

func InitWxAPP(corpID string) {

	client := workwx.New(corpID)

	//测试环境
	testInfo := WxEnv[WXENV_TEST]
	TestWarnWx = client.WithApp(testInfo.CorpSecret, testInfo.AgentId)
	if TestWarnWx == nil {
		log.Fatalf("ProdLargeAmountWX info error")
	}
	TestWarnWx.SpawnAccessTokenRefresher()

	//生产大金额告警
	prodLargeAmountInfo := WxEnv[WXENV_PROD_AMOUNT]
	ProdLargeAmountWX = client.WithApp(prodLargeAmountInfo.CorpSecret, prodLargeAmountInfo.AgentId)
	if ProdLargeAmountWX == nil {
		log.Fatalf("ProdLargeAmountWX info error")
	}
	ProdLargeAmountWX.SpawnAccessTokenRefresher()

}
