package logicpool

import (
	"sync"

	"github.com/PhamDuyKhang/userplayboar/internal/app/conf"
	"github.com/PhamDuyKhang/userplayboar/internal/app/db"
	"github.com/PhamDuyKhang/userplayboar/internal/app/errors"
	"github.com/PhamDuyKhang/userplayboar/internal/app/feature/hello"
	"github.com/PhamDuyKhang/userplayboar/internal/app/feature/organization"
	"github.com/PhamDuyKhang/userplayboar/internal/app/feature/usercrud"
	"github.com/PhamDuyKhang/userplayboar/internal/app/pkg/glog"
)

var (
	//HlSrv Hello service
	HlSrv hello.Service
	//EmployeeSvr service
	EmployeeSvr usercrud.EmployeeManager
	//Depart is a public service to handle business of department action
	Depart organization.ServiceI
)

var (
	o = sync.Once{}
)
var (
	logger = glog.New().WithPrefix("logic pool")
)

//NewLogicPool create all service logic once service is stared
func NewLogicPool(em *errors.AppErrors, conf *conf.Config) {
	o.Do(func() {
		newLogicPool(em, conf)
	})
}
func newLogicPool(em *errors.AppErrors, conf *conf.Config) {
	c, err := db.EstablishInfrastructure(conf)
	if err != nil {
		panic(err)
	}
	r := hello.NewRicher(c.MongoDBClient)
	HlSrv = hello.NewHelloService(r)

	or := organization.NewOrganizationMongo(c.MongoDBClient)
	Depart = organization.NewService(em, or)

	crud := usercrud.NewCrudMongo(c.MongoDBClient)
	EmployeeSvr = usercrud.NewService(crud, or)
}
