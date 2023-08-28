package dao

import (
	"context"
	"fmt"
	"go-zero/mall/user/dao/database"
	"go-zero/mall/user/internal/model"
)

type UserDao struct {
	*database.DBConn
}

func NewUserDao(conn *database.DBConn) *UserDao {
	return &UserDao{conn}
}

func (d *UserDao) Save(ctx context.Context, data *model.User) error {
	sql := fmt.Sprintf("insert into %s (name,gender) values (?,?)", data.TableName())
	result, err := d.Conn.ExecCtx(ctx, sql, data.Name, data.Gender)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	data.Id = id
	return nil
}
