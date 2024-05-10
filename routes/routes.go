package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {

	c.HTML(http.StatusOK, "form.page.tmpl", nil)

}

func Post(c *gin.Context) {

	data := c.PostForm("number")

	idx, _ := strconv.Atoi(data)

	// titles := []string{"Title-1", "Title-2", "Title-3", "Title-4"}
	// description := []string{"Description-1", "Description-2", "Description-3", "Description-4"}

	c.HTML(http.StatusOK, "news.page.tmpl", gin.H{
		"Title": fmt.Sprintf("Title - %d", idx),
		"Body":  fmt.Sprintf("Body - %d", idx),
	})

}
