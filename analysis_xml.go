package mysqlplus

import (
	"encoding/xml"
	"io/ioutil"
)

// sqlXml xml文件指定模型结构体
type sqlXml struct {
	XmlName string      `xml:"sqlXml"`    // sqlXml
	Name    string      `xml:"name,attr"` // name=?
	Insert  []insertXml `xml:"insert"`    // insert tag
	Update  []updateXml `xml:"update"`    // update tag
	Delete  []deleteXml `xml:"delete"`    // delete tag
	Select  []selectXml `xml:"select"`    // select tag
}

// insertXml insert tag
type insertXml struct {
	XmlName string `xml:"insert"`
	Name    string `xml:"name,attr"`
	Sql     string `xml:",chardata"`
}

// updateXml update tag
type updateXml struct {
	XmlName string `xml:"update"`
	Name    string `xml:"name,attr"`
	Sql     string `xml:",chardata"`
}

// deleteXml delete tag
type deleteXml struct {
	XmlName string `xml:"delete"`
	Name    string `xml:"name,attr"`
	Sql     string `xml:",chardata"`
}

// selectXml select tag
type selectXml struct {
	XmlName string `xml:"select"`
	Name    string `xml:"name,attr"`
	Sql     string `xml:",chardata"`
}

// AnalysisXml 解析xml
func AnalysisXml(source string) (*sqlXml, error) {
	centext, err := ioutil.ReadFile(source)
	if err != nil {
		return nil, err
	}

	var sqlXml sqlXml
	err = xml.Unmarshal(centext, &sqlXml)
	if err != nil {
		return nil, err
	}
	return &sqlXml, nil
}

// GetInsertMap 获取insert语句Map
func (sqlXml *sqlXml) GetInsertMap() map[string]string {
	result := make(map[string]string)
	for i := range sqlXml.Insert {
		insert := sqlXml.Insert[i]
		result[insert.Name] = insert.Sql
	}
	return result
}

// GetUpdateMap 获取update语句Map
func (sqlXml *sqlXml) GetUpdateMap() map[string]string {
	result := make(map[string]string)
	for i := range sqlXml.Update {
		update := sqlXml.Update[i]
		result[update.Name] = update.Sql
	}
	return result
}

// GetDeleteMap 获取delete语句Map
func (sqlXml *sqlXml) GetDeleteMap() map[string]string {
	result := make(map[string]string)
	for i := range sqlXml.Delete {
		delete := sqlXml.Delete[i]
		result[delete.Name] = delete.Sql
	}
	return result
}

// GetSelectMap 获取select语句Map
func (sqlXml *sqlXml) GetSelectMap() map[string]string {
	result := make(map[string]string)
	for i := range sqlXml.Select {
		selectSql := sqlXml.Select[i]
		result[selectSql.Name] = selectSql.Sql
	}
	return result
}
