package main

import (
	"blog/global"
	"blog/internal/model"
	"blog/internal/routers"
	"blog/pkg/logger"
	"blog/pkg/setting"
	"blog/pkg/tracer"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupsetting err: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setuplogger err: %v", err)
	}
	err = setupTracer()
	if err != nil {
		log.Fatalf("init.setuptracer err: %v", err)
	}
}

// @title 博客系统
// @version 1.0
// @description golang blog
// @termsOfService
func main() {
	global.Logger.Infof("%s:go-blog/%s", "eds", "blog")
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRoter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return nil
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return nil
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return nil
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return nil
	}
	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Tracer", &global.TracerSetting)
	if err != nil {
		return err
	}

	global.JWTSetting.Expire *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   600, //最大600M
		MaxAge:    10,  //最多10天
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}

func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer(
		global.TracerSetting.ServiceName,
		global.TracerSetting.AgentHostPort,
	)
	if err != nil {
		return err
	}
	global.Tracer = jaegerTracer
	return nil
}
