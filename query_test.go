package mysqlplus_test

import (
	"fmt"
	mysqlplus "mysql_plus"
	"testing"
)

func TestQuerySQL(t *testing.T) {
	q := mysqlplus.Query.CreateSQL("tb_dict_info", DictInfo{})
	q.Select()
	q.Where("id=? and dict_name like concat('%', ?, '%')", "42087eb009d32bfeaa699e29da379cf4", "es")
	sql, params := q.ToSQL()
	fmt.Println(sql)
	fmt.Println(params...)

	b := mysqlplus.Query.CreateSQL("tb_dict_info", DictInfo{})
	b.Update()
	b.Set("dict_name = ?,dict_value = ?", "name", "val")
	b.Where("id = ?", "42087eb009d32bfeaa699e29da379cf4")
	sql, params = b.ToSQL()

	fmt.Println(sql)
	fmt.Println(params...)
}
