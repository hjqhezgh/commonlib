// Title：文件工具类
//
// Description:
//
// Author:black
//
// Createtime:2013-07-30 16:31
//
// Version:1.0
//
// 修改历史:版本号 修改日期 修改人 修改说明
//
// 1.0 2013-07-30 16:31 black 创建文档
package tools

import (
	"os"
)

//判断文件夹是否存在
func IsDirExists(path string) bool {

	fi, err := os.Stat(path)

	if err != nil {
		return os.IsExist(err)
	} else {
		return fi.IsDir()
	}
}
