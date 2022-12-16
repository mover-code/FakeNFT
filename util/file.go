/***************************
@File        : file.go
@Time        : 2022/02/18 17:40:47
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : 文件操作
****************************/

package util

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Mkdir 创建目录
func Mkdir(name string) error {
	if !Exists(name) {
		return os.MkdirAll(name, os.ModePerm)
	}
	return nil
}

// WriteJson 写入文件
func Write(v interface{}, path string) (err error) {
	data, _ := json.Marshal(&v)

	// 将 JSON 格式数据写入当前目录下的文件（文件不存在会自动创建）
	return ioutil.WriteFile(path, data, 0644)
}

// GetChildFile 获取子目录
func GetChildFile(path string) ([]string, error) {
	childs := []string{}
	fileList, err := ioutil.ReadDir(path)

	if err == nil {
		for _, f := range fileList {
			childs = append(childs, f.Name())
		}
	}
	return childs, err
}

// SplitFileName 分割文件名
func SplitFileName(f string) []string {
	return strings.Split(f, ".")
}

// FindNum 从字符串中匹配第一个数字
func FindNum(s string) int {
	re, _ := regexp.Compile(`\d+`)

	//查找符合正则的第一个
	all := re.FindAll([]byte(s), -1)
	res, _ := strconv.Atoi(string(all[0]))
	if res <= 0 {
		res = 0
	}
	return res
}

func FindLevel(s string) string {
	if r := strings.Split(s, "&"); len(r) > 1 {
		return strings.Split(r[1], "-")[0]
	}
	return "-"
}

func FindLimit(s string) int {
	if r := strings.Split(s, "-"); len(r) > 1 {
		re, _ := regexp.Compile(`\d+`)

		//查找符合正则的第一个
		all := re.FindAll([]byte(r[1]), -1)
		res, _ := strconv.Atoi(string(all[0]))
		if res <= 0 {
			res = 0
		}
		return res
	}
	return 0
}
