package mysqlplus

import (
	"fmt"
	"reflect"
	"strings"
)

// modelCententInteface 模型处理中心
type modelCententInteface interface {
	getColumns(val interface{}) ([]string, []interface{}) // 获取字段数组及设值数组

	Insert(table string, val interface{}) (string, []interface{}) // Insert 自动生成Insert语句方法

	Update(table string, set interface{}, where string, whereParams ...interface{}) (string, []interface{}) // Update 自动生成Update语句方法
}

// modelCententStruct 模型处理中心结构体
type modelCententStruct struct {
}

var ModelCentent modelCententInteface = (*modelCententStruct)(nil)

// Insert 自动生成Insert语句方法
func (mc *modelCententStruct) Insert(table string, val interface{}) (string, []interface{}) {
	sql := `insert into %s (%s) values (%s)`

	wheres := make([]string, 0)

	columns, params := mc.getColumns(val)
	for i := 0; i < len(columns); i++ {
		wheres = append(wheres, "?")
	}

	sql = fmt.Sprintf(sql, table, strings.Join(columns, ","), strings.Join(wheres, ","))
	return sql, params
}

// Update 自动生成Update语句方法
func (mc *modelCententStruct) Update(table string, set interface{}, where string, whereParams ...interface{}) (string, []interface{}) {
	sql := `update table %s set %s `
	columns, resultParams := mc.getColumns(set)
	setStr := ""
	if len(columns) > 1 {
		setStr = strings.Join(columns, " = ?, ") + " = ?"
	} else {
		setStr = columns[0] + " = ?"
	}

	if where != "" {
		setStr = setStr + " where " + where
		resultParams = append(resultParams, whereParams...)
	}

	sql = fmt.Sprintf(sql, table, setStr)

	return sql, resultParams
}

func (mc *modelCententStruct) getColumns(val interface{}) ([]string, []interface{}) {

	columns := make([]string, 0)
	params := make([]interface{}, 0)

	typeOf := reflect.TypeOf(val)
	valueOf := reflect.ValueOf(val)

	for i := 0; i < typeOf.NumField(); i++ {
		column := typeOf.Field(i).Tag.Get("db")
		value := valueOf.Field(i).Interface()
		if !isBlank(reflect.ValueOf(value)) {
			columns = append(columns, column)
			params = append(params, value)
		}
	}

	return columns, params
}
