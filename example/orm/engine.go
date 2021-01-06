package orm

import (
	"database/sql"
	"goo/example/orm/log"
	"goo/example/orm/session"
)

// 用于管理 db 对象

type Engine struct {
	db *sql.DB
}

func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)

	if err != nil {
		log.Error(err)
		return
	}

	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}

	e = &Engine{db: db}
	log.Info("connect db success")
	return
}

func (e *Engine) Close() {
	if err := e.db.Close(); err != nil {
		log.Error("failed to disconnect databases")
	}
	log.Info("close databases success")
}

func (e *Engine) NewSession() *session.Session {
	return session.New(e.db)
}
