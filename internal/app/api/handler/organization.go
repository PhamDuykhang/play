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
	departID := c.Param("deparid")
	p, err := h.svr.GetDepartment(c, departID)
	if err != nil {
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
		logger.Error(err)
		rs.StatusCode = h.e.Common.Code
		rs.Message = h.e.Common.Message
		c.AbortWithStatusJSON(http.StatusInternalServerError, rs)
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
	prjID := c.Param("deparid")
	p, err := h.svr.RecursiveLookup(c, prjID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, rs)
		return
	}
	rs.StatusCode = h.e.Success.Code
	rs.Message = h.e.Success.Message
	rs.Data = p
	c.JSON(http.StatusOK, rs)
	return
}

//AddSkill add a skill into database
func (h *OrganizationHandler) AddSkill(c *gin.Context) {
	var rs Response
	var rq organization.SkillRQ
	err := c.BindJSON(&rq)
	if err != nil {
		rs.StatusCode = h.e.Request.RequestInvalid.Code
		rs.Message = h.e.Request.RequestInvalid.Message
		c.AbortWithStatusJSON(http.StatusBadRequest, rs)
		logger.Error(err)
		return
	}
	s := organization.Skill{
		SkillID:    rq.SkillID,
		SkillValue: rq.SkillValue,
	}
	p, err := h.svr.AddNewSkill(c, s)
	if err != nil {
		rs.StatusCode = h.e.Common.Code
		rs.Message = h.e.Common.Message
		logger.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, rs)
		return
	}
	rs.StatusCode = h.e.Success.Code
	rs.Message = h.e.Success.Message
	rs.Data = p
	c.JSON(http.StatusOK, rs)
	return
}

//ListSkill get all skill base  id  in database
func (h *OrganizationHandler) ListSkill(c *gin.Context) {
	var rs Response
	p, err := h.svr.GetListSkill(c)
	if err != nil {
		rs.StatusCode = h.e.Common.Code
		rs.Message = h.e.Common.Message
		c.AbortWithStatusJSON(http.StatusInternalServerError, rs)
		return
	}
	rs.StatusCode = h.e.Success.Code
	rs.Message = h.e.Success.Message
	rs.Data = p
	c.JSON(http.StatusOK, rs)
	return
}
