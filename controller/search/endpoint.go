package search

import (
	"go-api-se-project/app"
	"net/http"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Endpoint struct {
	cv *app.Configs
	em *app.ErrorMessage
}

func NewEndpoint(config *app.Configs, errorMessage *app.ErrorMessage) *Endpoint {
	return &Endpoint{
		cv: config,
		em: errorMessage,
	}
}

func (ep *Endpoint) SearchGetEndpoint(r *gin.Context) {
	defer r.Request.Body.Close()
	log.Info("Get Data")

	//call logic
	result, err := SearchLogic()
	if err != nil {
		r.JSON(http.StatusBadRequest,err)
		return
	}
	log.Infof("%+v", result)
	log.Info("Status: Success")
	log.Infoln("*********************************************************************")
	//output of item list
	r.JSON(http.StatusOK, result)
	//return
	
}
