package posts

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllPost(c *gin.Context) {
	pageIndexStr := c.Query("pageIndex")
	pageSizeStr := c.Query("pageSize")

	ctx := c.Request.Context()
	pageIndex, err := strconv.Atoi(pageIndexStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("invalid page index").Error(),
		})
		return
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("invalid size").Error(),
		})
		return
	}
	response, err := h.postSvc.GetAllPost(ctx, pageSize, pageIndex)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}
