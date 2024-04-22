package support_sql

import (
	"database/sql"
	"github.com/pkg/errors"
	"strings"
	"testing"
	"time"
)

type User struct {
	id         int
	username   string
	createTime time.Time
}

func TestGenerateCode_InsertBatch(t *testing.T) {
	GenerateCode_InsertBatch("t_test", User{}, "id")
}

func TestGenerateCode_UpdateBatch(t *testing.T) {
	GenerateCode_UpdateBatch("t_test", User{}, "id")
}

func TestGenerateCode_Query(t *testing.T) {
	GenerateCode_Query("t_test", User{}, "id")
}

func UpdateBatch_User(conn *sql.DB, datas ...User) (sql.Result, error) {
	if len(datas) == 0 {
		return nil, nil
	}
	sql := "update t_test set username=?,create_time=? where id=?;"
	prepare, err := conn.Prepare(strings.Repeat(sql, len(datas)))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var args []any
	for _, data := range datas {
		args = append(args, data.username, data.createTime, data.id)
	}
	exec, err := prepare.Exec(args...)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return exec, nil
}
