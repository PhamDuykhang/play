package usercrud

import "time"

type (
	//Employee hold employee information in service
	Employee struct {
		EmpID         string     `json:"emp_id,omitempty" bson:"emp_id,omitempty"`
		EmpName       string     `json:"emp_name,omitempty" bson:"emp_name,omitempty"`
		EmpDepartment string     `json:"emp_department,omitempty" bson:"emp_department,omitempty"`
		EmpRoom       string     `json:"emp_room,omitempty" bson:"emp_room,omitempty"`
		EmpBirthDate  *time.Time `json:"emp_birth_date,omitempty" bson:"emp_birth_date,omitempty"`
		Address       Address    `json:"address,omitempty" json:"address,omitempty"`
		PhoneNum      string     `json:"phone_num,omitempty" bson:"phone_num,omitempty"`
	}
	// Address to manage the address
	Address struct {
		HomeNo   string `json:"home_no,omitempty" bson:"home_no,omitempty"`
		Street   string `json:"street,omitempty" bson:"street,omitempty"`
		District string `json:"district,omitempty" bson:"district,omitempty"`
		Country  string `json:"country,omitempty" bson:"country,omitempty"`
	}
	//EmployeeRequest hold information for add new employee
	EmployeeRequest struct {
		EmpID         string     `json:"emp_id,omitempty" bson:"emp_id,omitempty"`
		EmpName       string     `json:"emp_name,omitempty" bson:"emp_name,omitempty"`
		EmpDepartment string     `json:"emp_department,omitempty" bson:"emp_department,omitempty"`
		EmpRoom       string     `json:"emp_room,omitempty" bson:"emp_room,omitempty"`
		EmpBirthDate  *time.Time `json:"emp_birth_date,omitempty" bson:"emp_birth_date,omitempty"`
		Address       Address    `json:"address,omitempty" json:"address,omitempty"`
		PhoneNum      string     `json:"phone_num,omitempty" bson:"phone_num,omitempty"`
	}
	//DeleteAndFindEmployeeRequest hold information to find and delete employee by id
	DeleteAndFindEmployeeRequest struct {
		EmpID string `json:"emp_id,omitempty" bson:"emp_id,omitempty"`
	}
)
