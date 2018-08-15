package daos

import (
	"database/sql"
	"geon/api3/models"
	"log"
)

//Transaction interface for transaction
type Transaction interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

// DepartmentDAO implements database operations for Department.
type DepartmentDAO struct {
}

// NewDepartmentDAO return department repositorie
func NewDepartmentDAO() *DepartmentDAO {
	return &DepartmentDAO{}
}

// Create insert new department
func (DepartmentDAO) Create(tx Transaction, dp *models.Department) error {
	_, err := tx.Exec("INSERT INTO Department (dept_id,dept_name) VALUES (@p1, @p2)",
		dp.DeptID, dp.DeptName)
	return err
}

// Update update info to depaertment
func (DepartmentDAO) Update(tx Transaction, dp *models.Department) error {
	_, err := tx.Exec("UPDATE Department SET dept_name = @p1 WHERE dept_id = @p2",
		dp.DeptName, dp.DeptID)
	return err
}

// Delete delete department
func (DepartmentDAO) Delete(tx Transaction, depID string) error {
	_, err := tx.Exec("DELETE FROM Department WHERE dept_id = @p1", depID)
	return err
}

// GetOne get one row of department
func (DepartmentDAO) GetOne(tx Transaction, depID string) (*models.Department, error) {
	var dp models.Department
	err := tx.QueryRow("SELECT TOP 1 * FROM Department WHERE dept_id = @p1", depID).Scan(&dp.DeptID, &dp.DeptName)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No department with that ID.")
			return nil, nil
		}
		return nil, err
	}
	return &dp, nil
}

// GetAll get all row of department
func (DepartmentDAO) GetAll(tx Transaction) (*[]*models.Department, error) {
	rows, err := tx.Query("SELECT * FROM Department")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var depts []*models.Department
	for rows.Next() {
		var dp models.Department
		err := rows.Scan(&dp.DeptID, &dp.DeptName)
		if err != nil {
			return nil, err
		}

		depts = append(depts, &dp)
	}
	return &depts, nil
}
