package controller

import (
	"datastore-service/constants"

	"github.com/gin-gonic/gin"
)

func ExposeEndpoints(r *gin.Engine, mode constants.ServiceMode) {
	r.GET("/health", GetHealth())

	if mode == constants.TestMode || mode == constants.DataStorePodMode {
		dataStoreGroup := r.Group("/api/v1/ds/")
		dataStoreGroup.GET("/health", GetHealth())
		dataStoreGroup.POST("/upload", UploadFileController())
		dataStoreGroup.GET("/download/:refId", DownloadFileController())
	}

	if mode == constants.TestMode || mode == constants.LoadBalancerMode {
		loadBalancerGroup := r.Group("/api/v1/lb/")
		loadBalancerGroup.GET("/health", GetHealth())
	}

}
