package handlers

import (
	"encoding/json"
	"geon/api3/daos"
	log "geon/api3/logger"
	"geon/api3/models"
	"geon/api3/response"
	"net/http"

	"github.com/gorilla/mux"
)

//DeptHandler : department handler
type DeptHandler struct {
	dao    *daos.DepartmentDAO
	logger *log.Logger
}

//NewDeptHandler : init new department handler
func NewDeptHandler(dps *daos.DepartmentDAO, lg *log.Logger) *DeptHandler {
	return &DeptHandler{dps, lg}
}

//SelectDept : get department info
func (h DeptHandler) SelectDept(tx daos.Transaction, w http.ResponseWriter, r *http.Request) error {
	deptID := mux.Vars(r)["depId"]
	dept, err := h.dao.GetOne(tx, deptID)
	if err != nil {
		response.InternalServerError(w)
		return err
	}

	response.ResponseStatusWithJSON(http.StatusOK, dept, w)
	return nil
}

//SelectAllDept : get all department info
func (h DeptHandler) SelectAllDept(tx daos.Transaction, w http.ResponseWriter, r *http.Request) error {
	depts, err := h.dao.GetAll(tx)
	if err != nil {
		response.InternalServerError(w)
		return err
	}

	response.ResponseStatusWithJSON(http.StatusOK, depts, w)
	return nil
}

//CreateDept : create new department
func (h DeptHandler) CreateDept(tx daos.Transaction, w http.ResponseWriter, r *http.Request) error {
	var dp models.Department
	err := json.NewDecoder(r.Body).Decode(&dp)
	if err != nil {
		response.InternalServerError(w)
		return err
	}

	err = h.dao.Create(tx, &dp)
	if err != nil {
		response.InternalServerError(w)
		return err
	}

	response.ResponseStatus(http.StatusCreated, w)
	return nil
}

//UpdateDept : update department
func (h DeptHandler) UpdateDept(tx daos.Transaction, w http.ResponseWriter, r *http.Request) error {
	var dp models.Department
	err := json.NewDecoder(r.Body).Decode(&dp)
	if err != nil {
		response.InternalServerError(w)
		return err
	}

	if dp.DeptID = mux.Vars(r)["depId"]; dp.DeptID == "" {
		response.ResponseStatus(http.StatusNotModified, w)
		return nil
	}

	err = h.dao.Update(tx, &dp)
	if err != nil {
		response.InternalServerError(w)
		return err
	}

	response.ResponseStatus(http.StatusOK, w)
	return nil
}

//DeleteDept : delete department
func (h DeptHandler) DeleteDept(tx daos.Transaction, w http.ResponseWriter, r *http.Request) error {
	deptID := mux.Vars(r)["depIda"]
	if deptID == "" {
		response.ResponseStatus(http.StatusNotModified, w)
		return nil
	}

	err := h.dao.Delete(tx, deptID)
	if err != nil {
		response.InternalServerError(w)
		return err
	}

	response.ResponseStatus(http.StatusOK, w)
	return nil
}
