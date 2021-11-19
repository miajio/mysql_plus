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
	sql, params, err := mysqlplus.ModelCentent.Insert("tb_dict_info", DictInfo{
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
	if err != nil {
		fmt.Printf("insert sql create fail:%v \n", err)
	}

	fmt.Println(sql)
	fmt.Println(params...)
}
