package global

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/group-coldwallet/qywx/model"
	"log"
	"os"
)

var Cfg *model.Config

func InitConfig(configFile string) {
	Cfg = new(model.Config)
	if configFile == "" {
		//默认读取上一层conf文件夹 配置文件
		dir, err := os.Getwd()
		if err != nil {
			log.Fatalf("laod config getpwd error:%s", err.Error())
		}
		configFile = fmt.Sprintf("%s/conf/%s.toml", dir, "application")
	}
	if _, err := toml.DecodeFile(configFile, Cfg); err != nil {
		log.Fatalf("laod config decode:%s error:%s", configFile, err.Error())
	}

	//解密参数
	if Cfg.Encryption {
		Cfg.Decrypt()
	}
}
