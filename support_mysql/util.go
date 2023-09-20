package support_mysql

import (
	"database/sql"
	"github.com/pkg/errors"
	"strings"
	"sync"
)

var cache = make(map[string]string)
var cacheLock = sync.Mutex{}

func getValTemplate(sqlPre string) string {
	valTemplate, ok := cache[sqlPre]
	if !ok {
		cacheLock.Lock()
		valTemplate, ok = cache[sqlPre]
		if !ok {
			paramNum := strings.Count(sqlPre, ",") + 1
			valTemplate = "(" + strings.Repeat(",?", paramNum)[1:] + ")"
		}
		cacheLock.Unlock()
	}
	return valTemplate
}

func BatchInsert[T any](db *sql.DB, sqlPre string, datas []T, batch int, getParamFn func(T) []any) error {
	valTemplate := getValTemplate(sqlPre)
	dataLen := len(datas)
	if dataLen == 0 {
		return nil
	}
	count := dataLen / batch
	leave := dataLen % batch
	var batchs []int
	for i := 0; i < count; i++ {
		batchs = append(batchs, batch)
	}
	if leave > 0 {
		batchs = append(batchs, leave)
	}
	for i, b := range batchs {
		valStr := strings.Repeat(","+valTemplate, b)[1:]
		s := sqlPre + " values " + valStr
		prepare, err := db.Prepare(s)
		if err != nil {
			return errors.WithStack(err)
		}
		var params []any
		for j := 0; j < b; j++ {
			params = append(params, getParamFn(datas[i*batch+j])...)
		}
		_, err = prepare.Exec(params...)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

func Insert(db *sql.DB, sqlPre string, args []any) error {
	valTemplate := getValTemplate(sqlPre)
	s := sqlPre + " values " + valTemplate
	prepare, err := db.Prepare(s)
	if err != nil {
		return errors.WithStack(err)
	}
	_, err = prepare.Exec(args...)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
