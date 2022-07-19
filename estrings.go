package estring

import (
	"strings"
)

/*
*  replaceString,指定替换第几个字符串
*
 */
func ReplaceOneString(s string, old string, news string, pos int) string {

	if !strings.Contains(s, old) {
		return ""
	}
	if pos > strings.Count(s, old) || pos <= 0 {
		return ""
	}
	list := strings.SplitAfter(s, old)
	list[pos-1] = strings.Replace(list[pos-1], old, news, 1)
	return strings.Join(list, "")
}

//根据空格转为list，没有空格就没有
func ToList(s string) []string {
	return strings.Fields(s)
}

//字符串是否在list里面
func IsStringInList(list []string, sub string) bool {
	tmplist := strings.Join(list, " ")
	if strings.Contains(tmplist, sub) {
		return true
	}
	return false
}
