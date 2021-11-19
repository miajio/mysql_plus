package mysqlplus

import (
	"fmt"
	"reflect"
	"strings"
)

// modelCententInteface 模型处理中心
type modelCententInteface interface {
	Insert(table string, val interface{}) (sql string, params []interface{}, err error) // Insert 自动生成Insert语句方法
}

// modelCententStruct 模型处理中心结构体
type modelCententStruct struct {
}

var ModelCentent modelCententInteface = (*modelCententStruct)(nil)

// Insert 自动生成Insert语句方法
func (mc *modelCententStruct) Insert(table string, val interface{}) (sql string, params []interface{}, err error) {
	sql = `insert into %s (%s) values (%s)`

	valOf := reflect.ValueOf(val)
	typeOf := reflect.TypeOf(val)

	columns := make([]string, 0)
	wheres := make([]string, 0)

	for i := 0; i < typeOf.NumField(); i++ {
		column := typeOf.Field(i).Tag.Get("db")
		param := valOf.Field(i).Interface()

		columns = append(columns, column)
		params = append(params, param)
		wheres = append(wheres, "?")
	}

	sql = fmt.Sprintf(sql, table, strings.Join(columns, ","), strings.Join(wheres, ","))
	return
}

// Update 自动生成Update语句方法
func (mc *modelCententStruct) Update(table string, val interface{}) (sql string, params []interface{}, err error) {
	sql = `update table %s set %s where %s`

	return
}
