package router

import (
	"context"
	"log"
	"mall_go/global"
	"net/http"
	"time"
)

var (
	HttpSrvHandler *http.Server
	port           string
)

func HttpServerRun() {

	r := InitRouter()
	port = global.ServerSetting.HttpPort
	HttpSrvHandler = &http.Server{
		Addr:           ":" + port,
		Handler:        r,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	go httpListen()
	log.Printf(" [INFO] HttpServerRun:%s\n", port)

}
func httpListen() {
	if e := HttpSrvHandler.ListenAndServe(); e != nil {
		log.Fatalf(" [ERROR] HttpServerRun:%s Error:%v\n", port, e)
	}
}
func httpsListen() {
	if e := HttpSrvHandler.ListenAndServeTLS("storage/cert.pem", "storage/key.pem"); e != nil {
		log.Fatalf(" [ERROR] HttpServerRun:%s e:%v\n", port, e)
	}
}
func HttpServerStop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if e := HttpSrvHandler.Shutdown(ctx); e != nil {
		log.Fatalf(" [ERROR] HttpServerStop err:%v\n", e)
	}
	log.Printf(" [INFO] HttpServerStop stopped\n")
}
