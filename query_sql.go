package mysqlplus

import (
	"fmt"
	"strings"
)

type querySQLInterface interface {
}

type QuerySQL struct {
	table       string        // 表名
	model       interface{}   // 内核主表模型
	params      []interface{} // 实际内核参数
	sql         string        // 最终sql
	where       string        // where条件
	whereParams []interface{} // where条件参数数组
	set         string        // set语句
	setParams   []interface{} // set对应参数数组
}

func (q *QuerySQL) Select() {
	c := Util.GetAllTags(q.model, "db")
	sql := `SELECT %s FROM %s`
	q.sql = fmt.Sprintf(sql, strings.Join(c, ","), q.table)
}

func (q *QuerySQL) Update() {
	sql := `UPDATE %s`
	q.sql = fmt.Sprintf(sql, q.table)
}

func (q *QuerySQL) Delete() {
	sql := `DELETE FROM %s`
	q.sql = fmt.Sprintf(sql, q.table)
}

func (q *QuerySQL) Where(where string, params ...interface{}) {
	q.where = where
	q.whereParams = params
}

func (q *QuerySQL) Set(set string, params ...interface{}) {
	q.set = set
	q.setParams = params
}

func (q *QuerySQL) ToSQL() (string, []interface{}) {
	result := q.sql
	params := make([]interface{}, 0)
	if q.set != "" {
		result = result + " SET " + q.set
		params = append(params, q.setParams...)
	}
	if q.where != "" {
		result = result + " WHERE " + q.where
		params = append(params, q.whereParams...)
	}
	return result, params
}
