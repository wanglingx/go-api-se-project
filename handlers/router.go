package handlers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-api-se-project/app"
	"net/http"
)

type route struct {
	Name        string
	Description string
	Method      string
	Pattern     string
	Endpoint    gin.HandlerFunc
	Validation  gin.HandlerFunc
}

type Routes struct {
	pingService []route
}

func (r Routes) InitTransactionRoute(cv *app.Configs, em *app.ErrorMessage) http.Handler {
	
	//ping := ping.NewEndpoint(cv, em)

	r.pingService = []route{
		{
			Name:        "Ping Pong : GET",
			Description: "Ping Pong : HealthCheck",
			Method:      http.MethodGet,
			Pattern:     "/",
			//Endpoint:    ping.PingGetEndpoint,
		},
		{
			Name:        "Ping Pong : GET Prams",
			Description: "Ping Pong : HealthCheck",
			Method:      http.MethodGet,
			Pattern:     "/:name",
		//	Endpoint:    ping.PingGetParamsEndpoint,
		},
	}
	ro := gin.New()

	ro.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT"},
		AllowHeaders:     []string{"Content-Type", "Authorization", "x-api-key"},
		AllowCredentials: true,
	}))

	store := ro.Group("/app/ping")
	for _, e := range r.pingService {
		store.Handle(e.Method, e.Pattern, e.Endpoint)
	}

	return ro
}