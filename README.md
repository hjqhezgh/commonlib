commonlib
======

string
======
//从指定位置开始截取固定长度<br/>
func Substr(str string, start, length int) string {}

//截取索引号之间的字符内容<br/>
func SubstrByStEd(str string, start, end int) string {}

file
======
//判断文件夹是否存在<br/>
func IsDirExists(path string) bool {}

sql
======
//数据库列填充<br/>
func PutRecord(rows *sql.Rows, args ...interface{}) error {}

web
======
//html模板合成方法<br/>
RenderTemplate(w http.ResponseWriter, r *http.Request, templateName string, data map[string]interface{}, funcMap template.FuncMap, templateFiles ...string) {}

//JSON数据输出，需要指定换行符号的替换符号<br/>
func OutputJson(w http.ResponseWriter, object interface{},newlineReplace string) {}