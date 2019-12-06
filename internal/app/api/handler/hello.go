package handler

import (
	"net/http"

	"github.com/PhamDuyKhang/userplayboar/internal/app/errors"
	"github.com/PhamDuyKhang/userplayboar/internal/app/feature/hello"
	"github.com/PhamDuyKhang/userplayboar/internal/app/pkg/glog"
	"github.com/teera123/gin"
)

type (
	//HelloH the struct hold information of hello handler
	HelloH struct {
		srv hello.Service
		em  *errors.AppErrors
	}
)

var logger = glog.New().WithPrefix("helloH")

//NewHello create hello instance
func NewHello(e *errors.AppErrors, s hello.Service) *HelloH {
	return &HelloH{
		srv: s,
		em:  e,
	}
}

//Hello save mgs into database and respond a hello world phase
func (h *HelloH) Hello(c *gin.Context) {
	st := h.srv.Say("hello")
	logger.Info(st)
	c.JSON(http.StatusOK, gin.H{"data": st})
	return
}
