package dao

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero/mall/user/dao/database"
	"go-zero/mall/user/internal/model"
)

type UserDao struct {
	*database.DBConn
}

var cacheUserIdPrefix = "user_"

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

func (d *UserDao) FindById(ctx context.Context, id int64) (user *model.User, err error) {
	user = &model.User{}
	query := fmt.Sprintf("select * from %s where id = ?", user.TableName())
	userIdKey := fmt.Sprintf("%s:%d", cacheUserIdPrefix, id)
	err = d.ConnCache.QueryRowCtx(ctx, user, userIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	return
}

func (d *UserDao) FindByName(ctx context.Context, name string) (user *model.User, err error) {
	user = &model.User{}
	query := fmt.Sprintf("select * from %s where name = ?", user.TableName())
	userIdKey := fmt.Sprintf("%s:%s", cacheUserIdPrefix, name)
	err = d.ConnCache.QueryRowCtx(ctx, user, userIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		return conn.QueryRowCtx(ctx, v, query, name)
	})
	return
}
