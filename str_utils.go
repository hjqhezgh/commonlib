// Title：字符串操作扩展工具类
//
// Description:
//
// Author:Black
//
// Createtime:2013-07-22 16:17
//
// Version:1.0
//
// 修改历史:版本号 修改日期 修改人 修改说明
//1.0 2013-05-23 13:32 Black 创建
package main

//从指定位置开始截取固定长度
func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

//截取索引号之间的字符内容
func SubstrByStEd(str string, start, end int) string {

	rs := []rune(str)

	return string(rs[start:end])
}
