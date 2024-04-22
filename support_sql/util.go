package support_sql

import (
	"bcd-util/util"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"reflect"
	"strings"
	"sync"
)

var cache = make(map[string]string)
var cacheLock = sync.RWMutex{}

func getValTemplate(sqlPre string) string {
	cacheLock.RLock()
	valTemplate, ok := cache[sqlPre]
	cacheLock.RUnlock()
	if !ok {
		cacheLock.Lock()
		valTemplate, ok = cache[sqlPre]
		if !ok {
			paramNum := strings.Count(sqlPre, ",") + 1
			valTemplate = "(" + strings.Repeat(",?", paramNum)[1:] + ")"
			cache[sqlPre] = valTemplate
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

func getFieldsAndColumns(bean any, ignoreFields ...string) (string, []reflect.StructField, []string) {
	set := make(map[string]any)
	for _, e := range ignoreFields {
		set[e] = nil
	}
	typeOf := reflect.TypeOf(bean)
	var fields []reflect.StructField
	var columns []string
	for i := range typeOf.NumField() {
		field := typeOf.Field(i)
		fieldName := field.Name
		_, ok := set[fieldName]
		if ok {
			continue
		}
		columnName := util.CamelCaseToSplitChar(fieldName, '_')
		fields = append(fields, field)
		columns = append(columns, columnName)
	}
	return typeOf.Name(), fields, columns
}

func GenerateCode_InsertBatch(table string, data any, ignoreFields ...string) {
	typeName, fields, columns := getFieldsAndColumns(data, ignoreFields...)
	if len(fields) == 0 {
		util.Log.Warnf("GenerateCode_InsertBatch table[%s] fields empty", table)
	}
	sql1 := fmt.Sprintf("insert into %s(%s) values ", table, strings.Join(columns, ","))
	sql2 := strings.Repeat(",?", len(columns))
	arr := make([]string, len(fields))
	for i, e := range fields {
		arr[i] = "data." + e.Name
	}
	sql := fmt.Sprintf(`func InsertBatch_%s(db *sql.DB, datas ...%s) (sql.Result, error) {
	if len(datas) == 0 {
		return nil, nil
	}
	valueSql := "%s"
	var args []any
	for _, data := range datas {
		args = append(args, %s)
	}
	prepare, err := db.Prepare("%s" + strings.Repeat(valueSql, len(datas))[1:])
	if err != nil {
		return nil, errors.WithStack(err)
	}
	exec, err := prepare.Exec(args)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return exec, nil
}`, typeName, typeName, sql2, strings.Join(arr, ","), sql1)
	println(sql)
}

func GenerateCode_UpdateBatch(table string, data any, ignoreFields ...string) {
	typeName, fields, columns := getFieldsAndColumns(data, ignoreFields...)
	if len(fields) == 0 {
		util.Log.Warnf("GenerateCode_UpdateBatch table[%s] fields empty", table)
	}
	arr1 := make([]string, len(fields))
	arr2 := make([]string, len(fields))
	for i := range fields {
		arr1[i] = columns[i] + "=?"
		arr2[i] = "data." + fields[i].Name
	}
	sql1 := fmt.Sprintf("update %s set %s where id=?;", table, strings.Join(arr1, ","))
	sql := fmt.Sprintf(`func UpdateBatch_%s(db *sql.DB, datas ...%s) (sql.Result, error) {
	if len(datas) == 0 {
		return nil, nil
	}
	sql := "%s"
	prepare, err := db.Prepare(strings.Repeat(sql, len(datas)))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var args []any
	for _, data := range datas {
		args = append(args, %s,data.id)
	}
	exec, err := prepare.Exec(args...)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return exec, nil
}`, typeName, typeName, sql1, strings.Join(arr2, ","))
	println(sql)
}

func GenerateCode_Query(table string, data any, ignoreFields ...string) {
	typeName, fields, columns := getFieldsAndColumns(data, ignoreFields...)
	if len(fields) == 0 {
		util.Log.Warnf("GenerateCode_Query table[%s] fields empty", table)
	}
	arr1 := make([]string, len(fields))
	arr2 := make([]string, len(fields))
	arr3 := make([]string, len(fields))
	for i, e := range fields {
		arr1[i] = "        var " + e.Name + " " + e.Type.String()
		arr2[i] = "&" + e.Name
		arr3[i] = e.Name + ":" + e.Name
	}
	sql := fmt.Sprintf(`func Query_%s(db *sql.DB) ([]%s, error) {
	prepare, err := db.Prepare("select %s from %s")
	if err != nil {
		return nil, errors.WithStack(err)
	}
	query, err := prepare.Query()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var res []%s
	for query.Next() {
%s
		err := query.Scan(%s)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		res = append(res, %s{%s})
	}
	return res, nil
}`, typeName, typeName, strings.Join(columns, ","), table, typeName, strings.Join(arr1, "\n"), strings.Join(arr2, ","), typeName, strings.Join(arr3, ","))
	println(sql)
}
