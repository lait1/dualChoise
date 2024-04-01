package common

import (
	"dualChoose/internal/config"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// MySQL connect
func SetDataBaseMySQL(logger *zap.Logger, cfg *config.Config) (db *sqlx.DB, err error) {
	//connection := cfg.MySQLUser + ":" + cfg.MySQLPassword + "@tcp(mysql:3306)/" + cfg.MySQLDatabase + "?parseTime=true"
	connection := "root:district13@tcp(10.52.1.60:3306)/best_option?parseTime=true"
	logger.Info(connection)
	logger.Info("MySQL connect", zap.String("connect", SafeMysqlCreds(connection)))
	db, err = sqlx.Open("mysql", connection)
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
