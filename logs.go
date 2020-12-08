//package utils

//import (
//	"github.com/astaxie/beego"
//	"github.com/astaxie/beego/logs"
//)
//
//func InitLogs() {
//	beego.BConfig.Log.AccessLogsFormat = ""
//	level := beego.AppConfig.String("logs::level")
//	logs.SetLogger(logs.AdapterMultiFile, `{"filename":"logs/admin.log",
//		"separate":["critical", "error", "warning", "info", "debug"],
//		"level":`+ level+ `,
//		"daily":true,
//		"maxdays":10}`)
//	logs.Async() //异步
//	//输出文件名和行号
//	logs.EnableFuncCallDepth(true)
//	logs.SetLogFuncCallDepth(3)
//}

package gUtils

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"path/filepath"
)

func InitLog(dir string, filename string) {
	EnsureDir(dir)
	filePath := filepath.Join(dir, filename)
	//logFile, err := os.OpenFile("logs/system.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	//if nil != err {
	//	panic(err)
	//}
	//log.SetOutput(io.MultiWriter(os.Stdout, logFile))
	//filePath := fmt.Sprintf("logs/%s.log", time.Now().Format("20060102"))
	logs.SetLogger(logs.AdapterConsole)
	logs.SetLogger(logs.AdapterFile, fmt.Sprintf(`{"filename":"%s"}`, filePath))
	//logs.Async(100) //异步
	//输出文件名和行号
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(4)
}

// Info compatibility alias for Warning()
func Info(f interface{}, v ...interface{}) {
	logs.Info(f, v...)
}

// Debug logs a message at debug level.
func Debug(f interface{}, v ...interface{}) {
	logs.Debug(f, v...)
}

func Error(f interface{}, v ...interface{}) {
	logs.Error(f, v...)
}
func Warning(f interface{}, v ...interface{}) {
	logs.Warning(f, v...)
}
