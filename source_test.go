package mysqlplus_test

import (
	"fmt"
	mysqlplus "mysql_plus"
	"testing"
)

func TestXml(t *testing.T) {
	mo := mysqlplus.Model{
		Source: "./test.xml",
		Table:  "tb_dict_info",
		Core:   DictInfo{},
	}
	sql, param, err := mo.Insert("add", nil)
	if err != nil {
		fmt.Printf("insert sql get fail:%v", err)
		return
	}
	fmt.Println(sql)
	fmt.Println(param...)
}
