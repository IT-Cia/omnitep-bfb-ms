package task

import (
	"gsm/client/utils"
	"gsm/protobuf/taskpb"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

var taskClient taskpb.TaskServiceClient

func MakeConnection(conn *grpc.ClientConn) {
	taskClient = taskpb.NewTaskServiceClient(conn)
}

func AddRoutes(router *gin.RouterGroup) {
	group := router.Group("task")

	group.POST("/", createTask)
	group.GET("/", getAllTasks)
	group.GET("/:id", getTask)
	group.PUT("/:id", updateTask)
	group.DELETE("/:id", deleteTask)
}

// @Tags			Authentication
// @Accept			json
// @Router  		/task [post]
// @Param 			data	body	taskpb.CreateRequest		true "data"
// @Success 200 {object} 		taskpb.CreateResponse
func createTask(c *gin.Context) {
	var req taskpb.CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := taskClient.Create(c.Request.Context(), &req)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Tags			Authentication
// @Accept			json
// @Router  		/task/{id} [get]
// @Param 			id	path	int		true	"id"
// @Success 200 {object} taskpb.GetResponse
func getTask(c *gin.Context) {
	id, err := utils.ParamInt64(c, "id")
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := taskClient.Get(c.Request.Context(), &taskpb.GetRequest{
		Id: id,
	})
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Tags			Authentication
// @Accept			json
// @Router  		/task/{id} [put]
// @Param 			id		path	int							true "id"
// @Param 			data	body	taskpb.UpdateRequest		true "data"
// @Success 200 {object} 			emptypb.Empty
func updateTask(c *gin.Context) {
	id, err := utils.ParamInt64(c, "id")
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var req taskpb.UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	req.Id = id

	res, err := taskClient.Update(c.Request.Context(), &req)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Tags			Authentication
// @Accept			json
// @Router  		/task/{id} [delete]
// @Param 			id		path	int							true "id"
// @Success 200 {object} 			emptypb.Empty
func deleteTask(c *gin.Context) {
	id, err := utils.ParamInt64(c, "id")
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := taskClient.Delete(c.Request.Context(), &taskpb.DeleteRequest{
		Id: id,
	})
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Tags			Authentication
// @Accept			json
// @Router  		/task [get]
// @Param 			page	query	int			false	"page"
// @Param 			size	query	int			false	"size"
// @Param 			term	query	string		false	"term"
// @Success 200 {object} taskpb.GetAllResponse
func getAllTasks(c *gin.Context) {
	page, err := utils.QueryInt64(c, "page")
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	size, err := utils.QueryInt64(c, "size")
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	term := c.Query("term")

	res, err := taskClient.GetAll(c.Request.Context(), &taskpb.GetAllRequest{
		Page: page,
		Size: size,
		Term: term,
	})

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}
