gotool
======

//从指定位置开始截取固定长度
func Substr(str string, start, length int) string {}

//截取索引号之间的字符内容
func SubstrByStEd(str string, start, end int) string {}

//判断文件夹是否存在
func IsDirExists(path string) bool {}

//数据库列填充
func PutRecord(rows *sql.Rows, args ...interface{}) error {}
