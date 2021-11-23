package mysqlplus

import (
	"fmt"
	"strings"
)

type querySQLInterface interface {
}

type QuerySQL struct {
	table  string        // 表名
	model  interface{}   // 内核主表模型
	params []interface{} // 实际内核参数
	sql    string        // 最终sql
}

func (q *QuerySQL) Select() {
	c := Util.GetAllTags(q.model, "db")
	sql := `select %s from %s`
	q.sql = fmt.Sprintf(sql, strings.Join(c, ","), q.table)
}
