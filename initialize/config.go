package initialize

import (
	"github.com/CodingJzy/library_backend/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"strings"
)

var v *viper.Viper

func LoadConfig() {
	v = viper.New()
	initConfig()
}

func initConfig() {
	// 设置配置文件
	v.SetConfigFile("./config.yaml")

	// 配置环境变量
	v.SetEnvPrefix("JW")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	// 加载配置
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("配置文件加载错误：%v\n", err)
	}

	if err := v.Unmarshal(&global.Config); err != nil {
		log.Fatalf("配置反序列化失败：%v\n", err)
	}

	// 监听配置
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		if err := v.Unmarshal(&global.Config); err != nil {
			log.Fatalf("配置反序列化失败：%v\n", err)
		}
	})
	log.Println("load config success")
}
