package usercrud

import (
	"context"

	"github.com/PhamDuyKhang/userplayboar/internal/app/feature/organization"
	"github.com/pkg/errors"
)

type (
	//UserManager is a infterace for repo layer
	UserManager interface {
		InsertUser(ctx context.Context, e Employee) (emp Employee, err error)
		UpdateUser(ctx context.Context, e Employee) (emp Employee, err error)
		DeleteUser(ctx context.Context, emID string) (err error)
		Find(ctx context.Context, emID string) (emp Employee, err error)
		FindAll(ctx context.Context) ([]Employee, error)
	}
	//EmployeeManager to expose to handler
	EmployeeManager interface {
		AddNewEmployee(ctx context.Context, e Employee) (newE Employee, err error)
		UpdateEmployee(ctx context.Context, e Employee) (newE Employee, err error)
		DeleteEmployee(ctx context.Context, emID string) (err error)
		FindEmployee(ctx context.Context, emID string) (EmployeeRes, error)
		GetAllEmployee(ctx context.Context) ([]EmployeeRes, error)
	}
	//Service make logice for api
	Service struct {
		repo   UserManager
		orRepo organization.RepoI
	}
)

//NewService init a service
func NewService(r UserManager, orRepo organization.RepoI) *Service {
	return &Service{
		repo:   r,
		orRepo: orRepo,
	}
}

//AddNewEmployee to add new employee to database
func (s *Service) AddNewEmployee(ctx context.Context, e Employee) (newE Employee, err error) {
	return s.repo.InsertUser(ctx, e)
}

//UpdateEmployee to add new employee to database
func (s *Service) UpdateEmployee(ctx context.Context, e Employee) (newE Employee, err error) {
	return s.repo.UpdateUser(ctx, e)
}

//DeleteEmployee to delete employee to database
func (s *Service) DeleteEmployee(ctx context.Context, emID string) (err error) {
	return s.repo.DeleteUser(ctx, emID)
}

//FindEmployee to add new employee to database
func (s *Service) FindEmployee(ctx context.Context, emID string) (EmployeeRes, error) {
	e, err := s.repo.Find(ctx, emID)
	if err != nil {
		return EmployeeRes{}, errors.Wrap(err, "can't find employee form database")
	}
	var listDepartment []string
	err = s.departmentLookup(e.EmpDepartment, &listDepartment)
	if err != nil {
		errors.Wrap(err, "can't make  employee department path")
	}
	logger.Infoc(ctx, "the department path %v", listDepartment)
	stack := make([]string, len(listDepartment))
	for i := range listDepartment {
		stack[len(listDepartment)-(i+1)] = listDepartment[i]
	}

	return EmployeeRes{
		EmpID:         e.EmpID,
		EmpName:       e.EmpName,
		EmpBirthDate:  e.EmpBirthDate,
		TechSkill:     e.TechSkill,
		FullAddress:   e.Address.HomeNo + "/" + e.Address.Street + "/" + e.Address.District + "/",
		PhoneNum:      e.PhoneNum,
		EmpDepartment: stack,
	}, nil

}

//GetAllEmployee get all employee in database
func (s *Service) GetAllEmployee(ctx context.Context) ([]EmployeeRes, error) {
	employees, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	empRes := make([]EmployeeRes, len(employees))

	for i := range employees {
		empRes[i] = EmployeeRes{
			EmpID:        employees[i].EmpID,
			EmpName:      employees[i].EmpName,
			EmpBirthDate: employees[i].EmpBirthDate,
			TechSkill:    employees[i].TechSkill,
			FullAddress:  employees[i].Address.HomeNo + "/" + employees[i].Address.Street + "/" + employees[i].Address.District + "/",
			PhoneNum:     employees[i].PhoneNum,
		}
	}
	return empRes, nil

}

func (s *Service) departmentLookup(departmentID string, stack *[]string) error {
	d, err := s.orRepo.FindDepartmentByID(context.Background(), departmentID)
	if err != nil {
		return errors.Wrap(err, "can't get department")
	}
	*stack = append(*stack, d.ID)
	if d.ParentID == "" {
		return nil
	}
	err = s.departmentLookup(d.ParentID, stack)
	if err != nil {
		return errors.Wrap(err, "can't get department")
	}
	return nil
}
