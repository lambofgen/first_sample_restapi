package middleware

import (
	"database/sql"
	"fmt"
	"geon/api3/daos"
	log "geon/api3/logger"
	"net/http"
	"time"
)

type Middleware struct {
	log *log.Logger
}

func New(log *log.Logger) *Middleware {
	return &Middleware{log}
}

func (m Middleware) Logging(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		l := m.log.WithRequest(r)
		l.Info("==>> request service operation")
		h.ServeHTTP(w, r)
		l.Info(fmt.Sprintf("<<== service response %v", time.Now().Sub(start)))
	})
}

func (m Middleware) Recover(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rc := recover(); rc != nil {
				m.log.WithRequest(r).Error(fmt.Sprintf("recovered : %s", rc))
			}
		}()
		h.ServeHTTP(w, r)
	})
}

//TxFunc type of custom handler transaction
type TxFunc func(f func(tx daos.Transaction, w http.ResponseWriter, r *http.Request) error) http.HandlerFunc

// Tx middleware for have transaction
func (m Middleware) Tx(db *sql.DB) TxFunc {
	return func(f func(tx daos.Transaction, w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {

			t, err := db.Begin()
			if err != nil {
				panic(err)
			}

			defer func() {
				l := m.log.WithRequest(r)
				if p := recover(); p != nil {
					t.Rollback()
					l.Info("transaction rollbacked")
					panic(p)
				} else if err != nil {
					t.Rollback()
					l.Info("transaction rollbacked")
					panic(err)
				} else {
					err = t.Commit()
					l.Info("transaction commited")
				}
			}()

			err = f(t, w, r)
		}
	}
}

// Ntx middleware for have non transaction
func (m Middleware) Ntx(db *sql.DB) TxFunc {
	return func(f func(tx daos.Transaction, w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			err := f(db, w, r)
			defer func() {
				if p := recover(); p != nil {
					panic(p)
				} else if err != nil {
					panic(err)
				}
			}()
		}
	}
}
