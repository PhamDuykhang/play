package organization

import (
	"context"

	"github.com/PhamDuyKhang/userplayboar/internal/app/errors"
	"github.com/PhamDuyKhang/userplayboar/internal/app/pkg/glog"
	"github.com/palantir/stacktrace"
)

var logger = glog.New().WithPrefix("hello")

type (
	//RepoI hold necessary func t use in servie
	RepoI interface {
		InsertDepartment(ctx context.Context, d Organization) (Organization, error)
		FindChildrentByParentID(ctx context.Context, pID string) ([]Organization, error)
		FindDepartmentByID(ctx context.Context, id string) (Organization, error)
		UpdateDepartment(ctx context.Context, d Organization) (Organization, error)
	}
	//ServiceI ..
	ServiceI interface {
		GetDepartment(ctx context.Context, id string) (DepartmentRes, error)
		UpdateDepartment(ctx context.Context, d Organization) (DepartmentRes, error)
		RecursiveLookup(ctx context.Context, id string) (RecursiveLookupRes, error)
		CreateDepartment(ctx context.Context, d Organization) (DepartmentRes, error)
	}
	//Service instance to handle business code
	Service struct {
		em *errors.AppErrors
		r  RepoI
	}
)

//NewService create service instance
func NewService(e *errors.AppErrors, r RepoI) *Service {
	return &Service{
		em: e,
		r:  r,
	}
}

//CreateDepartment update a existed department information from database base department's id
func (s *Service) CreateDepartment(ctx context.Context, d Organization) (DepartmentRes, error) {
	d, err := s.r.InsertDepartment(ctx, d)
	var res DepartmentRes
	if err != nil {
		logger.Errorc(ctx, "can't insert department")
		res.Code = s.em.Common.Code
		res.Message = s.em.Common.Message
		err := stacktrace.Propagate(err, "can't get department in database")
		return res, err
	}
	logger.Debugc(ctx, "insert department successfully")
	res.Code = s.em.Success.Code
	res.Message = s.em.Success.Message
	res.ID = d.ID
	res.Name = d.Name
	res.Type = d.Type
	res.MetaData = d.MetaData
	res.ParentID = d.ParentID
	return res, nil
}

//UpdateDepartment update a existed department information from database base department's id
func (s *Service) UpdateDepartment(ctx context.Context, d Organization) (DepartmentRes, error) {
	d, err := s.r.UpdateDepartment(ctx, d)
	var res DepartmentRes
	if err != nil {
		logger.Errorc(ctx, "can't update department")
		res.Code = s.em.Common.Code
		res.Message = s.em.Common.Message
		err := stacktrace.Propagate(err, "can't get department in database")
		return res, err
	}
	logger.Debugc(ctx, "update department successfully")
	res.Code = s.em.Success.Code
	res.Message = s.em.Success.Message
	res.ID = d.ID
	res.Name = d.Name
	res.Type = d.Type
	res.MetaData = d.MetaData
	res.ParentID = d.ParentID
	return res, nil
}

//GetDepartment get a department information from database base department's id
func (s *Service) GetDepartment(ctx context.Context, id string) (DepartmentRes, error) {
	d, err := s.r.FindDepartmentByID(ctx, id)
	var res DepartmentRes
	if err != nil {
		res.Code = s.em.Common.Code
		res.Message = s.em.Common.Message
		err := stacktrace.Propagate(err, "can't get department in database")
		return res, err
	}
	res.Code = s.em.Success.Code
	res.Message = s.em.Success.Message
	res.ID = d.ID
	res.Name = d.Name
	res.Type = d.Type
	res.MetaData = d.MetaData
	res.ParentID = d.ParentID
	return res, nil
}

//RecursiveLookup to get a tree of department with root node base on id of paramenter
func (s Service) RecursiveLookup(ctx context.Context, id string) (RecursiveLookupRes, error) {
	var res RecursiveLookupRes
	root, err := s.r.FindDepartmentByID(ctx, id)
	if err != nil {
		res.Code = s.em.Common.Code
		res.Message = s.em.Common.Message
		return res, stacktrace.Propagate(err, "can't get root department form database in database")
	}
	err = s.deepGetTree(ctx, &root)
	if err != nil {
		res.Code = s.em.Common.Code
		res.Message = s.em.Common.Message
		return res, stacktrace.Propagate(err, "can't get all department in data")
	}
	res.Object = root
	return res, nil
}

func (s Service) deepGetTree(ctx context.Context, root *Organization) error {

	child, err := s.r.FindChildrentByParentID(ctx, root.ID)
	if err != nil {
		return stacktrace.Propagate(err, "can't get list childrent department form database in database")
	}
	root.Childrent = child
	for idx := range child {
		err := s.deepGetTree(ctx, &child[idx])
		if err != nil {
			return err
		}
	}
	return nil
}
