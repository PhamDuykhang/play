package logicpool

import (
	"sync"

	"github.com/PhamDuyKhang/userplayboar/internal/app/conf"
	"github.com/PhamDuyKhang/userplayboar/internal/app/db"
	"github.com/PhamDuyKhang/userplayboar/internal/app/feature/hello"
)

var (
	//HlSrv Hello service
	HlSrv hello.Service
)

var (
	o = sync.Once{}
)

//NewLogicPool create all service logic once service is stared
func NewLogicPool(conf *conf.Config) {
	o.Do(func() {
		newLogicPool(conf)
	})
}
func newLogicPool(conf *conf.Config) {
	c, err := db.EstablishInfrastructure(conf)
	if err != nil {
		panic(err)
	}
	r := hello.NewRicher(c.MongoDBClient)
	HlSrv = hello.NewHelloService(r)

}
