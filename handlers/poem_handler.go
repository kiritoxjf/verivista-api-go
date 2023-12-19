package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"verivista-api-go/database"
	"verivista-api-go/interfaces"
)

type Poem struct {
	Title   string `json:"title"`
	Dynasty string `json:"dynasty"`
	Author  string `json:"author"`
	Content string `json:"content"`
	Tag     string `json:"tag"`
}

func PoemHandler(c *gin.Context) {
	DB := database.DBClient
	var poem Poem
	// 获取数据
	err := DB.QueryRow("SELECT title, dynasty, author, content, tag FROM t_poem WHERE id = (SELECT MAX(id) FROM t_poem)").Scan(&poem.Title, &poem.Dynasty, &poem.Author, &poem.Content, &poem.Tag)
	if err != nil {
		logrus.Errorln("[ERROR get poem]: ", err)
		c.JSON(http.StatusInternalServerError, interfaces.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "ERROR GET POEM",
		})
	}
	c.JSON(http.StatusOK, poem)
}
