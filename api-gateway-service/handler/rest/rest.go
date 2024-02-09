package rest

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nafisalfiani/p3-final-project/api-gateway-service/docs/swagger"
	"github.com/nafisalfiani/p3-final-project/api-gateway-service/handler/scheduler"
	"github.com/nafisalfiani/p3-final-project/api-gateway-service/usecase"
	"github.com/nafisalfiani/p3-final-project/lib/appcontext"
	"github.com/nafisalfiani/p3-final-project/lib/auth"
	"github.com/nafisalfiani/p3-final-project/lib/configreader"
	"github.com/nafisalfiani/p3-final-project/lib/log"
	"github.com/nafisalfiani/p3-final-project/lib/parser"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gopkg.in/yaml.v2"
)

const (
	infoRequest  string = `httpclient Sent Request: uri=%v method=%v`
	infoResponse string = `httpclient Received Response: uri=%v method=%v resp_code=%v`
)

var once = &sync.Once{}

type REST interface {
	Run()
}

type rest struct {
	http       *gin.Engine
	conf       Config
	confreader configreader.Interface
	log        log.Interface
	json       parser.JSONInterface
	auth       auth.Interface
	uc         *usecase.Usecases
	scheduler  scheduler.Interface
}

func Init(conf Config, confreader configreader.Interface, log log.Interface, json parser.JSONInterface, auth auth.Interface, uc *usecase.Usecases, scheduler scheduler.Interface) REST {
	r := &rest{}
	once.Do(func() {
		switch conf.Mode {
		case gin.ReleaseMode:
			gin.SetMode(gin.ReleaseMode)
		case gin.DebugMode, gin.TestMode:
			gin.SetMode(gin.TestMode)
		default:
			gin.SetMode("")
		}

		ginEngine := gin.New()

		r = &rest{
			conf:       conf,
			confreader: confreader,
			log:        log,
			auth:       auth,
			json:       json,
			http:       ginEngine,
			uc:         uc,
			scheduler:  scheduler,
		}

		// Set CORS
		switch r.conf.Cors.Mode {
		case "allowall":
			r.http.Use(cors.New(cors.Config{
				AllowAllOrigins: true,
				AllowHeaders:    []string{"*"},
				AllowMethods: []string{
					http.MethodHead,
					http.MethodGet,
					http.MethodPost,
					http.MethodPut,
					http.MethodPatch,
					http.MethodDelete,
				},
			}))
		default:
			r.http.Use(cors.New(cors.DefaultConfig()))
		}

		// Set Recovery
		r.http.Use(gin.Recovery())

		// Set Timeout
		r.http.Use(r.SetTimeout)

		r.Register()
	})

	return r
}

func (r *rest) Run() {
	// Create context that listens for the interrupt signal from the OS.
	c := appcontext.SetServiceVersion(context.Background(), r.conf.Meta.Version)
	ctx, stop := signal.NotifyContext(c, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	port := ":8080"
	if r.conf.Port != "" {
		port = fmt.Sprintf(":%s", r.conf.Port)
	}

	srv := &http.Server{
		Addr:              port,
		Handler:           r.http,
		ReadHeaderTimeout: 2 * time.Second,
	}

	// Initializing the server in a goroutine so that it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			r.log.Error(ctx, fmt.Sprintf("Serving HTTP error: %s", err.Error()))
		}
	}()
	r.log.Info(ctx, fmt.Sprintf("Listening and Serving HTTP on %s", srv.Addr))

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	r.log.Info(ctx, "Shutting down server...")

	// The context is used to inform the server it has timeout duration to finish the request it is currently handling
	quitctx, cancel := context.WithTimeout(c, r.conf.ShutdownTimeout)
	defer cancel()
	if err := srv.Shutdown(quitctx); err != nil {
		r.log.Fatal(quitctx, fmt.Sprintf("Server Shutdown: %s", err.Error()))
	}
	r.log.Info(quitctx, "Server Shut Down.")
}

