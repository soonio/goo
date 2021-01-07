package orm

import (
	"database/sql"
	"goo/example/orm/dialect"
	"goo/example/orm/log"
	"goo/example/orm/session"
)

// 负责与DB对象交互，链接，测试链接成功与否

type Engine struct {
	db      *sql.DB
	dialect dialect.Dialect
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

	dial, ok := dialect.GetDialect(driver)
	if !ok {
		log.Errorf("dialect %s Not Found", driver)
		return
	}

	e = &Engine{db: db, dialect: dial}
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
	return session.New(e.db, e.dialect)
}
