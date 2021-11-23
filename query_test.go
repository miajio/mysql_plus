package mysqlplus_test

import (
	"fmt"
	mysqlplus "mysql_plus"
	"testing"
)

func TestQuerySQL(t *testing.T) {
	// q := mysqlplus.Query.CreateSQL("tb_dict_info", DictInfo{}).Select("").Where("id= ? and dict_name like concat('%', ?, '%')", "42087eb009d32bfeaa699e29da379cf4", "es")
	// sql, params := q.ToSQL()
	// fmt.Println(sql)
	// fmt.Println(params...)

	// b := mysqlplus.Query.CreateSQL("tb_dict_info", DictInfo{}).Update().Set("dict_name = ?,dict_value = ?", "name", "val").Where("id = ?", "42087eb009d32bfeaa699e29da379cf4")
	// sql, params = b.ToSQL()

	// fmt.Println(sql)
	// fmt.Println(params...)

	c := mysqlplus.Query.CreateSQL("tb_dict_info", DictInfo{}).SelectJoin("a.id, b.id, b.name").Join(1, map[string]mysqlplus.SqlParam{
		"a": {
			Sql: "a.id = b.join_id",
		},
	}).Where("a.id = ?", "123456")
	sql, params := c.ToSQL()
	fmt.Println(sql)
	fmt.Println(params...)
}
