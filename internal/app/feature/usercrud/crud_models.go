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
		TechSkill     []string   `json:"tech_skill,omitempty" bson:"tech_skill,omitempty"`
		Address       Address    `json:"address,omitempty" bson:"address,omitempty"`
		FullAddress   string     `json:"full_address,omitempty" bson:"full_address,omitempty"`
		PhoneNum      string     `json:"phone_num,omitempty" bson:"phone_num,omitempty"`
	}
	// Address to manage the address
	Address struct {
		HomeNo   string `json:"home_no,omitempty" bson:"home_no,omitempty"`
		Street   string `json:"street,omitempty" bson:"street,omitempty"`
		District string `json:"district,omitempty" bson:"district,omitempty"`
		Country  string `json:"country,omitempty" bson:"country,omitempty"`
	}
	//TechSkill hold technical skill information of employee
	TechSkill struct {
		SkillID    string `json:"skill_id,omitempty"`
		SkillValue string `json:"skill_value,omitempty"`
	}
	//EmployeeRequest hold information for add new employee
	EmployeeRequest struct {
		EmpID         string     `json:"emp_id,omitempty"`
		EmpName       string     `json:"emp_name,omitempty"`
		EmpDepartment []string   `json:"emp_department,omitempty"`
		EmpRoom       string     `json:"emp_room,omitempty"`
		EmpBirthDate  *time.Time `json:"emp_birth_date,omitempty"`
		FullAddress   string     `json:"full_address,omitempty"`
		TechSkill     []string   `json:"tech_skill,omitempty"`
		PhoneNum      string     `json:"phone_num,omitempty" bson:"phone_num,omitempty"`
	}

	//EmployeeRes to wrap information and return it for FE
	EmployeeRes struct {
		EmpID         string     `json:"emp_id,omitempty"`
		EmpName       string     `json:"emp_name,omitempty"`
		EmpDepartment []string   `json:"emp_department,omitempty"`
		EmpRoom       string     `json:"emp_room,omitempty"`
		EmpBirthDate  *time.Time `json:"emp_birth_date,omitempty"`
		FullAddress   string     `json:"full_address,omitempty" bson:"full_address,omitempty"`
		PhoneNum      string     `json:"phone_num,omitempty"`
		TechSkill     []string   `json:"tech_skill,omitempty"`
	}
	//DeleteAndFindEmployeeRequest hold information to find and delete employee by id
	DeleteAndFindEmployeeRequest struct {
		EmpID string `json:"emp_id,omitempty" bson:"emp_id,omitempty" uri:"id"`
	}
)
