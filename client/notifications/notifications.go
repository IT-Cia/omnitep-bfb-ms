package notifications

import (
	"encoding/json"
	"gsm/go/notifications"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

var notificationsClient notifications.NotificationsClient

func MakeConnection(conn *grpc.ClientConn) {
	notificationsClient = notifications.NewNotificationsClient(conn)
}

func AddRoutes(router *gin.RouterGroup) {
	group := router.Group("notifications")

	group.POST("/", Send)
}

func Send(c *gin.Context) {
	var req notifications.SendRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := notificationsClient.Send(c.Request.Context(), &req)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	b, err := json.Marshal(res)
	jsonData := []byte(b)
	c.Data(http.StatusOK, "application/json", jsonData)
	//c.IndentedJSON(http.StatusOK, res)
	//c.JSON(http.StatusOK, res)
}
