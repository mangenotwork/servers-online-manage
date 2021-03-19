package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
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

//字符串转int
func Str2Int(s string) int {
	reg := regexp.MustCompile(`[0-9]+`)
	sList := reg.FindAllString(s, -1)
	if len(sList) == 0 {
		return 0
	}

	int64num, err := strconv.Atoi(sList[0])
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

//数字类字符串 转 int
func Num2Int(s string) int {
	innum, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return innum
}

//随机字符
// length 指定长度
func GetRandString(length int) string {
	strBytes := make([]byte, length)
	for i := 0; i < length; i++ {
		strBytes[i] = byte(rand.Intn(26) + 97)
	}
	return string(strBytes)
}

// 字符串Md5
func Str2MD5(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}
