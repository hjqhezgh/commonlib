// Title：
//
// Description:
//
// Author:black
//
// Createtime:2014-01-13 10:22
//
// Version:1.0
//
// 修改历史:版本号 修改日期 修改人 修改说明
//
// 1.0 2014-01-13 10:22 black 创建文档
package commonlib

import (
	"database/sql"
	"github.com/axgle/mahonia"
)

func ExportCsv(columnNames []string, sql string ,params []interface {},db *sql.DB) ([]byte ,error){

	type TmpString struct {
		Value string
	}

	enc:=mahonia.NewEncoder("gbk")

	rows, err := db.Query(sql, params...)

	if err != nil {
		return []byte{} , err
	}

	var buf = []byte{}

	for  index,name := range columnNames{
		buf = append(buf,[]byte(enc.ConvertString(name))...)
		if index !=len(columnNames)-1{
			buf = append(buf,[]byte(";")...)
		}
	}

	buf = append(buf,[]byte("\n")...)

	columns, err := rows.Columns()

	for rows.Next() {
		objects := []*TmpString{}
		tmpString := []interface{}{}

		for i := 0; i < len(columns); i++ {
			var tmp = new(TmpString)
			objects = append(objects, tmp)
			tmpString = append(tmpString, &tmp.Value)
		}

		err = PutRecord(rows, tmpString...)

		if err != nil {
			return nil, err
		}

		for index,object := range objects{
			buf = append(buf,[]byte(enc.ConvertString(object.Value))...)
			if index !=len(objects)-1{
				buf = append(buf,[]byte(";")...)
			}
		}

		buf = append(buf,[]byte("\n")...)
	}


	return buf,nil
}

