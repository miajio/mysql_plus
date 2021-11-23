package mysqlplus

import (
	"fmt"
	"strings"
)

// queryInterface 模型处理中心
type queryInterface interface {
	Insert(table string, val interface{}) (string, []interface{})                                           // Insert 自动生成Insert语句方法
	Update(table string, set interface{}, where string, whereParams ...interface{}) (string, []interface{}) // Update 自动生成Update语句方法
	Delete(table, where string, whereParams ...interface{}) (string, []interface{})                         // Delete 自动生成Delete语句方法
	Select(table, where string, model interface{}, whereParams ...interface{}) (string, []interface{})      // Select 自动生成Select语句方法
	CreateSQL(model interface{}) *QuerySQL                                                                  // CreateSQL 创建QuerySQL查询对象
}

// queryStruct 模型处理中心结构体
type queryStruct struct {
}

var Query queryInterface = (*queryStruct)(nil)

// Insert 自动生成Insert语句方法
func (mc *queryStruct) Insert(table string, val interface{}) (string, []interface{}) {
	sql := `insert into %s (%s) values (%s)`

	wheres := make([]string, 0)

	columns, params := Util.GetColumns(val, "db")
	for i := 0; i < len(columns); i++ {
		wheres = append(wheres, "?")
	}

	sql = fmt.Sprintf(sql, table, strings.Join(columns, ","), strings.Join(wheres, ","))
	return sql, params
}

// Update 自动生成Update语句方法
func (mc *queryStruct) Update(table string, set interface{}, where string, whereParams ...interface{}) (string, []interface{}) {
	sql := `update %s set %s `
	columns, resultParams := Util.GetColumns(set, "db")
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
func (mc *queryStruct) Delete(table string, where string, whereParams ...interface{}) (string, []interface{}) {
	sql := `delete from %s`
	if where != "" {
		sql = sql + " where %s"
		return fmt.Sprintf(sql, table, where), whereParams
	}

	return fmt.Sprintf(sql, table), whereParams
}

// Select 自动生成Select语句
func (mc *queryStruct) Select(table, where string, model interface{}, whereParams ...interface{}) (string, []interface{}) {
	sql := `select %s from %s`
	columns := strings.Join(Util.GetAllTags(model, "db"), ",")
	if where != "" {
		sql = sql + " where %s"
		return fmt.Sprintf(sql, columns, table, where), whereParams
	}

	return fmt.Sprintf(sql, columns, table), whereParams
}

// CreateSQL 创建QuerySQL查询对象
func (mc *queryStruct) CreateSQL(model interface{}) *QuerySQL {
	return &QuerySQL{
		model: model,
	}
}
