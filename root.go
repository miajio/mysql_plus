package mysqlplus

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type mySqlPlusInterface interface{}

type mySqlPlusStruct struct {
	isInit bool
	db     *sqlx.DB
}

var MySqlPlus mySqlPlusInterface = (*mySqlPlusStruct)(nil)

// Create 创建数据库连接对象
func (mySqlPlus *mySqlPlusStruct) Create(link string) error {
	db, err := sqlx.Open("mysql", link)
	if err != nil {
		return err
	}
	mySqlPlus.db = db
	mySqlPlus.isInit = true
	return nil
}

func (mySqlPlus *mySqlPlusStruct) SetMaxOpenConns(max int) {
	mySqlPlus.db.SetMaxOpenConns(max)
}

func (mySqlPlus *mySqlPlusStruct) SetMaxIdleConns(max int) {
	mySqlPlus.db.SetMaxIdleConns(max)
}

// func (mySqlPlus *mySqlPlusStruct) SetConnMaxLifetime(max time.Duration) {
// 	mySqlPlus.db.SetConnMaxLifetime(max)
// }
