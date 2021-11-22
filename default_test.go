package mysqlplus_test

import (
	"fmt"
	mysqlplus "mysql_plus"
	"strings"
	"testing"
	"time"

	"github.com/go-basic/uuid"
)

// DictInfo tb_dict_info 字典信息表
type DictInfo struct {
	Id           string `json:"id" db:"id"`                       // 字典id
	DictName     string `json:"dictName" db:"dict_name"`          // 字典名称
	DictKey      string `json:"dictKey" db:"dict_key"`            // 字典键
	DictValue    string `json:"dictValue" db:"dict_value"`        // 字典值
	DictBeforeId string `json:"dictBeforeId" db:"dict_before_id"` // 上级字典id
	CreateTime   int64  `json:"createTime" db:"create_time"`      // 创建时间
	UpdateTime   int64  `json:"updateTime" db:"update_time"`      // 修改时间
	CreateUser   string `json:"createUser" db:"create_user"`      // 创建者
	Status       int    `json:"status" db:"status"`               // 字典状态 1 正常 2 停用 3 删除
}

func TestInsert(t *testing.T) {
	sql, params := mysqlplus.Query.Insert("tb_dict_info", DictInfo{
		Id:           strings.ToUpper(strings.ReplaceAll(uuid.New(), "-", "")),
		DictName:     "DictName",
		DictKey:      "DictKey",
		DictValue:    "DictValue",
		DictBeforeId: "DictBeforeId",
		CreateTime:   time.Now().Unix(),
		UpdateTime:   0,
		CreateUser:   "CreateUser",
		Status:       1,
	})

	fmt.Println(sql)
	fmt.Println(params...)
}

func TestUpdate(t *testing.T) {
	sql, params := mysqlplus.Query.Update("tb_dict_info", DictInfo{
		// Id:           strings.ToUpper(strings.ReplaceAll(uuid.New(), "-", "")),
		DictName: "DictName",
		// DictKey:      "DictKey",
		DictValue:    "DictValue",
		DictBeforeId: "DictBeforeId",
		CreateTime:   time.Now().Unix(),
		UpdateTime:   0,
		CreateUser:   "CreateUser",
		Status:       1,
	}, "status = ? and ? > create_time", 1, time.Now().Unix())

	fmt.Println(sql)
	fmt.Println(params...)
}

func TestDelete(t *testing.T) {
	sql, params := mysqlplus.Query.Delete("tb_dict_info", "status = ? and ? > create_time", 1, time.Now().Unix())

	fmt.Println(sql)
	fmt.Println(params...)
}

func TestSelect(t *testing.T) {
	sql, params := mysqlplus.Query.Select("tb_dict_info", "status = ? and ? > create_time", DictInfo{}, 1, time.Now().Unix())

	fmt.Println(sql)
	fmt.Println(params...)
}

func TestDBSelect(t *testing.T) {
	mp, err := mysqlplus.MySqlPlus.Create("root:123456@tcp(localhost:3306)/mw?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("create error: %v", err)
		return
	}
	var res []DictInfo
	if err = mp.Select(&res, DictInfo{}, "tb_dict_info", "id=?", "06E67B7741E44874B784409B55DD4108"); err != nil {
		fmt.Printf("select error: %v", err)
		return
	}
	for i := range res {
		fmt.Println(res[i].Id)
	}
}

func TestDBUpdate(t *testing.T) {
	mp, err := mysqlplus.MySqlPlus.Create("root:123456@tcp(localhost:3306)/mw?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("create error: %v", err)
		return
	}
	if _, err = mp.Update("tb_dict_info", DictInfo{
		DictName: "hello",
	}, "id=?", "06E67B7741E44874B784409B55DD4108"); err != nil {
		fmt.Printf("update error: %v", err)
		return
	}
}

func TestDBDelete(t *testing.T) {
	mp, err := mysqlplus.MySqlPlus.Create("root:123456@tcp(localhost:3306)/mw?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("create error: %v", err)
		return
	}
	if _, err = mp.Delete("tb_dict_info", "id=?", "06E67B7741E44874B784409B55DD4108"); err != nil {
		fmt.Printf("delete error: %v", err)
		return
	}
}

func TestDBInsert(t *testing.T) {
	mp, err := mysqlplus.MySqlPlus.Create("root:123456@tcp(localhost:3306)/mw?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("create error: %v", err)
		return
	}
	if _, err = mp.Insert("tb_dict_info", DictInfo{
		Id:           mysqlplus.Util.UUIDReplaceAll("-", ""),
		DictName:     "DictName",
		DictKey:      "DictKey",
		DictValue:    "DictValue",
		DictBeforeId: "DictBeforeId",
		CreateTime:   time.Now().Unix(),
		UpdateTime:   0,
		CreateUser:   "CreateUser",
		Status:       1,
	}); err != nil {
		fmt.Printf("insert error: %v", err)
		return
	}
}
