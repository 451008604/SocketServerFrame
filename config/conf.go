package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/451008604/socketServerFrame/logs"
)

var configPath string // 配置的文件夹路径

type GlobalObj struct {
	Debug            bool   // 是否Debug模式
	Host             string // 当前服务主机IP
	TcpPort          string // 当前服务端口
	Name             string // 当前服务名称
	Version          string // 当前服务版本号
	MaxPackSize      int    // 传输数据包最大值
	MaxConn          int    // 当前服务允许的最大连接数
	WorkerPoolSize   int    // work池大小
	WorkerTaskMaxLen int    // work对应的执行队列内任务数量的上限
	MaxMsgChanLen    int    // 读写消息的通道最大缓冲数
	ProtocolIsJson   bool   // 是否采用Json协议
}

var globalObject *GlobalObj

func init() {
	globalObject = &GlobalObj{
		Debug:            false,
		Host:             "127.0.0.1",
		TcpPort:          "7777",
		Name:             "socketServerFrame",
		Version:          "v0.1",
		MaxPackSize:      4096,
		MaxConn:          10,
		WorkerPoolSize:   3,
		WorkerTaskMaxLen: 1024,
		MaxMsgChanLen:    100,
		ProtocolIsJson:   true,
	}

	globalObject.Reload()
	logs.SetPrintMode(globalObject.Debug)

	str, _ := json.Marshal(globalObject)
	logs.PrintLogInfo(fmt.Sprintf("服务配置参数：%v", string(str)))
}

// GetGlobalObject 获取全局配置对象
func GetGlobalObject() GlobalObj {
	return *globalObject
}

func (o *GlobalObj) Reload() {
	err := json.Unmarshal(getConfigDataToBytes("config.json"), &globalObject)
	logs.PrintLogErr(err)
}

// 获取配置数据到字节
func getConfigDataToBytes(configName string) []byte {
	if configPath == "" {
		// configPath = os.Getenv("GOPATH") + "/src/" + globalObject.Name + "/config/"
		configPath = "./config/"
	}

	bytes, err := ioutil.ReadFile(configPath + configName)
	logs.PrintLogPanic(err)
	return bytes
}
