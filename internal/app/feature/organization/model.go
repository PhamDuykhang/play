package organization

import "github.com/PhamDuyKhang/userplayboar/internal/app/errors"

const (
	//DeliveryCenter represents for a delivery center unit in company
	DeliveryCenter = "DeliveryCenter"
	//DeliveryGroup represents for a delivery group unit in company
	DeliveryGroup = "DeliveryGroup"
	//Project represents for project of dc or dg
	Project = "Project"
)

type (
	//Organization is a struct represents for department hierarchy in company
	Organization struct {
		ID        string         `json:"id,omitempty" bson:"id,omitempty"`
		Name      string         `json:"name,omitempty" bson:"name,omitempty"`
		Type      string         `json:"-" bson:"type,omitempty"`
		MetaData  interface{}    `json:"-" bson:"meta_data,omitempty"`
		ParentID  string         `json:"parent_id,omitempty" bson:"parent_id,omitempty"`
		Childrent []Organization `json:"childrent,omitempty" bson:"-"`
	}
	/*Response struct from service layer*/

	//RecursiveLookupRes to wrap data from database and remove necessary fields
	RecursiveLookupRes struct {
		errors.ErrorUnit
		Object interface{} `json:"childrent,omitempty" bson:"-"`
	}

	//DepartmentRes to wrap department data from database and remove necessary
	DepartmentRes struct {
		errors.ErrorUnit
		ID       string      `json:"id,omitempty" `
		Name     string      `json:"name,omitempty" `
		Type     string      `json:"type,omitempty"`
		MetaData interface{} `json:"meta_data,omitempty"`
		ParentID string      `json:"parent_id,omitempty" `
	}

	//DepartmentRQ the struct hold request information for update and add new department
	DepartmentRQ struct {
		ID        string         `json:"id,omitempty" bson:"id,omitempty"`
		Name      string         `json:"name,omitempty" bson:"name,omitempty"`
		Type      string         `json:"-" bson:"type,omitempty"`
		MetaData  interface{}    `json:"-" bson:"meta_data,omitempty"`
		ParentID  string         `json:"parent_id,omitempty" bson:"parent_id,omitempty"`
		Childrent []Organization `json:"childrent,omitempty" bson:"-"`
	}
)