func (r *rest) Register() {
	// server health and testing purpose
	r.http.GET("/ping", r.Ping)
	r.registerSwaggerRoutes()
	r.registerPlatformRoutes()

	commonPublicMiddlewares := gin.HandlersChain{
		r.addFieldsToContext, r.BodyLogger,
	}

	// auth api
	authv1 := r.http.Group("/auth/v1", commonPublicMiddlewares...)
	authv1.POST("/register", r.RegisterUser)
	authv1.GET("/verify-email/:id", r.VerifyEmail)
	authv1.POST("/login", r.Login)

	// webhooks
	webv1 := r.http.Group("/webhook/v1", commonPublicMiddlewares...)
	webv1.POST("/transaction", r.UpdatePaymentStatus)

	commonPrivateMiddlewares := append(commonPublicMiddlewares, r.VerifyUser)

	// register private middlewares
	v1 := r.http.Group("/api/v1", commonPrivateMiddlewares...)

	// product-service
	v1.GET("/category", r.ListCategory)
	v1.GET("/region", r.ListRegion)
	v1.GET("/ticket", r.ListOpenTicket)
	v1.GET("/ticket-sold", r.ListSoldTicketByMe)
	v1.GET("/ticket-bought", r.ListBoughtTicketByMe)
	v1.POST("/ticket", r.RegisterTicketForSale)
	v1.PUT("/ticket/:id", r.UpdateTicketInfo)
	v1.DELETE("/ticket/:id", r.TakeDownTicket)

	v1.GET("/wishlist", r.ListWishlist)
	v1.GET("/wishlist/:id", r.GetWishlistSubscriber)
	v1.POST("/wishlist", r.SubscribeToWishlist)
	v1.DELETE("/wishlist", r.UnsubscribeFromWishlist)

	// transaction-service
	v1.GET("/transaction", r.ListTransaction)
	v1.POST("/transaction", r.CreateTransaction)
	v1.PUT("/transaction", r.UpdateTransaction)
	v1.GET("/wallet", r.GetWallet)

	// scheduler
	v1.POST("/admin/scheduler/trigger", r.TriggerScheduler)
}

func (r *rest) registerSwaggerRoutes() {
	if r.conf.Swagger.Enabled {
		swagger.SwaggerInfo.Title = r.conf.Meta.Title
		swagger.SwaggerInfo.Description = r.conf.Meta.Description
		swagger.SwaggerInfo.Version = fmt.Sprintf("%s-%s", r.conf.Meta.Version, r.conf.Meta.Environment)
		swagger.SwaggerInfo.Host = r.conf.Meta.Host
		swagger.SwaggerInfo.BasePath = r.conf.Meta.BasePath

		swaggerAuth := gin.Accounts{
			r.conf.Swagger.BasicAuth.Username: r.conf.Swagger.BasicAuth.Password,
		}

		r.http.GET(fmt.Sprintf("/%s/*any", r.conf.Swagger.Path),
			gin.BasicAuthForRealm(swaggerAuth, "Restricted"),
			ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
}

func (r *rest) registerPlatformRoutes() {
	if r.conf.Platform.Enabled {
		platformAuth := gin.Accounts{
			r.conf.Platform.BasicAuth.Username: r.conf.Platform.BasicAuth.Password,
		}

		r.http.GET(fmt.Sprintf("/%s", r.conf.Platform.Path),
			gin.BasicAuthForRealm(platformAuth, "Restricted"),
			r.platformConfig)
	}
}

func (r *rest) platformConfig(ctx *gin.Context) {
	switch ctx.Query("output") {
	case "yaml":
		c, err := yaml.Marshal(r.confreader.AllSettings())
		if err != nil {
			r.httpRespError(ctx, err)
			return
		}
		ctx.String(http.StatusOK, string(c))
	default:
		ctx.IndentedJSON(http.StatusOK, r.confreader.AllSettings())
	}
}
