package models

type Department struct {
	DeptID   string `json:"deptId" db:"dept_id"`
	DeptName string `json:"deptName" db:"dept_name"`
}
