// KittenCore 的主函数所在包
package main

import (
	"bytes"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/FloatTech/zbputils/process"
	_ "github.com/Kittengarten/KittenCore/abuse"
	"github.com/Kittengarten/KittenCore/kitten"
	_ "github.com/Kittengarten/KittenCore/perf"
	_ "github.com/Kittengarten/KittenCore/sfacg"
	_ "github.com/Kittengarten/KittenCore/stack"

	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/music"
	_ "github.com/FloatTech/ZeroBot-Plugin/plugin/qqwife"

	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/driver"

	log "github.com/sirupsen/logrus"
)

// 颜色代码常量
const (
	colorCodePanic = "\x1b[1;31m" // color.Style{color.Bold, color.Red}.String()
	colorCodeFatal = "\x1b[1;31m" // color.Style{color.Bold, color.Red}.String()
	colorCodeError = "\x1b[31m"   // color.Style{color.Red}.String()
	colorCodeWarn  = "\x1b[33m"   // color.Style{color.Yellow}.String()
	colorCodeInfo  = "\x1b[37m"   // color.Style{color.White}.String()
	colorCodeDebug = "\x1b[32m"   // color.Style{color.Green}.String()
	colorCodeTrace = "\x1b[36m"   // color.Style{color.Cyan}.String()
	colorReset     = "\x1b[0m"
)

// 获取日志等级对应色彩代码
func getLogLevelColorCode(level log.Level) string {
	switch level {
	case log.PanicLevel:
		return colorCodePanic
	case log.FatalLevel:
		return colorCodeFatal
	case log.ErrorLevel:
		return colorCodeError
	case log.WarnLevel:
		return colorCodeWarn
	case log.InfoLevel:
		return colorCodeInfo
	case log.DebugLevel:
		return colorCodeDebug
	case log.TraceLevel:
		return colorCodeTrace
	default:
		return colorCodeInfo
	}
}

// LogFormat 日志输出样式
type LogFormat struct{}

// Format 设置该日志输出样式的具体样式
func (f LogFormat) Format(entry *log.Entry) ([]byte, error) {
	buf := new(bytes.Buffer)

	buf.WriteString(getLogLevelColorCode(entry.Level))

	buf.WriteByte('[')
	buf.WriteString(entry.Time.Format("2006-01-02 15:04:05"))
	buf.WriteString("] ")
	buf.WriteByte('[')
	buf.WriteString(strings.ToUpper(entry.Level.String()))
	buf.WriteString("]: ")
	buf.WriteString(entry.Message)
	buf.WriteString(" \n")

	buf.WriteString(colorReset)

	return buf.Bytes(), nil
}

func init() {
	config := kitten.LoadConfig()
	log.SetFormatter(&LogFormat{}) // 设置日志输出样式
	file, err := os.OpenFile(config.Log.Path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	mw := io.MultiWriter(os.Stdout, file)
	if kitten.Check(err) {
		log.SetOutput(mw)
	} else {
		log.Warn("写入日志失败了喵！")
	}
	log.SetLevel(log.TraceLevel) // 设置最低日志等级
}

func main() {
	// 处理panic，防止程序崩溃
	defer func() {
		if err := recover(); !kitten.Check(err) {
			log.Error("main函数有Bug喵！")
			log.Error(err)
		}
	}()

	config := kitten.LoadConfig()
	log.Info("已经载入配置了喵！")
	rand.Seed(time.Now().UnixNano()) // 全局重置随机数种子，插件无须再次使用

	zero.RunAndBlock(zero.Config{
		NickName:      config.NickName,
		CommandPrefix: config.CommandPrefix,
		SuperUsers:    config.SuperUsers,
		Driver: []zero.Driver{
			&driver.WSClient{
				// OneBot 正向WS 默认使用 6700 端口
				Url:         config.WebSocket.URL,
				AccessToken: config.WebSocket.AccessToken,
			},
		},
	}, process.GlobalInitMutex.Unlock)
}
