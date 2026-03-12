package handler

import (
	"net/http"
	"artifact-store/internal/api"
	"github.com/gin-gonic/gin"
)

type Server struct {}

var _ api.ServerInterface = (*Server)(nil)

func (s *Server) GetCharts(c *gin.Context) {
	charts := []api.Chart {
		{ Id: new(int64(1)), Name: new(string("test-chart")) },
	}
	c.JSON(http.StatusOK, charts)
}
