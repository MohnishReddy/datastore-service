package controller

import (
	"datastore-service/constants"
	"datastore-service/models"
	"datastore-service/services"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetHealth() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, &models.BasicResponse{
			Message: "Ping!",
		})
	}
}

func UploadFileController() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadGateway, &models.BasicResponse{
				Error: err.Error(),
			})
			return
		}

		compressionReq := c.Request.Header.Get(constants.COMPRESSION_REQ_HEADER)

		// File info
		fileId := c.Request.Header.Get(constants.FILE_ID_HEADER)
		if fileId == "" {
			fileId = "file_" + uuid.New().String()
		}
		fileName := c.Request.Header.Get(constants.FILE_NAME_HEADER)
		if fileName == "" {
			fileName = file.Filename
		}

		// Prev server Info
		prevServerUrl := c.Request.Header.Get(constants.PREV_SERVER_URL_HEADER)
		prevServerId := c.Request.Header.Get(constants.PREV_SERVER_ID_HEADER)
		prevDataRefId := c.Request.Header.Get(constants.PREV_DATA_REF_ID_HEADER)

		// Open the uploaded file
		uploadedFile, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, &models.BasicResponse{
				Error: err.Error(),
			})
			return
		}
		defer uploadedFile.Close()

		fileBytes := make([]byte, file.Size)

		_, err = io.ReadFull(uploadedFile, fileBytes)
		if err != nil {
			c.JSON(http.StatusInternalServerError, &models.BasicResponse{
				Error: err.Error(),
			})
			return
		}

		refId, dataUploaded, err := services.NewFileUploadHandler(prevServerUrl, prevServerId, prevDataRefId).HandleFileUpload(c, fileId, fileName, fileBytes, compressionReq == "true")
		if err != nil {
			c.JSON(http.StatusInternalServerError, &models.BasicResponse{
				Error: err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, &models.FileUploadResponse{
			FileId:          fileId,
			FileName:        fileName,
			FileSize:        len(fileBytes),
			DataUploaded:    dataUploaded,
			RefId:           *refId,
			CompressionUsed: compressionReq == "true",
		})
	}
}

func DownloadFileController() gin.HandlerFunc {
	return func(c *gin.Context) {
		refId := c.Params.ByName("refId")
		fileData, err := services.NewFileDownloadHandler().HandleFileDownload(c, refId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, &models.BasicResponse{
				Error: err.Error(),
			})
			return
		}

		var nextServerDetails []models.ServerDetails

		for _, s := range fileData.NextServerDetails {
			if nextServerDetails == nil {
				nextServerDetails = []models.ServerDetails{}
			}
			nextServerDetails = append(nextServerDetails, models.ServerDetails{
				ServerId:  s.ServerId,
				ServerUrl: s.ServerUrl,
				RefId:     s.RefId,
			})
		}

		c.JSON(http.StatusOK, &models.FileDownloadResponse{
			FileBytes:         &fileData.Data,
			FileId:            fileData.FileId,
			FileName:          fileData.FileName,
			NextServerDetails: &nextServerDetails,
		})
	}
}
