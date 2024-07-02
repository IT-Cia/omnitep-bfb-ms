package notifications

import (
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

	c.JSON(http.StatusOK, res)
}
