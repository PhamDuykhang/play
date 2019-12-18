package handler

import (
	"net/http"

	"github.com/PhamDuyKhang/userplayboar/internal/app/feature/usercrud"
	"github.com/teera123/gin"
)

type (
	//CRUDHandler handle all employee incoming request
	CRUDHandler struct {
		svr usercrud.EmployeeManager
	}
)

//NewCRUD create CRUD instance
func NewCRUD(s usercrud.EmployeeManager) *CRUDHandler {
	return &CRUDHandler{
		svr: s,
	}
}

//DeleteEmployee remove employee  form database and base in employee id
func (h CRUDHandler) DeleteEmployee(c *gin.Context) {
	var r usercrud.DeleteAndFindEmployeeRequest
	err := c.BindJSON(&r)
	var rs Response
	if err != nil {
		rs.StatusCode = http.StatusBadRequest
		rs.Message = "can't decode your request"
		c.JSON(http.StatusBadRequest, rs)
		return
	}
	err = h.svr.DeleteEmployee(c, r.EmpID)
	if err != nil {
		rs.StatusCode = http.StatusInternalServerError
		rs.Message = err.Error()
		c.JSON(http.StatusInternalServerError, rs)
		return
	}
	rs.StatusCode = http.StatusOK
	rs.Message = "your request is successfully"
	c.JSON(http.StatusOK, rs)
	return
}

//AddNewEmployee add employee form request
func (h CRUDHandler) AddNewEmployee(c *gin.Context) {
	var r usercrud.EmployeeRequest
	err := c.BindJSON(&r)
	var rs Response
	if err != nil {
		rs.StatusCode = http.StatusBadRequest
		rs.Message = "can't decode your request"
		c.JSON(http.StatusBadRequest, rs)
		return
	}
	emp := usercrud.Employee{
		EmpID:         r.EmpID,
		EmpName:       r.EmpName,
		EmpDepartment: r.EmpDepartment,
		EmpRoom:       r.EmpRoom,
		EmpBirthDate:  r.EmpBirthDate,
		Address:       r.Address,
		PhoneNum:      r.PhoneNum,
		TechSkill:     r.TechSkill,
	}
	emData, err := h.svr.AddNewEmployee(c, emp)
	if err != nil {
		rs.StatusCode = http.StatusInternalServerError
		rs.Message = err.Error()
		c.JSON(http.StatusInternalServerError, rs)
		return
	}
	rs.StatusCode = http.StatusOK
	rs.Message = "your request is successfully"
	rs.Data = emData
	c.JSON(http.StatusOK, rs)
	return
}

//UpdateEmployee update  employee infomation in database form request
func (h CRUDHandler) UpdateEmployee(c *gin.Context) {
	var r usercrud.EmployeeRequest
	err := c.BindJSON(&r)
	var rs Response
	if err != nil {
		rs.StatusCode = http.StatusBadRequest
		rs.Message = "can't decode your request"
		c.JSON(http.StatusBadRequest, rs)
		return
	}
	emp := usercrud.Employee{
		EmpID:         r.EmpID,
		EmpName:       r.EmpName,
		EmpDepartment: r.EmpDepartment,
		EmpRoom:       r.EmpRoom,
		EmpBirthDate:  r.EmpBirthDate,
		Address:       r.Address,
		PhoneNum:      r.PhoneNum,
	}
	emID, err := h.svr.UpdateEmployee(c, emp)
	if err != nil {
		rs.StatusCode = http.StatusInternalServerError
		rs.Message = err.Error()
		c.JSON(http.StatusInternalServerError, rs)
		return
	}
	rs.StatusCode = http.StatusOK
	rs.Message = "your request is successfully"
	rs.Data = emID
	c.JSON(http.StatusOK, rs)
	return
}

//FindEmployee update  employee infomation in database form request
func (h CRUDHandler) FindEmployee(c *gin.Context) {

	id := c.Param("id")
	var rs Response
	if id == "" {
		rs.StatusCode = http.StatusBadRequest
		rs.Message = "can't decode your request"
		c.JSON(http.StatusBadRequest, rs)
		return
	}

	employee, err := h.svr.FindEmployee(c, id)
	if err != nil {
		rs.StatusCode = http.StatusInternalServerError
		rs.Message = err.Error()
		c.JSON(http.StatusInternalServerError, rs)
		return
	}
	rs.StatusCode = http.StatusOK
	rs.Message = "your request is successfully"
	rs.Data = employee
	c.JSON(http.StatusOK, rs)
	return
}

//GetAllEmployee get all employee in database
func (h CRUDHandler) GetAllEmployee(c *gin.Context) {
	var rs Response
	employee, err := h.svr.GetAllEmployee(c)
	if err != nil {
		logger.Errorc(c, "get employee data error: %v", err)
		rs.StatusCode = http.StatusInternalServerError
		rs.Message = err.Error()
		c.JSON(http.StatusInternalServerError, rs)
		return
	}
	rs.StatusCode = http.StatusOK
	rs.Message = "your request is successfully"
	rs.Data = employee
	c.JSON(http.StatusOK, rs)
	return
}
