package routers

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type UriParam struct {
	Any string `uri:"any"`
}

func SocketIOJSFile(c *gin.Context) {
	var param UriParam
	if err := c.ShouldBindUri(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"httpCode": http.StatusBadRequest, "error": err.Error()})
		return
	}

	if strings.Contains(param.Any, "js") {
		data, err := os.ReadFile(fmt.Sprintf("../asset%s", param.Any))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"httpCode": http.StatusBadRequest, "error": err.Error()})
		}
		c.Writer.Header().Add("Content-Type", "application/javascript")
		c.Writer.Write(data)
		return
	}
}
