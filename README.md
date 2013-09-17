gotool
======

//从指定位置开始截取固定长度<br/>
func Substr(str string, start, length int) string {}

//截取索引号之间的字符内容<br/>
func SubstrByStEd(str string, start, end int) string {}

//判断文件夹是否存在<br/>
func IsDirExists(path string) bool {}

//数据库列填充<br/>
func PutRecord(rows *sql.Rows, args ...interface{}) error {}
