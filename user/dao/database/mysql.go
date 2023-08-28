package database

import "github.com/zeromicro/go-zero/core/stores/sqlx"

type DBConn struct {
	Conn sqlx.SqlConn
}

func Connect(dataSource string) *DBConn {
	sqlConn := sqlx.NewMysql(dataSource)
	return &DBConn{
		Conn: sqlConn,
	}
}
