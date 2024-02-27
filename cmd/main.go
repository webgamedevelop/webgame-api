//	@title			webgame-api
//	@version		1.0
//	@description	webgame-api docs
//	@contact.name	webgamedevelop
//	@contact.email	webgamedevelop@163.com
//	@contact.url	http://www.swagger.io/support
//	@host			localhost:8080
//	@BasePath		/api

package main

import (
	"context"
	"errors"
	"flag"
	"net/http"
	"sync"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	swaggofiles "github.com/swaggo/files"
	ginswagger "github.com/swaggo/gin-swagger"
	cliflag "k8s.io/component-base/cli/flag"
	"k8s.io/component-base/cli/globalflag"
	"k8s.io/component-base/version/verflag"
	"k8s.io/klog/v2"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/webgamedevelop/logger"

	"github.com/webgamedevelop/webgame-api/internal/handlers/docs"
	"github.com/webgamedevelop/webgame-api/internal/handlers/healthz"
	"github.com/webgamedevelop/webgame-api/internal/handlers/metrics"
)

func main() {
	var apiAddr string
	var swagHost string
	pflag.StringVar(&apiAddr, "api-bind-address", ":8080", "The address the api endpoint binds to.")
	pflag.StringVar(&swagHost, "swag-host", "localhost:8080", "swagger host")

	var versionFlag pflag.FlagSet
	verflag.AddFlags(&versionFlag)
	globalflag.AddGlobalFlags(pflag.CommandLine, "webgame-api")
	logger.InitFlags(flag.CommandLine)
	pflag.CommandLine.ParseErrorsWhitelist.UnknownFlags = true
	cliflag.InitFlags()

	if pflag.CommandLine.Changed("help") {
		pflag.Usage()
		return
	}

	verflag.PrintAndExitIfRequested()

	ctx := ctrl.SetupSignalHandler()
	setLogger(ctx)
	defer klog.Flush()

	router := gin.Default()

	// TODO
	//  log middleware
	//  recover middleware
	//  tracing route
	router.GET("/metrics", metrics.Metrics)
	router.GET("/healthz", healthz.Healthz)

	router.Use(cors.Default())
	router.Use(gin.BasicAuth(map[string]string{"admin": "admin12345"}))
	// TODO InspectRequest middleware

	docs.SwaggerInfo.Host = swagHost
	router.GET("/swagger/*any", ginswagger.WrapHandler(swaggofiles.Handler))

	// TODO add handlers here

	srv := http.Server{
		Addr:    apiAddr,
		Handler: router,
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			klog.Fatal(err)
			return
		}
		klog.Info("HTTP server shutdown")
	}()

	<-ctx.Done()
	klog.Info("context done, shutting down webgame-api http server")
	if err := srv.Shutdown(context.Background()); err != nil {
		klog.Error(err)
		return
	}

	wg.Wait()
	klog.Info("the Webgame-API HTTP server has shutdown normally")
}

func setLogger(ctx context.Context) {
	l, flush := logger.New(ctx, logger.DefaultEncoderConfig)
	klog.SetLoggerWithOptions(l, klog.FlushLogger(flush))
}
