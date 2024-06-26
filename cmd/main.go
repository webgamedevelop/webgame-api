package main

import (
	"context"
	"errors"
	"flag"
	"net/http"
	"os"
	"sync"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	swaggofiles "github.com/swaggo/files"
	ginswagger "github.com/swaggo/gin-swagger"
	_ "go.uber.org/automaxprocs"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	cliflag "k8s.io/component-base/cli/flag"
	"k8s.io/component-base/cli/globalflag"
	"k8s.io/component-base/version/verflag"
	"k8s.io/klog/v2"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/webgamedevelop/logger"
	webgamev1 "github.com/webgamedevelop/webgame/api/v1"

	"github.com/webgamedevelop/webgame-api/internal/handlers/api"
	apiv1 "github.com/webgamedevelop/webgame-api/internal/handlers/api/v1"
	"github.com/webgamedevelop/webgame-api/internal/handlers/docs"
	"github.com/webgamedevelop/webgame-api/internal/handlers/healthz"
	"github.com/webgamedevelop/webgame-api/internal/handlers/metrics"
	"github.com/webgamedevelop/webgame-api/internal/handlers/middleware"
	"github.com/webgamedevelop/webgame-api/internal/models"
	pkgclient "github.com/webgamedevelop/webgame-api/pkg/kubernetes/client"
	"github.com/webgamedevelop/webgame-api/pkg/validator"
)

var scheme = runtime.NewScheme()

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
	utilruntime.Must(webgamev1.AddToScheme(scheme))
}

func main() {
	var (
		apiAddr, swagHost, ginMode, adminEmail, adminPhone, adminPassword string
		importData, enableSwag                                            bool
	)

	pflag.StringVar(&apiAddr, "api-bind-address", ":8080", "The address the api endpoint binds to.")
	pflag.BoolVar(&enableSwag, "enable-swag", false, "Enable swagger.")
	pflag.StringVar(&swagHost, "swag-host", "localhost:8080", "Swagger host.")
	pflag.StringVar(&ginMode, "gin-mode", "release", "Gin mode, debug, release or test.")
	pflag.StringVar(&adminEmail, "init-admin-email", "18600001111@139.com", "Initial email for the admin user.")
	pflag.StringVar(&adminPhone, "init-admin-phone", "18600001111", "Initial phone number for the admin user.")
	pflag.StringVar(&adminPassword, "init-admin-password", "admin12345", "Initial password for the admin user.")
	pflag.BoolVar(&importData, "import-initialization-data", false, "Import initialization data, and exit.")

	var versionFlag pflag.FlagSet
	verflag.AddFlags(&versionFlag)
	globalflag.AddGlobalFlags(pflag.CommandLine, "webgame-api")
	logger.InitFlags(flag.CommandLine)
	models.InitFlags(flag.CommandLine)
	middleware.InitFlags(flag.CommandLine)
	pflag.CommandLine.ParseErrorsWhitelist.UnknownFlags = true
	cliflag.InitFlags()

	if pflag.CommandLine.Changed("help") {
		pflag.Usage()
		return
	}

	// usage: --version / --version=raw
	verflag.PrintAndExitIfRequested()

	ctx := ctrl.SetupSignalHandler()
	var err error

	if err = setLogger(ctx, ginMode); err != nil {
		klog.Error(err)
		return
	}

	defer klog.Flush()

	if err = models.Init(); err != nil {
		klog.Error(err)
		return
	}

	if err = models.Migrate(); err != nil {
		klog.Error(err)
		return
	}

	// import initialization data, and exit
	if importData {
		if err = models.Initialize(); err != nil {
			klog.Error(err)
			os.Exit(1)
		}
		klog.Info("initialization data imported")
		return
	}

	if err = models.InitAdminUser("admin", adminEmail, adminPhone, adminPassword); err != nil {
		klog.Error(err)
		return
	}

	if err = pkgclient.Init(ctrl.GetConfigOrDie(), client.Options{Scheme: scheme}); err != nil {
		klog.Error(err)
		return
	}

	var jwtMiddleware *jwt.GinJWTMiddleware
	if jwtMiddleware, err = middleware.JWT(); err != nil {
		klog.Error(err)
		return
	}

	if err = validator.RegisterValidation(); err != nil {
		klog.Error(err)
		return
	}

	// create http router
	router := gin.Default()
	router.Use(cors.Default())
	router.NoRoute(jwtMiddleware.MiddlewareFunc(), middleware.RouteNotFound)
	router.NoMethod(jwtMiddleware.MiddlewareFunc(), middleware.MethodNotAllowed)

	if enableSwag {
		docs.SwaggerInfo.Host = swagHost
		router.GET("/swagger/*any", ginswagger.WrapHandler(swaggofiles.Handler))
	}

	// add metrics and healthz handlers
	router.GET("/metrics", metrics.Metrics)
	router.GET("/healthz", healthz.Healthz)

	apiRouter := router.Group("/api")
	apiRouterV1 := apiRouter.Group("/v1")

	// add user api
	api.AttachUserAPI(apiRouterV1, &apiv1.User{}, jwtMiddleware)

	// add middlewares
	apiRouterV1.Use(jwtMiddleware.MiddlewareFunc())
	apiRouterV1.Use(middleware.InspectRequest())

	// add Resources api
	api.AttachResourceAPI(apiRouterV1, "/secret", &apiv1.Secret{})
	api.AttachResourceAPI(apiRouterV1, "/ingressclass", &apiv1.IngressClass{})
	api.AttachResourceAPI(apiRouterV1, "/webgame", &apiv1.Webgame{})

	srv := http.Server{
		Addr:    apiAddr,
		Handler: router,
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		// start http server
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

func setLogger(ctx context.Context, mode string) error {
	l, flush, err := logger.New(ctx, logger.DefaultEncoderConfig)
	if err != nil {
		return err
	}
	klog.SetLoggerWithOptions(l, klog.FlushLogger(flush))
	ctrl.SetLogger(l)
	gin.SetMode(mode)
	gin.DefaultWriter = logger.Writer()
	gin.DefaultErrorWriter = logger.Writer()
	return nil
}
