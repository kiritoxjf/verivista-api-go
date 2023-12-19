package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"verivista-api-go/database"
	"verivista-api-go/interfaces"
)

type ICP struct {
	License string `json:"license"`
}

func ICPHandler(c *gin.Context) {
	DB := database.DBClient
	var icp ICP
	err := DB.QueryRow("SELECT license FROM t_icp").Scan(&icp.License)
	if err != nil {
		logrus.Errorln("[ERROR get icp]: ", err)
		c.JSON(http.StatusInternalServerError, interfaces.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "ERROR GET ICP",
		})
	}
	c.JSON(http.StatusOK, icp)
}
