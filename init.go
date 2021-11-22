package mysqlplus

import (
	"fmt"
	"reflect"
	"strings"
)

// modelCententInteface 模型处理中心
type modelCententInteface interface {
	getColumns(val interface{}) ([]string, []interface{}) // 获取字段数组及设值数组

	getAllTags(val interface{}, tag string) []string // 获取当前结构体对应标签数据

	isBlank(value reflect.Value) bool // 判断值是否为空

	Insert(table string, val interface{}) (string, []interface{}) // Insert 自动生成Insert语句方法

	Update(table string, set interface{}, where string, whereParams ...interface{}) (string, []interface{}) // Update 自动生成Update语句方法

	Delete(table, where string, whereParams ...interface{}) (string, []interface{}) // Delete 自动生成Delete语句方法

	Select(table, where string, model interface{}, whereParams ...interface{}) (string, []interface{}) // Select 自动生成Select语句方法
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
	sql := `update %s set %s `
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

// Delete 自动生成Delete语句方法
func (mc *modelCententStruct) Delete(table string, where string, whereParams ...interface{}) (string, []interface{}) {
	sql := `delete from %s`
	if where != "" {
		sql = sql + " where %s"
		return fmt.Sprintf(sql, table, where), whereParams
	}

	return fmt.Sprintf(sql, table), whereParams
}

// Select 自动生成Select语句
func (mc *modelCententStruct) Select(table, where string, model interface{}, whereParams ...interface{}) (string, []interface{}) {
	sql := `select %s from %s`
	columns := strings.Join(mc.getAllTags(model, "db"), ",")
	if where != "" {
		sql = sql + " where %s"
		return fmt.Sprintf(sql, columns, table, where), whereParams
	}

	return fmt.Sprintf(sql, columns, table), whereParams
}

// getColumns 基础内核方法 获取参数数据
// 仅返回数据非空的字段数据及数据结果
func (mc *modelCententStruct) getColumns(val interface{}) ([]string, []interface{}) {

	columns := make([]string, 0)
	params := make([]interface{}, 0)

	typeOf := reflect.TypeOf(val)
	valueOf := reflect.ValueOf(val)

	for i := 0; i < typeOf.NumField(); i++ {
		column := typeOf.Field(i).Tag.Get("db")
		value := valueOf.Field(i).Interface()
		if !mc.isBlank(reflect.ValueOf(value)) {
			columns = append(columns, column)
			params = append(params, value)
		}
	}

	return columns, params
}

// getAllTags 基础内核方法 获取参数标签数据
// 获取所有字段数据及数据结果
func (mc *modelCententStruct) getAllTags(val interface{}, tag string) []string {

	columns := make([]string, 0)

	typeOf := reflect.TypeOf(val)

	for i := 0; i < typeOf.NumField(); i++ {
		column := typeOf.Field(i).Tag.Get("db")
		columns = append(columns, column)
	}

	return columns
}

// isBlank 判断值是否为空
func (mc *modelCententStruct) isBlank(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	}
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}
