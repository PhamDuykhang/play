package usercrud

import "context"

type (
	//UserManager is a infterace for repo layer
	UserManager interface {
		InsertUser(ctx context.Context, e Employee) (empID string, err error)
		UpdateUser(ctx context.Context, e Employee) (empID string, err error)
		DeleteUser(ctx context.Context, emID string) (err error)
		Find(ctx context.Context, emID string) (emp Employee, err error)
	}
	//EmployeeManager to expose to handler
	EmployeeManager interface {
		AddNewEmployee(ctx context.Context, e Employee) (empID string, err error)
		UpdateEmployee(ctx context.Context, e Employee) (empID string, err error)
		DeleteEmployee(ctx context.Context, emID string) (err error)
		FindEmployee(ctx context.Context, emID string) (emp Employee, err error)
	}
	//Service make logice for api
	Service struct {
		repo UserManager
	}
)

//NewService init a service
func NewService(r UserManager) *Service {
	return &Service{
		repo: r,
	}
}

//AddNewEmployee to add new employee to database
func (s *Service) AddNewEmployee(ctx context.Context, e Employee) (empID string, err error) {
	return s.repo.InsertUser(ctx, e)
}

//UpdateEmployee to add new employee to database
func (s *Service) UpdateEmployee(ctx context.Context, e Employee) (empID string, err error) {
	return s.repo.UpdateUser(ctx, e)
}

//DeleteEmployee to delete employee to database
func (s *Service) DeleteEmployee(ctx context.Context, emID string) (err error) {
	return s.repo.DeleteUser(ctx, emID)
}

//FindEmployee to add new employee to database
func (s *Service) FindEmployee(ctx context.Context, emID string) (emp Employee, err error) {
	return s.repo.Find(ctx, emID)
}
