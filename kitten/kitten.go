// Package kitten 包含了 KittenCore 以及各插件的核心依赖结构体、方法和函数
package kitten

import (
	"fmt"
	"net/url"

	"github.com/Kittengarten/KittenCore/kitten/core"

	"gopkg.in/yaml.v3"
)

const configFile = `config.yaml`

var (
	botConfig config    // 来自 Bot 的配置文件
	imagePath core.Path // 图片路径
	Weight    int       // 自身叠猫猫体重（0.1kg 数）
)

func init() {
	// 配置文件加载
	d, err := core.Path(configFile).Read()
	if nil != err {
		fmt.Println(err, `请配置`, configFile, `后重新启动喵！`)
	}
	if err := yaml.Unmarshal(d, &botConfig); nil != err {
		fmt.Println(err, `请按 YAML 格式配置`, configFile, `后重新启动喵！`)
	}
	if 0 == len(botConfig.SuperUsers) {
		fmt.Println(`请在`, configFile, `中配置 superusers 喵！`)
	}
	if 0 == len(botConfig.NickName) {
		botConfig.NickName = []string{`喵喵`}
	}
	if _, err := url.Parse(botConfig.WebSocket.URL); nil != err {
		fmt.Println(err, `请正确配置 `, configFile, ` 中的 websocket.url 喵！`)
	}
	if _, err := url.Parse(botConfig.WebUI.URL); nil != err {
		fmt.Println(err, `请正确配置`, configFile, `中的 webui.url 喵！`)
	}
	// 图片路径
	imagePath = core.FilePath(botConfig.Path, `image`)
}

// MainConfig 获取主配置
func MainConfig() config {
	return botConfig
}
