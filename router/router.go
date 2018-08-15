package router

import (
	"database/sql"
	"geon/api3/daos"
	"geon/api3/handlers"
	"geon/api3/logger"
	"geon/api3/middleware"

	"github.com/gorilla/mux"
)

//Mount router
func Mount(r *mux.Router, db *sql.DB) {
	// Logger
	log := logger.New()

	// Middleware
	md := middleware.New(log)
	tx := md.Tx(db)
	ntx := md.Ntx(db)

	// Department
	{
		handler := handlers.NewDeptHandler(daos.NewDepartmentDAO(), log)
		prefix := r.PathPrefix("/dept").Subrouter()

		prefix.HandleFunc("", tx(handler.CreateDept)).Methods("POST")
		prefix.HandleFunc("", ntx(handler.SelectAllDept)).Methods("GET")
		prefix.HandleFunc("/{depId}", tx(handler.UpdateDept)).Methods("PUT")
		prefix.HandleFunc("/{depId}", tx(handler.DeleteDept)).Methods("DELETE")
		prefix.HandleFunc("/{depId}", ntx(handler.SelectDept)).Methods("GET")
	}

	r.Use(md.Recover, md.Logging)
}
