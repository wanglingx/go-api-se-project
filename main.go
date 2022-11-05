package main

import (
	"fmt"
	"go-api-se-project/app"
	"net/http"
	handlers "go-api-se-project/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {
	cv := app.InitConfigs()
	log.Infof("Database Info: %+v", cv.ConfigInfo)

	em := app.InitErrorMessage(cv)
	app.InitDB(cv)

	r := handlers.Routes{}
	handleRoute := r.InitTransactionRoute(cv, em)

	AppHttp := &http.Server{
		Addr:    fmt.Sprint(":", 3000),
		Handler: handleRoute,
	}
    
    fmt.Printf("[Starting] [HOST]http://localhost%+v/\n" , AppHttp.Addr)

	if err := AppHttp.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Panicf("Server listen: %s\n", err)
	} 	else if err != nil {
			log.Panicf("Server listen error: %s\n", err)
		}
 
}
