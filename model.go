package mysqlplus

// Model 数据库模型结构体
type Model struct {
	Table     string            // 表名
	Source    string            // 来源
	Core      interface{}       // 内核结构体
	isInit    bool              // 是否初始化
	insertMap map[string]string // insert语句Map
	updateMap map[string]string // update语句Map
	deleteMap map[string]string // delete语句Map
	selectXml map[string]string // select语句Map

}

// init 初始化
func (model *Model) init() error {
	if !model.isInit {
		sqlXml, err := AnalysisXml(model.Source)
		if err != nil {
			return err
		}
		model.isInit = true
		model.insertMap = sqlXml.GetInsertMap()
		model.updateMap = sqlXml.GetUpdateMap()
		model.deleteMap = sqlXml.GetDeleteMap()
		model.selectXml = sqlXml.GetSelectMap()
	}
	return nil
}

// Insert 组装生成insert语句
// param:
// method 方法
// param 参数结构体
// result:
// string 返回的sql
// []interface 对应参数列表
func (model *Model) Insert(method string, param interface{}) (string, []interface{}, error) {
	if err := model.init(); err != nil {
		return "", nil, err
	}
	sql := model.insertMap[method]
	return sql, nil, nil
}
