// Title：提供web渲染常用的一些方法
//
// Description:
//
// Author:black
//
// Createtime:2013-09-17 14:37
//
// Version:1.0
//
// 修改历史:版本号 修改日期 修改人 修改说明
//
// 1.0 2013-09-17 14:37 black 创建文档
package commonlib

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
	"text/template"
)

//html模板合成方法
//参考的调用方式 RenderTemplate(w, r, "home.html", m, nil, "../template/home.html", "../template/nav.html")
func RenderTemplate(w http.ResponseWriter, r *http.Request, templateName string, data map[string]interface{}, funcMap template.FuncMap, templateFiles ...string) {

	var err error

	var t *template.Template

	var buf bytes.Buffer

	t = template.New(templateName)

	if funcMap != nil {
		//注册一些常用函数
		t = t.Funcs(funcMap)
	}

	t, err = t.ParseFiles(templateFiles...)

	if err != nil {
		panic(err)
	}

	err = t.Execute(&buf, data)

	if err != nil {
		panic(err)
	}

	w.Write(buf.Bytes())
}

//JSON数据输出，需要指定换行符号的替换符号
func OutputJson(w http.ResponseWriter, object interface{}, newlineReplace string) {
	b, err := json.Marshal(object)

	if err != nil {
		return
	}


	str := string(b)

	str = strings.Replace(str, "\n", newlineReplace, -1)
	str = strings.Replace(str, "\n\r", newlineReplace, -1)
	str = strings.Replace(str, "\r\n", newlineReplace, -1)
	str = strings.Replace(str, "\r", newlineReplace, -1)
	/*
	str = strings.Replace(str, "\"", "\\\"", -1)
	str = strings.Replace(str, "'", "\\'", -1)*/

	w.Write([]byte(str))
}
