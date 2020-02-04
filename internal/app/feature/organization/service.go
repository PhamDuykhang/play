package organization

import (
	"context"

	"github.com/PhamDuyKhang/userplayboar/internal/app/errors"
	"github.com/PhamDuyKhang/userplayboar/internal/app/pkg/glog"
	"github.com/palantir/stacktrace"
	pkgerr "github.com/pkg/errors"
)

var logger = glog.New().WithPrefix("organization")

type (
	//RepoI hold necessary func t use in servie
	RepoI interface {
		InsertDepartment(ctx context.Context, d Organization) (Organization, error)
		FindChildrentByParentID(ctx context.Context, pID string) ([]Organization, error)
		FindDepartmentByID(ctx context.Context, id string) (Organization, error)
		UpdateDepartment(ctx context.Context, d Organization) (Organization, error)

		/*Skill repo*/

		InsertSkill(ctx context.Context, sk Skill) (Skill, error)
		FindAllSkill(ctx context.Context) ([]Skill, error)
	}
	//ServiceI ..
	ServiceI interface {
		GetDepartment(ctx context.Context, id string) (DepartmentRes, error)
		UpdateDepartment(ctx context.Context, d Organization) (DepartmentRes, error)
		RecursiveLookup(ctx context.Context, id string) (RecursiveLookupRes, error)
		CreateDepartment(ctx context.Context, d Organization) (DepartmentRes, error)

		/*Skill*/

		AddNewSkill(ctx context.Context, s Skill) (SkillRs, error)
		GetListSkill(ctx context.Context) ([]SkillRs, error)
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
		err := stacktrace.Propagate(err, "can't get department in database")
		return res, err
	}
	logger.Debugc(ctx, "insert department successfully")
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
		err := stacktrace.Propagate(err, "can't get department in database")
		return res, err
	}
	logger.Debugc(ctx, "update department successfully")
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
		err := stacktrace.Propagate(err, "can't get department in database")
		return res, err
	}
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
		return res, stacktrace.Propagate(err, "can't get root department form database in database")
	}
	err = s.deepGetTree(ctx, &root)
	if err != nil {
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
	root.Children = child
	for idx := range child {
		err := s.deepGetTree(ctx, &child[idx])
		if err != nil {
			return err
		}
	}
	return nil
}

//AddNewSkill store a skill in to database
func (s Service) AddNewSkill(ctx context.Context, sk Skill) (SkillRs, error) {
	var skr SkillRs
	sk, err := s.r.InsertSkill(ctx, sk)
	if err != nil {
		return skr, pkgerr.Wrap(err, "can't add new skill into database")
	}
	skr.SkillID = sk.SkillID
	skr.SkillValue = sk.SkillValue
	return skr, nil
}

//GetListSkill get  list skill from database
func (s Service) GetListSkill(ctx context.Context) ([]SkillRs, error) {
	sk, err := s.r.FindAllSkill(ctx)
	if err != nil {
		return nil, pkgerr.Wrap(err, "can't add new skill into database")
	}
	var skr = make([]SkillRs, len(sk))
	for i := range sk {
		skr[i] = SkillRs{
			SkillID:    sk[i].SkillID,
			SkillValue: sk[i].SkillValue,
		}
	}
	return skr, nil
}
