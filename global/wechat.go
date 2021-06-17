package global

var WxEnv map[string]*EnvInfo

const (
	WXENV_TEST        = "test"
	WXENV_PROD_AMOUNT = "prod_large_amount"
	WXENV_PROD_WARN   = "prod_large_warn"
)

//#1：全员, 2：运维, 3：测试, 4：大额告警, 5：正式环境交易告警
//to_party: '3'
//# 1000003：生产环境交易告警，1000004：测试环境交易告警
//agent_id: ' 1000004'
//推送关系绑定
type EnvInfo struct {
	ToParty    []string
	AgentId    int64
	CorpSecret string
}

func init() {
	WxEnv = make(map[string]*EnvInfo)

	//测试环境交易告警
	//WxEnv[WXENV_TEST] = &EnvInfo{
	//	ToParty:    []string{"82"},
	//	AgentId:    1000004,
	//	CorpSecret: "IGv_VTjdWTAlmgpGti0Z3spSb2zficU0a7ykoMmfM7U",
	//}

	//xXumGCGXNBuZxH-M1J61SyLYw9o-yCB5GBjStrNaWRg
	//1000014

	WxEnv[WXENV_TEST] = &EnvInfo{
		ToParty:    []string{"83"},
		AgentId:    1000016,
		CorpSecret: "BF5P3g0biTS_5qAKt5tGA4qz39EBK8gqWiLkQqF6ZXQ",
	}
	//生成大金额告警
	//Env["prod_large_amount"] = &EnvInfo{
	//	ToParty: "4",
	//	AgentId: 1000007,
	//	CorpSecret:"Hwrsx2cS-lDcEJOymw9P_8zzKdjH8SSpZ0DZ8SxkPoE"
	//}

	//WxEnv[WXENV_PROD_AMOUNT] = &EnvInfo{
	//	ToParty:    []string{"4"},
	//	AgentId:    1000008,
	//	CorpSecret: "aGTno289e_LxbvST2OzVKOAQn3cBg-V7muISV3jF_S8",
	//}

	//WxEnv[WXENV_PROD_AMOUNT] = &EnvInfo{
	//	ToParty:    []string{"83"},
	//	AgentId:    1000016,
	//	CorpSecret: "BF5P3g0biTS_5qAKt5tGA4qz39EBK8gqWiLkQqF6ZXQ",
	//}

	WxEnv[WXENV_PROD_AMOUNT] = &EnvInfo{
		ToParty:    []string{"82"},
		AgentId:    1000014,
		CorpSecret: "xXumGCGXNBuZxH-M1J61SyLYw9o-yCB5GBjStrNaWRg",
	}

	//生产环境交易告警
	WxEnv[WXENV_PROD_WARN] = &EnvInfo{
		ToParty:    []string{"82"},
		AgentId:    1000014,
		CorpSecret: "xXumGCGXNBuZxH-M1J61SyLYw9o-yCB5GBjStrNaWRg",
	}

}
