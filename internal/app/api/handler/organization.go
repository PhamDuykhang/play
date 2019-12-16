package handler

import (
	"net/http"

	"github.com/PhamDuyKhang/userplayboar/internal/app/errors"
	"github.com/PhamDuyKhang/userplayboar/internal/app/feature/organization"
	"github.com/teera123/gin"
)

type (
	//OrganizationHandler handle all department incoming request
	OrganizationHandler struct {
		svr organization.ServiceI
		e   *errors.AppErrors
	}
)

//NewOrganizationHandler create OrganizationHandler instance
func NewOrganizationHandler(e *errors.AppErrors, s organization.ServiceI) *OrganizationHandler {
	return &OrganizationHandler{
		svr: s,
		e:   e,
	}
}

//GetDepartment get project base id  in database
func (h *OrganizationHandler) GetDepartment(c *gin.Context) {
	var rs Response
	departID := c.Param("id")
	p, err := h.svr.GetDepartment(c, departID)
	if err != nil {
		rs.StatusCode = p.Code
		rs.Message = p.Message
		logger.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, rs)
		return
	}
	rs.StatusCode = h.e.Success.Code
	rs.Message = h.e.Success.Message
	rs.Data = p
	c.JSON(http.StatusOK, rs)
	return
}

//CreateDepartment add  department in database
func (h *OrganizationHandler) CreateDepartment(c *gin.Context) {
	logger.Infoc(c, "insert new department")
	var r organization.DepartmentRQ
	var rs Response
	err := c.BindJSON(&r)
	if err != nil {
		rs.StatusCode = h.e.Request.RequestInvalid.Code
		rs.Message = h.e.Request.RequestInvalid.Message
		c.AbortWithStatusJSON(http.StatusBadRequest, rs)
		logger.Error(err)
		return
	}
	depart := organization.Organization{
		ID:       r.ID,
		Name:     r.Name,
		Type:     r.Type,
		MetaData: r.MetaData,
		ParentID: r.ParentID,
	}
	d, err := h.svr.CreateDepartment(c, depart)
	if err != nil {
		rs.StatusCode = d.Code
		rs.Message = d.Message
		logger.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, rs)
		return
	}
	rs.StatusCode = h.e.Success.Code
	rs.Message = h.e.Success.Message
	rs.Data = d
	c.JSON(http.StatusOK, rs)
	return
}

//UpdateDepartment update department base id  in database
func (h *OrganizationHandler) UpdateDepartment(c *gin.Context) {
	var rs Response
	var r organization.DepartmentRQ
	err := c.BindJSON(&r)
	if err != nil {
		rs.StatusCode = h.e.Request.RequestInvalid.Code
		rs.Message = h.e.Request.RequestInvalid.Message
		c.AbortWithStatusJSON(http.StatusBadRequest, rs)
		logger.Error(err)
		return
	}
	depart := organization.Organization{
		ID:       r.ID,
		Name:     r.Name,
		Type:     r.Type,
		MetaData: r.MetaData,
		ParentID: r.ParentID,
	}
	d, err := h.svr.UpdateDepartment(c, depart)
	if err != nil {
		rs.StatusCode = d.Code
		rs.Message = d.Message
		logger.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, rs)
		return
	}
	rs.StatusCode = h.e.Success.Code
	rs.Message = h.e.Success.Message
	rs.Data = d
	c.JSON(http.StatusOK, rs)
	return
}

//GetDepartmentTree get all department base root id  in database
func (h *OrganizationHandler) GetDepartmentTree(c *gin.Context) {
	var rs Response
	prjID := c.Param("id")
	p, err := h.svr.RecursiveLookup(c, prjID)
	if err != nil {
		rs.StatusCode = p.Code
		rs.Message = p.Message
		c.AbortWithStatusJSON(http.StatusBadRequest, rs)
		return
	}
	rs.StatusCode = h.e.Success.Code
	rs.Message = h.e.Success.Message
	rs.Data = p
	c.JSON(http.StatusOK, rs)
	return
}
