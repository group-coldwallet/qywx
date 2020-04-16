package model

import (
	"github.com/group-coldwallet/qywx/pkg/util"
)

type Config struct {
	Env        string     `toml:"env"`        //开发环境
	Http       HttpConfig `toml:"http"`       //http服务器配置
	Encryption bool       `toml:"encryption"` //加密启用
	Wechat     WechatQy   `toml:"wx"`
}

type HttpConfig struct {
	Port string `toml:"port"` //服务器端口
}

type WechatQy struct {
	Corpid string `toml:"corpid"`
	//Secret string `toml:"secret"`
}

func (c *Config) New() {
	c = &Config{
		Env: "prod",
		Http: HttpConfig{
			Port: "8080",
		},
		Encryption: false,
	}
}

func (c *Config) Decrypt() {
	aesKey := "cck+Bms+ec2w1PzwP8J2FngKa4EB/fPl"
	wxCorpid, _ := util.AesBase64Crypt([]byte(c.Wechat.Corpid), []byte(aesKey), false)
	//wxSecret, _ := util.AesBase64Crypt([]byte(c.Wechat.Secret), []byte(aesKey), false)
	c.Wechat.Corpid = string(wxCorpid)
	//c.Wechat.Secret = string(wxSecret)
}
