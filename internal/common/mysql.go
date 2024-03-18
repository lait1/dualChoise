package common

import (
	"dualChoose/internal/config"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// MySQL connect
func SetDataBaseMySQL(logger *zap.Logger, cfg *config.Config) (db *sqlx.DB, err error) {
	logger.Info("MySQL connect", zap.String("connect", SafeMysqlCreds(cfg.MySQLConnect)))
	db, err = sqlx.Open("mysql", cfg.MySQLConnect)
	if err != nil {
		return
	}
	db.SetMaxOpenConns(cfg.MySQLMaxConnections)
	db.SetMaxIdleConns(cfg.MySQLIdleConnections)

	if err = db.Ping(); err != nil {
		return
	}

	return
}

func SafeMysqlCreds(creds string) string {
	cfg, _ := mysql.ParseDSN(creds)
	cfg.Passwd = "___"
	return cfg.FormatDSN()
}
