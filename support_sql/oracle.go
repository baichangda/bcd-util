package support_sql

import (
	"database/sql"
	"github.com/pkg/errors"
	"strconv"
	"strings"
	"sync"
)

var cache_oracle = make(map[string]string)
var cacheLock_oracle = sync.RWMutex{}

func getValTemplate_oracle(sqlPre string) string {
	cacheLock_oracle.RLock()
	valTemplate, ok := cache_oracle[sqlPre]
	cacheLock_oracle.RUnlock()
	if !ok {
		cacheLock_oracle.Lock()
		valTemplate, ok = cache_oracle[sqlPre]
		if !ok {
			paramNum := strings.Count(sqlPre, ",") + 1
			sb := strings.Builder{}
			for i := range paramNum {
				sb.WriteString(",")
				sb.WriteString(":p" + strconv.Itoa(i+1))
			}
			valTemplate = "(" + sb.String()[1:] + ")"
			cache_oracle[sqlPre] = valTemplate
		}
		cacheLock_oracle.Unlock()
	}
	return valTemplate
}

// BatchInsert_oracle 批量插入oracle
// columns 所有列、格式为 column1,column2,...
func BatchInsert_oracle[T any](db *sql.DB, table string, columns string, datas []T, batch int, getParamFn func(T) []any) error {
	columnCount := strings.Count(columns, ",") + 1
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
	var err error
	sb := strings.Builder{}
	var prepare1 *sql.Stmt
	if dataLen >= batch {
		for i := 1; i <= batch; i++ {
			sb.WriteString("\nINTO " + table + "(" + columns + ") VALUES(")
			for j := 1; j <= columnCount; j++ {
				if j > 1 {
					sb.WriteString(",")
				}
				sb.WriteString(":p" + strconv.Itoa(i*j))
			}
			sb.WriteString(")")
		}
		insertSql := "INSERT ALL" + sb.String() + "\n" + "SELECT * FROM dual"
		prepare1, err = db.Prepare(insertSql)
		defer prepare1.Close()
		if err != nil {
			return errors.WithStack(err)
		}
		sb.Reset()
	}

	var prepare2 *sql.Stmt
	if leave > 0 {
		batchs = append(batchs, leave)
		for i := 1; i <= leave; i++ {
			sb.WriteString("\nINTO " + table + "(" + columns + ") VALUES(")
			for j := 1; j <= columnCount; j++ {
				if j > 1 {
					sb.WriteString(",")
				}
				sb.WriteString(":p" + strconv.Itoa(i*j))
			}
			sb.WriteString(")")
		}
		insertSql := "INSERT ALL" + sb.String() + "\n" + "SELECT * FROM dual"
		prepare2, err = db.Prepare(insertSql)
		defer prepare2.Close()
		if err != nil {
			return errors.WithStack(err)
		}
	}

	for i, b := range batchs {
		var prepare *sql.Stmt
		if b == batch {
			prepare = prepare1
		} else {
			prepare = prepare2
		}
		var params []any
		for j := 0; j < b; j++ {
			index := i*batch + j
			if index >= dataLen {
				println(index)
			}
			params = append(params, getParamFn(datas[index])...)
		}
		_, err = prepare.Exec(params...)
		if err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

// BatchInsert_oracle 逐条插入oracle
// sqlPre 格式为insert into table(column1,column2,...)
func Insert_oracle(db *sql.DB, sqlPre string, args ...[]any) error {
	valTemplate := getValTemplate_oracle(sqlPre)
	s := sqlPre + " values " + valTemplate
	prepare, err := db.Prepare(s)
	if err != nil {
		return errors.WithStack(err)
	}
	defer prepare.Close()
	for _, e := range args {
		_, err = prepare.Exec(e...)
		if err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}
