package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/extra-time-zone/xgin"
	"github.com/extra-time-zone/xgin/example/config"
	"github.com/extra-time-zone/xgin/example/controller"
	"github.com/extra-time-zone/xgin/example/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	//0、xglobal
	config.Init()

	//1、Init xGinEngine
	xgin.SetConfig(&xgin.Config{
		GinMode: gin.DebugMode,
		//LogFile: "./log/server.log",
	})

	xGinEngine := xgin.NewXGinEngine()
	xGinEngine.Use(middleware.Global())
	xGinEngine.Use(middleware.RecoverPanic())

	//最后调用xgin默认中间件
	xGinEngine.Use(xgin.DefaultMiddleware())

	xGinEngine.GET("/abc", &controller.ABCHandler{})

	xrouter := xGinEngine.Group("/n", middleware.LoginNotVerify())
	xrouter.GET("/test", &controller.TestNotHandler{})

	verifyRouter := xGinEngine.Group("/v", middleware.LoginVerify())
	verifyRouter.GET("/test", &controller.TestHandler{})

	//2、Run
	ipPort := fmt.Sprintf(`%v:%v`, config.Config.Http.Host, config.Config.Http.Port)
	//if err := xGinEngine.Run(ipPort); err != nil {
	//	log.Fatalln(err)
	//}

	srv := &http.Server{
		Addr:    ipPort,
		Handler: xGinEngine.GetGinEngine(),
	}

	//3、Start HTTP Server
	log.Printf("[XGIN Server] Listening and serving HTTP on %v\n", ipPort)
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				log.Print("[XGIN Server] Listen exit successful")
			} else {
				log.Fatalf("[XGIN Server] Listen fatal error: %s\n", err)
			}
		}
	}()

	//4、Capturing the exit signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sig := <-quit
	log.Printf("[XGIN Server] Received signal: %s, shutting down server...", sig)

	//5、Close gracefully
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			log.Println("[XGIN Server] Shut down timeout")
		} else {
			log.Fatalf("[XGIN Server] Forced to shut down, error: %+v", err)
		}
	}

	//6、Clean
	log.Println("[XGIN Server] Exit gracefully")
}
