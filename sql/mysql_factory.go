// Title：列赋值工具
//
// Description:数据库连接池
//
// Author:Bill
//
// Createtime:2013-07-24 15:09
//
// Version:1.0
//
// 修改历史:版本号 修改日期 修改人 修改说明
//
// 1.0 2013-07-24 15:09 Bill 创建文档
package sql

import (
	"database/sql"
	"strconv"
)

func PutRecord(rows *sql.Rows, args ...interface{}) error {

	values := make([]sql.RawBytes, len(args))

	scanArgs := make([]interface{}, len(values))

	for i := range values {
		scanArgs[i] = &values[i]
	}

	err := rows.Scan(scanArgs...)

	if err != nil {
		return err
	}

	var value string
	for i, col := range values {
		if col == nil {
			continue
		}
		value = string(col)
		putColData(args[i], value)
	}
	return nil
}

func putColData(arg interface{}, value string) {
	switch vtype := arg.(type) {
	case *string:
		*vtype = value
		break

	case *int:
		*vtype, _ = strconv.Atoi(value)
		break

	case *int64:
		i, _ := strconv.Atoi(value)
		*vtype = int64(i)
		break

	case *float32:
		i, _ := strconv.Atoi(value)
		*vtype = float32(i)
		break
	}
}
