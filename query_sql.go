package mysqlplus

import (
	"fmt"
	"strings"
)

type querySQLInterface interface {
}

type QuerySQL struct {
	table        string              // 表名
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
	Sql    string        // sql语句
	Params []interface{} // 对应参数数组
}

func (q *QuerySQL) Select(queryColumns string) *QuerySQL {
	c := Util.GetAllTags(q.model, "db")
	sql := `SELECT %s FROM %s`
	if queryColumns != "" {
		q.sql = fmt.Sprintf(sql, queryColumns, q.table)
	} else {
		q.sql = fmt.Sprintf(sql, strings.Join(c, ","), q.table)
	}
	q.group = 1
	return q
}

func (q *QuerySQL) SelectJoin(queryColumns string) *QuerySQL {
	q = q.Select(queryColumns)
	q.group = 4
	return q
}

// Join 关联查询逻辑
// joinType 0 cross join 1 inner join 2 left join 3 right join
func (q *QuerySQL) Join(joinType int, joinOn map[string]SqlParam) *QuerySQL {

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
		if sqlParam.Sql != "" {
			joinSql = joinSql + key + " ON " + sqlParam.Sql
			joinParams = append(joinParams, sqlParam.Params...)
		}
	}
	q.joinSql = joinSql
	q.joinParams = joinParams
	return q
}

func (q *QuerySQL) Update() *QuerySQL {
	sql := `UPDATE %s`
	q.sql = fmt.Sprintf(sql, q.table)
	q.group = 2
	return q
}

func (q *QuerySQL) Delete() *QuerySQL {
	sql := `DELETE FROM %s`
	q.sql = fmt.Sprintf(sql, q.table)
	q.group = 3
	return q
}

func (q *QuerySQL) Where(where string, params ...interface{}) *QuerySQL {
	q.where = where
	q.whereParams = params
	return q
}

func (q *QuerySQL) Set(set string, params ...interface{}) *QuerySQL {
	q.set = set
	q.setParams = params
	return q
}

func (q *QuerySQL) ToSQL() (string, []interface{}) {
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
