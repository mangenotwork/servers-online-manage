package utils

import (
	"regexp"
	"strconv"
)

//删除字符串前后两端的所有空格
func DeletePreAndSufSpace(str string) string {
	strList := []byte(str)
	spaceCount, count := 0, len(strList)
	for i := 0; i <= len(strList)-1; i++ {
		if strList[i] == 32 {
			spaceCount++
		} else {
			break
		}
	}

	strList = strList[spaceCount:]
	spaceCount, count = 0, len(strList)
	for i := count - 1; i >= 0; i-- {
		if strList[i] == 32 {
			spaceCount++
		} else {
			break
		}
	}

	return string(strList[:count-spaceCount])
}

//字符串转flot64
func Str2Flot64(s string) float64 {
	floatnum, err := strconv.ParseFloat(s, 64)
	if err != nil {
		floatnum = 0
	}
	return floatnum
}

//字符串转int64
func Str2Int64(s string) int64 {
	reg := regexp.MustCompile(`[0-9]+`)
	sList := reg.FindAllString(s, -1)
	if len(sList) == 0 {
		return 0
	}

	int64num, err := strconv.ParseInt(sList[0], 10, 64)
	if err != nil {
		return 0
	}
	return int64num
}

//数字类字符串 转 int64
func Num2Int64(s string) int64 {
	int64num, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return int64num
}
