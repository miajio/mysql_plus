package mysqlplus

import (
	"fmt"
	"strings"
)

type querySQLInterface interface {
	Select(queryColumns string) *querySQL                    // Select 创建查询方案
	SelectJoin(queryColumns ...string) *querySQL             // SelectJoin 创建关联关系的查询方案
	join(joinType int, joinOn map[string]SqlParam) *querySQL // Join 关联查询逻辑
	Join(joinOn map[string]SqlParam) *querySQL               // Join cross join
	InnerJoin(joinOn map[string]SqlParam) *querySQL          // InnerJoin inner join
	LeftJoin(joinOn map[string]SqlParam) *querySQL           // LeftJoin left join
	RightJoin(joinOn map[string]SqlParam) *querySQL          // RightJoin right join
	Update() *querySQL                                       // Update 创建修改方案
	Delete() *querySQL                                       // Delete 创建删除方案
	Where(where string, params ...interface{}) *querySQL     // Where where条件语句
	Set(set string, params ...interface{}) *querySQL         // Set set语句
	ToSQL() (string, []interface{})                          // ToSQL 转sql语句
}

// querySQL 生成SQL语句的核心类
type querySQL struct {
	table        string              // 表名
	as           string              // 别名
	model        interface{}         // 内核主表模型
	params       []interface{}       // 实际内核参数
	sql          string              // 最终sql
	where        string              // where条件
	whereParams  []interface{}       // where条件参数数组
	set          string              // set语句
	setParams    []interface{}       // set对应参数数组
	group        int                 // 类型 1 查询 2 修改 3 删除 4 表关联模式查询
	queryColumns string              // 查询列数据
	joinSql      string              // 表关联逻辑sql
	joinOn       map[string]SqlParam // 表关联逻辑模式 on key 别名 value 表关联关系sql
	joinParams   []interface{}       // 表关联逻辑模式 参数数组
	resultModel  interface{}         // 返回数据模型
}

type SqlParam struct {
	As     string        // as 别名
	On     string        // sql语句
	Params []interface{} // 对应参数数组
}

// CreateQuerySQL 创建QuerySQL查询对象
func CreateQuerySQL(table string, as string, model interface{}) *querySQL {
	return &querySQL{
		table: table,
		as:    as,
		model: model,
	}
}

// Select 创建查询方案
func (q *querySQL) Select(queryColumns string) *querySQL {
	c := Util.GetAllTags(q.model, "db")
	sql := `SELECT %s FROM %s`
	if queryColumns != "" {
		q.sql = fmt.Sprintf(sql, queryColumns, q.table)
	} else {
		q.sql = fmt.Sprintf(sql, strings.Join(c, ","), q.table)
	}
	if q.as != "" {
		q.sql = q.sql + " as " + q.as
	}
	q.group = 1
	return q
}

// SelectJoin 创建关联关系的查询方案
func (q *querySQL) SelectJoin(queryColumns ...string) *querySQL {
	q = q.Select(strings.Join(queryColumns, ","))
	q.group = 4
	return q
}

// join 关联查询逻辑
// joinType 0 cross join 1 inner join 2 left join 3 right join
func (q *querySQL) join(joinType int, joinOn map[string]SqlParam) *querySQL {
	joinSql := ""

	joinParams := make([]interface{}, 0)

	switch joinType {
	case 0:
		joinSql = " JOIN "
	case 1:
		joinSql = " INNER JOIN "
	case 2:
		joinSql = " LEFT JOIN "
	case 3:
		joinSql = " RIGHT JOIN "
	}

	for key := range joinOn {
		sqlParam := joinOn[key]
		if sqlParam.On != "" {
			as := ""
			if sqlParam.As != "" {
				as = " as " + sqlParam.As
			}

			joinSql = joinSql + key + as + " ON " + sqlParam.On
			joinParams = append(joinParams, sqlParam.Params...)
		}
	}
	q.joinSql = joinSql
	q.joinParams = joinParams
	return q
}

// Join cross join
func (q *querySQL) Join(joinOn map[string]SqlParam) *querySQL {
	return q.join(0, joinOn)
}

// InnerJoin inner join
func (q *querySQL) InnerJoin(joinOn map[string]SqlParam) *querySQL {
	return q.join(1, joinOn)
}

// LeftJoin left join
func (q *querySQL) LeftJoin(joinOn map[string]SqlParam) *querySQL {
	return q.join(2, joinOn)
}

// RightJoin right join
func (q *querySQL) RightJoin(joinOn map[string]SqlParam) *querySQL {
	return q.join(3, joinOn)
}

// Update 创建修改方案
func (q *querySQL) Update() *querySQL {
	sql := `UPDATE %s`
	q.sql = fmt.Sprintf(sql, q.table)
	q.group = 2
	return q
}

// Delete 创建删除方案
func (q *querySQL) Delete() *querySQL {
	sql := `DELETE FROM %s`
	q.sql = fmt.Sprintf(sql, q.table)
	q.group = 3
	return q
}

// Where where条件语句
func (q *querySQL) Where(where string, params ...interface{}) *querySQL {
	q.where = where
	q.whereParams = params
	return q
}

// Set set语句
func (q *querySQL) Set(set string, params ...interface{}) *querySQL {
	q.set = set
	q.setParams = params
	return q
}

// ToSQL 转sql语句
func (q *querySQL) ToSQL() (string, []interface{}) {
	result := q.sql
	params := make([]interface{}, 0)
	if q.group == 4 {
		result = result + q.joinSql
		params = append(params, q.joinParams...)
	}
	if q.group == 2 {
		if q.set != "" {
			result = result + " SET " + q.set
			params = append(params, q.setParams...)
		}
	}
	if q.where != "" {
		result = result + " WHERE " + q.where
		params = append(params, q.whereParams...)
	}
	return result, params
}
