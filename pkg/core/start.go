package core

import (
	"flag"
	"github.com/gin-gonic/gin"
	"log"
	"mall_go/global"
	"mall_go/internal/cache"
	"mall_go/internal/model"
	"mall_go/pkg/logger"
	"mall_go/pkg/setting"
	"mall_go/pkg/tracer"
	"os"
	"time"
)

func StartModule() {
	var e error
	if e = initFlag(); e != nil {
		log.Fatalf("initFlag e: %v", e)
	}
	if e = initSetting(); e != nil {
		log.Fatalf("initSetting e: %v", e)
	}
	if e = initValidate(); e != nil {
		log.Fatalf("initValidate e: %v", e)
	}

	if e = initLogger(); e != nil {
		log.Fatalf("initLogger e: %v", e)
	}
	e = initDBEngine()
	if e != nil {
		log.Fatalf("initDBEngine e: %v", e)
	}
	//e = initRedis()
	//if e != nil {
	//	log.Fatalf("initRedis e: %v", e)
	//}

	//e = initTracer()
	//if e != nil {
	//	log.Fatalf("initTracer e: %v", e)
	//}
	e = initTask()
	if e != nil {
		log.Fatalf("initTask e: %v", e)
	}
	initMsg()
	//msg.StartMsg()
}

var (
	port    string
	runMode string
)

func initFlag() error {
	flag.StringVar(&port, "port", "", "启动端口")
	flag.StringVar(&runMode, "mode", gin.Mode(), "启动模式")
	flag.Parse()

	return nil
}
func initSetting() (e error) {
	path, _ := os.Getwd()
	config := path + "/configs"
	s, e := setting.NewSetting(config, runMode, "yaml")
	if e != nil {
		return
	}
	e = s.ReadSection("Server", &global.ServerSetting)
	if e != nil {
		return
	}
	e = s.ReadSection("App", &global.AppSetting)
	if e != nil {
		return
	}

	if e = s.ReadSection("Database", &global.DatabaseSetting); e != nil {
		return
	}

	if e = s.ReadSection("Redis", &global.RedisSetting); e != nil {
		return
	}

	if e = s.ReadSection("JWT", &global.JWTSetting); e != nil {
		return
	}

	if e = s.ReadSection("Log", &global.LogSetting); e != nil {
		return
	}

	if e = s.ReadSection("Email", &global.EmailSetting); e != nil {
		return
	}

	global.AppSetting.DefaultContextTimeout *= time.Second
	global.JWTSetting.Expire *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	if port != "" {
		global.ServerSetting.HttpPort = port
	}
	if runMode != "" {
		global.ServerSetting.RunMode = runMode
	}

	return
}
func initLogger() (e error) {

	logSet := global.LogSetting
	global.Logger, e = logger.NewLogger(logSet.Formatter, logSet.Level, logSet.ReportCaller, logSet.SavePath)
	if e != nil {
		return
	}
	//logger.Log.AddHook(&logger.AppHook{})
	return
}

func initDBEngine() (e error) {

	global.DBEngine, e = model.NewDBEngine(global.DatabaseSetting)
	if e != nil {
		return
	}

	return
}
func initTracer() (e error) {
	jaegerTracer, _, e := tracer.NewJaegerTracer("s", "127.0.0.1:6831")
	if e != nil {
		return e
	}
	global.Tracer = jaegerTracer
	return nil
}
func initRedis() (e error) {
	global.RDB, e = cache.NewRedisClient(global.RedisSetting)
	if e != nil {
		return
	}
	return
}
func initTask() (e error) {
	go testTask()
	return
}
func testTask() {
	for {
		time.Sleep(1 * time.Second)

	}
}
func startMsg() {
	//RegisterNotification(1,testNotify)
}
