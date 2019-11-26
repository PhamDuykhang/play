package api

import (
	"net/http"

	"github.com/teera123/gin"

	"github.com/PhamDuyKhang/userplayboar/internal/app/api/handler"
	"github.com/PhamDuyKhang/userplayboar/internal/app/conf"
	"github.com/PhamDuyKhang/userplayboar/internal/app/logicpool"
	"github.com/PhamDuyKhang/userplayboar/internal/app/pkg/glog"
)

const (
	post = http.MethodPost
	get  = http.MethodGet
)

var (
	logger = glog.New().WithPrefix("api")
)

type (
	//Router hold all information of a router
	Router struct {
		desc     string
		path     string
		method   string
		endPoint gin.HandlerFunc
	}
)

//Init init  gin router instand for RESTfull service
func Init(cf *conf.Config) http.Handler {
	logicpool.NewLogicPool(cf)
	h := handler.NewHello(logicpool.HlSrv)
	r := []Router{
		{
			desc:     "ping/pong service",
			path:     "/ping",
			method:   get,
			endPoint: ping,
		},
		{
			desc:     "hello service",
			path:     "/hello",
			method:   get,
			endPoint: h.Hello,
		},
	}
	ro := gin.New()

	v1 := ro.Group("/v1")
	for _, rou := range r {
		v1.Handle(rou.method, rou.path, rou.endPoint)
	}
	return ro

}
func ping(c *gin.Context) {
	logger.Info(c.ClientIP())
	c.JSON(http.StatusOK, gin.H{"data": "pong"})
}
